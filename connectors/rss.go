package connectors

import (
	"github.com/SlyMarbo/rss"
	"github.com/kennygrant/sanitize"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"html"
	"log"
	"strconv"
	"strings"
	"time"
)

type Rss struct {
	Connector models.Connector
}

func (x Rss) Run(config *models.Config, connector models.Connector) {
	defer Recovery(config, connector)
	nextMarker := ""
	for {
		nextMarker = callRss(nextMarker, config, connector)
		time.Sleep(120 * time.Second)
	}
}

func (x Rss) Send(config *models.Config, connector models.Connector, message models.Message, target string) {
	return
}

func callRss(lastMarker string, config *models.Config, connector models.Connector) (nextMarker string) {
	var displayOnStart = 0
	if config.Debug {
		log.Print("Starting rss feed fetch for " + connector.ID)
	}
	feed, err := rss.Fetch(connector.Server)
	if err != nil {
		log.Print(err)
		return
	}
	var messages []models.Message
	if config.Debug {
		log.Print("Feed count for " + connector.ID + ": " + strconv.Itoa(len(feed.Items)))
	}
	for i := len(feed.Items) - 1; i >= 0; i-- {
		if config.Debug {
			log.Print("Feed " + connector.ID + " item #" + strconv.Itoa(i) + " marker " + feed.Items[i].Date.String())
		}
		if lastMarker == "" {
			lastMarker = feed.Items[displayOnStart].Date.String()
		}
		item := feed.Items[i]
		if item.Date.String() > lastMarker {
			status := "NONE"
			if connector.SuccessMatch != "" {
				if strings.Contains(item.Title, connector.SuccessMatch) {
					status = "SUCCESS"
				}
				if strings.Contains(item.Content, connector.SuccessMatch) {
					status = "SUCCESS"
				}
			}
			if connector.WarningMatch != "" {
				if strings.Contains(item.Title, connector.WarningMatch) {
					status = "WARN"
				}
				if strings.Contains(item.Title, connector.WarningMatch) {
					status = "WARN"
				}
			}
			if connector.FailureMatch != "" {
				if strings.Contains(item.Title, connector.FailureMatch) {
					status = "FAIL"
				}
				if strings.Contains(item.Title, connector.FailureMatch) {
					status = "FAIL"
				}
			}
			m := models.Message{
				Routes:      connector.Routes,
				Request:     "",
				Title:       connector.ID + " " + html.UnescapeString(sanitize.HTML(item.Title)),
				Description: html.UnescapeString(sanitize.HTML(item.Content)),
				Link:        item.Link,
				Status:      status,
			}
			messages = append(messages, m)
			if i == 0 {
				lastMarker = item.Date.String()
			}
		}
	}
	for _, m := range messages {
		commands.Parse(config, &m)
		Broadcast(config, m)
	}
	nextMarker = lastMarker
	if config.Debug {
		log.Print("Next marker for " + connector.ID + ": " + nextMarker)
	}
	return nextMarker
}
