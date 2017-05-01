# Hex

Hex is a chatops bot written in Go and is completely configuration driven. Contributions are welcome via pull requests. Hex was started as a way of getting DevOps tasks and feedback into Slack. There are a billion other bots, but we wanted to learn Go, so this was a fun way to learn it and meet our needs. The name "Hex" was chosen by @kcwinner because he is a big fan of the [_Ender's Game_ books](https://en.wikipedia.org/wiki/Hex_(Ender%27s_Game)). The name is not meant to be gender specific and can be effectively changed when you set your bot up.

* [Installing](#Installing)
* [Concepts](#Concepts)
* [Configuration](#Configuration)
* [Services](#Services)
* [Piplines](#Piplines)
* [Actions](#Actions)
* [Aliases](#Aliases)
* [Matching](#Matching)
* [Substitutions](#Substitutions)
* [Getting Involved](#getting-involved)


## Installing
Run with [Docker](https://hub.docker.com/r/hexbotio/hex/):
```
docker run --name my-hex -v /some/hex.json:/etc/hex.json -d hexbotio/hex:v2.0
```

Manually with Go 1.8.1:
```
go get github.com/hexbotio/hex
go install github.com/hexbotio/hex
```


## Concepts
* Action - A step that happens as part of a Pipeline.
* Input - A Service that is the trigger for a Pipeline.
* Message - The event from an Input that passes through a Pipeline.
* Output - A Service that takes the results of a Pipeline.
* Pipeline - The collection of actions that are triggered by one or more Inputs where the results go to one or more Outputs.
* Service - An enabled external service which Hex can interact with.


## Configuration
The configuration of Hex is via a json config file and will be looked for in this order:
* `HEX_CONFIG` - Environment variable
* `--config <file name>` - Command line parameter
* `./hex.json` - the location of the hex binary
* `~/hex.json` - the home directory of the user
* `/etc/hex.json` - the global config

The basic configuration file should include these elements:
```
{
  "BotName": "hex",
  "Debug": false,
  "LogFile": "/var/log/hex.log",
  "Aliases": [
  ],
  "Services": [
  ],
  "Pipelines": [
  ]
}
```

To protect sensitive data, you can use environment variables throughout the configuration file's Config values with the format `${SLACK_TOKEN}`. Additionally, environment variables can be used in the Command property of Pipeline Actions.

Hex has these builtin environment variables:

* HEX_CONFIG - The path to the configuration file
* HEX_LOGFILE - The path to the logfile
* HEX_WORKSPACE - The temporary workspace where pipelines can write to
* HEX_DEBUG - Set as true/false to globally turn on logging
* HEX_BOT_NAME - If you want to set your own bot name


## Services
Services are how you can activate features within Hex to get inputs and outputs with the outside world. Below are the list of supported Services and example configuration for each.
* [CLI Service](#cli-service)
* [File Service](#file-service)
* [RSS Service](#rss-service)
* [Scheduler Service](#scheduler-service)
* [Slack Service](#slack-service)
* [SSH Service](#ssh-service)
* [Twitter Service](#twitter-service)
* [Webhook Service](#webhook-service)
* [WinRM Service](#winrm-service)

#### CLI Service
This service enables the command line interface, not meant to be run as a daemon. It is helpful for debugging or just indulging your command line love.

Supports:
* Inputs
* Outputs

Example:
```
{
  "Services": [
    {"Type": "cli", "Name": "my cli", "Active": true}
  ],
}
```

#### File
This service enables files for reading or writing, though the file must reside on the same system that the Hex process runs on.

Supports:
* Inputs (if Mode is set to "r" and will optionally filter results to the Filter [mached](#matching) value)
* Outputs (if Mode is set to "w")

Example:
```
{
  "Services": [
    {"Type": "file", "Name": "my log", "Active": true,
      "Config": {
        "File": "/var/log/output.log",
        "Filter": "*error*",
        "Mode": "r"
      }
    }
  ],
}
```

#### RSS Service
This service listens for an RSS feed.

Supports:
* Inputs

Variables:
* `${hex.rss.title}` - RSS Feed title
* `${hex.rss.content}` - RSS Feed content
* `${hex.rss.date}` - RSS Feed date
* `${hex.rss.link}` - RSS Feed link

Example:
```
{
  "Services": [
    {"Type": "file", "Name": "my log", "Active": true,
      "Config": {
        "File": "/var/log/output.log",
        "Filter": "*error*",
        "Mode": "r"
      }
    }
  ],
}
```

#### Scheduler Service
This service fires a message on the cron schedule specified.

Supports:
* Inputs

Example:
```
{
  "Services": [
    {"Type": "scheduler", "Name": "every minute", "Active": false,
      "Config": {
        "Schedule": "0 * * * * *"
      }
    }
  ],
}
```

#### Slack Service
This service enables slack integration.

Supports:
* Inputs
* Outputs

Example:
```
{
  "Services": [
    {"Type": "slack", "Name": "my slack", "Active": true,
      "Config": {
        "Key": "${SLACK_TOKEN}",
        "Image": ":hex:",
        "SlackDebug": "false"
      }
    }
  ],
}
```

#### SSH Service
This service enables a server endpoint for SSH commands.

Supports:
* Actions

Example:
```
{
  "Services": [
    {"Type": "ssh", "Name": "my server", "Active": true,
      "Config": {
        "Server": "myserver.com",
        "Port": "22",
        "Login": "hex",
        "Pass": "${MYSERVER_PASSWORD}"
      }
    }
  ],
}
```

#### Twilio Service
This service will send SMS messages with an active Twilio account.

Supports:
* Outputs

Example:
```
{
  "Services": [
    {"Type": "twilio", "Name": "my twilio", "Active": true,
      "Config": {
        "Key": "${API_KEY}",
        "Pass": "${AUTH_TOKEN}",
        "From": "${FROM_NUMBER}"
      }
    }
  ],
}
```

#### Twitter Service
This service enables the twitter feed.

Supports:
* Inputs

Variables:
* `${hex.twitter.lang}` - Twitter tweet language

Example:
```
{
  "Services": [
    {"Type": "twitter", "Name": "my twitter", "Active": false,
      "Config": {
        "Key": "${APP_TWITTER_KEY}",
        "Secret": "${APP_TWITTER_SECRET}",
        "AccessToken": "${APP_ACCESS_TOKEN}",
        "AccessTokenSecret": "${APP_ACCESS_TOKEN_SECRET}",
        "Filter": "filter1,filter2"
      }
    }
  ],
}
```

#### Webhook Service
This service listens for incoming webhooks.

Supports:
* Inputs

Example:
```
{
  "Services": [
    {"Type": "webhook", "Name": "my webhook", "Active": true,
      "Config": {
        "Port": "8080"
      }
    }
  ],
}
```

#### WinRM Service
This service enables a server endpoint for WinRM commands.

Supports:
* Actions

Example:
```
{
  "Services": [
    {"Type": "winrm", "Name": "my server", "Active": true,
      "Config": {
        "Server": "myserver.com",
        "Port": "5985",
        "Login": "hex",
        "Pass": "${MYSERVER_PASSWORD}"
      }
    }
  ],
}
```


## Pipelines
A pipeline allows you to match incoming commands, schedules or webhooks to a series of actions.
* Name - A unique name for the Pipeline
* Active - A way to mark pipelines as being active or note (true/false)
* Alert - Treat this pipeline as an alert and only report on state change
* Inputs - One or more Service Inputs to match against. Wildcards of `*` can be used to match allows.
  * Type - Match against a Service Type
  * Name - Match against a Service Name
  * Target - Match against a Service Target (ex: Slack Channel)
  * Match - Match against the Message Input from the User
  * ACL - Match against a comma delimited list of Users or Targets to limit who can run the Pipeline
* Actions - One or more ordered Actions to be performed, see (Actions)[#Actions]
* Outputs - One or more Service Outputs to send the results of the Pipeline to
  * Name - The Service Name to send to
  * Targets - A comma delimited list of Targets to send to in the service (ex: Slack Channels)

Example Checking Disk space every minute:
```
{
  "Pipelines": [
    {
      "Name": "Check Disk",
      "Active": true,
      "Alert": true,
      "Inputs": [
        {"Type": "schedule", "Name": "every minute", "Target": "*", "Match": "*", "ACL": "*"}
      ],
      "Actions": [
        {
          "Name": "chk disk", "Type": "ssh", "Command": "/usr/lib/nagios/plugins/check_disk -w 15% -c 5%",
          "Service": "prd*", Success:"DISK OK"
        }
      ],
      "Outputs": [
        {"Name": "my slack", "Targets": "#alerts"}
      ]
    }
  ]
}
```

Example RSS Feed to a Sack Channel:
```
{
  "Pipelines": [
    {
      "Name": "AWS RSS Feeds",
      "Active": true,
      "Alert": false,
      "Inputs": [
        {"Type": "rss", "Name": "*", "Target": "*", "Match": "*", "ACL": "*"}
      ],
      "Actions": [
        {
          "Name": "Respond", "Type": "format", "Command": "${hex.rss.title}"
        }
      ],
      "Outputs": [
        {"Name": "my slack", "Targets": "#alerts"}
      ]
    }
  ]
}
```

## Actions
Actions are performed in the pipeline. Some actions depend on activated Services, such as SSH to perform their action against.
* [Format Action](#format-action)
* [SSH Service](#ssh-action)
* [WinRM Service](#winrm-action)

#### Format Action
This will format the input statement for the next action.
* Name - A name for the Action
* Type - "format"
* Command - The formating for the input rewrite following the [Substitutions](#substitutions) format
* Success - A string to [Match](#matching) against the results to determine if successful
* Failure - A string to [Match](#matching) against the results to determine if not successful
* RunOnFail - A true/false value to specify if the action should run if the previous action failed

Example:
```
"Actions": [
  {
    "Name": "My Message ID", "Type": "format", "Command": "Hey ${hex.user}, your Message ID is ${hex.id}"
  }
]
```

#### SSH Action
This Action will SSH out to the specified SSH Services and run the Command.
* Name - A name for the Action
* Type - "ssh"
* Command - The command to run on the target system which supports the [Substitutions](#substitutions) format
* Service - A string to [Match](#matching) against SSH Services defined
* Success - A string to [Match](#matching) against the results to determine if successful
* Failure - A string to [Match](#matching) against the results to determine if not successful
* RunOnFail - A true/false value to specify if the action should run if the previous action failed

Example:
```
"Actions": [
  {
    "Name": "chk disk", "Type": "ssh", "Command": "/usr/lib/nagios/plugins/check_disk -w 15% -c 5%",
    "Service": "prd*", Success:"DISK OK"
  }
]
```

#### WinRM Action
This Action will SSH out to the specified SSH Services and run the Command.
* Name - A name for the Action
* Type - "winrm"
* Command - The command to run on the target system which supports the [Substitutions](#substitutions) format
* Service - A string to [Match](#matching) against SSH Services defined
* Success - A string to [Match](#matching) against the results to determine if successful
* Failure - A string to [Match](#matching) against the results to determine if not successful
* RunOnFail - A true/false value to specify if the action should run if the previous action failed

Example:
```
"Actions": [
  {
    "Name": "C Directory", "Type": "ssh", "Command": "dir c:\",
    "Service": "qa-web"
  }
]
```


## Aliases
With Hex, you can create aliases for commands.

Example:
```
"Aliases": [
  {"Match": "hex monitor prod", "Output": "hex monitor prod1 && hex monitor prod2 && hex monitor prod3", "HideHelp": false},
  {"Match": "hex deploy*", "Output": "hex upgrade service ${1}", "Help": "hex deploy <service name>"}
]
```


## Matching
Hex uses a consistent string matching method.

Examples:
* `*failure*` - Match anywhere in a string
* `failure*` - Match at the beginning of a string
* `*failure` - Match at the end of a string
* `/fail(.+)/` - Regular expression matching


## Substitutions
Hex uses a consistent format for variable substitution. Each service may add additional variables for substitution.

Built-in Variables:
* `${ENVIRONMENT_VARIABLE}` - Environment variables are always available
* `${hex.id}` - A unique identifier for each message that passes through
* `${hex.timestamp}` - An epoch timestamp of the Message's creation
* `${hex.type}` - The type of Service which created the Message
* `${hex.name}` - The name of the Service which created the Message
* `${hex.target}` - The target from which the Message was created
* `${hex.user}` - The user who created the Message
* `${hex.input}` - The raw input of Message
* `${hex.input.0}` - The first word of the input
* `${hex.input.1}` - The second word of the input
* `${hex.input.2:*}` - The third through last word of the input
* `${hex.input.3:4}` - The fourth through fifth word of the input
* `${hex.input.json:web.api}` - The json value at {"web": {"api": "some value"}}
* `${hex.pipeline.name}` - The Name of the matching Pipeline
* `${hex.pipeline.alert}` - If this is set as an alert or not
* `${hex.pipeline.runid}` - A Unique identifier for each instance of a Pipeline run
* `${hex.pipeline.workspace}` - Location of the workspace for this Pipeline run


## Getting Involved
Development Environment
* Get your toes wet with Go
* Setup your Go 1.8.1 environment
* Pull the project with 'go get github.com/hexbotio/hex'
* Compile with 'go install hex.go'
* Use the sample _hex.json_ file checked in as a starting point
* Run your code and config with `go run hex.go --config ~/hex.json`

