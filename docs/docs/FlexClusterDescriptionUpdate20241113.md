# FlexClusterDescriptionUpdate20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the instance. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]

## Methods

### NewFlexClusterDescriptionUpdate20241113

`func NewFlexClusterDescriptionUpdate20241113() *FlexClusterDescriptionUpdate20241113`

NewFlexClusterDescriptionUpdate20241113 instantiates a new FlexClusterDescriptionUpdate20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexClusterDescriptionUpdate20241113WithDefaults

`func NewFlexClusterDescriptionUpdate20241113WithDefaults() *FlexClusterDescriptionUpdate20241113`

NewFlexClusterDescriptionUpdate20241113WithDefaults instantiates a new FlexClusterDescriptionUpdate20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlexClusterDescriptionUpdate20241113) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlexClusterDescriptionUpdate20241113) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlexClusterDescriptionUpdate20241113) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlexClusterDescriptionUpdate20241113) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetTags

`func (o *FlexClusterDescriptionUpdate20241113) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlexClusterDescriptionUpdate20241113) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlexClusterDescriptionUpdate20241113) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlexClusterDescriptionUpdate20241113) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *FlexClusterDescriptionUpdate20241113) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *FlexClusterDescriptionUpdate20241113) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *FlexClusterDescriptionUpdate20241113) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *FlexClusterDescriptionUpdate20241113) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


