package repositories

import (
	"bytes"
	"encoding/json"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"golang.org/x/net/context"
	"io/ioutil"
	"neopaginal-passanger/configs"
	"neopaginal-passanger/elasticsearch"
	"neopaginal-passanger/entity/elastic"
	"neopaginal-passanger/errors"
	"strings"
	"time"
)

type (
	Repository interface {
		Index(ctx context.Context, id string, document *elastic.Crawl) error
		BulkIndex(ctx context.Context, documents []*elastic.Crawl) error
	}

	crawlElasticRepository struct {
		elasticClient *elasticsearch7.Client
		indexName     string
	}
)

func (repo *crawlElasticRepository) Index(ctx context.Context, id string, document *elastic.Crawl) error {

	var (
		err      error
		response *esapi.Response
		start    time.Time
		logExtra = make(map[string]interface{})
	)

	req := esapi.IndexRequest{
		Index:      repo.indexName,
		DocumentID: document.Id,
		Body:       esutil.NewJSONReader(&document),
	}

	start = time.Now()

	if response, err = req.Do(ctx, repo.elasticClient); err != nil {
		return err
	}

	logExtra["InsertDuration"] = time.Since(start).Milliseconds()

	defer response.Body.Close()

	if response.IsError() {
		if response.StatusCode == 409 {
			contents, _ := ioutil.ReadAll(response.Body)
			return errors.New(elasticsearch.VersionConflictError, response.Status(), id, string(contents))
		}
		return errors.New(elasticsearch.InsertingError, response.Status(), id)
	}
	return nil
}

func (repo *crawlElasticRepository) BulkIndex(ctx context.Context, documents []*elastic.Crawl) error {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         repo.indexName,
		Client:        repo.elasticClient,
		FlushInterval: 30 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	length := len(documents)
	for i := 0; i < length; i++ {
		document := documents[i]

		data, err := json.Marshal(document)
		if err != nil {
			panic(err)
		}

		err = bi.Add(context.Background(), esutil.BulkIndexerItem{
			Action:          "index",
			DocumentID:      document.Id,
			Body:            bytes.NewReader(data),
			RetryOnConflict: nil,
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
			},
		})
	}

	if err := bi.Close(context.Background()); err != nil {
		panic(err)
	}

	return nil
}

func NewCrawlElasticRepository(appConfig configs.AppConfig, shouldCreateIndex bool) Repository {
	var (
		e7Client *elasticsearch7.Client
		err      error
	)

	cfg := elasticsearch7.Config{Addresses: appConfig.ApplicationSettings.ElasticNodes}

	if e7Client, err = elasticsearch7.NewClient(cfg); err != nil {
		panic(err)
	}

	if shouldCreateIndex {
		if err = EnsureIndex(appConfig.ApplicationSettings.ElasticIndexName, elasticsearch.CrawlIndexDefinition_v1, e7Client); err != nil {
			panic(err)
		}
	}

	return &crawlElasticRepository{elasticClient: e7Client, indexName: appConfig.ApplicationSettings.ElasticIndexName}
}

func EnsureIndex(indexName string, indexDefinition string, e7Client *elasticsearch7.Client) error {
	err, statusCode := CheckIndex(indexName, e7Client)
	if statusCode == 404 {
		err = CreateIndex(indexName, indexDefinition, e7Client)
	}

	return err
}

func CheckIndex(indexName string, e7Client *elasticsearch7.Client) (error, int) {

	req := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}

	res, err := req.Do(context.Background(), e7Client)
	statusCode := res.StatusCode
	defer res.Body.Close()

	return err, statusCode
}

func CreateIndex(indexName string, indexDefinition string, e7Client *elasticsearch7.Client) error {
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(indexDefinition),
	}

	res, err := req.Do(context.Background(), e7Client)
	if res.StatusCode == 400 {

		contents, _ := ioutil.ReadAll(res.Body)

		return errors.New(elasticsearch.IndexMappingError, res.StatusCode, string(contents))
	}
	defer res.Body.Close()

	return err
}
