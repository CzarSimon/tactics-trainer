package auth_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/role"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRBAC(t *testing.T) {
	assert := assert.New(t)
	r := httputil.NewRouter("httputil-test", func() error {
		return nil
	})

	creds := jwt.Credentials{
		Issuer: "test-issuer",
		Secret: "some-very-secret-secret",
	}

	rbac := auth.NewRBAC(creds)
	r.GET("/auth-required", rbac.Secure(scope.ListProblemSets), func(c *gin.Context) {
		_, err := auth.MustGetPrincipal(c)
		if err != nil {
			c.Error(err)
			return
		}

		httputil.SendOK(c)
	})

	jwtIssuer := jwt.NewIssuer(creds)

	anonymousToken, err := jwtIssuer.Issue(jwt.User{
		ID:    "some-id",
		Roles: []string{role.Anonymous},
	}, time.Hour)
	assert.NoError(err)

	userToken, err := jwtIssuer.Issue(jwt.User{
		ID:    "some-id",
		Roles: []string{role.User},
	}, time.Hour)
	assert.NoError(err)

	unknownToken, err := jwtIssuer.Issue(jwt.User{
		ID:    "some-id",
		Roles: []string{"UNKNOWN"},
	}, time.Hour)
	assert.NoError(err)

	req := testutil.CreateRequest(http.MethodGet, "/health", nil)
	res := testutil.PerformRequest(r, req)
	assert.Equal(http.StatusOK, res.Code)

	req = testutil.CreateRequest(http.MethodGet, "/auth-required", nil)
	res = testutil.PerformRequest(r, req)
	assert.Equal(http.StatusUnauthorized, res.Code)

	req = testutil.CreateRequest(http.MethodGet, "/auth-required", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", unknownToken))
	res = testutil.PerformRequest(r, req)
	assert.Equal(http.StatusForbidden, res.Code)

	req = testutil.CreateRequest(http.MethodGet, "/auth-required", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", anonymousToken))
	res = testutil.PerformRequest(r, req)
	assert.Equal(http.StatusForbidden, res.Code)

	req = testutil.CreateRequest(http.MethodGet, "/auth-required", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userToken))
	res = testutil.PerformRequest(r, req)
	assert.Equal(http.StatusOK, res.Code)
}
