#!/bin/bash
set -ueo pipefail

file_path="../internal/core/version.go"
versions_file_path="../openapi/versions.json"

# Use grep to extract the version string and store it in a variable
SDK_VERSION=$(grep -o 'Version = "[^"]*"' "$file_path" | awk -F'"' '{print $2}')
SDK_RESOURCE_VERSION=$(grep -o 'Resource = "[^"]*"' "$file_path" | awk -F'"' '{print $2}')

# Extract the minor and major parts of the version
SDK_MINOR_VERSION=$(echo "$SDK_VERSION" | awk -F'.' '{print $2}')
SDK_MAJOR_VERSION=$(echo "$SDK_VERSION" | awk -F'.' '{print $1}')

# Extract the last version from the versions.json file
HYPEN_RESOURCE_VERSION=$(cat $versions_file_path | jq -r '.versions."2.0" | .[-1]')
NEW_RESOURCE_VERSION=$(echo "$HYPEN_RESOURCE_VERSION" | tr -d '-')

echo "Extracted version from version.go file: '$SDK_VERSION'. Resource Version: '$SDK_RESOURCE_VERSION'"
echo "Extracted version versions.json: '$NEW_RESOURCE_VERSION'."


