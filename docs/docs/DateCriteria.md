# DateCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateField** | Pointer to **string** | Indexed database parameter that stores the date that determines when data moves to the online archive. MongoDB Cloud archives the data when the current date exceeds the date in this database parameter plus the number of days specified through the **expireAfterDays** parameter. Set this parameter when you set &#x60;\&quot;criteria.type\&quot; : \&quot;DATE\&quot;&#x60;. | [optional] 
**DateFormat** | Pointer to **string** | Syntax used to write the date after which data moves to the online archive. Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when **\&quot;criteria.type\&quot; : \&quot;DATE\&quot;**. You must set **\&quot;criteria.type\&quot; : \&quot;DATE\&quot;** if **\&quot;collectionType\&quot;: \&quot;TIMESERIES\&quot;**. | [optional] [default to "ISODATE"]
**ExpireAfterDays** | Pointer to **int** | Number of days after the value in the **criteria.dateField** when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set **\&quot;criteria.type\&quot; : \&quot;DATE\&quot;**. | [optional] 
**Type** | Pointer to **string** | Means by which MongoDB Cloud selects data to archive. Data can be chosen using the age of the data or a MongoDB query. **DATE** selects documents to archive based on a date. **CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn&#39;t support **CUSTOM** when &#x60;\&quot;collectionType\&quot;: \&quot;TIMESERIES\&quot;&#x60;. | [optional] 

## Methods

### NewDateCriteria

`func NewDateCriteria() *DateCriteria`

NewDateCriteria instantiates a new DateCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateCriteriaWithDefaults

`func NewDateCriteriaWithDefaults() *DateCriteria`

NewDateCriteriaWithDefaults instantiates a new DateCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateField

`func (o *DateCriteria) GetDateField() string`

GetDateField returns the DateField field if non-nil, zero value otherwise.

### GetDateFieldOk

`func (o *DateCriteria) GetDateFieldOk() (*string, bool)`

GetDateFieldOk returns a tuple with the DateField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateField

`func (o *DateCriteria) SetDateField(v string)`

SetDateField sets DateField field to given value.

### HasDateField

`func (o *DateCriteria) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetDateFormat

`func (o *DateCriteria) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *DateCriteria) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *DateCriteria) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *DateCriteria) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetExpireAfterDays

`func (o *DateCriteria) GetExpireAfterDays() int`

GetExpireAfterDays returns the ExpireAfterDays field if non-nil, zero value otherwise.

### GetExpireAfterDaysOk

`func (o *DateCriteria) GetExpireAfterDaysOk() (*int, bool)`

GetExpireAfterDaysOk returns a tuple with the ExpireAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfterDays

`func (o *DateCriteria) SetExpireAfterDays(v int)`

SetExpireAfterDays sets ExpireAfterDays field to given value.

### HasExpireAfterDays

`func (o *DateCriteria) HasExpireAfterDays() bool`

HasExpireAfterDays returns a boolean if a field has been set.

### GetType

`func (o *DateCriteria) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DateCriteria) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DateCriteria) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DateCriteria) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


