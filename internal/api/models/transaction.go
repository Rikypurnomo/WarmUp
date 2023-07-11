package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID int  `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`
	// ProductId Hapus
	// ProductID int     `json:"product_id"`
	// Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	// Tambah Cart one to many
	TotalPrice int `json:"total_price"`
}
