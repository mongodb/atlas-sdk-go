# StreamsConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the stream connection. In the case of the Sample type, this is the name of the sample source. | [optional] 
**Type** | Pointer to **string** | Type of the connection. Can be either Cluster or Kafka. | [optional] 
**ClusterName** | Pointer to **string** | Name of the cluster configured for this connection. | [optional] 
**DbRoleToExecute** | Pointer to [**DBRoleToExecute**](DBRoleToExecute.md) |  | [optional] 
**Authentication** | Pointer to [**StreamsKafkaAuthentication**](StreamsKafkaAuthentication.md) |  | [optional] 
**BootstrapServers** | Pointer to **string** | Comma separated list of server addresses. | [optional] 
**Config** | Pointer to **map[string]string** | A map of Kafka key-value pairs for optional configuration. This is a flat object, and keys can have &#39;.&#39; characters. | [optional] 
**Networking** | Pointer to [**StreamsKafkaNetworking**](StreamsKafkaNetworking.md) |  | [optional] 
**Security** | Pointer to [**StreamsKafkaSecurity**](StreamsKafkaSecurity.md) |  | [optional] 
**Headers** | Pointer to **map[string]string** | A map of key-value pairs that will be passed as headers for the request. | [optional] 
**Url** | Pointer to **string** | The url to be used for the request. | [optional] 
**Aws** | Pointer to [**StreamsAWSConnectionBaseConfig**](StreamsAWSConnectionBaseConfig.md) |  | [optional] 

## Methods

### NewStreamsConnection

`func NewStreamsConnection() *StreamsConnection`

NewStreamsConnection instantiates a new StreamsConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsConnectionWithDefaults

`func NewStreamsConnectionWithDefaults() *StreamsConnection`

NewStreamsConnectionWithDefaults instantiates a new StreamsConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsConnection) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsConnection) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsConnection) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetName

`func (o *StreamsConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamsConnection) HasName() bool`

HasName returns a boolean if a field has been set.
### GetType

`func (o *StreamsConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamsConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamsConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamsConnection) HasType() bool`

HasType returns a boolean if a field has been set.
### GetClusterName

`func (o *StreamsConnection) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *StreamsConnection) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *StreamsConnection) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *StreamsConnection) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetDbRoleToExecute

`func (o *StreamsConnection) GetDbRoleToExecute() DBRoleToExecute`

GetDbRoleToExecute returns the DbRoleToExecute field if non-nil, zero value otherwise.

### GetDbRoleToExecuteOk

`func (o *StreamsConnection) GetDbRoleToExecuteOk() (*DBRoleToExecute, bool)`

GetDbRoleToExecuteOk returns a tuple with the DbRoleToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbRoleToExecute

`func (o *StreamsConnection) SetDbRoleToExecute(v DBRoleToExecute)`

SetDbRoleToExecute sets DbRoleToExecute field to given value.

### HasDbRoleToExecute

`func (o *StreamsConnection) HasDbRoleToExecute() bool`

HasDbRoleToExecute returns a boolean if a field has been set.
### GetAuthentication

`func (o *StreamsConnection) GetAuthentication() StreamsKafkaAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *StreamsConnection) GetAuthenticationOk() (*StreamsKafkaAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *StreamsConnection) SetAuthentication(v StreamsKafkaAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *StreamsConnection) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.
### GetBootstrapServers

`func (o *StreamsConnection) GetBootstrapServers() string`

GetBootstrapServers returns the BootstrapServers field if non-nil, zero value otherwise.

### GetBootstrapServersOk

`func (o *StreamsConnection) GetBootstrapServersOk() (*string, bool)`

GetBootstrapServersOk returns a tuple with the BootstrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServers

`func (o *StreamsConnection) SetBootstrapServers(v string)`

SetBootstrapServers sets BootstrapServers field to given value.

### HasBootstrapServers

`func (o *StreamsConnection) HasBootstrapServers() bool`

HasBootstrapServers returns a boolean if a field has been set.
### GetConfig

`func (o *StreamsConnection) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StreamsConnection) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StreamsConnection) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *StreamsConnection) HasConfig() bool`

HasConfig returns a boolean if a field has been set.
### GetNetworking

`func (o *StreamsConnection) GetNetworking() StreamsKafkaNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *StreamsConnection) GetNetworkingOk() (*StreamsKafkaNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *StreamsConnection) SetNetworking(v StreamsKafkaNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *StreamsConnection) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.
### GetSecurity

`func (o *StreamsConnection) GetSecurity() StreamsKafkaSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *StreamsConnection) GetSecurityOk() (*StreamsKafkaSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *StreamsConnection) SetSecurity(v StreamsKafkaSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *StreamsConnection) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.
### GetHeaders

`func (o *StreamsConnection) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *StreamsConnection) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *StreamsConnection) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *StreamsConnection) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.
### GetUrl

`func (o *StreamsConnection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StreamsConnection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StreamsConnection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StreamsConnection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.
### GetAws

`func (o *StreamsConnection) GetAws() StreamsAWSConnectionBaseConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *StreamsConnection) GetAwsOk() (*StreamsAWSConnectionBaseConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *StreamsConnection) SetAws(v StreamsAWSConnectionBaseConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *StreamsConnection) HasAws() bool`

HasAws returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


