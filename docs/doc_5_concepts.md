# SDK concepts

##  Fluent and Struct Based API

Generated client support two different ways to execute API requests.
1. Fluent API: where users are supplying arguments using chain of functions (default)
2. Struct API: Suppling request data using an single go structure containing request body and arguments

### Fluent API example

Fluent API should be used by default to handle all requests.

```go
    projects, response, err := sdk.ProjectsApi.ListProjects(ctx).
	    ItemsPerPage(1).Execute()
```  

### Struct based API example

Struct based API is particularly useful for HTTP GET requests where we need to pass number of arguments to the function:

```go
	listParams := &admin.ListProjectsApiParams{ItemsPerPage: admin.PtrInt32(1)}
	projects, response, err := sdk.ProjectsApi.ListProjectsWithParams(ctx, listParams).Execute()
```    
