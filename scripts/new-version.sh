#!/bin/bash
set -ueo pipefail

file_path="./openapi/versions.json"
target_file_path="./internal/core/version.go"

last_element=`cat $file_path | jq -r '.versions."2.0" | .[-1]'`

# Remove hyphens from the last_element
cleaned_last_element=$(echo "$last_element" | tr -d '-')

# Set the extracted version as an environment variable
export CURRENT_RESOURCE_VERSION="$cleaned_last_element"

# Display the extracted version
echo "Current Resource Version: $CURRENT_RESOURCE_VERSION"

source ./scripts/extract-version.sh

# Update the version.go file with the new version
if [ "$CURRENT_RESOURCE_VERSION" == "$SDK_RESOURCE_VERSION" ]; then
	echo "Resource Version is already up to date. Changing minor version."
	# Extract the minor version from the SDK_VERSION
	minor_version=$(echo "$SDK_VERSION" | awk -F'.' '{print $2}')
	# Increment the minor version
	new_minor_version=$((minor_version + 1))
	# Update the SDK_VERSION
	export SDK_VERSION="v${CURRENT_RESOURCE_VERSION}001.${new_minor_version}.0"
else	
	# Update the SDK_VERSION
	echo "Resource Version is not up to date. Changing major version."
	export SDK_VERSION="v${CURRENT_RESOURCE_VERSION}001.0.0"
	echo "Modifying Resource Version across the repository."
	for file in $(find . -type f -name "*"); do
    	sed 's/$CURRENT_RESOURCE_VERSION/$SDK_RESOURCE_VERSION}/g' "${file}_tmp"
		mv "${file}_tmp" "${file}"
	done
fi

echo "Creating new version.go file with $SDK_VERSION and resource version: $CURRENT_RESOURCE_VERSION"

echo "package core

// Version of the SDK used for the autorelease process.
// When version is bumped in the PR it means that the PR is ready to be merged.
// For more information please see: https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_1_concepts.md
const (
	// SDK release tag version
	Version = \"${SDK_VERSION}\"
	// Resource Version
	Resource = \"$CURRENT_RESOURCE_VERSION\"
)" > $target_file_path

