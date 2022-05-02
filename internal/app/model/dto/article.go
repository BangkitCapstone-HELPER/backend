package dto

import (
	"encoding/json"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/datatypes"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
)

type (
	ArticleDTO struct {
		Title       string                 `json:"title"`
		Link        []string               `json:"link"`
		GuID        string                 `json:"guid"`
		PubDate     string                 `json:"pubDate"`
		Description map[string]interface{} `json:"description"`
		Enclosure   map[string]interface{} `json:"enclosure"`
	}
)

func (c ArticleDTO) ToDAO() dao.Article {
	var article map[string]interface{}
	data, _ := json.Marshal(c)
	_ = json.Unmarshal(data, &article)
	return dao.Article{
		Article: datatypes.JSONMap(article),
	}
}

func NewArticleDTO(article dao.Article) ArticleDTO {
	var result ArticleDTO
	newArticle, _ := article.Article.MarshalJSON()
	_ = json.Unmarshal(newArticle, &result)
	return result
}
