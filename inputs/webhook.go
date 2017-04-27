package inputs

import (
	"fmt"
	"github.com/projectjane/jane/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Webhook struct {
}

func (x Webhook) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)

	port, _ := strconv.Atoi(service.Config["Port"])

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: nil,
	}

	handle := func(w http.ResponseWriter, r *http.Request) {
		rawbody, err := ioutil.ReadAll(r.Body)
		body := string(rawbody)
		if err != nil {
			log.Print(err)
		}
		defer r.Body.Close()
		message := models.MakeMessage(service.Type, service.Name, r.RequestURI, r.RemoteAddr, body)
		inputMsgs <- message

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JaneBot"))
	}

	http.HandleFunc("/", handle)

	err := server.ListenAndServe()
	if err != nil {
		log.Print(err)
	}
}
