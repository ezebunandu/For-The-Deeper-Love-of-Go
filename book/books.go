package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

var catalog = map[string]Book{
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

func GetAllBooks(catalog map[string]Book) []Book {
	return slices.Collect(maps.Values(catalog))
}

func GetBook(catalog map[string]Book, id string) (Book, bool) {
	book, ok := catalog[id]
	return book, ok
}

func AddBook(catalog map[string]Book, b Book) {
	catalog[b.ID] = b
}

func GetCatalog() map[string]Book{
	return map[string]Book{
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
}