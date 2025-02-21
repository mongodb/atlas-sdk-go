# OrgUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user. | [readonly] 
**OrgMembershipStatus** | **string** | String enum that indicates whether the MongoDB Cloud user has a pending invitation to join the organization or they are already active in the organization. | [readonly] 
**Roles** | [**OrgUserRolesResponse**](OrgUserRolesResponse.md) |  | 
**TeamIds** | Pointer to **[]string** | List of unique 24-hexadecimal digit strings that identifies the teams to which this MongoDB Cloud user belongs. | [optional] [readonly] 
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

### NewOrgUserResponse

`func NewOrgUserResponse(id string, orgMembershipStatus string, roles OrgUserRolesResponse, username string, ) *OrgUserResponse`

NewOrgUserResponse instantiates a new OrgUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUserResponseWithDefaults

`func NewOrgUserResponseWithDefaults() *OrgUserResponse`

NewOrgUserResponseWithDefaults instantiates a new OrgUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrgUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgUserResponse) SetId(v string)`

SetId sets Id field to given value.

### GetOrgMembershipStatus

`func (o *OrgUserResponse) GetOrgMembershipStatus() string`

GetOrgMembershipStatus returns the OrgMembershipStatus field if non-nil, zero value otherwise.

### GetOrgMembershipStatusOk

`func (o *OrgUserResponse) GetOrgMembershipStatusOk() (*string, bool)`

GetOrgMembershipStatusOk returns a tuple with the OrgMembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipStatus

`func (o *OrgUserResponse) SetOrgMembershipStatus(v string)`

SetOrgMembershipStatus sets OrgMembershipStatus field to given value.

### GetRoles

`func (o *OrgUserResponse) GetRoles() OrgUserRolesResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgUserResponse) GetRolesOk() (*OrgUserRolesResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgUserResponse) SetRoles(v OrgUserRolesResponse)`

SetRoles sets Roles field to given value.

### GetTeamIds

`func (o *OrgUserResponse) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrgUserResponse) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrgUserResponse) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrgUserResponse) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.
### GetUsername

`func (o *OrgUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OrgUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OrgUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### GetInvitationCreatedAt

`func (o *OrgUserResponse) GetInvitationCreatedAt() time.Time`

GetInvitationCreatedAt returns the InvitationCreatedAt field if non-nil, zero value otherwise.

### GetInvitationCreatedAtOk

`func (o *OrgUserResponse) GetInvitationCreatedAtOk() (*time.Time, bool)`

GetInvitationCreatedAtOk returns a tuple with the InvitationCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationCreatedAt

`func (o *OrgUserResponse) SetInvitationCreatedAt(v time.Time)`

SetInvitationCreatedAt sets InvitationCreatedAt field to given value.

### HasInvitationCreatedAt

`func (o *OrgUserResponse) HasInvitationCreatedAt() bool`

HasInvitationCreatedAt returns a boolean if a field has been set.
### GetInvitationExpiresAt

`func (o *OrgUserResponse) GetInvitationExpiresAt() time.Time`

GetInvitationExpiresAt returns the InvitationExpiresAt field if non-nil, zero value otherwise.

### GetInvitationExpiresAtOk

`func (o *OrgUserResponse) GetInvitationExpiresAtOk() (*time.Time, bool)`

GetInvitationExpiresAtOk returns a tuple with the InvitationExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationExpiresAt

`func (o *OrgUserResponse) SetInvitationExpiresAt(v time.Time)`

SetInvitationExpiresAt sets InvitationExpiresAt field to given value.

### HasInvitationExpiresAt

`func (o *OrgUserResponse) HasInvitationExpiresAt() bool`

HasInvitationExpiresAt returns a boolean if a field has been set.
### GetInviterUsername

`func (o *OrgUserResponse) GetInviterUsername() string`

GetInviterUsername returns the InviterUsername field if non-nil, zero value otherwise.

### GetInviterUsernameOk

`func (o *OrgUserResponse) GetInviterUsernameOk() (*string, bool)`

GetInviterUsernameOk returns a tuple with the InviterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterUsername

`func (o *OrgUserResponse) SetInviterUsername(v string)`

SetInviterUsername sets InviterUsername field to given value.

### HasInviterUsername

`func (o *OrgUserResponse) HasInviterUsername() bool`

HasInviterUsername returns a boolean if a field has been set.
### GetCountry

`func (o *OrgUserResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrgUserResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrgUserResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrgUserResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.
### GetCreatedAt

`func (o *OrgUserResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgUserResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgUserResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrgUserResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetFirstName

`func (o *OrgUserResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrgUserResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrgUserResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrgUserResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.
### GetLastAuth

`func (o *OrgUserResponse) GetLastAuth() time.Time`

GetLastAuth returns the LastAuth field if non-nil, zero value otherwise.

### GetLastAuthOk

`func (o *OrgUserResponse) GetLastAuthOk() (*time.Time, bool)`

GetLastAuthOk returns a tuple with the LastAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuth

`func (o *OrgUserResponse) SetLastAuth(v time.Time)`

SetLastAuth sets LastAuth field to given value.

### HasLastAuth

`func (o *OrgUserResponse) HasLastAuth() bool`

HasLastAuth returns a boolean if a field has been set.
### GetLastName

`func (o *OrgUserResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrgUserResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrgUserResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrgUserResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.
### GetMobileNumber

`func (o *OrgUserResponse) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *OrgUserResponse) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *OrgUserResponse) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *OrgUserResponse) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


