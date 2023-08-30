default: build

build:
	go build -v ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w -e .

test:
	go test -v -cover -timeout=120s -parallel=4 ./...

submodules:
	@git submodule sync
	@git submodule update --init --recursive
	@git config core.hooksPath githooks
	@git config submodule.recurse true

.PHONY: build lint fmt test submodules
