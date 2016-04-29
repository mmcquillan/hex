package connectors

import (
	"github.com/projectjane/jane/models"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"log"
)

type Hipchat struct {
}

func (x Hipchat) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	//hc := hipchat.NewClient(connector.Key)
	return
}

func (x Hipchat) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Hipchat) Publish(connector models.Connector, message models.Message, target string) {
	hc := hipchat.NewClient(connector.Key)
	msg := &hipchat.NotificationRequest{Message: message.Out.Text}
	resp, err = hc.Room.Notification(target, msg)
	if err != nil {
		log.Print(err)
		log.Print(resp)
	}
}

func (x Hipchat) Help(connector models.Connector) (help string) {
	return
}
