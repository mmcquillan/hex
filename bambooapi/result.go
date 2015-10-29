package bambooapi

import (
	"bytes"
	"encoding/xml"
	"github.com/kennygrant/sanitize"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Key          string `xml:"buildResultKey"`
	PlanData     Plan   `xml:"plan"`
	Plan         string
	Number       string `xml:"buildNumber"`
	StartTime    string `xml:"prettyBuildStartedTime"`
	CompleteTime string `xml:"prettyBuildCompletedTime"`
	State        string `xml:"buildState"`
	Duration     string `xml:"buildDurationDescription"`
	Responsible  string `xml:"reasonSummary"`
}

type Plan struct {
	Name string `xml:"name,attr"`
	Key  string `xml:"key,attr"`
}

func GetResult(server string, user string, pass string, key string) (result Result) {
	url := "https://" + user + ":" + pass + "@" + server
	url += "/builds/rest/api/latest/result/" + key
	req, err := http.NewRequest("GET", url, bytes.NewBufferString(""))
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	xml.Unmarshal(body, &result)
	result.Plan = result.PlanData.Name
	result.Responsible = sanitize.HTML(result.Responsible)
	return result
}
