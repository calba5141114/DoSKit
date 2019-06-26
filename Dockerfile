# project base image (needs newer version)
FROM golang:1.11

# Maintain and Workdir
LABEL maintainer="Carlos Alba <soradev4@gmail.com>"
WORKDIR $GOPATH/src/github.com/calba5141114/DoSKit

# copy source from local host to Docker host
COPY . .

# get all dependencies