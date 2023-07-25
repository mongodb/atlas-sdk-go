#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
source "$script_path/extract-version.sh"

# Display the extracted version
echo "Current Resource Version: $CURRENT_RESOURCE_VERSION"

echo "Modifying $CURRENT_RESOURCE_VERSION Resource Version across the repository."

npm exec replace-in-file -- $CURRENT_RESOURCE_VERSION $SDK_RESOURCE_VERSION *.go,*.md,.*.mustache
 
