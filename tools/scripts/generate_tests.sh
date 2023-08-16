#!/usr/bin/env bash
set -o errexit
set -o nounset

OPENAPI_FOLDER=${OPENAPI_FOLDER:-./openapi}
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-atlas-api.yaml}
FOLDER=${FOLDER:-../test/generated}

transformed_file="atlas-api-transformed.yaml"
client_package="test"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Running Test Generation"

npm exec openapi-generator-cli -- generate \
    -c "./config/config.yaml" -i "$openapiFileLocation" -o "$FOLDER" \
    --package-name="$client_package" \
    --type-mappings=integer=int \
    --ignore-file-override=config/.go-ignore-tests

pushd $FOLDER
for file in *; do
    if [[ ! -f "$file" ]]; then
        continue
    fi

    if [[ ! "$file" == *.* ]]; then
        new_name="${file}.go"
        mv "$file" "$client_package/$new_name"
        echo "Added extension to: $new_name"
    fi
done
popd
