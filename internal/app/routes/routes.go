package routes

import (
	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewMenuRoutes),
	fx.Provide(NewArticleRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
//go:generate mockery --name=Route --case underscore --inpackage
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoute,
	menuRoutes MenuRoute,
	articleRoutes ArticleRoute,
) Routes {
	return Routes{
		userRoutes,
		menuRoutes,
		articleRoutes,
	}
}

// Setup all the route
func (a Routes) Setup() {
	for _, route := range a {
		route.Setup()
	}
}
