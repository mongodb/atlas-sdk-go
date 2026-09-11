# OrganizationMaintenanceSettingsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaveAssignmentMode** | Pointer to **string** | Mode configured for this organization that determines how maintenance waves are assigned to projects. Possible values are &#x60;MANUAL&#x60; and &#x60;ENV_TAG_MAPPING&#x60;. Defaults to &#x60;MANUAL&#x60; when unset. Set the organization to &#x60;ENV_TAG_MAPPING&#x60; mode to have Atlas derive the maintenance wave from each project&#39;s environment tag. The tag key must be &#x60;environment&#x60;. The tag value determines the wave: &#x60;development&#x60; or &#x60;test&#x60; maps to Wave 1, &#x60;staging&#x60; maps to Wave 2, and &#x60;production&#x60; maps to Wave 3. Only this field can be updated; Atlas derives read-only &#x60;effectiveWaveAssignmentMode&#x60; on GET responses and uses that value for scheduling when it differs from &#x60;waveAssignmentMode&#x60;. Omit this field to leave the current value unchanged. Specify null to reset to the default value (&#x60;MANUAL&#x60;). | [optional] 

## Methods

### NewOrganizationMaintenanceSettingsUpdateRequest

`func NewOrganizationMaintenanceSettingsUpdateRequest() *OrganizationMaintenanceSettingsUpdateRequest`

NewOrganizationMaintenanceSettingsUpdateRequest instantiates a new OrganizationMaintenanceSettingsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMaintenanceSettingsUpdateRequestWithDefaults

`func NewOrganizationMaintenanceSettingsUpdateRequestWithDefaults() *OrganizationMaintenanceSettingsUpdateRequest`

NewOrganizationMaintenanceSettingsUpdateRequestWithDefaults instantiates a new OrganizationMaintenanceSettingsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsUpdateRequest) GetWaveAssignmentMode() string`

GetWaveAssignmentMode returns the WaveAssignmentMode field if non-nil, zero value otherwise.

### GetWaveAssignmentModeOk

`func (o *OrganizationMaintenanceSettingsUpdateRequest) GetWaveAssignmentModeOk() (*string, bool)`

GetWaveAssignmentModeOk returns a tuple with the WaveAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsUpdateRequest) SetWaveAssignmentMode(v string)`

SetWaveAssignmentMode sets WaveAssignmentMode field to given value.

### HasWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsUpdateRequest) HasWaveAssignmentMode() bool`

HasWaveAssignmentMode returns a boolean if a field has been set.

### SetWaveAssignmentModeNil

`func (o *OrganizationMaintenanceSettingsUpdateRequest) SetWaveAssignmentModeNil()`

SetWaveAssignmentModeNil sets WaveAssignmentMode to an explicit JSON null when marshaled, overriding any value previously set with SetWaveAssignmentMode. Calling SetWaveAssignmentMode again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


