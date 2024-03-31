package config

import "github.com/spf13/viper"

type db struct {
	Host        string
	User        string
	Name        string
	Password    string
	Port        int
	AutoMigrate bool
}

func DB() db {
	return db{
		Host:        viper.GetString("MONGO_HOST"),
		User:        viper.GetString("MONGO_USER"),
		Name:        viper.GetString("MONGO_DATABASE"),
		Password:    viper.GetString("MONGO_PASSWORD"),
		Port:        viper.GetInt("MONGO_PORT"),
		AutoMigrate: viper.GetBool("MONGO_AUTO_MIGRATE"),
	}
}
