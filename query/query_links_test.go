package query

import (
	"testing"

	"github.com/SHA65536/WikipediaGoApi/region"
	"github.com/stretchr/testify/assert"
)

var inputQueryLinkTitle = "Turtle"
var inputQueryLinkCont = "37751|0|Painted_turtle"
var expectedInfoQueryUrl = "https://en.wikipedia.org/w/api.php?action=query&format=json&formatversion=2&plcontinue=37751%7C0%7CPainted_turtle&pllimit=max&plnamespace=0&prop=links&titles=Turtle"

func TestQueryLinksToString(t *testing.T) {
	assert := assert.New(t)
	val, err := LinksQueryToUrl(region.English, inputQueryLinkTitle, inputQueryLinkCont)
	assert.Nil(err, "should not error parsing query")
	assert.Equal(expectedInfoQueryUrl, val)
}
