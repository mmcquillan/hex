package connectors

import (
	"github.com/projectjane/jane/models"
	"log"
	"net/smtp"
)

type Email struct {
	Connector models.Connector
}

func (x Email) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Email) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Email) Publish(connector models.Connector, message models.Message, target string) {
	if target == "" {
		log.Print("No email provided to the email connector")
	} else {
		auth := smtp.PlainAuth("", connector.Login, connector.Pass, connector.Server)
		to := []string{target}
		msg := []byte("To: " + target + "\r\n" +
			"Subject: " + message.Out.Text + "\r\n" +
			"\r\n" + message.Out.Detail + "\r\n\r\n" + message.Out.Link + "\r\n")
		err := smtp.SendMail(connector.Server, auth, connector.From, to, msg)
		if err != nil {
			log.Print(err)
		}
	}
}
