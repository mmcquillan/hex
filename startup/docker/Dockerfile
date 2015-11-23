# jane docker file
FROM golang:1.5.1
MAINTAINER Project Jane
ENV GOBIN=/go/bin
RUN go get github.com/projectjane/jane
WORKDIR /go/src/github.com/projectjane/jane
ENTRYPOINT git pull && go install jane.go && /go/bin/jane

