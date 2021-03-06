# Configuration

Hex can run from environment variables, command line options or a configuration file. If you use a configuration file, use that as a the first argument to execute hex, such as `/usr/local/bin/hex /etc/hex/config.yaml`.

yaml config:
```
---
slack: true
slack_token: "${SLACK_TOKEN}"
scheduler: true
log_file: "/var/log/hex.log"
rules_dir: "/etc/hex/rules"
plugins_dir: "/etc/hex/plugins"
```

json config:
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
- CLI Param:  `--admins`
- Conf File: `admins`

### UserACL
- Description: Comma delimited list of users allowed to execute any rule
- Default: *
- Type: string
- Env Variable: `HEX_USER_ACL`
- CLI Param: `--user-acl`
- Conf File: `user_acl`

### ChannelACL
- Description: Comma delimited list of channels allowed to execute any rule
- Default: *
- Type: string
- Env Variable: `HEX_CHANNEL_ACL`
- CLI Param: `--channel-acl`
- Conf File: `channel_acl`

### Plugins Directory
- Description: The location of the hex plugins
- Default:
- Type: string
- Env Variable: `HEX_PLUGINS_DIR`
- CLI Param:  `--plugins-dir`
- Conf File: `plugins_dir`

### Rules Directory
- Description: The location of the hex rules
- Default:
- Type: string
- Env Variable: `HEX_RULES_DIR`
- CLI Param:  `--rules-dir`
- Conf File: `rules_dir`

### Log File
- Description: The logfile for the hexbot (empty is stdout)
- Default: 
- Type: string
- Env Variable: `HEX_LOG_FILE`
- CLI Param:  `--log-file`
- Conf File: `log_file`

### Log Level
- Description: The log level to write out (error, info, debug, trace)
- Default: `info`
- Type: string
- Env Variable: `HEX_LOG_LEVEL`
- CLI Param:  `--log-level`
- Conf File: `log_level`

### Bot Name
- Description: Bot Name in Slack
- Default: `@hex`
- Type: string
- Env Variable: `HEX_BOT_NAME`
- CLI Param:  `--bot-name`
- Conf File: `bot_name`

### Command Line Interface
- Description: Flag to enable the command line interface feature (for dev use)
- Default: `false`
- Type: bool
- Env Variable: `HEX_CLI`
- CLI Param:  `--cli`
- Conf File: `cli`

### Auditing
- Description: Flag to enable the auditing feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_AUDITING`
- CLI Param:  `--auditing`
- Conf File: `auditing`

### Auditing File
- Description: The location of the auditing log file (empty is stdout)
- Default:
- Type: string
- Env Variable: `HEX_AUDITING_FILE`
- CLI Param:  `--auditing-file`
- Conf File: `auditing_file`

### Slack
- Description: Flag to enable the slack feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_SLACK`
- CLI Param:  `--slack`
- Conf File: `slack`

### Slack Token
- Description: Slack token (created in slack)
- Default: 
- Type: string
- Env Variable: `HEX_SLACK_TOKEN`
- CLI Param:  `--slack-token`
- Conf File: `slack_token`

### Slack Icon
- Description: Slack icon to use for the bot
- Default: `:nut_and_bolt:`
- Type: string
- Env Variable: `HEX_SLACK_ICON`
- CLI Param:  `--slack-icon`
- Conf File: `slack_icon`

### Slack Debug
- Description: Flag to enable slack library level debug (for dev use)
- Default: `false`
- Type: bool
- Env Variable: `HEX_SLACK_DEBUG`
- CLI Param:  `--slack-debug`
- Conf File: `slack_debug`

### Scheduler
- Description: Flag to enable the scheduler feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_SCHEDULER`
- CLI Param:  `--scheduler`
- Conf File: `scheduler`

### Webhook
- Description: Flag to enable the webhook feature
- Default: `false`
- Type: bool
- Env Variable: `HEX_WEBHOOK`
- CLI Param:  `--webhook`
- Conf File: `webhook`

### Webhook Port
- Description: The port to listen on for webhook calls
- Default: `8000`
- Type: int
- Env Variable: `HEX_WEBHOOK_PORT`
- CLI Param:  `--webhook-port`
- Conf File: `webhook_port`

### Command
- Description: Command to execute and quit
- Default:
- Type: string
- Env Variable: `HEX_COMMAND`
- CLI Param:  `--command`
- Conf File: `command`

### Vars
- Description: Vars are a key/value set of vars which can be used through out rules for injecting sensitive data
- Default:
- Type: string key/value
- Note: These will be available as `${hex.var.KEYNAME}`

