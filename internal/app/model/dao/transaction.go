package dao

import (
	transaction_status "github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao/trxStatus"
	"gorm.io/gorm"
)

type TransactionItem struct {
	gorm.Model
	Menu          Menu `gorm:"foreignKey:MenuID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsMorning     bool `gorm:"default:false"`
	IsNoon        bool `gorm:"default:false"`
	IsAfternoon   bool `gorm:"default:false"`
	MenuID        uint
	TransactionID uint
	Count         int
}

type Transaction struct {
	gorm.Model
	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID"`
	Status           transaction_status.TransactionStatus
	Amount           int
	Address          string `gorm:"size:255"`
	UserID           uint
}
