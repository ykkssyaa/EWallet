package configs

import "github.com/spf13/viper"

func InitConfig() error {

	viper.SetConfigFile("config/.env")

	err := viper.ReadInConfig()
	return err
}
