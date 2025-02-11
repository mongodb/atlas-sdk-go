# FlexClusterDescription20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSettings** | Pointer to [**FlexBackupSettings20241113**](FlexBackupSettings20241113.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Flex cluster topology. | [optional] [readonly] [default to "REPLICASET"]
**ConnectionStrings** | Pointer to [**FlexConnectionStrings20241113**](FlexConnectionStrings20241113.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this instance. This parameter expresses its value in ISO 8601 format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the instance. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the instance runs. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the instance. | [optional] [readonly] 
**ProviderSettings** | [**FlexProviderSettings20241113**](FlexProviderSettings20241113.md) |  | 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of this instance. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the instance. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. | [optional] [readonly] [default to "LTS"]

## Methods

### NewFlexClusterDescription20241113

`func NewFlexClusterDescription20241113(providerSettings FlexProviderSettings20241113, ) *FlexClusterDescription20241113`

NewFlexClusterDescription20241113 instantiates a new FlexClusterDescription20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexClusterDescription20241113WithDefaults

`func NewFlexClusterDescription20241113WithDefaults() *FlexClusterDescription20241113`

NewFlexClusterDescription20241113WithDefaults instantiates a new FlexClusterDescription20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupSettings

`func (o *FlexClusterDescription20241113) GetBackupSettings() FlexBackupSettings20241113`

GetBackupSettings returns the BackupSettings field if non-nil, zero value otherwise.

### GetBackupSettingsOk

`func (o *FlexClusterDescription20241113) GetBackupSettingsOk() (*FlexBackupSettings20241113, bool)`

GetBackupSettingsOk returns a tuple with the BackupSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSettings

`func (o *FlexClusterDescription20241113) SetBackupSettings(v FlexBackupSettings20241113)`

SetBackupSettings sets BackupSettings field to given value.

### HasBackupSettings

`func (o *FlexClusterDescription20241113) HasBackupSettings() bool`

HasBackupSettings returns a boolean if a field has been set.
### GetClusterType

`func (o *FlexClusterDescription20241113) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *FlexClusterDescription20241113) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *FlexClusterDescription20241113) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *FlexClusterDescription20241113) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.
### GetConnectionStrings

`func (o *FlexClusterDescription20241113) GetConnectionStrings() FlexConnectionStrings20241113`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *FlexClusterDescription20241113) GetConnectionStringsOk() (*FlexConnectionStrings20241113, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *FlexClusterDescription20241113) SetConnectionStrings(v FlexConnectionStrings20241113)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *FlexClusterDescription20241113) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.
### GetCreateDate

`func (o *FlexClusterDescription20241113) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *FlexClusterDescription20241113) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *FlexClusterDescription20241113) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *FlexClusterDescription20241113) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.
### GetGroupId

`func (o *FlexClusterDescription20241113) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FlexClusterDescription20241113) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FlexClusterDescription20241113) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *FlexClusterDescription20241113) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *FlexClusterDescription20241113) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlexClusterDescription20241113) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlexClusterDescription20241113) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlexClusterDescription20241113) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLinks

`func (o *FlexClusterDescription20241113) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlexClusterDescription20241113) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlexClusterDescription20241113) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlexClusterDescription20241113) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMongoDBVersion

`func (o *FlexClusterDescription20241113) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *FlexClusterDescription20241113) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *FlexClusterDescription20241113) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *FlexClusterDescription20241113) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.
### GetName

`func (o *FlexClusterDescription20241113) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlexClusterDescription20241113) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlexClusterDescription20241113) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlexClusterDescription20241113) HasName() bool`

HasName returns a boolean if a field has been set.
### GetProviderSettings

`func (o *FlexClusterDescription20241113) GetProviderSettings() FlexProviderSettings20241113`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *FlexClusterDescription20241113) GetProviderSettingsOk() (*FlexProviderSettings20241113, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *FlexClusterDescription20241113) SetProviderSettings(v FlexProviderSettings20241113)`

SetProviderSettings sets ProviderSettings field to given value.

### GetStateName

`func (o *FlexClusterDescription20241113) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *FlexClusterDescription20241113) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *FlexClusterDescription20241113) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *FlexClusterDescription20241113) HasStateName() bool`

HasStateName returns a boolean if a field has been set.
### GetTags

`func (o *FlexClusterDescription20241113) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlexClusterDescription20241113) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlexClusterDescription20241113) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlexClusterDescription20241113) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *FlexClusterDescription20241113) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *FlexClusterDescription20241113) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *FlexClusterDescription20241113) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *FlexClusterDescription20241113) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.
### GetVersionReleaseSystem

`func (o *FlexClusterDescription20241113) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *FlexClusterDescription20241113) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *FlexClusterDescription20241113) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *FlexClusterDescription20241113) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


