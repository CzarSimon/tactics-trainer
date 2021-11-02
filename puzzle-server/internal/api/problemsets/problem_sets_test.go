package problemsets_test

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
	"github.com/CzarSimon/tactics-trainer/gopkg/auth"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/role"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/api/problemsets"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var jwtCreds = jwt.Credentials{
	Issuer: "problemsets_test",
	Secret: "1f49ac286fd8565b160a6712c48e88a9",
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

func TestCreateProblemSets_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodPost, "/v1/problem-sets", []string{
		role.Anonymous,
		"missing",
	})
}

func TestListProblemSets_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodGet, "/v1/problem-sets", []string{
		role.Anonymous,
		"missing",
	})
}

func TestGetProblemSet(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	svc, rbac := setupEnv(ctx)
	router := setupRouter(svc, rbac)

	userID := id.New()
	set := models.ProblemSet{
		ID:             id.New(),
		Name:           "ps-name",
		Themes:         []string{"passedPawn", "endgame"},
		RatingInterval: "1300 - 1500",
		UserID:         userID,
		PuzzleIDs:      []string{"puzzle-0", "puzzle-1"},
	}

	err := svc.ProblemSetRepo.Save(ctx, set)
	assert.NoError(err)

	req := testutil.CreateRequest(http.MethodGet, "/v1/problem-sets/"+set.ID, nil)
	attachAuthHeader(req, userID, role.User)
	res := testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)
	var body models.ProblemSet
	err = json.NewDecoder(res.Result().Body).Decode(&body)
	assert.NoError(err)
	assert.Equal(set, body)

	req = testutil.CreateRequest(http.MethodGet, "/v1/problem-sets/missing-set-id", nil)
	attachAuthHeader(req, userID, role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusNotFound, res.Code)

	req = testutil.CreateRequest(http.MethodGet, "/v1/problem-sets/"+set.ID, nil)
	attachAuthHeader(req, "other-user-id", role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusForbidden, res.Code)
}

func TestGetProblemSet_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodGet, "/v1/problem-sets/some-id", []string{
		role.Anonymous,
		"missing",
	})
}

func TestDeleteProblemSet_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodDelete, "/v1/problem-sets/some-id", []string{
		role.Anonymous,
		"missing",
	})
}

func TestUpdateProblemSet_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodPut, "/v1/problem-sets/set-id/puzzles/puzzle-id", []string{
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

func setupEnv(ctx context.Context) (*service.ProblemSetService, auth.RBAC) {
	db := testutil.InMemoryDB(true, "../../../resources/db/sqlite")
	problemSetRepo := repository.NewProblemSetRepository(db)
	svc := &service.ProblemSetService{
		ProblemSetRepo: problemSetRepo,
	}

	puzzleRepo := repository.NewPuzzleRepository(db)
	for _, p := range puzzles {
		err := puzzleRepo.Save(ctx, p)
		if err != nil {
			log.Fatal(err)
		}
	}

	return svc, auth.NewRBAC(jwtCreds)
}

func setupRouter(svc *service.ProblemSetService, rbac auth.RBAC) http.Handler {
	r := httputil.NewRouter("puzzle-server", func() error {
		return nil
	})
	problemsets.AttachController(svc, rbac, r)
	return r
}
