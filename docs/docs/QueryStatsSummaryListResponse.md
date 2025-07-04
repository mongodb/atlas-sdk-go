# QueryStatsSummaryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


