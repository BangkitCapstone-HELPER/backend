package routes

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

// PolicyRouteParams defines route for policy
type userRouteParams struct {
	fx.In

	Controller     controllers.UserController
	Handler        lib.HTTPServer
	AuthMiddleware middlewares.AuthMiddleware
}

// PolicyRoute defines route for policy
type UserRoute Route

// NewPolicyRoutes creates new instance of PolicyRouteImpl
func NewUserRoutes(pr userRouteParams) UserRoute {
	return &pr
}

// Setup PolicyRouteImpl
func (a *userRouteParams) Setup() {
	s := echo.New()
	s.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})
	r := a.Handler.RouterV1().Group("/user")
	r.POST("/register", a.Controller.CreateUser)
	r.GET("/:id", a.Controller.GetUserById, a.AuthMiddleware.Setup(constants.PermissionNonAdmin))
	r.POST("/password/change", a.Controller.ChangePassword)
	r.GET("/info", a.Controller.GetUser)
	r.POST("/", a.Controller.CreateUser)
	r.PATCH("/", a.Controller.UpdateUser)
	r.POST("/login", a.Controller.Login)
}
