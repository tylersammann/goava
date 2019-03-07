SHELL=/bin/bash

.PHONY: dep-ensure format lint test

dep-ensure:
	dep ensure

format:
	go fmt $$(go list ./... | grep -v /vendor/)

lint: dep-ensure
	go vet ./... && \
	golint -set_exit_status `go list ./... | grep -v /vendor/`

run-example: dep-ensure
	go run main.go

test:
	go test -v -count=1 ./...
