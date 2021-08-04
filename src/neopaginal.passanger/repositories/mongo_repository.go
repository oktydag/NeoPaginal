package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"neopaginal-passanger/configs"
	mongo2 "neopaginal-passanger/entity/mongo"
	mongoClient "neopaginal-passanger/mongo"
	"time"
)

const (
	collectionName_Crawl = "Crawl"
)

type CrawlMongoRepository interface {
	Crawls(ctx context.Context) ([]*mongo2.Crawl, error)
	GetByCrawledUrl(ctx context.Context, crawledUrl string) (*mongo2.Crawl, error )
}

type crawlMongoRepository struct {
	database *mongo.Database
}

func (self *crawlMongoRepository) GetByCrawledUrl(ctx context.Context, crawledUrl string) (*mongo2.Crawl, error) {
	collection := self.database.Collection(collectionName_Crawl)
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	query := bson.M{
		"crawledurl": crawledUrl,
	}

	cur, err := collection.Find(ctxWithTimeout, query)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	var crawl *mongo2.Crawl

	for cur.Next(ctx) {
		var result *mongo2.Crawl
		err = cur.Decode(&result)
		if err := cur.Err(); err != nil {
			return nil, err
		}

		crawl = result
	}

	return crawl, nil}

func (self *crawlMongoRepository) Crawls(ctx context.Context) ([]*mongo2.Crawl, error) {

	collection := self.database.Collection(collectionName_Crawl)
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var crawls []*mongo2.Crawl

	cur, err := collection.Find(ctxWithTimeout, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result *mongo2.Crawl
		err = cur.Decode(&result)
		if err := cur.Err(); err != nil {
			return nil, err
		}

		crawls = append(crawls, result)
	}

	return crawls, nil

}

func NewCrawlRepository(dbAsMongo *mongo.Database) CrawlMongoRepository {

	return &crawlMongoRepository{database: dbAsMongo}
}

func InitializeMongoRepository(appConfig configs.AppConfig) CrawlMongoRepository {
	crawlDbSecondary, err := mongoClient.NewDatabaseWithSecondary(appConfig.ApplicationSettings.MongoDBAddress, appConfig.ApplicationSettings.MongoDBDatabaseName)
	if err != nil {
		panic(err.Error())
	}

	var crawlRepository = NewCrawlRepository(crawlDbSecondary)

	return crawlRepository
}
