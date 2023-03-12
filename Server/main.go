package main

import (
	"flag"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/server"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	env := flag.String("e","development","")
	flag.Parse()

	config.Init(*env)
	database.Init()
	if err := server.Init(); err != nil {
		panic(err)
	}
}