package commands

import (
  "strings"
  "fmt"
  "github.com/mmcquillan/jane/models"
  "net/http"
  "encoding/xml"
  "io/ioutil"
  "log"
)

type Query struct {
  QueryResult string `xml:"success,attr"`
  PodList []Pod `xml:"pod"`
  Datatypes string `xml:"datatypes,attr"`
}

type Pod struct {
  Title string `xml:"title,attr"`
  SubpodList []Subpod `xml:"subpod"`
}

type Subpod struct {
  Title string `xml:"title,attr"`
  Plaintext string `xml:"plaintext"`
  Image Image `xml:"img"`
}

type Image struct {
  Source string `xml:"src,attr"`
}

var errorResult = "Error processing request"
var baseQuery = "http://api.wolframalpha.com/v2/query?appid="
var input = "&input="
var returnTypes = "&format=image,plaintext"

func Wolfram(msg string, command models.Command) (results string) {
  msg = strings.TrimSpace(strings.Replace(msg, command.Match, "", 1))
  encodedRequest := strings.Replace(msg, " ", "%20", -1)
  getRequest := baseQuery + command.ApiKey + input + encodedRequest + returnTypes

  resp, err := http.Get(getRequest)
  if err != nil {
    log.Print(err)
    return errorResult
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Print(err)
    return errorResult
  }

  var query Query
  err = xml.Unmarshal(body, &query)
  if err != nil {
    log.Print(err)
    return errorResult
  }

  if query.QueryResult == "false" {
    return "Failed to return a result"
  }

  var result string
  if query.Datatypes == "Math" {
    result = ParseWolframAlphaMathQuery(query)
  } else {
    result = ParseWolframAlphaQuery(query)
  }

  return result
}

func ParseWolframAlphaMathQuery(query Query) string {
  result := ""

  for _, pod := range query.PodList {
    if len(pod.SubpodList) == 0 {
      continue
    }
    if pod.Title != "Result" {
      continue
    }
    for _, subpod := range pod.SubpodList {
      if subpod.Plaintext == "{}" {
        continue
      }
      result += fmt.Sprintf("     %s\n", subpod.Plaintext)
    }
  }

  return result
}

func ParseWolframAlphaQuery(query Query) string {
  result := ""

  for _, pod := range query.PodList {
    if len(pod.SubpodList) == 0 {
      continue
    }
    if pod.Title == "Input interpretation" {
      continue
    }
    for _, subpod := range pod.SubpodList {
      if subpod.Plaintext == "{}" {
        continue
      }
      result += fmt.Sprintf("     %s\n", subpod.Plaintext)
      // result += fmt.Sprintf("     Image: %s\n\n", subpod.Image.Source)
    }
  }

  return result
}
