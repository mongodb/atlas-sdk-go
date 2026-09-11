#!/bin/bash

export VERSION_UPDATE_PATHS="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml,../**/*.md"
# Historical records that must preserve their original version references.
# replace-in-file CLI does not split comma-separated --ignore values; each pattern needs its own flag.
# Paths are matched against glob results relative to tools/ (the glob cwd), so files under
# tools/ are seen as `releaser/...`, not `../tools/releaser/...`; a `**/` prefix matches both.
export VERSION_IGNORE_FLAGS="--ignore=../docs/mig_*.md --ignore=**/breaking_changes/*.md --ignore=../compliance/**"

# don't change doc files to reduce Preview PR noise
export VERSION_UPDATE_PATHS_PREVIEW="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml"
