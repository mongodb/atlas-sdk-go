# RestoreJobFileHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | Human-readable label that identifies the hashed file. | [optional] [readonly] 
**Hash** | Pointer to **string** | Hashed checksum that maps to the restore file. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**TypeName** | Pointer to **string** | Human-readable label that identifies the hashing algorithm used to compute the hash value. | [optional] [readonly] 

## Methods

### NewRestoreJobFileHash

`func NewRestoreJobFileHash() *RestoreJobFileHash`

NewRestoreJobFileHash instantiates a new RestoreJobFileHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreJobFileHashWithDefaults

`func NewRestoreJobFileHashWithDefaults() *RestoreJobFileHash`

NewRestoreJobFileHashWithDefaults instantiates a new RestoreJobFileHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *RestoreJobFileHash) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *RestoreJobFileHash) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *RestoreJobFileHash) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *RestoreJobFileHash) HasFileName() bool`

HasFileName returns a boolean if a field has been set.
### GetHash

`func (o *RestoreJobFileHash) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RestoreJobFileHash) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RestoreJobFileHash) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *RestoreJobFileHash) HasHash() bool`

HasHash returns a boolean if a field has been set.
### GetLinks

`func (o *RestoreJobFileHash) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RestoreJobFileHash) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RestoreJobFileHash) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RestoreJobFileHash) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetTypeName

`func (o *RestoreJobFileHash) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *RestoreJobFileHash) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *RestoreJobFileHash) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *RestoreJobFileHash) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


