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
	var state = make(map[string]string)
	for _, chk := range connector.Checks {
		state[chk.Name] = "OK"
	}
	for {
		alerts := callWebsite(&state, connector)
		reportWebsite(alerts, &state, commandMsgs, connector)
		time.Sleep(60 * time.Second)
	}
}

func (x Website) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Website) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Website) Help(connector models.Connector) (help string) {
	return
}

func callWebsite(state *map[string]string, connector models.Connector) (alerts []string) {
	for _, chk := range connector.Checks {
		out := "OK"
		if connector.Debug {
			log.Print("Starting website call to " + chk.Check)
		}
		tran := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tran}
		res, err := client.Get(chk.Check)
		defer res.Body.Close()
		if err != nil {
			if connector.Debug {
				log.Print("Error call to " + chk.Check + " with " + err.Error())
			}
			out = "CRITICAL " + err.Error()
		} else {
			if connector.Debug {
				log.Print("Completed website call to " + chk.Check + " with " + res.Status)
			}
			if res.StatusCode == 200 {
				out = "OK " + chk.Check + " " + res.Status
			} else {
				out = "CRITICAL " + chk.Check + " " + res.Status
			}
		}
		if (*state)[chk.Name] != out && (*state)[chk.Name] != "OK" {
			if connector.Debug {
				log.Print("Reporting alert for " + chk.Name)
			}
			alerts = append(alerts, chk.Name)
		}
		(*state)[chk.Name] = out
	}
	return alerts
}

func reportWebsite(alerts []string, state *map[string]string, commandMsgs chan<- models.Message, connector models.Connector) {
	if connector.Debug {
		log.Print("Starting reporting on website results for " + connector.Server)
	}
	for _, a := range alerts {
		out := (*state)[a]
		var color = "NONE"
		if strings.Contains(out, "OK") {
			color = "SUCCESS"
		} else if strings.Contains(out, "CRITICAL") {
			color = "FAIL"
		} else {
			color = "NONE"
		}
		var m models.Message
		m.Routes = connector.Routes
		m.In.Process = false
		m.Out.Text = connector.ID + " " + a
		m.Out.Detail = out
		m.Out.Status = color
		commandMsgs <- m
	}
}
