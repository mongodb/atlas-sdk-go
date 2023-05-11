# Atlas-SDK-GO Migration Guide

This document is meant to help you migrate from [go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas) to [mongodb/atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go).

## Summary

End users can migrate from [mongodb/go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas) to [mongodb/atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go) partially
by using both librarires at the same time. 

An [go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas) will no longer receive major feature updates.
We strongly recommend updating to [mongodb/atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go) for the latest changes.

## Background

[mongodb/atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go) is based on Atlas V2 API. 
Atlas-sdk-go hides a complexity related to Versioned API by exposing versioned API as golang methods.
End users need to be aware that major sdk releases can introduce breaking changes only in the small subset of the Atlas APIs.

## Structural changes

SDK have been rewriten completely from scratch. 
Our team made attempts to minimize amount of changes required for the end users. 
brings a number of changes in how API requests are made. 

### Client Initialization

New SDK has different methods for the initialization of the clients. 

```go
import admin "go.mongodb.org/atlas-sdk/admin" 
sdk, err := admin.NewClient(
    // Authentication using ApiKey and ApiSecret
    admin.UseDigestAuth(apiKey, apiSecret))
```

> NOTE: Both SDKs use Digest based authentication. The same credentials will apply. 

Please follow [authentication guide](https://github.com/mongodb/atlas-sdk-go#authentication) for more information.

### Error handling 

Error handling requires developers to use dedicated methods for casting errors to APIError objects.

```go
    apiErr, _ := admin.AsError(err)
    log.Fatalf("Error when performing SDK request: %v", apiErr.GetDetail())
```
Please follow [error handling guide](https://github.com/mongodb/atlas-sdk-go#error-handling) for more information.

### Format of the API interface

API interface has changed to differentiate APIs from other methods in the SDK.  
Each API method has an Api suffix. For example:

`sdk.Projects` will now `sdk.ProjectsApi`

Each method now explains the object that is created. For example:

`sdk.Projects.create()` will become `sdk.ProjectsApi.createProject(...)`

Please refer to the [endpoint documentation](https://github.com/mongodb/go-client-mongodb-atlas/tree/main/mongodbatlasv2#documentation-for-api-endpoints) for more information.

### Different naming conventions for SDK methods

Model names and properties are formatted in PascalCase format for clarity and predictability of the methods and field names. 
For example, _ClusterAWSProviderSettings_ will become now _ClusterAwsProviderSettings_.  

The same applies to property names:

For example, `ID` fields will become `Id` etc. 

### Multiple choices when creating request body objects

New SDK improves clarity for request and response objects. For situations when the endpoint accepts multiple formats of the payload (polymorphism), users are able to specify instances of the Api models they want to use for a particular request. For example, when creating a cluster we can use one of the dedicated RegionConfigs objects (AWSRegionConfig, GCPRegionConfig, etc.): 


```go
//...
RegionConfig{
 	// Dedicated region config for AWS cloud
    AWSRegionConfig: &mongodbatlas.AWSRegionConfig{
        //AWS-specific fields are here
        RegionName:   &regionName, 
    },
}
//...
```
