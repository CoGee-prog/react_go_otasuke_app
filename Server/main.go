package main

import (
	"flag"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/server"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Parse()

	// 日本時間に設定
	time.Local = time.FixedZone("JST", 9*60*60)

	config.Init(*env)
	database.Init()
	database.Migration(models.OpponentRecruiting{})
	defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
