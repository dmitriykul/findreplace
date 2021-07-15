
export RELEASE_VERSION=1.1.9

GIT_COMMIT := $(shell git rev-parse HEAD)

export APP_VERSION=$(RELEASE_VERSION)
export APP_REVISION=$(GIT_COMMIT)

DOCKER_IMAGE_I18N=harbor.ispring.lan/infrastructure/findreplace:$(RELEASE_VERSION)

GOPATH := $(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

all: build test check

.PHONY: modules
modules:
	go mod tidy

clean:
	rm -f bin/findreplace
	rm -f bin/findreplacetests

.PHONY: build
build: modules
	bin/run-go-build cmd/findreplace bin/findreplace
	bin/run-go-build cmd/findreplacetests bin/findreplacetests

.PHONY: test
test:
	go test ./...
	bin/findreplacetests

.PHONY: check
check:
	golangci-lint run ./... --config .golangci.yml

.PHONY: prerelease_findreplace
prerelease_findreplace: build test check
	docker build -t "$(DOCKER_IMAGE_I18N)" -f docker/findreplace.Dockerfile .

.PHONY: release_findreplace
release_findreplace: prerelease_findreplace
	docker push "$(DOCKER_IMAGE_I18N)"	
