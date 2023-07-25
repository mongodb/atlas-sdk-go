#!/bin/bash
set -ueo pipefail

file_path="../internal/core/version.go"
versions_file_path="../openapi/versions.json"

# Use grep to extract the version string and store it in a variable
version=$(grep -o 'Version = "[^"]*"' "$file_path" | awk -F'"' '{print $2}')
resource_version=$(grep -o 'Resource = "[^"]*"' "$file_path" | awk -F'"' '{print $2}')

export SDK_VERSION="$version"
export SDK_RESOURCE_VERSION="$resource_version"

echo "Extracted version from version.go file: '$SDK_VERSION'. Resource Version: '$SDK_RESOURCE_VERSION'"

last_element=$(cat $versions_file_path | jq -r '.versions."2.0" | .[-1]')

# Remove hyphens from the last_element
cleaned_last_element=$(echo "$last_element" | tr -d '-')

# Set the extracted version as an environment variable
export CURRENT_RESOURCE_VERSION="$cleaned_last_element"

echo "Extracted version versions.json: '$CURRENT_RESOURCE_VERSION'."


