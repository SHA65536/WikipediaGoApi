package query

import (
	"net/url"

	"github.com/SHA65536/WikipediaGoApi/namespace"
	"github.com/SHA65536/WikipediaGoApi/region"
)

type QueryLinksResult struct {
	Cont  QueryLinksContinue `json:"continue"`
	Query QueryLinksQuery    `json:"query"`
}

type QueryLinksQuery struct {
	Pages []QueryLinksPage `json:"pages"`
}

type QueryLinksPage struct {
	Id    int                 `json:"pageid"`
	Ns    namespace.Namespace `json:"ns"`
	Title string              `json:"title"`
	Links []QueryLink         `json:"Links"`
}

type QueryLink struct {
	Ns    namespace.Namespace `json:"ns"`
	Title string              `json:"title"`
}

type QueryLinksContinue struct {
	Next string `json:"plcontinue"`
}

func LinksQueryToUrl(base region.Region, title string, cont string) (string, error) {
	var values = url.Values{}
	res, err := url.Parse(string(base))
	if err != nil {
		return "", err
	}

	values.Set("action", "query")
	values.Set("prop", "links")
	values.Set("plnamespace", "0")
	values.Set("formatversion", "2")
	values.Set("format", "json")
	values.Set("pllimit", "max")
	values.Set("titles", title)

	if cont != "" {
		values.Set("plcontinue", cont)
	}

	res.RawQuery = values.Encode()
	return res.String(), nil
}
