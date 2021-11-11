package cycles_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/id"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/role"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/api/cycles"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

const userID = "cycle-controller-test-user"

var jwtCreds = jwt.Credentials{
	Issuer: "problemsets_test",
	Secret: "1f49ac286fd8565b160a6712c48e88a9",
}

var problemSet = models.ProblemSet{
	ID:             id.New(),
	Name:           "ps-name",
	Themes:         []string{"passedPawn", "endgame"},
	RatingInterval: "1300 - 1500",
	UserID:         userID,
	PuzzleIDs:      []string{"puzzle-0", "puzzle-1"},
}

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

func TestGetCycle(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	svc, rbac := setupEnv(ctx)
	router := setupRouter(svc, rbac)

	cycle := models.Cycle{
		ID:              id.New(),
		Number:          1,
		ProblemSetID:    problemSet.ID,
		CurrentPuzzleID: "puzzle-2",
		CompleatedAt:    timeutil.Now(),
		CreatedAt:       timeutil.Now(),
		UpdatedAt:       timeutil.Now(),
	}

	err := svc.CycleRepo.Save(ctx, cycle)
	assert.NoError(err)

	path := fmt.Sprintf("/v1/cycles/%s", cycle.ID)
	req := testutil.CreateRequest(http.MethodGet, path, nil)
	attachAuthHeader(req, userID, role.User)
	res := testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)

	var body models.Cycle
	err = json.NewDecoder(res.Result().Body).Decode(&body)
	assert.NoError(err)
	assert.Equal(cycle, body)

	req = testutil.CreateRequest(http.MethodGet, path, nil)
	attachAuthHeader(req, "other-user-id", role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusForbidden, res.Code)

	req = testutil.CreateRequest(http.MethodGet, "/v1/cycles/missing-id", nil)
	attachAuthHeader(req, userID, role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusNotFound, res.Code)
}

func TestUpdateCycle_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodPut, "/v1/cycles/some-id", []string{
		role.Anonymous,
		"missing",
	})
}

func TestGetCycle_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodGet, "/v1/cycles/some-id", []string{
		role.Anonymous,
		"missing",
	})
}

func test_UnauthorizedAndForbidden(t *testing.T, method, path string, forbidden []string) {
	assert := assert.New(t)
	router := setupRouter(setupEnv(context.Background()))

	req := testutil.CreateRequest(method, path, nil)
	res := testutil.PerformRequest(router, req)
	assert.Equal(http.StatusUnauthorized, res.Code)

	for _, role := range forbidden {
		req = testutil.CreateRequest(method, path, nil)
		attachAuthHeader(req, id.New(), role)
		res = testutil.PerformRequest(router, req)
		assert.Equal(http.StatusForbidden, res.Code)
	}
}

func attachAuthHeader(req *http.Request, userId, role string) {
	token, _ := jwt.NewIssuer(jwtCreds).Issue(jwt.User{ID: userId, Roles: []string{role}}, time.Hour)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func setupEnv(ctx context.Context) (*service.CycleService, auth.RBAC) {
	db := testutil.InMemoryDB(true, "../../../resources/db/sqlite")
	problemSetRepo := repository.NewProblemSetRepository(db)
	puzzleRepo := repository.NewPuzzleRepository(db)
	cycleRepo := repository.NewCycleRepository(db)

	svc := &service.CycleService{
		CycleRepo:      cycleRepo,
		ProblemSetRepo: problemSetRepo,
	}

	for _, p := range puzzles {
		err := puzzleRepo.Save(ctx, p)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := problemSetRepo.Save(ctx, problemSet)
	if err != nil {
		log.Fatal(err)
	}

	return svc, auth.NewRBAC(jwtCreds)
}

func setupRouter(svc *service.CycleService, rbac auth.RBAC) http.Handler {
	r := httputil.NewRouter("puzzle-server", func() error {
		return nil
	})
	cycles.AttachController(svc, rbac, r)
	return r
}
