package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name" form:"name" validate:"required" gorm:"varchar(255)"`
	Price       int      `json:"price" form:"price" validate:"required" gorm:"price"`
	Description string   `json:"description" form:"description" validate:"required" gorm:"description"`
	CategoryID  int      `json:"category_id" form:"category_id"`
	Category    Category `json:"category"  gorm:"foreignkey:CategoryID"`
}
