package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID int  `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`
	TotalPrice int `json:"total_price"`
}
