package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"neopaginal-command/configs"
	"neopaginal-command/domain/entity"
	mongoClient "neopaginal-command/mongo"
	"time"
)

const (
	collectionName_Crawl = "Crawl"
)

type CrawlRepository interface {
	Save(ctx context.Context, crawl *entity.Crawl) (error)
	IsCrawled(ctx context.Context, crawledUrl string, groupId string) (bool, error)
}

type crawlRepository struct {
	database *mongo.Database
}

func (self *crawlRepository) IsCrawled(ctx context.Context, crawledUrl string, groupId string) (bool, error){

	query := bson.M{
		"crawledurl":            crawledUrl,
		"groupid":              groupId,
	}

	collection := self.database.Collection(collectionName_Crawl)
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cur, err := collection.Find(ctxWithTimeout, query)
	if err != nil {
		return false, err
	}

	defer cur.Close(ctx)

	var crawl *entity.Crawl
	for cur.Next(ctx) {
		var result *entity.Crawl
		err = cur.Decode(&result)
		if err := cur.Err(); err != nil {
			return false, err
		}
		crawl = result
	}

	if crawl == nil {return false, nil}
	return true, nil
}

func (self *crawlRepository) Save(ctx context.Context, crawl *entity.Crawl) error {

	collection := self.database.Collection(collectionName_Crawl)

	_, err := collection.InsertOne(ctx, crawl)

	if err != nil {
		return err
	}

	return nil
}

func NewCrawlRepository(dbAsMongo *mongo.Database) CrawlRepository {

	return &crawlRepository{database: dbAsMongo}
}

func InitializeRepository(appConfig configs.AppConfig) CrawlRepository {
	crawlDbSecondary, err := mongoClient.NewDatabaseWithSecondary(appConfig.ApplicationSettings.MongoDBAddress, appConfig.ApplicationSettings.MongoDBDatabaseName)
	if err != nil {
		panic(err.Error())
	}

	var crawlRepository = NewCrawlRepository(crawlDbSecondary)

	return crawlRepository
}
