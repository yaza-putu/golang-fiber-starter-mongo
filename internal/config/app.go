package config

import "github.com/spf13/viper"

type app struct {
	Name   string
	Debug  bool
	MaxCpu int
	ENV    string
	Key    string
	Port   int
	Lang   string
}

func App() app {
	return app{
		Name:   viper.GetString("APP_NAME"),
		Debug:  viper.GetBool("APP_DEBUG"),
		MaxCpu: viper.GetInt("APP_MAX_CPU"),
		ENV:    viper.GetString("APP_ENV"),
		Key:    viper.GetString("APP_KEY"),
		Port:   viper.GetInt("APP_PORT"),
		Lang:   viper.GetString("APP_LANG"),
	}
}
