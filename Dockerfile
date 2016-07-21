# jane docker file
FROM golang:1.5.4
MAINTAINER Project Jane
ENV GOBIN=/go/bin
ADD . /go/src/github.com/projectjane/jane
WORKDIR /go/src/github.com/projectjane/jane
RUN go get -u github.com/projectjane/jane && go install jane.go
ENTRYPOINT /go/bin/jane
