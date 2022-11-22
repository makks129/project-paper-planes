.PHONY: update
update:
	go get -u all && go mod vendor

.PHONY: test-setup
test-setup:
	@go run ./test/functional/main.go -phase=setup

.PHONY: test-teardown
test-teardown:
	@go run ./test/functional/main.go -phase=teardown

.PHONY: just-tests
test-only:
	@GO_ENV=test go test -run Test -v ./test/... -p 1

.PHONY: test
test:
	@make test-setup
	@make test-only
	@make test-teardown

# TODO fix coverage (go tool cover -func=coverage.out)
.PHONY: coverage
coverage:
	@make test-setup
	@GO_ENV=test go test -run Test -v ./test/... -p 1 -coverprofile=coverage.out
	@make test-teardown

.PHONY: lint
lint:
	@echo "Running lint..."
	@golangci-lint run --enable-all

.PHONY: install
install: mkdirs installlint

.PHONY: mkdirs
mkdirs:
	@mkdir ./.bin 2>/dev/null || true

.PHONY: installlint
installlint: mkdirs
	@echo "Installing linter..." || true
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./.bin "${LINT_VERSION}" >/dev/null 2>&1 || true
