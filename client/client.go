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

func (c *Client) GetQueryLinksWithContinue(title string) ([]query.QueryLink, error) {
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

func (c *Client) GetArticles(cont string) (*query.QueryAllResult, error) {
	var res = &query.QueryAllResult{}
	url, err := query.AllQueryToUrl(c.DefaultRegion, cont)
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

func (c *Client) GetAllArticles() *GetAllArticlesGen {
	return &GetAllArticlesGen{
		Client: c,
		Done:   false,
		Idx:    0,
		Cur:    make([]query.QueryAllPage, 0),
		Cont:   "",
	}
}

type GetAllArticlesGen struct {
	Client *Client
	Done   bool
	Idx    int
	Cur    []query.QueryAllPage
	Cont   string
}

func (g *GetAllArticlesGen) Next() (*query.QueryAllPage, bool) {
	if g.Done {
		return nil, false
	}
	for g.Idx == len(g.Cur) {
		res, err := g.Client.GetArticles(g.Cont)
		if err != nil {
			return nil, false
		}

		g.Cur = res.Query.Pages
		g.Cont = res.Cont.Next
		g.Idx = 0

		if len(res.Query.Pages) == 0 {
			if res.Cont.Next == "" {
				return nil, false
			}
			continue
		}

		if res.Cont.Next == "" {
			g.Done = true
		}
	}
	g.Idx++
	return &g.Cur[g.Idx-1], true
}

/*
func (c *Client) GetQueryLinksWithContinue(title string) ([]query.QueryLink, error) {
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
*/
