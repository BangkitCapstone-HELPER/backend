package middlewares

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"go.uber.org/fx"

	_ "github.com/BangkitCapstone-HELPER/backend/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type swaggerMiddlewareParams struct {
	fx.In

	Handler lib.HTTPServer
}

type SwaggerMiddleware Middleware

func NewSwaggerMiddleware(params swaggerMiddlewareParams) SwaggerMiddleware {
	return &params
}

func (p swaggerMiddlewareParams) Setup() {
	p.Handler.Engine().GET("/docs/*", echoSwagger.WrapHandler)
}
