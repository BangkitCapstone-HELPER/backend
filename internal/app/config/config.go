package config

import (
	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewViperLoader),
	fx.Provide(NewHTTP),
	fx.Provide(NewDatabase),
	fx.Provide(NewCache),
	fx.Provide(NewJWT),
)
