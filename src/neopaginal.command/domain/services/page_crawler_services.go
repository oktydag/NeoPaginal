package services

import (
	"github.com/89z/mech"
	"net/http"
	"net/url"
)

func UrlCrawl (url string) *mech.Node {
	response, err := http.Get(url)

	if err != nil {
	}

	defer response.Body.Close()
	document, err := mech.Parse(response.Body)

	if err != nil {
		panic(err)
	}

	return document
}

func ParseUrl(href string) *url.URL {
	parsedUrl, err := url.Parse(href)
	if err != nil {
		panic(err)
	}
	return parsedUrl
}
