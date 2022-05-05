package lib

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func NewDatabase(cfg config.DatabaseConfig) Database {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		panic("error connecting to database")
	}
	db.AutoMigrate(&dao.User{})
	db.AutoMigrate(&dao.Item{})
	db.AutoMigrate(&dao.Content{})
	db.AutoMigrate(&dao.DayMenu{})
	db.AutoMigrate(&dao.Menu{})
	db.AutoMigrate(&dao.Article{})
	db.AutoMigrate(&dao.Transaction{})
	return Database{
		Db: db,
	}
}
