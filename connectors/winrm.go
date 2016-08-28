package connectors

import (
	"log"
	"os"

	"github.com/masterzen/winrm"
	"github.com/projectjane/jane/models"
)

//WinRM Struct representing a WinRM Connector
type WinRM struct {
	Client *winrm.Client
}

//Listen Standard listen
func (x WinRM) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	var err error

	endpoint := winrm.NewEndpoint(connector.Server, 5985, false, false, nil, nil, nil, 0)
	x.Client, err = winrm.NewClient(endpoint, connector.Login, connector.Pass)
	if err != nil {
		log.Println("Error connecting to endpoint:", err)
	}
}

//Command Standard command parser
func (x WinRM) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	// shell, err := x.Client.CreateShell()
	// if err != nil {
	// 	log.Println("Error creating shell:", err)
	// }
	// defer shell.Close()

	if x.Client == nil {
		log.Println("Client is nil. Connecting...")

		var err error

		endpoint := winrm.NewEndpoint(connector.Server, 5985, false, false, nil, nil, nil, 0)
		x.Client, err = winrm.NewClient(endpoint, connector.Login, connector.Pass)
		if err != nil {
			log.Println("Error connecting to endpoint:", err)
		}

		log.Println(x.Client)
	}

	_, err := x.Client.RunWithInput("ipconfig", os.Stdout, os.Stderr, os.Stdin)
	if err != nil {
		panic(err)
	}

}

//Publish Not implemented
func (x WinRM) Publish(connector models.Connector, message models.Message, target string) {
	return
}

//Help Returns help information for the connector
func (x WinRM) Help(connector models.Connector) (help string) {
	help += "winrm"
	return help
}
