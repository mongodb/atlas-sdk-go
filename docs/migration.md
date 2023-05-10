# Migrating to Atlas-SDK-GO

Atlas SDK GO represents second generation of the Atlas API and clients.
This guide provides migration information for https://github.com/mongodb/go-client-mongodb-atlas users.

## Summary

Atlas SDK v2 is no longer written manually and it is entirely based on the Atlas API OpenAPI file giving users ability to automatically receive all relevant Atlas updates almost instantly after Atlas release.

End users can migrate from [mongodb/go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas) SDK to a new repository. 
If you are looking for integrating SDK into a new project please follow official getting started guide. 

An [go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas) will no longer receive major updates.
We strongly recommend updating to [mongodb/atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go) for the latest changes.


## Required concepts

Apart from being generated from the OpenAPI file, the new SDK is based on Atlas V2 API.  
That API introduces the concept of Resource Version: 
 
**Resource version**: A new version of a Public API resource in the versioned API usually represented as custom (e.g. `application/vnd.atlas.2023-06-01+json`) content type. Version is applied by _Accept _and_ Content-Type _headers.

Atlas-sdk-go hides a complexity related to Versioned API by exposing versioned API as golang methods.
End users need to be aware that major sdk releases can introduce breaking changes only in the small subset of the Atlas APIs.

### Structural changes

SDK v2 brings a number of changes in how API requests are made. 


### Client Initialization

New SDK has different methods for the initialization of the clients. 

```go
import mongodbatlas "go.mongodb.org/atlas/mongodbatlasv2" 
sdk, err := mongodbatlas.NewClient(
    // Authentication using ApiKey and ApiSecret
	mongodbatlas.UseDigestAuth(apiKey, apiSecret))
```

Please follow a detailed [guide](https://github.com/mongodb/atlas-sdk-go#authentication) for more information about authorization.

### Error handling 

Error handling requires developers to use dedicated methods for casting errors to APIError objects.
```go
	apiErr, _ := mongodbatlas.AsError(err)
    log.Fatalf("Error when performing SDK request: %v", apiErr.GetDetail())
```

### Format of the API interface

API interface has changed to differentiate APIs from other methods in the SDK.  
Each API method has an Api suffix. For example:

`sdk.Projects` will become `sdk.ProjectsAPI`

Each method also clearly explains the object we are creating. For example:

`sdk.Projects.create()` will become `sdk.ProjectsApi.createProject(...)`

**Notable differences:**

_AdvancedClusters_ are now _MultiCloudClustersApi_

_EncryptionsAtRest_ are now _EncryptionAtRestUsingCustomerKeyManagementApi_

_AccessListAPIKeys_ are now included in _ProgrammaticAPIKeysApi_

For a full list of the API please refer to the: [endpoint documentation](https://github.com/mongodb/go-client-mongodb-atlas/tree/main/mongodbatlasv2#documentation-for-api-endpoints)


### Different naming conventions for SDK methods

Model names and properties using now PascalCase format for clarity and predictability of the methods and field names. For example, _ClusterAWSProviderSettings_ will become now _ClusterAwsProviderSettings_.  


The same applies to property names

For example, `ID` fields will become `Id` etc. 


### Structural changes 

New SDK improves clarity for request and response objects. For situations when the endpoint accepts multiple formats of the payload, users are able to specify instances of the Api models they want to use for a particular request. For example, when creating a cluster we can use oneOf of the dedicated RegionConfigs objects (AWSRegionConfig, GCPRegionConfig, etc.). For example: 


```go
RegionConfig{
 	// Dedicated region config for AWS cloud
    AWSRegionConfig: &mongodbatlas.AWSRegionConfig{
        //AWS-specific fields are here
        RegionName:   &regionName, 
    },
}
```
