#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"
# shellcheck source=/dev/null
source "$script_path/version-paths.sh"

# # Display the extracted version
echo "Current Resource Version: $NEW_RESOURCE_VERSION"

major_version_breaking_changes="${SDK_MAJOR_VERSION: -3}"
major_version_bump=$((major_version_breaking_changes+1))
major_version_bump=$(printf "%03d" "$major_version_bump")
BUMPED_MAJOR_VERSION="v${SDK_RESOURCE_VERSION}${major_version_bump}"

echo "Modifying $SDK_MAJOR_VERSION to $BUMPED_MAJOR_VERSION Resource Version across the repository."
npm exec -c "replace-in-file /$SDK_MAJOR_VERSION/g $BUMPED_MAJOR_VERSION $VERSION_UPDATE_PATHS --isRegex"

echo "Creating empty breaking changes file for $BUMPED_MAJOR_VERSION"
echo "# Breaking Changes" > "$script_path/../breaking_changes/${BUMPED_MAJOR_VERSION}.md"

## Explicitly update version.go file
export SDK_VERSION="${BUMPED_MAJOR_VERSION}.0.0"
export NEW_RESOURCE_VERSION="${BUMPED_MAJOR_VERSION}"

echo "Updating version.go file: $SDK_VERSION"
target_file_path="$script_path/../../../internal/core/version.go"
envsubst < "$script_path/../templates/VERSION.tmpl" > "$target_file_path"

