#!/usr/bin/env bash
set -o errexit
set -o nounset

######################################################
# Generate client using OpenAPI generator
# Environment variables:
#   OPENAPI_FOLDER - folder containing openapi file
#   OPENAPI_FILE_NAME - openapi file name
#   SDK_FOLDER - folder location for generated client
######################################################

OPENAPI_FOLDER=${OPENAPI_FOLDER:-./openapi}
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-atlas-api.yaml}
DOC_FOLDER=${DOC_FOLDER:-./docs}

transformed_file="atlas-api-transformed.yaml"
client_package="admin"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Running Client Generation"

npm exec openapi-generator-cli -- generate \
    -c "./config/config.yaml" -i "$openapiFileLocation" -o "$DOC_FOLDER" \
    --package-name="$client_package" \
    --type-mappings=integer=int \
    --ignore-file-override=config/.go-ignore-docs

mv "$DOC_FOLDER"/README.md "$DOC_FOLDER"/doc_1_reference.md
