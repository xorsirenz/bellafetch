SHELL := /bin/sh
PACKAGE= bellafetch
SOURCES := $(wildcard *.go cmd/*.go)
VERSION=$(shell git describe --tags --long 2>/dev/null)
LDFLAGS=-ldflags "-X main.version=${VERSION}"

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
FILE_NAME="${PACKAGE}-${VERSION}-${GOOS}-${GOARCH}"

ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

.PHONY: all build clean

all: build

build: $(SOURCES)	
	go build ${LDFLAGS} -o ${PACKAGE} ./cmd

clean:
	rm -f ${PACKAGE}
