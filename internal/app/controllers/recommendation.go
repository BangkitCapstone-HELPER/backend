package controllers

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type recommendationControllerParams struct {
	fx.In
	Service services.RecommendationService
}

type RecommendationController interface {
	CreateArticle(ctx echo.Context) error
	GetArticle(ctx echo.Context) error
}

func NewRecommendationController(params recommendationControllerParams) RecommendationController {
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
// @Router /api/v1/recommendation/ [post]
func (c recommendationControllerParams) CreateArticle(ctx echo.Context) error {
	article := dto.RecommendationDTO{}

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
// @Router /api/v1/recommendation/ [get]
func (c recommendationControllerParams) GetArticle(ctx echo.Context) error {
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
