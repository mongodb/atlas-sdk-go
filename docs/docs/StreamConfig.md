# StreamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Tier** | Pointer to **string** | Selected tier for the Stream Instance. Configures Memory / VCPU allowances. | [optional] 

## Methods

### NewStreamConfig

`func NewStreamConfig() *StreamConfig`

NewStreamConfig instantiates a new StreamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConfigWithDefaults

`func NewStreamConfigWithDefaults() *StreamConfig`

NewStreamConfigWithDefaults instantiates a new StreamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamConfig) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamConfig) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamConfig) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetTier

`func (o *StreamConfig) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *StreamConfig) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *StreamConfig) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *StreamConfig) HasTier() bool`

HasTier returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


