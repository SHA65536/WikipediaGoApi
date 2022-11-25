package opensearch

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	ns "github.com/SHA65536/WikipediaGoApi/namespace"
	"github.com/SHA65536/WikipediaGoApi/region"
)

type SearchProfile string

const (
	Strict  SearchProfile = "strict"            // Strict profile with few punctuation characters removed but diacritics and stress marks are kept
	Normal  SearchProfile = "normal"            // Few punctuation characters, some diacritics and stopwords removed
	Fuzzy   SearchProfile = "fuzzy"             // Similar to normal with typo correction (two typos supported)
	Classic SearchProfile = "classic"           // Classic prefix, few punctuation characters and some diacritics removed
	Auto    SearchProfile = "engine_autoselect" // Let wikipedia decide
)

type RedirectType string

const (
	Return  RedirectType = "return"  // Shows non-canonical links
	Resolve RedirectType = "resolve" // Resolves redirects to canonical link
)

type OpenSearchArgs struct {
	Query      string         // String to search
	Namespaces []ns.Namespace // Filter result namespaces (Default 0)
	Limit      int            // Result limit (Default 10) (Max 500)
	Profile    SearchProfile  // Type of search
	Redirect   RedirectType   // Type of resolution (Default Resolve)
}

// Formats the search arguments into a url
func (args *OpenSearchArgs) ToQuery(base region.Region) (string, error) {
	var values = url.Values{}
	res, err := url.Parse(string(base))
	if err != nil {
		return "", err
	}

	values.Set("action", "opensearch")
	values.Set("formatversion", "2")
	values.Set("format", "json")

	if args.Query == "" {
		return "", fmt.Errorf("query must not be empty")
	}
	values.Set("search", args.Query)

	if len(args.Namespaces) != 0 {
		var nslist = make([]string, len(args.Namespaces))
		for i := range args.Namespaces {
			nslist[i] = args.Namespaces[i].String()
		}
		values.Set("namespace", strings.Join(nslist, "|"))
	}

	if args.Limit < 0 || args.Limit > 500 {
		return "", fmt.Errorf("limit must be between 0 and 500")
	} else if args.Limit > 0 {
		values.Set("limit", strconv.Itoa(args.Limit))
	}

	if args.Profile != "" {
		values.Set("profile", string(args.Profile))
	}

	if args.Redirect != "" {
		values.Set("redirects", string(args.Redirect))
	} else {
		values.Set("redirects", string(Resolve))
	}

	res.RawQuery = values.Encode()

	return res.String(), nil
}

type OpenSearchResult struct {
	Query  string   // User query
	Titles []string // List of titles
	Links  []string // List of links
}

func (r *OpenSearchResult) UnmarshalJSON(data []byte) error {
	var lists = []any{}
	err := json.Unmarshal(data, &lists)
	if err != nil {
		return err
	}
	if len(lists) != 4 {
		return fmt.Errorf("not enough arguments in return")
	}
	if val, ok := lists[0].(string); ok {
		r.Query = val
	} else {
		return fmt.Errorf("query not string")
	}

	if val, ok := lists[1].([]interface{}); ok {
		r.Titles = make([]string, len(val))
		for i := range val {
			if title, ok := val[i].(string); ok {
				r.Titles[i] = title
			} else {
				return fmt.Errorf("titles not list of strings")
			}
		}
	} else {
		return fmt.Errorf("titles not list of strings")
	}

	if val, ok := lists[3].([]interface{}); ok {
		r.Links = make([]string, len(val))
		for i := range val {
			if link, ok := val[i].(string); ok {
				r.Links[i] = link
			} else {
				return fmt.Errorf("links not list of strings")
			}
		}
	} else {
		return fmt.Errorf("links not list of strings")
	}

	return nil
}
