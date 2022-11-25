package wikipediagoapi

import (
	"net/http"

	"github.com/SHA65536/WikipediaGoApi/region"
)

type Client struct {
	Client        *http.Client
	DefaultRegion region.Region
}
