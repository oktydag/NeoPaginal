package elastic

type Crawl struct {
	Id            string    `json:"Id" bson:"_id"`
	GroupId       string    `json:"GroupId"`
	Step          int       `json:"Step"`
	BoundType     string    `json:"BoundType"`
	Urls          []string  `json:"Urls"`
	CrawledUrl    string    `json:"CrawledUrl"`
	IsCrawled     bool      `json:"IsCrawled"`
	InBoundCount  int      `json:"InBoundCount"`
	OutBoundCount int      `json:"OutBoundCount"`
}