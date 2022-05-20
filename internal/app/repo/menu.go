package repo

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm/clause"

	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type MenuRepo interface {
	GetMenu(id uint64) (dao.Menu, error)
	GetAllMenu() ([]dao.Menu, error)
	CreateMenu(user dao.Menu) (dao.Menu, error)
	CreateDayMenu(dayMenu dao.DayMenu) (dao.DayMenu, error)
	CreateItem(item dao.Item) (dao.Item, error)
	//UpdateUser(userId uint64, updateMap map[string]interface{}) (dao.User, error)
	DeleteMenu(id uint64) (dao.Menu, error)
}

type menuRepoParams struct {
	fx.In
	lib.Database
	Redis lib.Cache
}

func NewMenuRepo(params menuRepoParams) MenuRepo {
	return &params
}

func (p *menuRepoParams) GetMenu(id uint64) (dao.Menu, error) {
	menu := dao.Menu{}

	if err := p.Db.First(&menu, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Menu{}, constants.DatabaseRecordNotFound
		}

		return dao.Menu{}, err
	}
	return menu, nil
}

func (p *menuRepoParams) DeleteMenu(id uint64) (dao.Menu, error) {
	menu := dao.Menu{}

	if err := p.Db.Delete(&menu, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Menu{}, constants.DatabaseRecordNotFound
		}

		return dao.Menu{}, err
	}
	return menu, nil
}

func (p *menuRepoParams) GetAllMenu() ([]dao.Menu, error) {
	menus := []dao.Menu{}
	context := context.Background()
	val, err := p.Redis.Cache.Get(context, "menu").Result()
	if err == redis.Nil {
		if err := p.Db.Preload("DayMenus.Items").Preload(clause.Associations).Find(&menus).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return []dao.Menu{}, constants.DatabaseRecordNotFound
			}

			return []dao.Menu{}, err
		}
		json, err2 := json.Marshal(menus)
		if err2 != nil {
			return []dao.Menu{}, err2
		}
		err3 := p.Redis.Cache.Set(context, "menu", json, 10*time.Hour).Err()
		if err3 != nil {
			return []dao.Menu{}, err3
		}
	} else {
		err = json.Unmarshal([]byte(val), &menus)
		if err != nil {
			return []dao.Menu{}, err
		}
	}
	return menus, nil
}

func (p *menuRepoParams) CreateMenu(menu dao.Menu) (dao.Menu, error) {

	if err := p.Db.Create(&menu).Error; err != nil {
		return dao.Menu{}, err
	}

	return menu, nil

}

func (p *menuRepoParams) CreateDayMenu(dayMenu dao.DayMenu) (dao.DayMenu, error) {

	if err := p.Db.Create(&dayMenu).Error; err != nil {
		return dao.DayMenu{}, err
	}

	return dayMenu, nil

}

func (p *menuRepoParams) CreateItem(item dao.Item) (dao.Item, error) {

	if err := p.Db.Create(&item).Error; err != nil {
		return dao.Item{}, err
	}

	return item, nil

}
