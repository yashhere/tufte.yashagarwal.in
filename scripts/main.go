package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"scripts/pkg/library"
	"scripts/pkg/links"
)

func setupLogger() {
	mw := io.Writer(os.Stdout)
	log.SetOutput(mw)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	log.SetLevel(log.InfoLevel)
}

func main() {
	setupLogger()

	var err error
	destLibraryFile := "../data/library.json"
	err = library.FetchDataFromGsheets(destLibraryFile)
	if err != nil {
		log.Fatal(err)
	}

	destBookmarksFile := "../data/bookmarks.json"
	err = links.CreateBookMarksFile(destBookmarksFile)
	if err != nil {
		log.Fatal(err)
	}
}
