package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitializeConfigs() AppConfig {

	environmentName := os.Getenv("ENV")

	viper.SetConfigName(fmt.Sprintf("config.%s", environmentName))
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return appConfig
}
