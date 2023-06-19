# CloudAppUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | Two alphabet characters that identifies MongoDB Cloud user&#39;s geographic location. This parameter uses the ISO 3166-1a2 code format. | 
**CreatedAt** | Pointer to **time.Time** | Date and time when the current account is created. This value is in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**EmailAddress** | **string** | Email address that belongs to the MongoDB Cloud user. | 
**FirstName** | **string** | First or given name that belongs to the MongoDB Cloud user. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user. | [optional] [readonly] 
**LastAuth** | Pointer to **time.Time** | Date and time when the current account last authenticated. This value is in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**LastName** | **string** | Last name, family name, or surname that belongs to the MongoDB Cloud user. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MobileNumber** | **string** | Mobile phone number that belongs to the MongoDB Cloud user. | 
**Password** | **string** | Password applied with the username to log in to MongoDB Cloud. MongoDB Cloud does not return this parameter except in response to creating a new MongoDB Cloud user. Only the MongoDB Cloud user can update their password after it has been set from the MongoDB Cloud console. | 
**Roles** | Pointer to [**[]CloudAccessRoleAssignment**](CloudAccessRoleAssignment.md) | List of objects that display the MongoDB Cloud user&#39;s roles and the corresponding organization or project to which that role applies. A role can apply to one organization or one project but not both. | [optional] 
**TeamIds** | Pointer to **[]string** | List of unique 24-hexadecimal digit strings that identifies the teams to which this MongoDB Cloud user belongs. | [optional] [readonly] 
**Username** | **string** | Email address that represents the username of the MongoDB Cloud user. | 

## Methods

### NewCloudAppUser

`func NewCloudAppUser(country string, emailAddress string, firstName string, lastName string, mobileNumber string, password string, username string, ) *CloudAppUser`

NewCloudAppUser instantiates a new CloudAppUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAppUserWithDefaults

`func NewCloudAppUserWithDefaults() *CloudAppUser`

NewCloudAppUserWithDefaults instantiates a new CloudAppUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CloudAppUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CloudAppUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CloudAppUser) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCreatedAt

`func (o *CloudAppUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAppUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAppUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAppUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmailAddress

`func (o *CloudAppUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CloudAppUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CloudAppUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFirstName

`func (o *CloudAppUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CloudAppUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CloudAppUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetId

`func (o *CloudAppUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAppUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAppUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAppUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAuth

`func (o *CloudAppUser) GetLastAuth() time.Time`

GetLastAuth returns the LastAuth field if non-nil, zero value otherwise.

### GetLastAuthOk

`func (o *CloudAppUser) GetLastAuthOk() (*time.Time, bool)`

GetLastAuthOk returns a tuple with the LastAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuth

`func (o *CloudAppUser) SetLastAuth(v time.Time)`

SetLastAuth sets LastAuth field to given value.

### HasLastAuth

`func (o *CloudAppUser) HasLastAuth() bool`

HasLastAuth returns a boolean if a field has been set.

### GetLastName

`func (o *CloudAppUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CloudAppUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CloudAppUser) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetLinks

`func (o *CloudAppUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CloudAppUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CloudAppUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CloudAppUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMobileNumber

`func (o *CloudAppUser) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *CloudAppUser) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *CloudAppUser) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.


### GetPassword

`func (o *CloudAppUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CloudAppUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CloudAppUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRoles

`func (o *CloudAppUser) GetRoles() []CloudAccessRoleAssignment`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CloudAppUser) GetRolesOk() (*[]CloudAccessRoleAssignment, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CloudAppUser) SetRoles(v []CloudAccessRoleAssignment)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CloudAppUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTeamIds

`func (o *CloudAppUser) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *CloudAppUser) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *CloudAppUser) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *CloudAppUser) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUsername

`func (o *CloudAppUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudAppUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudAppUser) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


