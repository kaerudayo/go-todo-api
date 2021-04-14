FROM golang:1.16.2-alpine
RUN apk update && apk add git

ENV GO111MODULE=on

RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
