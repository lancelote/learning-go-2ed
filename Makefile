.DEFAULT_GOAL := vet
.PHONY:fmt vet lint test cov

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

cov:
	go test -v -cover -coverprofile=c.out ./...
	go tool cover -html=c.out
