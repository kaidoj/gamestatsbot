package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() *viper.Viper {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %v \n", err)
	}

	return viper.GetViper()
}
