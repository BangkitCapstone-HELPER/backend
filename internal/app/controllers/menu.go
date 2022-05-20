package controllers

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
	"strconv"
)

type menuControllerParams struct {
	fx.In
	//Redis   lib.Cache
	Service services.MenuService
}

type MenuController interface {
	CreateMenu(ctx echo.Context) error
	GetMenu(ctx echo.Context) error
	GetMenuById(ctx echo.Context) error
}

func NewMenuController(params menuControllerParams) MenuController {
	return &params
}

// CreateOrder godoc
// @Summary Create a new menu
// @Description Create a new menu with the input paylod
// @Tags menu
// @Accept  json
// @Produce  json
// @Param menu_info body dto.CreateMenuRequestDTO true "create menu"
// @Success 200 {object} dto.MenuDTO
// @Router /api/v1/menu/ [post]
func (c menuControllerParams) CreateMenu(ctx echo.Context) error {
	menu := dto.CreateMenuRequestDTO{}

	if err := ctx.Bind(&menu); err != nil {
		return err
	}

	result, err := c.Service.CreateMenu(menu)

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
// @Summary Get all menu
// @Description Get all menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Success 200 {object} []dto.MenuDTO
// @Router /api/v1/menu/ [get]
func (c menuControllerParams) GetMenu(ctx echo.Context) error {
	result, err := c.Service.GetAllMenu()
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
// @Summary Get menu by id
// @Description Get menu by id
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path string true "menu id"
// @Success 200 {object} dto.MenuDTO
// @Router /api/v1/menu/ [get]
func (c menuControllerParams) GetMenuById(ctx echo.Context) error {
	idInString := ctx.Param("id")
	id, _ := strconv.ParseUint(idInString, 10, 32)
	result, err := c.Service.GetMenu(id)
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
// @Summary Delete menu by id
// @Description Delete menu by id
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path string true "menu id"
// @Success 200 {object} dto.MenuDTO
// @Router /api/v1/menu/ [delete]
func (c menuControllerParams) DeleteMenuById(ctx echo.Context) error {
	idInString := ctx.Param("id")
	id, _ := strconv.ParseUint(idInString, 10, 32)
	_, err := c.Service.DeleteMenu(id)
	var resp lib.Response
	if err != nil {
		resp = lib.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		resp = lib.Response{
			Status:  http.StatusOK,
			Message: "Data delete Successfully",
		}
	}
	return resp.JSON(ctx)
}
