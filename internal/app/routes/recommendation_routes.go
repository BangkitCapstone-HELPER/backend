package routes

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"go.uber.org/fx"
)

// PolicyRouteParams defines route for policy
type recommendationRouteParams struct {
	fx.In

	Controller controllers.RecommendationController
	Handler    lib.HTTPServer
}

// PolicyRoute defines route for policy
type RecommendationRoute Route

// NewPolicyRoutes creates new instance of PolicyRouteImpl
func NewRecommendationRoutes(pr recommendationRouteParams) RecommendationRoute {
	return &pr
}

// Setup PolicyRouteImpl
func (a *recommendationRouteParams) Setup() {

	r := a.Handler.RouterV1().Group("/recommendation")
	r.POST("/", a.Controller.CreateArticle)
	r.GET("/", a.Controller.GetArticle)
}
