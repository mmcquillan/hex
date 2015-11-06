package listeners

import (
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
	"time"
)

func Monitor(config *models.Config, listener models.Listener) {
	defer Recovery(config, listener)
	var state = make(map[string]string)
	for _, chk := range listener.Checks {
		state[chk.Name] = "OK"
	}
	for {
		alerts := callMonitor(&state, config, listener)
		reportMonitor(alerts, &state, config, listener)
		time.Sleep(60 * time.Second)
	}
}

func callMonitor(state *map[string]string, config *models.Config, listener models.Listener) (alerts []string) {
	serverconn := true
	clientconn := &ssh.ClientConfig{
		User: listener.Login,
		Auth: []ssh.AuthMethod{
			ssh.Password(listener.Pass),
		},
	}
	if config.Debug {
		log.Print("Starting client connection for " + listener.Server)
	}
	client, err := ssh.Dial("tcp", listener.Server+":22", clientconn)
	if err != nil {
		log.Print(err)
	}
	if client == nil {
		serverconn = false
	} else {
		defer client.Close()
		for _, chk := range listener.Checks {
			if config.Debug {
				log.Print("Starting session connection for " + listener.Server + " " + chk.Name + " check")
			}
			session, err := client.NewSession()
			if err != nil {
				log.Print(err)
			}
			if session == nil {
				serverconn = false
			} else {
				defer session.Close()
				if config.Debug {
					log.Print("Starting session call for " + listener.Server + " " + chk.Name + " check")
				}
				b, err := session.CombinedOutput(chk.Check)
				if config.Debug {
					log.Print("Ending session call for " + listener.Server + " " + chk.Name + " check")
				}
				if err != nil {
					log.Print(err)
				}
				out := string(b[:])
				if config.Debug {
					log.Print("Session results for " + listener.Server + " " + chk.Name + ": " + out)
				}
				if (*state)[chk.Name] != out && !(strings.Contains(out, listener.SuccessMatch) && strings.Contains((*state)[chk.Name], listener.SuccessMatch)) {
					alerts = append(alerts, chk.Name)
					(*state)[chk.Name] = out
				}
			}
		}
	}
	if !serverconn {
		if config.Debug {
			log.Print("Cannot connect to server " + listener.Server)
		}
		out := "CRITICAL - Cannot connect to server " + listener.Server
		for _, chk := range listener.Checks {
			if (*state)[chk.Name] != out {
				alerts = append(alerts, chk.Name)
				(*state)[chk.Name] = out
			}
		}
	}
	return alerts
}

func reportMonitor(alerts []string, state *map[string]string, config *models.Config, listener models.Listener) {
	if config.Debug {
		log.Print("Starting reporting on monitroing results for " + listener.Server)
	}
	for _, a := range alerts {
		out := (*state)[a]
		var color = "NONE"
		if strings.Contains(out, listener.SuccessMatch) {
			color = "SUCCESS"
		} else if strings.Contains(out, listener.WarningMatch) {
			color = "WARN"
		} else if strings.Contains(out, listener.FailureMatch) {
			color = "FAIL"
		} else {
			color = "NONE"
		}
		for _, d := range listener.Destinations {
			if strings.Contains(out, d.Match) || d.Match == "*" {
				m := models.Message{
					Relays:      d.Relays,
					Target:      d.Target,
					Request:     "",
					Title:       listener.Name + " " + a,
					Description: out,
					Link:        "",
					Status:      color,
				}
				commands.Parse(config, m)
			}
		}
	}
}
