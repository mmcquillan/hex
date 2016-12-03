package connectors

import (
	"encoding/json"
	"fmt"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

// ImageMe Struct representing the image me connector
type ImageMe struct {
}

//Listen Not implemented
func (x ImageMe) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

// Command Takes in animateme or imageme command
func (x ImageMe) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if match, tokens := parse.Match("image me*", message.In.Text); match {
		message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags)
		message.Out.Text = callImageMe(tokens["*"], connector.KeyValues["Key"], connector.Pass, false)
		publishMsgs <- message
	}
	if match, tokens := parse.Match("animate me*", message.In.Text); match {
		message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags)
		message.Out.Text = callImageMe(tokens["*"], connector.KeyValues["Key"], connector.Pass, true)
		publishMsgs <- message
	}
}

// Publish Not implemented
func (x ImageMe) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

// Help Returns help data
func (x ImageMe) Help(connector models.Connector) (help []string) {
	help = make([]string, 0)
	help = append(help, "image me <image keywords> - pulls back an image url")
	help = append(help, "animate me <image keywords> - pulls back an animated gif url")
	return help
}

type searchResult struct {
	Items []items `json:"items"`
}

type items struct {
	Link string `json:"link"`
}

var imageClient = &http.Client{}

var baseURL = "https://www.googleapis.com/customsearch/v1?key="
var errorMessage = "Error retrieving image"
var animated bool

func callImageMe(msg string, apiKey string, cx string, animated bool) string {
	start := rand.Intn(3)
	if start < 1 {
		start = 1
	}

	cx = "&cx=" + cx
	returnFields := fmt.Sprintf("&fields=items(link)&start=%v", start)
	query := "&q=" + url.QueryEscape(msg)
	fields := "&searchType=image"
	if animated {
		fields += "&fileType=gif&hq=animated&tbs=itp:animated"
	}

	url := baseURL + apiKey + cx + returnFields + query + fields

	resp, err := imageClient.Get(url)
	if err != nil {
		log.Print(err)
		return errorMessage
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return findDeprecatedImage(msg, animated)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return errorMessage
	}

	var result searchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Print(err)
		return errorMessage
	}

	if len(result.Items) > 0 {
		randomLink := result.Items[rand.Intn(len(result.Items))]
		return randomLink.Link
	}

	return findDeprecatedImage(msg, animated)
}

type deprecatedResult struct {
	ResponseData responseData `json:"responseData"`
}

type responseData struct {
	Results []result `json:"results"`
}

type result struct {
	URL string `json:"url"`
}

func findDeprecatedImage(query string, animated bool) string {
	baseURL := "https://ajax.googleapis.com/ajax/services/search/images?v=1.0&rsz=8"
	if animated {
		baseURL += "&as_filetype=gif"
	}

	baseURL += "&q="
	searchURL := baseURL + url.QueryEscape(query)

	if animated {
		searchURL += url.QueryEscape(" animated")
	}

	resp, err := imageClient.Get(searchURL)
	if err != nil {
		log.Print(err)
		return errorMessage
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return errorMessage
	}

	var result deprecatedResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Print(err)
		return errorMessage
	}

	index := rand.Intn(len(result.ResponseData.Results))

	if len(result.ResponseData.Results) > 0 {
		return result.ResponseData.Results[index].URL
	}

	return "No results found"
}
