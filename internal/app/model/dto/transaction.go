package dto

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	transaction_status "github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao/trxStatus"
	"gorm.io/gorm"
	"time"
)

type (
	CreateTransactionRequestDTO struct {
		Address     string    `json:"address"`
		UserID      uint      `json:"user_id"`
		Amount      int       `json:"amount"`
		Lat         float64   `json:"lat"`
		Lng         float64   `json:"lng"`
		IsMorning   bool      `json:"is_morning"`
		IsNoon      bool      `json:"is_noon"`
		IsAfternoon bool      `json:"is_afternoon"`
		MenuID      uint      `json:"menu_id"`
		Count       int       `json:"count"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Upload      string    `json:"upload"`
	}
	//TransactionItemDTO struct {
	//	IsMorning   bool `json:"is_morning"`
	//	IsNoon      bool `json:"is_noon"`
	//	IsAfternoon bool `json:"is_afternoon"`
	//	MenuID      uint `json:"menu_id"`
	//	Count       int  `json:"count"`
	//}

	TransactionDTO struct {
		ID          uint                                 `json:"id"`
		Status      transaction_status.TransactionStatus `json:"status"`
		Address     string                               `json:"address"`
		UserID      uint                                 `json:"user_id"`
		Amount      int                                  `json:"amount"`
		IsMorning   bool                                 `json:"is_morning"`
		IsNoon      bool                                 `json:"is_noon"`
		IsAfternoon bool                                 `json:"is_afternoon"`
		MenuID      uint                                 `json:"menu_id"`
		Count       int                                  `json:"count"`
		Menu        MenuDTO                              `json:"menu"`
		Lat         float64                              `json:"lat"`
		Lng         float64                              `json:"lng"`
		Remaining   int                                  `json:"remaining"`
		Upload      string                               `json:"upload"`
		CreatedAt   time.Time                            `json:"created_at"`
		UpdatedAt   time.Time                            `json:"updated_at"`
	}

	UpdateTransactionDTO struct {
		ID     uint                                 `json:"id"`
		Status transaction_status.TransactionStatus `json:"status"`
		Upload string                               `json:"upload"`
	}
)

func (c CreateTransactionRequestDTO) ToDAO() dao.Transaction {
	//transactionItems := []dao.TransactionItem{}
	//
	//for _, transactionItem := range c.TransactionItems {
	//	newTransactionItem := dao.TransactionItem{
	//		IsMorning:   transactionItem.IsMorning,
	//		IsNoon:      transactionItem.IsNoon,
	//		IsAfternoon: transactionItem.IsAfternoon,
	//		MenuID:      transactionItem.MenuID,
	//		Count:       transactionItem.Count,
	//	}
	//
	//	transactionItems = append(transactionItems, newTransactionItem)
	//}

	return dao.Transaction{
		IsMorning:   c.IsMorning,
		IsNoon:      c.IsNoon,
		IsAfternoon: c.IsAfternoon,
		MenuID:      c.MenuID,
		Count:       c.Count,
		Status:      transaction_status.Pending,
		Address:     c.Address,
		Upload:      c.Upload,
		Model:       gorm.Model{},
		Menu:        dao.Menu{},
		Amount:      c.Amount,
		Lat:         c.Lat,
		Lng:         c.Lng,
	}
}

func NewTransactionDTO(transaction dao.Transaction) TransactionDTO {
	//transactionItems := []TransactionItemDTO{}
	//
	//for _, transactionItem := range transaction.TransactionItems {
	//	newTransactionItem := TransactionItemDTO{
	//		IsMorning:   transactionItem.IsMorning,
	//		IsNoon:      transactionItem.IsNoon,
	//		IsAfternoon: transactionItem.IsAfternoon,
	//		MenuID:      transactionItem.MenuID,
	//		Count:       transactionItem.Count,
	//	}
	//
	//	transactionItems = append(transactionItems, newTransactionItem)
	//}

	return TransactionDTO{
		IsMorning:   transaction.IsMorning,
		IsNoon:      transaction.IsNoon,
		IsAfternoon: transaction.IsAfternoon,
		MenuID:      transaction.MenuID,
		Count:       transaction.Count,
		Status:      transaction_status.Pending,
		Address:     transaction.Address,
		UserID:      transaction.UserID,
		Amount:      transaction.Amount,
		Remaining:   transaction.Remaining,
		Menu:        NewMenuDTO(transaction.Menu),
		Upload:      transaction.Upload,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}

func ToTransactionDTO(transaction dao.Transaction) TransactionDTO {
	//transactionItems := []TransactionItemDTO{}
	//
	//for _, transactionItem := range transaction.TransactionItems {
	//	newTransactionItem := TransactionItemDTO{
	//		IsMorning:   transactionItem.IsMorning,
	//		IsNoon:      transactionItem.IsNoon,
	//		IsAfternoon: transactionItem.IsAfternoon,
	//		MenuID:      transactionItem.MenuID,
	//		Count:       transactionItem.Count,
	//	}
	//
	//	transactionItems = append(transactionItems, newTransactionItem)
	//}

	return TransactionDTO{
		IsMorning:   transaction.IsMorning,
		IsNoon:      transaction.IsNoon,
		IsAfternoon: transaction.IsAfternoon,
		MenuID:      transaction.MenuID,
		Count:       transaction.Count,
		Status:      transaction.Status,
		Address:     transaction.Address,
		UserID:      transaction.UserID,
		Amount:      transaction.Amount,
		Remaining:   transaction.Remaining,
		Menu:        NewMenuDTO(transaction.Menu),
		Upload:      transaction.Upload,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}
