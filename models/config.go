package models

import (
	"github.com/hashicorp/go-hclog"
)

// Config struct
type Config struct {
	Version      string
	Logger       hclog.Logger
	Admins       string            `json:"admins" yaml:"admins"`
	ACL          string            `json:"acl" yaml:"acl"`
	PluginsDir   string            `json:"plugins_dir" yaml:"plugins_dir"`
	RulesDir     string            `json:"rules_dir" yaml:"rules_dir"`
	RulesGitUrl  string            `json:"rules_git_url" yaml:"rules_git_url"`
	LogFile      string            `json:"log_file" yaml:"log_file"`
	Debug        bool              `json:"debug" yaml:"debug"`
	Trace        bool              `json:"trace" yaml:"trace"`
	Quiet        bool              `json:"quiet" yaml:"quiet"`
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
