package books_test

import (
	"books"
	"cmp"
	"encoding/json"
	"io"
	"net"
	"net/http"
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
	err := c.AddBook(books.Book{
		ID:     "123",
		Title:  "One Day I will Write About This Place",
		Author: "Binyavinga Wainana",
		Copies: 1,
	})
    if err != nil {
        t.Fatal(err)
    }
	_, ok = c.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}

func getTestCatalog() *books.Catalog {
    catalog := books.NewCatalog()
    err := catalog.AddBook(books.Book{
        ID:     "abc",
        Title:  "Purple Hibiscus",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 1,
    })
    if err != nil {
        panic(err)
    }
    err = catalog.AddBook(books.Book{
        ID:     "xyz",
        Title:  "The Thing Around Your Neck",
        Author: "Chimamanda Ngozi Adichie",
        Copies: 1,
    })
    if err != nil {
        panic(err)
    }
    return  catalog
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

func TestSetCopies__OnCatalogModifiesSpecifiedBook(t *testing.T){
    t.Parallel()
    catalog := getTestCatalog()
    book, ok := catalog.GetBook("abc")
    if !ok {
        t.Fatal("book not found")
    }
    if book.Copies != 1 {
        t.Fatalf("want 1 copy before change, got %d", book.Copies)
    }
    err := catalog.SetCopies("abc", 2)
    if err != nil {
        t.Fatal(err)
    }
    book, ok = catalog.GetBook("abc")
    if !ok {
        t.Fatal("book not found")
    }
    if book.Copies != 2 {
        t.Fatalf("want 2 copies after change, got %d", book.Copies)
    }
}

func TestAddBook__ReturnsErrorWhenBookAlreadyInCatalog(t *testing.T){
    t.Parallel()
    catalog := getTestCatalog()
    _, ok := catalog.GetBook("abc")
    if !ok{
        t.Fatal("book not present")
    }
    err := catalog.AddBook(books.Book{
			ID:     "abc",
			Title:  "Purple Hibiscus",
			Author: "Chimamanda Ngozi Adichie",
			Copies: 1,},
        )

    if err == nil {
        t.Fatal("want error when book ID already in catalog, got nil")
    }
}

func TestSetCopies__IsRaceFree(t *testing.T){
    t.Parallel()
    catalog := getTestCatalog()
    go func(){
        for range 100 {
            err := catalog.SetCopies("abc", 0)
            if err != nil {
                panic(err)
            }
        }
    }()
    for range 100 {
        _, err := catalog.GetCopies("abc")
        if err != nil {
            t.Fatal(err)
        }
    }
}

func TestNewCatalog__CreatesEmptyCatalog(t *testing.T){
    t.Parallel()
    catalog := books.NewCatalog()
    books := catalog.GetAllBooks()
    if len(books) > 0 {
        t.Errorf("want empty catalog, got %#v", books)
    }
}

func TestServer__ListsAllBooks(t *testing.T){
    t.Parallel()
    addr := randomLocalAddr(t)
    go func(){
        err := books.ListenAndServe(addr, getTestCatalog())
        if err != nil {
            panic(err)
        }
    }()
    resp, err := http.Get("http://" + addr + "/list")
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("unexpected status %d", resp.StatusCode)
    }
    got := []books.Book{}
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatal(err)
    }
    err = json.Unmarshal(data, &got)
    if err != nil {
        t.Fatalf("%v in %q", err, data)
    }
    assertTestBooks(t, got)

}

func randomLocalAddr(t *testing.T) string {
    t.Helper()
    l, err := net.Listen("tcp", ":0")
    if err != nil {
        t.Fatal(err)
    }
    defer l.Close()
    return l.Addr().String()
}