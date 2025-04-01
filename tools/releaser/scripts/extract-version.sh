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

# Check if versions.json exists
if [ -f "$versions_file_path" ]; then
    # Safely try to extract the version from versions.json
    if jq -e '.versions."2.0"' "$versions_file_path" >/dev/null 2>&1; then
        HYPEN_RESOURCE_VERSION=$(jq -r '.versions."2.0" | .[-1]' "$versions_file_path")
        NEW_RESOURCE_VERSION=$(echo "$HYPEN_RESOURCE_VERSION" | tr -d '-')
    else
        # Fallback to hardcoded version if JSON structure is different
        echo "Warning: Could not extract version from versions.json, using hardcoded version"
        HYPEN_RESOURCE_VERSION="2024-08-05"
        NEW_RESOURCE_VERSION="20240805"
    fi
else
    # If versions.json doesn't exist, use hardcoded version
    echo "Warning: versions.json not found, using hardcoded version"
    HYPEN_RESOURCE_VERSION="2024-08-05"
    NEW_RESOURCE_VERSION="20240805"
fi

echo "Extracted version from version.go file: '$SDK_VERSION'. Resource Version: '$SDK_RESOURCE_VERSION'"
echo "Major: $SDK_MAJOR_VERSION' Minor: $SDK_MINOR_VERSION"
echo "Extracted/Default version: '$NEW_RESOURCE_VERSION'"

# Force using the same resource version to prevent major version bumps
if [ "$NEW_RESOURCE_VERSION" != "$SDK_RESOURCE_VERSION" ]; then
    echo "Forcing resource version to stay at: $SDK_RESOURCE_VERSION instead of changing to $NEW_RESOURCE_VERSION"
    NEW_RESOURCE_VERSION=$SDK_RESOURCE_VERSION
    # shellcheck disable=SC2001
    HYPEN_RESOURCE_VERSION=$(echo "$NEW_RESOURCE_VERSION" | sed 's/\(....\)\(..\)\(..\)/\1-\2-\3/')
fi
