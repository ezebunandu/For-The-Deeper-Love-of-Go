package main

import (
	"fmt"

	"books"
)

func main() {
	catalog := books.GetCatalog()
	for _, book := range catalog.GetAllBooks() {
		fmt.Println(book)
	}
}
