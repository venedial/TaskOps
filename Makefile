.PHONY: all
all: build

.PHONY: build
build:
	go build ./...

.PHONY: local
local:
	ENV=local air

.PHONY: clean
clean:
	go clean
