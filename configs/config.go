package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	ApiToken string `mapstructure:"API_TOKEN"`
}

func InitConfig() (cfg Config) {
	viper.SetConfigFile("./configs/envs/dev.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read config: " + err.Error())
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("failed to unmarshal config: " + err.Error())
	}

	return cfg
}
