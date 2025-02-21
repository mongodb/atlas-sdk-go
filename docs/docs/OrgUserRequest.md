# OrgUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | [**OrgUserRolesRequest**](OrgUserRolesRequest.md) |  | 
**TeamIds** | Pointer to **[]string** | List of unique 24-hexadecimal digit strings that identifies the teams to which this MongoDB Cloud user belongs. | [optional] 
**Username** | **string** | Email address that represents the username of the MongoDB Cloud user. | 

## Methods

### NewOrgUserRequest

`func NewOrgUserRequest(roles OrgUserRolesRequest, username string, ) *OrgUserRequest`

NewOrgUserRequest instantiates a new OrgUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUserRequestWithDefaults

`func NewOrgUserRequestWithDefaults() *OrgUserRequest`

NewOrgUserRequestWithDefaults instantiates a new OrgUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *OrgUserRequest) GetRoles() OrgUserRolesRequest`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgUserRequest) GetRolesOk() (*OrgUserRolesRequest, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgUserRequest) SetRoles(v OrgUserRolesRequest)`

SetRoles sets Roles field to given value.

### GetTeamIds

`func (o *OrgUserRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrgUserRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrgUserRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrgUserRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.
### GetUsername

`func (o *OrgUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OrgUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OrgUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


