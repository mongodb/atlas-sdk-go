# PaginatedApiAtlasFlexBackupSnapshot20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]FlexBackupSnapshot20241113**](FlexBackupSnapshot20241113.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Total number of documents available. MongoDB Cloud omits this value if &#x60;includeCount&#x60; is set to &#x60;false&#x60;. | [optional] [readonly] 

## Methods

### NewPaginatedApiAtlasFlexBackupSnapshot20241113

`func NewPaginatedApiAtlasFlexBackupSnapshot20241113() *PaginatedApiAtlasFlexBackupSnapshot20241113`

NewPaginatedApiAtlasFlexBackupSnapshot20241113 instantiates a new PaginatedApiAtlasFlexBackupSnapshot20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedApiAtlasFlexBackupSnapshot20241113WithDefaults

`func NewPaginatedApiAtlasFlexBackupSnapshot20241113WithDefaults() *PaginatedApiAtlasFlexBackupSnapshot20241113`

NewPaginatedApiAtlasFlexBackupSnapshot20241113WithDefaults instantiates a new PaginatedApiAtlasFlexBackupSnapshot20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetResults

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) GetResults() []FlexBackupSnapshot20241113`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) GetResultsOk() (*[]FlexBackupSnapshot20241113, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) SetResults(v []FlexBackupSnapshot20241113)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) HasResults() bool`

HasResults returns a boolean if a field has been set.
### GetTotalCount

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedApiAtlasFlexBackupSnapshot20241113) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


