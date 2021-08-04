package application_services

import (
	"neopaginal-command/configs"
	"neopaginal-command/domain"
)

func CrawlerAsRecursive (appConfig configs.AppConfig, crawlRepository domain.CrawlRepository, crawledList []string, step int,  groupId string) {

	var recursiveCrawledSlice func([]string, int)

	recursiveCrawledSlice = func(appendedCrawlList []string, step int) {

		if step > appConfig.ApplicationSettings.MaxCrawlLimit {
			return
		} else {

			var newCrawledUrls, newStep = AnalyzeCrawlByUrls(appendedCrawlList, step, appConfig, crawlRepository, groupId)
			recursiveCrawledSlice(newCrawledUrls, newStep)
		}
	}

	recursiveCrawledSlice(crawledList, step)
}
