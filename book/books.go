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

func (catalog Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func (catalog Catalog) GetBook(id string) (Book, bool) {
	book, ok := catalog[id]
	return book, ok
}

func (catalog Catalog) AddBook(b Book) error {
	_, ok := catalog.GetBook(b.ID)
	if ok {
		return fmt.Errorf("ID %q already in catalog", b.ID)
	}
	catalog[b.ID] = b
	return nil
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

func (catalog Catalog) Sync(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(catalog)
	if err != nil {
		return err
	}
	return nil
}

func (catalog Catalog) SetCopies(ID string, copies int) error {
	book, ok := catalog.GetBook(ID)
	if !ok {
		return fmt.Errorf("not found: %q", ID)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	catalog[ID] = book
	return nil
}
