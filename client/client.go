package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	search "github.com/SHA65536/WikipediaGoApi/opensearch"
	"github.com/SHA65536/WikipediaGoApi/query"
	"github.com/SHA65536/WikipediaGoApi/region"
)

type Client struct {
	Client        *http.Client
	DefaultRegion region.Region
}

func MakeClient() *Client {
	return &Client{
		Client:        http.DefaultClient,
		DefaultRegion: region.English,
	}
}

func (c *Client) GetOpenSearch(args search.OpenSearchArgs) (*search.OpenSearchResult, error) {
	var res = &search.OpenSearchResult{}
	url, err := args.ToQuery(c.DefaultRegion)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	return res, err
}

func (c *Client) GetQuerySearch(titles []string) (*query.QueryInfoResult, error) {
	var res = &query.QueryInfoResult{}
	url, err := query.InfoQueryToUrl(c.DefaultRegion, titles)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	return res, err
}

func (c *Client) GetQueryLinks(title, cont string) (*query.QueryLinksResult, error) {
	var res = &query.QueryLinksResult{}
	url, err := query.LinksQueryToUrl(c.DefaultRegion, title, cont)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	return res, err
}

func (c *Client) GetAllQueryLinks(title string) ([]query.QueryLink, error) {
	var stop bool
	var cont string
	var result = []query.QueryLink{}
	for !stop {
		stop = true
		var res, err = c.GetQueryLinks(title, cont)
		if err != nil {
			return nil, err
		}
		result = append(result, res.Query.Pages[0].Links...)
		if res.Cont.Next != "" {
			stop = false
			cont = res.Cont.Next
		}
	}
	return result, nil
}
