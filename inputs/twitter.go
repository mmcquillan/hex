package inputs

import (
	"log"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/projectjane/jane/models"
)

// Twitter Struct representing the
type Twitter struct {
}

// Input Input to the Twitter stream api
func (x Twitter) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)

	config := oauth1.NewConfig(service.Config["Key"], service.Config["Secret"])
	token := oauth1.NewToken(service.Config["AccessToken"], service.Config["AccessTokenSecret"])

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		message := models.MakeMessage(service.Type, service.Name, "", tweet.User.ScreenName, tweet.Text)
		message.Inputs["jane.twitter.lang"] = tweet.Lang
		inputMsgs <- message
	}

	demux.DM = func(dm *twitter.DirectMessage) {
		// do nothing
	}

	demux.Event = func(event *twitter.Event) {
		// do nothing
	}

	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         strings.Split(service.Config["Filter"], ","),
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Println(err)
		return
	}

	go demux.HandleChan(stream.Messages)
}
