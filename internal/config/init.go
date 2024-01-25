package config

import "github.com/spf13/viper"

func InitConfig() error {

	viper.SetConfigFile("internal/config/.env")

	err := viper.ReadInConfig()
	return err
}
