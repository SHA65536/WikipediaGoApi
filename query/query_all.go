package query

import (
	"net/url"

	"github.com/SHA65536/WikipediaGoApi/namespace"
	"github.com/SHA65536/WikipediaGoApi/region"
)

type QueryAllResult struct {
	Cont  QueryAllContinue `json:"continue"`
	Query QueryAllQuery    `json:"query"`
}

type QueryAllQuery struct {
	Pages []QueryAllPage `json:"pages"`
}

type QueryAllPage struct {
	Id    int                 `json:"pageid"`
	Ns    namespace.Namespace `json:"ns"`
	Title string              `json:"title"`
}

type QueryAllContinue struct {
	Next string `json:"gapcontinue"`
}

func AllQueryToUrl(base region.Region, cont string) (string, error) {
	var values = url.Values{}
	res, err := url.Parse(string(base))
	if err != nil {
		return "", err
	}

	values.Set("action", "query")
	values.Set("generator", "allpages")
	values.Set("gapnamespace", "0")
	values.Set("formatversion", "2")
	values.Set("format", "json")
	values.Set("gaplimit", "max")
	values.Set("gapfilterredir", "nonredirects")

	if cont != "" {
		values.Set("gapcontinue", cont)
	}

	res.RawQuery = values.Encode()
	return res.String(), nil
}

//https://en.wikipedia.org/w/api.php?action=query&generator=allpages&gapnamespace=0&gaplimit=max&format=json&formatversion=2&gapfilterredir=nonredirects&gapcontinue=`Usuman_dan_Muhammad_Fodio
