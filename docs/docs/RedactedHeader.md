# RedactedHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Header name. | [readonly] 
**Value** | **string** | Redacted header value. | [readonly] 

## Methods

### NewRedactedHeader

`func NewRedactedHeader(name string, value string, ) *RedactedHeader`

NewRedactedHeader instantiates a new RedactedHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedactedHeaderWithDefaults

`func NewRedactedHeaderWithDefaults() *RedactedHeader`

NewRedactedHeaderWithDefaults instantiates a new RedactedHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RedactedHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedactedHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedactedHeader) SetName(v string)`

SetName sets Name field to given value.

### GetValue

`func (o *RedactedHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RedactedHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RedactedHeader) SetValue(v string)`

SetValue sets Value field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


