SOURCE_FILES?=./...
GOLANGCI_VERSION=v2.1.6 # Also update golangci-lint GH action in pr.yml when updating this version

export GO111MODULE := on
export PATH := ./bin:$(PATH)

default: build

.PHONY: link-git-hooks
link-git-hooks:
	find .git/hooks -type l -exec rm {} \;
	find githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: build
build:
	go install $(SOURCE_FILES)

.PHONY: test
test:
	go test $(SOURCE_FILES) -timeout=30s -parallel=4 -race

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

.PHONY: install-golangci-lint
install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)

.PHONY: install-goimports
install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: tools
tools: install-golangci-lint install-goimports

.PHONY: addcopy
addcopy:
	@scripts/add-copy.sh

TAG=$(patsubst v%,%,$(shell git describe --tags --dirty --always))
.PHONY: check-version
check-version:
	scripts/check-version.sh "$(TAG)"

.PHONY: openapi-pipeline
openapi-pipeline: install-goimports
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

.PHONY: gen-purls
gen-purls:
	./scripts/compliance/gen-purls.sh

.PHONY: gen-sbom
gen-sbom:
	./scripts/compliance/gen-sbom.sh

.PHONY: gen-ssdlc-report
gen-ssdlc-report:
	./scripts/compliance/gen-ssdlc-report.sh

.PHONY: upload-sbom
upload-sbom:
	./scripts/compliance/upload-sbom.sh
