package internals

import (
	"log"
	"reflect"

	"github.com/projectjane/jane/models"
)

// Input interface
type Action interface {
	Act(message *models.Message, config *models.Config)
}

// List of Inputs
var List = make(map[string]reflect.Type)

func init() {
	List["help*"] = reflect.TypeOf(Help{})
	List["passwd"] = reflect.TypeOf(Passwd{})
	List["ping"] = reflect.TypeOf(Ping{})
	List["uptime"] = reflect.TypeOf(Uptime{})
	List["version"] = reflect.TypeOf(Version{})
	List["whoami"] = reflect.TypeOf(Whoami{})
}

// InternalHelp function
func InternalHelp(config *models.Config) (internal []string) {
	internal = make([]string, 6)
	internal[0] = config.BotName + " help <filter> - This help"
	internal[1] = config.BotName + " passwd - Password generator"
	internal[2] = config.BotName + " ping - Simple ping response for the bot"
	internal[3] = config.BotName + " uptime - Number of seconds process has been running"
	internal[4] = config.BotName + " version - Compiled version number/sha"
	internal[5] = config.BotName + " whoami - Your user name"
	return internal
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
