package services

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

// Wolfram Empty struct
type Wolfram struct {
}

// Input Not Implemented
func (x Wolfram) Input(inputMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
}

// Action Matches wolfram command and queries the wolfram api
func (x Wolfram) Action(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	if match, tokens := parse.Match("wolfram*", message.In.Text); match {
		message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags)
		message.Out.Text = callWolfram(tokens["*"], connector.Key)
		outputMsgs <- message
	}
}

// Output Not Implemented
func (x Wolfram) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	return
}

// Help Returns help information
func (x Wolfram) Help(connector models.Connector) (help []string) {
	help = make([]string, 0)
	help = append(help, "wolfram <query> - returns the wolfram alpha results from an api")
	return help
}

type query struct {
	QueryResult string `xml:"success,attr"`
	PodList     []pod  `xml:"pod"`
	Datatypes   string `xml:"datatypes,attr"`
}

type pod struct {
	Title      string   `xml:"title,attr"`
	SubpodList []subpod `xml:"subpod"`
}

type subpod struct {
	Title     string `xml:"title,attr"`
	Plaintext string `xml:"plaintext"`
	Image     image  `xml:"img"`
}

type image struct {
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

	var wolfQuery query
	err = xml.Unmarshal(body, &wolfQuery)
	if err != nil {
		log.Print(err)
		return errorResult
	}

	if wolfQuery.QueryResult == "false" {
		return "Failed to return a result"
	}

	var result string
	if wolfQuery.Datatypes == "Math" {
		result = parseWolframAlphaMathQuery(wolfQuery)
	} else {
		result = parseWolframAlphaQuery(wolfQuery)
	}

	return result
}

func parseWolframAlphaMathQuery(wolfQuery query) string {
	result := ""

	for _, pod := range wolfQuery.PodList {
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

func parseWolframAlphaQuery(wolfQuery query) string {
	result := ""

	for _, pod := range wolfQuery.PodList {
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
		}
	}

	return result
}
