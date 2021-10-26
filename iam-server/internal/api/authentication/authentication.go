package authentication

import (
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// controller http handler for authenication methods
type controller struct {
	svc *service.AuthenticationService
}

// AttachController attaches a controller to the specified route group.
func AttachController(svc *service.AuthenticationService, r gin.IRouter) {
	controller := &controller{
		svc: svc,
	}

	r.POST("/v1/signup", controller.signUp)
	r.POST("/v1/login", controller.login)
}

func (h *controller) signUp(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "authenication_controller_signup")
	defer span.Finish()

	req, err := parseAuthenticationRequest(c, true)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	res, err := h.svc.Signup(ctx, req)
	if err != nil {
		span.LogFields(log.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *controller) login(c *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "authenication_controller_login")
	defer span.Finish()

	err := httputil.NotImplementedf("login is not yet implemeted")
	c.Error(err)
}

func parseAuthenticationRequest(c *gin.Context, validatePassword bool) (models.AuthenticationRequest, error) {
	var body models.AuthenticationRequest
	err := c.BindJSON(&body)
	if err != nil {
		err = httputil.BadRequestf("failed to parse request body. %w", err)
		return models.AuthenticationRequest{}, err
	}

	err = body.Valid(validatePassword)
	if err != nil {
		return models.AuthenticationRequest{}, httputil.BadRequestError(err)
	}

	return body, nil
}
