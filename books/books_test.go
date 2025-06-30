package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestBookToString__FormatsBookInfoAsString(t *testing.T) {
    t.Parallel()
    input := books.Book{
        Title: "We Must Set Forth At Dawn",
        Author: "Wole Soyinka",
        Copies: 2,
    }
    want := "We Must Set Forth At Dawn by Wole Soyinka (copies: 2)"
    got := books.BookToString(input)

    if want != got {
        t.Fatalf("want %q, got %q", want, got)
    }
}

func TestGetAllBooks__ReturnsAllBooksInCatalog(t *testing.T){
    want := []books.Book{
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
    got := books.GetAllBooks()
    if !slices.Equal(want, got){
        t.Fatalf("want %#v, got %#v", want, got)
    }
}