package common

import (
  "github.com/projectjane/jane/models"
)

var CommandMsgs chan<- models.Message
var PublishMsgs chan<- models.Message

func StartUp(commandMsgs, publishMsgs chan<- models.Message) {
  CommandMsgs = commandMsgs
  PublishMsgs = publishMsgs
}
