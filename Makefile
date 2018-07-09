VERSION ?= latest # Should be git tag later on

IMAGE = radityasurya/pw-file
PKG = github.com/radityasurya/pw-file

LDFLAGS = "-s -w -X radityasurya/pw-file/pkg/version.Version=$(VERSION)"

OS ?= linux
ARCH ?= amd64

test:
	go test -race ./...

generate:
	go list ./... |grep -v vendor |xargs go generate

build:
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -installsuffix cgo -o bin/druid-task-producer -a -tags netgo -ldflags $(LDFLAGS) ./cmd

all: generate test build

image:
	docker build -t $(IMAGE):$(VERSION) .