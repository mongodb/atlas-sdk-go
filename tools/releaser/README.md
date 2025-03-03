# Go SDK Release Pipeline

## Context

Please refer to the [end user versioning guide](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_1_concepts.md#release-strategy-semantic-versioning)

## Release steps

The release is done as part of the 2 steps:

1. Automatic bump as part of the OpenAPI generation process that creates API bot update PR.
2. Release tag when update PR is merged

## Disabling automatic release

In situations where we do not want to perform automatic release please remove changes from `./internal/core/version.go` file.

## Manual Minor Release

Manual release is possible by updating `./internal/core/version.go` file with minor version bump.

## Automation Process Internals

Automation reads Resource Version in the `version.go` file and in the OpenAPI `versions.json`

### New Resource Version

For the new Resource Version introduced in the versions.json

- Reset the minor version
- Replace all occurrences of the Major version in the repository
- Update versions in versions.go
- Generate breaking changes file

### New SDK Generation without change in Rersource version

- Issue new Minor release
- Update versions in versions.go

### Breaking Changes Release

In situations when breaking changes are present automation will recognize them and automatically create major version bump.

- Automatically bump the Major version in versions.go
- Replace all occurrences of the Major version in the repository
- Generate breaking changes file

### Manual Breaking Changes Release

When introducing breaking changes outside generation process we need to manually run script:

1. Run

`make update-version`

2. Add breaking changes information to `./breaking _changes/{major_version}`
3. Run `npm run format` to ensure that file contains correct release notes.
4. Commit all the changes.

## Release Notes

If the release contains breaking changes (it is a major release)
we should provide a list of breaking changes notes under `./breaking _changes/{major_version}.md` file.
This file is automatically generated during release.

## SDK Preview

Preview generates SDK version by substituting `github.com/mongodb/atlas-sdk-go` as module path.
Version should never be committed to the main branch. This task should only be used in [Generate Preview SDK Github action](../../.github/workflows/autoupdate-preview.yaml) which creates a PR with latest changes.

## Folder structure

`./scripts` - release scripts. Please do not edit them without testing them on the fork repository.
`./breaking_changes` - folder containing a list of breaking changes. It is generally safe to prune old files in that folder as values are used during the release.
