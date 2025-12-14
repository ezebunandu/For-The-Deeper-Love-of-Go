package books_test

import (
	"books"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBookToString__FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := books.BookToString(input)
	if want != got {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestGetAllBooks__ReturnsAllBooks(t *testing.T){
	t.Parallel()
	want := []books.Book{
		{
			Title: "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		{
			Title: "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
	got := books.GetAllBooks()
	if !slices.Equal(want, got){
		t.Fatalf("want %#v, got %#v", want, got)
	}
}
