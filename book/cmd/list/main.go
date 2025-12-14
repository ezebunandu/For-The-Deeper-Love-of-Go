package main

import (
	"fmt"

	"books"
)

func main() {
	fmt.Println("Books in stock:")
	book := books.Book{
		Title:  "Engineering in Plain Sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}

	fmt.Println(books.BookToString(book))
}
