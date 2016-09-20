package connectors

import (
	"encoding/xml"
	"fmt"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Wolfram struct {
}

func (x Wolfram) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
}

func (x Wolfram) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if match, tokens := parse.Match("wolfram*", message.In.Text); match {
		message.In.Tags += "," + connector.Tags
		message.Out.Text = callWolfram(tokens["*"], connector.Key)
		publishMsgs <- message
	}
}

func (x Wolfram) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Wolfram) Help(connector models.Connector) (help string) {
	help += "wolfram <query> - returns the wolfram alpha results from an api\n"
	return help
}

type Query struct {
	QueryResult string `xml:"success,attr"`
	PodList     []Pod  `xml:"pod"`
	Datatypes   string `xml:"datatypes,attr"`
}

type Pod struct {
	Title      string   `xml:"title,attr"`
	SubpodList []Subpod `xml:"subpod"`
}

type Subpod struct {
	Title     string `xml:"title,attr"`
	Plaintext string `xml:"plaintext"`
	Image     Image  `xml:"img"`
}

type Image struct {
	Source string `xml:"src,attr"`
}

var errorResult = "Error processing request"
var baseQuery = "http://api.wolframalpha.com/v2/query?appid="
var input = "&input="
var returnTypes = "&format=image,plaintext"

func callWolfram(msg string, api string) (results string) {
	encodedRequest := url.QueryEscape(msg)
	getRequest := baseQuery + api + input + encodedRequest + returnTypes

	resp, err := http.Get(getRequest)
	if err != nil {
		log.Print(err)
		return errorResult
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return errorResult
	}

	var query Query
	err = xml.Unmarshal(body, &query)
	if err != nil {
		log.Print(err)
		return errorResult
	}

	if query.QueryResult == "false" {
		return "Failed to return a result"
	}

	var result string
	if query.Datatypes == "Math" {
		result = ParseWolframAlphaMathQuery(query)
	} else {
		result = ParseWolframAlphaQuery(query)
	}

	return result
}

func ParseWolframAlphaMathQuery(query Query) string {
	result := ""

	for _, pod := range query.PodList {
		if len(pod.SubpodList) == 0 {
			continue
		}
		if pod.Title != "Result" {
			continue
		}
		for _, subpod := range pod.SubpodList {
			if subpod.Plaintext == "{}" {
				continue
			}
			result += fmt.Sprintf("     %s\n", subpod.Plaintext)
		}
	}

	return result
}

func ParseWolframAlphaQuery(query Query) string {
	result := ""

	for _, pod := range query.PodList {
		if len(pod.SubpodList) == 0 {
			continue
		}
		if pod.Title == "Input interpretation" {
			continue
		}
		for _, subpod := range pod.SubpodList {
			if subpod.Plaintext == "{}" {
				continue
			}
			result += fmt.Sprintf("     %s\n", subpod.Plaintext)
			// result += fmt.Sprintf("     Image: %s\n\n", subpod.Image.Source)
		}
	}

	return result
}
