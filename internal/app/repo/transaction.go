package repo

import (
	"errors"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepo interface {
	GetTransactionByID(trxId uint) (dao.Transaction, error)
	GetTransactionByUID(userId uint) ([]dao.Transaction, error) // get by user id
	CreateTransaction(dao.Transaction) (dao.Transaction, error)
	UpdateTransaction(trxId uint, updateMap map[string]interface{}) (dao.Transaction, error)
}
type transactionRepoParams struct {
	fx.In
	lib.Database
}

func NewTransactionRepo(params transactionRepoParams) TransactionRepo {
	return &params
}
func (t transactionRepoParams) GetTransactionByID(trxId uint) (dao.Transaction, error) {
	trx := dao.Transaction{}

	if err := t.Db.Preload("Menus.DayMenus").Preload(clause.Associations).First(&trx, trxId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Transaction{}, constants.DatabaseRecordNotFound
		}

		return dao.Transaction{}, err
	}
	return trx, nil
}

func (t transactionRepoParams) GetTransactionByUID(userId uint) ([]dao.Transaction, error) {
	trx := []dao.Transaction{}

	if err := t.Db.Order("created_at").Preload("Menu.DayMenus").Preload(clause.Associations).Where("user_id = ? ", userId).Find(&trx).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dao.Transaction{}, constants.DatabaseRecordNotFound
		}

		return []dao.Transaction{}, err
	}
	return trx, nil
}

func (t transactionRepoParams) CreateTransaction(trx dao.Transaction) (dao.Transaction, error) {
	if err := t.Db.Create(&trx).Error; err != nil {
		return dao.Transaction{}, err
	}

	return trx, nil
}

func (p *transactionRepoParams) UpdateTransaction(trxId uint, updateMap map[string]interface{}) (dao.Transaction, error) {
	if err := p.Db.Model(&dao.Transaction{}).Where("id = ?", trxId).Updates(updateMap).Error; err != nil {
		return dao.Transaction{}, err
	}

	return p.GetTransactionByID(trxId)

}
