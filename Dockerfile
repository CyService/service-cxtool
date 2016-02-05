FROM golang:1.5.3

# Minimalistic REST API server for Cytoscape CI
WORKDIR /go

# This is the default GOPATH for this container.
ADD . /go/src/github.com/cytoscape-ci/service-cxtool
WORKDIR /go/src/github.com/cytoscape-ci/service-cxtool

# Install Go dependencies
RUN go get github.com/rs/cors
RUN go get github.com/cytoscape-ci/elsa-client/reg
RUN go get github.com/cytoscape-ci/cxtool

# Build the server for this environment
RUN go build app.go

EXPOSE 3000

# Run it!
CMD ./app -agent http://52.35.61.6:8080/registration
