# SkuResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Human-readable short summary of what this SKU represents. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique string that identifies the SKU. | [optional] [readonly] 

## Methods

### NewSkuResponse

`func NewSkuResponse() *SkuResponse`

NewSkuResponse instantiates a new SkuResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuResponseWithDefaults

`func NewSkuResponseWithDefaults() *SkuResponse`

NewSkuResponseWithDefaults instantiates a new SkuResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SkuResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkuResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkuResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SkuResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetId

`func (o *SkuResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SkuResponse) HasId() bool`

HasId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


