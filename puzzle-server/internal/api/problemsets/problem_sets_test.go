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
	"github.com/CzarSimon/httputil/client/rpc"
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

func TestCreateProblemSet(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	svc, rbac := setupEnv(ctx)
	router := setupRouter(svc, rbac)

	userID := id.New()
	reqBody := models.CreateProblemSetRequest{
		Name: "ps-name",
		Filter: models.PuzzleFilter{
			Themes:        []string{"passedPawn", "endgame"},
			MinRating:     1300,
			MaxRating:     1500,
			MinPopularity: 100,
			Size:          2,
		},
	}

	req, _ := rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/problem-sets", reqBody)
	attachAuthHeader(req, userID, role.User)
	res := testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)
	var resBody models.ProblemSet
	err := json.NewDecoder(res.Result().Body).Decode(&resBody)
	assert.NoError(err)
	assert.Equal(reqBody.Name, resBody.Name)
	assert.Len(resBody.PuzzleIDs, 1)
	assert.Equal(puzzles[2].ID, resBody.PuzzleIDs[0])

	storedSet, found, err := svc.ProblemSetRepo.Find(ctx, resBody.ID)
	assert.NoError(err)
	assert.True(found)
	assert.Equal(storedSet, resBody)

	reqBody.Name = ""
	req, _ = rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/problem-sets", reqBody)
	attachAuthHeader(req, userID, role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusBadRequest, res.Code)

	reqBody.Name = "valid-name"
	reqBody.Filter.MaxRating = 1000 // Lower than the minimum rating
	req, _ = rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/problem-sets", reqBody)
	attachAuthHeader(req, userID, role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusBadRequest, res.Code)

	reqBody.Filter.MaxRating = 1500
	reqBody.Filter.Size = 50000 // Very large problem size, potential DoS
	req, _ = rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/problem-sets", reqBody)
	attachAuthHeader(req, userID, role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusBadRequest, res.Code)

	reqBody.Filter.Size = 10
	reqBody.Filter.Themes = append(reqBody.Filter.Themes, "prophylaxis") // No matching puzzle with all these themes
	req, _ = rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/problem-sets", reqBody)
	attachAuthHeader(req, userID, role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusUnprocessableEntity, res.Code)
}

func TestCreateProblemSet_UnauthorizedAndForbidden(t *testing.T) {
	test_UnauthorizedAndForbidden(t, http.MethodPost, "/v1/problem-sets", []string{
		role.Anonymous,
		"missing",
	})
}

func TestListProblemSets(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	svc, rbac := setupEnv(ctx)
	router := setupRouter(svc, rbac)

	user1ID := id.New()
	user2ID := id.New()
	sets := []models.ProblemSet{
		{
			ID:             id.New(),
			Name:           "ps-name-0",
			Themes:         []string{"passedPawn", "endgame"},
			RatingInterval: "1300 - 1500",
			UserID:         user1ID,
			PuzzleIDs:      []string{"puzzle-0", "puzzle-1"},
		},
		{
			ID:             id.New(),
			Name:           "ps-name-1",
			Themes:         []string{},
			RatingInterval: "1500 - 1700",
			UserID:         user1ID,
			PuzzleIDs:      []string{"puzzle-2"},
		},
		{
			ID:             id.New(),
			Name:           "ps-name-2",
			Themes:         []string{},
			RatingInterval: "1500 - 1700",
			UserID:         user2ID,
			PuzzleIDs:      []string{"puzzle-2"},
		},
	}

	for _, set := range sets {
		err := svc.ProblemSetRepo.Save(ctx, set)
		assert.NoError(err)
	}

	req := testutil.CreateRequest(http.MethodGet, "/v1/problem-sets", nil)
	attachAuthHeader(req, user1ID, role.User)
	res := testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)
	var user1Sets []models.ProblemSet
	err := json.NewDecoder(res.Result().Body).Decode(&user1Sets)
	assert.NoError(err)
	assert.Len(user1Sets, 2)

	req = testutil.CreateRequest(http.MethodGet, "/v1/problem-sets", nil)
	attachAuthHeader(req, user2ID, role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)
	var user2Sets []models.ProblemSet
	err = json.NewDecoder(res.Result().Body).Decode(&user2Sets)
	assert.NoError(err)
	assert.Len(user2Sets, 1)
	assert.Equal(sets[2].ID, user2Sets[0].ID)

	req = testutil.CreateRequest(http.MethodGet, "/v1/problem-sets", nil)
	attachAuthHeader(req, "user-without-sets", role.User)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)
	var emptySetList []models.ProblemSet
	err = json.NewDecoder(res.Result().Body).Decode(&emptySetList)
	assert.NoError(err)
	assert.Len(emptySetList, 0)
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
	puzzleRepo := repository.NewPuzzleRepository(db)

	svc := &service.ProblemSetService{
		ProblemSetRepo: problemSetRepo,
		PuzzleRepo:     puzzleRepo,
	}

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
