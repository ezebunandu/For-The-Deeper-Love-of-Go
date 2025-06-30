package books

import "fmt"

type Book struct {
    ID string
    Title string
    Author string
    Copies int
}

var catalog = []Book{
    {
        ID: "1",
        Title: "Purple Hibiscus",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 23,
    },
    {
        ID: "2",
        Title: "The Thing Around Your Neck",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 21,
    },
}

func GetAllBooks() []Book{
    return  catalog
}

func GetBook(id string) (Book, bool) {
    for _, book := range catalog {
        if book.ID == id {
            return  book, true
        }
    }
    return Book{}, false
}

func BookToString(book Book) string {
    return  fmt.Sprintf("%v by %v (copies: %d)", book.Title, book.Author, book.Copies)
}