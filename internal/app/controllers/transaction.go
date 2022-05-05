package controllers

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type transactionControllerParams struct {
	fx.In
	Service services.TransactionService
	lib.JWT
}

type TransactionController interface {
	CreateTransaction(ctx echo.Context) error
	GetTransactionByUID(ctx echo.Context) error
	UpdateTransaction(ctx echo.Context) error
}

func NewTransactionController(params transactionControllerParams) TransactionController {
	return &params
}

// CreateOrder godoc
// @Summary Create a new transaction
// @Description Create a new transaction with the input paylod
// @Tags transaction
// @Accept  json
// @Produce  json
// @Param transaction_info body dto.CreateTransactionRequestDTO true "create transaction"
// @Success 200 {object} dto.TransactionDTO
// @Router /api/v1/transaction/ [post]
func (c transactionControllerParams) CreateTransaction(ctx echo.Context) error {
	transaction := dto.CreateTransactionRequestDTO{}

	if err := ctx.Bind(&transaction); err != nil {
		return err
	}

	token, _ := utils.ExtractToken(ctx)
	user, err := utils.GetUserFromToken(token, c.JWT)
	if err != nil {
		return err
	}
	transaction.UserID = uint(user.ID)
	result, err := c.Service.CreateTransaction(transaction)

	var resp lib.Response
	if err != nil {
		resp = lib.Response{
			Status:  http.StatusBadRequest,
			Data:    result,
			Message: err.Error(),
		}
	} else {
		resp = lib.Response{
			Status: http.StatusOK,
			Data:   result,
		}
	}
	return resp.JSON(ctx)
}

// CreateOrder godoc
// @Summary Get all transaction
// @Description Get all transaction
// @Tags transaction
// @Accept       json
// @Produce      json
// @Success 200 {object} []dto.TransactionDTO
// @Router /api/v1/transaction/ [get]
func (c transactionControllerParams) GetTransactionByUID(ctx echo.Context) error {
	token, _ := utils.ExtractToken(ctx)
	user, err := utils.GetUserFromToken(token, c.JWT)
	if err != nil {
		return err
	}
	result, err := c.Service.GetAllTransactionByUID(uint(user.ID))
	var resp lib.Response
	if err != nil {
		resp = lib.Response{
			Status:  http.StatusBadRequest,
			Data:    result,
			Message: err.Error(),
		}
	} else {
		resp = lib.Response{
			Status: http.StatusOK,
			Data:   result,
		}
	}
	return resp.JSON(ctx)
}

// CreateOrder godoc
// @Summary update transaction
// @Description update transaction
// @Tags transaction
// @Accept       json
// @Produce      json
// @Param transaction_info body dto.UpdateTransactionDTO true "create transaction"
// @Success 200 {object} []dto.TransactionDTO
// @Router /api/v1/transaction/ [patch]
func (c transactionControllerParams) UpdateTransaction(ctx echo.Context) error {
	transaction := dto.UpdateTransactionDTO{}

	if err := ctx.Bind(&transaction); err != nil {
		return err
	}
	result, err := c.Service.UpdateTransaction(transaction)
	var resp lib.Response
	if err != nil {
		resp = lib.Response{
			Status:  http.StatusBadRequest,
			Data:    result,
			Message: err.Error(),
		}
	} else {
		resp = lib.Response{
			Status: http.StatusOK,
			Data:   result,
		}
	}
	return resp.JSON(ctx)
}
