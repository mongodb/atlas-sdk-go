# AWSComputeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxInstanceSize** | Pointer to **string** | Maximum instance size to which your cluster can automatically scale. | [optional] 
**MinInstanceSize** | Pointer to **string** | Minimum instance size to which your cluster can automatically scale. | [optional] 

## Methods

### NewAWSComputeAutoScaling

`func NewAWSComputeAutoScaling() *AWSComputeAutoScaling`

NewAWSComputeAutoScaling instantiates a new AWSComputeAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSComputeAutoScalingWithDefaults

`func NewAWSComputeAutoScalingWithDefaults() *AWSComputeAutoScaling`

NewAWSComputeAutoScalingWithDefaults instantiates a new AWSComputeAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxInstanceSize

`func (o *AWSComputeAutoScaling) GetMaxInstanceSize() string`

GetMaxInstanceSize returns the MaxInstanceSize field if non-nil, zero value otherwise.

### GetMaxInstanceSizeOk

`func (o *AWSComputeAutoScaling) GetMaxInstanceSizeOk() (*string, bool)`

GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceSize

`func (o *AWSComputeAutoScaling) SetMaxInstanceSize(v string)`

SetMaxInstanceSize sets MaxInstanceSize field to given value.

### HasMaxInstanceSize

`func (o *AWSComputeAutoScaling) HasMaxInstanceSize() bool`

HasMaxInstanceSize returns a boolean if a field has been set.

### GetMinInstanceSize

`func (o *AWSComputeAutoScaling) GetMinInstanceSize() string`

GetMinInstanceSize returns the MinInstanceSize field if non-nil, zero value otherwise.

### GetMinInstanceSizeOk

`func (o *AWSComputeAutoScaling) GetMinInstanceSizeOk() (*string, bool)`

GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstanceSize

`func (o *AWSComputeAutoScaling) SetMinInstanceSize(v string)`

SetMinInstanceSize sets MinInstanceSize field to given value.

### HasMinInstanceSize

`func (o *AWSComputeAutoScaling) HasMinInstanceSize() bool`

HasMinInstanceSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


