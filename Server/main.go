package main

import (
	"react_go_otasuke_app/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	if err := server.Init(); err != nil {
		panic(err)
	}
}