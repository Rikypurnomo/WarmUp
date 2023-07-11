package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
