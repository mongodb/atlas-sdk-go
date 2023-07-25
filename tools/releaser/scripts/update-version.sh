#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
source "$script_path/extract-version.sh"
source $script_path/version-paths.sh

# Display the extracted version
echo "Current Resource Version: $NEW_RESOURCE_VERSION"

echo "Modifying $NEW_RESOURCE_VERSION to $SDK_RESOURCE_VERSION Resource Version across the repository."
npm exec -c "replace-in-file $NEW_RESOURCE_VERSION $SDK_RESOURCE_VERSION $VERSION_UPDATE_PATHS"
 
