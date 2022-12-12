package dto

type ReadContractsResponse struct {
	Page BasePageResponse
	Data []ReadContractResponse
}

type ReadContractResponse struct {
	ID          uint   															`json:"id"`
	CreatedAt   string 															`json:"createdAt"`
	Subscription ReadContractsSubscriptionResponse `json:"subscription"`
}

type ReadContractsSubscriptionResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Term        uint   `json:"term"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}