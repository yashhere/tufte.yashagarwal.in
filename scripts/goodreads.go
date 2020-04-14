package main

import (
  "encoding/json"
  "encoding/xml"
  "net/url"
  "path"
  "os"
  "fmt"
  "bytes"
  "mime/multipart"
  "net/http"
  "io/ioutil"
  "strconv"
  "strings"
  "time"
)

type goodreadsData struct {
  Title           string `json:"title"`
  Author          string `json:"author"`
  ImageUrl        string `json:"image_url"`
  CurrentPage     int64 `json:"current_page"`
  TotalPage       int64 `json:"total_page"`
  UpdatedAt       string `json:"updated_at"`
  FirstUpdatePage int64 `json:"first_update_location"`
  Url             string `json:"book_url"`
  PercentAtStart  int64 `json:"percent_start"`
  PercentCurrent  int64 `json:"percent_current"`
}

func formatDateTime(s string) string {
  t, err := time.Parse("2006-01-02T15:04:05-07:00", strings.TrimSpace(s))
  if err != nil {
    fmt.Println(err)
    return ""
  } else {
    return t.Format("02/01/2006")
  }
}

func createBookURL(bookTitle string) string {
  baseUrl, err := url.Parse("https://www.goodreads.com/book/title")
  if err != nil {
    fmt.Println("Malformed URL: ", err.Error())
    return ""
  }

  // Prepare Query Parameters
  params := url.Values{}
  params.Add("id", bookTitle)

  // Add Query Parameters to the URL
  baseUrl.RawQuery = params.Encode() // Escape Query Parameters

  return baseUrl.String()
}

func callAPI(url, key, userID, method string) ([]byte, error) {
  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  _ = writer.WriteField("shelf", "currently-reading")
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
  }

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------627876348093473599659484")

  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  defer res.Body.Close()

  return ioutil.ReadAll(res.Body)
}

func strToInt(s string) int64 {
  if i, err := strconv.ParseInt(s, 10, 64); err == nil {
    return i
  } else {
    return -1
  }
}

func getFirstUpdateOfMonth(id, key string) (int64) {
  url := fmt.Sprintf("https://www.goodreads.com/user_status/show/%s.xml?key=%s", id, key)
  body, err := callAPI(url, key, "", "GET")
  if err != nil {
    fmt.Errorf("%s", err)
    return -1
  }

  var c ReadStatus
  xml.Unmarshal(body, &c)

  var page string
  var index int = -1
  for idx, status := range c.UserStatus.UserStatus {
    t, _ := time.Parse("02/01/2006", formatDateTime(status.UpdatedAt))
    if time.Now().Month() == t.Month() {
      index = idx
      page = status.Page
      continue
    } else {
      break
    }
  }

  if index != -1 {
   return strToInt(page)
  }

  return -1
}

func main() {
  key := os.Getenv("GOODREADS_KEY")
  userID := os.Getenv("GOODREADS_ID")

  user_url := fmt.Sprintf("https://www.goodreads.com/user/show/%s?key=%s",userID, key)
  body, err := callAPI(user_url, key, userID, "GET")

  if err != nil {
    fmt.Errorf("%s", err)
  }
  var c GoodreadsUser
  xml.Unmarshal(body, &c)

  parsedJson := make([]*goodreadsData, 0)
  for _, book := range c.User.UserStatuses.UserStatus {
    j := new(goodreadsData)

    j.Title = book.Book.Title

    authorList := book.Book.Authors.Author

    if len(authorList) == 1 {
      j.Author = book.Book.Authors.Author[0].Name
    } else {
      j.Author = book.Book.Authors.Author[0].Name + " et al."
    }

    j.UpdatedAt = formatDateTime(book.UpdatedAt.Text)
    j.FirstUpdatePage = getFirstUpdateOfMonth(book.ID.Text, key)

    if i, err := strconv.ParseInt(book.Page.Text, 10, 64); err == nil {
      j.CurrentPage = i
    }

    j.CurrentPage = strToInt(book.Page.Text)
    j.TotalPage = strToInt(book.Book.NumPages.Text)

    u, err := url.Parse(book.Book.ImageURL)
    if err != nil {
      fmt.Println(err)
    }
    j.ImageUrl = strings.Replace(book.Book.ImageURL, path.Base(u.Path), book.Book.ID.Text + ".jpg", -1)
    j.Url = createBookURL(j.Title)

    if j.FirstUpdatePage == -1 {
      j.FirstUpdatePage = j.CurrentPage
    }
    parsedJson = append(parsedJson, j)

    j.PercentAtStart = int64(j.FirstUpdatePage*100/j.TotalPage)
    j.PercentCurrent = int64(j.CurrentPage*100/j.TotalPage)

    //b, err := json.Marshal(j)
    //if err != nil {
    //  fmt.Println(err)
    //  return
    //}
    //fmt.Println(string(b))
  }

  var file []byte
  file, err = json.MarshalIndent(parsedJson, "", " ")
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

  err = ioutil.WriteFile(dirName + "/goodreads.json", file, 0644)
  if err != nil {
  fmt.Println(err)
  }
}
