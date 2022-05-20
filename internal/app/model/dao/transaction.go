package dao

import (
	transaction_status "github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao/trxStatus"
	"gorm.io/gorm"
)

//type TransactionItem struct {
//	gorm.Model
//	Menu          Menu `gorm:"foreignKey:MenuID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
//	IsMorning     bool `gorm:"default:false"`
//	IsNoon        bool `gorm:"default:false"`
//	IsAfternoon   bool `gorm:"default:false"`
//	MenuID        uint
//	TransactionID uint
//	Count         int
//}

type Transaction struct {
	gorm.Model
	Menu          Menu `gorm:"foreignKey:MenuID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsMorning     bool `gorm:"default:false"`
	IsNoon        bool `gorm:"default:false"`
	IsAfternoon   bool `gorm:"default:false"`
	MenuID        uint
	TransactionID uint
	Count         int
	Remaining     int `gorm:"default:0"`
	Status        transaction_status.TransactionStatus
	Amount        int
	Address       string `gorm:"size:255"`
	UserID        uint
	Lat           float64 `gorm:"type:decimal(10,8)"`
	Lng           float64 `gorm:"type:decimal(11,8)"`
}
