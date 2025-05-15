# DatasetRetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModifiedDate** | Pointer to **time.Time** | Date when retention policy was last modified. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Units** | **string** | Quantity of time in which the Data Lake Pipeline measures dataset retention. | 
**Value** | **int** | Number that indicates the amount of days, weeks, or months that the Data Lake Pipeline will retain datasets. | 

## Methods

### NewDatasetRetentionPolicy

`func NewDatasetRetentionPolicy(units string, value int, ) *DatasetRetentionPolicy`

NewDatasetRetentionPolicy instantiates a new DatasetRetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetRetentionPolicyWithDefaults

`func NewDatasetRetentionPolicyWithDefaults() *DatasetRetentionPolicy`

NewDatasetRetentionPolicyWithDefaults instantiates a new DatasetRetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModifiedDate

`func (o *DatasetRetentionPolicy) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *DatasetRetentionPolicy) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *DatasetRetentionPolicy) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *DatasetRetentionPolicy) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.
### GetUnits

`func (o *DatasetRetentionPolicy) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *DatasetRetentionPolicy) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *DatasetRetentionPolicy) SetUnits(v string)`

SetUnits sets Units field to given value.

### GetValue

`func (o *DatasetRetentionPolicy) GetValue() int`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DatasetRetentionPolicy) GetValueOk() (*int, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DatasetRetentionPolicy) SetValue(v int)`

SetValue sets Value field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


