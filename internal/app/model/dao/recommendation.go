package dao

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/datatypes"
	"gorm.io/gorm"
)

type Recommendation struct {
	gorm.Model
	Recommendation datatypes.JSONMap
}
