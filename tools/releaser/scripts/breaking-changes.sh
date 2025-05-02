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

# Extract changes based on pattern
extractChanges() {
    local raw_changes=$1
    local pattern=$2
    local header=$3
    echo "$raw_changes" | awk -v pattern="$pattern" -v header="$header" '
        $0 ~ pattern {print header; collecting=1; next}
        collecting && /^#/ {collecting=0}
        collecting && NF {print "- "$0}
    '
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

BREAKING_CHANGES=$(extractChanges "$RAW_CHANGES" "## incompatible changes" "### incompatible changes")
NON_BREAKING_CHANGES=$(extractChanges "$RAW_CHANGES" "## compatible changes" "### compatible changes")

set -e
popd || exit

if [ -z "$BREAKING_CHANGES" ] && [ -z "$NON_BREAKING_CHANGES" ]; then
  echo "No changes detected"
else
  echo "Detected changes in the release"
  if [ -z "$TARGET_BREAKING_CHANGES_FILE" ]; then
    echo "Changes for the release:"
    [ -n "$BREAKING_CHANGES" ] && echo -e "\nBreaking Changes:\n$BREAKING_CHANGES"
    [ -n "$NON_BREAKING_CHANGES" ] && echo -e "\nNon-Breaking Changes:\n$NON_BREAKING_CHANGES"
  else
    # Only create breaking changes file for major version bumps
    if [ -n "$BREAKING_CHANGES" ]; then
      echo "Creating the breaking changes file with following breaking changes:"
      echo "$BREAKING_CHANGES"
      echo -e "# Breaking Changes\n## SDK changes\n$BREAKING_CHANGES\n## API Changelog\n https://www.mongodb.com/docs/atlas/reference/api-resources-spec/changelog" \
        >"$script_path/../breaking_changes/${TARGET_BREAKING_CHANGES_FILE}.md"
    fi
  fi
fi
