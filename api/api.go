package api

import (
  "net/http"
  "log"
  // "encoding/json"
  "github.com/projectjane/jane/api/routers"
  "github.com/projectjane/jane/api/common"
  "github.com/projectjane/jane/models"
)

func StartRestServer(commandMsgs, publishMsgs chan<- models.Message) {
  log.Println("*** Starting Rest Server ***")

  common.StartUp(commandMsgs, publishMsgs)

  router := routers.InitRoutes()

  server := &http.Server{
    Addr: ":8080",
    Handler: router,
  }

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal(err)
    // data.CloseDb()
  }
}
