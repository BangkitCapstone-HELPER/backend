package repo

import (
	"errors"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	transaction_status "github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao/trxStatus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	GetTransactionByID(trxId uint) (dao.Transaction, error)
	GetTransactionByUID(userId uint) ([]dao.Transaction, error) // get by user id
	CreateTransaction(dao.Transaction) (dao.Transaction, error)
	UpdateTransaction(trxId uint, status transaction_status.TransactionStatus) (dao.Transaction, error)
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

	if err := t.Db.First(&trx, trxId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Transaction{}, constants.DatabaseRecordNotFound
		}

		return dao.Transaction{}, err
	}
	return trx, nil
}

func (t transactionRepoParams) GetTransactionByUID(userId uint) ([]dao.Transaction, error) {
	trx := []dao.Transaction{}

	if err := t.Db.Where("id_user = ? ", userId).Find(&trx).Error; err != nil {
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

func (p *transactionRepoParams) UpdateTransaction(trxId uint, status transaction_status.TransactionStatus) (dao.Transaction, error) {
	trx := dao.Transaction{
		Model: gorm.Model{
			ID: trxId,
		},
	}
	if err := p.Db.First(&trx).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Transaction{}, constants.DatabaseRecordNotFound
		}
		return dao.Transaction{}, err
	}
	trx.Status = status

	err := p.Db.Save(&trx).Error
	if err != nil {
		return dao.Transaction{}, err
	}
	return trx, nil

}
