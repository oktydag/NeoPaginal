const esClient = require('../database/elasticsearchClient')
const config = require('../config.json')

class CrawlService {
    static async GetCrawlsWithBoundCounts() {
        const body = { }

        try {
            var source = ["CrawledUrl", "InBoundCount", "OutBoundCount", "Urls"]
            const resp = await searchDoc(config.ApplicationSettings.ElasticIndexName, body,source );
            var result = resp.hits.hits.map(hit => hit._source);

            return result;
        } catch (e) {
            console.log(e);
        }
    }

    static async GetCrawlByUrl(url) {
        const body = {
            query: {
                match_phrase: {
                    "CrawledUrl.keyword": url
                }
            }
        }
        try {
            var source = ["CrawledUrl", "InBoundCount", "OutBoundCount"]
            const searchResponse = await searchDoc(config.ApplicationSettings.ElasticIndexName, body, source );

            var result = searchResponse.hits.hits.map(hit => hit._source);

            return result;
        } catch (e) {
            console.log(e);
        }
    }
}


const searchDoc = async function(indexName, payload, source){
    return await esClient.search({
        index: indexName,
        body: payload,
        _source : source,
        size : 1000
    });
}


module.exports = CrawlService