# AvailableClustersDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentVersion** | Pointer to **string** | Version of MongoDB Agent that monitors/manages the cluster. | [optional] [readonly] 
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster. | [optional] [readonly] 
**DbSizeBytes** | Pointer to **int64** | Size of this database on disk at the time of the request expressed in bytes. | [optional] [readonly] 
**FeatureCompatibilityVersion** | **string** | Version of MongoDB [features](https://docs.mongodb.com/manual/reference/command/setFeatureCompatibilityVersion) that this cluster supports. | [readonly] 
**Managed** | **bool** | Flag that indicates whether Automation manages this cluster. | [readonly] 
**MongoDBVersion** | **string** | Version of MongoDB that this cluster runs. | [readonly] 
**Name** | **string** | Human-readable label that identifies this cluster. | [readonly] 
**OplogSizeMB** | Pointer to **int** | Size of the Oplog on disk at the time of the request expressed in MB. | [optional] [readonly] 
**Sharded** | **bool** | Flag that indicates whether someone configured this cluster as a sharded cluster.  - If &#x60;true&#x60;, this cluster serves as a sharded cluster. - If &#x60;false&#x60;, this cluster serves as a replica set. | [readonly] 
**ShardsSize** | Pointer to **int** | Number of shards that comprise this cluster. | [optional] [readonly] 
**TlsEnabled** | **bool** | Flag that indicates whether someone enabled TLS for this cluster. | [readonly] 

## Methods

### NewAvailableClustersDeployment

`func NewAvailableClustersDeployment(featureCompatibilityVersion string, managed bool, mongoDBVersion string, name string, sharded bool, tlsEnabled bool, ) *AvailableClustersDeployment`

NewAvailableClustersDeployment instantiates a new AvailableClustersDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableClustersDeploymentWithDefaults

`func NewAvailableClustersDeploymentWithDefaults() *AvailableClustersDeployment`

NewAvailableClustersDeploymentWithDefaults instantiates a new AvailableClustersDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentVersion

`func (o *AvailableClustersDeployment) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *AvailableClustersDeployment) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *AvailableClustersDeployment) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *AvailableClustersDeployment) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetClusterId

`func (o *AvailableClustersDeployment) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AvailableClustersDeployment) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AvailableClustersDeployment) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *AvailableClustersDeployment) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetDbSizeBytes

`func (o *AvailableClustersDeployment) GetDbSizeBytes() int64`

GetDbSizeBytes returns the DbSizeBytes field if non-nil, zero value otherwise.

### GetDbSizeBytesOk

`func (o *AvailableClustersDeployment) GetDbSizeBytesOk() (*int64, bool)`

GetDbSizeBytesOk returns a tuple with the DbSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbSizeBytes

`func (o *AvailableClustersDeployment) SetDbSizeBytes(v int64)`

SetDbSizeBytes sets DbSizeBytes field to given value.

### HasDbSizeBytes

`func (o *AvailableClustersDeployment) HasDbSizeBytes() bool`

HasDbSizeBytes returns a boolean if a field has been set.

### GetFeatureCompatibilityVersion

`func (o *AvailableClustersDeployment) GetFeatureCompatibilityVersion() string`

GetFeatureCompatibilityVersion returns the FeatureCompatibilityVersion field if non-nil, zero value otherwise.

### GetFeatureCompatibilityVersionOk

`func (o *AvailableClustersDeployment) GetFeatureCompatibilityVersionOk() (*string, bool)`

GetFeatureCompatibilityVersionOk returns a tuple with the FeatureCompatibilityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCompatibilityVersion

`func (o *AvailableClustersDeployment) SetFeatureCompatibilityVersion(v string)`

SetFeatureCompatibilityVersion sets FeatureCompatibilityVersion field to given value.


### GetManaged

`func (o *AvailableClustersDeployment) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *AvailableClustersDeployment) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *AvailableClustersDeployment) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetMongoDBVersion

`func (o *AvailableClustersDeployment) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *AvailableClustersDeployment) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *AvailableClustersDeployment) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.


### GetName

`func (o *AvailableClustersDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableClustersDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableClustersDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetOplogSizeMB

`func (o *AvailableClustersDeployment) GetOplogSizeMB() int`

GetOplogSizeMB returns the OplogSizeMB field if non-nil, zero value otherwise.

### GetOplogSizeMBOk

`func (o *AvailableClustersDeployment) GetOplogSizeMBOk() (*int, bool)`

GetOplogSizeMBOk returns a tuple with the OplogSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogSizeMB

`func (o *AvailableClustersDeployment) SetOplogSizeMB(v int)`

SetOplogSizeMB sets OplogSizeMB field to given value.

### HasOplogSizeMB

`func (o *AvailableClustersDeployment) HasOplogSizeMB() bool`

HasOplogSizeMB returns a boolean if a field has been set.

### GetSharded

`func (o *AvailableClustersDeployment) GetSharded() bool`

GetSharded returns the Sharded field if non-nil, zero value otherwise.

### GetShardedOk

`func (o *AvailableClustersDeployment) GetShardedOk() (*bool, bool)`

GetShardedOk returns a tuple with the Sharded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharded

`func (o *AvailableClustersDeployment) SetSharded(v bool)`

SetSharded sets Sharded field to given value.


### GetShardsSize

`func (o *AvailableClustersDeployment) GetShardsSize() int`

GetShardsSize returns the ShardsSize field if non-nil, zero value otherwise.

### GetShardsSizeOk

`func (o *AvailableClustersDeployment) GetShardsSizeOk() (*int, bool)`

GetShardsSizeOk returns a tuple with the ShardsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardsSize

`func (o *AvailableClustersDeployment) SetShardsSize(v int)`

SetShardsSize sets ShardsSize field to given value.

### HasShardsSize

`func (o *AvailableClustersDeployment) HasShardsSize() bool`

HasShardsSize returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *AvailableClustersDeployment) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *AvailableClustersDeployment) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *AvailableClustersDeployment) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


