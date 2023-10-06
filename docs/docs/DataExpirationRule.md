# DataExpirationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireAfterDays** | Pointer to **int** | Number of days used in the date criteria for nominating documents for deletion. | [optional] 

## Methods

### NewDataExpirationRule

`func NewDataExpirationRule() *DataExpirationRule`

NewDataExpirationRule instantiates a new DataExpirationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExpirationRuleWithDefaults

`func NewDataExpirationRuleWithDefaults() *DataExpirationRule`

NewDataExpirationRuleWithDefaults instantiates a new DataExpirationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireAfterDays

`func (o *DataExpirationRule) GetExpireAfterDays() int`

GetExpireAfterDays returns the ExpireAfterDays field if non-nil, zero value otherwise.

### GetExpireAfterDaysOk

`func (o *DataExpirationRule) GetExpireAfterDaysOk() (*int, bool)`

GetExpireAfterDaysOk returns a tuple with the ExpireAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfterDays

`func (o *DataExpirationRule) SetExpireAfterDays(v int)`

SetExpireAfterDays sets ExpireAfterDays field to given value.

### HasExpireAfterDays

`func (o *DataExpirationRule) HasExpireAfterDays() bool`

HasExpireAfterDays returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


