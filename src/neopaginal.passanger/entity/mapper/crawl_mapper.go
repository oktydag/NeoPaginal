package mapper

import (
	"neopaginal-passanger/entity/elastic"
	"neopaginal-passanger/entity/mongo"
)

func MapMongoCrawlerToElasticCrawler(mongoCrawler *mongo.Crawl, inBoundCount int, outBoundCount int) *elastic.Crawl {

	var crawlForElastic = &elastic.Crawl{
		Id:            mongoCrawler.Id,
		GroupId:       mongoCrawler.GroupId,
		Step:          mongoCrawler.Step,
		CrawledUrl:    mongoCrawler.CrawledUrl,
		BoundType:     string(mongoCrawler.BoundType),
		IsCrawled:     mongoCrawler.IsCrawled,
		Urls:          mongoCrawler.Urls,
		InBoundCount:  inBoundCount,
		OutBoundCount: outBoundCount,
	}

	return crawlForElastic
}
