#!/bin/bash
set +e
GOPATH=$(go env GOPATH)

# Inputs:
# GIT_BASE_REF - The base REF of the git repository
# Usually "${{ github.event.pull_request.base.sha || github.event.merge_group.base_sha }}" 
# TARGET_BREAKING_CHANGES_FILE - file to save breaking changes
script_path=$(dirname "$0")

echo "Installing go-apidiff"
go install github.com/joelanford/go-apidiff@latest

echo "Running breaking changes check"
BREAKING_CHANGES=$("$GOPATH/bin/go-apidiff" "$GIT_BASE_REF" --compare-imports="false" --print-compatible="false" --repo-path="../")

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
