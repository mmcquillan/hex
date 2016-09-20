package connectors

import (
	"crypto/tls"
	"github.com/projectjane/jane/models"
	"log"
	"net/http"
	"strings"
	"time"
)

type Website struct {
}

func (x Website) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	var state = "OK"
	for {
		out := "OK"
		if connector.Debug {
			log.Print("Starting website call to " + connector.Server)
		}
		tran := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tran}
		res, err := client.Get(connector.Server)
		if err != nil {
			if connector.Debug {
				log.Print("Error call to " + connector.Server + " with " + err.Error())
			}
			out = "CRITICAL " + err.Error()
		} else {
			defer res.Body.Close()
			if connector.Debug {
				log.Print("Completed website call to " + connector.Server + " with " + res.Status)
			}
			if res.StatusCode == 200 {
				out = "OK"
			} else {
				out = "CRITICAL " + res.Status
			}
		}
		if state != out {
			if connector.Debug {
				log.Print("Reporting alert for " + connector.ID)
			}
			var color = "NONE"
			if strings.Contains(out, "OK") {
				color = "SUCCESS"
			} else if strings.Contains(out, "CRITICAL") {
				color = "FAIL"
			} else {
				color = "NONE"
			}
			var m models.Message
			m.In.ConnectorType = connector.Type
			m.In.ConnectorID = connector.ID
			m.In.Tags = connector.Tags
			m.In.Process = false
			m.Out.Text = connector.ID
			m.Out.Detail = out
			m.Out.Status = color
			commandMsgs <- m
		}
		state = out
		time.Sleep(60 * time.Second)
	}
}

func (x Website) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Website) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Website) Help(connector models.Connector) (help string) {
	return
}
