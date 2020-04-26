package main

import (
	"scripts/pkg/library"
)

func main() {
	destFile := "../data/library.json"
	library.FetchDataFromGsheets(destFile)
}
