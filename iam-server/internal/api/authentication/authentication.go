package authentication

import (
	"github.com/CzarSimon/httputil"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// controller http handler for authenication methods
type controller struct{}

// AttachController attaches a controller to the specified route group.
func AttachController(r gin.IRouter) {
	controller := &controller{}

	g.GET("/v1/signup", controller.signUp)
	g.GET("/v1/login", controller.login)
}

func (h *controller) signUp(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "authenication_controller_signup")
	defer span.Finish()

	err := httputil.NotImplementedf("signup is not yet implemeted")
	c.Error(err)
}

func (h *controller) login(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "authenication_controller_login")
	defer span.Finish()

	err := httputil.NotImplementedf("login is not yet implemeted")
	c.Error(err)
}
