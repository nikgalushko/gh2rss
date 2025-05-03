B=$(shell git rev-parse --abbrev-ref HEAD)
BRANCH=$(subst /,-,$(B))
GITREV=$(shell git describe --abbrev=7 --always --tags)
REV=$(GITREV)-$(BRANCH)-$(shell date +%Y%m%d-%H:%M:%S)

all: build

build: info
	- cd app; CGO_ENABLED=0 go build -ldflags "-X main.revision=$(REV)" -o ../target/gh2rss

check:
	- cd app; golangci-lint run --out-format=tab --tests=false ./...

info:
	- @echo "revision $(REV)"

.PHONY: bin info
