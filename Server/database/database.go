package database

import (
	"react_go_otasuke_app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var d *gorm.DB

func Init(models ...interface{}) {
	c := config.GetConfig()
	var err error
	d, err := gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic(err)
	}

	d.AutoMigrate(models...)
}

func GetDB() *gorm.DB{
	return d
}

func Close() {
	d.Close()
}