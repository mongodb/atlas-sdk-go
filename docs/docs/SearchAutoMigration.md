# SearchAutoMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Flag that disables future automatic Search Node provisioning. Defaults to false on creation. Omission on update preserves the stored choice. Disabling retains existing Search Nodes and allows in-flight migrations to continue. Enabling does not bypass resource policies; if indexes exceed the migration threshold, the next eligible evaluation may start provisioning. | [optional] [default to false]

## Methods

### NewSearchAutoMigration

`func NewSearchAutoMigration() *SearchAutoMigration`

NewSearchAutoMigration instantiates a new SearchAutoMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAutoMigrationWithDefaults

`func NewSearchAutoMigrationWithDefaults() *SearchAutoMigration`

NewSearchAutoMigrationWithDefaults instantiates a new SearchAutoMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *SearchAutoMigration) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SearchAutoMigration) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SearchAutoMigration) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SearchAutoMigration) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *SearchAutoMigration) SetDisabledNil()`

SetDisabledNil sets Disabled to an explicit JSON null when marshaled, overriding any value previously set with SetDisabled. Calling SetDisabled again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


