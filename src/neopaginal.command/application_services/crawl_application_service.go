package application_services

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"neopaginal-command/configs"
	"neopaginal-command/domain"
)

func Crawl(appConfig configs.AppConfig, crawlRepository domain.CrawlRepository) bool {

	var firstStepUrl = appConfig.ApplicationSettings.InitialSchema + "://" + appConfig.ApplicationSettings.InitialHostName
	var groupId = uuid.NewV4().String()
	var initialStep = 1

	var initialStepUrls = WorkInitialStepCrawl(appConfig, crawlRepository, firstStepUrl, initialStep, groupId)

	CrawlerAsRecursive(appConfig, crawlRepository, initialStepUrls, initialStep+1, groupId)
	return true
}

func WorkInitialStepCrawl(appConfig configs.AppConfig, crawlRepository domain.CrawlRepository, firstStepUrl string, initialStep int, groupId string) []string {
	var initialStepUrls, initialCrawl = StepWorker(appConfig, firstStepUrl, initialStep, groupId)
	crawlRepository.Save(context.Background(), initialCrawl)

	return initialStepUrls
}