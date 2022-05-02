package repo

import (
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewUserRepo),
	fx.Provide(NewMenuRepo),
	fx.Provide(NewArticleRepo),
	fx.Provide(NewTransactionRepo),
)
