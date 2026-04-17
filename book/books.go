package books

import "fmt"

type Book struct {
	ID string
	Title  string
	Author string
	Copies int
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

var catalog = []Book{
	{
		ID: "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	{
		ID: "xyz",
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func GetAllBooks() []Book {
	return catalog
}
	
func GetBook(id string) (Book, bool) {
	for _, b := range catalog {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}