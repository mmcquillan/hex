package listeners

import (
	"github.com/SlyMarbo/rss"
	"github.com/mmcquillan/bambooapi"
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/outputs"
	"log"
	"strings"
)

func Bamboo(config *configs.Config, lastMarker string) (nextMarker string, messages []outputs.Message) {
	channels := config.BambooChannels
	url := "https://" + config.BambooUser + ":" + config.BambooPass + "@" + config.BambooUrl + "/builds/plugins/servlet/streams?local=true"
	feed, err := rss.Fetch(url)
	if err != nil {
		log.Println(err)
	}
	for i := len(feed.Items) - 1; i >= 0; i-- {
		item := feed.Items[i]
		if item.Date.String() > lastMarker {
			build := strings.Split(strings.Replace(item.Title, "<a href=\"https://"+config.BambooUrl+"/builds/browse/", "", 1), "\"")[0]
			link := "https://" + config.BambooUrl + "/builds/browse/" + build
			res := bambooapi.GetResult(config.BambooUrl, config.BambooUser, config.BambooPass, build)
			for planmatch, channel := range channels {
				if strings.Contains(res.Plan, planmatch) || planmatch == "*" {
					m := outputs.Message{channel, "Bamboo Build " + res.State, res.Plan + " #" + res.Number + " - " + res.Responsible, link, res.State}
					messages = append(messages, m)
				}
			}
			nextMarker = item.Date.String()
		}
	}
	return nextMarker, messages
}
