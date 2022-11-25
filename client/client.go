package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	search "github.com/SHA65536/WikipediaGoApi/opensearch"
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
