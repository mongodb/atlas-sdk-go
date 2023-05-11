# PemFileInfoView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]X509Certificate**](X509Certificate.md) | List of certificates in the file. | [optional] 
**FileName** | Pointer to **string** | Human-readable label given to the file. | [optional] 

## Methods

### NewPemFileInfoView

`func NewPemFileInfoView() *PemFileInfoView`

NewPemFileInfoView instantiates a new PemFileInfoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPemFileInfoViewWithDefaults

`func NewPemFileInfoViewWithDefaults() *PemFileInfoView`

NewPemFileInfoViewWithDefaults instantiates a new PemFileInfoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *PemFileInfoView) GetCertificates() []X509Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *PemFileInfoView) GetCertificatesOk() (*[]X509Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *PemFileInfoView) SetCertificates(v []X509Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *PemFileInfoView) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetFileName

`func (o *PemFileInfoView) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *PemFileInfoView) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *PemFileInfoView) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *PemFileInfoView) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


