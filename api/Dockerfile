FROM golang:latest

# Copy the local package files to the container's
ADD . /go/src/github.com/luuminhthai/rest-api-golang/api

# Install our dependencies
RUN go get github.com/kataras/iris
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson

# Setting up working directory
WORKDIR /go/src/github.com/luuminhthai/rest-api-golang/api

# Install api binary globally within container
RUN go install github.com/luuminhthai/rest-api-golang/api

# Set binary as entrypoint 
ENTRYPOINT /go/bin/api

# Expose default port (3000)
EXPOSE 8080