package services

import (
	"log"
	"reflect"
)

var List = make(map[string]reflect.Type)

func init() {
	List["cli"] = reflect.TypeOf(Cli{})
	List["client"] = reflect.TypeOf(Client{})
	List["bamboo"] = reflect.TypeOf(Bamboo{})
	List["email"] = reflect.TypeOf(Email{})
	List["exec"] = reflect.TypeOf(Exec{})
	List["file"] = reflect.TypeOf(File{})
	List["imageme"] = reflect.TypeOf(ImageMe{})
	List["jira"] = reflect.TypeOf(Jira{})
	List["log"] = reflect.TypeOf(Log{})
	List["slack"] = reflect.TypeOf(Slack{})
	List["response"] = reflect.TypeOf(Response{})
	List["rss"] = reflect.TypeOf(Rss{})
	List["website"] = reflect.TypeOf(Website{})
	List["wolfram"] = reflect.TypeOf(Wolfram{})
	List["redis"] = reflect.TypeOf(Redis{})
	List["server"] = reflect.TypeOf(Server{})
	List["webhook"] = reflect.TypeOf(Webhook{})
	List["winrm"] = reflect.TypeOf(WinRM{})
	List["twilio"] = reflect.TypeOf(Twilio{})
	List["twitter"] = reflect.TypeOf(Twitter{})

	// depricated
	List["exec2"] = reflect.TypeOf(Exec{})
	List["logging"] = reflect.TypeOf(File{})
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
