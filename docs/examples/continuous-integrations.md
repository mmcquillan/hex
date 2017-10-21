# Continuous Integrations

This document covers how Hex performs its continuous integrations for the project.

1. Enable the webhook functionality `"webhook": true, "webhook_port": 8080`
2. Ensure your Hex Bot is accessible from the internet `curl http://<your ip address>:8080/ping` which should result in a response like: `{"serviced_by":"HexBot", "message_id":"b7kjt2lmk9b382mh7k50"}`
3. In the Github repo you want to hook up, navigate to Settings > Webhooks > Add webhook
4. Set the Payload URL to `http://<your ip address>:8080/build/hex` (note, you should adjust this to whatever path you wish to use)
5. Set the Content type to `application/json`
6. Likely you just want the push events and will not need a secret at this time
7. Using the path you set in step 4 (`/build/hex`) we will construct a rule to build - this example is what is used for hex:
```
{
  "rule": "Build Hex",
  "match": "build hex",
  "url": "/build/hex",
  "channel": "#builds",
  "format": true,
  "actions": [
    { "type": "hex-local", "command": "env", "hide_output": true, "config": { 
        "env": "GOPATH=/tmp/build/${hex.id}; GOBIN=/tmp/build/${hex.id}/bin; PATH=/bin:/usr/bin:/usr/local/go/bin/",
        "dir": "/tmp/build"
      }
    },
    { "type": "hex-local", "command": "mkdir -p /tmp/build/${hex.id}", "hide_output": true, "last_config": true },
    { "type": "hex-local", "command": "echo Pulling; go get github.com/hexbotio/hex", "last_config": true },
    { "type": "hex-local", "command": "echo Testing; go test github.com/hexbotio/hex/parse", "last_config": true },
    { "type": "hex-local", "command": "echo Building; go build github.com/hexbotio/hex", "last_config": true },
    { "type": "hex-local", "command": "rm -rf /tmp/build/${hex.id}", "hide_output": true, "run_on_fail": true },
    { "type": "hex-response", "command": "Push by: ${hex.input.json:sender.login} ${hex.input.json:head_commit.message}", "run_on_fail": true }
  ]
}
```
8. Test by pushing code to the repo and watch the #builds channel (or whatever you set it to) for the build steps.

