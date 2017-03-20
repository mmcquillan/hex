package services

import (
	"github.com/SlyMarbo/rss"
	"github.com/kennygrant/sanitize"
	"github.com/projectjane/jane/models"
	"html"
	"log"
	"strconv"
	"time"
)

type Rss struct {
	Connector models.Connector
}

func (x Rss) Input(inputMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	nextMarker := ""
	for {
		nextMarker = callRss(nextMarker, inputMsgs, connector)
		time.Sleep(120 * time.Second)
	}
}

func (x Rss) Command(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Rss) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Rss) Help(connector models.Connector) (help []string) {
	return
}

func callRss(lastMarker string, inputMsgs chan<- models.Message, connector models.Connector) (nextMarker string) {
	var displayOnStart = 0
	if connector.Debug {
		log.Print("Starting rss feed fetch for " + connector.ID)
	}
	feed, err := rss.Fetch(connector.Server)
	if err != nil {
		log.Print(err)
		return
	}
	if connector.Debug {
		log.Print("Feed count for " + connector.ID + ": " + strconv.Itoa(len(feed.Items)))
	}
	for i := len(feed.Items) - 1; i >= 0; i-- {
		if connector.Debug {
			log.Print("Feed " + connector.ID + " item #" + strconv.Itoa(i) + " marker " + feed.Items[i].Date.String())
		}
		if lastMarker == "" {
			lastMarker = feed.Items[displayOnStart].Date.String()
		}
		item := feed.Items[i]
		if item.Date.String() > lastMarker {
			var m models.Message
			m.In.ConnectorType = connector.Type
			m.In.ConnectorID = connector.ID
			m.In.Tags = connector.Tags
			m.In.Process = false
			m.Out.Text = connector.ID + " " + html.UnescapeString(sanitize.HTML(item.Title))
			m.Out.Detail = html.UnescapeString(sanitize.HTML(item.Content))
			m.Out.Link = item.Link
			inputMsgs <- m
			if i == 0 {
				lastMarker = item.Date.String()
			}
		}
	}
	nextMarker = lastMarker
	if connector.Debug {
		log.Print("Next marker for " + connector.ID + ": " + nextMarker)
	}
	return nextMarker
}
