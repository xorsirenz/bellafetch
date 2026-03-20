SHELL := /bin/sh
PACKAGE= bellafetch
SOURCES := $(wildcard *.go cmd/*.go)
VERSION=$(shell git describe --tags --long 2>/dev/null)

TAGS=netgo,osusergo
LDFLAGS="-extldflags '-static' -s -w -X main.version=${VERSION}"

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

FILE_NAME="${PACKAGE}-${VERSION}-${GOOS}-${GOARCH}"

ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

.PHONY: all build clean

all: build

build: $(SOURCES)
	CGO_ENABLED=0 go build -a -tags ${TAGS} -ldflags ${LDFLAGS} -o ${PACKAGE} ./cmd

clean:
	rm -f ${PACKAGE}
