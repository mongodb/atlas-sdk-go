# StreamsDataProcessRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Label that identifies the cloud service provider where MongoDB Cloud performs stream processing. Currently, this parameter only supports AWS and AZURE. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Region** | **string** | Name of the cloud provider region hosting Atlas Stream Processing. | 

## Methods

### NewStreamsDataProcessRegion

`func NewStreamsDataProcessRegion(cloudProvider string, region string, ) *StreamsDataProcessRegion`

NewStreamsDataProcessRegion instantiates a new StreamsDataProcessRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsDataProcessRegionWithDefaults

`func NewStreamsDataProcessRegionWithDefaults() *StreamsDataProcessRegion`

NewStreamsDataProcessRegionWithDefaults instantiates a new StreamsDataProcessRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *StreamsDataProcessRegion) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *StreamsDataProcessRegion) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *StreamsDataProcessRegion) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetLinks

`func (o *StreamsDataProcessRegion) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsDataProcessRegion) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsDataProcessRegion) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsDataProcessRegion) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRegion

`func (o *StreamsDataProcessRegion) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsDataProcessRegion) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsDataProcessRegion) SetRegion(v string)`

SetRegion sets Region field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


