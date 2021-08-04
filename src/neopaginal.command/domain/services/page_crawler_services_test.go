package services

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var urlForParse *url.URL

func TestParseUrlShouldReturnUrl(t *testing.T) {

	urlForParse = ParseUrl("www.github.com")
	assert.NotEmpty(t, urlForParse, "url should not be empty")
}

func TestParseUrlShouldReturnCorrectUrlDetails(t *testing.T) {
	var href = "https://www.github.com"

	urlForParse = ParseUrl(href)

	assert.EqualValues(t, urlForParse.Hostname(), "www.github.com", "hostname should be equal !")
	assert.EqualValues(t, urlForParse.Scheme, "https", "schema should be equal !")
	assert.False(t, urlForParse.ForceQuery, "ForceQuery should be false !")

}
