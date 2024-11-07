#!/bin/bash
set -ueo pipefail

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"
# shellcheck source=/dev/null
source "$script_path/version-paths.sh"

OLD_PACKAGE="go.mongodb.org/atlas-sdk/${SDK_MAJOR_VERSION}"
NEW_PACKAGE="github.com/mongodb/atlas-sdk-go"
examples_path="$script_path/../../../examples"

echo "Delete specific version $SDK_MAJOR_VERSION from examples go.mod"
(cd "$examples_path" && go mod edit -droprequire="$OLD_PACKAGE")

echo "Modifying all instances from $OLD_PACKAGE to $NEW_PACKAGE across the repository."
npm install
npm exec -c "replace-in-file '/$OLD_PACKAGE/g' '$NEW_PACKAGE' $VERSION_UPDATE_PATHS_PREVIEW --isRegex"

echo "Add preview version to examples go.mod"
(cd "$examples_path" && go get $NEW_PACKAGE)

