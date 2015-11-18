package commands

import (
	"github.com/mmcquillan/jane/models"
	"math/rand"
	"strings"
	"regexp"
)

var synRegex = regexp.MustCompile(`(?i)(SYN-)[0-9]+`)

func Parse(config *models.Config, message *models.Message) {

	if message.Request != "" {

		// loop through and find command matches
		var match []models.Command
		for _, c := range config.Commands {
			if strings.HasPrefix(strings.ToLower(message.Request), strings.ToLower(c.Match)) {
				match = append(match, c)
			}
		}

		// if no match, just leave
		if len(match) == 0 {

			matches := synRegex.FindAllString(message.Request, -1)

			var command models.Command

			for _, c := range config.Commands {
				if strings.HasPrefix(strings.ToLower("jira"), strings.ToLower(c.Match)) {
					command = c
				}
			}

			if len(matches) > 0 {
				r, desc, link := Jira(matches[0], command)

				if r != "" {
					message.Title = r
				}

				if desc != "" {
					message.Description = desc
				}

				if link != "" {
					message.Link = link
				}
			}

			return
		}

		// if more than one match, pick a random one
		var i = 0
		if len(match) > 0 {
			i = rand.Intn(len(match))
		}

		// send to a command type
		var r string
		switch match[i].Type {
		case "response":
			r = Response(message.Request, match[i])
		case "exec":
			r = Exec(message.Request, match[i])
		case "reload":
			r = Reload(match[i], config)
		case "help":
			r = Help(config)
		case "wolfram":
			r = Wolfram(message.Request, match[i])
		case "jira":
			var desc, link string
			r, desc, link = Jira(message.Request, match[i])

			if desc != "" {
				message.Description = desc
			}
			if link != "" {
				message.Link = link
			}
		}

		// feedback
		message.Title = r

	}

}
