package core

import (
	"log"
	"strings"

	"github.com/projectjane/jane/actions"
	"github.com/projectjane/jane/internals"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
)

func Pipeline(inputMsgs <-chan models.Message, outputMsgs chan<- models.Message, config *models.Config) {
	for {
		message := <-inputMsgs
		if config.Debug {
			log.Printf("PipelineEval: %+v", message)
		}
		aliasMessages(&message, config)
		messages := splitMessages(message)
		for _, message := range messages {
			aliasMessages(&message, config)
			runInternals(message, outputMsgs, config)
			for _, pipeline := range config.Pipelines {
				if pipeline.Active {
					for _, input := range pipeline.Inputs {
						var matchPipeline = true

						// match by type
						if !(input.Type == message.Inputs["jane.type"] || input.Type == "*" || input.Type == "") {
							matchPipeline = false
						}

						// match by service name
						if !(input.Name == message.Inputs["jane.name"] || input.Name == "*" || input.Name == "") {
							matchPipeline = false
						}

						// match by target
						if !(parse.Match(input.Target, message.Inputs["jane.target"]) || input.Target == "*" || input.Target == "") {
							matchPipeline = false
						}

						// match by input
						if !(parse.Match(input.Match, message.Inputs["jane.input"]) || input.Match == "*" || input.Match == "") {
							matchPipeline = false
						}

						// match by ACL
						if !(input.ACL == "" || input.ACL == "*") {
							aclList := strings.Split(input.ACL, ",")
							for _, acl := range aclList {
								if message.Inputs["jane.user"] != strings.TrimSpace(acl) && message.Inputs["jane.target"] != strings.TrimSpace(acl) {
									matchPipeline = false
								}
							}
						}

						// if a match, then execute actions
						if matchPipeline {
							message.Inputs["jane.pipeline"] = pipeline.Name
							message.Outputs = pipeline.Outputs
							go runActions(pipeline.Actions, message, outputMsgs, config)
						}

					}
				}
			}
		}
	}
}

func runActions(actionList []models.Action, message models.Message, outputMsgs chan<- models.Message, config *models.Config) {
	for _, action := range actionList {
		if actions.Exists(action.Type) {
			actionService := actions.Make(action.Type).(actions.Action)
			if actionService != nil {
				if message.Success || action.RunOnFail {
					actionService.Act(action, &message, config)
				}
			}
		}
	}
	if config.Debug {
		log.Printf("PostAction: %+v", message)
	}
	outputMsgs <- message
}

func runInternals(message models.Message, outputMsgs chan<- models.Message, config *models.Config) {
	for internal, _ := range internals.List {
		if parse.Match(config.BotName+" "+internal, message.Inputs["jane.input"]) {
			internalService := internals.Make(internal).(internals.Action)
			if internalService != nil {
				internalService.Act(&message, config)
				var output models.Output
				output.Name = message.Inputs["jane.name"]
				output.Targets = message.Inputs["jane.target"]
				message.Inputs["jane.pipeline"] = "internal"
				message.Outputs = append(message.Outputs, output)
				if config.Debug {
					log.Printf("PostInternal: %+v", message)
				}
				outputMsgs <- message
			}
		}
	}
}

func aliasMessages(message *models.Message, config *models.Config) {
	for _, alias := range config.Aliases {
		if parse.Match(alias.Match, message.Inputs["jane.input"]) {
			message.Inputs["jane.input"] = parse.Substitute(alias.Output, message.Inputs)
		}
	}
}

func splitMessages(message models.Message) (msgs []models.Message) {
	if strings.Contains(message.Inputs["jane.input"], "&&") {
		cmds := strings.Split(message.Inputs["jane.input"], "&&")
		for _, cmd := range cmds {
			var m = message
			m.Inputs["jane.input"] = strings.TrimSpace(cmd)
			msgs = append(msgs, m)
		}
	} else {
		msgs = append(msgs, message)
	}
	return msgs
}
