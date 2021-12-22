
.PHONY: build test start dev watch debug
build:
	CGO_ENABLED=0 go build -o bin/sub main.go

test:
	go test ./...
	
start:
	bin/sub

dev: build start

watch:
	air

debug:
	dlv debug --headless --listen=:2345 --log --api-version=2