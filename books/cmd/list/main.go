package main

import (
    "fmt"
    "books"
)



func main(){
    catalog := books.GetCatalog()
    for _, book := range books.GetAllBooks(catalog){
        fmt.Println(book)
    }
}
