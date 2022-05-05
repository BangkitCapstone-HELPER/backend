package routes

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/controllers"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"go.uber.org/fx"
)

// PolicyRouteParams defines route for policy
type transactionRouteParams struct {
	fx.In

	Controller controllers.TransactionController
	Handler    lib.HTTPServer
}

// PolicyRoute defines route for policy
type TransactionRoute Route

// NewPolicyRoutes creates new instance of PolicyRouteImpl
func NewTransactionRoutes(pr transactionRouteParams) TransactionRoute {
	return &pr
}

// Setup PolicyRouteImpl
func (a *transactionRouteParams) Setup() {

	r := a.Handler.RouterV1().Group("/transaction")
	r.POST("/", a.Controller.CreateTransaction)
	r.GET("/", a.Controller.GetTransactionByUID)
	r.PATCH("/", a.Controller.UpdateTransaction)
}
