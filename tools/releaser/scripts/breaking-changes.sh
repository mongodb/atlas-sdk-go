#!/bin/bash
set +e
GOPATH=$(go env GOPATH)

# Inputs:
# GIT_BASE_REF - The base REF of the git repository (git rev-parse origin/main)
# Usually "${{ github.event.pull_request.base.sha || github.event.merge_group.base_sha }}" 
# TARGET_BREAKING_CHANGES_FILE - file to save breaking changes
script_path=$(dirname "$0")

echo "Installing go-apidiff"
go install github.com/joelanford/go-apidiff@latest > /dev/null

GIT_BASE_REF=${GIT_BASE_REF:-git rev-parse head || echo}

echo "Running breaking changes check for $GIT_BASE_REF"

pushd "$script_path/../../" || exit ## workaround for --repo-path="../" not working
BREAKING_CHANGES=$("$GOPATH/bin/go-apidiff" "$GIT_BASE_REF" --compare-imports="false" --print-compatible="false")
popd || exit

if [ -z "$BREAKING_CHANGES" ]; then
  echo "No breaking changes detected"
else
  if [ -z "$TARGET_BREAKING_CHANGES_FILE" ]; then
    echo "Breaking changes detected"
    echo "$BREAKING_CHANGES"
  else
    echo "Creating breaking changes file"
    echo -e "# Breaking Changes \n ## SDK changes\n$BREAKING_CHANGES\n## API Changelog\n https://www.mongodb.com/docs/atlas/reference/api-resources-spec/changelog" \
		> "$script_path/../breaking_changes/${TARGET_BREAKING_CHANGES_FILE}.md"
  fi
fi
