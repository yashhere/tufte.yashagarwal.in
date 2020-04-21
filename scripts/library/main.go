package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

type Library struct {
	BookID                  int    `json:"book_id"`
	Title                   string `json:"title"`
	Author                  string `json:"author"`
	ISBN                    int 	`json:"isbn"`
	ISBN13                  int64  `json:"isbn13"`
	MyRating                int    `json:"my_rating"`
	NumberOfPages           int    `json:"num_of_pages"`
	YearPublished           int    `json:"year_published"`
	DateRead                string `json:"date_read,omitempty"`
	DateAdded               string `json:"date_added,omitempty"`
	Bookshelves             string `json:"bookshelves"`
}

type GoodReadsData struct {
	BookID                   int     `json:"Book Id"`
	Title                    string  `json:"Title"`
	Author                   string  `json:"Author"`
	AuthorLF                 string  `json:"Author l-f"`
	AdditionalAuthors        string  `json:"Additional Authors"`
	ISBN                     int     `json:"ISBN"`
	ISBN13                   int64   `json:"ISBN13"`
	MyRating                 int     `json:"My Rating"`
	AverageRating            float64 `json:"Average Rating"`
	Publisher                string  `json:"Publisher"`
	Binding                  string  `json:"Binding"`
	NumberOfPages            int     `json:"Number of Pages"`
	YearPublished            int     `json:"Year Published"`
	OriginalPublicationYear  int     `json:"Original Publication Year"`
	DateRead                 string  `json:"Date Read"`
	DateAdded                string  `json:"Date Added"`
	Bookshelves              string  `json:"Bookshelves"`
	BookshelvesWithPositions string  `json:"Bookshelves with positions"`
	ExclusiveShelf           string  `json:"Exclusive Shelf"`
	MyReview                 string  `json:"My Review"`
	Spoiler                  string  `json:"Spoiler"`
	PrivateNotes             string  `json:"Private Notes"`
	ReadCount                int     `json:"Read Count"`
	RecommendedFor           string  `json:"Recommended For"`
	RecommendedBy            string  `json:"Recommended By"`
	OwnedCopies              int     `json:"Owned Copies"`
	OriginalPurchaseDate     string  `json:"Original Purchase Date"`
	OriginalPurchaseLocation string  `json:"Original Purchase Location"`
	Condition                string  `json:"Condition"`
	ConditionDescription     string  `json:"Condition Description"`
	BCID                     string  `json:"BCID"`
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
	file, err := ioutil.ReadFile("scripts/library/goodreads.json")
	if err != nil {
		fmt.Println(err)
	}
	data := []GoodReadsData{}

	_ = json.Unmarshal([]byte(file), &data)

	space := regexp.MustCompile(`\s+`)

	library := make([]*Library, 0)
	for _, d := range data {
		book := new(Library)
		book.BookID = d.BookID
		book.Title = space.ReplaceAllString(d.Title, " ")
		if d.AdditionalAuthors == "" {
			book.Author = space.ReplaceAllString(d.Author, " ")
		} else {
			book.Author = space.ReplaceAllString(d.Author, " ") + " et al."
		}

		book.ISBN = d.ISBN
		book.ISBN13 = d.ISBN13
		book.MyRating = d.MyRating
		book.NumberOfPages = d.NumberOfPages
		book.YearPublished = d.YearPublished

		book.DateRead = formatDateTime(d.DateRead)
		book.DateAdded = formatDateTime(d.DateAdded)

		readState := space.ReplaceAllString(strings.Replace(d.ExclusiveShelf, "-", "_", -1), " ")

		t, err := time.Parse("02-01-2006", book.DateRead)
		if err == nil {
			if (time.Now().Month() - t.Month()) <= 1 && time.Now().Year() == t.Year() {
				readState = "recently_finished"
			}
		}

		book.Bookshelves = readState
		library = append(library, book)
	}

	d, _ := json.MarshalIndent(library, "", " ")

	_ = ioutil.WriteFile("data/library.json", d, 0644)
	if err != nil {
		fmt.Println(err)
	}
}