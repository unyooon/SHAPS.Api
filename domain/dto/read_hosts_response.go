package dto

type ReadHostsResponse struct {
	Page BasePageResponse   `json:"page"`
	Data []ReadHostResponse `json:"data"`
}
