#!/bin/bash
set -ueo pipefail

## Target version file
target_file_path="../internal/core/version.go"

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"
# shellcheck source=/dev/null
source "$script_path/version-paths.sh"

# Update the version.go file with the new version
if [ "$NEW_RESOURCE_VERSION" == "$SDK_RESOURCE_VERSION" ]; then
	echo "Resource Version is already up to date. Changing minor version."
	# Increment the minor version
	new_minor_version=$((SDK_MINOR_VERSION + 1))
	# Update the SDK_VERSION
	SDK_VERSION="${SDK_MAJOR_VERSION}.${new_minor_version}.0"
else	 
	# Update the SDK_VERSION
	echo "Resource Version is not up to date. Changing major version."
	NEW_MAJOR_VERSION="${NEW_RESOURCE_VERSION}001"
	SDK_VERSION="v${NEW_MAJOR_VERSION}.0.0" 
	echo "Modifying $NEW_RESOURCE_VERSION to $SDK_RESOURCE_VERSION Resource Version across the repository."
	npm exec -c "replace-in-file /$SDK_MAJOR_VERSION/g $NEW_MAJOR_VERSION $VERSION_UPDATE_PATHS --isRegex"
fi 

echo "Creating new version.go file with $SDK_VERSION and resource version: $NEW_RESOURCE_VERSION"

export SDK_VERSION
export NEW_RESOURCE_VERSION

envsubst < "$script_path/../templates/VERSION.tmpl" > $target_file_path
