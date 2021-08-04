package entity

type Crawl struct {
	GroupId    string    `json:"GroupId"`
	Step       int       `json:"Step"`
	BoundType  BoundType `json:"BoundType"`
	Urls       []string  `json:"Urls"`
	CrawledUrl string    `json:"CrawledUrl"`
	IsCrawled  bool      `json:"IsCrawled"`
}

type CrawlBuilder struct {
	crawl *Crawl
}

type CrawlWorkerBuilder struct {
	CrawlBuilder
}

func NewCrawlBuilder() *CrawlBuilder {
	return &CrawlBuilder{crawl: &Crawl{}}
}

func (c *CrawlBuilder) Lives() *CrawlWorkerBuilder {
	return &CrawlWorkerBuilder{*c}
}

func (a *CrawlWorkerBuilder) WithGroupId(groupId string) *CrawlWorkerBuilder {
	a.crawl.GroupId = groupId
	return a
}


func (a *CrawlWorkerBuilder) WithStep(step int) *CrawlWorkerBuilder {
	a.crawl.Step = step
	return a
}

func (a *CrawlWorkerBuilder) WithBoundType(boundType BoundType) *CrawlWorkerBuilder {
	a.crawl.BoundType = boundType
	return a
}

func (a *CrawlWorkerBuilder) WithUrls(urls []string) *CrawlWorkerBuilder {
	a.crawl.Urls = urls
	return a
}

func (a *CrawlWorkerBuilder) WithCrawledUrl(crawledUrl string) *CrawlWorkerBuilder {
	a.crawl.CrawledUrl = crawledUrl
	return a
}

func (a *CrawlWorkerBuilder) Crawled() *CrawlWorkerBuilder {
	a.crawl.IsCrawled = true
	return a
}

func (b *CrawlBuilder) Build() *Crawl {
	return b.crawl
}

