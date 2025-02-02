# Contributing OpenAPI Generation Tools

## Changes in tools

We need to regenerate our SDK every time tools are changed as - changes might affect the generated SDK.
PR based checks will verify if changes are up to date.

```bash
make clean_and_generate
```

## OpenAPI generator

OpenAPI generator version is set in [./openapitools.json](./openapitools.json) file.

## Transformer

Follow [transformer readme](./transformer/README.md) for more information

## Modification of the custom templates

Some of the templates were modified to fit for our general needs.
We use mustache comments to mark templates as modified:

- `{{! X-XGEN-CUSTOM }}` - usually present at the top of the file - means that the whole file is currently customized.
- `{{! X-XGEN-MODIFIED }}` - is put into existing templates to note areas in the code that were customized.

## Generating custom files

Tools enable us to generate custom files using OpenAPI model.
We can for example generate unit tests or other helpers that are required.
To generate custom file add custom template edit config.yaml:

```yaml
files:
  myclient.mustache:
    templateType: SupportingFiles
    destinationFilename: yourresultfilename.go
```

## Custom templates

By default `./tools/config/go-templates` folder have all available templates as part of the openapi generator.
This enables us to verify new templates changes.
Templates are sourced from openapi generator:

https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go

## gorelease

We use [gorelease](https://pkg.go.dev/golang.org/x/exp/cmd/gorelease) to detect if there are breaking changes and a new major version release is needed.
In [Generate SDK Github action](../.github/workflows/autoupdate-prod.yaml), the last release is used for comparison.
