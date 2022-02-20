package dto

type ReadUserResponse struct {
	ID   string  `column:"id"`
	Icon *string `column:"icon"`
}
