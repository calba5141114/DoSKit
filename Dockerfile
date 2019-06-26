# project base image (needs newer version)
FROM golang:1.11

# Maintain and Workdir
LABEL maintainer="Carlos Alba <soradev4@gmail.com>"
WORKDIR $GOPATH/src/github.com/calba5141114/DoSKit

# copy source from local host to Docker host
COPY . .

# get all dependencies
RUN go get -d -v ./...

# install binary
RUN go install -v ./...

EXPOSE 8080

CMD ["DosKit"]