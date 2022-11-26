package query

import (
	"testing"

	"github.com/SHA65536/WikipediaGoApi/region"
	"github.com/stretchr/testify/assert"
)

var inputQueryAllCont = "\"Boss\"_Tweed"
var expectedAllQueryUrl = "https://en.wikipedia.org/w/api.php?action=query&format=json&formatversion=2&gapcontinue=%22Boss%22_Tweed&gapfilterredir=nonredirects&gaplimit=max&gapnamespace=0&generator=allpages"

func TestQueryAllToString(t *testing.T) {
	assert := assert.New(t)
	val, err := AllQueryToUrl(region.English, inputQueryAllCont)
	assert.Nil(err, "should not error parsing query")
	assert.Equal(expectedAllQueryUrl, val)
}
