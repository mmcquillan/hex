package outputs

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hexbotio/hex/models"
)

type Twilio struct {
}

func (x Twilio) Write(outputMsgs <-chan models.Message, service models.Service) {
	for {
		message := <-outputMsgs
		client := &http.Client{}
		textmsg := strings.Join(message.Response[:], "\n")
		accountSid := service.Config["Key"]
		authToken := service.Config["Pass"]
		urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

		values := url.Values{}
		values.Set("To", message.Inputs["hex.output"])
		values.Set("From", service.Config["From"])
		values.Set("Body", textmsg)

		req, err := http.NewRequest("POST", urlStr, strings.NewReader(values.Encode()))
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			log.Println(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Print("TwilioSent - " + string(body))
	}
}
