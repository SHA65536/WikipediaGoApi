package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/SHA65536/WikipediaGoApi/client"
	"github.com/schollz/progressbar/v3"
)

func main() {
	cl := client.MakeClient()

	f, err := os.Create("articles.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	var count int
	bar := progressbar.Default(6579999)
	for gen := cl.GetAllArticles(); true; count++ {
		if val, cont := gen.Next(); cont {
			if err := w.Write([]string{strconv.Itoa(val.Id), val.Title}); err != nil {
				panic(err)
			}
			bar.Add(1)
		} else {
			break
		}
	}
}
