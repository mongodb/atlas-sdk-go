# QueryStatsSummaryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Explanation of query shape insights availability limitations on the current tier. | [optional] [readonly] 
**QueryStatsCollection** | Pointer to [**QueryStatsCollectionResponse**](QueryStatsCollectionResponse.md) |  | [optional] 
**Scale** | Pointer to [**ApiClusterScale**](ApiClusterScale.md) |  | [optional] 
**Summaries** | Pointer to [**[]QueryStatsSummary**](QueryStatsSummary.md) | List of query shape statistic summaries from Query Shape Insights. | [optional] 

## Methods

### NewQueryStatsSummaryListResponse

`func NewQueryStatsSummaryListResponse() *QueryStatsSummaryListResponse`

NewQueryStatsSummaryListResponse instantiates a new QueryStatsSummaryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryStatsSummaryListResponseWithDefaults

`func NewQueryStatsSummaryListResponseWithDefaults() *QueryStatsSummaryListResponse`

NewQueryStatsSummaryListResponseWithDefaults instantiates a new QueryStatsSummaryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *QueryStatsSummaryListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QueryStatsSummaryListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QueryStatsSummaryListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QueryStatsSummaryListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *QueryStatsSummaryListResponse) SetMessageNil()`

SetMessageNil sets Message to an explicit JSON null when marshaled, overriding any value previously set with SetMessage. Calling SetMessage again clears the null override.

### GetQueryStatsCollection

`func (o *QueryStatsSummaryListResponse) GetQueryStatsCollection() QueryStatsCollectionResponse`

GetQueryStatsCollection returns the QueryStatsCollection field if non-nil, zero value otherwise.

### GetQueryStatsCollectionOk

`func (o *QueryStatsSummaryListResponse) GetQueryStatsCollectionOk() (*QueryStatsCollectionResponse, bool)`

GetQueryStatsCollectionOk returns a tuple with the QueryStatsCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatsCollection

`func (o *QueryStatsSummaryListResponse) SetQueryStatsCollection(v QueryStatsCollectionResponse)`

SetQueryStatsCollection sets QueryStatsCollection field to given value.

### HasQueryStatsCollection

`func (o *QueryStatsSummaryListResponse) HasQueryStatsCollection() bool`

HasQueryStatsCollection returns a boolean if a field has been set.

### SetQueryStatsCollectionNil

`func (o *QueryStatsSummaryListResponse) SetQueryStatsCollectionNil()`

SetQueryStatsCollectionNil sets QueryStatsCollection to an explicit JSON null when marshaled, overriding any value previously set with SetQueryStatsCollection. Calling SetQueryStatsCollection again clears the null override.

### GetScale

`func (o *QueryStatsSummaryListResponse) GetScale() ApiClusterScale`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *QueryStatsSummaryListResponse) GetScaleOk() (*ApiClusterScale, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *QueryStatsSummaryListResponse) SetScale(v ApiClusterScale)`

SetScale sets Scale field to given value.

### HasScale

`func (o *QueryStatsSummaryListResponse) HasScale() bool`

HasScale returns a boolean if a field has been set.

### SetScaleNil

`func (o *QueryStatsSummaryListResponse) SetScaleNil()`

SetScaleNil sets Scale to an explicit JSON null when marshaled, overriding any value previously set with SetScale. Calling SetScale again clears the null override.

### GetSummaries

`func (o *QueryStatsSummaryListResponse) GetSummaries() []QueryStatsSummary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *QueryStatsSummaryListResponse) GetSummariesOk() (*[]QueryStatsSummary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *QueryStatsSummaryListResponse) SetSummaries(v []QueryStatsSummary)`

SetSummaries sets Summaries field to given value.

### HasSummaries

`func (o *QueryStatsSummaryListResponse) HasSummaries() bool`

HasSummaries returns a boolean if a field has been set.

### SetSummariesNil

`func (o *QueryStatsSummaryListResponse) SetSummariesNil()`

SetSummariesNil sets Summaries to an explicit JSON null when marshaled, overriding any value previously set with SetSummaries. Calling SetSummaries again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


