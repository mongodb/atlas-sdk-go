# GroupMcpConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Unique identifier for the Service Account client associated with this MCP configuration. Use this Service Account to connect to the Atlas Remote MCP. | [optional] [readonly] 
**EgressClientId** | Pointer to **string** | Unique identifier for the egress Service Account client associated with this MCP configuration. This Service Account is managed by MongoDB Atlas. | [optional] [readonly] 
**IpAccessList** | Pointer to [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | List of IP access list entries that define allowed source addresses for this MCP configuration. | [optional] 
**McpConfigId** | Pointer to **string** | Unique identifier that identifies this MCP configuration. | [optional] [readonly] 
**McpConfigName** | Pointer to **string** | Human-readable name that identifies this MCP configuration. | [optional] 
**Roles** | Pointer to **[]string** | List of project roles associated with this MCP configuration. | [optional] 

## Methods

### NewGroupMcpConfigResponse

`func NewGroupMcpConfigResponse() *GroupMcpConfigResponse`

NewGroupMcpConfigResponse instantiates a new GroupMcpConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMcpConfigResponseWithDefaults

`func NewGroupMcpConfigResponseWithDefaults() *GroupMcpConfigResponse`

NewGroupMcpConfigResponseWithDefaults instantiates a new GroupMcpConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *GroupMcpConfigResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GroupMcpConfigResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GroupMcpConfigResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GroupMcpConfigResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *GroupMcpConfigResponse) SetClientIdNil()`

SetClientIdNil sets ClientId to an explicit JSON null when marshaled, overriding any value previously set with SetClientId. Calling SetClientId again clears the null override.

### GetEgressClientId

`func (o *GroupMcpConfigResponse) GetEgressClientId() string`

GetEgressClientId returns the EgressClientId field if non-nil, zero value otherwise.

### GetEgressClientIdOk

`func (o *GroupMcpConfigResponse) GetEgressClientIdOk() (*string, bool)`

GetEgressClientIdOk returns a tuple with the EgressClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressClientId

`func (o *GroupMcpConfigResponse) SetEgressClientId(v string)`

SetEgressClientId sets EgressClientId field to given value.

### HasEgressClientId

`func (o *GroupMcpConfigResponse) HasEgressClientId() bool`

HasEgressClientId returns a boolean if a field has been set.

### SetEgressClientIdNil

`func (o *GroupMcpConfigResponse) SetEgressClientIdNil()`

SetEgressClientIdNil sets EgressClientId to an explicit JSON null when marshaled, overriding any value previously set with SetEgressClientId. Calling SetEgressClientId again clears the null override.

### GetIpAccessList

`func (o *GroupMcpConfigResponse) GetIpAccessList() []ServiceAccountIPAccessListEntry`

GetIpAccessList returns the IpAccessList field if non-nil, zero value otherwise.

### GetIpAccessListOk

`func (o *GroupMcpConfigResponse) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool)`

GetIpAccessListOk returns a tuple with the IpAccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessList

`func (o *GroupMcpConfigResponse) SetIpAccessList(v []ServiceAccountIPAccessListEntry)`

SetIpAccessList sets IpAccessList field to given value.

### HasIpAccessList

`func (o *GroupMcpConfigResponse) HasIpAccessList() bool`

HasIpAccessList returns a boolean if a field has been set.

### SetIpAccessListNil

`func (o *GroupMcpConfigResponse) SetIpAccessListNil()`

SetIpAccessListNil sets IpAccessList to an explicit JSON null when marshaled, overriding any value previously set with SetIpAccessList. Calling SetIpAccessList again clears the null override.

### GetMcpConfigId

`func (o *GroupMcpConfigResponse) GetMcpConfigId() string`

GetMcpConfigId returns the McpConfigId field if non-nil, zero value otherwise.

### GetMcpConfigIdOk

`func (o *GroupMcpConfigResponse) GetMcpConfigIdOk() (*string, bool)`

GetMcpConfigIdOk returns a tuple with the McpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpConfigId

`func (o *GroupMcpConfigResponse) SetMcpConfigId(v string)`

SetMcpConfigId sets McpConfigId field to given value.

### HasMcpConfigId

`func (o *GroupMcpConfigResponse) HasMcpConfigId() bool`

HasMcpConfigId returns a boolean if a field has been set.

### SetMcpConfigIdNil

`func (o *GroupMcpConfigResponse) SetMcpConfigIdNil()`

SetMcpConfigIdNil sets McpConfigId to an explicit JSON null when marshaled, overriding any value previously set with SetMcpConfigId. Calling SetMcpConfigId again clears the null override.

### GetMcpConfigName

`func (o *GroupMcpConfigResponse) GetMcpConfigName() string`

GetMcpConfigName returns the McpConfigName field if non-nil, zero value otherwise.

### GetMcpConfigNameOk

`func (o *GroupMcpConfigResponse) GetMcpConfigNameOk() (*string, bool)`

GetMcpConfigNameOk returns a tuple with the McpConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpConfigName

`func (o *GroupMcpConfigResponse) SetMcpConfigName(v string)`

SetMcpConfigName sets McpConfigName field to given value.

### HasMcpConfigName

`func (o *GroupMcpConfigResponse) HasMcpConfigName() bool`

HasMcpConfigName returns a boolean if a field has been set.

### SetMcpConfigNameNil

`func (o *GroupMcpConfigResponse) SetMcpConfigNameNil()`

SetMcpConfigNameNil sets McpConfigName to an explicit JSON null when marshaled, overriding any value previously set with SetMcpConfigName. Calling SetMcpConfigName again clears the null override.

### GetRoles

`func (o *GroupMcpConfigResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupMcpConfigResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupMcpConfigResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupMcpConfigResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *GroupMcpConfigResponse) SetRolesNil()`

SetRolesNil sets Roles to an explicit JSON null when marshaled, overriding any value previously set with SetRoles. Calling SetRoles again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


