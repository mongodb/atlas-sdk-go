# EmployeeAccessGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | **time.Time** | Expiration date for the employee access grant. |
**GrantType** | **string** | Level of access to grant to MongoDB Employees. |
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewEmployeeAccessGrant

`func NewEmployeeAccessGrant(expirationTime time.Time, grantType string, ) *EmployeeAccessGrant`

NewEmployeeAccessGrant instantiates a new EmployeeAccessGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeAccessGrantWithDefaults

`func NewEmployeeAccessGrantWithDefaults() *EmployeeAccessGrant`

NewEmployeeAccessGrantWithDefaults instantiates a new EmployeeAccessGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *EmployeeAccessGrant) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *EmployeeAccessGrant) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *EmployeeAccessGrant) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### GetGrantType

`func (o *EmployeeAccessGrant) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *EmployeeAccessGrant) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *EmployeeAccessGrant) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### GetLinks

`func (o *EmployeeAccessGrant) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmployeeAccessGrant) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmployeeAccessGrant) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmployeeAccessGrant) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


