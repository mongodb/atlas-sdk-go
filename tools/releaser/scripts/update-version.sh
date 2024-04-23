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
# Remove leading zeros from the variable
major_version_breaking_changes=$((10#$major_version_breaking_changes))
echo $major_version_breaking_changes "major_version_breaking_changes"
# Change value
((major_version_breaking_changes++))
# Ensure the value stays within the range 1-999
major_version_breaking_changes=$((major_version_breaking_changes % 1000))

major_version_bump=$(printf "%03d" "$major_version_breaking_changes")
BUMPED_MAJOR_VERSION="v${SDK_RESOURCE_VERSION}${major_version_bump}"

mkdir -p "$script_path/../breaking_changes/"
export TARGET_BREAKING_CHANGES_FILE=${BUMPED_MAJOR_VERSION}
echo -e "# Breaking Changes\n## SDK changes\n ${BREAKING_CHANGES:-TODO} \n## API Changelog\n https://www.mongodb.com/docs/atlas/reference/api-resources-spec/changelog" \
		> "$script_path/../breaking_changes/${TARGET_BREAKING_CHANGES_FILE}.md"

echo "Modifying $SDK_MAJOR_VERSION to $BUMPED_MAJOR_VERSION Resource Version across the repository."
npm install
npm exec -c "replace-in-file /$SDK_MAJOR_VERSION/g $BUMPED_MAJOR_VERSION $VERSION_UPDATE_PATHS --isRegex"

## Explicitly update version.go file
export SDK_VERSION="${BUMPED_MAJOR_VERSION}.0.0"
export NEW_RESOURCE_VERSION="${SDK_RESOURCE_VERSION}"

echo "Updating version.go file: $SDK_VERSION"
target_file_path="$script_path/../../../internal/core/version.go"
envsubst < "$script_path/../templates/VERSION.tmpl" > "$target_file_path"

