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

func Test_problemSetRepo_Find(t *testing.T) {
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

	_, found, err := repo.Find(ctx, set.ID)
	assert.NoError(err)
	assert.False(found)

	err = repo.Save(ctx, set)
	assert.NoError(err)

	foundSet, found, err := repo.Find(ctx, set.ID)
	assert.NoError(err)
	assert.True(found)
	assert.Equal(set.ID, foundSet.ID)
	assert.Equal(set, foundSet)

	_, found, err = repo.Find(ctx, id.New())
	assert.NoError(err)
	assert.False(found)

	ctx, cancel := context.WithCancel(ctx)
	cancel()
	_, _, err = repo.Find(ctx, set.ID)
	assert.True(errors.Is(err, context.Canceled))
}

func Test_problemSetRepo_FindByUserID(t *testing.T) {
	assert := assert.New(t)
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	seedPuzzles(t, db)
	repo := repository.NewProblemSetRepository(db)
	ctx := context.Background()

	sets := []models.ProblemSet{
		{
			ID:             id.New(),
			Name:           "ps-name",
			Themes:         []string{},
			RatingInterval: "1300 - 1500",
			UserID:         "user-0",
			PuzzleIDs:      []string{"puzzle-0"},
		},
		{
			ID:             id.New(),
			Name:           "ps-name",
			Themes:         []string{},
			RatingInterval: "1300 - 1500",
			UserID:         "user-0",
			PuzzleIDs:      []string{"puzzle-1"},
		},
		{
			ID:             id.New(),
			Name:           "ps-name",
			Themes:         []string{},
			RatingInterval: "1300 - 1500",
			UserID:         "user-1",
			PuzzleIDs:      []string{"puzzle-2"},
		},
	}

	for _, set := range sets {
		err := repo.Save(ctx, set)
		assert.NoError(err)
	}

	found, err := repo.FindByUserID(ctx, "user-0")
	assert.NoError(err)
	assert.Len(found, 2)

	found, err = repo.FindByUserID(ctx, "user-1")
	assert.NoError(err)
	assert.Len(found, 1)
	assert.Equal(sets[2].ID, found[0].ID)
	assert.Len(found[0].PuzzleIDs, 0)

	found, err = repo.FindByUserID(ctx, "other-user")
	assert.NoError(err)
	assert.Len(found, 0)
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
