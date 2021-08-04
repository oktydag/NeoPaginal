package application_services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrawlerAsRecursiveShouldNotBeNil(t *testing.T) {

	var actualCrawlList, expectedCrawlList, step = MockInitializeForResursive()

	var recursiveCrawledSlice func([]string, int)

	recursiveCrawledSlice = func(appendedCrawlList []string, step int) {

		if step > appConfig.ApplicationSettings.MaxCrawlLimit {
			return
		} else {

			var newCrawledUrls, newStep = expectedCrawlList, step + 1
			recursiveCrawledSlice(newCrawledUrls, newStep)
			assert.NotNil(t, newCrawledUrls, "new list should not be nil")
		}
	}

	recursiveCrawledSlice(actualCrawlList, step)
}

func TestCrawlerAsRecursiveActualListShouldBeSameAsExpected(t *testing.T) {

	var actualCrawlList, expectedCrawlList, step = MockInitializeForResursive()

	var recursiveCrawledSlice func([]string, int)

	recursiveCrawledSlice = func(appendedCrawlList []string, step int) {

		if step > appConfig.ApplicationSettings.MaxCrawlLimit {
			return
		} else {

			var newCrawledUrls, newStep = expectedCrawlList, step + 1
			recursiveCrawledSlice(newCrawledUrls, newStep)

			assert.EqualValues(t, expectedCrawlList, newCrawledUrls, "new list should be as expected")
		}
	}

	recursiveCrawledSlice(actualCrawlList, step)
}

func TestCrawlerAsRecursiveStepShouldBeSameAsExpected(t *testing.T) {

	var actualCrawlList, expectedCrawlList, step = MockInitializeForResursive()

	var recursiveCrawledSlice func([]string, int)

	recursiveCrawledSlice = func(appendedCrawlList []string, step int) {

		if step > appConfig.ApplicationSettings.MaxCrawlLimit {
			return
		} else {

			var newCrawledUrls, newStep = expectedCrawlList, step + 1
			recursiveCrawledSlice(newCrawledUrls, newStep)

			assert.True(t, step+1 == newStep, "step should be as expected")
		}
	}

	recursiveCrawledSlice(actualCrawlList, step)
}

func MockInitializeForResursive() ([]string, []string, int) {

	var actualCrawlList = make([]string, 0)
	actualCrawlList = append(actualCrawlList, "https://www.facebook.com/cnnturk")
	actualCrawlList = append(actualCrawlList, "https://www.twitter.com/cnnturk")
	actualCrawlList = append(actualCrawlList, "https://www.instgram.com/cnnturk")

	var expectedCrawlList = make([]string, 0)
	expectedCrawlList = append(expectedCrawlList, "https://www.hepsiburada.com/televizyon")
	expectedCrawlList = append(expectedCrawlList, "https://www.hurriyet.com.tr/cnnturk")

	var step = 1
	return actualCrawlList, expectedCrawlList, step
}
