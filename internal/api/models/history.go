package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	User          User        `json:"user" gorm:"foreignKey:UserID"`
	UserID        int         `json:"user_id" form:"user_id"`
	TransactionID int         `json:"transaction_id" form:"transaction_id"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
}
