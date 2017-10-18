package commands

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hexbotio/hex/models"
)

func Version(message *models.Message, config models.Config) {
	response := "Version: Non Standard Build"
	if config.Version != "" {
		response = "Version: " + config.Version
	}
	var u url.URL
	u.Scheme = "https"
	u.Host = "hexbot.io"
	u.Path = "/downloads/latest"
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		config.Logger.Error("Version New Request - " + err.Error())
	} else {
		req.Header.Add("User-Agent", "HexBot")
		client := &http.Client{}
		client.Timeout = time.Duration(3) * time.Second
		resp, err := client.Do(req)
		if err != nil {
			config.Logger.Error("Version Response - " + err.Error())
		} else {
			if resp.StatusCode != 200 {
				config.Logger.Error("Version Response Error")
			} else {
				rawbody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					config.Logger.Error("Version Body Read - " + err.Error())
				}
				body := strings.TrimSpace(string(rawbody))
				if body != config.Version {
					response = response + " (latest is " + body + ")"
				} else {
					response = response + " (current)"
				}
			}
		}
	}
	message.Outputs = append(message.Outputs, models.Output{
		Rule:     "version",
		Response: response,
	})
}
