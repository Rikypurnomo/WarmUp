package dto

import "gorm.io/gorm"

type ProductResponse struct {
	gorm.Model
	Name         string `json:"name" form:"name" validate:"required" gorm:"varchar(255)"`
	Price        int    `json:"price" form:"price" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	CategoryName string `json:"category_name" gorm:"column:category_name"`
}
