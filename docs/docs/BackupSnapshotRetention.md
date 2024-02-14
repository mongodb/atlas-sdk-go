# BackupSnapshotRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**RetentionUnit** | **string** | Quantity of time in which MongoDB Cloud measures snapshot retention. | 
**RetentionValue** | **int** | Number that indicates the amount of days, weeks, months, or years that MongoDB Cloud retains the snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items. If the hourly policy item specifies a retention of two days, specify two days or greater for the retention of the weekly policy item. | 

## Methods

### NewBackupSnapshotRetention

`func NewBackupSnapshotRetention(retentionUnit string, retentionValue int, ) *BackupSnapshotRetention`

NewBackupSnapshotRetention instantiates a new BackupSnapshotRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSnapshotRetentionWithDefaults

`func NewBackupSnapshotRetentionWithDefaults() *BackupSnapshotRetention`

NewBackupSnapshotRetentionWithDefaults instantiates a new BackupSnapshotRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BackupSnapshotRetention) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BackupSnapshotRetention) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BackupSnapshotRetention) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BackupSnapshotRetention) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRetentionUnit

`func (o *BackupSnapshotRetention) GetRetentionUnit() string`

GetRetentionUnit returns the RetentionUnit field if non-nil, zero value otherwise.

### GetRetentionUnitOk

`func (o *BackupSnapshotRetention) GetRetentionUnitOk() (*string, bool)`

GetRetentionUnitOk returns a tuple with the RetentionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionUnit

`func (o *BackupSnapshotRetention) SetRetentionUnit(v string)`

SetRetentionUnit sets RetentionUnit field to given value.

### GetRetentionValue

`func (o *BackupSnapshotRetention) GetRetentionValue() int`

GetRetentionValue returns the RetentionValue field if non-nil, zero value otherwise.

### GetRetentionValueOk

`func (o *BackupSnapshotRetention) GetRetentionValueOk() (*int, bool)`

GetRetentionValueOk returns a tuple with the RetentionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionValue

`func (o *BackupSnapshotRetention) SetRetentionValue(v int)`

SetRetentionValue sets RetentionValue field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


