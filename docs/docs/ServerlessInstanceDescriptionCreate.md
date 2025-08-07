# ServerlessInstanceDescriptionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable label that identifies the serverless instance. | 
**ProviderSettings** | [**ServerlessProviderSettings**](ServerlessProviderSettings.md) |  | 
**ServerlessBackupOptions** | Pointer to [**ClusterServerlessBackupOptions**](ClusterServerlessBackupOptions.md) |  | [optional] 
**StateName** | Pointer to **string** | Human-readable label that indicates any current activity being taken on this cluster by the Atlas control plane. With the exception of CREATING and DELETING states, clusters should always be available and have a Primary node even when in states indicating ongoing activity.   - &#x60;IDLE&#x60;: Atlas is making no changes to this cluster and all changes requested via the UI or API can be assumed to have been applied.  - &#x60;CREATING&#x60;: A cluster being provisioned for the very first time returns state CREATING until it is ready for connections. Ensure IP Access List and DB Users are configured before attempting to connect.  - &#x60;UPDATING&#x60;: A change requested via the UI, API, AutoScaling, or other scheduled activity is taking place.  - &#x60;DELETING&#x60;: The cluster is in the process of deletion and will soon be deleted.  - &#x60;REPAIRING&#x60;: One or more nodes in the cluster are being returned to service by the Atlas control plane. Other nodes should continue to provide service as normal. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the serverless instance. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the serverless instance. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the serverless instance. If set to &#x60;false&#x60;, MongoDB Cloud will delete the serverless instance. | [optional] [default to false]

## Methods

### NewServerlessInstanceDescriptionCreate

`func NewServerlessInstanceDescriptionCreate(name string, providerSettings ServerlessProviderSettings, ) *ServerlessInstanceDescriptionCreate`

NewServerlessInstanceDescriptionCreate instantiates a new ServerlessInstanceDescriptionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstanceDescriptionCreateWithDefaults

`func NewServerlessInstanceDescriptionCreateWithDefaults() *ServerlessInstanceDescriptionCreate`

NewServerlessInstanceDescriptionCreateWithDefaults instantiates a new ServerlessInstanceDescriptionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerlessInstanceDescriptionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerlessInstanceDescriptionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerlessInstanceDescriptionCreate) SetName(v string)`

SetName sets Name field to given value.

### GetProviderSettings

`func (o *ServerlessInstanceDescriptionCreate) GetProviderSettings() ServerlessProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ServerlessInstanceDescriptionCreate) GetProviderSettingsOk() (*ServerlessProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ServerlessInstanceDescriptionCreate) SetProviderSettings(v ServerlessProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### GetServerlessBackupOptions

`func (o *ServerlessInstanceDescriptionCreate) GetServerlessBackupOptions() ClusterServerlessBackupOptions`

GetServerlessBackupOptions returns the ServerlessBackupOptions field if non-nil, zero value otherwise.

### GetServerlessBackupOptionsOk

`func (o *ServerlessInstanceDescriptionCreate) GetServerlessBackupOptionsOk() (*ClusterServerlessBackupOptions, bool)`

GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessBackupOptions

`func (o *ServerlessInstanceDescriptionCreate) SetServerlessBackupOptions(v ClusterServerlessBackupOptions)`

SetServerlessBackupOptions sets ServerlessBackupOptions field to given value.

### HasServerlessBackupOptions

`func (o *ServerlessInstanceDescriptionCreate) HasServerlessBackupOptions() bool`

HasServerlessBackupOptions returns a boolean if a field has been set.
### GetStateName

`func (o *ServerlessInstanceDescriptionCreate) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ServerlessInstanceDescriptionCreate) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ServerlessInstanceDescriptionCreate) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ServerlessInstanceDescriptionCreate) HasStateName() bool`

HasStateName returns a boolean if a field has been set.
### GetTags

`func (o *ServerlessInstanceDescriptionCreate) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerlessInstanceDescriptionCreate) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerlessInstanceDescriptionCreate) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerlessInstanceDescriptionCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *ServerlessInstanceDescriptionCreate) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ServerlessInstanceDescriptionCreate) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ServerlessInstanceDescriptionCreate) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ServerlessInstanceDescriptionCreate) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


