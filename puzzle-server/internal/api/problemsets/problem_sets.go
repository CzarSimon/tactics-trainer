package problemsets

import (
	"net/http"
	"strings"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// Controller http handler for problem sets
type controller struct {
	svc *service.ProblemSetService
}

// AttachController attaches a controller to the specified route group.
func AttachController(svc *service.ProblemSetService, rbac auth.RBAC, r gin.IRouter) {
	controller := &controller{
		svc: svc,
	}
	g := r.Group("/v1/problem-sets")
	secure := rbac.Secure

	g.GET("", secure(scope.ListProblemSets), controller.listProblemSets)
	g.POST("", secure(scope.CreateProblemSet), controller.createProblemSet)
	g.GET("/:setId", secure(scope.ReadProblemSet), controller.getProblemSet)
	g.DELETE("/:setId", secure(scope.DeleteProblemSet), controller.archiveProblemSet)
	g.GET("/:setId/cycles", secure(scope.ListProblemSetCycles), controller.listProblemSetCycles)
	g.POST("/:setId/cycles", secure(scope.CreateProblemSetCycle), controller.createProblemSetCycle)
}

func (h *controller) listProblemSets(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "problem_set_controller_list_problem_sets")
	defer span.Finish()

	includeArchived := parseFlag(c, "includeArchived")
	principal, err := auth.MustGetPrincipal(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	sets, err := h.svc.ListProblemSets(ctx, principal.ID, includeArchived)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, sets)
}

func (h *controller) createProblemSet(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "problem_set_controller_create_problem_set")
	defer span.Finish()

	req, err := parseCreateProblemSetRequest(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	principal, err := auth.MustGetPrincipal(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	set, err := h.svc.CreateProblemSet(ctx, req, principal.ID)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (h *controller) getProblemSet(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "problem_set_controller_get_problem_set")
	defer span.Finish()

	id := c.Param("setId")
	principal, err := auth.MustGetPrincipal(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	set, err := h.svc.GetProblemSet(ctx, id, principal.ID)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (h *controller) archiveProblemSet(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "problem_set_controller_archive_problem_set")
	defer span.Finish()

	id := c.Param("setId")
	principal, err := auth.MustGetPrincipal(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	err = h.svc.ArchiveProblemSet(ctx, id, principal.ID)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	httputil.SendOK(c)
}

func (h *controller) createProblemSetCycle(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "problem_set_controller_create_problem_set_cycle")
	defer span.Finish()

	principal, err := auth.MustGetPrincipal(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	id := c.Param("setId")
	cycle, err := h.svc.CreateProblemSetCycle(ctx, id, principal.ID)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, cycle)
}

func (h *controller) listProblemSetCycles(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "problem_set_controller_list_problem_set_cycles")
	defer span.Finish()

	principal, err := auth.MustGetPrincipal(c)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	id := c.Param("setId")
	onlyActive := parseFlag(c, "onlyActive")
	cycles, err := h.svc.ListProblemSetCycles(ctx, id, principal.ID, onlyActive)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, cycles)
}

func parseCreateProblemSetRequest(c *gin.Context) (models.CreateProblemSetRequest, error) {
	var body models.CreateProblemSetRequest
	err := c.BindJSON(&body)
	if err != nil {
		err = httputil.BadRequestf("failed to parse request body. %w", err)
		return models.CreateProblemSetRequest{}, err
	}

	err = body.Valid()
	if err != nil {
		err = httputil.BadRequestf("invalid %s. %w", body, err)
		return models.CreateProblemSetRequest{}, err
	}

	return body, nil
}

func parseFlag(c *gin.Context, name string) bool {
	val, ok := c.GetQuery(name)
	if !ok {
		return false
	}

	return strings.ToLower(val) == "true" || val == "1"
}

func notImplemented(c *gin.Context) {
	c.Error(httputil.NotImplementedError(nil))
}
