package opensearch

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/SHA65536/WikipediaGoApi/region"
	"github.com/stretchr/testify/assert"
)

//go:embed test_data.json
var testOpenSearchJson []byte

var expectedOpenSearchJson = OpenSearchResult{
	Query: "Te",
	Titles: []string{
		"Te",
		"Television show",
		"Texas",
		"Tesla, Inc.",
		"Tennessee",
		"Ted Bundy",
		"Ted Kaczynski",
		"Tencent",
		"Tennis",
		"Telugu language",
	},
	Links: []string{
		"https://en.wikipedia.org/wiki/Te",
		"https://en.wikipedia.org/wiki/Television_show",
		"https://en.wikipedia.org/wiki/Texas",
		"https://en.wikipedia.org/wiki/Tesla,_Inc.",
		"https://en.wikipedia.org/wiki/Tennessee",
		"https://en.wikipedia.org/wiki/Ted_Bundy",
		"https://en.wikipedia.org/wiki/Ted_Kaczynski",
		"https://en.wikipedia.org/wiki/Tencent",
		"https://en.wikipedia.org/wiki/Tennis",
		"https://en.wikipedia.org/wiki/Telugu_language",
	},
}

func TestUnmarshalOpenSearchResult(t *testing.T) {
	var Result OpenSearchResult
	assert := assert.New(t)
	err := json.Unmarshal(testOpenSearchJson, &Result)
	assert.Nil(err, "should not error parsing example")
	assert.Equal(expectedOpenSearchJson, Result, "should equal expected struct")
}

var expectedOpenSearchQuery = "https://en.wikipedia.org/w/api.php?action=opensearch&format=json&formatversion=2&limit=11&profile=fuzzy&redirects=resolve&search=Te"

var inputQueryArgs = OpenSearchArgs{
	Query:    "Te",
	Limit:    11,
	Profile:  Fuzzy,
	Redirect: Resolve,
}

func TestOpenSearchQueryToString(t *testing.T) {
	assert := assert.New(t)
	val, err := inputQueryArgs.ToQuery(region.English)
	assert.Nil(err, "should not error parsing query")
	assert.Equal(expectedOpenSearchQuery, val)
}
