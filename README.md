# atlas-sdk-go
[![PkgGoDev](https://pkg.go.dev/badge/go.mongodb.org/atlas-sdk)](https://pkg.go.dev/go.mongodb.org/atlas-sdk)
![CI](https://github.com/mongodb/atlas-sdk-go/workflows/CI/badge.svg)
![Go](https://img.shields.io/github/go-mod/go-version/mongodb/atlas-sdk-go)

A Go SDK for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

Note that `atlas-sdk-go` only supports the two most recent major versions of Go.

## Getting started

### Adding Dependency

```terminal
go get go.mongodb.org/atlas-sdk/v20240805002
```

### Using in the code

Construct a new Atlas SDK client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
import "go.mongodb.org/atlas-sdk/v20240805002/admin"

func example() {
	ctx := context.Background()

	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

	sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
	if err != nil {
		log.Fatalf("Error when instantiating new client: %v", err)
	}
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	if err != nil {
		log.Fatalf("Could not fetch projects: %v", err)
	}
	fmt.Printf("Response status: %v\n", response.Status)
}
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

## Test Support

[Testify Mock](https://pkg.go.dev/github.com/stretchr/testify/mock) files are generated using [Mockery](https://github.com/vektra/mockery) to help to unit test clients using Atlas Go SDK.

See [test example with mocks](./examples/mock/cluster_test.go)

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

## License

`atlas-sdk-go` is released under the Apache 2.0 license. See [LICENSE](LICENSE)
