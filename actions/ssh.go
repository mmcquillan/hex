package actions

import (
	"log"
	"time"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"golang.org/x/crypto/ssh"
)

//Exec struct
type Ssh struct {
}

//Action function
func (x Ssh) Act(action models.Action, message *models.Message, config *models.Config) {

	// have to lookup each matching server
	for _, service := range config.Services {
		if service.Active && service.Type == "ssh" {
			if parse.Match(action.Service, service.Name) {

				command := parse.Substitute(action.Command, message.Inputs)
				out := ""

				// setup the server connection
				serverconn := true
				clientconn := &ssh.ClientConfig{
					User: service.Config["Login"],
					Auth: []ssh.AuthMethod{
						ssh.Password(service.Config["Pass"]),
					},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				}
				port := "22"
				if service.Config["Port"] != "" {
					port = service.Config["Port"]
				}
				retryCounter := 1
				for retryCounter <= 3 {
					client, err := ssh.Dial("tcp", service.Config["Server"]+":"+port, clientconn)
					if err != nil {
						log.Print(err)
					}
					if client == nil {
						serverconn = false
					} else {
						defer client.Close()
						session, err := client.NewSession()
						if err != nil {
							log.Print(err)
						}
						if session == nil {
							serverconn = false
						} else {
							defer session.Close()
							b, err := session.CombinedOutput(command)
							if err != nil {
								log.Print(err)
							}
							out = string(b[:])
						}
					}
					if serverconn {
						retryCounter = 999
					} else {
						time.Sleep(time.Duration(3*retryCounter) * time.Second)
						retryCounter += 1
					}
				}
				if !serverconn {
					out = "ERROR - Cannot connect to server " + service.Name
				}
				message.Response = append(message.Response, out)
				message.Success = ActionSuccess(out, action.Success, action.Failure)
			}
		}
	}
}
