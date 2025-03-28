#!/bin/bash
set -eu

# Inputs:
TARGET_BREAKING_CHANGES_FILE=${TARGET_BREAKING_CHANGES_FILE:-""}
script_path=$(dirname "$0")

baseVersion() {
	local base_version
    # Create an isolated subshell to contain the sourcing to avoid overwritting version vars
	base_version=$(
        {
			# shellcheck source=/dev/null
            source "$script_path/extract-version.sh" >&2
            # Output only the variable we want
            echo "github.com/mongodb/atlas-sdk-go/$SDK_MAJOR_VERSION@$SDK_VERSION"
        }
    )
    local ret_val=$?
    if [ $ret_val -ne 0 ]; then
        echo "Error when sourcing extract-version.sh" >&2
        return $ret_val
    fi
    export BASE_VERSION="$base_version"
}

baseVersion

echo "Installing gorelease"
go install golang.org/x/exp/cmd/gorelease@latest >/dev/null

pushd "$script_path/../../../" || exit ## workaround for --repo-path="../" not working
echo "Changed directory to $(pwd)"
set +e

RAW_CHANGES=$(gorelease -base "$BASE_VERSION")
echo "Changes detected from BASE_VERSION $BASE_VERSION:"
echo "$RAW_CHANGES"

BREAKING_CHANGES=$(echo "$RAW_CHANGES" | awk '
    /## incompatible changes/ {print "### incompatible changes"; collecting=1; next}
    collecting && /^#/ {collecting=0}
    collecting && NF {print "- "$0}
')

set -e
popd || exit

if [ -z "$BREAKING_CHANGES" ]; then
  echo "No major breaking changes detected"
else
  echo "Detected major breaking changes in the release"
  if [ -z "$TARGET_BREAKING_CHANGES_FILE" ]; then
    echo "Breaking changes for the major release"
    echo "$BREAKING_CHANGES"
  else
    echo "Creating the breaking changes file with following breaking changes:"
    echo "$BREAKING_CHANGES"
    echo -e "# Breaking Changes\n## SDK changes\n$BREAKING_CHANGES\n## API Changelog\n https://www.mongodb.com/docs/atlas/reference/api-resources-spec/changelog" \
      >"$script_path/../breaking_changes/${TARGET_BREAKING_CHANGES_FILE}.md"
  fi
fi
