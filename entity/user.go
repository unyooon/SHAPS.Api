package entity

import "time"

type User struct {
	ID        string         `column:"id" gorm:"primary_key"`
	StripeId  string         `column:"stripeId"`
	ConnectId string         `column:"connectId"`
	Icon      *string        `column:"icon"`
	Name      *string        `column:"name"`
	Host      []Subscription `column:"host" gorm:"many2many:hosts"`
	Construct []Construct    `column:"construct"`
	CreatedAt time.Time      `column:"created_at"`
	UpdatedAt time.Time      `column:"updated_at"`
	DeletedAt *time.Time     `column:"deleted_at"`
}
