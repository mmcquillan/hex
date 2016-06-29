package connectors

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Webhook struct {
	CommandMsgs chan<- models.Message
	Connector   models.Connector
}

var webhook Webhook

func (x Webhook) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	webhook = x
	webhook.CommandMsgs = commandMsgs
	webhook.Connector = connector

	port, _ := strconv.Atoi(connector.Port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: nil,
	}

	if connector.Debug {
		log.Printf("Starting Webhook connector at: %s", server.Addr)
	}

	http.HandleFunc("/", webhookHandler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (x Webhook) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Webhook) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Webhook) Help(connector models.Connector) (help string) {
	return
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if webhook.Connector.Debug {
		log.Print("Webhook Incoming URL: " + url)
	}
	if strings.HasPrefix(url, "/command/") || strings.HasPrefix(url, "/publish/") {
		rawbody, err := ioutil.ReadAll(r.Body)
		body := string(rawbody)
		if err != nil {
			log.Print(err)
		}
		defer r.Body.Close()
		if webhook.Connector.Debug {
			log.Print("Webhook Incoming Body: " + body)
		}
		bodyParsed, err := gabs.ParseJSON([]byte(body))
		if err != nil {
			log.Print(err)
		}
		for _, c := range webhook.Connector.Commands {
			if match, _ := parse.Match(c.Match, url); match {
				if webhook.Connector.Debug {
					log.Print("Webhook Match: " + c.Match)
				}
				out := c.Output
				re := regexp.MustCompile("{(.*)}")
				subs := re.FindAllString(c.Output, -1)
				for _, sub := range subs {
					if webhook.Connector.Debug {
						log.Print("Webhook Sub: " + sub)
					}
					sub_clean := strings.Replace(sub, "{", "", -1)
					sub_clean = strings.Replace(sub_clean, "}", "", -1)
					value, ok := bodyParsed.Path(sub_clean).Data().(string)
					if ok {
						if webhook.Connector.Debug {
							log.Print("Webhook Val: " + value)
						}
						out = strings.Replace(out, sub, value, -1)
					}
				}
				out = strings.Replace(out, "{}", body, -1)
				if strings.HasPrefix(url, "/command/") {
					var m models.Message
					m.Routes = webhook.Connector.Routes
					m.In.Source = webhook.Connector.ID
					m.In.Text = out
					m.In.Process = true
					webhook.CommandMsgs <- m
				} else if strings.HasPrefix(url, "/publish/") {
					var color = "NONE"
					var match = false
					if match, _ = parse.Match(c.Green, out); match {
						color = "SUCCESS"
					}
					if match, _ = parse.Match(c.Yellow, out); match {
						color = "WARN"
					}
					if match, _ = parse.Match(c.Red, out); match {
						color = "FAIL"
					}
					var m models.Message
					m.Routes = webhook.Connector.Routes
					m.In.Source = webhook.Connector.ID
					m.In.Text = r.URL.Path
					m.In.Process = false
					m.Out.Text = c.Name
					m.Out.Detail = out
					m.Out.Status = color
					webhook.CommandMsgs <- m
				}
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JaneBot"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("¯\\_(ツ)_/¯"))
	}
}
