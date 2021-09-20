package puzzles_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/id"
	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/api/puzzles"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestGetPuzzles(t *testing.T) {
	assert := assert.New(t)
	svc := setupPuzzleService()
	router := setupRouter(svc)
	ctx := context.Background()

	puzzle := getPuzzle()

	path := fmt.Sprintf("/v1/puzzles/%s", puzzle.ID)
	req := testutil.CreateRequest(http.MethodGet, path, nil)
	res := testutil.PerformRequest(router, req)
	assert.Equal(http.StatusNotFound, res.Code)

	err := svc.PuzzleRepo.Save(ctx, puzzle)
	assert.NoError(err)

	req = testutil.CreateRequest(http.MethodGet, path, nil)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)
	var body models.Puzzle
	err = json.NewDecoder(res.Result().Body).Decode(&body)
	assert.NoError(err)
	assert.Equal(puzzle, body)

	notFoundPath := fmt.Sprintf("/v1/puzzles/%s", id.New())
	req = testutil.CreateRequest(http.MethodGet, notFoundPath, nil)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusNotFound, res.Code)
}

func getPuzzle() models.Puzzle {
	puzzleId := id.New()
	return models.Puzzle{
		ID:         puzzleId,
		ExternalID: fmt.Sprintf("ext-%s", puzzleId),
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
}

func setupPuzzleService() *service.PuzzleService {
	db := testutil.InMemoryDB(true, "../../../resources/db/sqlite")
	repo := repository.NewPuzzleRepository(db)
	return &service.PuzzleService{
		PuzzleRepo: repo,
	}
}

func setupRouter(svc *service.PuzzleService) http.Handler {
	r := httputil.NewRouter("puzzle-server", func() error {
		return nil
	})
	puzzles.AttachController(svc, r)
	return r
}
