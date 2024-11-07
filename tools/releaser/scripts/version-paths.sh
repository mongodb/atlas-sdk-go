#!/bin/bash

export VERSION_UPDATE_PATHS="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml,../**/*.md"

# don't change doc files to reduce Preview PR noise
export VERSION_UPDATE_PATHS_PREVIEW="../**/go.mod,../**/*.go,../tools/**/*.mustache,../.mockery.yaml"
