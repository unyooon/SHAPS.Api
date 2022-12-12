package dto

type ReadConstructsResponse struct {
	Page BasePageResponse
	Data []ReadConstructResponse
}

type ReadConstructResponse struct {
	ID          uint   															`json:"id"`
	CreatedAt   string 															`json:"createdAt"`
	Subscription ReadConstructsSubscriptionResponse `json:"subscription"`
}

type ReadConstructsSubscriptionResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Term        uint   `json:"term"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}