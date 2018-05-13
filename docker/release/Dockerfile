FROM golang:1.10-alpine3.7

WORKDIR /go/src/github.com/mvisonneau/marefo

RUN \
apk add --no-cache make git binutils

COPY Makefile Gopkg.lock Gopkg.toml .
RUN \
make setup ;\
make deps

COPY . .
RUN \
make build
