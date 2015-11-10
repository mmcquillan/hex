package connectors

import (
	"reflect"
)

var List = make(map[string]reflect.Type)

func init() {
	List["cli"] = reflect.TypeOf(Cli{})
	List["slack"] = reflect.TypeOf(Slack{})
	List["monitor"] = reflect.TypeOf(Monitor{})
	List["rss"] = reflect.TypeOf(Rss{})
}

func MakeConnector(connType string) interface{} {
	c := (reflect.New(List[connType]).Elem().Interface())
	return c
}
