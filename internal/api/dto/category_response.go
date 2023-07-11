package dto

import "gorm.io/gorm"

type CategoryResponse struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"type: varchar(255)" validate:"required"`
}
