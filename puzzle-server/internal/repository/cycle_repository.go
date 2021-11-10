package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/opentracing/opentracing-go"
)

// CycleRepository interface for storing and querying for stored cycles.
type CycleRepository interface {
	Save(ctx context.Context, c models.Cycle) error
	Find(ctx context.Context, id string) (models.Cycle, bool, error)
}

func NewCycleRepository(db *sql.DB) CycleRepository {
	return &cycleRepo{
		db: db,
	}
}

type cycleRepo struct {
	db *sql.DB
}

const saveCycleRepo = `INSERT INTO cycle(id, number, problem_set_id, current_puzzle_id, compleated_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

func (r *cycleRepo) Save(ctx context.Context, c models.Cycle) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cycle_repo_save")
	defer span.Finish()

	compleatedAt := sql.NullTime{
		Time:  c.CompleatedAt,
		Valid: !c.CompleatedAt.Equal(time.Time{}),
	}
	_, err := r.db.ExecContext(ctx, saveCycleRepo, c.ID, c.Number, c.ProblemSetID, c.CurrentPuzzleID, compleatedAt, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save %s: %w", c, err)
	}

	return nil
}

const findCycleQuery = `
	SELECT 
		id, 
		number, 
		problem_set_id, 
		current_puzzle_id, 
		compleated_at, 
		created_at, 
		updated_at
	FROM
		cycle
	WHERE 
		id = ?
`

func (r *cycleRepo) Find(ctx context.Context, id string) (models.Cycle, bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cycle_repo_find")
	defer span.Finish()

	var c models.Cycle
	var compleatedAt sql.NullTime
	err := r.db.QueryRowContext(ctx, findCycleQuery, id).Scan(
		&c.ID,
		&c.Number,
		&c.ProblemSetID,
		&c.CurrentPuzzleID,
		&compleatedAt,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return models.Cycle{}, false, nil
	} else if err != nil {
		return models.Cycle{}, false, fmt.Errorf("failed to query Cycle(id=%s): %w", id, err)
	}

	c.CompleatedAt = compleatedAt.Time
	return c, true, nil
}
