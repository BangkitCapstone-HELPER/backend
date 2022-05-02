package controllers

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type menuControllerParams struct {
	fx.In
	//Redis   lib.Cache
	Service services.MenuService
}

type MenuController interface {
	CreateMenu(ctx echo.Context) error
	GetMenu(ctx echo.Context) error
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
// @Produce  json
// @Success 200 {object} []dto.MenuDTO
// @Router /api/v1/menu/ [get]
func (c menuControllerParams) GetMenu(ctx echo.Context) error {
	menu := dto.CreateMenuRequestDTO{}
	//var result []dto.MenuDTO
	if err := ctx.Bind(&menu); err != nil {
		return err
	}
	result, err := c.Service.GetAllMenu()
	//context := context.Background()
	//val, err := c.Redis.Cache.Get(context, "menu").Result()
	//if err == redis.Nil {
	//	result, err = c.Service.GetAllMenu()
	//	if err != nil {
	//		return err
	//	}
	//	json, err2 := json.Marshal(result)
	//	if err2 != nil {
	//		return err2
	//	}
	//	err3 := c.Redis.Cache.Set(context, "menu", json, 10*time.Second).Err()
	//	if err3 != nil {
	//		return err3
	//	}
	//} else {
	//	err = json.Unmarshal([]byte(val), &result)
	//	if err != nil {
	//		return err
	//	}
	//}
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
