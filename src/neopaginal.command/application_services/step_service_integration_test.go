package application_services

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"neopaginal-command/configs"
	"neopaginal-command/domain"
	"neopaginal-command/domain/entity"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}

var appConfig configs.AppConfig
var firstStepUrl = ""
var initialStep = 0
var groupId = uuid.NewV4().String()
var crawlRepository domain.CrawlRepository

func setUp() {
	appConfig = MockInitializeAppConfig()
	firstStepUrl = "https://www.cnnturk.com"
	initialStep = 1
	crawlRepository = domain.InitializeRepository(appConfig)
}

func TestGeneratedStepWorkerShouldNotReturnNil(t *testing.T) {
	var stepUrls, crawl = StepWorker(appConfig, firstStepUrl, initialStep, groupId)

	assert.NotNilf(t, stepUrls,"Urls should not be nil")
	assert.NotNilf(t, crawl,"Urls should not be nil")
}

func TestGeneratedCrawlShouldBeAsExpected(t *testing.T) {
	var crawledUrls = make([]string, 0)
	crawledUrls = append(crawledUrls, "https://www.cnnturk.com/tv")
	crawledUrls = append(crawledUrls, "https://www.facebook.com/cnnturk")
	crawledUrls = append(crawledUrls, "https://www.twitter.com/cnnturk")

	crawlAsExpected := &entity.Crawl{
		Step:      1,
		GroupId:   groupId,
		BoundType: entity.InBound,
		Urls:      crawledUrls,
		CrawledUrl: firstStepUrl,
		IsCrawled: true,
	}
	var stepUrls, crawl = StepWorker(appConfig, firstStepUrl, initialStep, groupId)

	assert.EqualValues(t, crawl, crawlAsExpected, "crawl should be expected struct !")
	assert.NotNilf(t, stepUrls,"Urls should not be nil")
}


func TestAnalyzeCrawlByUrlsShouldNotBeNil(t *testing.T) {

	var appendedCrawlList, step = InitializeForAnalyzeCrawlByUrlsPayload()

	var newCrawledUrls, newStep = AnalyzeCrawlByUrls(appendedCrawlList, step, appConfig, crawlRepository, groupId)

	assert.NotNilf(t, newCrawledUrls,"Urls should not be nil")
	assert.NotNilf(t, newStep,"Urls should not be nil")
}

func TestStepShouldBeGivenEqualAsExpectedAfterAnalyzeCrawlByUrls(t *testing.T) {

	var appendedCrawlList, step = InitializeForAnalyzeCrawlByUrlsPayload()

	var newCrawledUrls, newStep = AnalyzeCrawlByUrls(appendedCrawlList, step, appConfig, crawlRepository, groupId)

	assert.True(t, newStep == 3 , "Step Should increase after analyze for next step")
	assert.NotNilf(t, newCrawledUrls,"Urls should not be nil")
}

func TestUrlsShouldBeGivenEqualAsExpectedAfterAnalyzeCrawlByUrls(t *testing.T) {

	var appendedCrawlList, step = InitializeForAnalyzeCrawlByUrlsPayload()
	var newCrawledUrls, newStep = AnalyzeCrawlByUrls(appendedCrawlList, step, appConfig, crawlRepository, groupId)

	var expectedCrawlList, exceptedStep = MockForAnalyzeCrawlByUrls()

	assert.EqualValues(t, expectedCrawlList, newCrawledUrls, "new crawl list should be expected struct !")
	assert.EqualValues(t, exceptedStep, newStep, "new step should be expected struct !")
}

func MockInitializeAppConfig() configs.AppConfig {
	viper.SetConfigName(fmt.Sprintf("config.%s", "QA"))
	viper.AddConfigPath("..")

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

func InitializeForAnalyzeCrawlByUrlsPayload() ([]string, int){
	var appendedCrawlList = make([]string, 0)
	appendedCrawlList = append(appendedCrawlList, "https://www.cnnturk.com/tv")
	appendedCrawlList = append(appendedCrawlList, "https://www.facebook.com/cnnturk")
	appendedCrawlList = append(appendedCrawlList, "https://www.twitter.com/cnnturk")

	var step = 2
	return appendedCrawlList, step
}

func MockForAnalyzeCrawlByUrls() ([]string, int){
	var expectedCrawlList = make([]string, 0)
	expectedCrawlList = append(expectedCrawlList, "https://www.facebook.com/cnnturk")
	expectedCrawlList = append(expectedCrawlList, "https://www.twitter.com/cnnturk")
	expectedCrawlList = append(expectedCrawlList, "https://flipboard.com/@CnnTurk")
	expectedCrawlList = append(expectedCrawlList, "https://www.facebook.com/")
	expectedCrawlList = append(expectedCrawlList, "https://www.facebook.com/recover/initiate?lwv=110&ars=royal_blue_bar")
	expectedCrawlList = append(expectedCrawlList, "https://www.facebook.com/privacy/explanation")

	var step = 3
	return expectedCrawlList, step
}
