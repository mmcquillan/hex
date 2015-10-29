package bambooapi

import (
	"bytes"
	"log"
	"net/http"
)

func Queue(server string, user string, pass string, buildkey string) {
	url := "https://" + user + ":" + pass + "@" + server
	url += "/builds/rest/api/latest/queue/" + buildkey
	url += "?executeAllStages=true&os_authType=basic"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(""))
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	res.Body.Close()
}
