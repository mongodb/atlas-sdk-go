# FederatedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | Email address of the MongoDB Cloud user linked to the federated organization. | 
**FederationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies the federation to which this MongoDB Cloud user belongs. | 
**FirstName** | **string** | First or given name that belongs to the MongoDB Cloud user. | 
**LastName** | **string** | Last name, family name, or surname that belongs to the MongoDB Cloud user. | 
**UserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this user. | [optional] [readonly] 

## Methods

### NewFederatedUser

`func NewFederatedUser(emailAddress string, federationSettingsId string, firstName string, lastName string, ) *FederatedUser`

NewFederatedUser instantiates a new FederatedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedUserWithDefaults

`func NewFederatedUserWithDefaults() *FederatedUser`

NewFederatedUserWithDefaults instantiates a new FederatedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *FederatedUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *FederatedUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *FederatedUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFederationSettingsId

`func (o *FederatedUser) GetFederationSettingsId() string`

GetFederationSettingsId returns the FederationSettingsId field if non-nil, zero value otherwise.

### GetFederationSettingsIdOk

`func (o *FederatedUser) GetFederationSettingsIdOk() (*string, bool)`

GetFederationSettingsIdOk returns a tuple with the FederationSettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSettingsId

`func (o *FederatedUser) SetFederationSettingsId(v string)`

SetFederationSettingsId sets FederationSettingsId field to given value.


### GetFirstName

`func (o *FederatedUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FederatedUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FederatedUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *FederatedUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FederatedUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FederatedUser) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserId

`func (o *FederatedUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FederatedUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FederatedUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FederatedUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


