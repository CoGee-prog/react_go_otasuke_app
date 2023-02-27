package config

import (
	"github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
	c = viper.New()
	c.SetConfigFile("yaml")
	c.SetConfigName(env)
	c.AddConfigPath("config/environments/")
	c.AddConfigPath("/run/secrets/")
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetConfig() *viper.Viper{
	return c
}
