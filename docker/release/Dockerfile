FROM golang:1.10-alpine3.7

WORKDIR /go/src/github.com/mvisonneau/marefo

RUN \
apk add --no-cache make git binutils

COPY Makefile .
RUN \
make setup

COPY . .
RUN \
make deps ;\
make build
