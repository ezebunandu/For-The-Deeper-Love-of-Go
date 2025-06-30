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
    got := books.GetAllBooks()
    if !slices.Equal(want, got){
        t.Fatalf("want %#v, got %#v", want, got)
    }
}

func TestGetBook__FindsBookInCatalogByID(t *testing.T){
    t.Parallel()
    want := books.Book{
            ID: "1",
            Title: "Purple Hibiscus",
            Author: "Chimamanda Ngozi Adichie",
            Copies: 23,
    }
    got, ok := books.GetBook("1")
    if !ok {
        t.Fatal("book not found")
    }
    if want != got {
        t.Fatalf("want %#v, got %#v", want, got)
    }

}

func TestGetBook__ReturnsFalseWhenBookNotInCatalog(t *testing.T){
    t.Parallel()
    _, ok := books.GetBook("nonexistent ID")
    if ok {
        t.Fatalf("want false for nonexistent ID, got true")
    }
}