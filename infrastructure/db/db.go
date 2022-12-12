package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"shaps.api/domain/setting"
	"shaps.api/entity"
)

type Db struct {
	Connection *gorm.DB
}

func NewDb(setting setting.Setting) *Db {
	db, err := gorm.Open("mssql", setting.Dsn)
	if err != nil {
		panic(err.Error())
	}

	d := &Db{
		Connection: db,
	}
	d.autoMigration()

	return d
}

func (d *Db) Connect() *gorm.DB {
	return d.Connection
}

func (d *Db) autoMigration() {
	d.Connection.AutoMigrate(&entity.Subscription{})
	d.Connection.AutoMigrate(&entity.User{})
	d.Connection.AutoMigrate(&entity.Contract{})
}
