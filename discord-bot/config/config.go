package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init(path string) *viper.Viper {
	/*gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/kaidoj/gamestatsbot/discord-bot")*/
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %v \n", err)
	}

	return viper.GetViper()
}
