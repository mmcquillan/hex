package services

import (
	"github.com/projectjane/jane/models"
	"log"
	"net/smtp"
	"strings"
)

type Email struct {
	Connector models.Connector
}

func (x Email) Input(inputMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Email) Command(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Email) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	for {
		message := <-outputMsgs
		for _, target := range strings.Split(message.Out.Target, ",") {
			if target == "" {
				log.Print("No email provided to the email connector")
			} else {
				auth := smtp.PlainAuth("", connector.Login, connector.Pass, connector.Server)
				to := []string{target}
				msg := []byte("To: " + target + "\r\n" +
					"Subject: " + message.Out.Text + "\r\n" +
					"\r\n" + message.Out.Detail + "\r\n\r\n" + message.Out.Link + "\r\n")
				err := smtp.SendMail(connector.Server+":"+connector.Port, auth, connector.From, to, msg)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func (x Email) Help(connector models.Connector) (help []string) {
	return
}
