# ApiAtlasResourcePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedByUser** | Pointer to [**ApiAtlasUserMetadata**](ApiAtlasUserMetadata.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date and time in UTC when the atlas resource policy was created. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the atlas resource policy. | [optional] [readonly] 
**LastUpdatedByUser** | Pointer to [**ApiAtlasUserMetadata**](ApiAtlasUserMetadata.md) |  | [optional] 
**LastUpdatedDate** | Pointer to **time.Time** | Date and time in UTC when the atlas resource policy was last updated. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that describes the atlas resource policy. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the organization the atlas resource policy belongs to. | [optional] [readonly] 
**Policies** | Pointer to [**[]ApiAtlasPolicy**](ApiAtlasPolicy.md) | List of policies that make up the atlas resource policy. | [optional] [readonly] 
**Version** | Pointer to **string** | A string that identifies the version of the atlas resource policy. | [optional] [readonly] 

## Methods

### NewApiAtlasResourcePolicy

`func NewApiAtlasResourcePolicy() *ApiAtlasResourcePolicy`

NewApiAtlasResourcePolicy instantiates a new ApiAtlasResourcePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasResourcePolicyWithDefaults

`func NewApiAtlasResourcePolicyWithDefaults() *ApiAtlasResourcePolicy`

NewApiAtlasResourcePolicyWithDefaults instantiates a new ApiAtlasResourcePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedByUser

`func (o *ApiAtlasResourcePolicy) GetCreatedByUser() ApiAtlasUserMetadata`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *ApiAtlasResourcePolicy) GetCreatedByUserOk() (*ApiAtlasUserMetadata, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *ApiAtlasResourcePolicy) SetCreatedByUser(v ApiAtlasUserMetadata)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *ApiAtlasResourcePolicy) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.
### GetCreatedDate

`func (o *ApiAtlasResourcePolicy) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiAtlasResourcePolicy) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiAtlasResourcePolicy) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiAtlasResourcePolicy) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.
### GetId

`func (o *ApiAtlasResourcePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasResourcePolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasResourcePolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasResourcePolicy) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLastUpdatedByUser

`func (o *ApiAtlasResourcePolicy) GetLastUpdatedByUser() ApiAtlasUserMetadata`

GetLastUpdatedByUser returns the LastUpdatedByUser field if non-nil, zero value otherwise.

### GetLastUpdatedByUserOk

`func (o *ApiAtlasResourcePolicy) GetLastUpdatedByUserOk() (*ApiAtlasUserMetadata, bool)`

GetLastUpdatedByUserOk returns a tuple with the LastUpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedByUser

`func (o *ApiAtlasResourcePolicy) SetLastUpdatedByUser(v ApiAtlasUserMetadata)`

SetLastUpdatedByUser sets LastUpdatedByUser field to given value.

### HasLastUpdatedByUser

`func (o *ApiAtlasResourcePolicy) HasLastUpdatedByUser() bool`

HasLastUpdatedByUser returns a boolean if a field has been set.
### GetLastUpdatedDate

`func (o *ApiAtlasResourcePolicy) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ApiAtlasResourcePolicy) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ApiAtlasResourcePolicy) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ApiAtlasResourcePolicy) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.
### GetName

`func (o *ApiAtlasResourcePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasResourcePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasResourcePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasResourcePolicy) HasName() bool`

HasName returns a boolean if a field has been set.
### GetOrgId

`func (o *ApiAtlasResourcePolicy) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiAtlasResourcePolicy) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiAtlasResourcePolicy) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApiAtlasResourcePolicy) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetPolicies

`func (o *ApiAtlasResourcePolicy) GetPolicies() []ApiAtlasPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApiAtlasResourcePolicy) GetPoliciesOk() (*[]ApiAtlasPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApiAtlasResourcePolicy) SetPolicies(v []ApiAtlasPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ApiAtlasResourcePolicy) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.
### GetVersion

`func (o *ApiAtlasResourcePolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiAtlasResourcePolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiAtlasResourcePolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiAtlasResourcePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


