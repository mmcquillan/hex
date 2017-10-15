# Hex

Hex is a chatops bot written in Go and is completely configuration driven. Contributions are welcome via pull requests. Hex was started as a way of getting DevOps tasks and feedback into Slack. There are a billion other bots, but we wanted to learn Go, so this was a fun way to learn it and meet our needs. 

http://hexbot.io


## Configuration

Hex can run from environment variables, command line options or a configuration file. If you use a configuration file, use that as a the first argument to execute hex, such as `/usr/local/bin/hex /etc/hex/config.json`. Below is a sample config file:

```
{
  "slack": true,
  "slack_token": "${SLACK_TOKEN}",
  "scheduler": true,
  "log_file": "/var/log/hex.log",
  "rules_dir": "/etc/hex/rules",
  "plugins_dir": "/etc/hex/plugins"
}
```

Other config options are lited below:

### Admins
- Description: Comma delimited list of admins (users or channels)
- Default:
- Type: string
- Env Variable: `HEX_ADMINS`
- CLI Param:  `-admins`
- Conf File: `admins`

### Plugins Directory
- Description: The location of the hex plugins
- Default: `/etc/hex/plugins`
- Type: string
- Env Variable: `HEX_PLUGINS_DIR`
- CLI Param:  `-plugins-dir`
- Conf File: `plugins_dir`

### Rules Directory
- Description: The location of the hex rules
- Default: `/etc/hex/rules`
- Type: string
- Env Variable: `HEX_RULES_DIR`
- CLI Param:  `-rules-dir`
- Conf File: `rules_dir`

### Log File
- Description: The logfile for the hexbot (empty is stdout)
- Default: 
- Type: string
- Env Variable: `HEX_LOG_FILE`
- CLI Param:  `-log-file`
- Conf File: `log_file`

### Debug
- Description: Flag to enable debug for logs
- Default: `false`
- Type: bool
- Env Variable: `HEX_DEBUG`
- CLI Param:  `-debug`
- Conf File: `debug`

### Bot Name
- Description: Bot Name in Slack
- Default: `@hex`
- Type: string
- Env Variable: `HEX_BOT_NAME`
- CLI Param:  `-bot-name`
- Conf File: `bot_name`

### Command Line Interface
- Description: Flag to enable the command line interface feature (for dev use)
- Default: `false`
- Type: bool
- Env Variable: `HEX_CLI`
- CLI Param:  `-cli`
- Conf File: `cli`

### Auditing
- Description: Flag to enable the auditing feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_AUDITING`
- CLI Param:  `-auditing`
- Conf File: `auditing`

### Auditing File
- Description: The location of the auditing log file
- Default:
- Type: string
- Env Variable: `HEX_AUDITING_FILE`
- CLI Param:  `-auditing-file`
- Conf File: `auditing_file`

### Slack
- Description: Flag to enable the slack feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_SLACK`
- CLI Param:  `-slack`
- Conf File: `slack`

### Slack Token
- Description: Slack token (created in slack)
- Default: 
- Type: string
- Env Variable: `HEX_SLACK_TOKEN`
- CLI Param:  `-slack-token`
- Conf File: `slack_token`

### Slack Icon
- Description: Slack icon to use for the bot
- Default: `:nut_and_bolt:`
- Type: string
- Env Variable: `HEX_SLACK_ICON`
- CLI Param:  `-slack-icon`
- Conf File: `slack_icon`

### Slack Debug
- Description: Flag to enable slack library level debug (for dev use)
- Default: `false`
- Type: bool
- Env Variable: `HEX_SLACK_DEBUG`
- CLI Param:  `-slack-debug`
- Conf File: `slack_debug`

### Scheduler
- Description: Flag to enable the scheduler feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_SCHEDULER`
- CLI Param:  `-scheduler`
- Conf File: `scheduler`

### Webhook
- Description: Flag to enable the webhook feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_WEBHOOK`
- CLI Param:  `-webhook`
- Conf File: `webhook`

### Webhook Port
- Description: The port to listen on for webhook calls
- Default: `8000`
- Type: int
- Env Variable: `HEX_WEBHOOK_PORT`
- CLI Param:  `-webhook-port`
- Conf File: `webhook_port`


## Rules

Each rule is a seperate json file which consists of some rule options and a series of actions to execute if the rule matches. A rule will generally be either a match from slack input, a schedule or a web url. Below is a sample rule:

```
{
  "rule": "Say Hello",
  "match": "hello",
  "actions": [
    {
      "type": "hex-response",
      "command": "Hello ${hex.user}!"
    }
  ]
}
```

### Name
- Config: `name`
- Description: Name of the rule, also used as the title for formatted output
- Default:
- Type: string

### Match
- Config: `match`
- Description: String to match with * as wild card or /../ as regular expression
- Default:
- Type: string

### Schedule
- Config: `schedule`
- Description: Cron style schedule with seconds
- Default:
- Type: string

### URL
- Config: `url`
- Description: URL to match for incoming webhooks
- Default:
- Type: string

### ACL
- Config: `acl`
- Description: A comma delimited list of users and channels allowed to execute the rule
- Default: *
- Type: string

### Channel
- Config: `channel`
- Description: A channel to send output to if a schedule or if you want a place for all output to be copied to
- Default:
- Type: string

### Format
- Config: `format`
- Description: Flag to format output when displaying in slack
- Default: false
- Type: bool

### Help
- Config: `help`
- Description: Custom help to display for the rule
- Default:
- Type: string

### Hide
- Config: `hide`
- Description: Flag for displaying help or not when user lists commands
- Default: false
- Type: bool

### Active
- Config: `active`
- Description: Flag for if the rule is run or not
- Default: true
- Type: bool

### Debug
- Config: `debug`
- Description: Flag for extra debug output in the logs
- Default: false
- Type: bool

### Actions.Type
- Config: `type`
- Description: The type of action, also known as the plugin to execute
- Default:
- Type: string

### Actions.Command
- Config: `command`
- Description: The command to give the plugin to resolve
- Default:
- Type: string

### Actions.OutputToVar
- Config: `output_to_var`
- Description: An option to take the output of the action and save it to the `hex.outputs.<action number>.response` var
- Default: false
- Type: bool

### Actions.RunOnFail
- Config: `run_on_fail`
- Description: A flag to let this action run if previous steps have failed
- Default: false
- Type: bool

### Actions.LastConfig
- Config: `last_config`
- Description: A flag to reuse the previous actions configuration (as a time saver)
- Default: false
- Type: bool

### Actions.Config
- Config: `config`
- Description: A set of key/value string pairs that can be plugin specific
- Default:
- Type: string key/value

