package dto

type CartRequest struct {
	ProductID int `json:"product_id" gorm:"type: int"`
	UserID    int `json:"user_id"`
}
