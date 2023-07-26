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

RELEASE_NOTES=$(envsubst < "$script_path/../templates/RELEASE_NOTES.tmpl")
RELEASE_NOTES=$RELEASE_NOTES\n$(cat $script_path/../breaking_changes/${SDK_MAJOR_VERSION}.md)

echo "RELEASE_NOTES=${RELEASE_NOTES}" >> $GITHUB_ENV
echo $RELEASE_NOTES
