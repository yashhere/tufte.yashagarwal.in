package main

import (
  "encoding/json"
  "encoding/xml"
  "os"
  "fmt"
  "bytes"
  "mime/multipart"
  "net/http"
  "net/url"
  "io/ioutil"
  "path"
  "strings"
)

type GoodreadsResponse struct {
  XMLName xml.Name `xml:"GoodreadsResponse"`
  Text    string   `xml:",chardata"`
  Request struct {
    Text           string `xml:",chardata"`
    Authentication string `xml:"authentication"`
    Key            string `xml:"key"`
    Method         string `xml:"method"`
  } `xml:"Request"`
  Books []struct {
    Text        string `xml:",chardata"`
    Start       string `xml:"start,attr"`
    End         string `xml:"end,attr"`
    Total       string `xml:"total,attr"`
    Numpages    string `xml:"numpages,attr"`
    Currentpage string `xml:"currentpage,attr"`
    Book        []struct {
      Text string `xml:",chardata"`
      ID   struct {
        Text string `xml:",chardata"`
        Type string `xml:"type,attr"`
      } `xml:"id"`
      ISBN struct {
        Text string `xml:",chardata"`
        Nil  string `xml:"nil,attr"`
      } `xml:"isbn"`
      Isbn13 struct {
        Text string `xml:",chardata"`
        Nil  string `xml:"nil,attr"`
      } `xml:"isbn13"`
      TextReviewsCount struct {
        Text string `xml:",chardata"`
        Type string `xml:"type,attr"`
      } `xml:"text_reviews_count"`
      URI                string `xml:"uri"`
      Title              string `xml:"title"`
      TitleWithoutSeries string `xml:"title_without_series"`
      ImageURL           string `xml:"image_url"`
      SmallImageURL      string `xml:"small_image_url"`
      LargeImageURL      string `xml:"large_image_url"`
      Link               string `xml:"link"`
      NumPages           string `xml:"num_pages"`
      Format             string `xml:"format"`
      EditionInformation string `xml:"edition_information"`
      Publisher          string `xml:"publisher"`
      PublicationDay     string `xml:"publication_day"`
      PublicationYear    string `xml:"publication_year"`
      PublicationMonth   string `xml:"publication_month"`
      AverageRating      string `xml:"average_rating"`
      RatingsCount       string `xml:"ratings_count"`
      Description        string `xml:"description"`
      Authors            struct {
        Text   string `xml:",chardata"`
        Author struct {
          Text     string `xml:",chardata"`
          ID       string `xml:"id"`
          Name     string `xml:"name"`
          Role     string `xml:"role"`
          ImageURL struct {
            Text    string `xml:",chardata"`
            Nophoto string `xml:"nophoto,attr"`
          } `xml:"image_url"`
          SmallImageURL struct {
            Text    string `xml:",chardata"`
            Nophoto string `xml:"nophoto,attr"`
          } `xml:"small_image_url"`
          Link             string `xml:"link"`
          AverageRating    string `xml:"average_rating"`
          RatingsCount     string `xml:"ratings_count"`
          TextReviewsCount string `xml:"text_reviews_count"`
        } `xml:"author"`
      } `xml:"authors"`
      Published string `xml:"published"`
      Work      struct {
        Text string `xml:",chardata"`
        ID   string `xml:"id"`
        URI  string `xml:"uri"`
      } `xml:"work"`
    } `xml:"book"`
  } `xml:"books"`
}

type jsonStruct struct {
  Title       string `json:"title"`
  ImageUrl       string `json:"image_url"`
}

func main() {
  key := os.Getenv("GOODREADS_KEY")
  user_id := os.Getenv("GOODREADS_ID")
  goodreads_url := fmt.Sprintf("https://www.goodreads.com/review/list/%s.xml?key=%s",user_id, key)
  method := "GET"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  _ = writer.WriteField("shelf", "currently-reading")
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
  }

  client := &http.Client {
  }
  req, err := http.NewRequest(method, goodreads_url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------627876348093473599659484")

  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  if err != nil {
    fmt.Errorf("%s", err)
  }
  var c GoodreadsResponse
  xml.Unmarshal(body, &c)

  parsedJson := make([]*jsonStruct, 0)
  for _, book := range c.Books {
    for _, b := range book.Book {
      j := new(jsonStruct)

      u, err := url.Parse(b.ImageURL)
      if err != nil {
        fmt.Println(err)
      }

      j.Title = b.Title
      j.ImageUrl = strings.Replace(b.ImageURL, path.Base(u.Path), b.ID.Text + ".jpg", -1)

      parsedJson = append(parsedJson, j)
    }
  }

  var file []byte
  file, err = json.MarshalIndent(parsedJson, "", " ")
  if err != nil {
    fmt.Println(err)
  }

  err = ioutil.WriteFile("./data/goodreads.json", file, 0644)
  if err != nil {
    fmt.Println(err)
  }
}
