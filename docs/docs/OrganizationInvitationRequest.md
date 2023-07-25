# OrganizationInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** | One or more organization or project level roles to assign to the MongoDB Cloud user. | [optional] 
**TeamIds** | Pointer to **[]string** | List of teams to which you want to invite the desired MongoDB Cloud user. | [optional] 
**Username** | Pointer to **string** | Email address that belongs to the desired MongoDB Cloud user. | [optional] 

## Methods

### NewOrganizationInvitationRequest

`func NewOrganizationInvitationRequest() *OrganizationInvitationRequest`

NewOrganizationInvitationRequest instantiates a new OrganizationInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationRequestWithDefaults

`func NewOrganizationInvitationRequestWithDefaults() *OrganizationInvitationRequest`

NewOrganizationInvitationRequestWithDefaults instantiates a new OrganizationInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *OrganizationInvitationRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrganizationInvitationRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrganizationInvitationRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrganizationInvitationRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetTeamIds

`func (o *OrganizationInvitationRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrganizationInvitationRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrganizationInvitationRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrganizationInvitationRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.
### GetUsername

`func (o *OrganizationInvitationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OrganizationInvitationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OrganizationInvitationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OrganizationInvitationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


