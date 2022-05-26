package dao

import "gorm.io/gorm"

type File struct {
	gorm.Model

	FileCode         string `gorm:"size:127"`
	OriginalFileName string `gorm:"size:255"`
	Extension        string `gorm:"size:15"`
}
