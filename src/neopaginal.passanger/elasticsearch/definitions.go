package elasticsearch

const CrawlAnalyzerName = "claims_analyzer"

var CrawlIndexDefinition_v1 = `
{
    "settings": {
        "number_of_shards": 3,
        "number_of_replicas": 1,
        "index": {
            "analysis": {
                "analyzer": {
                    "` + CrawlAnalyzerName + `": {
                        "tokenizer": "whitespace",
                        "filter": [
                            "lowercase",
                            "asciifolding"
                        ]
                    }
                }
            }
        }
    },
    "mappings": {
        "properties": {
            "Id": {
                "type": "keyword"
            },
            "GroupId": {
                "type": "keyword"
            },
            "Step": {
                "type": "integer"
            },
            "BoundType": {
                "type": "keyword"
            },
			"CrawledUrl": {
                "type": "keyword",
				"index" : "false"
            },
            "IsCrawled": {
                "type": "boolean"
            },
            "Urls": {
                "type": "keyword"
            },
            "InBoundCount": {
                "type": "integer"
            },
            "OutBoundCount": {
                "type": "integer"
            }
        }
    }
}
`