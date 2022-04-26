package middlewares

import "go.uber.org/fx"

type Middleware interface {
	Setup()
}

var Module = fx.Options(
	fx.Provide(NewSwaggerMiddleware),
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewMiddlewares),
	fx.Provide(NewPrivacyMiddleware),
)

// Middlewares contains multiple middleware
type Middlewares []Middleware

func NewMiddlewares(
	swaggerMiddleware SwaggerMiddleware,
) Middlewares {
	return Middlewares{
		swaggerMiddleware,
	}
}

// Setup sets up middlewares
func (a Middlewares) Setup() {
	for _, mdw := range a {
		mdw.Setup()
	}
}
