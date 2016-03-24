package connectors

import (
	"github.com/hpcloud/tail"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
)

type Logging struct {
}

func (x Logging) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	t, err := tail.TailFile(connector.File, tail.Config{Follow: true, Location: SeekInfo{Offset: 0, Whence: 2}})
	if err != nil {
		log.Print(err)
	}
	for line := range t.Lines {
		for _, chk := range connector.Checks {
			if match, _ := parse.Match(chk.Check, line.Text); match {
				var m models.Message
				m.Routes = connector.Routes
				m.In.Process = false
				m.Out.Text = connector.File + ": " + chk.Name
				m.Out.Detail = line.Text
				commandMsgs <- m
			}

		}
	}
}

func (x Logging) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Logging) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Logging) Help(connector models.Connector) (help string) {
	return
}
