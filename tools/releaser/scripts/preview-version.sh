#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"
# shellcheck source=/dev/null
source "$script_path/version-paths.sh"

NEW_VERSION="preview"
echo "Modifying all instances of version from $SDK_RESOURCE_VERSION to $NEW_VERSION across the repository."
npm install
npm exec -c "replace-in-file /$SDK_MAJOR_VERSION/g $NEW_VERSION $VERSION_UPDATE_PATHS_PREVIEW --isRegex"
