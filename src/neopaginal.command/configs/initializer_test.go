package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}

var appConfig AppConfig

func setUp() {
	appConfig = MockInitializeAppConfig()
}

func TestAppConfigCouldBeRead(t *testing.T) {
	assert.NotNilf(t, appConfig, "appConfig is nill")
}

func TestValuesShouldBeTrueInConfigs(t *testing.T) {
	assert.EqualValues(t, "https", appConfig.ApplicationSettings.InitialSchema, "schema is not match")
	assert.True(t, appConfig.ApplicationSettings.MaxCrawlLimit == 3, "MaxCrawlLimit is not correct")
	assert.False(t, appConfig.ApplicationSettings.MaxCrawlLimit == 50, "MaxCrawlLimit should be given expected")
}

func MockInitializeAppConfig() AppConfig {
	viper.SetConfigName(fmt.Sprintf("config.%s", "QA"))
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()
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