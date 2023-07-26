#!/bin/bash
set -ueo pipefail
 
GITHUB_ENV=${GITHUB_ENV:-.env}

## Scripts sets environemnt for github actions

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"

## Set the environment variables used 
export SDK_VERSION=${SDK_VERSION}
export HYPEN_RESOURCE_VERSION=${HYPEN_RESOURCE_VERSION}

echo "SDK_VERSION=${SDK_VERSION}" >> $GITHUB_ENV
echo "HYPEN_RESOURCE_VERSION=${HYPEN_RESOURCE_VERSION}" >> $GITHUB_ENV
echo "RELEASE_TAG=${SDK_VERSION}" >> $GITHUB_ENV

EOF=$(dd if=/dev/urandom bs=15 count=1 status=none | base64)

RELEASE_NOTES=$(envsubst < "$script_path/../templates/RELEASE_NOTES.tmpl")
RELEASE_NOTES=$RELEASE_NOTES\n$(cat $script_path/../breaking_changes/${SDK_MAJOR_VERSION}.md)

## Multiline string
echo "RELEASE_NOTES<<$EOF" >> $GITHUB_ENV
echo $RELEASE_NOTES >> $GITHUB_ENV
echo "$EOF" >> "$GITHUB_ENV"
