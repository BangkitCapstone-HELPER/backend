package repo

import (
	"errors"

	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	e "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUser(id uint64) (dao.User, error)
	GetUserByEmail(email string) (dao.User, error)
	ChangePassword(id uint64, password string) (dao.User, error)
	CreateUser(user dao.User) (dao.User, error)
	UpdateUser(userId uint64, updateMap map[string]interface{}) (dao.User, error)
}

type userRepoParams struct {
	fx.In

	lib.Database
	lib.Hash
}

func NewUserRepo(params userRepoParams) UserRepo {
	return &params
}

func (p *userRepoParams) GetUser(id uint64) (dao.User, error) {
	user := dao.User{}

	if err := p.Db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.User{}, constants.DatabaseRecordNotFound
		}

		return dao.User{}, err
	}
	return user, nil
}

func (p *userRepoParams) GetUserByEmail(email string) (dao.User, error) {
	user := dao.User{}
	if err := p.Db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.User{}, e.ErrUserEmailNotFound
		}

		return dao.User{}, err
	}
	return user, nil
}
func (p *userRepoParams) ChangePassword(id uint64, password string) (dao.User, error) {

	user := dao.User{}
	if err := p.Db.Model(&dao.User{}).Where("id = ?", id).Updates(map[string]interface{}{"password": password}).Error; err != nil {
		return dao.User{}, err
	}
	return user, nil
}

func (p *userRepoParams) UpdateUser(userId uint64, updateMap map[string]interface{}) (dao.User, error) {
	if err := p.Db.Model(&dao.User{}).Where("id = ?", userId).Updates(updateMap).Error; err != nil {
		return dao.User{}, err
	}

	return p.GetUser(userId)

}

func (p *userRepoParams) CreateUser(user dao.User) (dao.User, error) {

	if err := p.Db.Create(&user).Error; err != nil {
		return dao.User{}, err
	}

	return user, nil

}
