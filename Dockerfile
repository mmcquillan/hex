# hex docker file
FROM golang:1.8.1
MAINTAINER Project Hex
ENV GOBIN=/go/bin
ADD . /go/src/github.com/hexbotio/hex
WORKDIR /go/src/github.com/hexbotio/hex
RUN go get github.com/hexbotio/hex \
    && go install -ldflags "-X main.version=$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse HEAD | cut -c1-8)" hex.go
ENTRYPOINT /go/bin/hex
