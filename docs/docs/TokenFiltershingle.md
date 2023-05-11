# TokenFiltershingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxShingleSize** | **int** | Value that specifies the maximum number of tokens per shingle. This value must be greater than or equal to **minShingleSize**. | 
**MinShingleSize** | **int** | Value that specifies the minimum number of tokens per shingle. This value must be less than or equal to **maxShingleSize**. | 
**Type** | **string** | Human-readable label that identifies this token filter type. | 

## Methods

### NewTokenFiltershingle

`func NewTokenFiltershingle(maxShingleSize int, minShingleSize int, type_ string, ) *TokenFiltershingle`

NewTokenFiltershingle instantiates a new TokenFiltershingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFiltershingleWithDefaults

`func NewTokenFiltershingleWithDefaults() *TokenFiltershingle`

NewTokenFiltershingleWithDefaults instantiates a new TokenFiltershingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxShingleSize

`func (o *TokenFiltershingle) GetMaxShingleSize() int`

GetMaxShingleSize returns the MaxShingleSize field if non-nil, zero value otherwise.

### GetMaxShingleSizeOk

`func (o *TokenFiltershingle) GetMaxShingleSizeOk() (*int, bool)`

GetMaxShingleSizeOk returns a tuple with the MaxShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxShingleSize

`func (o *TokenFiltershingle) SetMaxShingleSize(v int)`

SetMaxShingleSize sets MaxShingleSize field to given value.


### GetMinShingleSize

`func (o *TokenFiltershingle) GetMinShingleSize() int`

GetMinShingleSize returns the MinShingleSize field if non-nil, zero value otherwise.

### GetMinShingleSizeOk

`func (o *TokenFiltershingle) GetMinShingleSizeOk() (*int, bool)`

GetMinShingleSizeOk returns a tuple with the MinShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinShingleSize

`func (o *TokenFiltershingle) SetMinShingleSize(v int)`

SetMinShingleSize sets MinShingleSize field to given value.


### GetType

`func (o *TokenFiltershingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenFiltershingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenFiltershingle) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


