package models

import (
	"strconv"
	"time"

	"github.com/rs/xid"
)

// Message struct
type Message struct {
	Inputs   map[string]string
	Success  bool
	Response []string
	Outputs  []Output
}

// MakeMessage function
func MakeMessage(Type string, Name string, Target string, User string, Input string) (message Message) {
	message.Inputs = make(map[string]string)
	message.Inputs["hex.id"] = xid.New().String()
	message.Inputs["hex.type"] = Type
	message.Inputs["hex.name"] = Name
	message.Inputs["hex.target"] = Target
	message.Inputs["hex.user"] = User
	message.Inputs["hex.input"] = Input
	message.Inputs["hex.timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	message.Success = true
	message.Response = make([]string, 0)
	return message
}
