# WikipediaGoApi
Wikipedia Go Api is a simple wrapper for some wikipedia api endpoints


## OpenSearch
Search for articles given a part of the title.
Example:
```go
cl := client.MakeClient()
res, err := cl.GetOpenSearch(opensearch.OpenSearchArgs{
    Query: "Te",
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", res)
```

## Query
Search for article info using titles
```go
cl := client.MakeClient()
res, err := cl.GetQuerySearch([]string{"Albert Einstein", "Reptile"})
if err != nil {
    panic(err)
}
fmt.Printf("%+v", res)
```

Search for links within an article
```go
cl := client.MakeClient()
res, err := cl.GetQueryLinksWithContinue("Turtle")
if err != nil {
    panic(err)
}
fmt.Printf("%+v", res)
```

List all pages
```go
cl := client.MakeClient()
var count int
for gen := cl.GetAllArticles(); count < 10000; count++ {
    if val, cont := gen.Next(); cont {
        fmt.Printf("%04d: %v\n", count, val)
    } else {
        break
    }
}
```