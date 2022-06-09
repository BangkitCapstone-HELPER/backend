package dto

import (
	"encoding/json"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/datatypes"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
)

type (
	RecommendationDTO struct {
		Title       string                 `json:"title"`
		Link        []string               `json:"link"`
		GuID        string                 `json:"guid"`
		PubDate     string                 `json:"pubDate"`
		Description map[string]interface{} `json:"description"`
		Enclosure   map[string]interface{} `json:"enclosure"`
	}
)

func (c RecommendationDTO) ToDAO() dao.Recommendation {
	var article map[string]interface{}
	data, _ := json.Marshal(c)
	_ = json.Unmarshal(data, &article)
	return dao.Recommendation{
		Recommendation: datatypes.JSONMap(article),
	}
}

func NewRecommendationDTO(article dao.Recommendation) RecommendationDTO {
	var result RecommendationDTO
	newArticle, _ := article.Recommendation.MarshalJSON()
	_ = json.Unmarshal(newArticle, &result)
	return result
}
