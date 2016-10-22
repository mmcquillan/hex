package connectors

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
)

var something = ""
var something2 = ""

// Twitter Empty struct
type Twitter struct {
	Client          *twitter.Client
	CommandMessages chan<- models.Message
}

// Listen Listen to the Twitter stream api
func (x Twitter) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	if x.Client == nil {
		client := setupTwitterClient(connector)
		x.Client = client
		x.CommandMessages = commandMsgs

		x.listenToStream(connector)
	}

	return
}

// Command Twitter command to post a tweet from the app
func (x Twitter) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	for _, c := range connector.Commands {
		if match, _ := parse.Match(c.Match, message.In.Text); match {

			log.Println(match)

			// message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags+","+c.Tags)
			// message.Out.Text = fmt.Sprintf("Redis Server: %s\nStatus:%s", connector.Server, status.String())
			// publishMsgs <- message
			return
		}
	}

	return
}

// Publish Twitter publishes out tweets
func (x Twitter) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

// Help Twitter help information
func (x Twitter) Help(connector models.Connector) (help []string) {
	return help
}

func setupTwitterClient(connector models.Connector) *twitter.Client {
	config := oauth1.NewConfig(connector.Key, connector.Secret)
	token := oauth1.NewToken(connector.AccessToken, connector.AccessTokenSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	return client
}

func (x Twitter) listenToStream(connector models.Connector) {
	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		if connector.Debug {
			log.Println(tweet.Text)
		}

		var m models.Message
		m.In.ConnectorType = webhook.Connector.Type
		m.In.ConnectorID = webhook.Connector.ID
		m.In.Tags = parse.TagAppend("", connector.Tags)
		m.In.Text = tweet.Text
		m.In.Process = false

		x.CommandMessages <- m
	}

	demux.DM = func(dm *twitter.DirectMessage) {
		if connector.Debug {
			log.Println(dm.SenderID)
		}
	}

	demux.Event = func(event *twitter.Event) {
		if connector.Debug {
			log.Printf("%#v\n", event)
		}
	}

	log.Println("Starting Twitter Stream...")

	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"DevOps", "Docker", "Azure"},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := x.Client.Streams.Filter(filterParams)
	if err != nil {
		log.Println(err)
		return
	}

	go demux.HandleChan(stream.Messages)
}

func (x Twitter) postTweet() {

}
