GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BUILD_DATE:=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT:=$(shell git rev-parse HEAD)
VERSION_MAJOR:=$(shell cat main.go | awk '/versionMajor = / {print $$3}')
VERSION_MINOR:=$(shell cat main.go | awk '/versionMinor = / {print $$3}')
VERSION_PATCH:=$(shell cat main.go | awk '/versionPatch = / {print $$3}')
VERSION:=$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)
LFLAGS=-ldflags "-X main.buildDate=${BUILD_DATE} -X main.commit=${COMMIT}"

DOCKER_IMAGE=greenstatic/fri-restavracija123-slack-bot

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)

.PHONY: docker
docker:
	docker build -t $(DOCKER_IMAGE):$(VERSION) -t $(DOCKER_IMAGE):latest -f ./Dockerfile .

.PHONY: docker-build
docker-build:
	GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 $(GOBUILD) -a -installsuffix cgo -o bot_amd64 $(LFLAGS) .
