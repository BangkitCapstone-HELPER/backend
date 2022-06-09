package services

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/repo"
	"go.uber.org/fx"
)

type recommendationServiceParams struct {
	fx.In
	ArticleRepo repo.RecommendationRepo
}

type RecommendationService interface {
	CreateArticle(user dto.RecommendationDTO) (dto.RecommendationDTO, error)
	GetAllArticle() ([]dto.RecommendationDTO, error)
}

func NewRecommendationService(params recommendationServiceParams) RecommendationService {
	return &params
}

func (u *recommendationServiceParams) CreateArticle(article dto.RecommendationDTO) (dto.RecommendationDTO, error) {
	newArticle, err := u.ArticleRepo.CreateArticle(article.ToDAO())

	return dto.NewRecommendationDTO(newArticle), err
}

func (u *recommendationServiceParams) GetAllArticle() ([]dto.RecommendationDTO, error) {
	articles, err := u.ArticleRepo.GetAllArticle()

	var result []dto.RecommendationDTO
	for _, article := range articles {
		result = append(result, dto.NewRecommendationDTO(article))
	}

	return result, err
}
