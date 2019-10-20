all: build

build:
	protoc --gogofaster_out=. ./*/proto/*.proto
	go build ./...

install:
	go install ./...