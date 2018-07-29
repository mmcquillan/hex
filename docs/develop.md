# Develop

Get involved with the development of Hex or it's plugins.


### Hex Bot

- Setup your Go 1.9.1 environment
- Use [govendor](https://github.com/kardianos/govendor) for managing go packages
- Pull the project with `go get github.com/mmcquillan/hex`
- Setup a local configuration such as `~/hex/config.json` with [configuration settings](configuration.md)
- Run your instance with `go run $GOPATH/src/github.com/mmcquillan/hex/hex.go ~/hex/config.json`


### Plugins

Plugins are a great way of adding functionality to Hex to increase its capability. A very simple plugin to emulate is the hex-response plugin which just returns formatted text. You can get started with that here:

https://github.com/mmcquillan/hex-response


### Testing

There are unit tests written for some of the parse methods. In additon there is a testing framework to run integration tests against the bot and the various plugins.

https://github.com/mmcquillan/thext

