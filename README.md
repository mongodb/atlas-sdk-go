# atlas-sdk-go
[![PkgGoDev](https://pkg.go.dev/badge/go.mongodb.org/atlas-sdk)](https://pkg.go.dev/go.mongodb.org/atlas-sdk)
![CI](https://github.com/mongodb/atlas-sdk-go/workflows/CI/badge.svg)

> NOTE: THIS REPOSITORY IS NOT READY FOR PUBLIC CONSUMPTION.
> For official Atlas Golang SDK please go to https://github.com/mongodb/go-client-mongodb-atlas

A Go SDK for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

Note that `atlas-sdk-go` only supports the two most recent major versions of Go.

## Getting started

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

## Documentation

Please refer to the [docs](./docs)

## Examples

Example for creating an dedicated MongoDB cluster on AWS infrastructure

```bash
go run ./examples/example_cluster_aws.go
```

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

## License

`atlas-sdk-go` is released under the Apache 2.0 license. See [LICENSE](LICENSE)
