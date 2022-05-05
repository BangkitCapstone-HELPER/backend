package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	GetAllArticle() ([]dao.Article, error)
	CreateArticle(article dao.Article) (dao.Article, error)
}

type articleRepoParams struct {
	fx.In
	lib.Database
	Redis lib.Cache
}

func NewArticleRepo(params articleRepoParams) ArticleRepo {
	return &params
}

func (p *articleRepoParams) GetAllArticle() ([]dao.Article, error) {
	articles := []dao.Article{}
	context := context.Background()
	val, err := p.Redis.Cache.Get(context, "article").Result()
	if err == redis.Nil {
		if err := p.Db.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return []dao.Article{}, constants.DatabaseRecordNotFound
			}

			return []dao.Article{}, err
		}
		json, err2 := json.Marshal(articles)
		if err2 != nil {
			return []dao.Article{}, err2
		}
		err3 := p.Redis.Cache.Set(context, "article", json, 10*time.Hour).Err()
		if err3 != nil {
			return []dao.Article{}, err3
		}
	} else {
		err = json.Unmarshal([]byte(val), &articles)
		if err != nil {
			return []dao.Article{}, err
		}
	}
	return articles, nil
}

func (p *articleRepoParams) CreateArticle(article dao.Article) (dao.Article, error) {

	if err := p.Db.Create(&article).Error; err != nil {
		return dao.Article{}, err
	}

	return article, nil

}
