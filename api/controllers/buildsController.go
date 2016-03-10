package controllers

import (
  "net/http"
  "encoding/json"
  // "github.com/gorilla/mux"

  "github.com/projectjane/jane/api/common"
  "github.com/projectjane/jane/models"

)

func StartBuild(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var buildPlan models.BuildPlan

  err := json.NewDecoder(r.Body).Decode(&buildPlan)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
  } else {

    var msg models.Message
    msg.In.Text = "bamboo build " + buildPlan.PlanKey

    common.PublishMsgs<- msg

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Started build"))
  }
}


// func (x Bamboo) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
// 	if strings.HasPrefix(message.In.Text, "bamboo status") {
// 		commandDeployStatus(message, publishMsgs, connector)
// 		commandBuildStatus(message, publishMsgs, connector)
// 	}
// 	if strings.HasPrefix(message.In.Text, "bamboo build") {
// 		commandBuild(message, publishMsgs, connector)
// 	}
// }
