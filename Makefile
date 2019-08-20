# This adds versioning by using the most recent commit 
# Adds the versioning into the TAG shell variable
# All this file is is a bunch of bash commands
# $(GOOS) $(GOBUILD) -ldflags "-X main.version=$(TAG)" -o $(BINARY_NAME) .
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

#More parameters
BINARY_NAME=main
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
#Turns out you need to compile the binary with lunx architecture otherwise it will fail on the docker image running ubuntu
GOOS=GOOS=linux

test:
	#TODO add some tests smh

build: 
		$(GOOS) $(GOBUILD) -ldflags "-X main.version=$(TAG)" -o $(BINARY_NAME) .

clean: 
		$(GOCLEAN)

run: build
		./$(BINARY_NAME)

pack: build
		docker build -t gcr.io/youtube-data-235117/trbackend:$(TAG) .

upload: pack
		docker push gcr.io/youtube-data-235117/trbackend:$(TAG)

deprun: pack
		docker run  --rm -p 8080:8080 gcr.io/youtube-data-235117/trbackend:$(TAG)

deps:
		#TODO: Add go get all the dependencies