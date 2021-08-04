package configs

type AppConfig struct {
	ApplicationSettings ApplicationSettings
}

type ApplicationSettings struct {
	InitialSchema       string
	InitialHostName     string
	MongoDBAddress      string
	MongoDBDatabaseName string
	MaxCrawlLimit		int
}
