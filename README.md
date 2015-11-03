# Jane


## Getting Started
* This is developed using Go 1.5.1
* Pull the project with `go get github.com/mmcquillan/jane`
* Use the sample.config for an example configuration file


## Configuration
The entire configuration of the site is done via a json configuraiton file. The configuration file is expected to be named 'jane.config' and will be looked for in this order:
* ./jane.config - the location of the jane binary
* ~/jane.config - the home directory of the user
* /etc/jane.config - the global config


## Listeners
Listeners are what Jane uses to pull in information and listen for commands. The Relays specify where the results from the input should be written to or * for all. The Target can specify a channel in the case of slack.

### Command Line listener
`{"Type": "cli", "Name": "cli", "Relays": "cli", "Active": false}`

### Slack Listener
`{"Type": "slack", "Name": "slack", "Resource": "xxxSlackTokenxxx", "Relays": "slack", "Active": true }`

### RSS Listener
`{"Type": "rss", "Name": "AWS EC2", "Resource": "http://status.aws.amazon.com/rss/ec2-us-east-1.rss", "Relays": "*", "Target": "#devops", "SuccessMatch": "", "FailureMatch": "", "Active": true }`

### Monitor Listener
`{"Type": "monitor", "Name": "Prod Elasticsearch", "Resource": "user:password@prod.server.com|/usr/lib/nagios/plugins/check_procs -C elasticsearch -c1:1", "Relays": "*", "Target": "#devops", "Active": true}`


## Commands
Commands are what execute or respond to requests by listeners.


## Relays
Relays are how information gets communicated out from Jane.

