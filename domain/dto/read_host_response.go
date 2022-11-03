package dto

type ReadHostResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Term        uint   `json:"term"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}
