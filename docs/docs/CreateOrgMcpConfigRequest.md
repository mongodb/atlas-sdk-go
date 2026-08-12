# CreateOrgMcpConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAccessList** | Pointer to [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | List of IP access list entries that define allowed source addresses for this MCP configuration. | [optional] 
**McpConfigName** | **string** | Human-readable name that identifies this MCP configuration. | 
**Roles** | **[]string** | List of organization roles to assign to this MCP configuration. | 

## Methods

### NewCreateOrgMcpConfigRequest

`func NewCreateOrgMcpConfigRequest(mcpConfigName string, roles []string, ) *CreateOrgMcpConfigRequest`

NewCreateOrgMcpConfigRequest instantiates a new CreateOrgMcpConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgMcpConfigRequestWithDefaults

`func NewCreateOrgMcpConfigRequestWithDefaults() *CreateOrgMcpConfigRequest`

NewCreateOrgMcpConfigRequestWithDefaults instantiates a new CreateOrgMcpConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAccessList

`func (o *CreateOrgMcpConfigRequest) GetIpAccessList() []ServiceAccountIPAccessListEntry`

GetIpAccessList returns the IpAccessList field if non-nil, zero value otherwise.

### GetIpAccessListOk

`func (o *CreateOrgMcpConfigRequest) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool)`

GetIpAccessListOk returns a tuple with the IpAccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessList

`func (o *CreateOrgMcpConfigRequest) SetIpAccessList(v []ServiceAccountIPAccessListEntry)`

SetIpAccessList sets IpAccessList field to given value.

### HasIpAccessList

`func (o *CreateOrgMcpConfigRequest) HasIpAccessList() bool`

HasIpAccessList returns a boolean if a field has been set.

### SetIpAccessListNil

`func (o *CreateOrgMcpConfigRequest) SetIpAccessListNil()`

SetIpAccessListNil sets IpAccessList to an explicit JSON null when marshaled, overriding any value previously set with SetIpAccessList. Calling SetIpAccessList again clears the null override.

### GetMcpConfigName

`func (o *CreateOrgMcpConfigRequest) GetMcpConfigName() string`

GetMcpConfigName returns the McpConfigName field if non-nil, zero value otherwise.

### GetMcpConfigNameOk

`func (o *CreateOrgMcpConfigRequest) GetMcpConfigNameOk() (*string, bool)`

GetMcpConfigNameOk returns a tuple with the McpConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpConfigName

`func (o *CreateOrgMcpConfigRequest) SetMcpConfigName(v string)`

SetMcpConfigName sets McpConfigName field to given value.

### GetRoles

`func (o *CreateOrgMcpConfigRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateOrgMcpConfigRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateOrgMcpConfigRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


