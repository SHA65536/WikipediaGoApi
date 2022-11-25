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
	res3, err := cl.GetAllQueryLinks("Turtle")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(res3), res3)
}
