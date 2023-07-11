package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UserID        int         `json:"user_id"`
	User          User        `json:"user" gorm:"foreignKey:UserID"`
	ProductID     int         `json:"product_id"`
	Product       Product     `json:"product" gorm:"foreignKey:ProductID"`
	TransactionID int         `json:"transaction_id"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
	Category      Category    `json:"category" gorm:"foreignKey:CategoryID"`
	CategoryID    int         `json:"category_id"`
}
