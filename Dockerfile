FROM golang:1.13-alpine AS build-context

RUN apk update && \
    apk add build-base bind=9.14.8-r5 git

COPY . /go/src/dns_api

WORKDIR /go/src/dns_api

RUN go mod tidy && go build

# PRODUCTION IMAGE

FROM alpine:3.11

RUN apk update && \
    apk add bind=9.14.8-r5

COPY config/named.conf /etc/bind/named.conf
COPY config/myzone.com.zone /etc/bind/zone/myzone.com.zone

COPY --from=build-context /go/src/dns_api/dns_api .

CMD named && /dns_api
