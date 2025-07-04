package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
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

func (c Catalog) AddBook (b Book) error {
    _, ok := c.GetBook(b.ID)
    if ok {
        return  fmt.Errorf("book already present in catalog")
    }
    c[b.ID] = b
    return  nil
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

func (b *Book) SetCopies(copies int) error {
    if copies < 0 {
        return fmt.Errorf("invalid copies %d", copies)
    }
    b.Copies = copies
    return nil
}

func OpenCatalog(path string) (Catalog, error) {
    file, err := os.Open(path)
    if err != nil {
        return  nil, err
    }
    defer file.Close()
    catalog := Catalog{}
    err = json.NewDecoder(file).Decode(&catalog)
    if err != nil {
        return nil, err
    }
    return  catalog, nil
}

func (c Catalog) Sync(path string) error {
    file, err := os.Create(path)
    if err != nil {
        return  err
    }
    err = json.NewEncoder(file).Encode(c)
    if err != nil {
        return  err
    }
    return  nil
}

func (c Catalog) SetCopies(ID string, copies int) error {
    book, ok := c.GetBook(ID)
    if !ok {
        return fmt.Errorf("ID %q not in found", ID)
    }
    err := book.SetCopies(copies)
    if err != nil {
        return  err
    }
    c[ID] = book
    return  nil
}