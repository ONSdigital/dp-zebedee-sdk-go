.PHONY: test
test:
	go test -count=1 -race -cover ./...

.PHONY: audit
audit:
	go list -json -m all | nancy sleuth

.PHONY: build
build:
	go build ./...

.PHONY: lint
lint: ## Used in ci to run linters against Go code
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2
	golangci-lint run ./...