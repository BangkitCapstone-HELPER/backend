package controllers

import (
	"net/http"
	"strconv"

	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type userControllerParams struct {
	fx.In

	Service services.UserService
	lib.Hash
	lib.JWT
}

type UserController interface {
	GetUser(ctx echo.Context) error
	GetUserById(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	ChangePassword(ctx echo.Context) error
	Login(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
}

func NewUserController(params userControllerParams) UserController {
	return &params
}

// CreateOrder godoc
// @Summary Get User
// @Description Get User
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "user token"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {object} lib.Response
// @Router /api/v1/user/info [get]
func (c userControllerParams) GetUser(ctx echo.Context) error {
	token, _ := utils.ExtractToken(ctx)
	user, err := utils.GetUserFromToken(token, c.JWT)

	var resp lib.Response
	if err != nil {
		resp = lib.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
	} else {
		resp = lib.Response{
			Status: http.StatusOK,
			Data: dto.UserDTO{
				Name:        user.Name,
				Email:       user.Email,
				ID:          uint64(user.ID),
				Address:     user.Address,
				PhoneNumber: user.PhoneNumber,
			},
		}
	}
	return resp.JSON(ctx)
}

// CreateOrder godoc
// @Summary Update User
// @Description Update User
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "user token"
// @Param user_info body dto.UpdateUserDTO true "update user"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {object} lib.Response
// @Router /api/v1/user [patch]
func (c userControllerParams) UpdateUser(ctx echo.Context) error {
	updateUser := dto.UpdateUserDTO{}
	ctx.Bind(&updateUser)
	token, _ := utils.ExtractToken(ctx)
	user, err := utils.GetUserFromToken(token, c.JWT)

	if err != nil {
		return err
	}
	res, err := c.Service.UpdateUser(uint(user.ID), updateUser)
	if err != nil {
		return err
	}
	return lib.Response{
		Status: http.StatusOK,
		Data:   res,
	}.JSON(ctx)
}

// CreateOrder godoc
// @Summary Get User By ID
// @Description Get User By ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Param Authorization header string true "user token"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {object} lib.Response
// @Router /api/v1/user/ [get]
func (c userControllerParams) GetUserById(ctx echo.Context) error {
	idInString := ctx.Param("id")
	id, _ := strconv.ParseUint(idInString, 10, 32)
	user, err := c.Service.GetUserById(id)

	var resp lib.Response
	if err != nil {
		resp = lib.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
	} else {
		resp = lib.Response{
			Status: http.StatusOK,
			Data: dto.UserDTO{
				Name:        user.Name,
				Email:       user.Email,
				ID:          uint64(user.ID),
				Address:     user.Address,
				PhoneNumber: user.PhoneNumber,
			},
		}
	}
	return resp.JSON(ctx)
}

// CreateOrder godoc
// @Summary Create a new user
// @Description Create a new user with the input paylod
// @Tags user
// @Accept  json
// @Produce  json
// @Param user_info body dto.CreateUserRequestDTO true "create user"
// @Success 200 {object} dto.UserDTO
// @Router /api/v1/user/ [post]
func (c userControllerParams) CreateUser(ctx echo.Context) error {
	user := dto.CreateUserRequestDTO{}

	if err := ctx.Bind(&user); err != nil {
		return err
	}

	result, err := c.Service.CreateUser(user)

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
// @Summary Change Password
// @Description Change Password
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "user token"
// @Param user_info body dto.ChangePasswordRequest true "create user"
// @Success 200 {object} dto.UserDTO
// @Router /api/v1/user/password/change [post]
func (c userControllerParams) ChangePassword(ctx echo.Context) error {
	changePasswordRequest := dto.ChangePasswordRequest{}

	if err := ctx.Bind(&changePasswordRequest); err != nil {
		return err
	}
	token, _ := utils.ExtractToken(ctx)
	user, err := utils.GetUserFromToken(token, c.JWT)
	result, err := c.Service.ChangePassword(user.ID, changePasswordRequest)

	var resp lib.Response
	if err != nil {
		resp = lib.Response{
			Status:  http.StatusBadRequest,
			Data:    result,
			Message: err.Error(),
		}
	} else {
		resp = lib.Response{
			Status:  http.StatusOK,
			Message: "Password Change Successfully",
		}
	}
	return resp.JSON(ctx)
}

// CreateOrder godoc
// @Summary Login
// @Tags user
// @Accept  json
// @Produce  json
// @Param login_info body dto.LoginRequest true "user login info"
// @Success 200 {object} dto.LoginResponse
// @Router /api/v1/user/login [post]
func (c userControllerParams) Login(ctx echo.Context) error {
	loginRequest := dto.LoginRequest{}
	ctx.Bind(&loginRequest)
	res, err := c.Service.Login(loginRequest)
	if err != nil {
		return err
	}

	return lib.Response{
		Status: http.StatusOK,
		Data:   res,
	}.JSON(ctx)
}
