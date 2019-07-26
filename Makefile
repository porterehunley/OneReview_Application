# This adds versioning by using the most recent commit 
# Adds the versioning into the TAG shell variable
# All this file is is a bunch of bash commands
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

#More parameters
BINARY_NAME=main
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test

test:
	#TODO add some tests smh

build: 
		$(GOBUILD) -ldflags "-X main.version=$(TAG)" -o $(BINARY_NAME) .

clean: 
		$(GOCLEAN)

run: build
		./$(BINARY_NAME)

deps:
		#TODO: Add go get all the dependencies