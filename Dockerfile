FROM golang:1.9.1
MAINTAINER Matt McQuillan
ENV GOBIN=/go/bin
RUN go get -d github.com/mmcquillan/hex
RUN go install github.com/mmcquillan/hex
ENTRYPOINT /go/bin/hex
