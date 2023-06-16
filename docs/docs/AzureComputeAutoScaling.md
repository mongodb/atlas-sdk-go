# AzureComputeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxInstanceSize** | Pointer to **string** | Maximum instance size to which your cluster can automatically scale. | [optional] 
**MinInstanceSize** | Pointer to **string** | Minimum instance size to which your cluster can automatically scale. | [optional] 

## Methods

### NewAzureComputeAutoScaling

`func NewAzureComputeAutoScaling() *AzureComputeAutoScaling`

NewAzureComputeAutoScaling instantiates a new AzureComputeAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeAutoScalingWithDefaults

`func NewAzureComputeAutoScalingWithDefaults() *AzureComputeAutoScaling`

NewAzureComputeAutoScalingWithDefaults instantiates a new AzureComputeAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxInstanceSize

`func (o *AzureComputeAutoScaling) GetMaxInstanceSize() string`

GetMaxInstanceSize returns the MaxInstanceSize field if non-nil, zero value otherwise.

### GetMaxInstanceSizeOk

`func (o *AzureComputeAutoScaling) GetMaxInstanceSizeOk() (*string, bool)`

GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceSize

`func (o *AzureComputeAutoScaling) SetMaxInstanceSize(v string)`

SetMaxInstanceSize sets MaxInstanceSize field to given value.

### HasMaxInstanceSize

`func (o *AzureComputeAutoScaling) HasMaxInstanceSize() bool`

HasMaxInstanceSize returns a boolean if a field has been set.

### GetMinInstanceSize

`func (o *AzureComputeAutoScaling) GetMinInstanceSize() string`

GetMinInstanceSize returns the MinInstanceSize field if non-nil, zero value otherwise.

### GetMinInstanceSizeOk

`func (o *AzureComputeAutoScaling) GetMinInstanceSizeOk() (*string, bool)`

GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstanceSize

`func (o *AzureComputeAutoScaling) SetMinInstanceSize(v string)`

SetMinInstanceSize sets MinInstanceSize field to given value.

### HasMinInstanceSize

`func (o *AzureComputeAutoScaling) HasMinInstanceSize() bool`

HasMinInstanceSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


