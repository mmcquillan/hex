# Run as Systemd

This covers how Hex is setup to run as a systemd service.

1. Create a Hex user `sudo useradd -r -s /bin/false hex`
2. Download the latest Hex binary from https://hexbot.io/downloads/
3. Place the `hex` binary in `/usr/local/bin`
4. Create a directory for the Hex config `sudo mkdir /etc/hex`
5. Create a directory for the Hex rules `sudo mkdir /etc/hex/rules`
6. Create a directory for the Hex plugins `sudo mkdir /etc/hex/plugins`
7. Create a log file `sudo touch /var/log/hex.log`
8. Change ownership for the log file `sudo chown hex:hex /var/log/hex.log`
9. Create a Hex config file such as `/etc/hex/config.json`:
```
{
  "slack": true,
  "slack_token": "<INSERT YOUR SLACK TOKEN>",
  "log_file": "/var/log/hex.log",
  "rules_dir": "/etc/hex/rules",
  "plugins_dir": "/etc/hex/plugins"
}
```
10. Create a systemd service file `/etc/systemd/system/hex.service`:
```
[Unit]
Description=HexBot

[Service]
User=hex
Group=hex
Type=simple
ExecStart=/usr/local/bin/hex /etc/hex/config.json
Restart=on-failure

[Install]
WantedBy=multi-user.target
```
11. Register the service file `sudo systemctl enable hex.service`
12. Create your first Hex rule `/etc/hex/rules/hello.json`:
```
{
  "rule": "hello",
  "match": "hello",
  "actions": [
    { "type": "hex-response", "command": "Hello ${hex.user}!" }
  ]
}
```
13. Change the ownership of all Hex configuration `sudo chown -R hex:hex /etc/hex`
14. Startup the Hex service `sudo systemctl start hex`
15. Check in Slack if the Hex bot shows and ask `@hex hello` and it should respond.
