package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
    ID string
    Title string
    Author string
    Copies int
}

var catalog = map[string]Book{
    "abc": 
    {
        ID: "abc",
        Title: "Purple Hibiscus",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 23,
    },
    "xyz":
    {
        ID: "xyz",
        Title: "The Thing Around Your Neck",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 21,
    },
}

func GetAllBooks() []Book{
    return slices.Collect(maps.Values(catalog))
}

func GetBook(ID string) (Book, bool) {
    book, ok := catalog[ID]
    return  book, ok
}

func BookToString(book Book) string {
    return  fmt.Sprintf("%v by %v (copies: %d)", book.Title, book.Author, book.Copies)
}