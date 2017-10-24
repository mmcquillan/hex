# Monitoring a Server

Hex will allow you to do some very basic server monitoring. Here we have an example of how Hex isused to alert if there is a critical update for our server.

```
{
  "rule": "apt-check",
  "match": "apt-check",
  "format": true,
  "output_fail_only": true,
  "schedule": "0 0 */1 * * *",
  "channel": "#alerts",
  "actions": [
    { "type": "hex-local", "command": "/usr/lib/nagios/plugins/check_apt" }
  ]
}
```
