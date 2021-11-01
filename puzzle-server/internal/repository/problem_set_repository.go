package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/opentracing/opentracing-go"
)

// PuzzleRepository interface for storing and querying for stored puzzles.
type ProblemSetRepository interface {
	Save(ctx context.Context, p models.ProblemSet) error
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
	INSERT INTO problem_set(id, name, description, themes, rating_interval, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

func saveProblemSet(ctx context.Context, tx *sql.Tx, p models.ProblemSet) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "save_problem_set")
	defer span.Finish()

	_, err := tx.ExecContext(ctx, saveProblemSetQuery, p.ID, p.Name, p.Description, encodeThemes(p.Themes), p.RatingInterval, p.UserID, p.CreatedAt, p.UpdatedAt)
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
		return fmt.Errorf("failed to %s: %w", p, err)
	}

	return nil
}
