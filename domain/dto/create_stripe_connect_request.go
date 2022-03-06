package dto

type CreateStripeConnectRequest struct {
	Mcc                   string `column:"mcc"`
	BuisinessUrl          string `column:"buisinessUrl"`
	ProductionDescription string `column:"productionDescription"`
	FirstName             string `column:"firstName"`
	LastName              string `column:"lastName"`
	FirstNameKana         string `column:"firstNameKana"`
	LastNameKana          string `column:"lastNameKana"`
	DobDay                int64  `column:"dobDay"`
	DobMonth              int64  `column:"dobMonth"`
	DobYear               int64  `column:"dobYear"`
	PostalCode            string `column:"postalCode"`
	Line1                 string `column:"line1"`
	Line2                 string `column:"line2"`
	Line2Kana             string `column:"line2"`
	Email                 string `column:"email"`
	Phone                 string `column:"phone"`
}
