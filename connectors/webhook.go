package connectors

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/projectjane/jane/models"
)

type Webhook struct {
	CommandMsgs chan<- models.Message
	PublishMsgs chan<- models.Message
}

var webhook Webhook

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var segs []string
	webhookString := r.URL.Path[9:]
	segs = strings.Split(webhookString, "/")
	log.Println(segs)
	log.Println(len(segs))
	if len(segs) < 2 {

		log.Println("About to split")
		segs = strings.Split(webhookString, "+")
		log.Println(segs)
		log.Println(len(segs))

		if len(segs) < 1 {
			w.WriteHeader(http.StatusNotFound)
			log.Println("Route not found")
			fmt.Fprintf(w, "Route not found")
			return
		}
	}

	if segs[1] == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Empty webhook data")
		fmt.Fprintf(w, "Empty webhook data")
		return
	}

	command := strings.Join(segs[2:], " ")
	log.Println(command)

	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, commands)
}

func (x Webhook) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	x.CommandMsgs = commandMsgs
	webhook = x

	if connector.Debug {
		log.Println("Starting Webhook connector...")
	}

	port, _ := strconv.Atoi(connector.Port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: nil,
	}

	log.Println(server.Addr)

	http.HandleFunc("/webhook/", webhookHandler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (x Webhook) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if message.In.Process {
		for _, c := range connector.Commands {
			if strings.HasPrefix(strings.ToLower(message.In.Text), strings.ToLower(c.Match)) {
				environment := Environment{
					Address:  connector.Server,
					Password: connector.Pass,
					DB:       0,
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
	help += fmt.Sprintf("Webhooks enable at %s:%s/webhook/\n", connector.Server, connector.Port)
	return help
}
