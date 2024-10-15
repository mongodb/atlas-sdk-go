#!/bin/bash
set -eu
GOPATH=$(go env GOPATH)

# Inputs:
# TARGET_BREAKING_CHANGES_FILE - file to save breaking changes
script_path=$(dirname "$0")

echo "Installing go-apidiff"
go install github.com/joelanford/go-apidiff@latest > /dev/null

TARGET_BREAKING_CHANGES_FILE=${TARGET_BREAKING_CHANGES_FILE:-""}

pushd "$script_path/../../../" || exit ## workaround for --repo-path="../" not working
echo "Changed directory to $(pwd)"
set +e
BREAKING_CHANGES=$("$GOPATH/bin/go-apidiff" main --compare-imports="false" --print-compatible="false")
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
      > "$script_path/../breaking_changes/${TARGET_BREAKING_CHANGES_FILE}.md"
    fi
fi
