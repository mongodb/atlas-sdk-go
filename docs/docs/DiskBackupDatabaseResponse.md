# DiskBackupDatabaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies the database within the snapshot. | [readonly] 

## Methods

### NewDiskBackupDatabaseResponse

`func NewDiskBackupDatabaseResponse(name string, ) *DiskBackupDatabaseResponse`

NewDiskBackupDatabaseResponse instantiates a new DiskBackupDatabaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupDatabaseResponseWithDefaults

`func NewDiskBackupDatabaseResponseWithDefaults() *DiskBackupDatabaseResponse`

NewDiskBackupDatabaseResponseWithDefaults instantiates a new DiskBackupDatabaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DiskBackupDatabaseResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupDatabaseResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupDatabaseResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupDatabaseResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *DiskBackupDatabaseResponse) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetName

`func (o *DiskBackupDatabaseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskBackupDatabaseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskBackupDatabaseResponse) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


