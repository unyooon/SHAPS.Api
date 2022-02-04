package infrastructure

import (
	"os"

	"github.com/jinzhu/gorm"
)

type Db struct {
	Host       string
	UserName   string
	Password   string
	DbName     string
	Connection *gorm.DB
}

func NewDb() *Db {
	d := &Db{
		Host:     os.Getenv("DB_HOST"),
		UserName: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASS"),
		DbName:   os.Getenv("DB_NAME"),
	}

	// TODO: 後で実装
	// db, err := gorm.Open("")
	// d.Connection = db
	d.autoMigration()

	return d
}

func (d *Db) Connect() *gorm.DB {
	return d.Connection
}

func (d *Db) autoMigration() {

}
