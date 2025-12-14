package books_test

import (
	"books"
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
