package library

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	api "scripts/api/library"
	utils "scripts/pkg/utils"
)

func FetchDataFromGsheets(destFile string) {
	sheets := []string{"read", "unread", "currently_reading", "queued"}

	gsheet_id, ok := os.LookupEnv("GSHEET_ID")
	if ok != true {
		fmt.Errorf("%s", "GSHEET_ID is not set.")
		os.Exit(1)
	}

	baseUrl := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=", gsheet_id)

	books := make([]*api.Books, 0)
	for _, sheet := range sheets {
		sheetUrl := fmt.Sprintf(baseUrl + "%s", sheet)
		resp, err := http.Get(sheetUrl)
		if err != nil {
			fmt.Errorf("Could not download %s", sheet)
			continue
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Errorf("%s", err)
		}

		r, err := csv.NewReader(bytes.NewReader(responseData)).ReadAll()
		if err != nil {
			continue
		}

		for idx, record := range r {
			if idx == 0 {
				continue
			}

			book := new(api.Books)

			book.BookID = record[0]
			book.Title = record[1]
			book.Author = record[2]

			if record[3] != "" && record[3] != "0" {
				book.ISBN = record[3]
			}

			if record[4] != "" && record[4] != "0" {
				book.ISBN13 = record[4]
			}

			if record[5] == "" || record[5] == "0" {
				fmt.Errorf("Rating does not seem to be valid for %s\n", book.Title)
			} else if myRating, err := utils.ConvertStringToNumber(record[5]); err != nil {
				fmt.Errorf("Rating does not seem to be valid for %s\n", book.Title)
			} else {
				book.MyRating = myRating
			}

			if record[6] == "" || record[6] == "0" {
				fmt.Errorf("NumberOfPages does not seem to be valid for %s\n", book.Title)
			} else if numberOfPages, err := utils.ConvertStringToNumber(record[6]); err != nil {
				fmt.Errorf("NumberOfPages does not seem to be valid for %s\n", book.Title)
			} else {
				book.NumberOfPages = numberOfPages
			}

			if record[7] == "" || record[7] == "0" {
				fmt.Errorf("YearPublished does not seem to be valid for %s\n", book.Title)
			} else if yearPublished, err := utils.ConvertStringToNumber(record[7]); err != nil {
				fmt.Errorf("YearPublished does not seem to be valid for %s\n", book.Title)
			} else {
				book.YearPublished = yearPublished
			}

			if record[8] == "" {
				fmt.Errorf("DateRead does not seem to be valid for %s\n", book.Title)
			} else if dateRead, err := time.Parse("02/01/2006", strings.TrimSpace(record[8])); err != nil {
				fmt.Errorf("error DateRead does not seem to be valid for %s\n", book.Title)
			} else {
				book.DateRead = &dateRead
			}

			if record[9] == "" {
				fmt.Errorf("DateAdded does not seem to be valid for %s\n", book.Title)
			} else if dateAdded, err := time.Parse("02/01/2006", strings.TrimSpace(record[9])); err != nil {
				fmt.Errorf("DateAdded does not seem to be valid for %s\n", book.Title)
			} else {
				book.DateAdded = &dateAdded
			}

			if record[9] == "" {
				fmt.Errorf("DateRead does not seem to be valid for %s\n", book.Title)
			} else if book.DateRead != nil && book.DateRead.Year() == time.Now().Year() && (time.Now().Month() - book.DateRead.Month()) <= 1 {
				book.Bookshelves = "recently_finished"
			} else {
				book.Bookshelves = record[10]
			}

			if len(record) == 12  {
				book.BlogLink = record[11]
			}

			books = append(books, book)
		}
	}

	file, _ := json.MarshalIndent(books, "", " ")

	err := ioutil.WriteFile(destFile, file, 0644)
	if err != nil {
		fmt.Errorf("%s", err)
	}
}