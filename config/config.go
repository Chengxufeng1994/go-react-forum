package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(configName string) *viper.Viper {
	var err error

	config = viper.New()
	config.SetConfigName(configName)
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")

	err = config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config.WatchConfig()

	return config
}

func GetConfig() *viper.Viper {
	return config
}
