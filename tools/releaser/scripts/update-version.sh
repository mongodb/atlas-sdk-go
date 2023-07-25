#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
source "$script_path/extract-version.sh"
source "$script_path/version-paths.sh"

# Display the extracted version
echo "Current Resource Version: $NEW_RESOURCE_VERSION"

major_version_breaking_changes="${SDK_MAJOR_VERSION: -3}"
major_version_bump=$((major_version_breaking_changes+1))
major_version_bump=$(printf "%03d" "$major_version_bump")
echo $major_version_bump
BUMPED_MAJOR_VERSION="v${SDK_RESOURCE_VERSION}${major_version_bump}"

echo "Modifying $SDK_MAJOR_VERSION to $BUMPED_MAJOR_VERSION Resource Version across the repository."
npm exec -c "replace-in-file $SDK_MAJOR_VERSION $BUMPED_MAJOR_VERSION $VERSION_UPDATE_PATHS"
