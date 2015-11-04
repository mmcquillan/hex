package listeners

import (
	"bytes"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
	"time"
)

func Monitor(config *models.Config, listener models.Listener) {
	defer Recovery(config, listener)
	var state = "OK"
	user, pass, server, chk := monitorResource(listener.Resource)
	clicon := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}
	for {
		var report = false
		var out = ""
		client, err := ssh.Dial("tcp", server+":22", clicon)
		if err != nil {
			log.Println(err)
		}
		if client == nil {
			out = "CRITICAL - Client cannot connect"
		} else {
			session, err := client.NewSession()
			if err != nil {
				log.Println(err)
			}
			if session == nil {
				out = "CRITICAL - Session cannot connect"
			} else {
				defer session.Close()
				var b bytes.Buffer
				session.Stdout = &b
				session.Run(chk)
				//err := ^^^
				//if err != nil {
				// leave for future debug logging
				//log.Println(err)
				//}
				out = b.String()
			}
		}
		var color = "NONE"
		if strings.Contains(out, "OK") {
			if state != "OK" {
				report = true
				state = "OK"
				color = "SUCCESS"
			}
		} else if strings.Contains(out, "WARNING") {
			if state != "WARNING" {
				report = true
				state = "WARNING"
				color = "WARN"
			}
		} else if strings.Contains(out, "CRITICAL") {
			if state != "CRITICAL" {
				report = true
				state = "CRITICAL"
				color = "FAIL"
			}
		} else {
			if state != "UNKNOWN" {
				report = true
				state = "UNKNOWN"
				color = "NONE"
			}
		}
		if report {
			for _, d := range listener.Destinations {
				if strings.Contains(out, d.Match) || d.Match == "*" {
					m := models.Message{
						Relays:      d.Relays,
						Target:      d.Target,
						Request:     "",
						Title:       listener.Name,
						Description: out,
						Link:        "",
						Status:      color,
					}
					commands.Parse(config, m)
				}
			}
		}
		time.Sleep(90 * time.Second)
	}
}

func monitorResource(resource string) (user string, pass string, server string, chk string) {
	if strings.Contains(resource, ":") {
		u := strings.Split(resource, ":")
		user = u[0]
		resource = strings.TrimSpace(strings.Replace(resource, user+":", "", 1))
	}
	if strings.Contains(resource, "@") {
		p := strings.Split(resource, "@")
		pass = p[0]
		resource = strings.TrimSpace(strings.Replace(resource, pass+"@", "", 1))
	}
	if strings.Contains(resource, "|") {
		s := strings.Split(resource, "|")
		server = s[0]
		chk = strings.TrimSpace(strings.Replace(resource, server+"|", "", 1))
	}
	return user, pass, server, chk
}
