# X509Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewX509Certificate

`func NewX509Certificate() *X509Certificate`

NewX509Certificate instantiates a new X509Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateWithDefaults

`func NewX509CertificateWithDefaults() *X509Certificate`

NewX509CertificateWithDefaults instantiates a new X509Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *X509Certificate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *X509Certificate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *X509Certificate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *X509Certificate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetNotAfter

`func (o *X509Certificate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *X509Certificate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *X509Certificate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *X509Certificate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *X509Certificate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *X509Certificate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


