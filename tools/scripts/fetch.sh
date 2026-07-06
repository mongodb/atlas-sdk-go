#!/usr/bin/env bash

set -o errexit
set -o nounset

#########################################################
# Fetch openapi from remote file
# Environment variables:
#   CURRENT_REVISION - current revision of the versioned API
#   OPENAPI_FILE_NAME - openapi file name to use
#   OPENAPI_FOLDER - folder for saving openapi file
#   TARGET_API_VERSION - optional; when set to "preview" fetches the latest
#     stable dated spec AND the preview delta spec (openapi-preview.yaml) and
#     merges the preview operations/schemas on top of the stable spec. This
#     produces a superset SDK: every stable endpoint plus the preview
#     media-type (application/vnd.atlas.preview+json) ones. The preview file is
#     only a small delta, so it must be merged rather than used on its own.
#########################################################

## Input variables with defaults

## OpenAPI file (latest)
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-"atlas-api.yaml"}

## Folder used for fetching files
OPENAPI_FOLDER=${OPENAPI_FOLDER:-"../openapi"}

## Optional target API version. When "preview", fetch the preview spec.
TARGET_API_VERSION=${TARGET_API_VERSION:-""}


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

## Dynamic Versioned API Version (latest stable, used as the base spec)
CURRENT_API_REVISION=$(jq -r '.[-1]' < "./${versions_file}")

echo "Fetching OAS file for version $CURRENT_API_REVISION"
openapi_url=${API_BASE_URL}/openapi-${CURRENT_API_REVISION}.yaml

echo "Fetching api from $openapi_url to $OPENAPI_FILE_NAME"

curl --show-error --fail --silent -o "$OPENAPI_FILE_NAME" "$openapi_url"

if [ "${TARGET_API_VERSION}" = "preview" ]; then
  # Additionally fetch the preview delta spec so we can merge its preview
  # media-type operations/schemas on top of the stable base spec.
  PREVIEW_FILE_NAME="atlas-api-preview.yaml"
  preview_url=${API_BASE_URL}/openapi-preview.yaml
  echo "Fetching preview delta from $preview_url to $PREVIEW_FILE_NAME"
  curl --show-error --fail --silent -o "$PREVIEW_FILE_NAME" "$preview_url"
fi

# Return to the tools directory (script's invocation dir) so the merge step below
# runs npm from there and asdf resolves the right toolchain.
popd

if [ "${TARGET_API_VERSION}" = "preview" ]; then
  # Merge the preview delta onto the stable base spec (in place), producing a
  # superset spec for the rest of the generation pipeline. Mirrors the
  # transformer invocation used in transform.sh.
  echo "Merging preview delta into $OPENAPI_FILE_NAME"
  npm install
  npm exec --prefix .. atlas-openapi-transformer -- merge-preview \
    --base "${OPENAPI_FOLDER}/${OPENAPI_FILE_NAME}" \
    --preview "${OPENAPI_FOLDER}/atlas-api-preview.yaml" \
    --output "${OPENAPI_FOLDER}/${OPENAPI_FILE_NAME}"
  rm -f "${OPENAPI_FOLDER}/atlas-api-preview.yaml"
fi
