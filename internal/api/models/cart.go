package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    int     `json:"user_id"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Status    bool    `json:"status"`
}
