export OPENAPI_FOLDER=../openapi
export SDK_FOLDER=../admin

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
	gofmt -w ./${SDK_FOLDER}/*.go

.PHONY: clean_and_generate
clean_and_generate: clean_client generate_client
