package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type Rss struct {
	XMLName    xml.Name `xml:"rss"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	Atom       string   `xml:"atom,attr"`
	Dc         string   `xml:"dc,attr"`
	Letterboxd string   `xml:"letterboxd,attr"`
	Channel    struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Href string `xml:"href,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description string `xml:"description"`
		Item        []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			Guid  struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate      string `xml:"pubDate"`
			WatchedDate  string `xml:"watchedDate"`
			Rewatch      string `xml:"rewatch"`
			FilmTitle    string `xml:"filmTitle"`
			FilmYear     string `xml:"filmYear"`
			MemberRating string `xml:"memberRating"`
			Description  string `xml:"description"`
			Creator      string `xml:"creator"`
		} `xml:"item"`
	} `xml:"channel"`
}

type letterboxdData struct {
	Title       	string `json:"title"`
	Rating       	string `json:"rating"`
	WatchedDate  	string `json:"watched_date"`
	Image 			string `json:"image"`
}

func readRSSFromUrl(url string) (*Rss, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data *Rss
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	url := "https://letterboxd.com/yagr/rss/"
	data, err := readRSSFromUrl(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res []*letterboxdData

	var re = regexp.MustCompile(`(?m)src\s*=\s*"(.+?)"`)
	for _, movie := range data.Channel.Item {
		t, _ := time.Parse("2006-01-02", strings.TrimSpace(movie.WatchedDate))
		now := time.Now()

		if t.Month() == 3 && t.Year() == now.Year() {
			imageUrl := re.FindAllString(movie.Description, -1)[0]				// taking only first occurrence
			imageUrl = strings.Replace(imageUrl, "src=", "", 1)			// removing "src="
			imageUrl = imageUrl[1 : len(imageUrl)-1]

			r := new(letterboxdData)
			r.Title = movie.FilmTitle
			r.WatchedDate = t.Format("02/01/2006")
			r.Image = imageUrl
			r.Rating = movie.MemberRating

			res = append(res, r)
		}
	}

	var file []byte
	file, err = json.MarshalIndent(res, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	dirName := "./data"
	_, err = os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			fmt.Errorf("%s", errDir)
			os.Exit(1)
		}
	}

	err = ioutil.WriteFile(dirName + "/letterboxd.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}