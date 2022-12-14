package main

import (
	"fmt"

	"github.com/SHA65536/WikipediaGoApi/client"
	"github.com/SHA65536/WikipediaGoApi/opensearch"
)

func main() {
	cl := client.MakeClient()
	res1, err := cl.GetOpenSearch(opensearch.OpenSearchArgs{
		Query: "Te",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n\n", res1)

	res2, err := cl.GetQuerySearch([]string{"Albert Einstein", "Reptile"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n\n", res2)

	res3, err := cl.GetQueryLinksWithContinue("🙏🏿")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(res3), res3)

	var count int
	for gen := cl.GetAllArticles(); count < 10000; count++ {
		if val, cont := gen.Next(); cont {
			fmt.Printf("%04d: %v\n", count, val)
		} else {
			break
		}
	}
}
