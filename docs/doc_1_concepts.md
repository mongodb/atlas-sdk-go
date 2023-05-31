# Concepts used in Atlas SDK GO

## Introduction

For each API SDK provides number of methods that users can call to fetch or modify data on the backend.
SDK offers flexibility in how various arguments can be passed to the client.

## Fetching Data from the Backend

To fetch data execute `{Api}.{Operation}WithParams()` method:

```go 
	// 1. Calling API method
	request := sdk.ProjectsApi.ListProjectsWithParams(ctx,
		// 2. We passing struct with all optional query parameters to the request
		&admin.ListProjectsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		})

	// 3. Values in requests can be also supplied using individual methods
	// This can be helpful when passing request objects to other methods. 
	projects, response, err := request.ItemsPerPage(10).Execute()
	examples.HandleErr(err, response)
```

Alternatively you can use shorter `{Operation}()` method using builder pattern to supply all arguments:

```go
    projects, response, err := sdk.ProjectsApi.ListProjects(ctx).
	    ItemsPerPage(1).Execute()
```

Note: Path parameters are always required and they will be supplied directly in the `{Operation}()` method:

Note: SDK supplies default values for both query and post objects.


## Performing Data Modification

 Use `{Operation}()` method to perform modifications. For example:


```go
 groupInvitationRequest := *admin.NewGroupInvitationRequest() 
 resp, r, err := sdk.ProjectsApi.CreateProjectInvitation(context.Background(), groupId, &groupInvitationRequest).Execute()
```

## Example

For more information about usage please refer to the [basic example](../examples/)
