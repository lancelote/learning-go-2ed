.DEFAULT_GOAL := vet
.PHONY:fmt vet lint

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

lint:
	golangci-lint run
	revive
	staticcheck ./...
	go vet ./...
