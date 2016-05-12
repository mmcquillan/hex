package connectors

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/projectjane/jane/models"
)

// Twilio Struct for manipulating the webhook connector
type Twilio struct {
}

// Listen Twilio listener
func (x Twilio) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

// Command Twilio command parser
func (x Twilio) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

// Publish Twilio publisher
func (x Twilio) Publish(connector models.Connector, message models.Message, target string) {
	client := &http.Client{}

	accountSid := connector.Key
	authToken := connector.Pass
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	body := buildURL(target, message.Out.Text)

	req, _ := http.NewRequest("POST", urlStr, body)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client.Do(req)
}

// Help Twilio help information
func (x Twilio) Help(connector models.Connector) (help string) {
	help += fmt.Sprintf("Webhooks enabled at %s:%s/webhook/\n", connector.Server, connector.Port)
	return help
}

func buildURL(toNumber, body string) *strings.Reader {
	values := url.Values{}
	values.Set("TO", toNumber)
	values.Set("FROM", "MY NUMBER")
	values.Set("Body", body)
	rb := *strings.NewReader(values.Encode())
	return &rb
}
