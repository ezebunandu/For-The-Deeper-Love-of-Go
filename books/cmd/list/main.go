package main

import (
    "fmt"
    "books"
)



func main(){
    catalog, err := books.OpenCatalog("testdata/catalog")
    if err != nil {
        fmt.Printf("opening catalog: %v\n", err)
        return
    }
    for _, book := range catalog.GetAllBooks(){
        fmt.Println(book)
    }
}
