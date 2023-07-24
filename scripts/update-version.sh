#!/bin/bash
set -ueo pipefail

source ./scripts/extract-version.sh

# Display the extracted version
echo "Current Resource Version: $CURRENT_RESOURCE_VERSION"

echo "Modifying $CURRENT_RESOURCE_VERSION Resource Version across the repository."
find . -type f -name "*" -not -path '*/\.*' -exec grep -l "$CURRENT_RESOURCE_VERSION" {} + | while read -r file; do
    echo "Modifying $file"
    sed "s/$CURRENT_RESOURCE_VERSION/$SDK_RESOURCE_VERSION/g" "$file" > "${file}_tmp"
    mv "${file}_tmp" "$file"
done
