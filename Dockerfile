FROM golang:1.5.1

ADD . /go/src/github.com/DreamItGetIT/sqs-initialiser
WORKDIR /go/src/github.com/DreamItGetIT/sqs-initialiser
ENV GO15VENDOREXPERIMENT=1
RUN mkdir /tools && go build -o /tools/sqsinit create_queues.go

ENTRYPOINT ["/tools/sqsinit"]
