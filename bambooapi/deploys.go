package bambooapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Deploys []struct {
	Environmentstatuses []struct {
		Environment struct {
			ID                  int    `json:"id"`
			Name                string `json:"name"`
			Deploymentprojectid int    `json:"deploymentProjectId"`
		} `json:"environment"`
		Deploymentresult struct {
			Deploymentversion struct {
				ID                 int    `json:"id"`
				Name               string `json:"name"`
				Creationdate       int64  `json:"creationDate"`
				Creatorusername    string `json:"creatorUserName"`
				Creatordisplayname string `json:"creatorDisplayName"`
				Planbranchname     string `json:"planBranchName"`
				Agezeropoint       int64  `json:"ageZeroPoint"`
			} `json:"deploymentVersion"`
			Deploymentversionname string `json:"deploymentVersionName"`
			ID                    int    `json:"id"`
			Deploymentstate       string `json:"deploymentState"`
			Lifecyclestate        string `json:"lifeCycleState"`
			Starteddate           int64  `json:"startedDate"`
			Queueddate            int64  `json:"queuedDate"`
			Executeddate          int64  `json:"executedDate"`
			Finisheddate          int64  `json:"finishedDate"`
			Reasonsummary         string `json:"reasonSummary"`
		} `json:"deploymentResult"`
	} `json:"environmentStatuses"`
}

func DeployResults(server string, user string, pass string) (deploys Deploys) {
	url := "https://" + user + ":" + pass + "@" + server
	//url += "/builds/rest/api/latest/deploy/dashboard/" + projectid
	url += "/builds/rest/api/latest/deploy/dashboard/"
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
	json.Unmarshal(body, &deploys)
	return deploys
}
