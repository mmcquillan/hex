#!/bin/bash
set -ex
rm -rf ./out
go test ./... -v -parallel 1
mkdir out
export GOBIN="$(pwd)/out"
go install hex.go
mkdir out/plugins
export GOBIN="$GOBIN/plugins"
go install github.com/mmcquillan/hex-local
go install github.com/mmcquillan/hex-response
go install github.com/mmcquillan/hex-ssh
go install github.com/mmcquillan/hex-twilio
go install github.com/mmcquillan/hex-validate
go install github.com/mmcquillan/hex-winrm
mkdir out/rules
cp ./rules/* ./out/rules/
