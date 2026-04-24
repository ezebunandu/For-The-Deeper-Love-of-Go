package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

func (book Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func (book *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	return nil
}

type Catalog map[string]Book

var catalog = Catalog{
	"abc": {

		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	"xyz": {
		ID:     "xyz",
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func (c Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(c))
}

func (c Catalog) GetBook(id string) (Book, bool) {
	book, ok := catalog[id]
	return book, ok
}

func (c Catalog) AddBook(b Book) {
	catalog[b.ID] = b
}

func OpenCatalog(path string) (Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	c := Catalog{}
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
