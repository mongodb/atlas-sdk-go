#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"

# Check if there are non-breaking changes
if [ ! -z "$NON_BREAKING_CHANGES" ]; then
    # Increment minor version
    new_minor_version=$((SDK_MINOR_VERSION + 1))
    new_version="${SDK_MAJOR_VERSION}.${new_minor_version}.0"
    
    echo "Updating version to $new_version due to non-breaking changes"
    
    # Update version.go file
    sed -i "s/Version = \"$SDK_VERSION\"/Version = \"$new_version\"/" "../internal/core/version.go"
    
    # Create a commit with the version update
    git add "../internal/core/version.go"
    git commit -m "chore: bump minor version to $new_version for non-breaking changes"
    git push
fi 