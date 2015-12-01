package connectors

import (
	"github.com/projectjane/jane/models"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
	"time"
)

type Monitor struct {
	Connector models.Connector
}

func (x Monitor) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	var state = make(map[string]string)
	for _, chk := range connector.Checks {
		state[chk.Name] = "OK"
	}
	for {
		alerts := callMonitor(&state, connector)
		reportMonitor(alerts, &state, commandMsgs, connector)
		time.Sleep(60 * time.Second)
	}
}

func (x Monitor) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Monitor) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Monitor) Help(connector models.Connector) (help string) {
	return
}

type alert struct {
	state string
	check string
	text  string
}

func callMonitor(state *map[string]string, connector models.Connector) (alerts []alert) {
	serverconn := true
	clientconn := &ssh.ClientConfig{
		User: connector.Login,
		Auth: []ssh.AuthMethod{
			ssh.Password(connector.Pass),
		},
	}
	if connector.Debug {
		log.Print("Starting client connection for " + connector.Server)
	}
	client, err := ssh.Dial("tcp", connector.Server+":22", clientconn)
	if err != nil {
		log.Print(err)
	}
	if client == nil {
		serverconn = false
	} else {
		defer client.Close()
		for _, chk := range connector.Checks {
			if connector.Debug {
				log.Print("Starting session connection for " + connector.Server + " " + chk.Name + " check")
			}
			session, err := client.NewSession()
			if err != nil {
				log.Print(err)
			}
			if session == nil {
				serverconn = false
			} else {
				defer session.Close()
				if connector.Debug {
					log.Print("Starting session call for " + connector.Server + " " + chk.Name + " check")
				}
				b, err := session.CombinedOutput(chk.Check)
				if connector.Debug {
					log.Print("Ending session call for " + connector.Server + " " + chk.Name + " check")
				}
				if err != nil && connector.Debug {
					log.Print(err)
				}
				out := string(b[:])
				if connector.Debug {
					log.Print("Session results for " + connector.Server + " " + chk.Name + ": " + out)
				}
				newState := "UNKNOWN"
				if strings.Contains(out, connector.SuccessMatch) {
					newState = connector.SuccessMatch
				} else if strings.Contains(out, connector.WarningMatch) {
					newState = connector.WarningMatch
				} else if strings.Contains(out, connector.FailureMatch) {
					newState = connector.FailureMatch
				}
				if (*state)[chk.Name] != newState {
					a := alert{state: newState, check: chk.Name, text: out}
					alerts = append(alerts, a)
					(*state)[chk.Name] = newState
				}
			}
		}
	}
	if !serverconn {
		if connector.Debug {
			log.Print("Cannot connect to server " + connector.Server)
		}
		out := "CRITICAL - Cannot connect to server " + connector.Server
		for _, chk := range connector.Checks {
			if (*state)[chk.Name] != "CRITICAL" {
				a := alert{state: "CRITICAL", check: chk.Name, text: out}
				alerts = append(alerts, a)
				(*state)[chk.Name] = "CRITICAL"
			}
		}
	}
	return alerts
}

func reportMonitor(alerts []alert, state *map[string]string, commandMsgs chan<- models.Message, connector models.Connector) {
	if connector.Debug {
		log.Print("Starting reporting on monitroing results for " + connector.Server)
	}
	for _, a := range alerts {
		var color = "NONE"
		if a.state == connector.SuccessMatch {
			color = "SUCCESS"
		} else if a.state == connector.WarningMatch {
			color = "WARN"
		} else if a.state == connector.FailureMatch {
			color = "FAIL"
		} else {
			color = "NONE"
		}
		var m models.Message
		m.Routes = connector.Routes
		m.In.Process = false
		m.Out.Text = connector.ID + " " + a.check
		m.Out.Detail = a.text
		m.Out.Status = color
		commandMsgs <- m
	}
}
