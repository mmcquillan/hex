# Substitutions

This substitution method is used for constructing commands passed to plugins.

- `${ENVIRONMENT_VARIABLE}` - Environment variables are always available
- `${hex.id}` - A unique identifier for each incoming message
- `${hex.service}` - The type of service a message originates from (cli, scheduler, slack, webhook)
- `${hex.input}` - The raw input of the Message
- `${hex.input.0}` - The first word of the input
- `${hex.input.1}` - The second word of the input
- `${hex.input.2:*}` - The third through last word of the input
- `${hex.input.3:4}` - The fourth through fifth word of the input
- `${hex.input.5|pancakes}` - The sixth word of the input or "pancakes" if not found
- `${hex.input.json:web.api}` - The json value at `{"web": {"api": "some value"} }`
- `${hex.user}` - The user who created the Message
- `${hex.channel}` - The originating channel for a slack input
- `${hex.hostname}` - The hostname for a CLI input
- `${hex.schedule}` - The schedule for a scheduler input
- `${hex.url}` - The URL for a webhook input
- `${hex.ipaddress}` - The originating IP for a webhook input
- `${hex.rule.runid}` - A unique identifier for each matched rule
- `${hex.rule.name}` - The name of the matched rule
- `${hex.rule.format}` - The flag for if a rule is formatted or not
- `${hex.rule.channel}` - The default channel for a rule
- `${hex.output.0.duration}` - The time in seconds for the first action to execute
- `${hex.output.0.response}` - The response for the first action if "output_to_var" is true
- `${hex.output.N.duration}` - The time in seconds for the N'th action to execute
- `${hex.output.N.duration}` - The response for the N'th action if "output_to_var" is true

