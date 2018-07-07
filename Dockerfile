FROM golang:1.10.3-alpine

LABEL maintainer "aimof <aimof.aimof@gmail.com>"

RUN apk update --no-cache && \
    apk add --no-cache git && \
    go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/aimof/discat/
WORKDIR /go/src/github.com/aimof/discat/cmd/discat/

RUN dep ensure && \
    go build

CMD ./discat
