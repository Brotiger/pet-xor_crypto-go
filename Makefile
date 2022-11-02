.PHONY: build
build:
	go build -v ./cmd/xor

.DEFAULT_GOAL := build