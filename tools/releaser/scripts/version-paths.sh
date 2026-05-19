#!/bin/bash

export VERSION_UPDATE_PATHS="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml,../**/*.md"
# Historical records that must preserve their original version references.
# replace-in-file CLI does not split comma-separated --ignore values; each pattern needs its own flag.
export VERSION_IGNORE_FLAGS="--ignore=../docs/mig_*.md --ignore=../tools/releaser/breaking_changes/*.md"

# don't change doc files to reduce Preview PR noise
export VERSION_UPDATE_PATHS_PREVIEW="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml"
