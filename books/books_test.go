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
		Copies: 1,
	}
	want := "We Must Set Forth At Dawn by Wole Soyinka (copies: 1)"
	got := input.String()

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooks__ReturnsAllBooksInCatalog(t *testing.T) {
	t.Parallel()
    catalog := getTestCatalog()
    got := catalog.GetAllBooks()
    assertTestBooks(t, got) 
}

func TestGetBook__FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	c := getTestCatalog()
	want := books.Book{
		ID:     "abc",
		Title:  "Purple Hibiscus",
		Author: "Chimamanda Ngozi Adichie",
		Copies: 1,
	}
	got, ok := c.GetBook("abc")
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
	_, ok := c.GetBook("nonexistent ID")
	if ok {
		t.Fatalf("want false for nonexistent ID, got true")
	}
}

func TestAddBook__AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	c := getTestCatalog()
	_, ok := c.GetBook("123")
	if ok {
		t.Fatal("book already present")
	}
	c.AddBook(books.Book{
		ID:     "123",
		Title:  "One Day I will Write About This Place",
		Author: "Binyavinga Wainana",
		Copies: 1,
	})
	_, ok = c.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}

func getTestCatalog() books.Catalog {
	return books.Catalog{
		"abc": {
			ID:     "abc",
			Title:  "Purple Hibiscus",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 1,
		},
		"xyz": {
			ID:     "xyz",
			Title:  "The Thing Around Your Neck",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 1,
		},
	}
}

func TestSetCopies__SetsNumberOfCopiesToGivenvalue(t *testing.T){
    t.Parallel()
    book :=books.Book{
        Copies: 1,
    }
    err := book.SetCopies(2)
    if err != nil {
        t.Fatal(err)
    }
    if book.Copies != 2 {
        t.Errorf("want 12 copies, got %d", book.Copies)
    }
}

func TestSetCopies__ReturnsErrorIfCopiesNegative(t *testing.T){
    t.Parallel()
    book := books.Book{}
    err := book.SetCopies(-1)
    if err == nil {
        t.Error("want error for negative copies, got nil")
    }
}

func TestOpenCatalog__ReadsSameDataWrittenBySync(t *testing.T){
    t.Parallel()
    catalog := getTestCatalog()
    path := t.TempDir() + "/catalog"
    err := catalog.Sync(path)
    if err != nil {
        t.Fatal(err)
    }
    newCatalog, err := books.OpenCatalog(path)
    if err != nil {
        t.Fatal(err)
    }
    got := newCatalog.GetAllBooks()
    assertTestBooks(t, got)
}

func assertTestBooks(t *testing.T, got []books.Book) {
    t.Helper()
    want := []books.Book{
		{
			ID:     "abc",
			Title:  "Purple Hibiscus",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 1,
		},
		{
			ID:     "xyz",
			Title:  "The Thing Around Your Neck",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 1,
		},
	}
    slices.SortFunc(got, func(a, b books.Book) int {
        return cmp.Compare(a.ID, b.ID)
    })
    if ! slices.Equal(want, got) {
        t.Fatalf("want %#v, got %#v", want, got)
    }
}