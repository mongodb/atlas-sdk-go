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

# Function to extract changes by section
extract_changes() {
    local section=$1
    echo "$RAW_CHANGES" | awk -v section="$section" '
        $0 ~ section {print "### " section; collecting=1; next}
        collecting && /^#/ {collecting=0}
        collecting && NF {print "- "$0}
    '
}

# Extract different types of changes
BREAKING_CHANGES=$(extract_changes "incompatible changes")
NEW_FEATURES=$(extract_changes "new features")
BUG_FIXES=$(extract_changes "bug fixes")
DEPRECATIONS=$(extract_changes "deprecations")
OTHER_CHANGES=$(extract_changes "other changes")

# Combine non-breaking changes for release notes
NON_BREAKING_CHANGES=""
[ -n "$NEW_FEATURES" ] && NON_BREAKING_CHANGES+="\n## New Features\n$NEW_FEATURES"
[ -n "$BUG_FIXES" ] && NON_BREAKING_CHANGES+="\n## Bug Fixes\n$BUG_FIXES"
[ -n "$DEPRECATIONS" ] && NON_BREAKING_CHANGES+="\n## Deprecations\n$DEPRECATIONS"
[ -n "$OTHER_CHANGES" ] && NON_BREAKING_CHANGES+="\n## Other Changes\n$OTHER_CHANGES"

set -e
popd || exit

if [ -z "$BREAKING_CHANGES" ] && [ -z "$NON_BREAKING_CHANGES" ]; then
  echo "No changes detected"
else
  echo "Detected changes in the release"
  if [ -z "$TARGET_BREAKING_CHANGES_FILE" ]; then
    echo "Breaking changes for the release:"
    [ -n "$BREAKING_CHANGES" ] && printf "\nBreaking Changes:\n%s" "$BREAKING_CHANGES"
    echo "Non-breaking changes for the release:"
    [ -n "$NON_BREAKING_CHANGES" ] && printf "\nNon-Breaking Changes:%s" "$NON_BREAKING_CHANGES"
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
