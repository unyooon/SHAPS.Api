package entity

import "time"

type User struct {
	ID          string         `column:"id" gorm:"primary_key"`
	DisplayName string         `column:"display_name"`
	Icon        string         `column:"icon"`
	Host        []Subscription `column:"host" gorm:"many2many:hosts"`
	Construct   []Subscription `column:"contract" gorm:"many2many:constructs"`
	CreatedAt   time.Time      `column:"created_at"`
	UpdatedAt   time.Time      `column:"updated_at"`
	DeletedAt   *time.Time     `column:"deleted_at"`
}
