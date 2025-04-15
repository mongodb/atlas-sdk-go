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

## Base URL
API_BASE_URL=${API_BASE_URL:-"https://cloud.mongodb.com"}

## Folder used for fetching files
OPENAPI_FOLDER=${OPENAPI_FOLDER:-"../openapi"}

## S3 bucket where the spec is hosted
S3_BUCKET=${S3_BUCKET:-"mongodb-mms-prod-build-server"}

versions_url="${API_BASE_URL}/api/openapi/versions"
# versions_file="versions.json"

pushd "${OPENAPI_FOLDER}"
echo "Fetching versions from $versions_url"

# curl --show-error --fail --silent -o "${versions_file}" \
#      -H "Accept: application/json" "${versions_url}"

## Dynamic Versioned API Version
# CURRENT_API_REVISION=$(jq -r '.versions."2.0" | .[-1]' <"./${versions_file}")

# echo "Fetching OpenAPI release sha"
# sha=$(curl --show-error --fail --silent -H "Accept: text/plain" "${API_BASE_URL}/api/private/unauth/version")

# echo "Fetching OAS file for ${sha}"
# openapi_url="https://${S3_BUCKET}.s3.amazonaws.com/openapi/${sha}-v2-${CURRENT_API_REVISION}.yaml"
openapi_url="https://raw.githubusercontent.com/mongodb/openapi/refs/heads/main/openapi/v2/openapi-2023-11-15.yaml"

echo "Fetching api from $openapi_url to $OPENAPI_FILE_NAME"

curl --show-error --fail --silent -o "$OPENAPI_FILE_NAME" "$openapi_url"

popd -0
