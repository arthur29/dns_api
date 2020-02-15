FROM golang:1.13-alpine

RUN apk update && \
    apk add bind=9.14.8-r5 && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/project
