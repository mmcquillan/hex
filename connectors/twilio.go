package connectors

import (
	"fmt"
	"io/ioutil"
	"log"
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
	values := buildURL(target, connector.From, message.Out.Text)

	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(values.Encode()))
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if connector.Debug {
		log.Println(string(body))
	}
}

// Help Twilio help information
func (x Twilio) Help(connector models.Connector) (help string) {
	help += fmt.Sprintf("Twilioooo\n", connector.Server, connector.Port)
	return help
}

func buildURL(toNumber, fromNumber, body string) url.Values {
	values := url.Values{}
	values.Set("To", toNumber)
	values.Set("From", fromNumber)
	values.Set("Body", body)
	return values
}
