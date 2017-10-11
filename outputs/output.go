package outputs

import (
	"reflect"

	"github.com/hexbotio/hex/models"
)

// Input interface
type Output interface {
	Write(outputMsgs <-chan models.Message, config models.Config)
}

// List of Inputs
var List = make(map[string]reflect.Type)

func init() {
	List["cli"] = reflect.TypeOf(Cli{})
	List["slack"] = reflect.TypeOf(Slack{})
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
