package models

import (
	"github.com/hashicorp/go-hclog"
)

// Config struct
type Config struct {
	Version      string
	Logger       hclog.Logger
	Admins       string            `json:"admins" yaml:"admins"`
	UserACL      string            `json:"user_acl" yaml:"user_acl"`
	ChannelACL   string            `json:"channel_acl" yaml:"channel_acl"`
	PluginsDir   string            `json:"plugins_dir" yaml:"plugins_dir"`
	RulesDir     string            `json:"rules_dir" yaml:"rules_dir"`
	LogFile      string            `json:"log_file" yaml:"log_file"`
	LogLevel     string            `json:"log_level" yaml:"log_level"`
	BotName      string            `json:"bot_name" yaml:"bot_name"`
	CLI          bool              `json:"cli" yaml:"cli"`
	Auditing     bool              `json:"auditing" yaml:"auditing"`
	AuditingFile string            `json:"auditing_file" yaml:"auditing_file"`
	Slack        bool              `json:"slack" yaml:"slack"`
	SlackToken   string            `json:"slack_token" yaml:"slack_token"`
	SlackIcon    string            `json:"slack_icon" yaml:"slack_icon"`
	SlackDebug   bool              `json:"slack_debug" yaml:"slack_debug"`
	Scheduler    bool              `json:"scheduler" yaml:"scheduler"`
	Webhook      bool              `json:"webhook" yaml:"webhook"`
	WebhookPort  int               `json:"webhook_port" yaml:"webhook_port"`
	Command      string            `json:"command" yaml:"command"`
	Vars         map[string]string `json:"vars" yaml:"vars"`
}
