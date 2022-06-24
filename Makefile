PROJECT_PATH=${PWD}

.PHONY: generate
generate:
	echo ${PROJECT_PATH}
	go run generate/go-generate.go ${PROJECT_PATH}

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	docker run --rm -v ${PROJECT_PATH}:/app -w /app ghcr.io/haproxytech/go-linter:1.46.2 -v --timeout 5m --color always --max-issues-per-linter 0 --max-same-issues 0

.PHONY: lint-local
lint-local:
	golangci-lint run
