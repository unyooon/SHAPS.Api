package dto

type CreateStripeConnectRequest struct {
	Mcc                   string `json:"mcc" binding:"required"`
	BuisinessUrl          string `json:"buisinessUrl" binding:"required"`
	ProductionDescription string `json:"productionDescription" binding:"required"`
	FirstName             string `json:"firstName" binding:"required"`
	LastName              string `json:"lastName" binding:"required"`
	FirstNameKana         string `json:"firstNameKana" binding:"required"`
	LastNameKana          string `json:"lastNameKana" binding:"required"`
	DobDay                int64  `json:"dobDay" binding:"required"`
	DobMonth              int64  `json:"dobMonth" binding:"required"`
	DobYear               int64  `json:"dobYear" binding:"required"`
	PostalCode            string `json:"postalCode" binding:"required"`
	Line1                 string `json:"line1" binding:"required"`
	Line2                 string `json:"line2" binding:"required"`
	Line2Kana             string `json:"line2Kana" binding:"required"`
	Email                 string `json:"email" binding:"required"`
	Phone                 string `json:"phone" binding:"required"`
}
