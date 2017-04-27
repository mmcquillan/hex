package outputs

import (
	"log"
	"reflect"

	"github.com/hexbotio/hex/models"
)

// Input interface
type Output interface {
	Write(outputMsgs <-chan models.Message, service models.Service)
}

// List of Inputs
var List = make(map[string]reflect.Type)

func init() {
	List["cli"] = reflect.TypeOf(Cli{})
	List["file"] = reflect.TypeOf(File{})
	List["slack"] = reflect.TypeOf(Slack{})
	List["twilio"] = reflect.TypeOf(Twilio{})
}

// Exists function
func Exists(connType string) (exists bool) {
	_, exists = List[connType]
	return exists
}

// MakeService
func Make(connType string) interface{} {
	if ct, ok := List[connType]; ok {
		c := (reflect.New(ct).Elem().Interface())
		return c
	} else {
		return nil
	}
}

// Recovery
func Recovery(service models.Service) {
	msg := "Panic - " + service.Name + " " + service.Type
	if r := recover(); r != nil {
		log.Print(msg, r)
	}
}
