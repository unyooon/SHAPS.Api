package dto

type CreateSubscriptionRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Term        uint   `json:"term" binding:"required"`
	Description string `json:"description" binding:"required"`
}
