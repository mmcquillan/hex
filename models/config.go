package models

import (
	"github.com/hashicorp/go-hclog"
)

// Config struct
type Config struct {
	Version      string
	Logger       hclog.Logger
	Admins       string            `json:"admins"`
	ACL          string            `json:"acl"`
	PluginsDir   string            `json:"plugins_dir"`
	RulesDir     string            `json:"rules_dir"`
	LogFile      string            `json:"log_file"`
	Debug        bool              `json:"debug"`
	Trace        bool              `json:"trace"`
	Quiet        bool              `json:"quiet"`
	BotName      string            `json:"bot_name"`
	CLI          bool              `json:"cli"`
	Auditing     bool              `json:"auditing"`
	AuditingFile string            `json:"auditing_file"`
	Slack        bool              `json:"slack"`
	SlackToken   string            `json:"slack_token"`
	SlackIcon    string            `json:"slack_icon"`
	SlackDebug   bool              `json:"slack_debug"`
	Scheduler    bool              `json:"scheduler"`
	Webhook      bool              `json:"webhook"`
	WebhookPort  int               `json:"webhook_port"`
	Command      string            `json:"command"`
	Vars         map[string]string `json:"vars"`
}
