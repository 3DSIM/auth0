FROM golang:1.10

ENV APPLICATION_NAME=auth0

# Install dep
RUN go get -u github.com/golang/dep/cmd/dep

# Install application
RUN mkdir -p $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
COPY . $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
WORKDIR $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
RUN dep ensure -v

# Run tests
RUN go test ./... -cover
