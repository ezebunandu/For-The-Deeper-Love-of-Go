package main

import (
    "books"
    "fmt"
)

func main() {
    book, ok := books.GetBook("1")
    if !ok {
        fmt.Println("Sorry, couldn't find that book in the catalog")
        return
    }
    fmt.Println(books.BookToString(book))
}