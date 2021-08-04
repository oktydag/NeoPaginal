package mongo

type Crawl struct {
	Id         string    `json:"Id" bson:"_id"`
	GroupId    string    `json:"GroupId"`
	Step       int       `json:"Step"`
	BoundType  BoundType `json:"BoundType"`
	Urls       []string  `json:"Urls"`
	CrawledUrl string    `json:"CrawledUrl"`
	IsCrawled  bool      `json:"IsCrawled"`
}

type BoundType string

const (
	InitialWithoutBound BoundType = "InitialWithoutBound"
	InBound             BoundType = "InBound"
	OutBound            BoundType = "OutBound"
)
