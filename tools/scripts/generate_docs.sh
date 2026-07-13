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

# The spec has been growing by roughly 1KB/day; SnakeYAML's parser (bundled
# with the generator) refuses YAML documents over 3MB by default. Raise that
# ceiling well ahead of time so growth doesn't silently break generation.
export JAVA_OPTS="${JAVA_OPTS:-} -DmaxYamlCodePoints=20000000"

echo "# Running Client Generation"

npm exec openapi-generator-cli -- generate \
    -c "./config/config.yaml" -i "$openapiFileLocation" -o "$DOC_FOLDER" \
    --package-name="$client_package" \
    --type-mappings=integer=int \
    --type-mappings=object=interface{} \
    --type-mappings=interface{}=any \
    --type-mappings=file=io.ReadCloser \
    --ignore-file-override=config/.go-ignore-docs

mv "$DOC_FOLDER"/README.md "$DOC_FOLDER"/doc_last_reference.md

# Replace default import prefix in examples so docs compile with this SDK.
find "$DOC_FOLDER/docs" -name "*.md" -exec perl -pi -e "s/openapiclient/${client_package}/g" {} +

# The Go generator names API doc files with an "Api" suffix (e.g. ThirdPartyIntegrationsApi.md)
# but links/classnames use the Go type name, which has an "API" suffix. Fix these broken links.
find "$DOC_FOLDER/docs" -name "*.md" -exec perl -pi -e 's/([A-Za-z]+)API\.md/$1Api.md/g' {} +
perl -pi -e 's/([A-Za-z]+)API\.md/$1Api.md/g' "$DOC_FOLDER/doc_last_reference.md"
