# atlas-sdk-go
[![PkgGoDev](https://pkg.go.dev/badge/go.mongodb.org/atlas-sdk)](https://pkg.go.dev/go.mongodb.org/atlas-sdk)
![CI](https://github.com/mongodb/atlas-sdk-go/workflows/CI/badge.svg)

A Go SDK for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

Note that `atlas-sdk-go` only supports the two most recent major versions of Go.

## Getting started

### Adding Dependency

```
go install go.mongodb.org/atlas-sdk@v0.16.0
```

### Using in the code

```go
import "go.mongodb.org/atlas-sdk/admin"
```

Construct a new Atlas SDK client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
import "go.mongodb.org/atlas-sdk/admin"

apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
```

For documentation about obtaining apiKey and apiSecret go to
https://docs.atlas.mongodb.com/configure-api-access.

The services of a client divide the API into logical chunks and correspond to
the structure of the Atlas API documentation at
https://www.mongodb.com/docs/atlas/reference/api-resources-spec/.

## Documentation

Please refer to the [docs](./docs)

## Examples

See [examples](./examples)

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

## License

`atlas-sdk-go` is released under the Apache 2.0 license. See [LICENSE](LICENSE)
