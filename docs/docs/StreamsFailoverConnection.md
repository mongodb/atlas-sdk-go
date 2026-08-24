# StreamsFailoverConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the connection. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the stream connection. | [optional] 
**Region** | Pointer to **string** | Connection region. | [optional] 
**State** | Pointer to **string** | Connection state. | [optional] [readonly] 
**Type** | Pointer to **string** | Connection type. | [optional] 
**ClusterGroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that contains the configured cluster. Required if the ID does not match the project containing the streams workspace. You must first enable the organization setting. | [optional] 
**ClusterName** | Pointer to **string** | Name of the cluster configured for this connection. | [optional] 
**DbRoleToExecute** | Pointer to [**DBRoleToExecute**](DBRoleToExecute.md) |  | [optional] 
**Authentication** | Pointer to [**StreamsKafkaAuthentication**](StreamsKafkaAuthentication.md) |  | [optional] 
**BootstrapServers** | Pointer to **string** | Comma separated list of server addresses. | [optional] 
**Config** | Pointer to **map[string]string** | Map of Kafka key-value pairs for optional configuration. This object is flat, and keys can have &#39;.&#39; characters. | [optional] 
**Networking** | Pointer to [**StreamsKafkaNetworking**](StreamsKafkaNetworking.md) |  | [optional] 
**Security** | Pointer to [**StreamsKafkaSecurity**](StreamsKafkaSecurity.md) |  | [optional] 

## Methods

### NewStreamsFailoverConnection

`func NewStreamsFailoverConnection() *StreamsFailoverConnection`

NewStreamsFailoverConnection instantiates a new StreamsFailoverConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsFailoverConnectionWithDefaults

`func NewStreamsFailoverConnectionWithDefaults() *StreamsFailoverConnection`

NewStreamsFailoverConnectionWithDefaults instantiates a new StreamsFailoverConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamsFailoverConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsFailoverConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsFailoverConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamsFailoverConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StreamsFailoverConnection) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetLinks

`func (o *StreamsFailoverConnection) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsFailoverConnection) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsFailoverConnection) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsFailoverConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StreamsFailoverConnection) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetName

`func (o *StreamsFailoverConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsFailoverConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsFailoverConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamsFailoverConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StreamsFailoverConnection) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.

### GetRegion

`func (o *StreamsFailoverConnection) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsFailoverConnection) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsFailoverConnection) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamsFailoverConnection) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *StreamsFailoverConnection) SetRegionNil()`

SetRegionNil sets Region to an explicit JSON null when marshaled, overriding any value previously set with SetRegion. Calling SetRegion again clears the null override.

### GetState

`func (o *StreamsFailoverConnection) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamsFailoverConnection) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamsFailoverConnection) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StreamsFailoverConnection) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *StreamsFailoverConnection) SetStateNil()`

SetStateNil sets State to an explicit JSON null when marshaled, overriding any value previously set with SetState. Calling SetState again clears the null override.

### GetType

`func (o *StreamsFailoverConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamsFailoverConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamsFailoverConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamsFailoverConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *StreamsFailoverConnection) SetTypeNil()`

SetTypeNil sets Type to an explicit JSON null when marshaled, overriding any value previously set with SetType. Calling SetType again clears the null override.

### GetClusterGroupId

`func (o *StreamsFailoverConnection) GetClusterGroupId() string`

GetClusterGroupId returns the ClusterGroupId field if non-nil, zero value otherwise.

### GetClusterGroupIdOk

`func (o *StreamsFailoverConnection) GetClusterGroupIdOk() (*string, bool)`

GetClusterGroupIdOk returns a tuple with the ClusterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroupId

`func (o *StreamsFailoverConnection) SetClusterGroupId(v string)`

SetClusterGroupId sets ClusterGroupId field to given value.

### HasClusterGroupId

`func (o *StreamsFailoverConnection) HasClusterGroupId() bool`

HasClusterGroupId returns a boolean if a field has been set.

### SetClusterGroupIdNil

`func (o *StreamsFailoverConnection) SetClusterGroupIdNil()`

SetClusterGroupIdNil sets ClusterGroupId to an explicit JSON null when marshaled, overriding any value previously set with SetClusterGroupId. Calling SetClusterGroupId again clears the null override.

### GetClusterName

`func (o *StreamsFailoverConnection) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *StreamsFailoverConnection) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *StreamsFailoverConnection) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *StreamsFailoverConnection) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *StreamsFailoverConnection) SetClusterNameNil()`

SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled, overriding any value previously set with SetClusterName. Calling SetClusterName again clears the null override.

### GetDbRoleToExecute

`func (o *StreamsFailoverConnection) GetDbRoleToExecute() DBRoleToExecute`

GetDbRoleToExecute returns the DbRoleToExecute field if non-nil, zero value otherwise.

### GetDbRoleToExecuteOk

`func (o *StreamsFailoverConnection) GetDbRoleToExecuteOk() (*DBRoleToExecute, bool)`

GetDbRoleToExecuteOk returns a tuple with the DbRoleToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbRoleToExecute

`func (o *StreamsFailoverConnection) SetDbRoleToExecute(v DBRoleToExecute)`

SetDbRoleToExecute sets DbRoleToExecute field to given value.

### HasDbRoleToExecute

`func (o *StreamsFailoverConnection) HasDbRoleToExecute() bool`

HasDbRoleToExecute returns a boolean if a field has been set.

### SetDbRoleToExecuteNil

`func (o *StreamsFailoverConnection) SetDbRoleToExecuteNil()`

SetDbRoleToExecuteNil sets DbRoleToExecute to an explicit JSON null when marshaled, overriding any value previously set with SetDbRoleToExecute. Calling SetDbRoleToExecute again clears the null override.

### GetAuthentication

`func (o *StreamsFailoverConnection) GetAuthentication() StreamsKafkaAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *StreamsFailoverConnection) GetAuthenticationOk() (*StreamsKafkaAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *StreamsFailoverConnection) SetAuthentication(v StreamsKafkaAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *StreamsFailoverConnection) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *StreamsFailoverConnection) SetAuthenticationNil()`

SetAuthenticationNil sets Authentication to an explicit JSON null when marshaled, overriding any value previously set with SetAuthentication. Calling SetAuthentication again clears the null override.

### GetBootstrapServers

`func (o *StreamsFailoverConnection) GetBootstrapServers() string`

GetBootstrapServers returns the BootstrapServers field if non-nil, zero value otherwise.

### GetBootstrapServersOk

`func (o *StreamsFailoverConnection) GetBootstrapServersOk() (*string, bool)`

GetBootstrapServersOk returns a tuple with the BootstrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServers

`func (o *StreamsFailoverConnection) SetBootstrapServers(v string)`

SetBootstrapServers sets BootstrapServers field to given value.

### HasBootstrapServers

`func (o *StreamsFailoverConnection) HasBootstrapServers() bool`

HasBootstrapServers returns a boolean if a field has been set.

### SetBootstrapServersNil

`func (o *StreamsFailoverConnection) SetBootstrapServersNil()`

SetBootstrapServersNil sets BootstrapServers to an explicit JSON null when marshaled, overriding any value previously set with SetBootstrapServers. Calling SetBootstrapServers again clears the null override.

### GetConfig

`func (o *StreamsFailoverConnection) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StreamsFailoverConnection) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StreamsFailoverConnection) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *StreamsFailoverConnection) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *StreamsFailoverConnection) SetConfigNil()`

SetConfigNil sets Config to an explicit JSON null when marshaled, overriding any value previously set with SetConfig. Calling SetConfig again clears the null override.

### GetNetworking

`func (o *StreamsFailoverConnection) GetNetworking() StreamsKafkaNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *StreamsFailoverConnection) GetNetworkingOk() (*StreamsKafkaNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *StreamsFailoverConnection) SetNetworking(v StreamsKafkaNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *StreamsFailoverConnection) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### SetNetworkingNil

`func (o *StreamsFailoverConnection) SetNetworkingNil()`

SetNetworkingNil sets Networking to an explicit JSON null when marshaled, overriding any value previously set with SetNetworking. Calling SetNetworking again clears the null override.

### GetSecurity

`func (o *StreamsFailoverConnection) GetSecurity() StreamsKafkaSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *StreamsFailoverConnection) GetSecurityOk() (*StreamsKafkaSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *StreamsFailoverConnection) SetSecurity(v StreamsKafkaSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *StreamsFailoverConnection) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *StreamsFailoverConnection) SetSecurityNil()`

SetSecurityNil sets Security to an explicit JSON null when marshaled, overriding any value previously set with SetSecurity. Calling SetSecurity again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


