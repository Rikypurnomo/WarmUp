package models_server

type Response struct {
	Meta
	Data interface{} `json:"data,omitempty"`
}

type ResponsePagination struct {
	Meta
	Pagination MetaPagination `json:"pagination"`
	Data       interface{}    `json:"data"`
}

type Meta struct {
	Status       bool   `json:"status"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type MetaPagination struct {
	CurrentPage int64 `json:"current_page"`
	NextPage    int64 `json:"next_page"`
	PrevPage    int64 `json:"prev_page"`
	TotalPages  int64 `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
}
