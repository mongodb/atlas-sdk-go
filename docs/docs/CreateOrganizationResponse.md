# CreateOrganizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to [**ApiKeyUserDetails**](ApiKeyUserDetails.md) |  | [optional] 
**FederationSettingsId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the federation that you linked the newly created organization to. | [optional] [readonly] 
**OrgOwnerId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user that you assigned the Organization Owner role in the new organization. | [optional] [readonly] 
**Organization** | Pointer to [**AtlasOrganization**](AtlasOrganization.md) |  | [optional] 

## Methods

### NewCreateOrganizationResponse

`func NewCreateOrganizationResponse() *CreateOrganizationResponse`

NewCreateOrganizationResponse instantiates a new CreateOrganizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationResponseWithDefaults

`func NewCreateOrganizationResponseWithDefaults() *CreateOrganizationResponse`

NewCreateOrganizationResponseWithDefaults instantiates a new CreateOrganizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CreateOrganizationResponse) GetApiKey() ApiKeyUserDetails`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateOrganizationResponse) GetApiKeyOk() (*ApiKeyUserDetails, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateOrganizationResponse) SetApiKey(v ApiKeyUserDetails)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateOrganizationResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetFederationSettingsId

`func (o *CreateOrganizationResponse) GetFederationSettingsId() string`

GetFederationSettingsId returns the FederationSettingsId field if non-nil, zero value otherwise.

### GetFederationSettingsIdOk

`func (o *CreateOrganizationResponse) GetFederationSettingsIdOk() (*string, bool)`

GetFederationSettingsIdOk returns a tuple with the FederationSettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSettingsId

`func (o *CreateOrganizationResponse) SetFederationSettingsId(v string)`

SetFederationSettingsId sets FederationSettingsId field to given value.

### HasFederationSettingsId

`func (o *CreateOrganizationResponse) HasFederationSettingsId() bool`

HasFederationSettingsId returns a boolean if a field has been set.

### GetOrgOwnerId

`func (o *CreateOrganizationResponse) GetOrgOwnerId() string`

GetOrgOwnerId returns the OrgOwnerId field if non-nil, zero value otherwise.

### GetOrgOwnerIdOk

`func (o *CreateOrganizationResponse) GetOrgOwnerIdOk() (*string, bool)`

GetOrgOwnerIdOk returns a tuple with the OrgOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgOwnerId

`func (o *CreateOrganizationResponse) SetOrgOwnerId(v string)`

SetOrgOwnerId sets OrgOwnerId field to given value.

### HasOrgOwnerId

`func (o *CreateOrganizationResponse) HasOrgOwnerId() bool`

HasOrgOwnerId returns a boolean if a field has been set.

### GetOrganization

`func (o *CreateOrganizationResponse) GetOrganization() AtlasOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CreateOrganizationResponse) GetOrganizationOk() (*AtlasOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CreateOrganizationResponse) SetOrganization(v AtlasOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CreateOrganizationResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


