check: lint test

fmt:
	go fmt ./...

fmtfix:
	golangci-lint run --fix -E gofmt,gofumpt,goimports

lint:
	golangci-lint run

test:
	go test ./...