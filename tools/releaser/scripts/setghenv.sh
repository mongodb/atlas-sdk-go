#!/bin/bash

set -ueo pipefail
 
GITHUB_ENV=${GITHUB_ENV:-.env}

## Scripts sets environemnt for github actions

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"

## Set the environment variables used 
export SDK_VERSION="${SDK_VERSION}"
export HYPEN_RESOURCE_VERSION="${HYPEN_RESOURCE_VERSION}"
{
  echo "SDK_VERSION=${SDK_VERSION}"
  echo "HYPEN_RESOURCE_VERSION=${HYPEN_RESOURCE_VERSION}"
  echo "RELEASE_TAG=${SDK_VERSION}"
} >> "$GITHUB_ENV"

EOF=$(dd if=/dev/urandom bs=15 count=1 status=none | base64)

RELEASE_NOTES=$(envsubst < "$script_path/../templates/RELEASE_NOTES.tmpl")

# Check if this is a major version bump
if [[ "$SDK_VERSION" =~ ^v[0-9]+\.0\.0$ ]]; then
    # For major versions, include breaking changes from the file
    breaking_changes_path="$script_path/../breaking_changes/${SDK_MAJOR_VERSION}.md"
    if [ -f "$breaking_changes_path" ]; then
        echo "Found breaking changes file for $SDK_MAJOR_VERSION"
        BREAKING_CHANGES=$(cat "$breaking_changes_path")
        RELEASE_NOTES=$(echo -e "${RELEASE_NOTES}\n\n${BREAKING_CHANGES}")
    fi
else
    # For minor versions, include non-breaking changes directly
    if [ ! -z "$NON_BREAKING_CHANGES" ]; then
        echo "Adding non-breaking changes to release notes"
        RELEASE_NOTES=$(echo -e "${RELEASE_NOTES}\n\n## Changes\n${NON_BREAKING_CHANGES}")
    fi
fi

## Multiline string
{
  echo "RELEASE_NOTES<<$EOF"
  echo "$RELEASE_NOTES"
  echo "$EOF"
} >> "$GITHUB_ENV"
