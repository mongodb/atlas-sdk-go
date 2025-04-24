# X509CertificateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Certificate content. | [optional] 
**NotAfter** | Pointer to **time.Time** | Latest date that the certificate is valid. | [optional] 
**NotBefore** | Pointer to **time.Time** | Earliest date that the certificate is valid. | [optional] 

## Methods

### NewX509CertificateUpdate

`func NewX509CertificateUpdate() *X509CertificateUpdate`

NewX509CertificateUpdate instantiates a new X509CertificateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateUpdateWithDefaults

`func NewX509CertificateUpdateWithDefaults() *X509CertificateUpdate`

NewX509CertificateUpdateWithDefaults instantiates a new X509CertificateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *X509CertificateUpdate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *X509CertificateUpdate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *X509CertificateUpdate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *X509CertificateUpdate) HasContent() bool`

HasContent returns a boolean if a field has been set.
### GetNotAfter

`func (o *X509CertificateUpdate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *X509CertificateUpdate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *X509CertificateUpdate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *X509CertificateUpdate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.
### GetNotBefore

`func (o *X509CertificateUpdate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *X509CertificateUpdate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *X509CertificateUpdate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *X509CertificateUpdate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


