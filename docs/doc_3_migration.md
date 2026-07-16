# Migrate from the Go HTTP Client to the Atlas Go SDK

Use this guide to migrate from the Go HTTP client ([go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas)) to the Atlas Go SDK ([mongodb/atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go)).

The [go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas) is deprecated and doesn't receive major feature updates. We strongly recommend migrating to the [Atlas Go SDK](https://github.com/mongodb/atlas-sdk-go) for the latest changes. You can also migrate partially by using both libraries at the same time. 

## Background

The Atlas Go SDK ([mongodb/atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go)) is based on the Atlas Admin API V2. 
The Atlas Go SDK simplifies the complexity of the versioned API by exposing the versioned API as Golang methods. Major SDK releases can introduce breaking changes only in a small subset of the Atlas Admin API endpoints.

## Structural Changes

The Atlas Go SDK doesn't rely on the deprecated [go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas). It changes how API requests are made, but minimizes the changes required for the end users.

### Client Initialization

The Atlas Go SDK has different methods for the initialization of the clients:

```go
// Surrounding code ignored for brevity

sdk, err := admin.NewClient(
    // Authentication using ApiKey and ApiSecret
    admin.UseDigestAuth(apiKey, apiSecret))
```

Note: Both the deprecated [go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas) and the Atlas Go SDK use Digest-based authentication. The same credentials apply. 

To learn more, see [Authenticate using the Atlas Go SDK](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_4_authentication.md).

### Error Handling 

Error handling requires developers to use dedicated methods for casting errors to API error objects:

```go
// Surrounding code omitted for brevity

apiErr, _ := admin.AsError(err)
log.Fatalf("Error when performing SDK request: %v", apiErr.GetDetail())
```
To learn more, see [Error Handling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_2_error_handling.md).

### Format of the API Interface

The Atlas Go SDK changes the API interface to differentiate APIs from other methods.  

Each API method has an API suffix. For example:

`sdk.Projects` will now `sdk.ProjectsApi`

Each method now explains the object that is created. For example:

`sdk.Projects.create()` will become `sdk.ProjectsApi.createProject(...)`

To learn more, see the [Endpoint Documentation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_last_reference.md).

### Different Naming Conventions for SDK Methods

Model names and properties are formatted in PascalCase format for clarity and predictability of the methods and field names. 
For example, _ClusterAWSProviderSettings_ will become now _ClusterAwsProviderSettings_.  

The same applies to property names. For example, `ID` fields will become `Id` etc. 

### Multiple Choices when Creating Request Body Objects

The Atlas Go SDK improves the clarity for request and response objects. For situations when the endpoint accepts multiple formats of the payload (polymorphism), you can specify instances of the API models that you want to use for a particular request. For example, when creating a cluster you can use one of the dedicated RegionConfigs objects (AWSRegionConfig, GCPRegionConfig, etc.): 


```go
// Surrounding code omitted for brevity

RegionConfig{
 	// Dedicated region config for AWS cloud
    AWSRegionConfig: &mongodbatlas.AWSRegionConfig{
        //AWS-specific fields are here
        RegionName:   &regionName, 
    },
}
```
