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
go install go.mongodb.org/atlas
```

### Using in the code

```go
mongodbatlas import "go.mongodb.org/atlas/mongodbatlasv2"
```

Construct a new Atlas client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
	mongodbatlas import "go.mongodb.org/atlas/mongodbatlasv2"

   	apiKey := os.Getenv("MDB_API_KEY")
	apiSecret := os.Getenv("MDB_API_SECRET")

	sdk := mongodbatlas.NewClient(mongodbatlas.UseDigestAuth(apiKey, apiSecret))
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
```

For documentation about obtaining private and public API token go to
https://docs.atlas.mongodb.com/configure-api-access.
The services of a client divide the API into logical chunks and correspond to
the structure of the Atlas API documentation at
https://www.mongodb.com/docs/atlas/reference/api-resources-spec/.

## Examples

Example for creating an dedicated MongoDB cluster on AWS infrastructure

```
go run ./examples/example_cluster_aws.go
```

### Authentication

The `atlas-sdk-go` library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle Digest Access authentication for
you. The easiest way to do this is using the [digest](https://github.com/mongodb-forks/digest)
library, but you can always use any other library that provides an `http.Client`.
If you have a private and public API token pair (https://docs.atlas.mongodb.com/configure-api-access),
you can use it with the digest library like:

```go
import (
    "context"
    "log"

    "github.com/mongodb-forks/digest"
    "go.mongodb.org/atlas/mongodbatlas"
)

func main() {
    t := digest.NewTransport("your public key", "your private key")
    tc, err := t.Client()
    if err != nil {
        log.Fatalf(err.Error())
    }

    client := mongodbatlas.NewClient(tc)
    projects, _, err := client.Projects.GetAllProjects(context.Background(), nil)
}
```

Note that when using an authenticated Client, all calls made by the client will
include the specified tokens. Therefore, authenticated clients should
almost never be shared between different users.

## Error handling

Fetching error code:
```go
import errors "go.mongodb.org/atlas/mongodbatlasv2"

apiError := errors.AsError(err)
fmt.Println(apiError)
```

Checking if error code exists:
```go
import errors "go.mongodb.org/atlas/mongodbatlasv2"


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
