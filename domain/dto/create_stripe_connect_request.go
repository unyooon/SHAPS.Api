package dto

type CreateStripeConnectRequest struct {
	Mcc                   string `column:"mcc" binding:"required"`
	BuisinessUrl          string `column:"buisinessUrl" binding:"required"`
	ProductionDescription string `column:"productionDescription" binding:"required"`
	FirstName             string `column:"firstName" binding:"required"`
	LastName              string `column:"lastName" binding:"required"`
	FirstNameKana         string `column:"firstNameKana" binding:"required"`
	LastNameKana          string `column:"lastNameKana" binding:"required"`
	DobDay                int64  `column:"dobDay" binding:"required"`
	DobMonth              int64  `column:"dobMonth" binding:"required"`
	DobYear               int64  `column:"dobYear" binding:"required"`
	PostalCode            string `column:"postalCode" binding:"required"`
	Line1                 string `column:"line1" binding:"required"`
	Line2                 string `column:"line2" binding:"required"`
	Line2Kana             string `column:"line2" binding:"required"`
	Email                 string `column:"email" binding:"required"`
	Phone                 string `column:"phone" binding:"required"`
}
