.PHONY: get_deps fmt
.DEFAULT_GOAL := build
tests: lint test

test:
	@echo MARK: unit tests
	go test -v $(shell go list ./... | grep -v /tmp | grep -v /tests) -timeout 60s -race -covermode=atomic -coverprofile=coverage.out

install_linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2

lint-todo:
	@echo MARK: make lint todo

lint:
	golangci-lint run

get_deps:
	go mod tidy

fmt:
	@gofmt -l -w -s .
	@goimports -w .

comments:
	@echo MARK: update comments
	@gocmt -i -d .

