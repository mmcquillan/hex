package listeners

import (
	"github.com/SlyMarbo/rss"
	"github.com/mmcquillan/jane/bambooapi"
	"github.com/mmcquillan/jane/configs"
	"log"
	"strings"
)

func Bamboo(config *configs.Config, lastMarker string) (nextMarker string, messages []Message) {
	channels := config.BambooChannels
	url := "https://" + config.BambooUser + ":" + config.BambooPass + "@prysminc.atlassian.net/builds/plugins/servlet/streams?local=true"
	feed, err := rss.Fetch(url)
	if err != nil {
		log.Println(err)
	}
	for i := len(feed.Items) - 1; i >= 0; i-- {
		item := feed.Items[i]
		if item.Date.String() > lastMarker {
			build := strings.Split(strings.Replace(item.Title, "<a href=\"https://prysminc.atlassian.net/builds/browse/", "", 1), "\"")[0]
			link := "https://prysminc.atlassian.net/builds/browse/" + build
			res := bambooapi.GetResult("prysminc.atlassian.net", config.BambooUser, config.BambooPass, build)
			for planmatch, channel := range channels {
				if strings.Contains(res.Plan, planmatch) || planmatch == "*" {
					m := Message{channel, "Bamboo Build " + res.State, res.Plan + " #" + res.Number + " - " + res.Responsible, link, res.State}
					messages = append(messages, m)
				}
			}
			nextMarker = item.Date.String()
		}
	}
	return nextMarker, messages
}
