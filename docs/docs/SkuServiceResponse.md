# SkuServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable name of the SKU service as shown on invoices. | [optional] [readonly] 

## Methods

### NewSkuServiceResponse

`func NewSkuServiceResponse() *SkuServiceResponse`

NewSkuServiceResponse instantiates a new SkuServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuServiceResponseWithDefaults

`func NewSkuServiceResponseWithDefaults() *SkuServiceResponse`

NewSkuServiceResponseWithDefaults instantiates a new SkuServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SkuServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkuServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkuServiceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SkuServiceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SkuServiceResponse) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


