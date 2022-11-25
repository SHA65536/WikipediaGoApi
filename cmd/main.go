package main

import (
	"fmt"

	wiki "github.com/SHA65536/WikipediaGoApi/client"
	"github.com/SHA65536/WikipediaGoApi/opensearch"
)

func main() {
	client := wiki.MakeClient()
	res, err := client.GetOpenSearch(opensearch.OpenSearchArgs{
		Query: "Te",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}
