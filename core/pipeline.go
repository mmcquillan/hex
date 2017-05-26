package core

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hexbotio/hex/actions"
	"github.com/hexbotio/hex/commands"
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/mohae/deepcopy"
	"github.com/rs/xid"
)

var state = make(map[string]models.State)

func Pipeline(inputMsgs <-chan models.Message, outputMsgs chan<- models.Message, config *models.Config) {
	initState(config)
	for {
		message := <-inputMsgs
		if config.Debug {
			log.Printf("PipelineEval: %+v", message)
		}
		aliasMessages(&message, config)
		messages := splitMessages(message)
		for _, message := range messages {
			aliasMessages(&message, config)
			runCommands(message, outputMsgs, config)
			for _, pipeline := range config.Pipelines {
				if pipeline.Active {
					for _, input := range pipeline.Inputs {
						var matchPipeline = true

						// match by type
						if !(input.Type == message.Inputs["hex.type"] || input.Type == "*" || input.Type == "") {
							matchPipeline = false
						}

						// match by service name
						if !(input.Name == message.Inputs["hex.name"] || input.Name == "*" || input.Name == "") {
							matchPipeline = false
						}

						// match by target
						if !(parse.Match(input.Target, message.Inputs["hex.target"]) || input.Target == "*" || input.Target == "") {
							matchPipeline = false
						}

						// match by input
						if !(parse.Match(input.Match, message.Inputs["hex.input"]) || input.Match == "*" || input.Match == "") {
							matchPipeline = false
						}

						// match by ACL
						if !(input.ACL == "" || input.ACL == "*") {
							matchAcl := false
							aclList := strings.Split(input.ACL, ",")
							for _, acl := range aclList {
								if message.Inputs["hex.user"] == strings.TrimSpace(acl) || message.Inputs["hex.target"] == strings.TrimSpace(acl) {
									matchAcl = true
								}
							}
							if !matchAcl {
								matchPipeline = false
							}
						}

						// if a match, then execute actions
						if matchPipeline {
							m := deepcopy.Copy(message).(models.Message)
							m.Inputs["hex.botname"] = config.BotName
							m.Inputs["hex.pipeline.name"] = pipeline.Name
							m.Inputs["hex.pipeline.alert"] = strconv.FormatBool(pipeline.Alert)
							m.Inputs["hex.pipeline.runid"] = xid.New().String()
							m.Inputs["hex.pipeline.workspace"] = config.Workspace + config.BotName + m.Inputs["hex.pipeline.runid"]
							m.Outputs = pipeline.Outputs
							go runActions(pipeline, m, outputMsgs, config)
						}

					}
				}
			}
		}
	}
}

func runActions(pipeline models.Pipeline, message models.Message, outputMsgs chan<- models.Message, config *models.Config) {
	if pipeline.Alert && state[pipeline.Name].Running {
		return
	}
	setRunning(pipeline.Name, true)
	for _, action := range pipeline.Actions {
		if actions.Exists(action.Type) {
			actionService := actions.Make(action.Type).(actions.Action)
			if actionService != nil {
				if message.Success || action.RunOnFail {
					actionService.Act(action, &message, config)
				}
			}
		}
	}
	setLastRun(pipeline.Name)
	if _, err := os.Stat(message.Inputs["hex.pipeline.workspace"]); err == nil {
		err := os.RemoveAll(message.Inputs["hex.pipeline.workspace"])
		if err != nil {
			log.Print("ERROR - Cleaning Workspace: " + message.Inputs["hex.pipeline.workspace"])
			log.Print(err)
		}
	}
	if config.Debug {
		log.Printf("PostAction: %+v", message)
	}
	if pipeline.Alert {
		if state[pipeline.Name].Success != message.Success {
			outputMsgs <- message
			setState(pipeline.Name, message.Success)
			setLastAlert(pipeline.Name)
		} else if !message.Success && (time.Now().Unix()-state[pipeline.Name].LastAlert) > (15*60) {
			outputMsgs <- message
			setLastAlert(pipeline.Name)
		}
	} else {
		outputMsgs <- message
	}
	setRunning(pipeline.Name, false)
}

func runCommands(message models.Message, outputMsgs chan<- models.Message, config *models.Config) {
	for command, _ := range commands.List {
		if parse.Match(command, message.Inputs["hex.input"]) {
			commandService := commands.Make(command).(commands.Action)
			if commandService != nil {
				commandService.Act(&message, config)
				var output models.Output
				output.Name = message.Inputs["hex.name"]
				output.Targets = message.Inputs["hex.target"]
				message.Inputs["hex.pipeline.name"] = "command"
				message.Outputs = append(message.Outputs, output)
				if config.Debug {
					log.Printf("PostCommand: %+v", message)
				}
				outputMsgs <- message
			}
		}
	}
}

func aliasMessages(message *models.Message, config *models.Config) {
	for _, alias := range config.Aliases {
		if parse.Match(alias.Match, message.Inputs["hex.input"]) {
			message.Inputs["hex.input"] = parse.Substitute(alias.Output, message.Inputs)
		}
	}
}

func splitMessages(message models.Message) (msgs []models.Message) {
	if strings.Contains(message.Inputs["hex.input"], "&&") {
		cmds := strings.Split(message.Inputs["hex.input"], "&&")
		for _, cmd := range cmds {
			var m = message
			m.Inputs["hex.input"] = strings.TrimSpace(cmd)
			msgs = append(msgs, m)
		}
	} else {
		msgs = append(msgs, message)
	}
	return msgs
}

func initState(config *models.Config) {
	for _, p := range config.Pipelines {
		state[p.Name] = models.State{
			LastRun:    0,
			LastChange: 0,
			LastAlert:  0,
			Success:    true,
			Running:    false,
		}
	}
}

func setRunning(pipeline string, running bool) {
	s := state[pipeline]
	s.Running = running
	state[pipeline] = s
}

func setState(pipeline string, success bool) {
	s := state[pipeline]
	if s.Success != success {
		s.LastChange = time.Now().Unix()
	}
	s.Success = success
	state[pipeline] = s
}

func setLastRun(pipeline string) {
	s := state[pipeline]
	s.LastRun = time.Now().Unix()
	state[pipeline] = s
}

func setLastAlert(pipeline string) {
	s := state[pipeline]
	s.LastAlert = time.Now().Unix()
	state[pipeline] = s
}
