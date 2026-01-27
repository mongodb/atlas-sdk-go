# ActivityFeedLinkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | **string** | Shareable link to the activity feed with pre-applied filters. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewActivityFeedLinkResponse

`func NewActivityFeedLinkResponse(link string, ) *ActivityFeedLinkResponse`

NewActivityFeedLinkResponse instantiates a new ActivityFeedLinkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityFeedLinkResponseWithDefaults

`func NewActivityFeedLinkResponseWithDefaults() *ActivityFeedLinkResponse`

NewActivityFeedLinkResponseWithDefaults instantiates a new ActivityFeedLinkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *ActivityFeedLinkResponse) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ActivityFeedLinkResponse) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ActivityFeedLinkResponse) SetLink(v string)`

SetLink sets Link field to given value.

### GetLinks

`func (o *ActivityFeedLinkResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActivityFeedLinkResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActivityFeedLinkResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActivityFeedLinkResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


