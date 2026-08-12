# CreateGroupMcpConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAccessList** | Pointer to [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | List of IP access list entries that define allowed source addresses for this MCP configuration. | [optional] 
**McpConfigName** | **string** | Human-readable name that identifies this MCP configuration. | 
**Roles** | **[]string** | List of project roles to assign to this MCP configuration. | 

## Methods

### NewCreateGroupMcpConfigRequest

`func NewCreateGroupMcpConfigRequest(mcpConfigName string, roles []string, ) *CreateGroupMcpConfigRequest`

NewCreateGroupMcpConfigRequest instantiates a new CreateGroupMcpConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupMcpConfigRequestWithDefaults

`func NewCreateGroupMcpConfigRequestWithDefaults() *CreateGroupMcpConfigRequest`

NewCreateGroupMcpConfigRequestWithDefaults instantiates a new CreateGroupMcpConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAccessList

`func (o *CreateGroupMcpConfigRequest) GetIpAccessList() []ServiceAccountIPAccessListEntry`

GetIpAccessList returns the IpAccessList field if non-nil, zero value otherwise.

### GetIpAccessListOk

`func (o *CreateGroupMcpConfigRequest) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool)`

GetIpAccessListOk returns a tuple with the IpAccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessList

`func (o *CreateGroupMcpConfigRequest) SetIpAccessList(v []ServiceAccountIPAccessListEntry)`

SetIpAccessList sets IpAccessList field to given value.

### HasIpAccessList

`func (o *CreateGroupMcpConfigRequest) HasIpAccessList() bool`

HasIpAccessList returns a boolean if a field has been set.

### SetIpAccessListNil

`func (o *CreateGroupMcpConfigRequest) SetIpAccessListNil()`

SetIpAccessListNil sets IpAccessList to an explicit JSON null when marshaled, overriding any value previously set with SetIpAccessList. Calling SetIpAccessList again clears the null override.

### GetMcpConfigName

`func (o *CreateGroupMcpConfigRequest) GetMcpConfigName() string`

GetMcpConfigName returns the McpConfigName field if non-nil, zero value otherwise.

### GetMcpConfigNameOk

`func (o *CreateGroupMcpConfigRequest) GetMcpConfigNameOk() (*string, bool)`

GetMcpConfigNameOk returns a tuple with the McpConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpConfigName

`func (o *CreateGroupMcpConfigRequest) SetMcpConfigName(v string)`

SetMcpConfigName sets McpConfigName field to given value.

### GetRoles

`func (o *CreateGroupMcpConfigRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateGroupMcpConfigRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateGroupMcpConfigRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


