package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path + "/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

}
