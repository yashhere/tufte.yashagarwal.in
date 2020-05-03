package main

import (
	"fmt"

	"scripts/pkg/library"
	"scripts/pkg/links"
)

func main() {
	var err error
	destLibraryFile := "../data/library.json"
	err = library.FetchDataFromGsheets(destLibraryFile)
	if err != nil {
		fmt.Println(err)
	}

	destBookmarksFile := "../data/bookmarks.json"
	err = links.CreateBookMarksFile(destBookmarksFile)
	if err != nil {
		fmt.Println(err)
	}
}
