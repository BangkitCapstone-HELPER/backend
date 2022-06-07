package routes

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/middlewares"
	"go.uber.org/fx"
)

type fileRouteParams struct {
	fx.In

	Controller controllers.FileController
	Handler    lib.HTTPServer
	Middleware middlewares.AuthMiddleware
}

// PolicyRoute defines route for policy
type FileRoute Route

// NewPolicyRoutes creates new instance of PolicyRouteImpl
func NewFileRoutes(pr fileRouteParams) FileRoute {
	return &pr
}

// Setup PolicyRouteImpl
func (a *fileRouteParams) Setup() {

	r := a.Handler.RouterV1().Group("/file")
	//r.GET("/:id", a.Controller.GetFile)
	r.POST("/", a.Controller.UploadFile)
	r.POST("/predict/", a.Controller.PredictImage)
}
