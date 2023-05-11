# GCPComputeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxInstanceSize** | Pointer to **string** | Maximum instance size to which your cluster can automatically scale. | [optional] 
**MinInstanceSize** | Pointer to **string** | Minimum instance size to which your cluster can automatically scale. | [optional] 

## Methods

### NewGCPComputeAutoScaling

`func NewGCPComputeAutoScaling() *GCPComputeAutoScaling`

NewGCPComputeAutoScaling instantiates a new GCPComputeAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPComputeAutoScalingWithDefaults

`func NewGCPComputeAutoScalingWithDefaults() *GCPComputeAutoScaling`

NewGCPComputeAutoScalingWithDefaults instantiates a new GCPComputeAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxInstanceSize

`func (o *GCPComputeAutoScaling) GetMaxInstanceSize() string`

GetMaxInstanceSize returns the MaxInstanceSize field if non-nil, zero value otherwise.

### GetMaxInstanceSizeOk

`func (o *GCPComputeAutoScaling) GetMaxInstanceSizeOk() (*string, bool)`

GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceSize

`func (o *GCPComputeAutoScaling) SetMaxInstanceSize(v string)`

SetMaxInstanceSize sets MaxInstanceSize field to given value.

### HasMaxInstanceSize

`func (o *GCPComputeAutoScaling) HasMaxInstanceSize() bool`

HasMaxInstanceSize returns a boolean if a field has been set.

### GetMinInstanceSize

`func (o *GCPComputeAutoScaling) GetMinInstanceSize() string`

GetMinInstanceSize returns the MinInstanceSize field if non-nil, zero value otherwise.

### GetMinInstanceSizeOk

`func (o *GCPComputeAutoScaling) GetMinInstanceSizeOk() (*string, bool)`

GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstanceSize

`func (o *GCPComputeAutoScaling) SetMinInstanceSize(v string)`

SetMinInstanceSize sets MinInstanceSize field to given value.

### HasMinInstanceSize

`func (o *GCPComputeAutoScaling) HasMinInstanceSize() bool`

HasMinInstanceSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


