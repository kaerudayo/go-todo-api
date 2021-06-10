
FROM golang:1.16.3-alpine
RUN apk update

ENV GOBIN=/go/bin
ENV GO111MODULE=on
ENV GOPATH=

RUN mkdir /go/app
ADD . /go/app

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/cespare/reflex

CMD reflex -r '(\.go$|go\.mod)' -s go run app/main.go
