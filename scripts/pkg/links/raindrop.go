package links

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"scripts/pkg/utils"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

func callApi(url string, method string, token string, page int, perpage int) ([]byte, error) {
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

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("perpage", strconv.Itoa(perpage))
	req.URL.RawQuery = q.Encode()

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

func downloadBookmarks(token string, page int, perpage int) ([]byte, error) {
	baseUrl := "https://api.raindrop.io/rest/v1/raindrops/0"
	method := "GET"

	return callApi(baseUrl, method, token, page, perpage)
}

func CreateBookMarksFile(destFile string) error {
	raindropToken, ok := os.LookupEnv("RAINDROP_TOKEN")
	if ok != true {
		return errors.New("Raindrop token is not set.")
	}

	perpage := 50
	var bookmarks []interface{}
	for page := 0; ; page++ {
		data, err := downloadBookmarks(raindropToken, page, perpage)
		if err != nil {
			return err
		}

		var bm map[string]interface{}
		err = json.Unmarshal(data, &bm)
		if err != nil {
			return err
		}

		b := bm["items"].([]interface{})
		bookmarks = append(bookmarks, b...)
		if len(b) < perpage {
			break
		}
	}

	tagsToOmit := []string{
		"do-not-publish",
	}

	links := make([]*Link, 0)
	for _, b := range bookmarks {
		omittedTagFound := false
		for _, tag := range b.(map[string]interface{})["tags"].([]interface{}) {
			if utils.FindString(tagsToOmit, tag.(string)) {
				omittedTagFound = true
				break
			}
		}

		if omittedTagFound {
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

	log.Infof("Writing [%d] links to %s.", len(links), path.Base(destFile))
	file, err := json.MarshalIndent(links, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(destFile, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
