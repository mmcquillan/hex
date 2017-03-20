package services

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
)

// Twitter Struct representing the
type Twitter struct {
	StreamClient    *twitter.Client
	TweetClient     *twitter.Client
	CommandMessages chan<- models.Message
	Connector       models.Connector
}

// Input Input to the Twitter stream api
func (x Twitter) Input(inputMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	if x.StreamClient == nil {
		client := newTwitterClient(connector)
		x.StreamClient = client
		x.CommandMessages = inputMsgs

		x.listenToStream(connector)
	}

	return
}

// Command Twitter command to post a tweet from the app
func (x Twitter) Command(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	for _, c := range connector.Commands {
		if match, tokens := parse.Match(c.Match, message.In.Text); match {

			message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags+","+c.Tags)

			tweet := parse.Substitute(c.Output, tokens)

			if x.TweetClient == nil {
				client := newTwitterClient(connector)
				x.TweetClient = client
			}

			err := x.postTweet(tweet)
			if err != nil {
				message.Out.Text = "Failed to post tweet."
			} else {
				message.Out.Text = "Successfully posted tweet."
			}

			outputMsgs <- message

			return
		}
	}

	return
}

// Output Twitter publishes out tweets
func (x Twitter) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	for {
		message := <-outputMsgs

		msg := message.Out.Text

		if x.TweetClient == nil {
			client := newTwitterClient(connector)
			x.TweetClient = client
		}

		x.postTweet(msg)
	}
}

// Help Twitter help information
func (x Twitter) Help(connector models.Connector) (help []string) {
	return help
}

func newTwitterClient(connector models.Connector) *twitter.Client {
	config := oauth1.NewConfig(connector.Key, connector.Secret)
	token := oauth1.NewToken(connector.KeyValues["AccessToken"], connector.KeyValues["AccessTokenSecret"])

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

		if tweet.Lang == "en" {
			var m models.Message
			m.In.ConnectorType = webhook.Connector.Type
			m.In.ConnectorID = webhook.Connector.ID
			m.In.Tags = parse.TagAppend("", connector.Tags)
			m.In.Text = tweet.Text
			m.Out.Text = tweet.User.ScreenName
			m.Out.Detail = fmt.Sprintf("```%s ```", tweet.Text)
			m.In.Process = false

			x.CommandMessages <- m
		}
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
		Track:         connector.Filter,
		StallWarnings: twitter.Bool(true),
	}

	stream, err := x.StreamClient.Streams.Filter(filterParams)
	if err != nil {
		log.Println(err)
		return
	}

	go demux.HandleChan(stream.Messages)
}

func (x Twitter) postTweet(message string) error {
	if x.TweetClient != nil {
		_, _, err := x.TweetClient.Statuses.Update(message, nil)
		if err != nil {
			log.Println("Error posting Twitter status:", err)
			return err
		}
	} else {
		log.Println("TweetClient null")
	}

	return nil
}
