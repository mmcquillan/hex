package models

import (
	"time"

	"github.com/rs/xid"
)

var PASS = "pass"
var WARN = "warn"
var FAIL = "fail"

// Message struct
type Message struct {
	Debug      bool
	CreateTime int64
	Attributes map[string]string
	Outputs    []Output
}

type Output struct {
	Rule      string
	StartTime int64
	EndTime   int64
	State     string
	Response  string
}

func MessageID() string {
	return xid.New().String()
}

func MessageTimestamp() int64 {
	return time.Now().Unix()
}

// MakeMessage function
func NewMessage() (message Message) {
	message.CreateTime = MessageTimestamp()
	message.Attributes = make(map[string]string)
	message.Attributes["hex.id"] = MessageID()
	message.Debug = false
	message.Outputs = make([]Output, 0)
	return message
}
