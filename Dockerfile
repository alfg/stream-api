FROM golang:1.7-alpine

EXPOSE 4000

ADD . /go/src/stream-api

RUN apk add --update ca-certificates git gcc g++ && \
    rm -rf /var/cache/apk/* && \
    cd /go/src/stream-api && \
    go get -d -v ./... && \
    go build -o /usr/bin/stream-api . && \
    apk del git gcc g++ && \
    rm -rf /var/cache/apk/* && \
    rm -rf /go

ENTRYPOINT ["/usr/bin/stream-api"]
