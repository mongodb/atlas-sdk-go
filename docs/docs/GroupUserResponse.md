# GroupUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user. | [readonly] 
**OrgMembershipStatus** | **string** | String enum that indicates whether the MongoDB Cloud user has a pending invitation to join the organization or they are already active in the organization. | [readonly] 
**Roles** | **[]string** | One or more project-level roles assigned to the MongoDB Cloud user. | [readonly] 
**Username** | **string** | Email address that represents the username of the MongoDB Cloud user. | [readonly] 
**InvitationCreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud sent the invitation. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**InvitationExpiresAt** | Pointer to **time.Time** | Date and time when the invitation from MongoDB Cloud expires. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**InviterUsername** | Pointer to **string** | Username of the MongoDB Cloud user who sent the invitation to join the organization. | [optional] [readonly] 
**Country** | Pointer to **string** | Two-character alphabetical string that identifies the MongoDB Cloud user&#39;s geographic location. This parameter uses the ISO 3166-1a2 code format. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud created the current account. This value is in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FirstName** | Pointer to **string** | First or given name that belongs to the MongoDB Cloud user. | [optional] [readonly] 
**LastAuth** | Pointer to **time.Time** | Date and time when the current account last authenticated. This value is in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**LastName** | Pointer to **string** | Last name, family name, or surname that belongs to the MongoDB Cloud user. | [optional] [readonly] 
**MobileNumber** | Pointer to **string** | Mobile phone number that belongs to the MongoDB Cloud user. | [optional] [readonly] 

## Methods

### NewGroupUserResponse

`func NewGroupUserResponse(id string, orgMembershipStatus string, roles []string, username string, ) *GroupUserResponse`

NewGroupUserResponse instantiates a new GroupUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserResponseWithDefaults

`func NewGroupUserResponseWithDefaults() *GroupUserResponse`

NewGroupUserResponseWithDefaults instantiates a new GroupUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupUserResponse) SetId(v string)`

SetId sets Id field to given value.

### GetOrgMembershipStatus

`func (o *GroupUserResponse) GetOrgMembershipStatus() string`

GetOrgMembershipStatus returns the OrgMembershipStatus field if non-nil, zero value otherwise.

### GetOrgMembershipStatusOk

`func (o *GroupUserResponse) GetOrgMembershipStatusOk() (*string, bool)`

GetOrgMembershipStatusOk returns a tuple with the OrgMembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipStatus

`func (o *GroupUserResponse) SetOrgMembershipStatus(v string)`

SetOrgMembershipStatus sets OrgMembershipStatus field to given value.

### GetRoles

`func (o *GroupUserResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupUserResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupUserResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### GetUsername

`func (o *GroupUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GroupUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GroupUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### GetInvitationCreatedAt

`func (o *GroupUserResponse) GetInvitationCreatedAt() time.Time`

GetInvitationCreatedAt returns the InvitationCreatedAt field if non-nil, zero value otherwise.

### GetInvitationCreatedAtOk

`func (o *GroupUserResponse) GetInvitationCreatedAtOk() (*time.Time, bool)`

GetInvitationCreatedAtOk returns a tuple with the InvitationCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationCreatedAt

`func (o *GroupUserResponse) SetInvitationCreatedAt(v time.Time)`

SetInvitationCreatedAt sets InvitationCreatedAt field to given value.

### HasInvitationCreatedAt

`func (o *GroupUserResponse) HasInvitationCreatedAt() bool`

HasInvitationCreatedAt returns a boolean if a field has been set.
### GetInvitationExpiresAt

`func (o *GroupUserResponse) GetInvitationExpiresAt() time.Time`

GetInvitationExpiresAt returns the InvitationExpiresAt field if non-nil, zero value otherwise.

### GetInvitationExpiresAtOk

`func (o *GroupUserResponse) GetInvitationExpiresAtOk() (*time.Time, bool)`

GetInvitationExpiresAtOk returns a tuple with the InvitationExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationExpiresAt

`func (o *GroupUserResponse) SetInvitationExpiresAt(v time.Time)`

SetInvitationExpiresAt sets InvitationExpiresAt field to given value.

### HasInvitationExpiresAt

`func (o *GroupUserResponse) HasInvitationExpiresAt() bool`

HasInvitationExpiresAt returns a boolean if a field has been set.
### GetInviterUsername

`func (o *GroupUserResponse) GetInviterUsername() string`

GetInviterUsername returns the InviterUsername field if non-nil, zero value otherwise.

### GetInviterUsernameOk

`func (o *GroupUserResponse) GetInviterUsernameOk() (*string, bool)`

GetInviterUsernameOk returns a tuple with the InviterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterUsername

`func (o *GroupUserResponse) SetInviterUsername(v string)`

SetInviterUsername sets InviterUsername field to given value.

### HasInviterUsername

`func (o *GroupUserResponse) HasInviterUsername() bool`

HasInviterUsername returns a boolean if a field has been set.
### GetCountry

`func (o *GroupUserResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GroupUserResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GroupUserResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GroupUserResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.
### GetCreatedAt

`func (o *GroupUserResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupUserResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupUserResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupUserResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetFirstName

`func (o *GroupUserResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GroupUserResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GroupUserResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GroupUserResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.
### GetLastAuth

`func (o *GroupUserResponse) GetLastAuth() time.Time`

GetLastAuth returns the LastAuth field if non-nil, zero value otherwise.

### GetLastAuthOk

`func (o *GroupUserResponse) GetLastAuthOk() (*time.Time, bool)`

GetLastAuthOk returns a tuple with the LastAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuth

`func (o *GroupUserResponse) SetLastAuth(v time.Time)`

SetLastAuth sets LastAuth field to given value.

### HasLastAuth

`func (o *GroupUserResponse) HasLastAuth() bool`

HasLastAuth returns a boolean if a field has been set.
### GetLastName

`func (o *GroupUserResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GroupUserResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GroupUserResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GroupUserResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.
### GetMobileNumber

`func (o *GroupUserResponse) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *GroupUserResponse) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *GroupUserResponse) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *GroupUserResponse) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


