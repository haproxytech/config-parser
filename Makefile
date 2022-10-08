PROJECT_PATH=${PWD}

.PHONY: generate
generate:
	go install mvdan.cc/gofumpt@latest
	echo ${PROJECT_PATH}
	go run generate/go-generate.go ${PROJECT_PATH}
	gofumpt -l -w .

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	docker run --rm -v ${PROJECT_PATH}:/app -w /app ghcr.io/haproxytech/go-linter:1.50.0 -v --timeout 5m --color always --max-issues-per-linter 0 --max-same-issues 0

.PHONY: lint-local
lint-local:
	golangci-lint run
