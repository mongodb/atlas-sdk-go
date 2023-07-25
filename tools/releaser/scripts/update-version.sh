#!/bin/bash
set -ueo pipefail

source ./scripts/extract-version.sh

# Display the extracted version
echo "Current Resource Version: $CURRENT_RESOURCE_VERSION"

echo "Modifying $CURRENT_RESOURCE_VERSION Resource Version across the repository."

npm exec replace-in-file $CURRENT_RESOURCE_VERSION $SDK_RESOURCE_VERSION ./**/*.go,./**/*.md,./**/*.mustache
 
