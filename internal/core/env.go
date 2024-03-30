package core

import (
	"fmt"

	"github.com/spf13/viper"
)

func Env() {
	// initial config
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error read config file : %w", err))
	}
}
