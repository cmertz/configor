default: fmt test lint

test:
	.ci/go-test 74

fmt:
	.ci/go-fmt

lint:
	golangci-lint run ./...
