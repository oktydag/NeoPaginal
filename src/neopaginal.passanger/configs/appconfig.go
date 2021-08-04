package configs

type AppConfig struct {
	ApplicationSettings ApplicationSettings
}

type ApplicationSettings struct {
	MongoDBAddress      string
	MongoDBDatabaseName string
	ElasticNodes        []string
	ElasticIndexName    string
	ChunkSize           int
}
