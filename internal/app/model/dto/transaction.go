package dto

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	transaction_status "github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao/trxStatus"
)

type (
	CreateTransactionRequestDTO struct {
		Address          string               `json:"address"`
		UserID           uint                 `json:"user_id"`
		TransactionItems []TransactionItemDTO `gorm:"foreignKey:TransactionID"`
	}
	TransactionItemDTO struct {
		IsMorning   bool `json:"is_morning"`
		IsNoon      bool `json:"is_noon"`
		IsAfternoon bool `json:"is_afternoon"`
		MenuID      uint `json:"menu_id"`
		Count       int  `json:"count"`
	}

	TransactionDTO struct {
		ID               uint                                 `json:"id"`
		Status           transaction_status.TransactionStatus `json:"status"`
		Address          string                               `json:"address"`
		UserID           uint                                 `json:"user_id"`
		Amount           int                                  `json:"amount"`
		TransactionItems []TransactionItemDTO                 `gorm:"foreignKey:TransactionID"`
	}

	UpdateTransactionDTO struct {
		ID     uint                                 `json:"id"`
		Status transaction_status.TransactionStatus `json:"status"`
	}
)

func (c CreateTransactionRequestDTO) ToDAO() dao.Transaction {
	transactionItems := []dao.TransactionItem{}

	for _, transactionItem := range c.TransactionItems {
		newTransactionItem := dao.TransactionItem{
			IsMorning:   transactionItem.IsMorning,
			IsNoon:      transactionItem.IsNoon,
			IsAfternoon: transactionItem.IsAfternoon,
			MenuID:      transactionItem.MenuID,
			Count:       transactionItem.Count,
		}

		transactionItems = append(transactionItems, newTransactionItem)
	}

	return dao.Transaction{
		TransactionItems: transactionItems,
		Status:           transaction_status.Pending,
		Address:          c.Address,
		UserID:           c.UserID,
	}
}

func NewTransactionDTO(transaction dao.Transaction) TransactionDTO {
	transactionItems := []TransactionItemDTO{}

	for _, transactionItem := range transaction.TransactionItems {
		newTransactionItem := TransactionItemDTO{
			IsMorning:   transactionItem.IsMorning,
			IsNoon:      transactionItem.IsNoon,
			IsAfternoon: transactionItem.IsAfternoon,
			MenuID:      transactionItem.MenuID,
			Count:       transactionItem.Count,
		}

		transactionItems = append(transactionItems, newTransactionItem)
	}

	return TransactionDTO{
		TransactionItems: transactionItems,
		Status:           transaction_status.Pending,
		Address:          transaction.Address,
		UserID:           transaction.UserID,
	}
}

func ToTransactionDTO(transaction dao.Transaction) TransactionDTO {
	transactionItems := []TransactionItemDTO{}

	for _, transactionItem := range transaction.TransactionItems {
		newTransactionItem := TransactionItemDTO{
			IsMorning:   transactionItem.IsMorning,
			IsNoon:      transactionItem.IsNoon,
			IsAfternoon: transactionItem.IsAfternoon,
			MenuID:      transactionItem.MenuID,
			Count:       transactionItem.Count,
		}

		transactionItems = append(transactionItems, newTransactionItem)
	}

	return TransactionDTO{
		TransactionItems: transactionItems,
		Status:           transaction.Status,
		Address:          transaction.Address,
		UserID:           transaction.UserID,
	}
}
