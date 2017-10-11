package models

import (
	"github.com/hashicorp/go-plugin"
	"github.com/hexbotio/hex-plugin"
)

// Rule Struct
type Plugin struct {
	Name   string
	Path   string
	Client *plugin.Client
	Action hexplugin.Action
}
