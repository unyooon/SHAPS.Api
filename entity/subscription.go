package entity

import "github.com/jinzhu/gorm"

type Subscription struct {
	gorm.Model
	Name        string `column:"name"`
	Price       int    `column:"price"`
	Term        uint   `column:"term"`
	Description string `column:"description"`
	CreatedBy   string `column:"created_by"`
}
