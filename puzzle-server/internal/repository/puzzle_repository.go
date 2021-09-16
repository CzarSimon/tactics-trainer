package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/opentracing/opentracing-go"
)

// PuzzleRepository interface for storing and querying for stored puzzles.
type PuzzleRepository interface {
	Save(ctx context.Context, p models.Puzzle) error
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

	_, err := r.db.ExecContext(ctx, saveQuery, p.ID, p.ExternalID, p.FEN, p.EncodeMoves(), p.Rating, p.RatingDeviation, p.Popularity, p.EncodeThemes(), p.GameURL, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}
