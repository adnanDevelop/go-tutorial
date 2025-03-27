package utils

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	TotalData   int `json:"total_data"`
}

type Response struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"` // Array of objects
	Pagination *Pagination `json:"pagination"`
}

type ShortResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type BadRequest struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
