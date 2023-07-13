# ClusterDescriptionConnectionStringsPrivateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | Pointer to **string** | Private endpoint-aware connection string that uses the &#x60;mongodb://&#x60; protocol to connect to MongoDB Cloud through a private endpoint. | [optional] [readonly] 
**Endpoints** | Pointer to [**[]ClusterDescriptionConnectionStringsPrivateEndpointEndpoint**](ClusterDescriptionConnectionStringsPrivateEndpointEndpoint.md) | List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].connectionString** or **connectionStrings.privateEndpoint[n].srvConnectionString**. | [optional] [readonly] 
**SrvConnectionString** | Pointer to **string** | Private endpoint-aware connection string that uses the &#x60;mongodb+srv://&#x60; protocol to connect to MongoDB Cloud through a private endpoint. The &#x60;mongodb+srv&#x60; protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don&#39;t need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your application supports it. If it doesn&#39;t, use connectionStrings.privateEndpoint[n].connectionString. | [optional] [readonly] 
**SrvShardOptimizedConnectionString** | Pointer to **string** | Private endpoint-aware connection string optimized for sharded clusters that uses the &#x60;mongodb+srv://&#x60; protocol to connect to MongoDB Cloud through a private endpoint. If the connection string uses this Uniform Resource Identifier (URI) format, you don&#39;t need to change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your application and Atlas cluster supports it. If it doesn&#39;t, use and consult the documentation for connectionStrings.privateEndpoint[n].srvConnectionString. | [optional] [readonly] 
**Type** | Pointer to **string** | MongoDB process type to which your application connects. Use &#x60;MONGOD&#x60; for replica sets and &#x60;MONGOS&#x60; for sharded clusters. | [optional] [readonly] 

## Methods

### NewClusterDescriptionConnectionStringsPrivateEndpoint

`func NewClusterDescriptionConnectionStringsPrivateEndpoint() *ClusterDescriptionConnectionStringsPrivateEndpoint`

NewClusterDescriptionConnectionStringsPrivateEndpoint instantiates a new ClusterDescriptionConnectionStringsPrivateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDescriptionConnectionStringsPrivateEndpointWithDefaults

`func NewClusterDescriptionConnectionStringsPrivateEndpointWithDefaults() *ClusterDescriptionConnectionStringsPrivateEndpoint`

NewClusterDescriptionConnectionStringsPrivateEndpointWithDefaults instantiates a new ClusterDescriptionConnectionStringsPrivateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.
### GetEndpoints

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetEndpoints() []ClusterDescriptionConnectionStringsPrivateEndpointEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetEndpointsOk() (*[]ClusterDescriptionConnectionStringsPrivateEndpointEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) SetEndpoints(v []ClusterDescriptionConnectionStringsPrivateEndpointEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.
### GetSrvConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetSrvConnectionString() string`

GetSrvConnectionString returns the SrvConnectionString field if non-nil, zero value otherwise.

### GetSrvConnectionStringOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetSrvConnectionStringOk() (*string, bool)`

GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) SetSrvConnectionString(v string)`

SetSrvConnectionString sets SrvConnectionString field to given value.

### HasSrvConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) HasSrvConnectionString() bool`

HasSrvConnectionString returns a boolean if a field has been set.
### GetSrvShardOptimizedConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetSrvShardOptimizedConnectionString() string`

GetSrvShardOptimizedConnectionString returns the SrvShardOptimizedConnectionString field if non-nil, zero value otherwise.

### GetSrvShardOptimizedConnectionStringOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetSrvShardOptimizedConnectionStringOk() (*string, bool)`

GetSrvShardOptimizedConnectionStringOk returns a tuple with the SrvShardOptimizedConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvShardOptimizedConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) SetSrvShardOptimizedConnectionString(v string)`

SetSrvShardOptimizedConnectionString sets SrvShardOptimizedConnectionString field to given value.

### HasSrvShardOptimizedConnectionString

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) HasSrvShardOptimizedConnectionString() bool`

HasSrvShardOptimizedConnectionString returns a boolean if a field has been set.
### GetType

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterDescriptionConnectionStringsPrivateEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


