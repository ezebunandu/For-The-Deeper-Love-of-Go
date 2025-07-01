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

type Catalog map[string]Book

func (c Catalog) GetAllBooks() []Book{
    return slices.Collect(maps.Values(c))
}

func (c Catalog) GetBook(ID string) (Book, bool) {
    book, ok := c[ID]
    return  book, ok
}

func (b Book) String() string {
    return  fmt.Sprintf("%v by %v (copies: %d)", b.Title, b.Author, b.Copies)
}

func (c Catalog) AddBook(b Book){
    c[b.ID] = b
}

func GetCatalog() Catalog {
    return Catalog{
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