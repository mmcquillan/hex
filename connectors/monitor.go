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

func callMonitor(state *map[string]string, connector models.Connector) (alerts []string) {
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
				if (*state)[chk.Name] != out && !(strings.Contains(out, connector.SuccessMatch) && strings.Contains((*state)[chk.Name], connector.SuccessMatch)) {
					alerts = append(alerts, chk.Name)
					(*state)[chk.Name] = out
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
			if (*state)[chk.Name] != out {
				alerts = append(alerts, chk.Name)
				(*state)[chk.Name] = out
			}
		}
	}
	return alerts
}

func reportMonitor(alerts []string, state *map[string]string, commandMsgs chan<- models.Message, connector models.Connector) {
	if connector.Debug {
		log.Print("Starting reporting on monitroing results for " + connector.Server)
	}
	for _, a := range alerts {
		out := (*state)[a]
		var color = "NONE"
		if strings.Contains(out, connector.SuccessMatch) {
			color = "SUCCESS"
		} else if strings.Contains(out, connector.WarningMatch) {
			color = "WARN"
		} else if strings.Contains(out, connector.FailureMatch) {
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
