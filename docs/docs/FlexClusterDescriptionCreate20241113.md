# FlexClusterDescriptionCreate20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies the instance. | 
**ProviderSettings** | [**FlexProviderSettingsCreate20241113**](FlexProviderSettingsCreate20241113.md) |  | 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the instance. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]

## Methods

### NewFlexClusterDescriptionCreate20241113

`func NewFlexClusterDescriptionCreate20241113(name string, providerSettings FlexProviderSettingsCreate20241113, ) *FlexClusterDescriptionCreate20241113`

NewFlexClusterDescriptionCreate20241113 instantiates a new FlexClusterDescriptionCreate20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexClusterDescriptionCreate20241113WithDefaults

`func NewFlexClusterDescriptionCreate20241113WithDefaults() *FlexClusterDescriptionCreate20241113`

NewFlexClusterDescriptionCreate20241113WithDefaults instantiates a new FlexClusterDescriptionCreate20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlexClusterDescriptionCreate20241113) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlexClusterDescriptionCreate20241113) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlexClusterDescriptionCreate20241113) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlexClusterDescriptionCreate20241113) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetName

`func (o *FlexClusterDescriptionCreate20241113) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlexClusterDescriptionCreate20241113) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlexClusterDescriptionCreate20241113) SetName(v string)`

SetName sets Name field to given value.

### GetProviderSettings

`func (o *FlexClusterDescriptionCreate20241113) GetProviderSettings() FlexProviderSettingsCreate20241113`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *FlexClusterDescriptionCreate20241113) GetProviderSettingsOk() (*FlexProviderSettingsCreate20241113, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *FlexClusterDescriptionCreate20241113) SetProviderSettings(v FlexProviderSettingsCreate20241113)`

SetProviderSettings sets ProviderSettings field to given value.

### GetTags

`func (o *FlexClusterDescriptionCreate20241113) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlexClusterDescriptionCreate20241113) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlexClusterDescriptionCreate20241113) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlexClusterDescriptionCreate20241113) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *FlexClusterDescriptionCreate20241113) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *FlexClusterDescriptionCreate20241113) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *FlexClusterDescriptionCreate20241113) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *FlexClusterDescriptionCreate20241113) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


