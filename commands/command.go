package commands

import (
	"reflect"

	"github.com/hexbotio/hex/models"
)

// Input interface
type Action interface {
	Act(message *models.Message, rules *map[string]models.Rule, config models.Config)
}

// List of Inputs
var List = make(map[string]reflect.Type)

func init() {
	List["help*"] = reflect.TypeOf(Help{})
	List["ping"] = reflect.TypeOf(Ping{})
	List["rules"] = reflect.TypeOf(Rules{})
	List["version"] = reflect.TypeOf(Version{})
}

// commandHelp function
func CommandHelp(config models.Config) (command []string) {
	command = make([]string, 4)
	command[0] = "help <filter> - This help"
	command[1] = "ping - Simple ping response for the bot"
	command[2] = "rules - dump of loaded rules"
	command[3] = "version - Compiled version number/sha"
	return command
}

// Exists function
func Exists(connType string) (exists bool) {
	_, exists = List[connType]
	return exists
}

// Make
func Make(connType string) interface{} {
	if ct, ok := List[connType]; ok {
		c := (reflect.New(ct).Elem().Interface())
		return c
	} else {
		return nil
	}
}
