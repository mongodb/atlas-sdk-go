#!/usr/bin/env bash
set -o errexit
set -o nounset

OPENAPI_FOLDER=${OPENAPI_FOLDER:-./openapi}
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-atlas-api.yaml}
FOLDER=${FOLDER:-../test/generated}

transformed_file="atlas-api-transformed.yaml"
client_package="admin"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Running Test Generation"

npm exec openapi-generator-cli -- generate \
    -c "./config/config.yaml" -i "$openapiFileLocation" -o "$FOLDER" \
    --package-name="$client_package" \
    --type-mappings=integer=int \
    --type-mappings=object=any \
    --type-mappings=file=io.ReadCloser \
    --ignore-file-override=config/.go-ignore-tests

