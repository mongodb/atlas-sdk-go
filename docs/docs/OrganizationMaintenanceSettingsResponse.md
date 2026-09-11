# OrganizationMaintenanceSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveWaveAssignmentMode** | Pointer to **string** | Wave assignment mode that Atlas uses when scheduling maintenance for projects in this organization. This read-only field takes precedence over &#x60;waveAssignmentMode&#x60; for scheduling. It matches &#x60;waveAssignmentMode&#x60; except when cross-organization maintenance sequencing is enabled and this organization is a linked non-paying organization, in which case it reflects the paying organization&#39;s &#x60;waveAssignmentMode&#x60;. Possible values are &#x60;MANUAL&#x60; and &#x60;ENV_TAG_MAPPING&#x60;. Defaults to &#x60;MANUAL&#x60; when no mode is configured. Omitted from GET responses when maintenance sequencing is disabled for this organization. | [optional] [readonly] 
**WaveAssignmentMode** | Pointer to **string** | Mode explicitly configured for this organization that determines how maintenance waves are assigned to projects. Possible values are &#x60;MANUAL&#x60; and &#x60;ENV_TAG_MAPPING&#x60;. Omitted from the response when no explicit preference has been set; in that case &#x60;effectiveWaveAssignmentMode&#x60; reflects the value the system uses (&#x60;MANUAL&#x60; by default). Atlas uses read-only &#x60;effectiveWaveAssignmentMode&#x60; (not this field) for scheduling. In a cross-organization billing hierarchy, a linked non-paying organization cannot update its &#x60;effectiveWaveAssignmentMode&#x60; field, which inherits from the paying organization&#39;s &#x60;waveAssignmentMode&#x60;. In this case, a linked non-paying organization&#39;s &#x60;effectiveWaveAssignmentMode&#x60; and &#x60;waveAssignmentMode&#x60; might differ. | [optional] 

## Methods

### NewOrganizationMaintenanceSettingsResponse

`func NewOrganizationMaintenanceSettingsResponse() *OrganizationMaintenanceSettingsResponse`

NewOrganizationMaintenanceSettingsResponse instantiates a new OrganizationMaintenanceSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMaintenanceSettingsResponseWithDefaults

`func NewOrganizationMaintenanceSettingsResponseWithDefaults() *OrganizationMaintenanceSettingsResponse`

NewOrganizationMaintenanceSettingsResponseWithDefaults instantiates a new OrganizationMaintenanceSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsResponse) GetEffectiveWaveAssignmentMode() string`

GetEffectiveWaveAssignmentMode returns the EffectiveWaveAssignmentMode field if non-nil, zero value otherwise.

### GetEffectiveWaveAssignmentModeOk

`func (o *OrganizationMaintenanceSettingsResponse) GetEffectiveWaveAssignmentModeOk() (*string, bool)`

GetEffectiveWaveAssignmentModeOk returns a tuple with the EffectiveWaveAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsResponse) SetEffectiveWaveAssignmentMode(v string)`

SetEffectiveWaveAssignmentMode sets EffectiveWaveAssignmentMode field to given value.

### HasEffectiveWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsResponse) HasEffectiveWaveAssignmentMode() bool`

HasEffectiveWaveAssignmentMode returns a boolean if a field has been set.

### SetEffectiveWaveAssignmentModeNil

`func (o *OrganizationMaintenanceSettingsResponse) SetEffectiveWaveAssignmentModeNil()`

SetEffectiveWaveAssignmentModeNil sets EffectiveWaveAssignmentMode to an explicit JSON null when marshaled, overriding any value previously set with SetEffectiveWaveAssignmentMode. Calling SetEffectiveWaveAssignmentMode again clears the null override.

### GetWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsResponse) GetWaveAssignmentMode() string`

GetWaveAssignmentMode returns the WaveAssignmentMode field if non-nil, zero value otherwise.

### GetWaveAssignmentModeOk

`func (o *OrganizationMaintenanceSettingsResponse) GetWaveAssignmentModeOk() (*string, bool)`

GetWaveAssignmentModeOk returns a tuple with the WaveAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsResponse) SetWaveAssignmentMode(v string)`

SetWaveAssignmentMode sets WaveAssignmentMode field to given value.

### HasWaveAssignmentMode

`func (o *OrganizationMaintenanceSettingsResponse) HasWaveAssignmentMode() bool`

HasWaveAssignmentMode returns a boolean if a field has been set.

### SetWaveAssignmentModeNil

`func (o *OrganizationMaintenanceSettingsResponse) SetWaveAssignmentModeNil()`

SetWaveAssignmentModeNil sets WaveAssignmentMode to an explicit JSON null when marshaled, overriding any value previously set with SetWaveAssignmentMode. Calling SetWaveAssignmentMode again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


