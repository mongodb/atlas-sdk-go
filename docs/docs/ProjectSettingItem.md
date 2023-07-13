# ProjectSettingItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Flag that indicates whether someone enabled the regionalized private endpoint setting for the specified project.  - Set this value to &#x60;true&#x60; to enable regionalized private endpoints. This allows you to create more than one private endpoint in a cloud provider region. You need to enable this setting to connect to multi-region and global MongoDB Cloud sharded clusters. Enabling regionalized private endpoints introduces the following limitations:   - Your applications must use the new connection strings for existing multi-region and global sharded clusters. This might cause downtime.   - Your MongoDB Cloud project can&#39;t contain replica sets nor can you create new replica sets in this project.    - You can&#39;t disable this setting if you have:     - more than one private endpoint in more than one region     - more than one private endpoint in one region and one private endpoint in one or more regions.  - Set this value to &#x60;false&#x60; to disable regionalized private endpoints. | 

## Methods

### NewProjectSettingItem

`func NewProjectSettingItem(enabled bool, ) *ProjectSettingItem`

NewProjectSettingItem instantiates a new ProjectSettingItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSettingItemWithDefaults

`func NewProjectSettingItemWithDefaults() *ProjectSettingItem`

NewProjectSettingItemWithDefaults instantiates a new ProjectSettingItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ProjectSettingItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProjectSettingItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProjectSettingItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


