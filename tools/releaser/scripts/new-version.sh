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

	echo "Print breaking changes"
	# shellcheck source=/dev/null
	source "$script_path/breaking-changes.sh"
	if [ -n "$BREAKING_CHANGES" ]; then
		echo "BREAKING CHANGES DETECTED FOR NON MAJOR VERSION BUMP"
		# source "$script_path/update-version.sh"	# we don't want to update the version so commenting this out for now
		# exit 0;
	fi
else
	# Update the SDK_VERSION
	echo "Resource Version is not up to date. Changing major version."
	NEW_MAJOR_VERSION="v${NEW_RESOURCE_VERSION}001"
	SDK_VERSION="${NEW_MAJOR_VERSION}.0.0"
	echo "generate breaking changes file"
	export TARGET_BREAKING_CHANGES_FILE=${NEW_MAJOR_VERSION}
	# shellcheck source=/dev/null
	source "$script_path/breaking-changes.sh"

	echo "Modifying all instances of version from $SDK_RESOURCE_VERSION to $NEW_RESOURCE_VERSION across the repository."
	npm install
	npm exec -c "replace-in-file /$SDK_MAJOR_VERSION/g $NEW_MAJOR_VERSION $VERSION_UPDATE_PATHS --isRegex"
fi

echo "Creating new version.go file with $SDK_VERSION and resource version: $NEW_RESOURCE_VERSION"

export SDK_VERSION
export NEW_RESOURCE_VERSION

envsubst <"$script_path/../templates/VERSION.tmpl" >$target_file_path
