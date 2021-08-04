package main

import (
	"context"
	"fmt"
	"neopaginal-passanger/application_services"
	"neopaginal-passanger/configs"
	"neopaginal-passanger/entity/elastic"
	"neopaginal-passanger/entity/mapper"
	"neopaginal-passanger/repositories"
)

func main() {

	var appConfig = configs.InitializeConfigs()

	var crawlMongoRepository = repositories.InitializeMongoRepository(appConfig)
	var crawlElasticRepository = repositories.NewCrawlElasticRepository(appConfig, false)

	var allCrawls, _ = crawlMongoRepository.Crawls(context.Background())

	var elasticBulkDocuments = []*elastic.Crawl{}

	for _, crawl := range allCrawls {

		var inBoundCount, outBoundCount = application_services.CalculateBound(crawlMongoRepository, crawl.Urls)

		var elasticCrawl = mapper.MapMongoCrawlerToElasticCrawler(crawl, inBoundCount, outBoundCount)

		elasticBulkDocuments = append(elasticBulkDocuments, elasticCrawl)

		if len(elasticBulkDocuments) == appConfig.ApplicationSettings.ChunkSize {
			err := crawlElasticRepository.BulkIndex(context.Background(), elasticBulkDocuments)
			if err != nil {
				panic(err)
			}

			elasticBulkDocuments = nil // if not nil, memory keeps allocation.
			elasticBulkDocuments = []*elastic.Crawl{}
		}
	}

	if len(elasticBulkDocuments) > 0 {
		err := crawlElasticRepository.BulkIndex(context.Background(), elasticBulkDocuments)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Migration Completed Successfully.")
}
