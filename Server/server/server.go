package server

import (
	"fmt"
	"react_go_otasuke_app/config"
)

func Init() error {
	router, err := NewRouter()
	if err != nil {
		return err
	}
	c := config.GetConfig()
	router.Run(fmt.Sprintf(":%s",c.GetString("server.port")))

	return nil
}
