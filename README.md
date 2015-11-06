# Jane

Jane is a bot to pull information and conduct operational activities in your chatops scenario - even in a command line way. This bot is written in go and is made to be configuration driven. Contributions are welcome via pull requests. If you want to know why the name 'Jane' was chosen, talk to @kcwinner.



## Getting Started
* This is developed using Go 1.5.1
* Pull the project with 'go get github.com/mmcquillan/jane'
* Compile with 'go install jane.go'
* Use the sample.config for an example configuration file
* Use the startup.conf as an upstart script to start/stop/restart on Linux systems
* You can rename your bot by setting the top-level Name configuration setting



## Configuration
The entire configuration of the site is done via a json config file. The configuration file is expected to be named 'jane.config' and will be looked for in this order:
* ./jane.config - the location of the jane binary
* ~/jane.config - the home directory of the user
* /etc/jane.config - the global config



## Listeners
Listeners are what Jane uses to pull in information and listen for commands. The Relays specify where the results from the input should be written to or * for all. The Target can specify a channel in the case of slack.

### Command Line listener
`{"Type": "cli", "Name": "cli", "Active": false,
 "Destinations": [{"Match": "*", "Relays": "cli", "Target": ""}]}`

### Slack Listener
`{"Type": "slack", "Name": "slack", "Active": true,
    "Key": "<SlackToken>",
    "Destinations": [
      {"Match": "*", "Relays": "slack", "Target": ""}
    ]
  },`

### RSS Listener
`{"Type": "rss", "Name": "Bamboo Build", "Active": true,
    "Server": "https://BambooUser:BambooPass@somecompany.atlassian.net/builds/plugins/servlet/streams?local=true",
    "SuccessMatch": "successful", "FailureMatch": "fail",
    "Destinations": [
      {"Match": "*", "Relays": "slack", "Target": "#devops"},
      {"Match": "NextGen", "Relays": "slack", "Target": "#nextgen"}
    ]
  },

  {"Type": "rss", "Name": "AWS EC2", "Active": true,
    "Server": "http://status.aws.amazon.com/rss/ec2-us-east-1.rss",
    "Destinations": [
      {"Match": "*", "Relays": "slack", "Target": "#devops"}
    ]
  }`

### Monitor Listener
Note, this is currently setup to execute a nagios style monitoring script and interpret the results as the example shows below.

`{"Type": "monitor", "Name": "Elasticsearch Node", "Active": true,
    "Server": "elasticsearch1.somecompany.com", "Login": "jane", "Pass": "abc123",
    "SuccessMatch": "OK", "WarningMatch": "WARNING", "FailureMatch": "CRITICAL",
    "Checks": [
      {"Name": "Apt Check", "Check": "/usr/lib/nagios/plugins/check_apt"},
      {"Name": "Disk Check", "Check": "/usr/lib/nagios/plugins/check_disk -w10% -c5% -A"},
      {"Name": "Elasticsearch Check", "Check": "/usr/lib/nagios/plugins/check_procs -a elasticsearch -c1:1"}
    ],
    "Destinations": [
      {"Match": "*", "Relays": "slack", "Target": "#devops"},
      {"Match": "*", "Relays": "slack", "Target": "@matt"}
    ]
  }`



## Commands
Commands are what execute or respond to requests by listeners.

### Response Command
`{"Type": "response", "Match": "motivate", "Output": "You can _do it_ %msg%!"}`

### Exec Command
`{"Type": "exec", "Match": "big", "Output": "```%stdout%```", "Cmd": "/usr/bin/figlet", "Args": "-w80 %msg%"}`

### Help Command
`{"Type": "help", "Match": "help"}`

### Reload Command
`{"Type": "reload", "Match": "reload", "Output": "Reloading command configuration"}`



## Relays
Relays are how information gets communicated out from Jane.

### Slack Relay
`{"Type": "slack", "Image": ":speech_balloon:", "Resource": "xxxSlackTokenxxx", "Active": true}`

### Command Line Relay
`{"Type": "cli", "Image": "", "Resource": "", "Active": false}`

