package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID          string      `column:"id" gorm:"primary_key"`
	DisplayName string      `column:"display_name"`
	Icon        string      `column:"icon"`
	Host        []Host      `column:"host"`
	Construct   []Construct `column:"contract"`
}
