SOURCE_FILES?=./...
GOLANGCI_VERSION=v1.59.0 # Also update golangci-lint GH action in pr.yml when updating this version
GOIMPORTS_VERSION=v0.21.0
GOAPIDIFF_VERSION=v0.8.2
COVERAGE=coverage.out

export GO111MODULE := on
export PATH := ./bin:$(PATH)


.PHONY: link-git-hooks
link-git-hooks:
	find .git/hooks -type l -exec rm {} \;
	find githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: test
test:
	go test $(SOURCE_FILES) -coverprofile $(COVERAGE) -timeout=30s -parallel=4 -cover -race

.PHONY: test-examples
test-examples:
	cd examples && go test ./...

.PHONY: fmt
fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w  ./**/**/*.go

.PHONY: lint-fix
lint-fix:
	./bin/golangci-lint run --fix

.PHONY: lint
lint:
	./bin/golangci-lint run $(SOURCE_FILES)

.PHONY: check
check: test lint-fix

.PHONY: tools
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)
	go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)
	go install github.com/joelanford/go-apidiff@$(GOAPIDIFF_VERSION)

.PHONY: addcopy
addcopy:
	@scripts/add-copy.sh

TAG=$(patsubst v%,%,$(shell git describe --tags --dirty --always))
.PHONY: check-version
check-version:
	scripts/check-version.sh "$(TAG)"

.PHONY: openapi-pipeline
openapi-pipeline: tools
	echo "Running OpenAPI Generation and Validation process"
	$(MAKE) -C tools clean_client
	echo "Running client generation"
	$(MAKE) -C tools generate_client
	echo "Validating generated SDK"
	go test ./test

.PHONY: gen-docs
gen-docs:
	$(MAKE) -C tools generate_docs
	./scripts/toc.sh
