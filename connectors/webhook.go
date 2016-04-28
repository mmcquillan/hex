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
	Connector   models.Connector
}

var webhook Webhook

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var segs []string

	// get everything passed /webhook/
	webhookString := r.URL.Path[9:]

	log.Println("About to split on plus")
	segs = strings.Split(webhookString, "+")

	log.Println(segs)
	if len(segs) < 2 || segs[1] == "" {

		log.Println("About to split on slash")
		segs = strings.Split(webhookString, "/")
		log.Println(segs)

		if len(segs) < 1 {
			w.WriteHeader(http.StatusNotFound)
			log.Println("Route not found")
			fmt.Fprintf(w, "Route not found")
			return
		}
	}

	if segs[0] == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Empty webhook data")
		fmt.Fprintf(w, "Empty webhook data")
		return
	}

	command := strings.Join(segs[0:], " ")
	log.Println(command)

	var m models.Message
	m.In.Source = webhook.Connector.ID
	m.In.Text = command
	m.In.Process = true
	webhook.CommandMsgs <- m

	w.WriteHeader(http.StatusOK)
}

func (x Webhook) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	x.CommandMsgs = commandMsgs
	x.Connector = connector
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
	log.Println("Processing command...")
	log.Println(message.In.Text)
	if message.In.Process {
		for _, c := range connector.Commands {
			if strings.HasPrefix(strings.ToLower(message.In.Text), strings.ToLower(c.Match)) {
				msg := strings.TrimSpace(strings.Replace(message.In.Text, c.Match, "", 1))
				log.Printf("Publishing... %s", msg)
				message.Out.Text = msg
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
