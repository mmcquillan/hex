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

// Twilio Empty struct
type Twilio struct {
}

// Listen Twilio does not listen for anything right now
func (x Twilio) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

// Command Twilio does not process any commands right now
func (x Twilio) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

// Publish Twilio publisher to push messages via Twilio REST Api
func (x Twilio) Publish(connector models.Connector, message models.Message, target string) {
	client := &http.Client{}

	req, err := buildRequest(target, message.Out.Text, connector)
	if err != nil {
		log.Println(err)
	}

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
	help += fmt.Sprintf("Twilio messages sent from %s\n", connector.From)
	return help
}

func buildRequest(toNumber, body string, connector models.Connector) (*http.Request, error) {
	accountSid := connector.Key
	authToken := connector.Pass
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	values := url.Values{}
	values.Set("To", toNumber)
	values.Set("From", connector.From)
	values.Set("Body", body)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(values.Encode()))
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}
