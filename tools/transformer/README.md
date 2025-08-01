## Atlas SDK Transformation Engine

## Why transforming OpenAPI

OpenAPI file returned by major cloud providers is pretty generic to suit the needs of various tools and projects.
When generating SDK we want to typically finetune the experience to:

- Avoid unnecessary breaking changes
- Flatten structure when overly complex (oneOf, allOf)
- Simplify models - names etc.

The transformation engine enables us to perform that actions by processing OpenAPI file for each generated client.
Additionally, framework enables us to extend it by additional rules per usage to finetune the experience as we wish.

## When to add transformation

Transformations should be generic - they should be reusable across OpenAPI files and be generally an unopinionated way of
handling OpenAPI file. That means that they would improve any public OpenAPI for.
Parameters for transformations (actual scripts) can be more finetuned to current OpenAPI file.

## Usage

1. Run the transformation engine

```
OPENAPI_FILE=openapi.yaml node ./transform.js
```

> NOTE: the provided file will be overwritten
> NOTE: engine supports yaml openapi files

## Reviewing and changing parameters for transformations

See [atlas transformations](./src/atlasTransformations.js) to review transformations
that are applied to atlas.

> NOTE: Order of the transformations matters.

## Implemented transformations

1. oneOf transformation
   Applied to all objects that meet either of the following criteria:

- Have "oneOf" field

For parent model containing multiple children:

- Moves all child enum or property fields into parent properties.
- Removes redundant oneOf fields on parent

2. allOf transformation.

Applied to all objects that meet the following criteria:

> Have a "properties" field
> Have the "oneOf" field

For parent model containing multiple children:

- Moves all parent property fields into children.
- Removes redundant allOf fields on children

3. Discriminator transformation

Applied to all objects that meet the following criteria:

> Have a "discriminator" field
> Missing "oneOf" field

For the parent model containing discriminator

- Fails if discriminator is missing mapping field (invalid case)
- Uses discriminator mapping in order to add oneOf field

4. Model name transformation

> components.schemas is has models

For each model:

- Remove prefix or suffix from model name
- Ignore if model name is not matching prefix or suffix

5. Remove enums transformation

For each model we will effectively remove all enum instances.
Enums in current form are hard to ensure backwards compatibility.

5. Remove nullable transformation

This transformation removes all openapi `nullable` fields from the schema.
Useful for situations for schemas where nullability handling is not desired.

6. Apply Operation ID Overrides

This transformation updates the operation IDs in the OpenAPI file if the operation has an `x-xgen-operation-id-override` extension. Then, it removes the Operation ID override extension from the OpenAPI file.

## Transformation Validation

Transformation engine does perform validation for invalid cases.
If transformation fails it usually means that OpenAPI file is not correct and has issues that require human attention.

The process of fixing OpenAPI always involves:

1. Finding the path of the failure in the generated OpenAPI file
2. Typical errors print path to the OpenAPI file that caused the issue:

For example:

```
   paths.'/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports'(post).requestBody.content is missing
```

3. We look into the line number of the URL path and search for it in the OpenAPI file. In the yaml section of our file, we need to find operationId a property.

## Known problems

1. "Missing object reference"

One of the Java annotations has a hardcoded $ref value that points to an object that is missing. Please review your changes for potential invalid references.

- $ref: "#/components/schemas/InvalidObject"

2. Missing oneOf

"OpenAPI object schema contains discriminator but missing oneOf or discriminator mapping. Please consider adding a oneOf or discriminator mapping section to the object"
Discriminator needs a mapping section and oneOf to reflect the proper inheritance structure
discriminator:
mapping:
...
oneOf:

3. Recursive reference in mapping
   "discriminator.mapping contains $ref to itself"
   Discriminator base objects cannot be mapped to themselves. This causes circular dependencies in many open-source tools

```
       ReplicaSet:
             discriminator:
                       mapping:
                             SET: "#/components/schemas/ReplicaSet"
```

## Adjusting logging

`XGEN_LOGGING_LEVEL` env variable can be used to specify logging levels:

- info
- warn
- debug
- error

Default: 'warn"

## Model ignore

Some transformations have ability to work with ignore list.
When writing transformation developers can use `isModelIgnored` function.
See [Name transformation](./src/transformations/name.js) for example.
