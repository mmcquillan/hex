package connectors

import (
	"github.com/SlyMarbo/rss"
	"github.com/kennygrant/sanitize"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"html"
	"log"
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

func callRss(lastMarker string, config *models.Config, connector models.Connector) (nextMarker string) {
	var displayOnStart = 0
	if config.Debug {
		log.Print("Starting rss feed fetch for " + connector.Server)
	}
	feed, err := rss.Fetch(connector.Server)
	if err != nil {
		log.Print(err)
		return
	}
	var messages []models.Message
	if config.Debug {
		log.Print("Feed count for " + connector.Server + ": " + string(len(feed.Items)))
	}
	for i := len(feed.Items) - 1; i >= 0; i-- {
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
			for _, r := range connector.Routes {
				if strings.Contains(item.Title, r.Match) || r.Match == "*" {
					m := models.Message{
						Relays:      r.Relays,
						Target:      r.Target,
						Request:     "",
						Title:       connector.Name + " " + html.UnescapeString(sanitize.HTML(item.Title)),
						Description: html.UnescapeString(sanitize.HTML(item.Content)),
						Link:        item.Link,
						Status:      status,
					}
					messages = append(messages, m)
				}
			}
			if i == 0 {
				lastMarker = item.Date.String()
			}
		}
	}
	for _, m := range messages {
		commands.Parse(config, m)
	}
	nextMarker = lastMarker
	if config.Debug {
		log.Print("Next marker for " + connector.Server + ": " + nextMarker)
	}
	return nextMarker
}
