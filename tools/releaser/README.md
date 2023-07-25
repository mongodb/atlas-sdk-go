# Go SDK Release Pipeline

## Context

Please refer to the [end user versioning guide](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_1_concepts.md#release-strategy-semantic-versioning)

## Release steps

Release is done as part of the 2 steps:

1. Automatic bump as part of the OpenAPI generation process that creates API bot update PR.
2. Release tag when update PR is merged

## Disabling automatic release

In situations where we do not want to perform automatic release please remove changes from `./internal/core/version.go` file.

## Manual Minor Release

Manual release is possible by updating `./internal/core/version.go` file with minor version bump.

## Manual Major Release

For major releases (breaking changes) introduced outside new Resource Version.

1. Please update `./internal/core/version.go` file with major version bump.
2. Run `make update-version` command

## Automation Process Internals

Automation reads Resource Version in the `version.go` file and in the OpenAPI `versions.json` 

### New Resource Version 

For new Resource Version introduced in the versions.json
- Reset minor version
- Replace all occurences of the Major version in the repository
- Update versions in versions.go

### New Generation

- Issue new Minor release
- Update versions in versions.go

### Manual Breaking Changes Release

- Read versions in versions.go
- Replace all occurences of the Major version in the repository
