package library

import (
	"time"
)

type Books struct {
	BookID        string    `json:"book_id,omitempty"`
	Title         string `json:"title,omitempty"`
	Author        string `json:"author,omitempty"`
	ISBN          string    `json:"isbn,omitempty"`
	ISBN13        string  `json:"isbn13,omitempty"`
	MyRating      int    `json:"my_rating,omitempty"`
	NumberOfPages int    `json:"num_of_pages,omitempty"`
	YearPublished int    `json:"year_published,omitempty"`
	DateRead      *time.Time `json:"date_read,omitempty"`
	DateAdded     *time.Time  `json:"date_added,omitempty"`
	Bookshelves   string `json:"bookshelves,omitempty"`
	BlogLink   string `json:"blog_link,omitempty"`
}
