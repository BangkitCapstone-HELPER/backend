package controllers

import (
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewMenuController),
	fx.Provide(NewArticleController),
	fx.Provide(NewTransactionController),
	fx.Provide(NewFileController),
)
