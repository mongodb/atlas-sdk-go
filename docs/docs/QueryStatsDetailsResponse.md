# QueryStatsDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstSeen** | Pointer to [**QueryShapeSeenMetadata**](QueryShapeSeenMetadata.md) |  | [optional] 
**LastSeen** | Pointer to [**QueryShapeSeenMetadata**](QueryShapeSeenMetadata.md) |  | [optional] 
**QueryStats** | Pointer to [**QueryStatsSummary**](QueryStatsSummary.md) |  | [optional] 

## Methods

### NewQueryStatsDetailsResponse

`func NewQueryStatsDetailsResponse() *QueryStatsDetailsResponse`

NewQueryStatsDetailsResponse instantiates a new QueryStatsDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryStatsDetailsResponseWithDefaults

`func NewQueryStatsDetailsResponseWithDefaults() *QueryStatsDetailsResponse`

NewQueryStatsDetailsResponseWithDefaults instantiates a new QueryStatsDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstSeen

`func (o *QueryStatsDetailsResponse) GetFirstSeen() QueryShapeSeenMetadata`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *QueryStatsDetailsResponse) GetFirstSeenOk() (*QueryShapeSeenMetadata, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *QueryStatsDetailsResponse) SetFirstSeen(v QueryShapeSeenMetadata)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *QueryStatsDetailsResponse) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.
### GetLastSeen

`func (o *QueryStatsDetailsResponse) GetLastSeen() QueryShapeSeenMetadata`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *QueryStatsDetailsResponse) GetLastSeenOk() (*QueryShapeSeenMetadata, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *QueryStatsDetailsResponse) SetLastSeen(v QueryShapeSeenMetadata)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *QueryStatsDetailsResponse) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.
### GetQueryStats

`func (o *QueryStatsDetailsResponse) GetQueryStats() QueryStatsSummary`

GetQueryStats returns the QueryStats field if non-nil, zero value otherwise.

### GetQueryStatsOk

`func (o *QueryStatsDetailsResponse) GetQueryStatsOk() (*QueryStatsSummary, bool)`

GetQueryStatsOk returns a tuple with the QueryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStats

`func (o *QueryStatsDetailsResponse) SetQueryStats(v QueryStatsSummary)`

SetQueryStats sets QueryStats field to given value.

### HasQueryStats

`func (o *QueryStatsDetailsResponse) HasQueryStats() bool`

HasQueryStats returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


