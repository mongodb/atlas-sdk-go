# PaginatedShardKeyAnalysisOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | [**[]ShardKeyAnalysisOperationResponse**](ShardKeyAnalysisOperationResponse.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [readonly] 
**TotalCount** | Pointer to **int** | Total number of documents available. When &#x60;includeCount&#x60; is set to &#x60;false&#x60;, MongoDB Cloud may omit this value or return it when the count is available without additional calculation. The total number is an estimate and may not be exact. | [optional] [readonly] 

## Methods

### NewPaginatedShardKeyAnalysisOperationResponse

`func NewPaginatedShardKeyAnalysisOperationResponse(results []ShardKeyAnalysisOperationResponse, ) *PaginatedShardKeyAnalysisOperationResponse`

NewPaginatedShardKeyAnalysisOperationResponse instantiates a new PaginatedShardKeyAnalysisOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedShardKeyAnalysisOperationResponseWithDefaults

`func NewPaginatedShardKeyAnalysisOperationResponseWithDefaults() *PaginatedShardKeyAnalysisOperationResponse`

NewPaginatedShardKeyAnalysisOperationResponseWithDefaults instantiates a new PaginatedShardKeyAnalysisOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedShardKeyAnalysisOperationResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedShardKeyAnalysisOperationResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedShardKeyAnalysisOperationResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedShardKeyAnalysisOperationResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *PaginatedShardKeyAnalysisOperationResponse) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetResults

`func (o *PaginatedShardKeyAnalysisOperationResponse) GetResults() []ShardKeyAnalysisOperationResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedShardKeyAnalysisOperationResponse) GetResultsOk() (*[]ShardKeyAnalysisOperationResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedShardKeyAnalysisOperationResponse) SetResults(v []ShardKeyAnalysisOperationResponse)`

SetResults sets Results field to given value.

### GetTotalCount

`func (o *PaginatedShardKeyAnalysisOperationResponse) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedShardKeyAnalysisOperationResponse) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedShardKeyAnalysisOperationResponse) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedShardKeyAnalysisOperationResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *PaginatedShardKeyAnalysisOperationResponse) SetTotalCountNil()`

SetTotalCountNil sets TotalCount to an explicit JSON null when marshaled, overriding any value previously set with SetTotalCount. Calling SetTotalCount again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


