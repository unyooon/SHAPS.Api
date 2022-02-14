package db

import "github.com/jinzhu/gorm"

type DbInterface interface {
	Connect() *gorm.DB
}
