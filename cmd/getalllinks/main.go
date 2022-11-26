package main

import (
	"encoding/csv"
	"os"
	"sync"

	"github.com/SHA65536/WikipediaGoApi/client"
	"github.com/schollz/progressbar/v3"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

var workers = 10

func main() {
	cl := client.MakeClient()

	f, err := os.Open("articles.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	articles, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	myCsv, err := ccsv.NewCsvWriter("links.csv")
	if err != nil {
		panic(err)
	}
	defer myCsv.Close()

	bar := progressbar.Default(int64(len(articles)), "Articles Parsed")

	var wg sync.WaitGroup
	chunkSize := (len(articles) + workers - 1) / workers
	for i := 0; i < len(articles); i += chunkSize {
		i := i
		wg.Add(1)
		end := i + chunkSize
		if end > len(articles) {
			end = len(articles)
		}
		go func() {
			defer wg.Done()
			for j := range articles[i:end] {
				links, err := cl.GetQueryLinksWithContinue(articles[i:end][j][1])
				if err != nil {
					panic(err)
				}
				for _, link := range links {
					myCsv.Write([]string{articles[i:end][j][1], link.Title})
				}
				bar.Add(1)
			}
		}()
	}
	wg.Wait()
}
