package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"

	gocmp "github.com/google/go-cmp/cmp"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := input.String()
	if want != got {
		t.Fatal(gocmp.Diff(want, got))
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	bookList := catalog.GetAllBooks()
	assertTestBooks(t, bookList)
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := books.Book{
		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}
	got, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not found")
	}
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("book already present")
	}
	b := books.Book{
		ID:     "123",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	}
	catalog.AddBook(b)
	_, ok = catalog.GetBook(b.ID)
	if !ok {
		t.Fatal("added book not found")
	}
}

func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T) {
	t.Parallel()
	book := books.Book{
		Copies: 5,
	}
	err := book.SetCopies(12)
	if err != nil {
		t.Fatal(err)
	}
	if book.Copies != 12 {
		t.Errorf("want 12 copies, got %d", book.Copies)
	}
}

func TestSetCopies_ReturnsErrorGivenNegativeCopies(t *testing.T) {
	t.Parallel()
	book := books.Book{}
	err := book.SetCopies(-1)
	if err == nil {
		t.Fatal("want error given negative copies but got none")
	}
}

func TestOpenCatalog_ReadsSameDataWrittenBySync(t *testing.T) {
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
	bookList := newCatalog.GetAllBooks()
	assertTestBooks(t, bookList)
}

// func TestSync_WritesCatalogDataToFile(t *testing.T) {
// 	t.Parallel()
// 	catalog := getTestCatalog()
// 	err := catalog.Sync("testdata/catalog.new")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	newCatalog, err := books.OpenCatalog("testdata/catalog.new")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	bookList := newCatalog.GetAllBooks()
// 	assertTestBooks(t, bookList)
// }

func getTestCatalog() books.Catalog {
	return books.Catalog{
		"abc": {

			ID:     "abc",
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		"xyz": {
			ID:     "xyz",
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
}

func assertTestBooks(t *testing.T, got []books.Book){
	t.Helper()
	want := []books.Book{
		{
			ID:     "abc",
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		{
			ID:     "xyz",
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}
