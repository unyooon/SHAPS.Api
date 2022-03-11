package dto

type CreateSubscriptionRequest struct {
	Name        string `column:"name" binding:"required"`
	Price       int    `column:"price" binding:"required"`
	Term        uint   `column:"term" binding:"required"`
	Description string `column:"description" binding:"required"`
}
