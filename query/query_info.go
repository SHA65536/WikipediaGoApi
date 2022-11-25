package query

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	ns "github.com/SHA65536/WikipediaGoApi/namespace"
	"github.com/SHA65536/WikipediaGoApi/region"
)

type QueryInfoResult struct {
	Pages []QueryInfoResultPage
}

type QueryInfoResultPage struct {
	Id        int                      `json:"pageid"`
	Ns        ns.Namespace             `json:"ns"`
	Title     string                   `json:"title"`
	Thumbnail QueryInfoResultThumbnail `json:"thumbnail"`
	Touched   time.Time                `json:"touched"`
	URL       string                   `json:"canonicalurl"`
	Extract   string                   `json:"extract"`
}

type QueryInfoResultThumbnail struct {
	Source string `json:"source"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (r *QueryInfoResult) UnmarshalJSON(data []byte) error {
	var Temp = struct {
		Query struct {
			Res []QueryInfoResultPage `json:"pages"`
		} `json:"query"`
	}{}
	err := json.Unmarshal(data, &Temp)
	if err != nil {
		return err
	}
	r.Pages = Temp.Query.Res
	return nil
}

func InfoQueryToUrl(base region.Region, titles []string) (string, error) {
	var values = url.Values{}
	res, err := url.Parse(string(base))
	if err != nil {
		return "", err
	}

	values.Set("action", "query")
	values.Set("prop", "pageimages|info|extracts")
	values.Set("exchars", "256")
	values.Set("pithumbsize", "100")
	values.Set("inprop", "url")
	values.Set("formatversion", "2")
	values.Set("format", "json")
	values.Set("explaintext", "true")
	values.Set("exintro", "true")

	if len(titles) == 0 {
		return "", fmt.Errorf("at least 1 title must be provided")
	}
	values.Set("titles", strings.Join(titles, "|"))

	res.RawQuery = values.Encode()
	return res.String(), nil
}
