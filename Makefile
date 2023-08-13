PROJECT_PATH=${PWD}
GOLANGCI_LINT_VERSION=1.54.1

.PHONY: generate
generate:
	go install mvdan.cc/gofumpt@latest
	echo ${PROJECT_PATH}
	go run generate/*.go ${PROJECT_PATH}
	gofumpt -l -w .

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	cd bin;GOLANGCI_LINT_VERSION=${GOLANGCI_LINT_VERSION} sh lint-check.sh
	bin/golangci-lint run --timeout 5m --color always --max-issues-per-linter 0 --max-same-issues 0
