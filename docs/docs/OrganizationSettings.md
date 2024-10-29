# OrganizationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccessListRequired** | Pointer to **bool** | Flag that indicates whether to require API operations to originate from an IP Address added to the API access list for the specified organization. | [optional] 
**MultiFactorAuthRequired** | Pointer to **bool** | Flag that indicates whether to require users to set up Multi-Factor Authentication (MFA) before accessing the specified organization. To learn more, see: https://www.mongodb.com/docs/atlas/security-multi-factor-authentication/. | [optional] 
**RestrictEmployeeAccess** | Pointer to **bool** | Flag that indicates whether to block MongoDB Support from accessing Atlas infrastructure and cluster logs for any deployment in the specified organization without explicit permission. Once this setting is turned on, you can grant MongoDB Support a 24-hour bypass access to the Atlas deployment to resolve support issues. To learn more, see: https://www.mongodb.com/docs/atlas/security-restrict-support-access/. | [optional] 
**SecurityContact** | Pointer to **string** | String that specifies a single email address for the specified organization to receive security-related notifications. Specifying a security contact does not grant them authorization or access to Atlas for security decisions or approvals. An empty string is valid and clears the existing security contact (if any). | [optional] 

## Methods

### NewOrganizationSettings

`func NewOrganizationSettings() *OrganizationSettings`

NewOrganizationSettings instantiates a new OrganizationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSettingsWithDefaults

`func NewOrganizationSettingsWithDefaults() *OrganizationSettings`

NewOrganizationSettingsWithDefaults instantiates a new OrganizationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccessListRequired

`func (o *OrganizationSettings) GetApiAccessListRequired() bool`

GetApiAccessListRequired returns the ApiAccessListRequired field if non-nil, zero value otherwise.

### GetApiAccessListRequiredOk

`func (o *OrganizationSettings) GetApiAccessListRequiredOk() (*bool, bool)`

GetApiAccessListRequiredOk returns a tuple with the ApiAccessListRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessListRequired

`func (o *OrganizationSettings) SetApiAccessListRequired(v bool)`

SetApiAccessListRequired sets ApiAccessListRequired field to given value.

### HasApiAccessListRequired

`func (o *OrganizationSettings) HasApiAccessListRequired() bool`

HasApiAccessListRequired returns a boolean if a field has been set.
### GetMultiFactorAuthRequired

`func (o *OrganizationSettings) GetMultiFactorAuthRequired() bool`

GetMultiFactorAuthRequired returns the MultiFactorAuthRequired field if non-nil, zero value otherwise.

### GetMultiFactorAuthRequiredOk

`func (o *OrganizationSettings) GetMultiFactorAuthRequiredOk() (*bool, bool)`

GetMultiFactorAuthRequiredOk returns a tuple with the MultiFactorAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiFactorAuthRequired

`func (o *OrganizationSettings) SetMultiFactorAuthRequired(v bool)`

SetMultiFactorAuthRequired sets MultiFactorAuthRequired field to given value.

### HasMultiFactorAuthRequired

`func (o *OrganizationSettings) HasMultiFactorAuthRequired() bool`

HasMultiFactorAuthRequired returns a boolean if a field has been set.
### GetRestrictEmployeeAccess

`func (o *OrganizationSettings) GetRestrictEmployeeAccess() bool`

GetRestrictEmployeeAccess returns the RestrictEmployeeAccess field if non-nil, zero value otherwise.

### GetRestrictEmployeeAccessOk

`func (o *OrganizationSettings) GetRestrictEmployeeAccessOk() (*bool, bool)`

GetRestrictEmployeeAccessOk returns a tuple with the RestrictEmployeeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictEmployeeAccess

`func (o *OrganizationSettings) SetRestrictEmployeeAccess(v bool)`

SetRestrictEmployeeAccess sets RestrictEmployeeAccess field to given value.

### HasRestrictEmployeeAccess

`func (o *OrganizationSettings) HasRestrictEmployeeAccess() bool`

HasRestrictEmployeeAccess returns a boolean if a field has been set.
### GetSecurityContact

`func (o *OrganizationSettings) GetSecurityContact() string`

GetSecurityContact returns the SecurityContact field if non-nil, zero value otherwise.

### GetSecurityContactOk

`func (o *OrganizationSettings) GetSecurityContactOk() (*string, bool)`

GetSecurityContactOk returns a tuple with the SecurityContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContact

`func (o *OrganizationSettings) SetSecurityContact(v string)`

SetSecurityContact sets SecurityContact field to given value.

### HasSecurityContact

`func (o *OrganizationSettings) HasSecurityContact() bool`

HasSecurityContact returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


