FROM golang:1.9
RUN mkdir -p /go/src/github.com/Alagou/AlagouAPI
WORKDIR /go/src/github.com/Alagou/AlagouAPI
ENV GOPATH /go