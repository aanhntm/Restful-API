package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DEFINED_DB_DRIVER"`
	ServerAddress string `mapstructure:"DEFINED_SERVER_ADDRESS"`
	DBSource      string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("app")
	v.SetConfigType("env")

	v.AutomaticEnv()
	v.ReadInConfig()

	err = v.Unmarshal(&config)
	return
}
