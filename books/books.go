package books

import "fmt"

type Book struct {
    Title string
    Author string
    Copies int
}

var catalog = []Book{
    {
        Title: "Purple Hibiscus",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 23,
    },
    {
        Title: "The Thing Around Your Neck",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 21,
    },
}

func GetAllBooks() []Book{
    return  catalog
}

func BookToString(book Book) string {
    return  fmt.Sprintf("%v by %v (copies: %d)", book.Title, book.Author, book.Copies)
}