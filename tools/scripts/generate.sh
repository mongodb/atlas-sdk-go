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
SDK_FOLDER=${SDK_FOLDER:-./}
transformed_file="atlas-api-transformed.yaml"
client_package="admin"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Running generation pipeline"
echo "# Running transformation based on $OPENAPI_FILE_NAME to the $transformed_file"
cp "$OPENAPI_FOLDER/$OPENAPI_FILE_NAME" "$openapiFileLocation"

npm install
npm run sdk:transform -- "$openapiFileLocation"

echo "# Running OpenAPI generator validation"
npm exec openapi-generator-cli -- validate -i "$openapiFileLocation"

echo "# Running Client Generation"

npm exec openapi-generator-cli -- generate \
    -c "./config/config.yaml" -i "$openapiFileLocation" -o "$SDK_FOLDER" \
    --package-name="$client_package" \
    --type-mappings=integer=int \
    --type-mappings=object=interface{} \
    --type-mappings=file=io.ReadCloser \
    --ignore-file-override=config/.go-ignore

# Check if goimports exists and is executable
if ! command -v goimports &> /dev/null; then
  echo "Error: goimports command not found. Please install goimports by running 'make install-goimports' before running this script."
  exit 1
fi

goimports -w "$SDK_FOLDER"
