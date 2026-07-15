# OrgDelegationSettingsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedMcpAccess** | Pointer to **string** | Policy that controls how MCP (Model Context Protocol) delegated access is permitted within this organization. Possible values are &#x60;DISALLOWED&#x60;, &#x60;READ_ONLY&#x60;, and &#x60;READ_WRITE&#x60;. Defaults to &#x60;DISALLOWED&#x60;. | [optional] 
**DelegatedPartnerAccess** | Pointer to **string** | Policy that controls whether partner delegated access is permitted within this organization. Possible values are &#x60;DISALLOWED&#x60; and &#x60;READ_WRITE&#x60;. Defaults to &#x60;DISALLOWED&#x60;. | [optional] 
**IdleRefreshTokenLifetime** | Pointer to **int64** | Maximum number of seconds a refresh token may be idle before it expires. Omit to leave unchanged; set to null to reset to the system default. Must be between 1 and 31536000 (1 year) when provided. | [optional] 
**MaximumRefreshTokenLifetime** | Pointer to **int64** | Maximum lifetime of a refresh token in seconds, regardless of activity. Omit to leave unchanged; set to null to reset to the system default. Must be between 1 and 31536000 (1 year) when provided. | [optional] 

## Methods

### NewOrgDelegationSettingsUpdateRequest

`func NewOrgDelegationSettingsUpdateRequest() *OrgDelegationSettingsUpdateRequest`

NewOrgDelegationSettingsUpdateRequest instantiates a new OrgDelegationSettingsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgDelegationSettingsUpdateRequestWithDefaults

`func NewOrgDelegationSettingsUpdateRequestWithDefaults() *OrgDelegationSettingsUpdateRequest`

NewOrgDelegationSettingsUpdateRequestWithDefaults instantiates a new OrgDelegationSettingsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedMcpAccess

`func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedMcpAccess() string`

GetDelegatedMcpAccess returns the DelegatedMcpAccess field if non-nil, zero value otherwise.

### GetDelegatedMcpAccessOk

`func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedMcpAccessOk() (*string, bool)`

GetDelegatedMcpAccessOk returns a tuple with the DelegatedMcpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMcpAccess

`func (o *OrgDelegationSettingsUpdateRequest) SetDelegatedMcpAccess(v string)`

SetDelegatedMcpAccess sets DelegatedMcpAccess field to given value.

### HasDelegatedMcpAccess

`func (o *OrgDelegationSettingsUpdateRequest) HasDelegatedMcpAccess() bool`

HasDelegatedMcpAccess returns a boolean if a field has been set.
### GetDelegatedPartnerAccess

`func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedPartnerAccess() string`

GetDelegatedPartnerAccess returns the DelegatedPartnerAccess field if non-nil, zero value otherwise.

### GetDelegatedPartnerAccessOk

`func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedPartnerAccessOk() (*string, bool)`

GetDelegatedPartnerAccessOk returns a tuple with the DelegatedPartnerAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedPartnerAccess

`func (o *OrgDelegationSettingsUpdateRequest) SetDelegatedPartnerAccess(v string)`

SetDelegatedPartnerAccess sets DelegatedPartnerAccess field to given value.

### HasDelegatedPartnerAccess

`func (o *OrgDelegationSettingsUpdateRequest) HasDelegatedPartnerAccess() bool`

HasDelegatedPartnerAccess returns a boolean if a field has been set.
### GetIdleRefreshTokenLifetime

`func (o *OrgDelegationSettingsUpdateRequest) GetIdleRefreshTokenLifetime() int64`

GetIdleRefreshTokenLifetime returns the IdleRefreshTokenLifetime field if non-nil, zero value otherwise.

### GetIdleRefreshTokenLifetimeOk

`func (o *OrgDelegationSettingsUpdateRequest) GetIdleRefreshTokenLifetimeOk() (*int64, bool)`

GetIdleRefreshTokenLifetimeOk returns a tuple with the IdleRefreshTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleRefreshTokenLifetime

`func (o *OrgDelegationSettingsUpdateRequest) SetIdleRefreshTokenLifetime(v int64)`

SetIdleRefreshTokenLifetime sets IdleRefreshTokenLifetime field to given value.

### HasIdleRefreshTokenLifetime

`func (o *OrgDelegationSettingsUpdateRequest) HasIdleRefreshTokenLifetime() bool`

HasIdleRefreshTokenLifetime returns a boolean if a field has been set.
### GetMaximumRefreshTokenLifetime

`func (o *OrgDelegationSettingsUpdateRequest) GetMaximumRefreshTokenLifetime() int64`

GetMaximumRefreshTokenLifetime returns the MaximumRefreshTokenLifetime field if non-nil, zero value otherwise.

### GetMaximumRefreshTokenLifetimeOk

`func (o *OrgDelegationSettingsUpdateRequest) GetMaximumRefreshTokenLifetimeOk() (*int64, bool)`

GetMaximumRefreshTokenLifetimeOk returns a tuple with the MaximumRefreshTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRefreshTokenLifetime

`func (o *OrgDelegationSettingsUpdateRequest) SetMaximumRefreshTokenLifetime(v int64)`

SetMaximumRefreshTokenLifetime sets MaximumRefreshTokenLifetime field to given value.

### HasMaximumRefreshTokenLifetime

`func (o *OrgDelegationSettingsUpdateRequest) HasMaximumRefreshTokenLifetime() bool`

HasMaximumRefreshTokenLifetime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


