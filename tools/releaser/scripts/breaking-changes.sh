#!/bin/bash
set -eu
GOPATH=$(go env GOPATH)

# Inputs:
# GIT_BASE_REF - The base REF of the git repository (git rev-parse origin/main)
# Usually "${{ github.event.pull_request.base.sha || github.event.merge_group.base_sha }}" 
# TARGET_BREAKING_CHANGES_FILE - file to save breaking changes
script_path=$(dirname "$0")

echo "Installing go-apidiff"
go install github.com/joelanford/go-apidiff@latest > /dev/null

current_ref=$(git rev-parse main || echo)
GIT_BASE_REF=${GIT_BASE_REF:-$current_ref}
TARGET_BREAKING_CHANGES_FILE=${TARGET_BREAKING_CHANGES_FILE:""}

echo "Running breaking changes check for $GIT_BASE_REF"

pushd "$script_path/../../../" || exit ## workaround for --repo-path="../" not working
echo "Changed directory to $(pwd)"
set +e
BREAKING_CHANGES=$("$GOPATH/bin/go-apidiff" "$GIT_BASE_REF" --compare-imports="false" --print-compatible="false" )
set -e
popd || exit
cmd_return_code=$?

echo "Breaking changes finished with $cmd_return_code return code"

if [ -z "$BREAKING_CHANGES" ]; then
    echo "No major breaking changes detected"
    exit $cmd_return_code
fi

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
 
exit $cmd_return_code
