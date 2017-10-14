package core

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hexbotio/hex/models"
)

func Config(config *models.Config, version string) {

	// start with defaults
	config.Version = version
	config.PluginsDir = "/etc/hex/plugins"
	config.RulesDir = "/etc/hex/rules"
	config.LogFile = ""
	config.WorkspaceDir = "/tmp"
	config.Debug = false
	config.BotName = "@hex"
	config.CLI = false
	config.Auditing = false
	config.AuditingFile = ""
	config.Slack = false
	config.SlackToken = ""
	config.SlackIcon = ":nut_and_bolt:"
	config.SlackDebug = false
	config.Scheduler = false
	config.Webhook = false
	config.WebhookPort = 8000

	// version and exit
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Print("HexBot " + config.Version + "\n")
		os.Exit(0)
	}

	// evaluate for config file
	if len(os.Args) > 1 && FileExists(os.Args[1]) {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal("Config File Read - ", err)
		}
		err = json.Unmarshal(file, &config)
		if err != nil {
			log.Fatal("Config File Unmarshal - ", err)
		}
	}

	// environment
	if os.Getenv("HEX_RULES_DIR") != "" {
		config.RulesDir = os.Getenv("HEX_RULES_DIR")
	}
	if os.Getenv("HEX_PLUGINS_DIR") != "" {
		config.PluginsDir = os.Getenv("HEX_PLUGINS_DIR")
	}
	if os.Getenv("HEX_LOG_FILE") != "" {
		config.LogFile = os.Getenv("HEX_LOG_FILE")
	}
	if os.Getenv("HEX_WORKSPACE_DIR") != "" {
		config.WorkspaceDir = os.Getenv("HEX_WORKSPACE_DIR")
	}
	if strings.ToUpper(os.Getenv("HEX_DEBUG")) == "TRUE" {
		config.Debug = true
	}
	if os.Getenv("HEX_BOT_NAME") != "" {
		config.BotName = os.Getenv("HEX_BOT_NAME")
	}
	if strings.ToUpper(os.Getenv("HEX_CLI")) == "TRUE" {
		config.CLI = true
	}
	if strings.ToUpper(os.Getenv("HEX_AUDITING")) == "TRUE" {
		config.Auditing = true
	}
	if os.Getenv("HEX_AUDITING_FILE") != "" {
		config.AuditingFile = os.Getenv("HEX_AUDITING_FILE")
	}
	if strings.ToUpper(os.Getenv("HEX_SLACK")) == "TRUE" {
		config.Slack = true
	}
	if os.Getenv("HEX_SLACK_TOKEN") != "" {
		config.SlackToken = os.Getenv("HEX_SLACK_TOKEN")
	}
	if os.Getenv("HEX_SLACK_ICON") != "" {
		config.SlackIcon = os.Getenv("HEX_SLACK_ICON")
	}
	if strings.ToUpper(os.Getenv("HEX_SLACK_DEBUG")) == "TRUE" {
		config.SlackDebug = true
	}
	if strings.ToUpper(os.Getenv("HEX_SCHEDULER")) == "TRUE" {
		config.Scheduler = true
	}
	if strings.ToUpper(os.Getenv("HEX_WEBHOOK")) == "TRUE" {
		config.Webhook = true
	}
	if os.Getenv("HEX_WEBHOOK_PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("HEX_WEBHOOK_PORT"))
		if err != nil {
			log.Fatal("Webhook Port is not a Number")
		}
		config.WebhookPort = port
	}

	// flags
	RulesDir := flag.String("rules-dir", config.RulesDir, "Rules Directory [/etc/hex/rules]")
	PluginsDir := flag.String("plugins-dir", config.PluginsDir, "Plugins Directory [/etc/hex/plugins]")
	LogFile := flag.String("log-file", config.LogFile, "Log File")
	WorkspaceDir := flag.String("workspace-dir", config.WorkspaceDir, "Workspace Directory [/tmp]")
	Debug := flag.Bool("debug", config.Debug, "Debug [false]")
	BotName := flag.String("bot-name", config.BotName, "Bot Name [hex]")
	CLI := flag.Bool("cli", config.CLI, "CLI [false]")
	Auditing := flag.Bool("auditing", config.Auditing, "Audting [false]")
	AuditingFile := flag.String("auditing-file", config.AuditingFile, "Auditing File")
	Slack := flag.Bool("slack", config.Slack, "Slack [false]")
	SlackToken := flag.String("slack-token", config.SlackToken, "Slack Token")
	SlackIcon := flag.String("slack-icon", config.SlackIcon, "Slack Icon [:nut_and_bolt:]")
	SlackDebug := flag.Bool("slack-debug", config.SlackDebug, "Slack Debug [false]")
	Scheduler := flag.Bool("scheduler", config.Scheduler, "Scheduler [false]")
	Webhook := flag.Bool("webhook", config.Webhook, "Webhook [false]")
	WebhookPort := flag.Int("webhook-port", config.WebhookPort, "Webhook Port [8000]")
	flag.Parse()

	// set flags
	config.RulesDir = *RulesDir
	config.PluginsDir = *PluginsDir
	config.LogFile = *LogFile
	config.WorkspaceDir = *WorkspaceDir
	config.Debug = *Debug
	config.BotName = *BotName
	config.CLI = *CLI
	config.Auditing = *Auditing
	config.AuditingFile = *AuditingFile
	config.Slack = *Slack
	config.SlackToken = *SlackToken
	config.SlackIcon = *SlackIcon
	config.SlackDebug = *SlackDebug
	config.Scheduler = *Scheduler
	config.Webhook = *Webhook
	config.WebhookPort = *WebhookPort

}
