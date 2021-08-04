package application_services

import (
	"context"
	"neopaginal-command/configs"
	"neopaginal-command/domain"
	"neopaginal-command/domain/entity"
	"neopaginal-command/domain/services"
	"strings"
)

func StepWorker(appConfig configs.AppConfig, url string, step int, groupId string) ([]string, *entity.Crawl){
	var crawledUrls = make([]string, 0)

	var parsedUrl = services.ParseUrl(url)
	var boundService = services.NewBoundService(appConfig)
	var bountTypeOf, _ = boundService.BoundDecider(parsedUrl)

	crawlBuilder := entity.NewCrawlBuilder()
	crawlBuilder.Lives().
		WithBoundType(bountTypeOf).
		Crawled().
		WithStep(step).
		WithGroupId(groupId).
		WithCrawledUrl(url)

	var mechDocument = services.UrlCrawl(url)

	a := mechDocument.ByTag("a")
	for a.Scan() {
		href := a.Attr("href")

		if !IsValidUrl(href, url) {
			continue
		}

		if len(crawledUrls) == appConfig.ApplicationSettings.MaxCrawlLimit {
			break
		}
		crawledUrls = append(crawledUrls, href)
	}

	crawlBuilder.Lives().WithUrls(crawledUrls)
	crawl := crawlBuilder.Build()

	return crawledUrls, crawl
}

func IsValidUrl(href string, initialUrl string) bool {
	if strings.Trim(href, "/") == initialUrl {
		return false
	}
	if strings.Contains(href, "https://") {
		return true
	}
	return false
}


func AnalyzeCrawlByUrls(appendedUrls []string, stepCounter int, appConfig configs.AppConfig, crawlRepository domain.CrawlRepository, groupId string ) ([]string, int ){
	sliceThatWillCrawl := []string{}

	for _, url := range appendedUrls {
		// check url is Crawled already
		var isUrlCrawledAlready, _ = crawlRepository.IsCrawled(context.Background(), url, groupId)
		if isUrlCrawledAlready {
			continue
		}

		var internalCrawledUrls, crawl = StepWorker(appConfig, url, stepCounter, groupId )
		crawlRepository.Save(context.Background(), crawl)

		sliceThatWillCrawl = append(sliceThatWillCrawl,internalCrawledUrls...)
	}

	return sliceThatWillCrawl, stepCounter+1
}