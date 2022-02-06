package infrastructure

import "github.com/jinzhu/gorm"

type DbInterface interface {
	Connect() *gorm.DB
}
