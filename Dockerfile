FROM golang:1.11

LABEL maintainer="Carlos Alba <soradev4@gmail.com>"
WORKDIR $GOPATH/src/github.com/calba5141114/DoSKit

COPY . .
