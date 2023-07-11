package dto

type History struct {
	Name          string `json:"name"`
	User          string `json:"user" gorm:"column:full_name"`
	TransactionID int    `json:"transaction_id" gorm:"column:transaction_id"`
	CategoryName  string `json:"category_name" gorm:"column:category_name"`
	// TotalPrice    int    `json:"total_price"`
}

type ResponseHistory struct {
	History []History `json:"history"`
}
