package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var c *viper.Viper

func init() {
	// .envファイルをロードする
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("APP_ENV"))
	c = viper.New()
	c.SetConfigFile("yml")
	c.SetConfigName(os.Getenv("APP_ENV"))
	c.AddConfigPath("config/environments/")
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetConfig() *viper.Viper {
	return c
}
