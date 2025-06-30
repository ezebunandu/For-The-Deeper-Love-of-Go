package main

import (
    "fmt"
    "books"
)



func main(){
    for _, book := range books.GetAllBooks(){
        fmt.Println(books.BookToString(book))
    }
}
