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

type Catalog struct {
	mu   *sync.RWMutex
	data map[string]Book
	Path string
}

func NewCatalog() *Catalog {
	return &Catalog{
		mu:   &sync.RWMutex{},
		data: map[string]Book{},
	}
}

func (catalog *Catalog) GetAllBooks() []Book {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	return slices.Collect(maps.Values(catalog.data))
}

func (catalog *Catalog) GetBook(id string) (Book, bool) {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	book, ok := catalog.data[id]
	return book, ok
}

func (catalog *Catalog) AddBook(b Book) error {
	_, ok := catalog.GetBook(b.ID)
	if ok {
		return fmt.Errorf("ID %q already in catalog", b.ID)
	}
	catalog.mu.Lock()
	defer catalog.mu.Unlock()
	catalog.data[b.ID] = b
	return nil
}

func OpenCatalog(path string) (*Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := NewCatalog()
	err = json.NewDecoder(file).Decode(&catalog.data)
	if err != nil {
		return nil, err
	}
	catalog.Path = path
	return catalog, nil
}

func (catalog *Catalog) Sync() error {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	file, err := os.Create(catalog.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(catalog.data)
	if err != nil {
		return err
	}
	return nil
}

func (catalog *Catalog) SetCopies(ID string, copies int) error {
	book, ok := catalog.GetBook(ID)
	if !ok {
		return fmt.Errorf("not found: %q", ID)
	}
	catalog.mu.Lock()
	defer catalog.mu.Unlock()
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	catalog.data[ID] = book
	return nil
}

func (catalog *Catalog) GetCopies(ID string) (int, error) {
	book, ok := catalog.GetBook(ID)
	if !ok {
		return 0, fmt.Errorf("ID %q not found", ID)
	}
	return book.Copies, nil
}
