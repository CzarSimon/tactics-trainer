package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/httputil/logger"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/role"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracelog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
)

var log = logger.GetDefaultLogger("gopkg/auth")

const principalKey = "tactics-trainer:gopkg:auth:principalKey"

// GetPrincipal returns the authenticated user if exists.
func GetPrincipal(c *gin.Context) (jwt.User, bool) {
	val, ok := c.Get(principalKey)
	if !ok {
		return jwt.User{}, false
	}

	user, ok := val.(jwt.User)
	return user, ok
}

// MustGetPrincipal returns the authenticated user or an error if none exists.
func MustGetPrincipal(c *gin.Context) (jwt.User, error) {
	principal, ok := GetPrincipal(c)
	if !ok {
		return jwt.User{}, fmt.Errorf("failed to parse prinipal from authenticated request")
	}

	return principal, nil
}

// RBAC adds role based access controll checks extracting roles from jwt.
type RBAC struct {
	Verifier jwt.Verifier
}

// NewRBAC creates a new RBAC struct with sane defaults.
func NewRBAC(creds jwt.Credentials) RBAC {
	return RBAC{
		Verifier: jwt.NewVerifier(creds, time.Minute),
	}
}

// Secure checks if a request was made with a jwt containing a specified list of roles.
func (r *RBAC) Secure(requiredScope scope.Scope) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := extractUserFromRequest(c, r.Verifier)
		if err != nil {
			logError(c, err)
			c.AbortWithStatusJSON(err.Status, err)
			return
		}

		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span.SetBaggageItem("user-id", user.ID)
			span.SetBaggageItem("user-roles", strings.Join(user.Roles, ";"))
		}
		c.Set(principalKey, user)

		for _, roleName := range user.Roles {
			r, ok := role.Get(roleName)
			if !ok {
				continue
			}

			if r.HasScope(requiredScope) {
				c.Next()
				return
			}
		}

		err = httputil.Forbiddenf("%s %s access denied for %s", c.Request.Method, c.Request.URL.Path, user)
		logError(c, err)
		c.AbortWithStatusJSON(err.Status, err)
	}
}

func extractUserFromRequest(c *gin.Context, verifier jwt.Verifier) (jwt.User, *httputil.Error) {
	token, err := exctractToken(c)
	if err != nil {
		return jwt.User{}, err
	}

	user, jwtErr := verifier.Verify(token)
	if jwtErr != nil {
		return jwt.User{}, httputil.UnauthorizedError(jwtErr)
	}

	return user, nil
}

func exctractToken(c *gin.Context) (string, *httputil.Error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", httputil.Unauthorizedf("no authorization header provided")
	}

	token := strings.Replace(header, "Bearer ", "", 1)
	return token, nil
}

func logError(c *gin.Context, err *httputil.Error) {
	span := opentracing.SpanFromContext(c.Request.Context())
	if span != nil {
		span.LogFields(tracelog.Error(err))
		ext.HTTPStatusCode.Set(span, uint16(err.Status))
	}

	if err.Status < 500 {
		log.Info(err.Message,
			zap.Int("status", err.Status),
			zap.String("errorId", err.ID),
			zap.Error(err.Err))
		return
	}
	log.Error(err.Message,
		zap.Int("status", err.Status),
		zap.String("errorId", err.ID),
		zap.Error(err.Err))
}
