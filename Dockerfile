# jane docker file
FROM golang:1.7.0
MAINTAINER Project Jane
ENV GOBIN=/go/bin
ADD . /go/src/github.com/projectjane/jane
WORKDIR /go/src/github.com/projectjane/jane
RUN go get github.com/projectjane/jane \
    && go install -ldflags "-X main.version=$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse HEAD | cut -c1-8)" jane.go
ENTRYPOINT /go/bin/jane
