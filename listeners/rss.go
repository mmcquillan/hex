package listeners

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

func Rss(config *models.Config, listener models.Listener) {
	var displayOnStart = 0
	lastMarker := ""
	for {
		feed, err := rss.Fetch(listener.Resource)
		if err != nil {
			log.Println(err)
			return
		}
		var messages []models.Message
		for i := len(feed.Items) - 1; i >= 0; i-- {
			if lastMarker == "" {
				lastMarker = feed.Items[displayOnStart].Date.String()
			}
			item := feed.Items[i]
			if item.Date.String() > lastMarker {
				status := "NONE"
				if listener.SuccessMatch != "" {
					if strings.Contains(item.Title, listener.SuccessMatch) {
						status = "SUCCESS"
					}
					if strings.Contains(item.Content, listener.SuccessMatch) {
						status = "SUCCESS"
					}
				}
				if listener.FailureMatch != "" {
					if strings.Contains(item.Title, listener.FailureMatch) {
						status = "FAIL"
					}
					if strings.Contains(item.Title, listener.FailureMatch) {
						status = "FAIL"
					}
				}
				for _, d := range listener.Destinations {
					if strings.Contains(item.Title, d.Match) || d.Match == "*" {
						m := models.Message{
							Relays:      d.Relays,
							Target:      d.Target,
							Request:     "",
							Title:       listener.Name + " " + html.UnescapeString(sanitize.HTML(item.Title)),
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
		time.Sleep(120 * time.Second)
	}
}
