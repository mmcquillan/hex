package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/mmcquillan/hex/models"
	"github.com/mmcquillan/hex/parse"
	"github.com/mmcquillan/matcher"
	"gopkg.in/yaml.v2"
)

// Config func
func Config(version string) (config models.Config) {

	// set version default
	if version == "" {
		version = "unknown"
	}

	// defaults
	config.Logger = hclog.New(nil)
	config.Version = version
	config.Admins = ""
	config.UserACL = "*"
	config.ChannelACL = "*"
	config.PluginsDir = ""
	config.RulesDir = ""
	config.LogFile = ""
	config.LogLevel = "info"
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
	config.Command = ""

	// saved
	_, _, values := matcher.Matcher("<bin> [config] [--]", strings.Join(os.Args, " "))
	if values["config"] != "" {
		fileName := values["config"]
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			config.Logger.Error("ERROR: Config File Read - " + err.Error())
			os.Exit(1)
		}
		subFile := parse.SubstituteEnv(string(file[:]))
		if strings.HasSuffix(fileName, ".json") {
			err = json.Unmarshal([]byte(subFile), &config)
			if err != nil {
				config.Logger.Error("ERROR: Config File json Unmarshal - " + err.Error())
				os.Exit(1)
			}
		} else if strings.HasSuffix(fileName, ".yaml") || strings.HasSuffix(fileName, ".yml") {
			err = yaml.Unmarshal([]byte(subFile), &config)
			if err != nil {
				config.Logger.Error("ERROR: Config File yaml Unmarshal - " + err.Error())
				os.Exit(1)
			}
		} else {
			config.Logger.Error("ERROR: Config File Unknown Type")
			os.Exit(1)
		}
	}

	// reflect on things
	obj := reflect.TypeOf(config)

	// env vars
	envVarPrefix := "HEX_"
	for i := 0; i < obj.NumField(); i++ {
		f := obj.Field(i)
		if f.Tag.Get("json") != "" {
			env := strings.ToUpper(envVarPrefix + f.Tag.Get("json"))
			v := reflect.ValueOf(&config).Elem().FieldByName(f.Name)
			if v.CanSet() && os.Getenv(env) != "" {
				switch fmt.Sprintf("%v", f.Type) {
				case "string":
					v.SetString(os.Getenv(env))
				case "int":
					if val, err := strconv.ParseInt(os.Getenv(env), 10, 32); err == nil {
						v.SetInt(val)
					}
				case "bool":
					if val, err := strconv.ParseBool(os.Getenv(env)); err == nil {
						v.SetBool(val)
					}
				}
			}
		}
	}

	// flags
	for i := 0; i < obj.NumField(); i++ {
		f := obj.Field(i)
		if f.Tag.Get("json") != "" {
			flag := strings.Replace(f.Tag.Get("json"), "_", "-", -1)
			v := reflect.ValueOf(&config).Elem().FieldByName(f.Name)
			if v.CanSet() && values[flag] != "" {
				switch fmt.Sprintf("%v", f.Type) {
				case "string":
					v.SetString(values[flag])
				case "int":
					if val, err := strconv.ParseInt(values[flag], 10, 32); err == nil {
						v.SetInt(val)
					}
				case "bool":
					if val, err := strconv.ParseBool(values[flag]); err == nil {
						v.SetBool(val)
					}
				}
			}
		}
	}

	// validate
	if config.UserACL == "" {
		config.Logger.Warn("WARNING: Setting a blank User ACL will result in nothing happening.")
	}
	if config.ChannelACL == "" {
		config.Logger.Warn("WARNING: Setting a blank Channel ACL will result in nothing happening.")
	}
	if !parse.Member("ERROR,INFO,DEBUG,TRACE", strings.ToUpper(config.LogLevel)) {
		config.Logger.Error("ERROR: Log Level should be error, info, debug or trace.")
		os.Exit(1)
	}
	if config.Slack && config.SlackToken == "" {
		config.Logger.Error("ERROR: Slack is enabled, but no Slack Token is specified.")
		os.Exit(1)
	}
	if config.Auditing && config.AuditingFile == "" {
		config.Logger.Error("ERROR: Auditing is enabled, but no Auditing File is specified.")
		os.Exit(1)
	}

	return config

}
