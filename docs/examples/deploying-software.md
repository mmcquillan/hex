# Deploying Software

With Hex's ability to connect to remote systems, you can easily deploy software by running a series of steps. Taking in arguments as part of the command, you can get more sophisticated by setting target deploy environments. Here is a somple example of a rule that deploys the Hex website.

1. Ensure you have configured the webhook feature.
2. Ensure you have setup a Github webhook as described [here](continuous-integrations.md).
3. Add a rule that reacts to the webhook and deploys the new site:
```
{
  "rule": "Deploy Web",
  "match": "deploy web",
  "url": "/github/hex-web",
  "channel": "#builds",
  "format": true,
  "actions": [
    { "type": "hex-local", "command": "env", "hide_output": true, "config": { 
        "env": "PATH=/bin:/usr/bin",
        "dir": "/var/www/hexbot.io"
      }
    },
    { "type": "hex-local", "command": "echo Deploying; git pull", "last_config": true },
    { "type": "hex-response", "command": "Push by: ${hex.input.json:sender.login}", "run_on_fail": true }
  ]
}
```
