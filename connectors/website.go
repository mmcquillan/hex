package connectors

import (
	"crypto/tls"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"log"
	"net/http"
	"strings"
	"time"
)

type Website struct {
	Connector models.Connector
}

func (x Website) Listen(config *models.Config, connector models.Connector) {
	defer Recovery(config, connector)
	var state = make(map[string]string)
	for _, chk := range connector.Checks {
		state[chk.Name] = "OK"
	}
	for {
		alerts := callWebsite(&state, config, connector)
		reportWebsite(alerts, &state, config, connector)
		time.Sleep(60 * time.Second)
	}
}

func (x Website) Command(config *models.Config, message *models.Message) {
	return
}

func (x Website) Publish(config *models.Config, connector models.Connector, message models.Message, target string) {
	return
}

func callWebsite(state *map[string]string, config *models.Config, connector models.Connector) (alerts []string) {
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

func reportWebsite(alerts []string, state *map[string]string, config *models.Config, connector models.Connector) {
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
		m := models.Message{
			Routes:      connector.Routes,
			Request:     "",
			Title:       connector.ID + " " + a,
			Description: out,
			Link:        "",
			Status:      color,
		}
		commands.Parse(config, &m)
		Broadcast(config, m)
	}
}
