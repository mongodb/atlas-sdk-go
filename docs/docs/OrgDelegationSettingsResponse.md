# OrgDelegationSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedMcpAccess** | Pointer to **string** | Policy that controls how MCP (Model Context Protocol) delegated access is permitted within this organization. Possible values are &#x60;DISALLOWED&#x60;, &#x60;READ_ONLY&#x60;, and &#x60;READ_WRITE&#x60;. Defaults to &#x60;DISALLOWED&#x60;. | [optional] 
**DelegatedPartnerAccess** | Pointer to **string** | Policy that controls whether partner delegated access is permitted within this organization. Possible values are &#x60;DISALLOWED&#x60; and &#x60;READ_WRITE&#x60;. Defaults to &#x60;DISALLOWED&#x60;. | [optional] 
**IdleRefreshTokenLifetime** | Pointer to **int64** | Maximum number of seconds a refresh token may be idle before it expires. When not set, the system default applies. | [optional] 
**MaximumRefreshTokenLifetime** | Pointer to **int64** | Maximum lifetime of a refresh token in seconds, regardless of activity. When not set, the system default applies. | [optional] 

## Methods

### NewOrgDelegationSettingsResponse

`func NewOrgDelegationSettingsResponse() *OrgDelegationSettingsResponse`

NewOrgDelegationSettingsResponse instantiates a new OrgDelegationSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgDelegationSettingsResponseWithDefaults

`func NewOrgDelegationSettingsResponseWithDefaults() *OrgDelegationSettingsResponse`

NewOrgDelegationSettingsResponseWithDefaults instantiates a new OrgDelegationSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedMcpAccess

`func (o *OrgDelegationSettingsResponse) GetDelegatedMcpAccess() string`

GetDelegatedMcpAccess returns the DelegatedMcpAccess field if non-nil, zero value otherwise.

### GetDelegatedMcpAccessOk

`func (o *OrgDelegationSettingsResponse) GetDelegatedMcpAccessOk() (*string, bool)`

GetDelegatedMcpAccessOk returns a tuple with the DelegatedMcpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMcpAccess

`func (o *OrgDelegationSettingsResponse) SetDelegatedMcpAccess(v string)`

SetDelegatedMcpAccess sets DelegatedMcpAccess field to given value.

### HasDelegatedMcpAccess

`func (o *OrgDelegationSettingsResponse) HasDelegatedMcpAccess() bool`

HasDelegatedMcpAccess returns a boolean if a field has been set.
### GetDelegatedPartnerAccess

`func (o *OrgDelegationSettingsResponse) GetDelegatedPartnerAccess() string`

GetDelegatedPartnerAccess returns the DelegatedPartnerAccess field if non-nil, zero value otherwise.

### GetDelegatedPartnerAccessOk

`func (o *OrgDelegationSettingsResponse) GetDelegatedPartnerAccessOk() (*string, bool)`

GetDelegatedPartnerAccessOk returns a tuple with the DelegatedPartnerAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedPartnerAccess

`func (o *OrgDelegationSettingsResponse) SetDelegatedPartnerAccess(v string)`

SetDelegatedPartnerAccess sets DelegatedPartnerAccess field to given value.

### HasDelegatedPartnerAccess

`func (o *OrgDelegationSettingsResponse) HasDelegatedPartnerAccess() bool`

HasDelegatedPartnerAccess returns a boolean if a field has been set.
### GetIdleRefreshTokenLifetime

`func (o *OrgDelegationSettingsResponse) GetIdleRefreshTokenLifetime() int64`

GetIdleRefreshTokenLifetime returns the IdleRefreshTokenLifetime field if non-nil, zero value otherwise.

### GetIdleRefreshTokenLifetimeOk

`func (o *OrgDelegationSettingsResponse) GetIdleRefreshTokenLifetimeOk() (*int64, bool)`

GetIdleRefreshTokenLifetimeOk returns a tuple with the IdleRefreshTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleRefreshTokenLifetime

`func (o *OrgDelegationSettingsResponse) SetIdleRefreshTokenLifetime(v int64)`

SetIdleRefreshTokenLifetime sets IdleRefreshTokenLifetime field to given value.

### HasIdleRefreshTokenLifetime

`func (o *OrgDelegationSettingsResponse) HasIdleRefreshTokenLifetime() bool`

HasIdleRefreshTokenLifetime returns a boolean if a field has been set.
### GetMaximumRefreshTokenLifetime

`func (o *OrgDelegationSettingsResponse) GetMaximumRefreshTokenLifetime() int64`

GetMaximumRefreshTokenLifetime returns the MaximumRefreshTokenLifetime field if non-nil, zero value otherwise.

### GetMaximumRefreshTokenLifetimeOk

`func (o *OrgDelegationSettingsResponse) GetMaximumRefreshTokenLifetimeOk() (*int64, bool)`

GetMaximumRefreshTokenLifetimeOk returns a tuple with the MaximumRefreshTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRefreshTokenLifetime

`func (o *OrgDelegationSettingsResponse) SetMaximumRefreshTokenLifetime(v int64)`

SetMaximumRefreshTokenLifetime sets MaximumRefreshTokenLifetime field to given value.

### HasMaximumRefreshTokenLifetime

`func (o *OrgDelegationSettingsResponse) HasMaximumRefreshTokenLifetime() bool`

HasMaximumRefreshTokenLifetime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


