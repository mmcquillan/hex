# Hex

Hex is a chatops bot written in Go and is completely configuration driven. Contributions are welcome via pull requests. Hex was started as a way of getting DevOps tasks and feedback into Slack. There are a billion other bots, but we wanted to learn Go, so this was a fun way to learn it and meet our needs. 

http://hexbot.io


## Configuration

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

### Workspace Directory
- Description: The location of the workspace for each rule
- Default: `/tmp`
- Type: string
- Env Variable: `HEX_WORKSPACE_DIR`
- CLI Param:  `-workspace-dir`
- Conf File: `workspace_dir`

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



