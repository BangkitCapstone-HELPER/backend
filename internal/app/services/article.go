package services

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/repo"
	"go.uber.org/fx"
)

type articleServiceParams struct {
	fx.In
	ArticleRepo repo.ArticleRepo
}

type ArticleService interface {
	CreateArticle(user dto.ArticleDTO) (dto.ArticleDTO, error)
	GetAllArticle() ([]dto.ArticleDTO, error)
}

func NewArticleService(params articleServiceParams) ArticleService {
	return &params
}

func (u *articleServiceParams) CreateArticle(article dto.ArticleDTO) (dto.ArticleDTO, error) {
	newArticle, err := u.ArticleRepo.CreateArticle(article.ToDAO())

	return dto.NewArticleDTO(newArticle), err
}

func (u *articleServiceParams) GetAllArticle() ([]dto.ArticleDTO, error) {
	articles, err := u.ArticleRepo.GetAllArticle()

	var result []dto.ArticleDTO
	for _, article := range articles {
		result = append(result, dto.NewArticleDTO(article))
	}

	return result, err
}
