package commands

import (
  "fmt"
  "net/http"
  "encoding/base64"
  "encoding/json"
  "io/ioutil"
  "strings"
  "github.com/mmcquillan/jane/models"
)

type Ticket struct {
  Key string `json:"key"`
  Fields Fields `json:"fields"`
}

type Fields struct {
  Summary string `json:"summary"`
  Assignee Assignee `json:"assignee"`
  Status Status `json:"status"`
  Description string `json:"description"`
  Creator Creator `json:"creator"`
  Priority Priority `json:"priority"`
}

type Assignee struct {
  Name string `json:"name"`
}

type Creator struct {
  Name string `json:"name"`
}

type IssueType struct {
  Name string `json:"name"`
}

type Status struct {
  Description string `json:"description"`
  Name string `json:"name"`
}

type Priority struct {
  Name string `json:"name"`
}

func Jira(msg string, command models.Command) string {
  msg = strings.TrimSpace(strings.Replace(msg, command.Match, "", 1))
  client := &http.Client {

  }

  baseUrl := command.Args
  issueNumber := msg

  auth := command.ApiKey
  encodedAuth := EncodeB64(auth)

  req, err := http.NewRequest("GET", baseUrl + "rest/api/2/issue/" + issueNumber, nil)
  if err != nil {
    fmt.Println(err)
  } else {
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", "Basic " + encodedAuth)
  }

  response, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }

  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  }

  var ticket Ticket
  json.Unmarshal(body, &ticket)

  link := baseUrl + "browse/" + issueNumber

  return fmt.Sprintf("Status: %s\nAssignee: %s\nPriority: %s\nSummary: %s\nDescription: %s\n\n%s\n",
    ticket.Fields.Status.Name, ticket.Fields.Assignee.Name, ticket.Fields.Priority.Name,
    ticket.Fields.Summary, ticket.Fields.Description, link)
}

func EncodeB64(message string) string {
    base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
    base64.StdEncoding.Encode(base64Text, []byte(message))
    return string(base64Text)
}
