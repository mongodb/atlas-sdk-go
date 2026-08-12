# OrgMcpConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAccessList** | Pointer to [**[]ServiceAccountIPAccessListEntry**](ServiceAccountIPAccessListEntry.md) | List of IP access list entries that define allowed source addresses for this MCP configuration. If provided, replaces the existing IP access list. | [optional] 
**McpConfigName** | Pointer to **string** | Updated human-readable name for this MCP configuration. | [optional] 
**Roles** | Pointer to **[]string** | List of organization roles associated with this MCP configuration. If provided, replaces the existing list of roles. | [optional] 

## Methods

### NewOrgMcpConfigUpdateRequest

`func NewOrgMcpConfigUpdateRequest() *OrgMcpConfigUpdateRequest`

NewOrgMcpConfigUpdateRequest instantiates a new OrgMcpConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMcpConfigUpdateRequestWithDefaults

`func NewOrgMcpConfigUpdateRequestWithDefaults() *OrgMcpConfigUpdateRequest`

NewOrgMcpConfigUpdateRequestWithDefaults instantiates a new OrgMcpConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAccessList

`func (o *OrgMcpConfigUpdateRequest) GetIpAccessList() []ServiceAccountIPAccessListEntry`

GetIpAccessList returns the IpAccessList field if non-nil, zero value otherwise.

### GetIpAccessListOk

`func (o *OrgMcpConfigUpdateRequest) GetIpAccessListOk() (*[]ServiceAccountIPAccessListEntry, bool)`

GetIpAccessListOk returns a tuple with the IpAccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessList

`func (o *OrgMcpConfigUpdateRequest) SetIpAccessList(v []ServiceAccountIPAccessListEntry)`

SetIpAccessList sets IpAccessList field to given value.

### HasIpAccessList

`func (o *OrgMcpConfigUpdateRequest) HasIpAccessList() bool`

HasIpAccessList returns a boolean if a field has been set.

### SetIpAccessListNil

`func (o *OrgMcpConfigUpdateRequest) SetIpAccessListNil()`

SetIpAccessListNil sets IpAccessList to an explicit JSON null when marshaled, overriding any value previously set with SetIpAccessList. Calling SetIpAccessList again clears the null override.

### GetMcpConfigName

`func (o *OrgMcpConfigUpdateRequest) GetMcpConfigName() string`

GetMcpConfigName returns the McpConfigName field if non-nil, zero value otherwise.

### GetMcpConfigNameOk

`func (o *OrgMcpConfigUpdateRequest) GetMcpConfigNameOk() (*string, bool)`

GetMcpConfigNameOk returns a tuple with the McpConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpConfigName

`func (o *OrgMcpConfigUpdateRequest) SetMcpConfigName(v string)`

SetMcpConfigName sets McpConfigName field to given value.

### HasMcpConfigName

`func (o *OrgMcpConfigUpdateRequest) HasMcpConfigName() bool`

HasMcpConfigName returns a boolean if a field has been set.

### SetMcpConfigNameNil

`func (o *OrgMcpConfigUpdateRequest) SetMcpConfigNameNil()`

SetMcpConfigNameNil sets McpConfigName to an explicit JSON null when marshaled, overriding any value previously set with SetMcpConfigName. Calling SetMcpConfigName again clears the null override.

### GetRoles

`func (o *OrgMcpConfigUpdateRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgMcpConfigUpdateRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgMcpConfigUpdateRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrgMcpConfigUpdateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *OrgMcpConfigUpdateRequest) SetRolesNil()`

SetRolesNil sets Roles to an explicit JSON null when marshaled, overriding any value previously set with SetRoles. Calling SetRoles again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


