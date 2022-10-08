package handlers

type Item struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `swaggerignore:"true"`
}
