GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run
GOBUILDFLAGS=-ldflags="-s -w -X main.Version=$(VERSION)"
PROJECT=circa10a/google-home-aws-news
BINARY=webhook
VERSION=0.2.4

# First target for travis ci
test:
	$(GOCMD) test -v ./...

build:
	$(GOBUILD) $(GOBUILDFLAGS) -o $(BINARY)

run:
	$(GORUN) .

compile:
	GOOS=linux GOARCH=amd64 go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-amd64
	# GOOS=linux GOARCH=arm go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-arm
	# GOOS=linux GOARCH=arm64 go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-arm64
	# GOOS=darwin GOARCH=amd64 go build $(GOBUILDFLAGS) -o bin/$(BINARY)-darwin-amd64

clean:
	$(GOCLEAN)
	rm -f $(BINARY)
