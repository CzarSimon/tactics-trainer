package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/opentracing/opentracing-go"
)

const (
	separator = ","
)

// PuzzleRepository interface for storing and querying for stored puzzles.
type PuzzleRepository interface {
	Save(ctx context.Context, p models.Puzzle) error
	Find(ctx context.Context, id string) (models.Puzzle, bool, error)
	FindByFilter(ctx context.Context, f models.PuzzleFilter) ([]models.Puzzle, error)
}

func NewPuzzleRepository(db *sql.DB) PuzzleRepository {
	return &puzzleRepo{
		db: db,
	}
}

type puzzleRepo struct {
	db *sql.DB
}

const saveQuery = `
	INSERT INTO puzzle(id, external_id, fen, moves, rating, rating_deviation, popularity, themes, game_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

// Save stores a puzzle in an sql database.
func (r *puzzleRepo) Save(ctx context.Context, p models.Puzzle) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "puzzle_repo_save")
	defer span.Finish()

	_, err := r.db.ExecContext(ctx, saveQuery, p.ID, p.ExternalID, p.FEN, encodeMoves(p.Moves), p.Rating, p.RatingDeviation, p.Popularity, encodeThemes(p.Themes), p.GameURL, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}

const findQuery = `
	SELECT
		id,
		external_id,
		fen, 
		moves,
		rating, 
		rating_deviation, 
		popularity,
		themes,  
		game_url, 
		created_at, 
		updated_at
	FROM 
		puzzle
	WHERE
		id = ?`

// Find looks up puzzle in a SQL database by its id
func (r *puzzleRepo) Find(ctx context.Context, id string) (models.Puzzle, bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "puzzle_repo_find")
	defer span.Finish()

	var p models.Puzzle
	var moveStr string
	var themeStr string
	err := r.db.QueryRowContext(ctx, findQuery, id).Scan(
		&p.ID,
		&p.ExternalID,
		&p.FEN,
		&moveStr,
		&p.Rating,
		&p.RatingDeviation,
		&p.Popularity,
		&themeStr,
		&p.GameURL,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return models.Puzzle{}, false, nil
	}
	if err != nil {
		return models.Puzzle{}, false, fmt.Errorf("failed to query for puzzle with id = %s. Error: %w", id, err)
	}

	p.Moves = decodeMoves(moveStr)
	p.Themes = decodeThemes(themeStr)
	return p, true, nil
}

const findPuzzlesByFilterQuery = `
	SELECT
		id,
		external_id,
		fen, 
		moves,
		rating, 
		rating_deviation, 
		popularity,
		themes,  
		game_url, 
		created_at, 
		updated_at
	FROM 
		puzzle
	WHERE
		rating >= ?
		AND rating <= ?
		AND popularity >= ?`

func (r *puzzleRepo) FindByFilter(ctx context.Context, f models.PuzzleFilter) ([]models.Puzzle, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "puzzle_repo_find_by_filter")
	defer span.Finish()

	query, values := createThemeFilterQuery(f)
	rows, err := r.db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, fmt.Errorf("failed to query puzzles by filter %s: %w", f, err)
	}
	defer rows.Close()

	puzzles := make([]models.Puzzle, 0)
	var p models.Puzzle
	var moveStr string
	var themeStr string
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.ExternalID, &p.FEN, &moveStr, &p.Rating, &p.RatingDeviation, &p.Popularity, &themeStr, &p.GameURL, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row. Error: %w, %s", err, f)
		}

		p.Moves = decodeMoves(moveStr)
		p.Themes = decodeThemes(themeStr)
		puzzles = append(puzzles, p)
	}

	return puzzles, nil
}

func createThemeFilterQuery(f models.PuzzleFilter) (string, []interface{}) {
	conditions := make([]string, len(f.Themes))
	filterValues := make([]interface{}, 3+len(f.Themes))
	filterValues[0] = f.MinRating
	filterValues[1] = f.MaxRating
	filterValues[2] = f.MinPopularity

	for i, theme := range f.Themes {
		conditions[i] = "themes LIKE ?"
		filterValues[3+i] = "%" + encodeTheme(theme) + "%"
	}

	filterValues = append(filterValues, f.Size)

	var themeCondition string
	if len(f.Themes) > 0 {
		themeCondition = fmt.Sprintf("AND %s", strings.Join(conditions, " AND "))
	}

	query := findPuzzlesByFilterQuery + themeCondition + " LIMIT ?"
	return query, filterValues
}

func encodeMoves(moves []string) string {
	return strings.Join(moves, separator)
}

func decodeMoves(s string) []string {
	return strings.Split(s, separator)
}

func encodeThemes(themes []string) string {
	encoded := make([]string, len(themes))
	for i, theme := range themes {
		encoded[i] = encodeTheme(theme)
	}
	return strings.Join(encoded, separator)
}

func encodeTheme(theme string) string {
	return fmt.Sprintf("[%s]", theme)
}

func decodeThemes(s string) []string {
	cleanStr := strings.ReplaceAll(strings.ReplaceAll(s, "[", ""), "]", "")
	return strings.Split(cleanStr, separator)
}
