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


func GetAllBooks(catalog map[string]Book) []Book{
    return slices.Collect(maps.Values(catalog))
}

func GetBook(catalog map[string]Book, ID string) (Book, bool) {
    book, ok := catalog[ID]
    return  book, ok
}

func (b Book) String() string {
    return  fmt.Sprintf("%v by %v (copies: %d)", b.Title, b.Author, b.Copies)
}

func AddBook(catalog map[string]Book, b Book){
    catalog[b.ID] = b
}

func GetCatalog() map[string]Book {
    return map[string]Book{
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
}