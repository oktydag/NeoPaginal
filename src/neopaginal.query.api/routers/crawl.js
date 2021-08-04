const express = require('express')
const router = new express.Router()
const CrawlController = require('../controllers/v1/CrawlController')
const basicAuth = require('../middleware/authentication')

const prefix = "/api/v1"
const crawlByUrl = prefix + "/crawl"
const crawlsWithBoundCounts = prefix + "/crawl/bound-counts"

router.get(crawlByUrl, basicAuth, async (req, res) => {
    try {
        const crawls = await CrawlController.CrawlByUrl(req.query.url)

        if (!crawls || Object.entries(crawls).length === 0) {
            return res.status(404).send()
        }

        res.send(crawls)
    } catch (e) {
        res.status(500).send()
    }
})

router.get(crawlsWithBoundCounts, basicAuth, async (req, res) => {
    try {
        const crawlwWithBounds = await CrawlController.BoundCounts()

        if (!crawlwWithBounds || Object.entries(crawlwWithBounds).length === 0) {
            return res.status(404).send()
        }

        res.send(crawlwWithBounds)
    } catch (e) {
        res.status(500).send()
    }
})

module.exports = router