package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"fullname"  binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}
