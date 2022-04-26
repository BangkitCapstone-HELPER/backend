package lib

import (
	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewHTTPHandler),
	fx.Provide(NewHash),
	fx.Provide(NewJWT),
)
