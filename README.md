# Jane

Jane is a chatops bot written in Go and is completely configuration driven. Contributions are welcome via pull requests. Jane was started as a way of getting DevOps tasks and feedback into Slack. There are a billion other bots, but we wanted to learn Go, so this was a fun way to learn it and meet our needs. The name "Jane" was chosen by @kcwinner because he is a big fan of the [_Ender's Game_ books](https://en.wikipedia.org/wiki/Jane_(Ender%27s_Game)). The name is not meant to be gender specific and can be effectively changed when you set your bot up.


## Running

### Docker
Use the [Docker Hub Repo](https://hub.docker.com/r/projectjane/jane/) then run:

```
docker run --name my-jane -v /some/jane.json:/etc/jane.json -d projectjane/jane
```

### Install
You can install Go and build Jane. Look at example startup scripts under the startup directory.

### Configuration
The configuration of Jane is via a json config file. The configuration file is expected to be named 'jane.json' and will be looked for in this order:
* --config <file name> - Pass in a configuration file location as a command line parameter
* ./jane.json - the location of the jane binary
* ~/jane.json - the home directory of the user
* /etc/jane.json - the global config


### Basic Configuration File
The basic configurtion file should include these elements:

```
{
  "LogFile": "/var/log/jane.log",
  "Aliases": [
  ],
  "Connectors": [
  ],
  "Routes": [
  ]
}
```


### Environment Variables
To protect sensitive data, you can set Connector Server, Port, Login and Pass as an environment variable, should the connector support those values. Use the format `${SLACK_TOKEN}` to use the environment variable SLACK_TOKEN. Additionally, environment variables can be used in output values on many connectors as well.


## Connectors
Connectors are what Jane uses to pull in information, interpret it, and issue out a response. The Routes specify where the results from the input should be written to or * for all. The Target can specify a channel in the case of Slack. 

For the connector configuration, when adding routes, you must specify the ID of the connector you want to route the response to.

Supported connectors:
* [bamboo](#bamboo-connector) - Atlassian Bamboo integration
* [cli](#cli-connector) - Command line interface
* [email](#email-connector) - Email
* [exec](#exec-connector) - Execution of commands with monitoring capability
* [file](#file-connector) - File watcher
* [imageme](#imageme-connector) - Pull back images or animated gifs
* [jira](#jira-connector) - Atlassian Jira integration
* [log](#log-connector) - Log connector to audit Jane calls
* [response](#response-connector) - Text Responses
* [rss](#rss-connector) - RSS Feed
* [slack](#slack-connector) - Slack chat
* [twilio](#twilio-connector) - send SMS alerts
* [website](#website-connector) - Monitor return code of websites
* [webhook](#webhook-connector) - Listener for webhooks
* [wolfram](#wolfram-connector) - Execute queries against Wolfram Alpha


### Bamboo Connector

This connector was written to integrate bamboo builds. It was written against the bamboo cloud/hosted solution, but should be compatible with installed versions. This connector listens for and displays builds and deploys in addition to letting you execute the builds (but not deploys because the bamboo api is dated).

#### Example:

```
{"Type": "bamboo", "ID": "bamboo server", "Active": true, "Debug": true,
   "Server": "<URL>.atlassian.net", "Login": "<JIRA USER>", "Pass": "<JIRA PASS>"
}
```

#### Usage:
* Run builds: `bamboo build <build key>`
* Get build status: `bamboo status <environment or build key>`
* Make sure to sepecify where you want the build and deploy messages to end up in the routes

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'exec2'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Server_ - The bamboo server address
* _Login_ - The bamboo user to login with
* _Pass_ - The bamboo password to connect with
* _Users_ - List of users who can execute the commands in this connector [security](#security)


### Cli Connector

This connector runs Jane via the command line interface instead of as a daemon and is helpful for debugging, or just indulging your command line love.

#### Example:

```
{"Type": "cli", "ID": "term-bot", "Active": true, "Debug": false
}
```

#### Usage:
* Use the command prompt to type your command and enter to send it.

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'exec2'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs


### Email Connector

*NOTE - This is experimental and untested*
This connector allows for the sending of emails. Point the connector to a valid SMTP server.

#### Example:

```
{"Type": "email", "ID": "EmailServer", "Active": false,
  "Server": "smtp-server.myserver.com", "Port": "465",
  "Login": "smtpuser", "Pass": "smtppassword",
  "From": "jane@myserver.com"
}
```

#### Usage:
* Make sure to sepecify the to address in the target for routes to send emails

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'exec2'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Server_ - The SMTP server
* _Port_ - The port for the SMTP server
* _Login_ - The SMTP login
* _Pass_ - The SMTP password


### Exec Connector

This connector provides a single means of making local and remote calls to Linux systems. You can allow these calls to be made by command, but also mark the calls with the RunCheck property to set Jane to check them. This combined with the interpreter for output, makes it a very capable monitoring platform.

#### Example:

```
{"Type": "exec", "ID": "Elastic Search", "Active": true,
  "Server": "elasticsearch1.somecompany.com", "Port": "22", "Login": "jane", "Pass": "abc123",
  "Commands": [
    {
        "Name": "Apt Check",
        "Match": "jane elasticsearch1 aptcheck",
        "Output": "```${STDOUT}```",
        "Cmd": "/usr/lib/nagios/plugins/check_apt",
        "Args": "",
        "HideHelp": false,
        "Help": "jane elasticsearch1 aptcheck - To check our elasticsearch!",
        "RunCheck": true,
        "Interval": 1,
        "Remind": 15,
        "Green": "*OK*",
        "Yellow": "*WARNING*",
        "Red": "*CRITICAL*"
    }
  ]
}
```

#### Usage:
* To make local calls to the system, leave out the Server, Port, Login, Pass values.

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'exec'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Server_ - The server address or IP to connect to
* _Port_ - The port number to connect to (Default: 22)
* _Login_ - The user to login with
* _Pass_ - The password to connect with
* _Users_ - List of users who can execute the commands in this connector [security](#security)
* _Commands_ - One or more commands to execute against the defined server
  * _Name_ - Readable name of check
  * _Match_ - Command [match](#matching)
  * _Output_ - Formatting for the output of the command, use `${STDOUT}` as the output
  * _Cmd_ - The command to execute (do not include arguments)
  * _Args_ - The arguments, space deliminated (you can access anything after the match above with positional variables like ${1}, ${2}, etc or ${*} for all input after the match)
  * _HideHelp_ - A boolean to show or hide the help when displaying help (Default: false)
  * _Help_ - Optional help text, otherwise it'll show the Match value
  * _RunCheck_ - A boolean that will have Jane periodically run this (Default: false)
  * _Interval_ - An integer that is the number of minutes between checks when RunCheck is true (Default: 1)
  * _Remind_ - An integer which is the number of units of Interval to wait before reminding of a non-Green status, with Zero being no reminders (Default: 0)
  * _Green_ - A [match](#matching) to identify what is in a green state
  * _Yellow_ - A [match](#matching) to identify what is in a yellow state
  * _Red_ - A [match](#matching) to identify what is in a red state


### File Connector
The File watch connector will monitor a local file and throw an alert anytime a matched string is detected.

#### Example:

```
{"Type": "file", "ID": "jane log", "Active": true,
  "File": "/home/matt/test.log",
  "Commands": [
    {"Name": "Starting Jane", "Match": "*stopping*"},
    {"Name": "Starting Jane", "Match": "*starting*"}
  ]
}
```

#### Usage:
* This is currently limited to the system in which Jane runs on

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'logging'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _File_ - The file to watch (make sure the Jane process has permission)
* _Commands_ - One or more commands to execute against the defined server
  * _Name_ - Readable name of check
  * _Match_ - String to match in the file [match](#matching)


### ImageMe Connector
Description

#### Example:

```
{"Type": "imageme", "ID": "imageme", "Active": true,
  "Key": "<GOOGLE API KEY>", "Pass": "<GOOGLE API PASS>"
}
```

#### Usage:
* Type `imageme <some text>` for an image url to be returned
* Type `animateme <some text>` for an animated gif to be returned

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'imageme'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Key_ - The Google API Key
* _Pass_ - The Google API Pass
* _Users_ - List of users who can execute the commands in this connector [security](#security)


### Jira Connector
This connector will integrate with your Jira server.

#### Example:

```
{"Type": "jira", "ID": "jira", "Active": true, "Debug": true,
    "Server": "<URL>.atlassian.net", "Login": "<JIRA USER>", "Pass": "<JIRA PASS>"
}
```

#### Usage:
* When a Jira ticket is detected, a link to the ticket will be provided
* You can also create a Jira issue `jira create <issueType> <project key> <summary>`

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'jira'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Server_ - The server url of your Jira instance
* _Login_ - A Jira user login
* _Pass_ - A Jira password
* _Users_ - List of users who can execute the commands in this connector [security](#security)


### Log Connector
The Log connector will write message output to a log file, usually for auditing purposes.

#### Example:

```
{"Type": "log", "ID": "audit", "Active": true,
  "File": "/home/matt/messages.log"
}
```

#### Usage:
* This is currently limited to the system in which Jane runs on

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'logging'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _File_ - The file to watch (make sure the Jane process has permission)


### Response Connector
Jane is capable of responding to questions or outputing fixed phrases.

#### Example:

```
{"Type": "response", "ID": "Text Response", "Active": true, "Debug": false,
    "Commands": [
        {"Match": "jane rules", "Output": "*The Three Laws of DevOps Robotics*\n\n1. A robot may not injure a production environment or, through inaction, allow a production environment to come to harm.\n\n2. A robot must obey the orders given it by a command line interface except where such orders would conflict with the First Law.\n\n3. A robot must protect its own production existence as long as such protection does not conflict with the First or Second Laws."}
    ]
}
```

#### Usage:
* You can integrate the input with the output by using the substitution values `${1}` for the first word, or `${*}` for the entire string

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'response'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Users_ - List of users who can execute the commands in this connector [security](#security)
* _Commands_ - One or more commands to execute against the defined server
  * _Match_ - String to [match](#matching)
  * _Output_ - String to output, with [substitutions](#substitutions)
  * _HideHelp_ - A boolean to show or hide the help when displaying help (Default: false)
  * _Help_ - Optional help text, otherwise it'll show the Match value


### RSS Connector
Pull in RSS feeds directly into your bot and see what's going on around the web. Some suggestions are cloud or vendor status pages.

#### Example:

```
{"Type": "rss", "ID": "AWS EC2", "Active": true,
    "Server": "http://status.aws.amazon.com/rss/ec2-us-east-1.rss"
}
```

#### Usage:
* Some RSS feeds are empty and cause problems for parsing, report any that give you issues

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'rss'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Server_ - The server address or IP to connect to


### Slack Connector
Slack integration allows for Jane to be interacted with through Slack by both sending and receiving messages.

#### Example:

```
{"Type": "slack", "ID": "slack", "Active": true,
  "Key": "<SlackToken>", "Image": ":game_die:"
}
```

#### Usage:
* Setup Jane as a new integration to get a Slack Token
* Add Jane to the channels you wish for the bot to listen to
* You can even direct message Jane
* Be sure to set the Route Target for specifying where listeners send their messages

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'slack'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Key_ - The Slack Token
* _Image_ - The Slack image to use for messages


### Twilio Connector
Twilio provides SMS and Phone integration.

#### Example:

```
{"Type": "twilio", "ID": "twilio", "Active": true,
  "Key": "<API_KEY>",
  "Pass": "<AUTH_TOKEN>",
  "From": "<FROM_NUMBER>"
}
```

#### Usage:
* For the target in other connector's routes, you can set the target phone number

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'twilio'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Key_ - The Twilio API Key
* _Pass_ - The Twilio Authorization Token
* _From_ - The Number to send from


### Website Connector
Basic website monitoring that throws alerts when it does not get a 200 OK status from the web server.

#### Example:

```
{"Type": "website", "ID": "Website Monitor", "Active": true,
  "Server": "https://google.com"
}
```

#### Usage:
* Test the URL to make sure you are pointed to one that gets back a 200 response

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'website'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Server_ - The website to monitor


### Webhook Connector

This connector opens a port for Jane to receive webhook calls. Webhooks calls are matched against the command list matches. Json can be interpreted and used to substitute into the output string. 


#### Example:

```
{"Type": "webhook", "ID": "Integrations", "Active": true, "Debug": true,
  "Port": "8080",
  "Commands": [
    {
        "Name": "Loggly Alerts",
        "Match": "/loggly/alerts",
        "Process": false,
        "Output": "```${alert_name} - ${search_link}```",
        "Red": "*alert*"
    },
    {
        "Name": "Git Commits",
        "Match": "/git/commit",
        "Process": true,
        "Output": "jane build stuff"
    },
    {
        "Name": "Messages",
        "Match": "/messages",
        "Process": false,
        "Output": "${?}"
    }
  ]
}
```

#### Usage:
* For this connector, the `${?}` represents everything in the URL after the ?

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'exec2'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Port_ - The port number to listen to (should be above 1024 if not running as a privledged user)
* _Users_ - List of users who can execute the commands in this connector [security](#security)
* _Commands_ - One or more commands to match the incoming webhook
  * _Name_ - Name of the matching webhook check
  * _Match_ - Webhook URL [match](#matching) (this will always be after the server name and port)
  * _Process_ - This defines if the incoming message should be processed by the other connector commands (true) or just published out to the routes (false) (Default: false)
  * _Output_ - This is the formatting for the output. Use the [json parsing rules](https://github.com/Jeffail/gabs#parsing-and-searching-json), '${?}' to output the query string or the [substitutions](#substitutions).
  * _Green_ - A [match](#matching) to identify what is in a green state
  * _Yellow_ - A [match](#matching) to identify what is in a yellow state
  * _Red_ - A [match](#matching) to identify what is in a red state


### Wolfram Connector
This is an integration to the Wolfram Alpha API

#### Example:

```
{"Type": "wolfram", "ID": "wolf", "Active": true, "Debug": true,
  "Key": "<WOLFRAM API KEY>"
}
```

#### Usage:
* Type `wolfram <query>` to get results back

#### Fields:
* _Type_ - This specifies the type of connector, in this case, 'wolfram'
* _ID_ - This should be a unique identifier for this connector
* _Active_ - This is a boolean value to set this connector to be activated
* _Debug_ - This is a boolean value to set if the connector shows debug information in the logs
* _Key_ - Wolfram API Key


## Core Concepts


### Aliases

With Jane, you can create aliases for commands.

#### Example:

```
"Aliases": [
  {"Match": "jane build app", "Output": "bamboo build KEY-PLAN"},
  {"Match": "jane monitor prod", "Output": "jane monitor prod1 && jane monitor prod2 && jane monitor prod3"},
  {"Match": "jane deploy*", "Output": "jane upgrade service ${1}"}
]
```

#### Fields:
* _Match_ - This will match incoming commands with the [match](#matching) rules
* _Output_ - This is the command that will be run supports [substitutions](#substitutions)
* _HideHelp_ - A boolean to show or hide the help when displaying help (Default: false)
* _Help_ - Optional help text, otherwise it'll show the Match value


### Routes

Routes can exist for connectors that listen to or interpret commands. Routes can have more than one connector if you would like to send messages to more than one place. Jane also matches on the routes to filter which messages get sent.

#### Example:

```
"Routes": [
    {
        "Matches": [{"ConnectorType": "*", "ConnectorID": "jane-slack", "Target": "*", "User": "*", "Message": "*"}],
        "Connectors": "jane-slack", "Targets": "#alerts,*"
    }
]
```

#### Usage:
* Some connector publishers allow you to specify a Target, such as Slack which uses a target for a channel
* Connectors and Targets allow for a comma seperated list of each
* Message follows the Jane [match rules](#matching)
* All other fields use an exact match or * for wild card
* A "*" as a Slack target will return the message to the originating user or channel

#### Fields:
* _Matches_ - Multiple matches against messages
  * _ConnectorType_ - The type such as "slack" or "rss"
  * _ConnectorID_ - The unique name for a specific connector
  * _Target_ - The incoming target (such as channel in slack)
  * _User_ - The user the message originated from
  * _Message_ - This will match the message or any message with "*" using the [match](#matching)
* _Connectors_ - A comma seperated list of connector name (ID) or "*" to match all connectors
* _Target_ - A comma seperated list of target which is connector specific or "*" for all


### Matching

Jane uses a consistent string matching method throughout.

#### Examples:

`*failure*` - Match anywhere in a string

`failure*` - Match at the beginning of a string

`*failure` - Match at the end of a string

`/fail(.+)/` - Regular expression matching


### Substitutions

Jane uses a format for substitutions which is implemented in most areas. Each connector can add it's own subsitution values.

#### Examples:

`${ENVIRONMENT_VARIABLE}` - Any environment variable available to Jane can be substituted

`${0}` - This is the matched value for commands

`${1}` - This is the first token after the matched command

`${2} ... ${n}` - This is the second token after the matched command (repeats for as many tokens as are after)

`${*}` - This is all tokens after the matched command


### Security

The way of securing who can execute actions via Jane is by setting an optional list of users who are allowed to run commands on connectors that implement commands.

#### Example:

```
"Users": "matt,ken,joe"
```

#### Usage:
* This only applies to connectors that implement commands which users can execute
* The list of users is comma delimited
* The user name is dependant on the connector type, you can run "jane whoami" to get your name


## Architecture

Jane makes heavy use of the Go thread and channel features. Each connector can implement one of the three phases of the Jane bot messaging - Listeners, Commands and Publishers

### Listeners

Listeners are implemented to be long running tasks that take input externally. Examples of this are listening to an RSS feed or connecting to the Slack API. When a listener gets an event it is interested in, it creates a new message and passes it through the command messaging channel where it can be further acted on. Each message will be processed by the commands except when the message's process flag is set to false, in which case it passes through to the Publishers.

### Commands

Commands are implemented to act upon messages. They can do any task based on the message, make changes to the message and pass the message through to the Publishers. It is worth noting that each message gets potentially processed by all Commands, but it is up to the Commands to decide if it should act or ignore. Once the command is complete, it will pass the message to the publisher messaging channel where it will be processed by the Publishers.

### Publishers

Publishers are a means of communicating back out to the world. A publisher will take the message handed to it, format it, and send it through its implemented publish method.


## Getting Involved

### Development Environment
* Get your toes wet with Go
* Setup your Go 1.5.3 environment
* Pull the project with 'go get github.com/projectjane/jane'
* Compile with 'go install jane.go'
* Use the sample _jane.json_ file checked in as a starting point
* Run your code and config with `go run jane.go --config ~/jane.json`
