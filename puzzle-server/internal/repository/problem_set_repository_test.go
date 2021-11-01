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

var puzzles = []models.Puzzle{
	{
		ID:         "puzzle-0",
		ExternalID: "ext-id-0",
		FEN:        "fen-str",
		Moves:      []string{"e2e4", "e7e5", "g1f3"},
		Rating:     1000,
		Popularity: 100,
		Themes:     []string{"mateIn2", "short", "sacrifice"},
	},
	{
		ID:         "puzzle-1",
		ExternalID: "ext-id-1",
		FEN:        "fen-str",
		Moves:      []string{"e2e4", "e7e5", "g1f3"},
		Rating:     1200,
		Popularity: 100,
		Themes:     []string{"mateIn3", "long"},
	},
	{
		ID:         "puzzle-2",
		ExternalID: "ext-id-2",
		FEN:        "fen-str",
		Moves:      []string{"e2e4", "e7e5", "g1f3"},
		Rating:     1400,
		Popularity: 100,
		Themes:     []string{"passedPawn", "endgame"},
	},
}

func Test_problemSetRepo_Save(t *testing.T) {
	assert := assert.New(t)
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	seedPuzzles(t, db)
	repo := repository.NewProblemSetRepository(db)
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

	var foundID string
	err := db.QueryRow("SELECT id FROM problem_set WHERE id = ?", set.ID).Scan(&foundID)
	assert.Equal(sql.ErrNoRows, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM problem_set_puzzle WHERE problem_set_id = ?", set.ID).Scan(&count)
	assert.NoError(err)
	assert.Equal(0, count)

	err = repo.Save(ctx, set)
	assert.NoError(err)

	err = db.QueryRow("SELECT id FROM problem_set WHERE id = ?", set.ID).Scan(&foundID)
	assert.NoError(err)
	assert.Equal(set.ID, foundID)

	err = db.QueryRow("SELECT COUNT(*) FROM problem_set_puzzle WHERE problem_set_id = ?", set.ID).Scan(&count)
	assert.NoError(err)
	assert.Equal(3, count)

	err = repo.Save(ctx, set)
	assert.Error(err)

	ctx, cancel := context.WithCancel(ctx)
	cancel()
	err = repo.Save(ctx, set)
	assert.True(errors.Is(err, context.Canceled))
}

func seedPuzzles(t *testing.T, db *sql.DB) {
	assert := assert.New(t)
	repo := repository.NewPuzzleRepository(db)
	ctx := context.Background()

	for _, p := range puzzles {
		err := repo.Save(ctx, p)
		assert.NoError(err)
	}
}
