package dto

type CreateSubscriptionRequest struct {
	Name        string `column:""`
	Price       int    `column:"price"`
	Term        uint   `column:"term"`
	Description string `column:"description"`
}
