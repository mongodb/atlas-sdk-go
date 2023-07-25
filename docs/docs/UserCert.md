# UserCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique 24-hexadecimal character string that identifies this certificate. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this certificate. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MonthsUntilExpiration** | Pointer to **int** | Number of months that the certificate remains valid until it expires. | [optional] [default to 3]
**NotAfter** | Pointer to **time.Time** | Date and time when this certificate expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Subject** | Pointer to **string** | Subject Alternative Name associated with this certificate. This parameter expresses its value as a distinguished name as defined in [RFC 2253](https://tools.ietf.org/html/2253). | [optional] [readonly] 

## Methods

### NewUserCert

`func NewUserCert() *UserCert`

NewUserCert instantiates a new UserCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCertWithDefaults

`func NewUserCertWithDefaults() *UserCert`

NewUserCertWithDefaults instantiates a new UserCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserCert) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCert) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCert) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserCert) HasId() bool`

HasId returns a boolean if a field has been set.
### GetCreatedAt

`func (o *UserCert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserCert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserCert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserCert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetGroupId

`func (o *UserCert) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserCert) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserCert) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UserCert) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetLinks

`func (o *UserCert) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserCert) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserCert) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserCert) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMonthsUntilExpiration

`func (o *UserCert) GetMonthsUntilExpiration() int`

GetMonthsUntilExpiration returns the MonthsUntilExpiration field if non-nil, zero value otherwise.

### GetMonthsUntilExpirationOk

`func (o *UserCert) GetMonthsUntilExpirationOk() (*int, bool)`

GetMonthsUntilExpirationOk returns a tuple with the MonthsUntilExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthsUntilExpiration

`func (o *UserCert) SetMonthsUntilExpiration(v int)`

SetMonthsUntilExpiration sets MonthsUntilExpiration field to given value.

### HasMonthsUntilExpiration

`func (o *UserCert) HasMonthsUntilExpiration() bool`

HasMonthsUntilExpiration returns a boolean if a field has been set.
### GetNotAfter

`func (o *UserCert) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *UserCert) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *UserCert) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *UserCert) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.
### GetSubject

`func (o *UserCert) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UserCert) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UserCert) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UserCert) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


