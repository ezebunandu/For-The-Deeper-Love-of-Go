package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"sync"

)

type Book struct {
    ID string
    Title string
    Author string
    Copies int
}

type Catalog struct {
    mu *sync.RWMutex
    data map[string]Book
}

func NewCatalog() *Catalog {
    return &Catalog{
        mu: &sync.RWMutex{},
        data: map[string]Book{},
    }
}

func (c *Catalog) GetAllBooks() []Book{
    c.mu.RLock()
    defer c.mu.RUnlock()
    return slices.Collect(maps.Values(c.data))
}

func (c *Catalog) GetBook(ID string) (Book, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    book, ok := c.data[ID]
    return  book, ok
}

func (b Book) String() string {
    return  fmt.Sprintf("%v by %v (copies: %d)", b.Title, b.Author, b.Copies)
}

func (c *Catalog) AddBook (b Book) error {
    c.mu.Lock()
    defer c.mu.Unlock()
    _, ok := c.data[b.ID]
    if ok {
        return  fmt.Errorf("book already present in catalog")
    }
    c.data[b.ID] = b
    return  nil
}

func GetCatalog() Catalog {
    return  Catalog{
        data: map[string]Book{
            "abc": {
                ID: "abc",
                Title: "Purple Hibiscus",
                Author: "Chimamanda Ngozi Adichie",
                Copies: 23,
            },
            "xyz": {
                ID: "xyz",
                Title: "The Thing Around Your Neck",
                Author: "Chimamanda Ngozi Adichie",
                Copies: 21,
            },
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

func OpenCatalog(path string) (*Catalog, error) {
    file, err := os.Open(path)
    if err != nil {
        return  nil, err
    }
    defer file.Close()
    catalog := NewCatalog()
    err = json.NewDecoder(file).Decode(&catalog.data)
    if err != nil {
        return nil, err
    }
    return  catalog, nil
}

func (c *Catalog) Sync(path string) error {
    c.mu.RLock()
    defer c.mu.RUnlock()
    file, err := os.Create(path)
    if err != nil {
        return  err
    }
    defer file.Close()
    err = json.NewEncoder(file).Encode(c.data)
    if err != nil {
        return  err
    }
    return  nil
}

func (c *Catalog) SetCopies(ID string, copies int) error {
    c.mu.Lock()
    defer c.mu.Unlock()
    book, ok := c.data[ID]
    if !ok {
        return fmt.Errorf("ID %q not in found", ID)
    }
    err := book.SetCopies(copies)
    if err != nil {
        return  err
    }
    c.data[ID] = book
    return  nil
}

func (catalog *Catalog) GetCopies(ID string) (int, error){
    catalog.mu.RLock()
    defer catalog.mu.RUnlock()
    book, ok := catalog.data[ID]
    if !ok {
        return  0, fmt.Errorf("ID %q not found", ID)
    }
    return  book.Copies, nil
}

