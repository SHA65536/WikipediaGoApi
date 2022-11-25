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
fmt.Printf("%+v", res2)
```

## WIP
### Parse
https://en.wikipedia.org/w/api.php?action=help&modules=parse
- title/pageid