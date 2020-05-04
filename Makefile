GO = GO111MODULE=on go

## display this help message
help:
	@echo ''
	@echo 'make targets for the project:'
	@echo
	@echo 'Usage:'
	@echo '  ## Develop / Test Commands'
	@echo '  all             Run all the commands.'
	@echo '  clean           Run clean up.'
	@echo '  setup           Run setup of tools needed. This should only be executed once.'
	@echo '  fmt             Run code formatter.'
	@echo '  check           Run static code analysis (lint).'
	@echo '  test            Run tests on project.'
	@echo '  deps            Run all dependencies that are needed.'
	@echo ''

all: clean deps fmt check test

clean:
	@echo "==> Cleaning..."
	rm -f coverage.out
	rm -f report.json

deps:
	@echo "==> Getting Dependencies..."
	${GO} mod tidy
	${GO} mod download

test:
	@echo "==> Testing..."
	CGO_ENABLED=0 ${GO} test -v -covermode=atomic -count=1 ./... -coverprofile coverage.out
	${GO} test -v -race -covermode=atomic -count=1 ./...  -json > report.json
	${GO} tool cover -func=coverage.out

fmt:
	@echo "==> Code Formatting..."
	${GO} fmt ./...

check: fmt
	@echo "==> Code Check..."
	golangci-lint run --fast --tests

clean_cache:
	@${GO} clean -cache -modcache -i -r -x

## run this to install the needed tools in this make file (for the lint and check targets).
## You should only need to do this once.
setup:
	@echo "==> Setup..."
	${GO} get -u github.com/golangci/golangci-lint/cmd/golangci-lint

