# ServerlessInstanceDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStrings** | Pointer to [**ServerlessInstanceDescriptionConnectionStrings**](ServerlessInstanceDescriptionConnectionStrings.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this serverless instance. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the serverless instance. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the serverless instance runs. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the serverless instance. | [optional] [readonly] 
**ProviderSettings** | [**ServerlessProviderSettings**](ServerlessProviderSettings.md) |  | 
**ServerlessBackupOptions** | Pointer to [**ClusterServerlessBackupOptions**](ClusterServerlessBackupOptions.md) |  | [optional] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of the serverless instance. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the serverless instance. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the serverless instance. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the serverless instance. If set to &#x60;false&#x60;, MongoDB Cloud will delete the serverless instance. | [optional] [default to false]

## Methods

### NewServerlessInstanceDescription

`func NewServerlessInstanceDescription(providerSettings ServerlessProviderSettings, ) *ServerlessInstanceDescription`

NewServerlessInstanceDescription instantiates a new ServerlessInstanceDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstanceDescriptionWithDefaults

`func NewServerlessInstanceDescriptionWithDefaults() *ServerlessInstanceDescription`

NewServerlessInstanceDescriptionWithDefaults instantiates a new ServerlessInstanceDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStrings

`func (o *ServerlessInstanceDescription) GetConnectionStrings() ServerlessInstanceDescriptionConnectionStrings`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *ServerlessInstanceDescription) GetConnectionStringsOk() (*ServerlessInstanceDescriptionConnectionStrings, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *ServerlessInstanceDescription) SetConnectionStrings(v ServerlessInstanceDescriptionConnectionStrings)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *ServerlessInstanceDescription) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetCreateDate

`func (o *ServerlessInstanceDescription) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ServerlessInstanceDescription) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ServerlessInstanceDescription) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ServerlessInstanceDescription) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ServerlessInstanceDescription) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerlessInstanceDescription) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerlessInstanceDescription) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ServerlessInstanceDescription) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ServerlessInstanceDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerlessInstanceDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerlessInstanceDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerlessInstanceDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ServerlessInstanceDescription) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerlessInstanceDescription) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerlessInstanceDescription) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerlessInstanceDescription) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMongoDBVersion

`func (o *ServerlessInstanceDescription) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ServerlessInstanceDescription) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ServerlessInstanceDescription) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ServerlessInstanceDescription) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetName

`func (o *ServerlessInstanceDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerlessInstanceDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerlessInstanceDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerlessInstanceDescription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ServerlessInstanceDescription) GetProviderSettings() ServerlessProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ServerlessInstanceDescription) GetProviderSettingsOk() (*ServerlessProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ServerlessInstanceDescription) SetProviderSettings(v ServerlessProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.


### GetServerlessBackupOptions

`func (o *ServerlessInstanceDescription) GetServerlessBackupOptions() ClusterServerlessBackupOptions`

GetServerlessBackupOptions returns the ServerlessBackupOptions field if non-nil, zero value otherwise.

### GetServerlessBackupOptionsOk

`func (o *ServerlessInstanceDescription) GetServerlessBackupOptionsOk() (*ClusterServerlessBackupOptions, bool)`

GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessBackupOptions

`func (o *ServerlessInstanceDescription) SetServerlessBackupOptions(v ClusterServerlessBackupOptions)`

SetServerlessBackupOptions sets ServerlessBackupOptions field to given value.

### HasServerlessBackupOptions

`func (o *ServerlessInstanceDescription) HasServerlessBackupOptions() bool`

HasServerlessBackupOptions returns a boolean if a field has been set.

### GetStateName

`func (o *ServerlessInstanceDescription) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ServerlessInstanceDescription) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ServerlessInstanceDescription) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ServerlessInstanceDescription) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### GetTags

`func (o *ServerlessInstanceDescription) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerlessInstanceDescription) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerlessInstanceDescription) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerlessInstanceDescription) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationProtectionEnabled

`func (o *ServerlessInstanceDescription) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ServerlessInstanceDescription) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ServerlessInstanceDescription) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ServerlessInstanceDescription) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


