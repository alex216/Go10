GO := go
SOURCES := btreedeletenode.go

all: build

test:
	$(GO) test -v --cover ./...

build: $(SOURCES)
	$(GO) build -v

clean:
	rm -rf ex10

.PHONY: all test build clean
