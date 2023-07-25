# AtlasOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization. | [optional] [readonly] 
**IsDeleted** | Pointer to **bool** | Flag that indicates whether this organization has been deleted. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies the organization. | 

## Methods

### NewAtlasOrganization

`func NewAtlasOrganization(name string, ) *AtlasOrganization`

NewAtlasOrganization instantiates a new AtlasOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasOrganizationWithDefaults

`func NewAtlasOrganizationWithDefaults() *AtlasOrganization`

NewAtlasOrganizationWithDefaults instantiates a new AtlasOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AtlasOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtlasOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtlasOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AtlasOrganization) HasId() bool`

HasId returns a boolean if a field has been set.
### GetIsDeleted

`func (o *AtlasOrganization) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AtlasOrganization) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AtlasOrganization) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AtlasOrganization) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.
### GetLinks

`func (o *AtlasOrganization) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AtlasOrganization) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AtlasOrganization) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AtlasOrganization) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetName

`func (o *AtlasOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtlasOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtlasOrganization) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


