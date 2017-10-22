FROM golang:1.9.1
MAINTAINER HexBot.io
ENV GOBIN=/go/bin
RUN go get -d github.com/hexbotio/hex
RUN go install github.com/hexbotio/hex
ENTRYPOINT /go/bin/hex
