# ProviderInstanceSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableRegions** | Pointer to [**[]AvailableRegion**](AvailableRegion.md) | List of regions that this cloud provider supports for this instance size. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the instance size or cluster tier. | [optional] [readonly] 

## Methods

### NewProviderInstanceSize

`func NewProviderInstanceSize() *ProviderInstanceSize`

NewProviderInstanceSize instantiates a new ProviderInstanceSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderInstanceSizeWithDefaults

`func NewProviderInstanceSizeWithDefaults() *ProviderInstanceSize`

NewProviderInstanceSizeWithDefaults instantiates a new ProviderInstanceSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableRegions

`func (o *ProviderInstanceSize) GetAvailableRegions() []AvailableRegion`

GetAvailableRegions returns the AvailableRegions field if non-nil, zero value otherwise.

### GetAvailableRegionsOk

`func (o *ProviderInstanceSize) GetAvailableRegionsOk() (*[]AvailableRegion, bool)`

GetAvailableRegionsOk returns a tuple with the AvailableRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRegions

`func (o *ProviderInstanceSize) SetAvailableRegions(v []AvailableRegion)`

SetAvailableRegions sets AvailableRegions field to given value.

### HasAvailableRegions

`func (o *ProviderInstanceSize) HasAvailableRegions() bool`

HasAvailableRegions returns a boolean if a field has been set.

### GetName

`func (o *ProviderInstanceSize) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderInstanceSize) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderInstanceSize) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderInstanceSize) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


