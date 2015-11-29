package connectors

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/projectjane/jane/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Jira struct {
}

func (x Jira) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Jira) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	var jiraRegex = regexp.MustCompile("[a-zA-Z]{2,12}-[0-9]{1,10}")
	issues := jiraRegex.FindAllString(message.In.Text, -1)
	for _, issue := range issues {
		if connector.Debug {
			log.Print("Jira match: " + issue)
		}
		client := &http.Client{}
		auth := EncodeB64(connector.Login + ":" + connector.Pass)
		req, err := http.NewRequest("GET", "https://"+connector.Server+"/rest/api/2/issue/"+issue, nil)
		if err != nil {
			log.Print(err)
		} else {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Basic "+auth)
		}
		response, err := client.Do(req)
		if err != nil {
			log.Print(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Print(err)
		}
		var ticket Ticket
		json.Unmarshal(body, &ticket)
		if connector.Debug {
			log.Printf("Jira result: %+v", ticket)
		}
		if ticket.Fields.Status.Name == "" {
			return
		}
		message.Out.Link = "https://" + connector.Server + "/browse/" + issue
		message.Out.Text = strings.ToUpper(issue) + " - " + ticket.Fields.Summary
		message.Out.Detail = fmt.Sprintf("Status: %s\nPriority: %s\nAssignee: %s\n",
			ticket.Fields.Status.Name, ticket.Fields.Priority.Name, ticket.Fields.Assignee.DisplayName)
		publishMsgs <- message
	}
}

func (x Jira) Publish(connector models.Connector, message models.Message, target string) {
	return
}

type Ticket struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Summary     string   `json:"summary"`
	Status      Status   `json:"status"`
	Description string   `json:"description"`
	Priority    Priority `json:"priority"`
	Assignee    Assignee `json:"assignee"`
}

type Status struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Priority struct {
	Name string `json:"name"`
}

type Assignee struct {
	DisplayName string `json:"displayName"`
}

func EncodeB64(message string) string {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return string(base64Text)
}
