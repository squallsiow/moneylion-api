FROM golang:latest 
ADD . /go/src/github.com/moneylion-api
WORKDIR  /go/src/github.com/moneylion-api

RUN make setenv
RUN make build

ENTRYPOINT ./moneylion-api

