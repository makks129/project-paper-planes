# TODO Setup pre-commit hooks

.PHONY: update
update:
	go get -u all && go mod vendor

.PHONY: test
test:
	@echo "Running tests...\n"
	@GO_ENV=test go test -run Test -v ./test/... -p 1

.PHONY: coverage
coverage:
	@echo "Running test coverage..."
	@GO_ENV=test go test -run Test -v ./test/... -p 1 -coverprofile=coverage.out

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
