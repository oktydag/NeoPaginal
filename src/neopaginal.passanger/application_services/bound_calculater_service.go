package application_services

import (
	"golang.org/x/net/context"
	"neopaginal-passanger/entity/mongo"
	"neopaginal-passanger/repositories"
)

func CalculateBound(crawlMongoRepository repositories.CrawlMongoRepository, urls []string) (int, int) {
	var inBoundCount = 0
	var outBoundCount = 0

	for _, url := range urls {
		var crawl, _ = crawlMongoRepository.GetByCrawledUrl(context.Background(), url)

		if crawl == nil {
			continue
		}

		if crawl.BoundType == mongo.InBound {
			inBoundCount++
		} else if crawl.BoundType == mongo.OutBound {
			outBoundCount++
		}
	}

	return inBoundCount, outBoundCount
}
