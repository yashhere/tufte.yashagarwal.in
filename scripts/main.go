package main

import (
	"scripts/pkg/library"

	"fmt"

	"scripts/pkg/links"
)

func main() {
	var err error
	destLibraryFile := "../static/data/library.json"
	err = library.FetchDataFromGsheets(destLibraryFile)
	if err != nil {
		fmt.Println(err)
	}

	destBookmarksFile := "../static/data/bookmarks.json"
	err = links.CreateBookMarksFile(destBookmarksFile)
	if err != nil {
		fmt.Println(err)
	}
}
