package books

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
    addr string
}

func NewClient(addr string) *Client {
    return  &Client{
        addr: addr,
    }
}

func (client *Client) GetBook(ID string) (Book, error) {
    path := fmt.Sprintf("http://%s/v1/find/%s", client.addr, ID)
    resp, err := http.Get(path)
    if err != nil {
        return Book{}, err
    }
    defer resp.Body.Close()
    if resp.StatusCode == http.StatusNotFound {
        return Book{}, fmt.Errorf("%q not found", ID)
    }
    if resp.StatusCode != http.StatusOK {
        return Book{}, fmt.Errorf("unexpected status %q", resp.Status)
    }
    book := Book{}
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return  Book{}, err
    }
    err = json.Unmarshal(data, &book)
    if err != nil {
        return Book{}, fmt.Errorf("%v in %q", err, data)
    }
    return  book, nil
}

func(client *Client) GetAllBooks() ([]Book, error){
    path := fmt.Sprintf("http://%s/v1/list", client.addr)
    resp, err := http.Get(path)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status %q", resp.Status)
    }
    books := []Book{}
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return  nil, err
    }
    err = json.Unmarshal(data, &books)
    if err != nil {
        return nil, fmt.Errorf("%v in %q", err, data)
    }
    return  books, nil
}