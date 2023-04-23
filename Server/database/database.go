package database

import (
	"react_go_otasuke_app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var d *gorm.DB

func Init() {
	c := config.GetConfig()
	var err error
	d, err = gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic(err)
	}
}

func Migration(models ...interface{}) {
	d.AutoMigrate(models...)
}

func GetDB() *gorm.DB {
	if d == nil {
		Init()
	}

	return d
}

func Close() {
	d.Close()
}
