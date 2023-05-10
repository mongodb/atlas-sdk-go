# atlas-sdk-go
[![PkgGoDev](https://pkg.go.dev/badge/go.mongodb.org/atlas-sdk)](https://pkg.go.dev/go.mongodb.org/atlas-sdk)
![CI](https://github.com/mongodb/atlas-sdk-go/workflows/CI/badge.svg)

> NOTE: THIS REPOSITORY IS NOT READY FOR PUBLIC CONSUMPTION.
> For official Atlas Golang SDK please go to https://github.com/mongodb/go-client-mongodb-atlas

A Go SDK for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

Note that `atlas-sdk-go` only supports the two most recent major versions of Go.

## Usage

### Adding Dependency

```
go install go.mongodb.org/atlas-sdk
```

### Using in the code

```go
import "go.mongodb.org/atlas-sdk/admin"
```

Construct a new Atlas SDK client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
import "go.mongodb.org/atlas-sdk/admin"

apiKey := os.Getenv("MDB_API_KEY")
apiSecret := os.Getenv("MDB_API_SECRET")

sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
```

For documentation about obtaining private and public API token go to
https://docs.atlas.mongodb.com/configure-api-access.
The services of a client divide the API into logical chunks and correspond to
the structure of the Atlas API documentation at
https://www.mongodb.com/docs/atlas/reference/api-resources-spec/.

## Examples

Example for creating an dedicated MongoDB cluster on AWS infrastructure

```bash
go run ./examples/example_cluster_aws.go
```

## Authentication

The `atlas-sdk-go` library uses Digest authentication. 
To obtain authentication tokens users can use Atlas UI or Atlas CLI 
For more information please follow: https://www.mongodb.com/docs/atlas/api/api-authentication,

## Error Handling

SDK enables users to obtain detailed information about errors returned from backend.
Errors are represented by [ErrorObject](./admin/model_error.go).
Users should rely on the error code for detection of specific error cases.

### Fetching Error Object
```go
import errors "go.mongodb.org/atlas-sdk/admin"

projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
apiError := errors.AsError(err)
fmt.Println(apiError)
```

### Checking for existence of specific error code
```go
import errors "go.mongodb.org/atlas-sdk/admin"


if errors.IsErrorCode(err, "code"){
 // Do something
}
```

## Advanced usage

###  Fluent and Struct Based API

Generated client support two different ways to execute API requests.
1. Fluent API: where users are supplying arguments using chain of functions (default)
2. Struct API: Suppling request data using an single go structure containing request body and arguments

#### Fluent API example

Fluent API should be used by default to handle all requests.

```go
    projects, response, err := sdk.ProjectsApi.ListProjects(ctx).
	    ItemsPerPage(1).Execute()
```  

#### Struct based API example

Struct based API is particularly useful for HTTP GET requests where we need to pass number of arguments to the function without checking 
```go
	listParams := &mongodbatlas.ListProjectsApiParams{ItemsPerPage: mongodbatlas.PtrInt32(1)}
	projects, response, err := sdk.ProjectsApi.ListProjectsWithParams(ctx, listParams).Execute()
```    

> NOTE: Struct based API is an still experimental feature.

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

## License

`atlas-sdk-go` is released under the Apache 2.0 license. See [LICENSE](LICENSE)
