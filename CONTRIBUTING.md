# Contributing to atlas-sdk-go

Thanks for your interest in contributing to `atlas-sdk-go`,
The following is a set of guidelines for contributing to this project.
These are just guidelines, not rules. Use your best judgment, and
feel free to propose changes to this document in a pull request.

Note that `atlas-sdk-go` is an evolving project, so expect things to change over
time as the team learns, listens and refines how we work with the community.

## What should I know before I get started?

### Code of Conduct

This project has adopted the [MongoDB Code of Conduct](https://www.mongodb.com/community-code-of-conduct).
By participating, you are expected to uphold this code.
If you see any violations of the above or have any other concerns or questions please contact us
using the following email alias: [community-conduct@mongodb.com](mailto:community-conduct@mongodb.com).

## How Can I Contribute?

### Reporting Bugs or Suggesting Enhancements

Before reporting a bug or suggesting an enhancements please perform a
**[cursory search](https://github.com/mongodb/atlas-sdk-go/issues)**
to see if the bug or enhancement has already been suggested. If it has, add a
:thumbsup: to indicate your interest in it, or comment if there is additional
information you would like to add.

Bugs and enhancement suggestions are tracked as [GitHub issues](https://guides.github.com/features/issues/).

Simply create an issue on the [project issue tracker](https://github.com/mongodb/atlas-sdk-go/issues/new)
and fill our the form.

Some additional advice:

* **Use a clear and descriptive title**
* **Provide a step-by-step description of the issue**
  This additional context helps the maintainers understand the enhancement from
  your perspective

### Before Submitting A New Pull Request

Before creating a pull request consider opening a new issue first,
please check [this section](#reporting-bugs-or-suggesting-enhancements) on how we track issues for this project.

### Autoclose stale issues and PRs

- After 30 days of no activity (no comments or commits on an issue/PR) we automatically tag it as "stale" and add a message: ```This issue/PR has gone 30 days without any activity and meets the project's definition of "stale". This will be auto-closed if there is no new activity over the next 60 days. If the issue is still relevant and active, you can simply comment with a "bump" to keep it open, or add the label "not_stale". Thanks for keeping our repository healthy!```
- After 60 more days of no activity we automatically close the issue/PR.

## Maintainer's Guide

Reviewers, please ensure that the CLA has been signed by referring to [the contributors tool](https://contributors.corp.mongodb.com/) (internal link).

### Development setup

#### Prerequisite Tools
- [Git](https://git-scm.com/)
- [Golang (at least Go 1.21)](https://golang.org/dl/)
    - Repository uses [asdf](https://asdf-vm.com) for golang version management
- Run `make tools` to install all required binaries

#### Environment
- Fork the repository.
- Clone your forked repository locally.
- We use Go Modules to manage dependencies, so you can develop outside of your `$GOPATH`.
- We use [golangci-lint](https://github.com/golangci/golangci-lint) to lint our code, you can install it locally via `make tools`.
- For pull requests to be accepted, contributors must sign [MongoDB CLA](https://www.mongodb.com/legal/contributor-agreement).

### Building and testing

The following is a short list of commands that can be run in the root of the project directory:

- Run `make test` to run all unit tests.
- Run `make lint` to validate against our linting rules.

We provide a git pre-commit hook to format and check the code, to install it run `make link-git-hooks`

## Repository structure

Repository contains following SDKs

- `auth`: Atlas Auth SDK
- `admin`: Atlas Admin SDK
- `tools`: tooling used to generate Atlas Admin SDK. See [./tools](./tools) for more information.
- `examples`: SDK examples
- `internal`: folder for hosting non user facing SDK methods and core helpers
- `mockadmin`: [Testify mocks](https://pkg.go.dev/github.com/stretchr/testify/mock) generated with [Mockery](https://github.com/vektra/mockery) for API interfaces

## VSCode debugging configuration
VSCode developers might use following debugging configuration for transformer and go debugging
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Go App",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${file}"
        },
        {
            "type": "node",
            "request": "launch",
            "name": "Transform",
            "skipFiles": [
                "<node_internals>/**"
            ],
            "env": {
                "OPENAPI_FILE_NAME": "./openapi/atlas-api.yaml"
            },
            "args": ["./openapi/atlas-api.yaml"],
            "program": "${workspaceFolder}/tools/transformer/src/transform"
        }
    ]
}

```

## Documentation

Documentation can be regenerated using `make gen-docs`

### Folder Structure

- `./docs` - mostly manual documentation (only README.md and reference is generated)
- `./docs/docs/`- generated documentation for each SDK API

### Changes in manual documentation

We recommend to create separate PR containing only documentation files each time manual documentation is changed.
Documentation team will be assigned as reviewer automatically and will help with review of the PR.


## Third Party Dependencies and Vulnerability Scanning

We scan our dependencies for vulnerabilities and incompatible licenses using [Snyk](https://snyk.io/).
To run Snyk locally please follow their [CLI reference](https://support.snyk.io/hc/en-us/articles/360003812458-Getting-started-with-the-CLI).

We also use Kondukto to scan for third-party dependency vulnerabilities. Kondukto creates tickets in MongoDB's issue tracking system for any vulnerabilities found.

### SBOM and Compliance
We generate Software Bill of Materials (SBOM) files for each release as part of MongoDB's SSDLC initiative. SBOM Lite files are automatically generated and included as release artifacts. Compliance reports are generated after each release and stored in the compliance/<release-version> directory.

Augmented SBOMs can be generated on customer request for any released version. This can only be done by MongoDB employees as it requires access to our GitHub workflow.

### Papertrail Integration
All releases are recorded using a MongoDB-internal application called Papertrail. This records various pieces of information about releases, including the date and time of the release, who triggered the release (by pushing to Evergreen), and a checksum of each release file.

This is done automatically as part of the release.
