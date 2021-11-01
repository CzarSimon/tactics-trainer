package problemsets

import (
	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Controller http handler for problem sets
type controller struct{}

// AttachController attaches a controller to the specified route group.
func AttachController(rbac auth.RBAC, r gin.IRouter) {
	controller := &controller{}
	g := r.Group("/v1/problem-sets")
	secure := rbac.Secure

	g.GET("", secure(scope.ListProblemSets), notImplemented)
	g.POST("", secure(scope.CreateProblemSet), controller.CreateProblemSet)
	g.GET("/:setId", secure(scope.ReadProblemSet), notImplemented)
	g.DELETE("/:setId", secure(scope.DeleteProblemSet), notImplemented)
	g.PUT("/:setId/puzzles/:puzzleId", secure(scope.UpdateProblemSet), notImplemented)
}

func (h *controller) CreateProblemSet(c *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "problem_set_controller_create_problem_set")
	defer span.Finish()

	notImplemented(c)
}

func notImplemented(c *gin.Context) {
	c.Error(httputil.NotImplementedError(nil))
}
