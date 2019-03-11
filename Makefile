SHELL=/bin/bash

.PHONY: dep-ensure format lint test run-example fail-prev pull-request

dep-ensure:
	dep ensure

format:
	go fmt $$(go list ./... | grep -v /vendor/)

lint: dep-ensure
	go vet ./... && \
		gometalinter.v2 ./...

test:
	go test -v -count=1 ./...

run-example: dep-ensure
	go run main.go

fail-prev:
	status=$$(git status --porcelain); \
	if ! test "x$${status}" = x; then \
		echo >&2 "Unexpected file changes detected"; \
		false; \
	fi

pull-request: dep-ensure fail-prev format fail-prev lint test
