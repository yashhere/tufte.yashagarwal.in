package library

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	api "scripts/api/library"
	utils "scripts/pkg/utils"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func FetchDataFromGsheets(destFile string) error {
	sheets := []string{"read", "unread", "currently_reading", "queued"}

	gsheet_id, ok := os.LookupEnv("GSHEET_ID")
	if ok != true {
		return errors.New("GSHEET_ID is not set.")
	}

	baseUrl := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=", gsheet_id)

	books := make([]*api.Books, 0)
	for _, sheet := range sheets {
		sheetUrl := fmt.Sprintf(baseUrl+"%s", sheet)
		resp, err := http.Get(sheetUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		r, err := csv.NewReader(bytes.NewReader(responseData)).ReadAll()
		if err != nil {
			return err
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
				log.Debugf("Rating does not seem to be valid for %s\n", book.Title)
			} else if myRating, err := utils.ConvertStringToNumber(record[5]); err != nil {
				log.Debugf("Rating does not seem to be valid for %s\n", book.Title)
			} else {
				book.MyRating = myRating
			}

			if record[6] == "" || record[6] == "0" {
				log.Debugf("NumberOfPages does not seem to be valid for %s\n", book.Title)
			} else if numberOfPages, err := utils.ConvertStringToNumber(record[6]); err != nil {
				log.Debugf("NumberOfPages does not seem to be valid for %s\n", book.Title)
			} else {
				book.NumberOfPages = numberOfPages
			}

			if record[7] == "" || record[7] == "0" {
				log.Debugf("YearPublished does not seem to be valid for %s\n", book.Title)
			} else if yearPublished, err := utils.ConvertStringToNumber(record[7]); err != nil {
				log.Debugf("YearPublished does not seem to be valid for %s\n", book.Title)
			} else {
				book.YearPublished = yearPublished
			}

			if record[8] == "" {
				log.Debugf("DateRead does not seem to be valid for %s\n", book.Title)
			} else if dateRead, err := time.Parse("02/01/2006", strings.TrimSpace(record[8])); err != nil {
				log.Debugf("error DateRead does not seem to be valid for %s\n", book.Title)
			} else {
				book.DateRead = &dateRead
			}

			if record[9] == "" {
				log.Debugf("DateAdded does not seem to be valid for %s\n", book.Title)
			} else if dateAdded, err := time.Parse("02/01/2006", strings.TrimSpace(record[9])); err != nil {
				log.Debugf("DateAdded does not seem to be valid for %s\n", book.Title)
			} else {
				book.DateAdded = &dateAdded
			}

			if record[9] == "" {
				log.Debugf("DateRead does not seem to be valid for %s\n", book.Title)
			} else if book.DateRead != nil && book.DateRead.Year() == time.Now().Year() && (int(time.Now().Sub(*book.DateRead).Hours()/24) <= 30) {
				book.Bookshelves = "recently_finished"
			} else {
				book.Bookshelves = record[10]
			}

			if len(record) == 12 {
				book.BlogLink = record[11]
			}

			books = append(books, book)
		}
	}

	log.Infof("Writing [%d] books to %s.", len(books), path.Base(destFile))
	file, _ := json.MarshalIndent(books, "", " ")

	err := ioutil.WriteFile(destFile, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
