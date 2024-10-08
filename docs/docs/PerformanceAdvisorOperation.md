# PerformanceAdvisorOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Predicates** | Pointer to [**[]any**](any.md) | List that contains the search criteria that the query uses. To use the values in key-value pairs in these predicates requires **Project Data Access Read Only** permissions or greater. Otherwise, MongoDB Cloud redacts these values. | [optional] [readonly] 
**Stats** | Pointer to [**PerformanceAdvisorOpStats**](PerformanceAdvisorOpStats.md) |  | [optional] 

## Methods

### NewPerformanceAdvisorOperation

`func NewPerformanceAdvisorOperation() *PerformanceAdvisorOperation`

NewPerformanceAdvisorOperation instantiates a new PerformanceAdvisorOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorOperationWithDefaults

`func NewPerformanceAdvisorOperationWithDefaults() *PerformanceAdvisorOperation`

NewPerformanceAdvisorOperationWithDefaults instantiates a new PerformanceAdvisorOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredicates

`func (o *PerformanceAdvisorOperation) GetPredicates() []any`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *PerformanceAdvisorOperation) GetPredicatesOk() (*[]any, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *PerformanceAdvisorOperation) SetPredicates(v []any)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *PerformanceAdvisorOperation) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.
### GetStats

`func (o *PerformanceAdvisorOperation) GetStats() PerformanceAdvisorOpStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *PerformanceAdvisorOperation) GetStatsOk() (*PerformanceAdvisorOpStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *PerformanceAdvisorOperation) SetStats(v PerformanceAdvisorOpStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *PerformanceAdvisorOperation) HasStats() bool`

HasStats returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


