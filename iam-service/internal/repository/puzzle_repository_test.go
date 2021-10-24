package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_puzzleRepo_Save(t *testing.T) {
	assert := assert.New(t)
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	repo := repository.NewPuzzleRepository(db)
	ctx := context.Background()

	puzzle := models.Puzzle{
		ID:         "puzzle-0",
		ExternalID: "ext-id-0",
		FEN:        "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		Moves: []string{
			"e2e4",
			"e7e5",
			"g1f3",
		},
		Rating:          1500,
		RatingDeviation: 100,
		Popularity:      10,
		Themes: []string{
			"mateIn2",
			"short",
			"sacrifice",
		},
		GameURL:   "https://some.url/some/id",
		CreatedAt: timeutil.Now(),
		UpdatedAt: timeutil.Now(),
	}

	var foundID string
	err := db.QueryRow("SELECT id FROM puzzle WHERE id = ?", puzzle.ID).Scan(&foundID)
	assert.Equal(sql.ErrNoRows, err)

	err = repo.Save(ctx, puzzle)
	assert.NoError(err)

	var moves string
	var themes string
	err = db.QueryRow("SELECT id, moves, themes FROM puzzle WHERE id = ?", puzzle.ID).Scan(&foundID, &moves, &themes)
	assert.NoError(err)
	assert.Equal("puzzle-0", foundID)
	assert.Equal("e2e4,e7e5,g1f3", moves)
	assert.Equal("[mateIn2],[short],[sacrifice]", themes)

	ctx, cancel := context.WithCancel(ctx)
	cancel()
	puzzle.ID = "other-id"

	err = repo.Save(ctx, puzzle)
	assert.Error(err)
	assert.True(errors.Is(err, context.Canceled))
}

func Test_puzzleRepo_Find(t *testing.T) {
	assert := assert.New(t)
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	repo := repository.NewPuzzleRepository(db)
	ctx := context.Background()

	puzzle := models.Puzzle{
		ID:         "puzzle-0",
		ExternalID: "ext-id-0",
		FEN:        "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		Moves: []string{
			"e2e4",
			"e7e5",
			"g1f3",
		},
		Rating:          1500,
		RatingDeviation: 100,
		Popularity:      10,
		Themes: []string{
			"mateIn2",
			"short",
			"sacrifice",
		},
		GameURL:   "https://some.url/some/id",
		CreatedAt: timeutil.Now(),
		UpdatedAt: timeutil.Now(),
	}

	_, found, err := repo.Find(ctx, puzzle.ID)
	assert.NoError(err)
	assert.False(found)

	err = repo.Save(ctx, puzzle)
	assert.NoError(err)

	p, found, err := repo.Find(ctx, puzzle.ID)
	assert.NoError(err)
	assert.True(found)
	assert.Equal(puzzle, p)

	_, found, err = repo.Find(ctx, "other-id")
	assert.NoError(err)
	assert.False(found)

	ctx, cancel := context.WithCancel(ctx)
	cancel()

	_, _, err = repo.Find(ctx, puzzle.ID)
	assert.Error(err)
	assert.True(errors.Is(err, context.Canceled))
}
