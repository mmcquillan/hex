package inputs

import (
	"github.com/SlyMarbo/rss"
	"github.com/kennygrant/sanitize"
	"github.com/hexbotio/hex/models"
	"html"
	"log"
	"time"
)

type Rss struct {
}

func (x Rss) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)
	nextMarker := ""
	for {
		nextMarker = callRss(nextMarker, inputMsgs, service)
		time.Sleep(120 * time.Second)
	}
}

func callRss(lastMarker string, inputMsgs chan<- models.Message, service models.Service) (nextMarker string) {
	var displayOnStart = 0
	feed, err := rss.Fetch(service.Config["URL"])
	if err != nil {
		log.Print(err)
		return
	}
	for i := len(feed.Items) - 1; i >= 0; i-- {
		if lastMarker == "" {
			lastMarker = feed.Items[displayOnStart].Date.String()
		}
		item := feed.Items[i]
		if item.Date.String() > lastMarker {
			title := html.UnescapeString(sanitize.HTML(item.Title))
			message := models.MakeMessage(service.Type, service.Name, "", "", title)
			message.Inputs["hex.rss.title"] = title
			message.Inputs["hex.rss.content"] = html.UnescapeString(sanitize.HTML(item.Content))
			message.Inputs["hex.rss.date"] = item.Date.String()
			message.Inputs["hex.rss.link"] = item.Link
			inputMsgs <- message
			if i == 0 {
				lastMarker = item.Date.String()
			}
		}
	}
	nextMarker = lastMarker
	return nextMarker
}
