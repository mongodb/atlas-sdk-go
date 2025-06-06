#!/usr/bin/env bash

set -o errexit
set -o nounset

#########################################################
# Fetch openapi from remote file
# Environment variables:
#   CURRENT_REVISION - current revision of the versioned API
#   OPENAPI_FILE_NAME - openapi file name to use
#   OPENAPI_FOLDER - folder for saving openapi file
#   S3_BUCKET - S3 bucket where the spec is hosted
#########################################################

## Input variables with defaults

## OpenAPI file (latest)
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-"atlas-api.yaml"}

## Folder used for fetching files
OPENAPI_FOLDER=${OPENAPI_FOLDER:-"../openapi"}


# Base URL for fetching the openapi file
API_BASE_URL=${API_BASE_URL:-https://raw.githubusercontent.com/mongodb/openapi/refs/heads/main/openapi/v2}
versions_file="versions.json"

pushd "${OPENAPI_FOLDER}"
versions_url="${API_BASE_URL}/${versions_file}"
echo "Fetching versions from $versions_url"

curl --show-error --fail --silent -o "${versions_file}" \
     -H "Accept: application/json" "${versions_url}"

# Remove "preview" and versions with ".upcoming" suffix from versions file and update the file
jq 'map(select(. != "preview" and (endswith(".upcoming") | not)))' < "./${versions_file}" > "./${versions_file}.tmp"
mv "./${versions_file}.tmp" "./${versions_file}"

## Dynamic Versioned API Version
CURRENT_API_REVISION=$(jq -r '.[-1]' < "./${versions_file}")

echo "Fetching OAS file for version $CURRENT_API_REVISION"
openapi_url=${API_BASE_URL}/openapi-${CURRENT_API_REVISION}.yaml

echo "Fetching api from $openapi_url to $OPENAPI_FILE_NAME"

curl --show-error --fail --silent -o "$OPENAPI_FILE_NAME" "$openapi_url"

popd -0 
