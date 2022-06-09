package repo

import (
	"errors"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type RecommendationRepo interface {
	GetAllArticle() ([]dao.Recommendation, error)
	CreateArticle(article dao.Recommendation) (dao.Recommendation, error)
}

type recommendationRepoParams struct {
	fx.In
	lib.Database
	Redis lib.Cache
}

func NewRecommendationRepo(params recommendationRepoParams) RecommendationRepo {
	return &params
}

func (p *recommendationRepoParams) GetAllArticle() ([]dao.Recommendation, error) {
	articles := []dao.Recommendation{}
	//context := context.Background()
	//val, err := p.Redis.Cache.Get(context, "article").Result()
	//if err == redis.Nil {
	if err := p.Db.Find(&articles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dao.Recommendation{}, constants.DatabaseRecordNotFound
		}

		return []dao.Recommendation{}, err
	}
	//json, err2 := json.Marshal(articles)
	//if err2 != nil {
	//	return []dao.Article{}, err2
	//}
	//err3 := p.Redis.Cache.Set(context, "article", json, 10*time.Hour).Err()
	//	if err != nil {
	//		return []dao.Article{}, err
	//	}
	//} else {
	//	err = json.Unmarshal([]byte(val), &articles)
	//	if err != nil {
	//		return []dao.Article{}, err
	//	}
	//}
	return articles, nil
}

func (p *recommendationRepoParams) CreateArticle(article dao.Recommendation) (dao.Recommendation, error) {

	if err := p.Db.Create(&article).Error; err != nil {
		return dao.Recommendation{}, err
	}

	return article, nil

}
