package connectors

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
		log.Print(err)
	}
}

func (x Webhook) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Webhook) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Webhook) Help(connector models.Connector) (help []string) {
	return
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	reqUrl := r.URL.Path
	reqQs, err := url.QueryUnescape(r.URL.RawQuery)
	if err != nil {
		log.Print(err)
	}
	if webhook.Connector.Debug {
		log.Print("Webhook Incoming URL: " + reqUrl)
	}
	rawbody, err := ioutil.ReadAll(r.Body)
	body := string(rawbody)
	if err != nil {
		log.Print(err)
	}
	defer r.Body.Close()
	if webhook.Connector.Debug {
		log.Print("Webhook Incoming Body: " + body)
	}

	// for new relic (lame)
	if strings.HasPrefix(body, "deployment=%7B") {
		body, err = url.QueryUnescape(strings.Replace(body, "deployment=", "", 1))
		if err != nil {
			log.Print(err)
		}
	}
	if strings.HasPrefix(body, "alert=%7B") {
		body, err = url.QueryUnescape(strings.Replace(body, "alert=", "", 1))
		if err != nil {
			log.Print(err)
		}
	}

	bodyParsed, err := gabs.ParseJSON([]byte(body))
	isJson := true
	if err != nil {
		log.Print(err)
		isJson = false
	}
	for _, c := range webhook.Connector.Commands {
		if match, _ := parse.Match(c.Match, reqUrl); match {
			if webhook.Connector.Debug {
				log.Print("Webhook Match: " + c.Match)
			}
			tokens := make(map[string]string)
			tokens["?"] = reqQs
			tokens["*"] = body
			if isJson {
				if match, subs := parse.SubstitutionVars(c.Output); match {
					for _, sub := range subs {
						if webhook.Connector.Debug {
							log.Print("Webhook Sub: " + sub)
						}
						value, ok := bodyParsed.Path(parse.Strip(sub)).Data().(string)
						if ok {
							if webhook.Connector.Debug {
								log.Print("Webhook Val: " + value)
							}
							tokens[parse.Strip(sub)] = value
						}
					}

				}
			}
			out := parse.Substitute(c.Output, tokens)
			if c.Process {
				var m models.Message
				m.In.ConnectorType = webhook.Connector.Type
				m.In.ConnectorID = webhook.Connector.ID
				m.In.Tags = webhook.Connector.Tags
				m.In.Text = out
				m.In.Process = true
				webhook.CommandMsgs <- m
			} else {
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
				m.In.ConnectorType = webhook.Connector.Type
				m.In.ConnectorID = webhook.Connector.ID
				m.In.Tags = webhook.Connector.Tags
				m.In.Text = reqUrl
				m.In.Process = false
				if c.Name != "" {
					m.Out.Text = c.Name
					m.Out.Detail = out
				} else {
					m.Out.Text = out
				}
				m.Out.Status = color
				webhook.CommandMsgs <- m
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("JaneBot"))
}
