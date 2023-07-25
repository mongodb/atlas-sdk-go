#!/bin/bash
set -ueo pipefail

target_file_path="../internal/core/version.go"

script_path=$(dirname "$0")
source $script_path/extract-version.sh

# Update the version.go file with the new version
if [ "$CURRENT_RESOURCE_VERSION" == "$SDK_RESOURCE_VERSION" ]; then
	echo "Resource Version is already up to date. Changing minor version."
	# Extract the minor version from the SDK_VERSION
	minor_version=$(echo "$SDK_VERSION" | awk -F'.' '{print $2}')
	major_version=$(echo "$SDK_VERSION" | awk -F'.' '{print $1}')
	# Increment the minor version
	new_minor_version=$((minor_version + 1))
	# Update the SDK_VERSION
	export SDK_VERSION="${major_version}.${new_minor_version}.0"
else	
	# Update the SDK_VERSION
	echo "Resource Version is not up to date. Changing major version."
	export SDK_VERSION="v${CURRENT_RESOURCE_VERSION}001.0.0"
	echo "Modifying $CURRENT_RESOURCE_VERSION Resource Version across the repository."
	npm exec replace-in-file $CURRENT_RESOURCE_VERSION $SDK_RESOURCE_VERSION ./* --ignore=*/node_modules/*, ./.* --dry
fi

echo "Creating new version.go file with $SDK_VERSION and resource version: $CURRENT_RESOURCE_VERSION"

echo "package core

// Version of the SDK used for the autorelease process.
// When version is bumped in the PR it means that the PR is ready to be merged.
// For more information please see: https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_1_concepts.md
const (
	// SDK release tag version.
	Version = \"${SDK_VERSION}\"
	// Resource Version.
	Resource = \"$CURRENT_RESOURCE_VERSION\"
)" > $target_file_path

