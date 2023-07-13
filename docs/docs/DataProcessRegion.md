# DataProcessRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the Cloud service provider where you wish to store your archived data. | [optional] 
**Region** | Pointer to **string** | Human-readable label that identifies the geographic location of the region where you wish to store your archived data. | [optional] 

## Methods

### NewDataProcessRegion

`func NewDataProcessRegion() *DataProcessRegion`

NewDataProcessRegion instantiates a new DataProcessRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProcessRegionWithDefaults

`func NewDataProcessRegionWithDefaults() *DataProcessRegion`

NewDataProcessRegionWithDefaults instantiates a new DataProcessRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DataProcessRegion) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DataProcessRegion) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DataProcessRegion) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DataProcessRegion) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetRegion

`func (o *DataProcessRegion) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DataProcessRegion) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DataProcessRegion) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DataProcessRegion) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


