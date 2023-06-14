# CreateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to [**CreateOrganizationApiKey**](CreateOrganizationApiKey.md) |  | [optional] 
**FederationSettingsId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the federation to link the newly created organization to. If specified, the proposed Organization Owner of the new organization must have the Organization Owner role in an organization associated with the federation. | [optional] 
**Name** | **string** | Human-readable label that identifies the organization. | 
**OrgOwnerId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user that you want to assign the Organization Owner role. This user must be a member of the same organization as the calling API key. If you provide &#x60;federationSettingsId&#x60;,  this user must instead have the Organization Owner role on an organization in the specified federation. This parameter is required only when you authenticate with Programmatic API Keys. | [optional] 

## Methods

### NewCreateOrganizationRequest

`func NewCreateOrganizationRequest(name string, ) *CreateOrganizationRequest`

NewCreateOrganizationRequest instantiates a new CreateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationRequestWithDefaults

`func NewCreateOrganizationRequestWithDefaults() *CreateOrganizationRequest`

NewCreateOrganizationRequestWithDefaults instantiates a new CreateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CreateOrganizationRequest) GetApiKey() CreateOrganizationApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateOrganizationRequest) GetApiKeyOk() (*CreateOrganizationApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateOrganizationRequest) SetApiKey(v CreateOrganizationApiKey)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateOrganizationRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetFederationSettingsId

`func (o *CreateOrganizationRequest) GetFederationSettingsId() string`

GetFederationSettingsId returns the FederationSettingsId field if non-nil, zero value otherwise.

### GetFederationSettingsIdOk

`func (o *CreateOrganizationRequest) GetFederationSettingsIdOk() (*string, bool)`

GetFederationSettingsIdOk returns a tuple with the FederationSettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSettingsId

`func (o *CreateOrganizationRequest) SetFederationSettingsId(v string)`

SetFederationSettingsId sets FederationSettingsId field to given value.

### HasFederationSettingsId

`func (o *CreateOrganizationRequest) HasFederationSettingsId() bool`

HasFederationSettingsId returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrgOwnerId

`func (o *CreateOrganizationRequest) GetOrgOwnerId() string`

GetOrgOwnerId returns the OrgOwnerId field if non-nil, zero value otherwise.

### GetOrgOwnerIdOk

`func (o *CreateOrganizationRequest) GetOrgOwnerIdOk() (*string, bool)`

GetOrgOwnerIdOk returns a tuple with the OrgOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgOwnerId

`func (o *CreateOrganizationRequest) SetOrgOwnerId(v string)`

SetOrgOwnerId sets OrgOwnerId field to given value.

### HasOrgOwnerId

`func (o *CreateOrganizationRequest) HasOrgOwnerId() bool`

HasOrgOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


