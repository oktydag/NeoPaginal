package services

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"neopaginal-command/configs"
	"neopaginal-command/domain/entity"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}

var crawlUrl url.URL
var appConfig configs.AppConfig
var boundServices BoundService

func setUp() {
	crawlUrl.Host = "https://www.cnnturk.com"
	appConfig = MockInitializeAppConfig()
	boundServices = NewBoundService(appConfig)
}

func TestBoundShouldBeDefaultDueToInitialUrl(t *testing.T) {
	crawlUrl.Path = "/tv"

	var bountTypeOf, _ = boundServices.BoundDecider(&crawlUrl)
	assert.Equal(t, entity.InitialWithoutBound, bountTypeOf, "Bound is not calculated correctly")

}

func TestBoundShouldBeInBound(t *testing.T) {
	var bountTypeOf, _ = boundServices.BoundDecider(&crawlUrl)
	assert.True(t, bountTypeOf == entity.InBound, "Bound is not calculated correctly")
}

func TestBoundShouldBeOutBound(t *testing.T) {
	crawlUrl.Host = "https://www.cnnteeeeurk.com"
	crawlUrl.Path = "/treter"

	var bountTypeOf, _ = boundServices.BoundDecider(&crawlUrl)
	assert.True(t, bountTypeOf == entity.OutBound, "Bound is not calculated correctly")
}

func MockInitializeAppConfig() configs.AppConfig {
	viper.SetConfigName(fmt.Sprintf("config.%s", "QA"))
	viper.AddConfigPath("../..")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var appConfig configs.AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return appConfig
}
