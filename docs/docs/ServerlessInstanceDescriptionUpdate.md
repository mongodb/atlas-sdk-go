# ServerlessInstanceDescriptionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerlessBackupOptions** | Pointer to [**ServerlessBackupOptions**](ServerlessBackupOptions.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the serverless instance. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the serverless instance. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the serverless instance. If set to &#x60;false&#x60;, MongoDB Cloud will delete the serverless instance. | [optional] [default to false]

## Methods

### NewServerlessInstanceDescriptionUpdate

`func NewServerlessInstanceDescriptionUpdate() *ServerlessInstanceDescriptionUpdate`

NewServerlessInstanceDescriptionUpdate instantiates a new ServerlessInstanceDescriptionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstanceDescriptionUpdateWithDefaults

`func NewServerlessInstanceDescriptionUpdateWithDefaults() *ServerlessInstanceDescriptionUpdate`

NewServerlessInstanceDescriptionUpdateWithDefaults instantiates a new ServerlessInstanceDescriptionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerlessBackupOptions

`func (o *ServerlessInstanceDescriptionUpdate) GetServerlessBackupOptions() ServerlessBackupOptions`

GetServerlessBackupOptions returns the ServerlessBackupOptions field if non-nil, zero value otherwise.

### GetServerlessBackupOptionsOk

`func (o *ServerlessInstanceDescriptionUpdate) GetServerlessBackupOptionsOk() (*ServerlessBackupOptions, bool)`

GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessBackupOptions

`func (o *ServerlessInstanceDescriptionUpdate) SetServerlessBackupOptions(v ServerlessBackupOptions)`

SetServerlessBackupOptions sets ServerlessBackupOptions field to given value.

### HasServerlessBackupOptions

`func (o *ServerlessInstanceDescriptionUpdate) HasServerlessBackupOptions() bool`

HasServerlessBackupOptions returns a boolean if a field has been set.

### GetTags

`func (o *ServerlessInstanceDescriptionUpdate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerlessInstanceDescriptionUpdate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerlessInstanceDescriptionUpdate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerlessInstanceDescriptionUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationProtectionEnabled

`func (o *ServerlessInstanceDescriptionUpdate) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ServerlessInstanceDescriptionUpdate) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ServerlessInstanceDescriptionUpdate) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ServerlessInstanceDescriptionUpdate) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


