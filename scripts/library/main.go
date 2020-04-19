package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Library struct {
	BookID                  int    `json:"book_id"`
	Title                   string `json:"title"`
	Author                  string `json:"author"`
	ISBN                    string `json:"isbn"`
	ISBN13                  int64  `json:"isbn13"`
	MyRating                int    `json:"my_rating"`
	NumberOfPages           int    `json:"num_of_pages"`
	YearPublished           int    `json:"year_published"`
	DateRead                string `json:"date_read,omitempty"`
	DateAdded               string `json:"date_added,omitempty"`
	Bookshelves             string `json:"bookshelves"`
}

type GoodReadsData struct {
	BookID                  int    `json:"Book Id"`
	Title                   string `json:"Title"`
	Author                  string `json:"Author"`
	AuthorLF                string `json:"Author l-f"`
	AdditionalAuthors       string `json:"Additional Authors"`
	ISBN                    string `json:"ISBN"`
	ISBN13                  int64  `json:"ISBN13"`
	MyRating                int    `json:"My Rating"`
	NumberOfPages           int    `json:"Number of Pages"`
	YearPublished           int    `json:"Year Published"`
	OriginalPublicationYear int    `json:"Original Publication Year"`
	DateRead                string `json:"Date Read"`
	DateAdded               string `json:"Date Added"`
	Bookshelves             string `json:"Bookshelves"`
	ExclusiveShelf          string `json:"Exclusive Shelf"`
}

// https://stackoverflow.com/a/44359967/5042046
func PrettyJson(data interface{}) (string, error) {
	const (
		empty = ""
		tab   = "\t"
	)

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func formatDateTime(s string) string {
	if s == "" {
		return ""
	}
	t, err := time.Parse("2006/01/02", strings.TrimSpace(s))
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		return t.Format("02-01-2006")
	}
}

func main() {
	file, err := ioutil.ReadFile("scripts/library/goodreads_data.json")
	if err != nil {
		fmt.Println(err)
	}
	data := []GoodReadsData{}

	_ = json.Unmarshal([]byte(file), &data)

	library := make(map[string][]*Library)
	for _, d := range data {
		book := new(Library)
		book.BookID = d.BookID
		book.Title = d.Title
		if d.AdditionalAuthors == "" {
			book.Author = d.Author
		} else {
			book.Author = d.Author + " et al."
		}

		book.ISBN = d.ISBN
		book.ISBN13 = d.ISBN13
		book.MyRating = d.MyRating
		book.NumberOfPages = d.NumberOfPages
		book.YearPublished = d.YearPublished

		book.DateRead = formatDateTime(d.DateRead)
		book.DateAdded = formatDateTime(d.DateAdded)

		readState := strings.Replace(d.ExclusiveShelf, "-", "_", -1)
		book.Bookshelves = readState

		library[readState] = append(library[readState], book)
	}

	d, _ := json.MarshalIndent(library, "", " ")

	_ = ioutil.WriteFile("data/library.json", d, 0644)
	if err != nil {
		fmt.Println(err)
	}
}