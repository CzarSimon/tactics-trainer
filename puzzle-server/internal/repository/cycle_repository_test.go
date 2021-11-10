package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/CzarSimon/httputil/id"
	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/stretchr/testify/assert"
)

func Test_cycleRepo_Save(t *testing.T) {
	assert := assert.New(t)
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	seedPuzzles(t, db)
	psRepo := repository.NewProblemSetRepository(db)
	cycleRepo := repository.NewCycleRepository(db)
	ctx := context.Background()

	set := models.ProblemSet{
		ID:             id.New(),
		Name:           "ps-name",
		Themes:         []string{"passedPawn", "endgame"},
		RatingInterval: "1300 - 1500",
		CreatedAt:      timeutil.Now(),
		UpdatedAt:      timeutil.Now(),
		UserID:         "user-0",
		PuzzleIDs: []string{
			"puzzle-0",
			"puzzle-1",
			"puzzle-2",
		},
	}

	err := psRepo.Save(ctx, set)
	assert.NoError(err)

	cycle := models.Cycle{
		ID:              id.New(),
		Number:          1,
		ProblemSetID:    set.ID,
		CurrentPuzzleID: "puzzle-0",
		CreatedAt:       timeutil.Now(),
		UpdatedAt:       timeutil.Now(),
	}

	var foundID string
	err = db.QueryRow("SELECT id FROM cycle WHERE id = ?", cycle.ID).Scan(&foundID)
	assert.Equal(sql.ErrNoRows, err)

	err = cycleRepo.Save(ctx, cycle)
	assert.NoError(err)

	err = db.QueryRow("SELECT id FROM cycle WHERE id = ?", cycle.ID).Scan(&foundID)
	assert.NoError(err)
	assert.Equal(cycle.ID, foundID)

	cycle.ID = id.New()
	cycle.Number = 2
	cycle.ProblemSetID = "missing-id"
	err = cycleRepo.Save(ctx, cycle)
	assert.Error(err)

	cycle.ProblemSetID = set.ID
	cycle.CurrentPuzzleID = "missing-id"
	err = cycleRepo.Save(ctx, cycle)
	assert.Error(err)

	cycle.CurrentPuzzleID = "puzzle-1"
	cycle.Number = 1
	err = cycleRepo.Save(ctx, cycle)
	assert.Error(err)

	ctx, cancel := context.WithCancel(ctx)
	cancel()
	err = cycleRepo.Save(ctx, cycle)
	assert.True(errors.Is(err, context.Canceled))
}
