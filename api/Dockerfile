FROM golang:latest

# Copy the local package files to the container's
ADD . /go/src/github.com/luuminhthai/rest-api-golang/api

# Install our dependencies
RUN go get github.com/go-martini/martini
RUN go get github.com/martini-contrib/binding
RUN go get github.com/martini-contrib/render
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson

# Install api binary globally within container
RUN go install github.com/luuminhthai/rest-api-golang/api

# Set binary as entrypoint 
ENTRYPOINT /go/bin/api

# Expose default port (3000)
EXPOSE 3000