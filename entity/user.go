package entity

import "time"

type User struct {
	ID          string      `column:"id" gorm:"primary_key"`
	DisplayName string      `column:"display_name"`
	Icon        string      `column:"icon"`
	Host        []Host      `column:"host"`
	Construct   []Construct `column:"contract"`
	CreatedAt   time.Time   `column:"created_at"`
	UpdatedAt   time.Time   `column:"updated_at"`
	DeletedAt   *time.Time  `column:"deleted_at"`
}
