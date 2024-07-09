GO := go
SOURCES := btreedeletenode.go

all: build

test:
	$(GO) test -v ./...

build:
	$(GO) build -v

init:
	$(GO) mod init ex10

clean:
	rm -rf ex10 go.mod go.sum

re: clean init build

.PHONY: all test build clean
