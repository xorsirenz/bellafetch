SHELL := /bin/sh
PACKAGE= bellafetch
SOURCES := $(wildcard *.go cmd/*.go)
VERSION ?= $(shell git describe --tags --long --dirty 2>/dev/null)

TAGS=netgo,osusergo
LDFLAGS_BASE=-extldflags '-static' -s -w -X main.version=${VERSION}
EXTRA_LDFLAGS ?=
LDFLAGS="${LDFLAGS_BASE} ${EXTRA_LDFLAGS}"

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

FILE_NAME="${PACKAGE}-${VERSION}-${GOOS}-${GOARCH}"

ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

.PHONY: all build clean

all: build

build: $(SOURCES)
	CGO_ENABLED=0 go build -a -tags ${TAGS} -ldflags ${LDFLAGS} -trimpath -o ${PACKAGE} ./cmd/bellafetch/

clean:
	rm -f ${PACKAGE}
