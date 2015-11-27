package commands

import (
  "fmt"
  "net/http"
  "net/url"
  "encoding/json"
  "io/ioutil"
  "math/rand"
  "github.com/projectjane/jane/models"
  "log"
  "strings"
)

type SearchResult struct {
  Items []Items `json:"items"`
}

type Items struct {
  Link string `json:"link"`
}

var client = &http.Client {

}

var baseUrl = "https://www.googleapis.com/customsearch/v1?key="
var cx = "&cx=001525354505425389698:itoebrubry4"

func ImageMe(msg string, command models.Command) string {
  msg = strings.TrimSpace(strings.Replace(msg, command.Match, "", 1))
  start := rand.Intn(3)

  apiKey := command.ApiKey
  returnFields := fmt.Sprintf("&fields=items(link)&start=%v", start)
  query := "&q=" + url.QueryEscape(msg)
  fields := "&searchType=image"
  url := baseUrl + apiKey + cx + returnFields + query + fields

  resp, err := client.Get(url)
  if err != nil {
    log.Print(err)
    return "Error retrieving image"
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    FindDeprecatedImage(msg)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Print(err)
    return "Error retrieving image"
  }

  var result SearchResult
  err = json.Unmarshal(body, &result)
  if err != nil {
    log.Print(err)
    return "Error retrieving image"
  }

  if len(result.Items) > 0 {
    randomLink := result.Items[rand.Intn(len(result.Items))]
    return randomLink.Link
  } else {
    return FindDeprecatedImage(msg)
  }
}

type DeprecatedResult struct {
  ResponseData ResponseData `json:"responseData"`
}

type ResponseData struct {
  Results []Result `json:"results"`
}

type Result struct {
  Url string `json:"url"`
}

func FindDeprecatedImage(query string) string {
  baseUrl := "https://ajax.googleapis.com/ajax/services/search/images?v=1.0&rsz=8&q="
  url := baseUrl + url.QueryEscape(query)

  resp, err := client.Get(url)
  if err != nil {
    log.Print(err)
    return "Error retrieving image"
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Print(err)
    return "Error retrieving image"
  }

  var result DeprecatedResult
  err = json.Unmarshal(body, &result)
  if err != nil {
    log.Print(err)
    return "Error retrieving image"
  }

  index := rand.Intn(len(result.ResponseData.Results))

  if len(result.ResponseData.Results) > 0 {
    return result.ResponseData.Results[index].Url
  }

  return "No results found"
}
