package services

import (
	"log"
	"reflect"

	"github.com/projectjane/jane/models"
)

var List = make(map[string]reflect.Type)

func init() {
	List["cli"] = reflect.TypeOf(Cli{})
	List["email"] = reflect.TypeOf(Email{})
	List["exec"] = reflect.TypeOf(Exec{})
	List["file"] = reflect.TypeOf(File{})
	List["imageme"] = reflect.TypeOf(ImageMe{})
	List["jira"] = reflect.TypeOf(Jira{})
	List["log"] = reflect.TypeOf(Log{})
	List["slack"] = reflect.TypeOf(Slack{})
	List["response"] = reflect.TypeOf(Response{})
	List["rss"] = reflect.TypeOf(Rss{})
	List["wolfram"] = reflect.TypeOf(Wolfram{})
	List["webhook"] = reflect.TypeOf(Webhook{})
	List["winrm"] = reflect.TypeOf(WinRM{})
	List["twilio"] = reflect.TypeOf(Twilio{})
	List["twitter"] = reflect.TypeOf(Twitter{})
}

func MakeService(connType string) interface{} {
	if ct, ok := List[connType]; ok {
		c := (reflect.New(ct).Elem().Interface())
		return c
	} else {
		log.Print("Error in configuration, connector type '" + connType + "' not supported")
		log.Fatal("Exiting due to configuration error")
		return nil
	}
}

func Recovery(connector models.Connector) {
	msg := "Panic - " + connector.ID + " " + connector.Type + " Connector"
	if r := recover(); r != nil {
		log.Print(msg, r)
	}
}
