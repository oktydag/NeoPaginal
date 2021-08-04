const crawlService = require('../../services/crawlService')


class CrawlController{
    static async BoundCounts()
    {
        var crawlsWithBoundsCount = await crawlService.GetCrawlsWithBoundCounts();

        return crawlsWithBoundsCount;
    }

    static async CrawlByUrl(url)
    {
        var crawls = await crawlService.GetCrawlByUrl(url);
        return crawls;
    }
}

module.exports = CrawlController;
