export OPENAPI_FOLDER=../openapi
export SDK_FOLDER=../admin
export DOC_FOLDER=../docs

.DEFAULT_GOAL := generate_client

.PHONY: clean_client
clean_client:
	rm -Rf ../admin

.PHONY: fetch_openapi
fetch_openapi:
	./scripts/fetch.sh

.PHONY: transform_openapi
transform_openapi:
	./scripts/transform.sh

.PHONY: generate_client
generate_client:
	./scripts/generate.sh

.PHONY: clean_and_generate
clean_and_generate: clean_client generate_client

.PHONY: generate_docs
generate_docs:	
	rm -Rf ../docs/docs/*
	./scripts/generate_docs.sh

.PHONY: generate_tests
generate_tests:	
	rm -Rf ../test/generated/test*
	./scripts/generate_tests.sh

.PHONY: new-release
new-release:
	./releaser/scripts/new-version.sh

.PHONY: update-version
update-version:
	## When manually updating major version
	./releaser/scripts/update-version.sh

.PHONY: generate_mocks
generate_mocks:	
	go install github.com/vektra/mockery/v2@v2.42.0 && mockery

