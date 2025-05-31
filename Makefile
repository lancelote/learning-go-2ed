.DEFAULT_GOAL := vet
.PHONY:fmt vet lint test

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

lint:
	golangci-lint run
	revive
	staticcheck ./...
	go vet ./...

test:
	go test ./...
