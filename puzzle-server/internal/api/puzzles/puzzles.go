package puzzles

import (
	"net/http"

	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	tracelog "github.com/opentracing/opentracing-go/log"
)

// Controller http handler for puzzles
type Controller struct {
	svc *service.PuzzleService
}

// AttachController attaches a controller to the specified route group.
func AttachController(svc *service.PuzzleService, r gin.IRouter) {
	controller := &Controller{
		svc: svc,
	}
	g := r.Group("/v1/puzzles")

	g.GET("/:id", controller.GetPuzzle)
}

func (h *Controller) GetPuzzle(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "puzzle_controller_get_puzzle")
	defer span.Finish()

	id := c.Param("id")
	puzzle, err := h.svc.GetPuzzle(ctx, id)
	if err != nil {
		span.LogFields(tracelog.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, puzzle)
}
