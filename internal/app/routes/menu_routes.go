package routes

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"go.uber.org/fx"
)

// PolicyRouteParams defines route for policy
type menuRouteParams struct {
	fx.In

	Controller controllers.MenuController
	Handler    lib.HTTPServer
}

// PolicyRoute defines route for policy
type MenuRoute Route

// NewPolicyRoutes creates new instance of PolicyRouteImpl
func NewMenuRoutes(pr menuRouteParams) MenuRoute {
	return &pr
}

// Setup PolicyRouteImpl
func (a *menuRouteParams) Setup() {

	r := a.Handler.RouterV1().Group("/menu")
	r.POST("/", a.Controller.CreateMenu)
	r.GET("/", a.Controller.GetMenu)
}
