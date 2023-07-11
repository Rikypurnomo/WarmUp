package dto

type CartResponse struct {
	Name         string `json:"name"`
	User         string `json:"full_name" gorm:"column:full_name"`
	Quantity     int    `json:"quantity"`
	CategoryName string `json:"category_name" gorm:"column:category_name"`
}

type ResponseCarts struct {
	CartResponse []CartResponse `json:"cart_response"`
	TotalPrice   int            `json:"total_price"`
}
