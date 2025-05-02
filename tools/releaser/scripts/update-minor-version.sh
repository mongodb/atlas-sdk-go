#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"
# shellcheck source=/dev/null
source "$script_path/version-paths.sh"

# Display the extracted version
echo "Current Resource Version: $NEW_RESOURCE_VERSION"

# Increment minor version
new_minor=$((SDK_MINOR_VERSION + 1))

# Create new version string
BUMPED_VERSION="${SDK_MAJOR_VERSION}.${new_minor}.0"

echo "Modifying $SDK_MAJOR_VERSION.$SDK_MINOR_VERSION.0 to $BUMPED_VERSION across the repository."
npm install
npm exec -c "replace-in-file /$SDK_MAJOR_VERSION\.$SDK_MINOR_VERSION\.0/g $BUMPED_VERSION $VERSION_UPDATE_PATHS --isRegex"

## Explicitly update version.go file
export SDK_VERSION="${BUMPED_VERSION}"

echo "Updating version.go file: $SDK_VERSION"
target_file_path="$script_path/../../../internal/core/version.go"
envsubst < "$script_path/../templates/VERSION.tmpl" > "$target_file_path" 