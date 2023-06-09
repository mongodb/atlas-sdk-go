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

The Atlas Go SDK follows the Semantic Versioning (SemVer) specification. The versioning scheme used for this SDK is as follows: `vYYYYMMDDXXX.Y.Z`, where:

- `YYYYMMDDXXX` represents the major version number. It includes the date of the [Versioned API Resource Version](https://www.mongodb.com/docs/atlas/api/versioned-api-overview/) that the SDK uses followed by three digits for other possible breaking changes.
- `Y` represents the minor version number, indicating non-breaking iterations of the same Versioned API Resource.
- `Z` represents the patch version number, indicating fixes in the SDK that do not affect users.

### Versioning Rules

#### Major Version (vYYYYMMDDXXX.0.0)

A major version increment signifies breaking changes in the SDK. The rules for major version increments are as follows:

1. The major version must include the date of the Versioned API and three digits for other possible breaking changes.
2. When a new Versioned API Resource Version is introduced, the SDK undergoes breaking changes and the version identifier is incremented. For example, `vYYYYMMDD` will become `v20300101` if a major version was released one year later.
3. If there are other significant breaking changes in the SDK that are unrelated to the Versioned API, increment the last three digits of the major version.


#### Minor Version (vYYYYMMDDXXX.Y.0)

A minor version increment represents the Atlas Go SDK release based on iterations of the targeted Versioned API. When MongoDB adds new features or enhancements to the SDK that are backward-compatible with the previous minor version, increment the minor version.

#### Patch Version (vYYYYMMDDXXX.Y.Z)

A patch version increment indicates fixes and improvements in the SDK that do not affect users.

## Example Version: v20230201001.0.0

Let's break down the example version `v20230201001.0.0` to understand its components:

- Major version: `v20230201001.0.0`
  - `v`: Indicates the start of the version number.
  - `20230201`: The date of the Versioned API version that this SDK is using, in the format of `YYYYMMDD`.
  - `001` starts from 001 and is incremented by 1 for every non-backward-compatible iteration targeting the specific API resource version.

The version `v20230201001.0.0` represents the initial release of the Golang SDK library for the Versioned API dated February 1, 2023, with no other breaking changes, iterations, or fixes.
