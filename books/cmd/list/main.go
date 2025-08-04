package main

import (
	"books"
	"fmt"
	"os"
)



func main(){
    client := books.NewClient("localhost:3000")
    books, err := client.GetAllBooks()
    if err != nil {
        fmt.Fprintf(os.Stderr, "err: %#v", err)
        return
    }
    for _, book := range books {
        fmt.Println(book)
    }
}
