package repo

import (
	"errors"
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
	CreateContent(content dao.Content) (dao.Content, error)
	CreateItem(item dao.Item) (dao.Item, error)
	//UpdateUser(userId uint64, updateMap map[string]interface{}) (dao.User, error)
	//DeleteMenu(id uint64)()
}

type menuRepoParams struct {
	fx.In
	lib.Database
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

func (p *menuRepoParams) GetAllMenu() ([]dao.Menu, error) {
	menus := []dao.Menu{}

	if err := p.Db.Preload("DayMenus.Contents.Items").Preload(clause.Associations).Find(&menus).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dao.Menu{}, constants.DatabaseRecordNotFound
		}

		return []dao.Menu{}, err
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

func (p *menuRepoParams) CreateContent(content dao.Content) (dao.Content, error) {

	if err := p.Db.Create(&content).Error; err != nil {
		return dao.Content{}, err
	}

	return content, nil

}

func (p *menuRepoParams) CreateItem(item dao.Item) (dao.Item, error) {

	if err := p.Db.Create(&item).Error; err != nil {
		return dao.Item{}, err
	}

	return item, nil

}
