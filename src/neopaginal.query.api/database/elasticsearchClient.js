const elasticsearch = require("elasticsearch")
const config = require('../config.json')

const esClient = elasticsearch.Client({
    host: config.ApplicationSettings.ElasticNodes
})

module.exports = esClient;