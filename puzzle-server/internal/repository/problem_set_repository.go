package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/opentracing/opentracing-go"
)

// PuzzleRepository interface for storing and querying for stored puzzles.
type ProblemSetRepository interface {
	Save(ctx context.Context, p models.ProblemSet) error
	Find(ctx context.Context, id string) (models.ProblemSet, bool, error)
	FindByUserID(ctx context.Context, userID string) ([]models.ProblemSet, error)
	Update(ctx context.Context, p models.ProblemSet) error
}

func NewProblemSetRepository(db *sql.DB) ProblemSetRepository {
	return &problemSetRepo{
		db: db,
	}
}

type problemSetRepo struct {
	db *sql.DB
}

func (r *problemSetRepo) Save(ctx context.Context, p models.ProblemSet) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_repo_save")
	defer span.Finish()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = saveProblemSet(ctx, tx, p)
	if err != nil {
		dbutil.Rollback(tx)
		return err
	}

	for i := range p.PuzzleIDs {
		psp := p.ProblemSetPuzzle(i)
		err = saveProblemSetPuzzle(ctx, tx, psp)
		if err != nil {
			dbutil.Rollback(tx)
			return err
		}
	}

	return tx.Commit()
}

const saveProblemSetQuery = `
	INSERT INTO problem_set(id, name, description, themes, rating_interval, user_id, archived, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

func saveProblemSet(ctx context.Context, tx *sql.Tx, p models.ProblemSet) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "save_problem_set")
	defer span.Finish()

	_, err := tx.ExecContext(ctx, saveProblemSetQuery, p.ID, p.Name, p.Description, encodeThemes(p.Themes), p.RatingInterval, p.UserID, p.Archived, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}

const saveProblemSetPuzzleQuery = `
	INSERT INTO problem_set_puzzle(id, puzzle_id, problem_set_id, number) VALUES (?, ?, ?, ?)`

func saveProblemSetPuzzle(ctx context.Context, tx *sql.Tx, p models.ProblemSetPuzzle) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "save_problem_set_puzzle")
	defer span.Finish()

	_, err := tx.ExecContext(ctx, saveProblemSetPuzzleQuery, p.ID, p.PuzzleID, p.ProblemSetID, p.Number)
	if err != nil {
		return fmt.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}

func (r *problemSetRepo) Find(ctx context.Context, id string) (models.ProblemSet, bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_repo_find")
	defer span.Finish()

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true, Isolation: sql.LevelReadCommitted})
	if err != nil {
		return models.ProblemSet{}, false, err
	}
	defer dbutil.Rollback(tx)

	set, found, err := findProblemSet(ctx, tx, id)
	if !found {
		return models.ProblemSet{}, found, err
	}

	puzzleIDs, err := findProblemSetPuzzleIDs(ctx, tx, id)
	if err != nil {
		return models.ProblemSet{}, false, err
	}

	set.PuzzleIDs = puzzleIDs
	return set, true, nil
}

const findProblemSetQuery = `
	SELECT id, name, description, themes, rating_interval, user_id, archived, created_at, updated_at FROM problem_set WHERE id = ?`

func findProblemSet(ctx context.Context, tx *sql.Tx, id string) (models.ProblemSet, bool, error) {
	var p models.ProblemSet
	var themeStr string
	var archived sql.NullBool
	err := tx.QueryRowContext(ctx, findProblemSetQuery, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&themeStr,
		&p.RatingInterval,
		&p.UserID,
		&archived,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return models.ProblemSet{}, false, nil
	} else if err != nil {
		return models.ProblemSet{}, false, fmt.Errorf("failed to query ProblemSet(id=%s): %w", id, err)
	}

	p.Themes = decodeThemes(themeStr)
	p.Archived = archived.Bool
	return p, true, nil
}

const findProblemSetPuzzleIDsQuery = `SELECT puzzle_id FROM problem_set_puzzle WHERE problem_set_id = ? ORDER BY number ASC`

func findProblemSetPuzzleIDs(ctx context.Context, tx *sql.Tx, problemSetId string) ([]string, error) {
	rows, err := tx.QueryContext(ctx, findProblemSetPuzzleIDsQuery, problemSetId)
	if err != nil {
		return nil, fmt.Errorf("failed to query problem_set_puzzle with problem_set_id=%s: %w", problemSetId, err)
	}
	defer rows.Close()

	puzzleIDs := make([]string, 0)
	var puzzleID string
	for rows.Next() {
		err := rows.Scan(&puzzleID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan problem_set_puzzle row. Error: %w, problem_set_id=%s", err, problemSetId)
		}

		puzzleIDs = append(puzzleIDs, puzzleID)
	}

	return puzzleIDs, nil
}

const findProblemSetsByUserIDQuery = `
	SELECT id, name, description, themes, rating_interval, user_id, archived, created_at, updated_at FROM problem_set WHERE user_id = ? ORDER BY created_at ASC`

func (r *problemSetRepo) FindByUserID(ctx context.Context, userID string) ([]models.ProblemSet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_repo_find_by_user_id")
	defer span.Finish()

	rows, err := r.db.QueryContext(ctx, findProblemSetsByUserIDQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query problem_set by user_id=%s: %w", userID, err)
	}
	defer rows.Close()

	sets := make([]models.ProblemSet, 0)
	var s models.ProblemSet
	var themeStr string
	var archived sql.NullBool
	for rows.Next() {
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Description,
			&themeStr,
			&s.RatingInterval,
			&s.UserID,
			&archived,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan problem_set row. Error: %w, user_id=%s", err, userID)
		}

		s.Themes = decodeThemes(themeStr)
		s.PuzzleIDs = make([]string, 0)
		s.Archived = archived.Bool
		sets = append(sets, s)
	}

	return sets, nil
}

const updateProblemSetQuery = `
	UPDATE
		problem_set
	SET    
		name = ?, 
		description = ?, 
		themes = ?,
		rating_interval = ?,
		archived = ?,
		updated_at = ?
	WHERE 
		id = ?
`

func (r *problemSetRepo) Update(ctx context.Context, p models.ProblemSet) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_repo_update")
	defer span.Finish()

	themeStr := encodeThemes(p.Themes)
	_, err := r.db.ExecContext(ctx, updateProblemSetQuery, p.Name, p.Description, themeStr, p.RatingInterval, p.Archived, timeutil.Now(), p.ID)
	if err != nil {
		return fmt.Errorf("failed to update %s: %w", p, err)
	}

	return nil
}
