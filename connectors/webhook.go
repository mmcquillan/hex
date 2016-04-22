package connectors

import (
	"log"
	"strings"
	"fmt"
  "net/http"
	"github.com/projectjane/jane/models"
)

type Webhook struct {
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
  segs := strings.Split(r.URL.Path, "/")
  if len(segs) < 2 {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Route not found")
    return
  }

  // command := segs[2]
  // provider := segs[3]

}

func (x Webhook) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
  defer Recovery(connector)
	if connector.Debug {
		log.Println("Starting Webhook connector...")
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", connector.Port),
		Handler: nil,
	}

  http.HandleFunc("/webhook/", webhookHandler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// for {
	// 	select {
	// 	case msg := <-rtm.IncomingEvents:
	// 		switch ev := msg.Data.(type) {
	// 		case *slack.MessageEvent:
	// 			if ev.User != "" {
  //
	// 				if connector.Debug {
	// 					log.Print("Evaluating incoming slack message")
	// 				}
  //
	// 				var r []models.Route
	// 				r = append(r, models.Route{Match: "*", Connectors: connector.ID, Target: ev.Channel})
	// 				for _, cr := range connector.Routes {
	// 					r = append(r, cr)
	// 				}
  //
	// 				var m models.Message
	// 				m.Routes = r
	// 				m.In.Source = connector.ID
	// 				m.In.User = ev.User
	// 				m.In.Text = html.UnescapeString(ev.Text)
	// 				m.In.Process = true
	// 				commandMsgs <- m
  //
	// 			}
	// 		}
	// 	}
	// }
}

func (x Webhook) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if message.In.Process {
		for _, c := range connector.Commands {
			if strings.HasPrefix(strings.ToLower(message.In.Text), strings.ToLower(c.Match)) {
				environment := Environment {
					Address: connector.Server,
					Password: connector.Pass,
					DB: 0,
				}

				status := FlushDb(environment)
				log.Println(status.String())
				message.Out.Text = fmt.Sprintf("Redis Server: %s\nStatus:%s", connector.Server, status.String())
				publishMsgs <- message
				return
			}
		}
	}
}

func (x Webhook) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Webhook) Help(connector models.Connector) (help string) {
	help += "jane flushdb <environment> - flushes the environments redis db\n"
	return help
}
