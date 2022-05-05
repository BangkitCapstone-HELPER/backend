package controllers

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type articleControllerParams struct {
	fx.In
	Service services.ArticleService
}

type ArticleController interface {
	CreateArticle(ctx echo.Context) error
	GetArticle(ctx echo.Context) error
}

func NewArticleController(params articleControllerParams) ArticleController {
	return &params
}

// CreateOrder godoc
// @Summary Create a new article
// @Description Create a new article with the input paylod
// @Tags article
// @Accept  json
// @Produce  json
// @Param article_info body dto.ArticleDTO true "create article"
// @Success 200 {object} dto.ArticleDTO
// @Router /api/v1/article/ [post]
func (c articleControllerParams) CreateArticle(ctx echo.Context) error {
	article := dto.ArticleDTO{}

	if err := ctx.Bind(&article); err != nil {
		return err
	}

	result, err := c.Service.CreateArticle(article)

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
// @Summary Get all article
// @Description Get all article
// @Tags article
// @Accept       json
// @Produce      json
// @Success 200 {object} []dto.ArticleDTO
// @Router /api/v1/article/ [get]
func (c articleControllerParams) GetArticle(ctx echo.Context) error {
	menu := dto.CreateMenuRequestDTO{}
	if err := ctx.Bind(&menu); err != nil {
		return err
	}
	result, err := c.Service.GetAllArticle()
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
