package cycles

import (
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// Controller http handler for cycles
type controller struct {
	svc *service.CycleService
}

// AttachController attaches a controller to the specified route group.
func AttachController(svc *service.CycleService, rbac auth.RBAC, r gin.IRouter) {
	h := &controller{
		svc: svc,
	}
	g := r.Group("/v1/cycles")
	secure := rbac.Secure

	g.GET("/:id", secure(scope.ReadCycle), h.getCycle)
	g.PUT("/:id", secure(scope.UpdateCycle), notImplemented)
}

func (h *controller) getCycle(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "cycle_controller_get_cycle")
	defer span.Finish()

	principal, err := auth.MustGetPrincipal(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	id := c.Param("id")
	cycle, err := h.svc.GetCycle(ctx, id, principal.ID)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, cycle)
}

func notImplemented(c *gin.Context) {
	c.Error(httputil.NotImplementedError(nil))
}
