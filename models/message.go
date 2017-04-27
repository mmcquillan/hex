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
	message.Inputs["jane.id"] = xid.New().String()
	message.Inputs["jane.type"] = Type
	message.Inputs["jane.name"] = Name
	message.Inputs["jane.target"] = Target
	message.Inputs["jane.user"] = User
	message.Inputs["jane.input"] = Input
	message.Inputs["jane.timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	message.Success = true
	message.Response = make([]string, 0)
	return message
}
