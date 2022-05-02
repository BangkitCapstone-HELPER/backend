package routes

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"go.uber.org/fx"
)

// PolicyRouteParams defines route for policy
type articleRouteParams struct {
	fx.In

	Controller controllers.ArticleController
	Handler    lib.HTTPServer
}

// PolicyRoute defines route for policy
type ArticleRoute Route

// NewPolicyRoutes creates new instance of PolicyRouteImpl
func NewArticleRoutes(pr articleRouteParams) ArticleRoute {
	return &pr
}

// Setup PolicyRouteImpl
func (a *articleRouteParams) Setup() {

	r := a.Handler.RouterV1().Group("/article")
	r.POST("/", a.Controller.CreateArticle)
	r.GET("/", a.Controller.GetArticle)
}
