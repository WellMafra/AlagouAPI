FROM golang:1.9

ARG app_env
ENV APP_ENV $app_env

COPY ./ /go/src/github.com/Alagou/AlagouAPI
WORKDIR /go/src/github.com/Alagou/AlagouAPI
ENV GOPATH /go

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = development ]; \
	then \
	go get github.com/pilu/fresh && \
	fresh; \
	else \
	ls -la; \
	AlagouAPI; \
	fi
