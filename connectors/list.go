package connectors

import (
	"log"
	"reflect"
)

var List = make(map[string]reflect.Type)

func init() {
	List["cli"] = reflect.TypeOf(Cli{})
	List["email"] = reflect.TypeOf(Email{})
	List["monitor"] = reflect.TypeOf(Monitor{})
	List["slack"] = reflect.TypeOf(Slack{})
	List["rss"] = reflect.TypeOf(Rss{})
	List["website"] = reflect.TypeOf(Website{})
}

func MakeConnector(connType string) interface{} {
	if ct, ok := List[connType]; ok {
		c := (reflect.New(ct).Elem().Interface())
		return c
	} else {
		log.Print("Error in configuration, connector type '" + connType + "' not supported")
		log.Fatal("Exiting due to configuration error")
		return nil
	}
}
