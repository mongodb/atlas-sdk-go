# AvailableCloudProviderRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | Flag that indicates whether the cloud provider sets this region as its default. AWS defaults to US_EAST_1, GCP defaults to CENTRAL_US, and AZURE defaults to US_WEST_2. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the supported region. | [optional] [readonly] 

## Methods

### NewAvailableCloudProviderRegion

`func NewAvailableCloudProviderRegion() *AvailableCloudProviderRegion`

NewAvailableCloudProviderRegion instantiates a new AvailableCloudProviderRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableCloudProviderRegionWithDefaults

`func NewAvailableCloudProviderRegionWithDefaults() *AvailableCloudProviderRegion`

NewAvailableCloudProviderRegionWithDefaults instantiates a new AvailableCloudProviderRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *AvailableCloudProviderRegion) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *AvailableCloudProviderRegion) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *AvailableCloudProviderRegion) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *AvailableCloudProviderRegion) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *AvailableCloudProviderRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableCloudProviderRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableCloudProviderRegion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AvailableCloudProviderRegion) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


