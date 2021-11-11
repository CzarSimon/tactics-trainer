package cycles

import (
	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/service"
	"github.com/gin-gonic/gin"
)

// Controller http handler for cycles
type controller struct {
	svc *service.CycleService
}

// AttachController attaches a controller to the specified route group.
func AttachController(svc *service.CycleService, rbac auth.RBAC, r gin.IRouter) {
	g := r.Group("/v1/cycles")
	secure := rbac.Secure

	g.GET("/:id", secure(scope.ReadCycle), notImplemented)
	g.PUT("/:id", secure(scope.UpdateCycle), notImplemented)
}

func notImplemented(c *gin.Context) {
	c.Error(httputil.NotImplementedError(nil))
}
