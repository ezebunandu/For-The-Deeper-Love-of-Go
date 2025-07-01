package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func TestBookToString__FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "We Must Set Forth At Dawn",
		Author: "Wole Soyinka",
		Copies: 2,
	}
	want := "We Must Set Forth At Dawn by Wole Soyinka (copies: 2)"
	got := input.String()

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooks__ReturnsAllBooksInCatalog(t *testing.T) {
	c := getTestCatalog()
	want := []books.Book{
		{
			ID:     "abc",
			Title:  "Purple Hibiscus",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 23,
		},
		{
			ID:     "xyz",
			Title:  "The Thing Around Your Neck",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 21,
		},
	}
	got := books.GetAllBooks(c)
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.ID, b.ID)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook__FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	c := getTestCatalog()
	want := books.Book{
		ID:     "abc",
		Title:  "Purple Hibiscus",
		Author: "Chimamanda Ngozi Adichie",
		Copies: 23,
	}
	got, ok := books.GetBook(c, "abc")
	if !ok {
		t.Fatal("book not found")
	}
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}

func TestGetBook__ReturnsFalseWhenBookNotInCatalog(t *testing.T) {
	t.Parallel()
	c := getTestCatalog()
	_, ok := books.GetBook(c, "nonexistent ID")
	if ok {
		t.Fatalf("want false for nonexistent ID, got true")
	}
}

func TestAddBook__AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	c := getTestCatalog()
	_, ok := books.GetBook(c, "123")
	if ok {
		t.Fatal("book already present")
	}
	books.AddBook(c, books.Book{
		ID:     "123",
		Title:  "One Day I will Write About This Place",
		Author: "Binyavinga Wainana",
		Copies: 8,
	})
	_, ok = books.GetBook(c, "123")
	if !ok {
		t.Fatal("added book not found")
	}
}

func getTestCatalog() map[string]books.Book {
	return map[string]books.Book{
		"abc": {
			ID:     "abc",
			Title:  "Purple Hibiscus",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 23,
		},
		"xyz": {
			ID:     "xyz",
			Title:  "The Thing Around Your Neck",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 21,
		},
	}
}
