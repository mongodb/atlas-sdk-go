# ApiHostViewAtlas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this MongoDB process. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the &#x60;mongod&#x60; or &#x60;mongos&#x60;. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). | [optional] [readonly] 
**Id** | Pointer to **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**LastPing** | Pointer to **time.Time** | Date and time when MongoDB Cloud received the last ping for this MongoDB process. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Links** | Pointer to [**[]LinkAtlas**](LinkAtlas.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Port** | Pointer to **int** | Internet Assigned Numbers Authority (IANA) port on which the MongoDB process listens for requests. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set that contains this process. This resource returns this parameter if this process belongs to a replica set. | [optional] [readonly] 
**ShardName** | Pointer to **string** | Human-readable label that identifies the shard that contains this process. This resource returns this value only if this process belongs to a sharded cluster. | [optional] [readonly] 
**TypeName** | Pointer to **string** | Type of MongoDB process that MongoDB Cloud tracks. MongoDB Cloud returns new processes as **NO_DATA** until MongoDB Cloud completes deploying the process. | [optional] [readonly] 
**UserAlias** | Pointer to **string** | Human-readable label that identifies the cluster node. MongoDB Cloud sets this hostname usually to the standard hostname for the cluster node. It appears in the connection string for a cluster instead of the value of the hostname parameter. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of MongoDB that this process runs. | [optional] [readonly] 

## Methods

### NewApiHostViewAtlas

`func NewApiHostViewAtlas() *ApiHostViewAtlas`

NewApiHostViewAtlas instantiates a new ApiHostViewAtlas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostViewAtlasWithDefaults

`func NewApiHostViewAtlasWithDefaults() *ApiHostViewAtlas`

NewApiHostViewAtlasWithDefaults instantiates a new ApiHostViewAtlas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ApiHostViewAtlas) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiHostViewAtlas) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiHostViewAtlas) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiHostViewAtlas) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetGroupId

`func (o *ApiHostViewAtlas) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiHostViewAtlas) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiHostViewAtlas) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiHostViewAtlas) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetHostname

`func (o *ApiHostViewAtlas) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiHostViewAtlas) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiHostViewAtlas) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApiHostViewAtlas) HasHostname() bool`

HasHostname returns a boolean if a field has been set.
### GetId

`func (o *ApiHostViewAtlas) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiHostViewAtlas) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiHostViewAtlas) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiHostViewAtlas) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLastPing

`func (o *ApiHostViewAtlas) GetLastPing() time.Time`

GetLastPing returns the LastPing field if non-nil, zero value otherwise.

### GetLastPingOk

`func (o *ApiHostViewAtlas) GetLastPingOk() (*time.Time, bool)`

GetLastPingOk returns a tuple with the LastPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPing

`func (o *ApiHostViewAtlas) SetLastPing(v time.Time)`

SetLastPing sets LastPing field to given value.

### HasLastPing

`func (o *ApiHostViewAtlas) HasLastPing() bool`

HasLastPing returns a boolean if a field has been set.
### GetLinks

`func (o *ApiHostViewAtlas) GetLinks() []LinkAtlas`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiHostViewAtlas) GetLinksOk() (*[]LinkAtlas, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiHostViewAtlas) SetLinks(v []LinkAtlas)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiHostViewAtlas) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetPort

`func (o *ApiHostViewAtlas) GetPort() int`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiHostViewAtlas) GetPortOk() (*int, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiHostViewAtlas) SetPort(v int)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiHostViewAtlas) HasPort() bool`

HasPort returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *ApiHostViewAtlas) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ApiHostViewAtlas) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ApiHostViewAtlas) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *ApiHostViewAtlas) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.
### GetShardName

`func (o *ApiHostViewAtlas) GetShardName() string`

GetShardName returns the ShardName field if non-nil, zero value otherwise.

### GetShardNameOk

`func (o *ApiHostViewAtlas) GetShardNameOk() (*string, bool)`

GetShardNameOk returns a tuple with the ShardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardName

`func (o *ApiHostViewAtlas) SetShardName(v string)`

SetShardName sets ShardName field to given value.

### HasShardName

`func (o *ApiHostViewAtlas) HasShardName() bool`

HasShardName returns a boolean if a field has been set.
### GetTypeName

`func (o *ApiHostViewAtlas) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ApiHostViewAtlas) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ApiHostViewAtlas) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *ApiHostViewAtlas) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.
### GetUserAlias

`func (o *ApiHostViewAtlas) GetUserAlias() string`

GetUserAlias returns the UserAlias field if non-nil, zero value otherwise.

### GetUserAliasOk

`func (o *ApiHostViewAtlas) GetUserAliasOk() (*string, bool)`

GetUserAliasOk returns a tuple with the UserAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAlias

`func (o *ApiHostViewAtlas) SetUserAlias(v string)`

SetUserAlias sets UserAlias field to given value.

### HasUserAlias

`func (o *ApiHostViewAtlas) HasUserAlias() bool`

HasUserAlias returns a boolean if a field has been set.
### GetVersion

`func (o *ApiHostViewAtlas) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiHostViewAtlas) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiHostViewAtlas) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiHostViewAtlas) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


