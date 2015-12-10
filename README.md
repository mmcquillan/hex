# Jane

Jane is a bot to pull information and conduct operational activities in your chatops scenario - even in a command line way. This bot is written in go and is made to be configuration driven. Contributions are welcome via pull requests. If you want to know why the name 'Jane' was chosen, talk to @kcwinner.



## Getting Started
* This is developed using Go 1.5.1
* Pull the project with 'go get github.com/projectjane/jane'
* Compile with 'go install jane.go'
* Use the samples in the startup folder for different environments


## Configuration
The entire configuration of the site is done via a json config file. The configuration file is expected to be named 'jane.config' and will be looked for in this order:
* -config config.json - Pass in a configuration file location as a command line parameter
* ./jane.json - the location of the jane binary
* ~/jane.json - the home directory of the user
* /etc/jane.json - the global config


## Connectors
Connectors are what Jane uses to pull in information, interpret them and issue out a response. The Routes specify where the results from the input should be written to or * for all. The Target can specify a channel in the case of slack. To add a new connector, Put them in the connectors folder and make an entry in connectors/list.go.

For the connector configuration, when adding routes, you must specify the ID of the connector you want to route response to.

Supported connectors:
* cli - Command line interface
* bamboo - Atlassian Bamboo integration
* email - Email
* exec - Execution of applications
* imageme - Pull back images or animated gifs
* jira - Atlassian Jira integration
* monitor - Monitor of systems
* slack - Slack chat
* response - Text Responses
* rss - RSS Feed
* website - Monitor return code of websites
* wolfram - Execute queries against Wolfram Alpha
