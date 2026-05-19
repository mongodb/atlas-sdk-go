#!/bin/bash

export VERSION_UPDATE_PATHS="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml,../**/*.md"
# These files are excluded because they are historical records that intentionally
# reference specific versions; updating them would corrupt the version history they document.
export VERSION_IGNORE_PATHS="../docs/mig_*.md,../tools/releaser/breaking_changes/*.md"

# don't change doc files to reduce Preview PR noise
export VERSION_UPDATE_PATHS_PREVIEW="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml"
