# Learn Concepts for the Atlas Go SDK

## Introduction

You can use the Atlas Go SDK to fetch or modify data from the Atlas Admin API.
The Atlas Go SDK is flexible because it accepts many different types of arguments.

## Fetching Data from the Back End

To fetch data, you can execute the `{Api}.{Operation}WithParams()` method:

```go 
	// 1. Calling API method
	request := sdk.ProjectsApi.ListProjectsWithParams(ctx,
		// 2. We passing struct with all optional query parameters to the request
		&admin.ListProjectsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		})

	// 3. You can also supply values in requests using individual methods
	// This can be helpful when passing request objects to other methods. 
	projects, response, err := request.ItemsPerPage(10).Execute()
	examples.HandleErr(err, response)
```

Alternatively, you can use the shorter `{Operation}()` method with a builder pattern to supply all arguments:

```go
    projects, response, err := sdk.ProjectsApi.ListProjects(ctx).
	    ItemsPerPage(1).Execute()
```

Note: The Atlas Go SDK requires path parameters and they must be provided directly in the `{Operation}()` method.

Note: The Atlas Go SDK supplies default values for both query and post objects.


## Performing Data Modification

Use the `{Operation}()` method to perform modifications. For example:


```go
 groupInvitationRequest := *admin.NewGroupInvitationRequest() 
 resp, r, err := sdk.ProjectsApi.CreateProjectInvitation(context.Background(), groupId, &groupInvitationRequest).Execute()
```

## Experimental Methods

Please note that we have some methods marked as experimental, denoted by the [experimental] tag in the operation description. 
This signifies that the method might be changed in the future without compatibility guarantees.

If you encounter any problems with methods marked as experimental, feel free to raise a [Github issue](https://github.com/mongodb/atlas-sdk-go/issues/new/choose) and the team will work to resolve it.

If you belive a method should be marked as stable, feel free to raise a PR appending the method's OperationId to our [operations.stable.json](https://github.com/mongodb/atlas-sdk-go/blob/main/tools/transformer/src/operations.stable.json) set.

## Example

To learn more about using the SDK, see the [basic example](https://github.com/mongodb/atlas-sdk-go/blob/main/examples/basic/basic.go).

## Release Strategy (Semantic Versioning)

Atlas Go SDK is following the Semantic Versioning (SemVer) specification. The versioning scheme used for this SDK is as follows: `vYYYYMMDDXXX.Y.Z`, where:

- `YYYYMMDDXXX` represents the major version number and includes the date of the [Versioned API Resource Version](https://www.mongodb.com/docs/atlas/api/versioned-api-overview/) used in the SDK followed by three digits for other possible breaking changes.
- `Y` represents the minor version number, indicating non-breaking iterations of the same Versioned API Resource.
- `Z` represents the patch version number, indicating non-user-affecting fixes in the SDK.

### Versioning Rules

#### Major Version (vYYYYMMDDXXX.0.0)

A major version increment signifies breaking changes in the SDK. The rules for major version increments are as follows:

1. When the Versioned API version is introduced SDK undergoes breaking changes, the major version must be incremented.
For example `vYYYYMMDD` will become `v20300101`
2. If there are other significant breaking changes in the SDK, not related to the Versioned API, the major version must also be incremented in the last 3 digits.
3. The major version must include the date of the Versioned API and three digits for other possible breaking changes.

#### Minor Version (vYYYYMMDDXXX.Y.0)

A minor version increment represents iterations of the same Versioned API. The rules for minor version increments are as follows:

1. When new features or enhancements are added to the SDK that are backward-compatible with the previous minor version, the minor version must be incremented.

#### Patch Version (vYYYYMMDDXXX.Y.Z)

A patch version increment indicates non-user-affecting fixes and improvements in the SDK. 

## Example Version: v20230201001.0.0

Let's break down the example version `v20230201001.0.0` to understand its components:

- Major version: `v20230201001.0.0`
  - `v`: Indicates the start of the version number.
  - `20230201001`: The date of the Versioned API that this SDK is using, in the format of `YYYYMMDD`.
  - `0`: Reserved for future
  - `0`: Reserved for future
  - `0`: Breaking change within that resource version

In summary, the example version `v20230201001.0.0` represents the initial release of the Golang SDK library for the Versioned API dated February 1, 2023, with no other breaking changes, iterations, or fixes.


