# AuthenticatedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Email address that represents the username of the MongoDB Cloud user. | [optional] [readonly] 

## Methods

### NewAuthenticatedUser

`func NewAuthenticatedUser() *AuthenticatedUser`

NewAuthenticatedUser instantiates a new AuthenticatedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatedUserWithDefaults

`func NewAuthenticatedUserWithDefaults() *AuthenticatedUser`

NewAuthenticatedUserWithDefaults instantiates a new AuthenticatedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AuthenticatedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthenticatedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthenticatedUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuthenticatedUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AuthenticatedUser) SetUsernameNil()`

SetUsernameNil sets Username to an explicit JSON null when marshaled, overriding any value previously set with SetUsername. Calling SetUsername again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


