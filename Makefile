SHELL=/bin/bash

.PHONY: dep-ensure format lint test

dep-ensure:
	dep ensure

format:
	go fmt $$(go list ./... | grep -v /vendor/)

lint: dep-ensure
	go vet ./... && \
		gometalinter.v2 ./...

run-example: dep-ensure
	go run main.go

test:
	go test -v -count=1 ./...
