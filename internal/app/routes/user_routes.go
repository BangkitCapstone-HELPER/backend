package routes

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/middlewares"
	"go.uber.org/fx"
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

	r := a.Handler.RouterV1().Group("/user")
	r.POST("/register", a.Controller.CreateUser)
	r.GET("/:id", a.Controller.GetUserById, a.AuthMiddleware.Setup(constants.PermissionNonAdmin))
	r.GET("/info", a.Controller.GetUser)
	r.POST("/", a.Controller.CreateUser)
	r.POST("/login", a.Controller.Login)
}
