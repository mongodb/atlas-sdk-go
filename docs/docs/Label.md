# Label

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key for the metadata file that MongoDB Cloud uploads to the bucket when the export job finishes. | [optional] 
**Value** | Pointer to **string** | Value for the key to include in file that MongoDB Cloud uploads to the bucket when the export job finishes. | [optional] 

## Methods

### NewLabel

`func NewLabel() *Label`

NewLabel instantiates a new Label object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelWithDefaults

`func NewLabelWithDefaults() *Label`

NewLabelWithDefaults instantiates a new Label object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Label) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Label) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Label) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Label) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *Label) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Label) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Label) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Label) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


