package dao

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/datatypes"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Article datatypes.JSONMap
}
