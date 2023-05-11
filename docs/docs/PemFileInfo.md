# PemFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]X509Certificate**](X509Certificate.md) |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 

## Methods

### NewPemFileInfo

`func NewPemFileInfo() *PemFileInfo`

NewPemFileInfo instantiates a new PemFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPemFileInfoWithDefaults

`func NewPemFileInfoWithDefaults() *PemFileInfo`

NewPemFileInfoWithDefaults instantiates a new PemFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *PemFileInfo) GetCertificates() []X509Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *PemFileInfo) GetCertificatesOk() (*[]X509Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *PemFileInfo) SetCertificates(v []X509Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *PemFileInfo) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetFileName

`func (o *PemFileInfo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *PemFileInfo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *PemFileInfo) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *PemFileInfo) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


