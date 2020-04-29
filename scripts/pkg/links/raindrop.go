package links

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	utils "scripts/pkg/utils"
)

type Link struct {
	Title     string     `json:"title,omitempty"`
	Url       string     `json:"url,omitempty"`
	DateAdded *time.Time `json:"date_added,omitempty"`
	Excerpt   string     `json:"excerpt,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
	Domain    string     `json:"domain,omitempty"`
	Important bool       `json:"important,omitempty"`
}

type Collection struct {
	Name string `json:"name,omitempty"`
	Id   int    `json:"id,omitempty"`
}

func callApi(url string, method string, token string) ([]byte, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func getCollections(token string, collectionToOmit []string) ([]int, error) {
	url := "https://api.raindrop.io/rest/v1/collections"
	method := "GET"

	body, err := callApi(url, method, token)
	if err != nil {
		return nil, err
	}

	// file, err := ioutil.ReadFile("pkg/links/collections.json")
	// if err != nil {
	// 	return nil, err
	// }

	var c map[string]interface{}

	// err = json.Unmarshal([]byte(file), &c)
	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	collections := make([]int, 0)
	for _, entry := range c["items"].([]interface{}) {

		// if the collectionName of entry == any of collectionToOmit, then don't include that id in collections
		if utils.FindString(collectionToOmit, strings.ToLower(entry.(map[string]interface{})["title"].(string))) {
			continue
		}

		collections = append(collections, int(entry.(map[string]interface{})["_id"].(float64)))
	}

	return collections, nil
}

func downloadBookmarks(token string) ([]byte, error) {
	url := "https://api.raindrop.io/rest/v1/raindrops/0"
	method := "GET"

	return callApi(url, method, token)
}

func CreateBookMarksFile(destFile string) error {
	raindropToken, ok := os.LookupEnv("RAINDROP_TOKEN")
	if ok != true {
		return errors.New("Raindrop token is not set.")
	}

	bookmarksData, err := downloadBookmarks(raindropToken)

	if err != nil {
		return err
	}

	var bm map[string]interface{}
	err = json.Unmarshal(bookmarksData, &bm)

	// fmt.Println(utils.PrettyJson(bm))
	// file, err := ioutil.ReadFile("pkg/links/bookmarks_data.json")
	if err != nil {
		return err
	}

	// var bm map[string]interface{}

	// _ = json.Unmarshal([]byte(file), &bm)

	collectionToOmit := []string{"unread", "to read later", "import"}
	collections, err := getCollections(raindropToken, collectionToOmit)

	if err != nil {
		return err
	}

	links := make([]*Link, 0)
	for _, b := range bm["items"].([]interface{}) {
		if !utils.FindInt(collections, int(b.(map[string]interface{})["collectionId"].(float64))) {
			continue
		}

		dateAdded, err := time.Parse(time.RFC3339, strings.TrimSpace(b.(map[string]interface{})["created"].(string)))
		if err != nil {
			continue
		}

		link := new(Link)
		link.Title = b.(map[string]interface{})["title"].(string)
		link.Url = b.(map[string]interface{})["link"].(string)
		link.DateAdded = &dateAdded
		link.Excerpt = b.(map[string]interface{})["excerpt"].(string)

		for _, tag := range b.(map[string]interface{})["tags"].([]interface{}) {
			link.Tags = append(link.Tags, tag.(string))
		}

		link.Domain = b.(map[string]interface{})["domain"].(string)

		if b.(map[string]interface{})["important"] != nil {
			link.Important = b.(map[string]interface{})["important"].(bool)
		}
		links = append(links, link)
	}

	file, _ := json.MarshalIndent(links, "", " ")

	err = ioutil.WriteFile(destFile, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
