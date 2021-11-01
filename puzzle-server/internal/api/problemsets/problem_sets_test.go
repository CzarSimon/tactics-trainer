package problemsets_test

import (
	"fmt"
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
	"github.com/stretchr/testify/assert"
)

var jwtCreds = jwt.Credentials{
	Issuer: "problemsets_test",
	Secret: "1f49ac286fd8565b160a6712c48e88a9",
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
	rbac := setupEnv()
	router := setupRouter(rbac)

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

func setupEnv() auth.RBAC {
	return auth.NewRBAC(jwtCreds)
}

func setupRouter(rbac auth.RBAC) http.Handler {
	r := httputil.NewRouter("puzzle-server", func() error {
		return nil
	})
	problemsets.AttachController(rbac, r)
	return r
}
