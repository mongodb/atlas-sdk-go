#!/bin/bash
set -ueo pipefail

file_path="./internal/core/version.go"

# Use grep to extract the version string and store it in a variable
version=$(grep -o 'Version = "[^"]*"' "$file_path" | awk -F'"' '{print $2}')
resource_version=$(grep -o 'Resource = "[^"]*"' "$file_path" | awk -F'"' '{print $2}')

export SDK_VERSION="$version"
export SDK_RESOURCE_VERSION="$resource_version"

echo "Extracted version from version.go file: '$SDK_VERSION'. Resource Version: '$SDK_RESOURCE_VERSION'"
