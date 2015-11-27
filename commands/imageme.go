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
var errorMessage = "Error retrieving image"
var animated bool

func ImageMe(msg string, command models.Command) string {
  msg = strings.TrimSpace(strings.Replace(msg, command.Match, "", 1))
  start := rand.Intn(3)
  if start < 1 {
    start = 1
  }

  if strings.Contains(command.Match, "animate") {
    animated = true;
  }

  apiKey := command.ApiKey
  cx := "&cx=" + command.Args
  returnFields := fmt.Sprintf("&fields=items(link)&start=%v", start)
  query := "&q=" + url.QueryEscape(msg)
  fields := "&searchType=image"
  if animated {
    fields += "&fileType=gif&hq=animated&tbs=itp:animated"
  }

  url := baseUrl + apiKey + cx + returnFields + query + fields

  resp, err := client.Get(url)
  if err != nil {
    log.Print(err)
    fmt.Println(err)
    return errorMessage
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    FindDeprecatedImage(msg, animated)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Print(err)
    return errorMessage
  }

  var result SearchResult
  err = json.Unmarshal(body, &result)
  if err != nil {
    log.Print(err)
    fmt.Println(err)
    return errorMessage
  }

  if len(result.Items) > 0 {
    randomLink := result.Items[rand.Intn(len(result.Items))]
    return randomLink.Link
  } else {
    return FindDeprecatedImage(msg, animated)
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

func FindDeprecatedImage(query string, animated bool) string {
  baseUrl := "https://ajax.googleapis.com/ajax/services/search/images?v=1.0&rsz=8"
  if animated {
    baseUrl += "&as_filetype=gif"
  }

  baseUrl += "&q="
  searchUrl := baseUrl + url.QueryEscape(query)

  if animated {
    searchUrl += url.QueryEscape(" animated")
  }

  resp, err := client.Get(searchUrl)
  if err != nil {
    log.Print(err)
    return errorMessage
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Print(err)
    return errorMessage
  }

  var result DeprecatedResult
  err = json.Unmarshal(body, &result)
  if err != nil {
    log.Print(err)
    return errorMessage
  }

  index := rand.Intn(len(result.ResponseData.Results))

  if len(result.ResponseData.Results) > 0 {
    return result.ResponseData.Results[index].Url
  }

  return "No results found"
}
