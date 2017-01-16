package main

import (
    "fmt"
    "github.com/lasalefamine/goquote"
    "log"
)

func main() {
    items, err := goquote.GetQuote("http://quotesondesign.com/wp-json/posts?filter[orderby]=rand&filter[posts_per_page]=1")
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item)
    }
}
