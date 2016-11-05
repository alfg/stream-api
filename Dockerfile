FROM golang:1.5.1

ADD . /go/src/golang-restful-starter
WORKDIR /go/src/golang-restful-starter
RUN go get -d -v
RUN go build
ENTRYPOINT /go/src/golang-restful-starter/golang-restful-starter

EXPOSE 4000
