# PerformanceAdvisorSlowQueryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlowQueries** | Pointer to [**[]PerformanceAdvisorSlowQuery**](PerformanceAdvisorSlowQuery.md) | List of operations that the Performance Advisor detected that took longer to execute than a specified threshold. | [optional] [readonly] 

## Methods

### NewPerformanceAdvisorSlowQueryList

`func NewPerformanceAdvisorSlowQueryList() *PerformanceAdvisorSlowQueryList`

NewPerformanceAdvisorSlowQueryList instantiates a new PerformanceAdvisorSlowQueryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorSlowQueryListWithDefaults

`func NewPerformanceAdvisorSlowQueryListWithDefaults() *PerformanceAdvisorSlowQueryList`

NewPerformanceAdvisorSlowQueryListWithDefaults instantiates a new PerformanceAdvisorSlowQueryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlowQueries

`func (o *PerformanceAdvisorSlowQueryList) GetSlowQueries() []PerformanceAdvisorSlowQuery`

GetSlowQueries returns the SlowQueries field if non-nil, zero value otherwise.

### GetSlowQueriesOk

`func (o *PerformanceAdvisorSlowQueryList) GetSlowQueriesOk() (*[]PerformanceAdvisorSlowQuery, bool)`

GetSlowQueriesOk returns a tuple with the SlowQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowQueries

`func (o *PerformanceAdvisorSlowQueryList) SetSlowQueries(v []PerformanceAdvisorSlowQuery)`

SetSlowQueries sets SlowQueries field to given value.

### HasSlowQueries

`func (o *PerformanceAdvisorSlowQueryList) HasSlowQueries() bool`

HasSlowQueries returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


