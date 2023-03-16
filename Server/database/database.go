package database

import (
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/models"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var d *gorm.DB

func Init() {
	c := config.GetConfig()
	var err error
	d, err := gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic(err)
	}

	allModels := getModles(models.AllModels{})
	d.AutoMigrate(allModels)
}

func getModles(models models.AllModels) []string {
	rtModels := reflect.TypeOf(models)
	allModels := make([]string, rtModels.NumField())
	for i := 0; i < rtModels.NumField(); i++ {
		allModels[i] = rtModels.Field(i).Name
	}
	return allModels
}
