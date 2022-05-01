package dao

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	// ID          uuid.UUID `gorm:"index;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string `gorm:"not null;size:255"`
	ContentID 		uint
}

type Content struct{
	gorm.Model
	Kind 		string `gorm:"size:255"`
	Items 		[]Item `gorm:"foreignKey:ContentID"`
	DayMenuID 	uint
}

type DayMenu struct{
	gorm.Model
	Image 		string `gorm:"size:255"`
	Day 		string `gorm:"index;size:255"`
	Contents 	[]Content `gorm:"foreignKey:DayMenuID"`
	MenuID 		uint
}

type Menu struct{
	gorm.Model
	Title 		string `gorm:"index;size:255"`
	DayMenus 	[]DayMenu `gorm:"foreignKey:MenuID"`
}
