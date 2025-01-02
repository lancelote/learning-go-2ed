.DEFAULT_GOAL := vet
.PHONY:fmt vet

fmt:
	go fmt ./...

vet: fmt
	go vet ./...
