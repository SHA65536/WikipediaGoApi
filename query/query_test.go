package query

import (
	_ "embed"
	"encoding/json"
	"testing"
	"time"

	"github.com/SHA65536/WikipediaGoApi/namespace"
	"github.com/SHA65536/WikipediaGoApi/region"
	"github.com/stretchr/testify/assert"
)

//go:embed test_data.json
var testQueryJson []byte

var expectedOpenSearchJson = QueryResult{
	[]QueryResultPage{
		{
			Id:    736,
			Ns:    namespace.Main,
			Title: "Albert Einstein",
			Thumbnail: QueryResultThumbnail{
				Source: "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Einstein_1921_by_F_Schmutzer_-_restoration.jpg/76px-Einstein_1921_by_F_Schmutzer_-_restoration.jpg",
				Width:  76,
				Height: 100,
			},
			Touched: time.Date(2022, time.November, 25, 14, 45, 30, 0, time.UTC),
			URL:     "https://en.wikipedia.org/wiki/Albert_Einstein",
			Extract: "Albert Einstein ( EYEN-styne; German: [ˈalbɛʁt ˈʔaɪnʃtaɪn] (listen); 14 March 1879 – 18 April 1955) was a German-born theoretical physicist, widely acknowledged to be one of the greatest and most influential physicists of all time. Einstein is best known for...",
		},
		{
			Id:    25409,
			Ns:    namespace.Main,
			Title: "Reptile",
			Thumbnail: QueryResultThumbnail{
				Source: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/43/Reptiles_2021_collage.jpg/100px-Reptiles_2021_collage.jpg",
				Width:  100,
				Height: 91,
			},
			Touched: time.Date(2022, time.November, 24, 20, 15, 13, 0, time.UTC),
			URL:     "https://en.wikipedia.org/wiki/Reptile",
			Extract: "Reptiles, as most commonly defined are the animals in the class Reptilia ( rep-TIL-ee), a grouping comprising all sauropsids except birds. Living reptiles comprise turtles, crocodilians, squamates (lizards and snakes) and rhynchocephalians (tuatara). As of...",
		},
	},
}

func TestUnmarshalQueryResult(t *testing.T) {
	var Result QueryResult
	assert := assert.New(t)
	err := json.Unmarshal(testQueryJson, &Result)
	assert.Nil(err, "should not error parsing example")
	assert.Equal(expectedOpenSearchJson, Result, "should equal expected struct")
}

var expectedQueryUrl = "https://en.wikipedia.org/w/api.php?action=query&exchars=256&exintro=true&explaintext=true&format=json&formatversion=2&inprop=url&pithumbsize=100&prop=pageimages%7Cinfo%7Cextracts&titles=Albert+Einstein%7CReptile"

var inputQueryTitles = []string{"Albert Einstein", "Reptile"}

func TestQueryQueryToString(t *testing.T) {
	assert := assert.New(t)
	val, err := QueryToUrl(region.English, inputQueryTitles)
	assert.Nil(err, "should not error parsing query")
	assert.Equal(expectedQueryUrl, val)
}
