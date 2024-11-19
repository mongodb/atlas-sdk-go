# FlexBackupRestoreJob20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryType** | Pointer to **string** | Means by which this resource returns the snapshot to the requesting MongoDB Cloud user. | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | Date and time when the download link no longer works. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the restore job. | [optional] [readonly] 
**InstanceName** | Pointer to **string** | Human-readable label that identifies the source instance. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**ProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project from which the restore job originated. | [optional] [readonly] 
**RestoreFinishedDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud completed writing this snapshot. MongoDB Cloud changes the status of the restore job to &#x60;CLOSED&#x60;. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**RestoreScheduledDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud will restore this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**SnapshotFinishedDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud completed writing this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot to restore. | [optional] [readonly] 
**SnapshotUrl** | Pointer to **string** | Internet address from which you can download the compressed snapshot files. The resource returns this parameter when  &#x60;\&quot;deliveryType\&quot; : \&quot;DOWNLOAD\&quot;&#x60;. | [optional] [readonly] 
**Status** | Pointer to **string** | Phase of the restore workflow for this job at the time this resource made this request. | [optional] [readonly] 
**TargetDeploymentItemName** | Pointer to **string** | Human-readable label that identifies the instance or cluster on the target project to which you want to restore the snapshot. You can restore the snapshot to another flex or dedicated cluster tier. | [optional] [readonly] 
**TargetProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that contains the instance or cluster to which you want to restore the snapshot. | [optional] [readonly] 

## Methods

### NewFlexBackupRestoreJob20241113

`func NewFlexBackupRestoreJob20241113() *FlexBackupRestoreJob20241113`

NewFlexBackupRestoreJob20241113 instantiates a new FlexBackupRestoreJob20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexBackupRestoreJob20241113WithDefaults

`func NewFlexBackupRestoreJob20241113WithDefaults() *FlexBackupRestoreJob20241113`

NewFlexBackupRestoreJob20241113WithDefaults instantiates a new FlexBackupRestoreJob20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryType

`func (o *FlexBackupRestoreJob20241113) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *FlexBackupRestoreJob20241113) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *FlexBackupRestoreJob20241113) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *FlexBackupRestoreJob20241113) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.
### GetExpirationDate

`func (o *FlexBackupRestoreJob20241113) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FlexBackupRestoreJob20241113) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FlexBackupRestoreJob20241113) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FlexBackupRestoreJob20241113) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.
### GetId

`func (o *FlexBackupRestoreJob20241113) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlexBackupRestoreJob20241113) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlexBackupRestoreJob20241113) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlexBackupRestoreJob20241113) HasId() bool`

HasId returns a boolean if a field has been set.
### GetInstanceName

`func (o *FlexBackupRestoreJob20241113) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *FlexBackupRestoreJob20241113) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *FlexBackupRestoreJob20241113) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *FlexBackupRestoreJob20241113) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.
### GetLinks

`func (o *FlexBackupRestoreJob20241113) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlexBackupRestoreJob20241113) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlexBackupRestoreJob20241113) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlexBackupRestoreJob20241113) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetProjectId

`func (o *FlexBackupRestoreJob20241113) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FlexBackupRestoreJob20241113) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FlexBackupRestoreJob20241113) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FlexBackupRestoreJob20241113) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.
### GetRestoreFinishedDate

`func (o *FlexBackupRestoreJob20241113) GetRestoreFinishedDate() time.Time`

GetRestoreFinishedDate returns the RestoreFinishedDate field if non-nil, zero value otherwise.

### GetRestoreFinishedDateOk

`func (o *FlexBackupRestoreJob20241113) GetRestoreFinishedDateOk() (*time.Time, bool)`

GetRestoreFinishedDateOk returns a tuple with the RestoreFinishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFinishedDate

`func (o *FlexBackupRestoreJob20241113) SetRestoreFinishedDate(v time.Time)`

SetRestoreFinishedDate sets RestoreFinishedDate field to given value.

### HasRestoreFinishedDate

`func (o *FlexBackupRestoreJob20241113) HasRestoreFinishedDate() bool`

HasRestoreFinishedDate returns a boolean if a field has been set.
### GetRestoreScheduledDate

`func (o *FlexBackupRestoreJob20241113) GetRestoreScheduledDate() time.Time`

GetRestoreScheduledDate returns the RestoreScheduledDate field if non-nil, zero value otherwise.

### GetRestoreScheduledDateOk

`func (o *FlexBackupRestoreJob20241113) GetRestoreScheduledDateOk() (*time.Time, bool)`

GetRestoreScheduledDateOk returns a tuple with the RestoreScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreScheduledDate

`func (o *FlexBackupRestoreJob20241113) SetRestoreScheduledDate(v time.Time)`

SetRestoreScheduledDate sets RestoreScheduledDate field to given value.

### HasRestoreScheduledDate

`func (o *FlexBackupRestoreJob20241113) HasRestoreScheduledDate() bool`

HasRestoreScheduledDate returns a boolean if a field has been set.
### GetSnapshotFinishedDate

`func (o *FlexBackupRestoreJob20241113) GetSnapshotFinishedDate() time.Time`

GetSnapshotFinishedDate returns the SnapshotFinishedDate field if non-nil, zero value otherwise.

### GetSnapshotFinishedDateOk

`func (o *FlexBackupRestoreJob20241113) GetSnapshotFinishedDateOk() (*time.Time, bool)`

GetSnapshotFinishedDateOk returns a tuple with the SnapshotFinishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFinishedDate

`func (o *FlexBackupRestoreJob20241113) SetSnapshotFinishedDate(v time.Time)`

SetSnapshotFinishedDate sets SnapshotFinishedDate field to given value.

### HasSnapshotFinishedDate

`func (o *FlexBackupRestoreJob20241113) HasSnapshotFinishedDate() bool`

HasSnapshotFinishedDate returns a boolean if a field has been set.
### GetSnapshotId

`func (o *FlexBackupRestoreJob20241113) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *FlexBackupRestoreJob20241113) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *FlexBackupRestoreJob20241113) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *FlexBackupRestoreJob20241113) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.
### GetSnapshotUrl

`func (o *FlexBackupRestoreJob20241113) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *FlexBackupRestoreJob20241113) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *FlexBackupRestoreJob20241113) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.

### HasSnapshotUrl

`func (o *FlexBackupRestoreJob20241113) HasSnapshotUrl() bool`

HasSnapshotUrl returns a boolean if a field has been set.
### GetStatus

`func (o *FlexBackupRestoreJob20241113) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlexBackupRestoreJob20241113) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlexBackupRestoreJob20241113) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlexBackupRestoreJob20241113) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetTargetDeploymentItemName

`func (o *FlexBackupRestoreJob20241113) GetTargetDeploymentItemName() string`

GetTargetDeploymentItemName returns the TargetDeploymentItemName field if non-nil, zero value otherwise.

### GetTargetDeploymentItemNameOk

`func (o *FlexBackupRestoreJob20241113) GetTargetDeploymentItemNameOk() (*string, bool)`

GetTargetDeploymentItemNameOk returns a tuple with the TargetDeploymentItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeploymentItemName

`func (o *FlexBackupRestoreJob20241113) SetTargetDeploymentItemName(v string)`

SetTargetDeploymentItemName sets TargetDeploymentItemName field to given value.

### HasTargetDeploymentItemName

`func (o *FlexBackupRestoreJob20241113) HasTargetDeploymentItemName() bool`

HasTargetDeploymentItemName returns a boolean if a field has been set.
### GetTargetProjectId

`func (o *FlexBackupRestoreJob20241113) GetTargetProjectId() string`

GetTargetProjectId returns the TargetProjectId field if non-nil, zero value otherwise.

### GetTargetProjectIdOk

`func (o *FlexBackupRestoreJob20241113) GetTargetProjectIdOk() (*string, bool)`

GetTargetProjectIdOk returns a tuple with the TargetProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProjectId

`func (o *FlexBackupRestoreJob20241113) SetTargetProjectId(v string)`

SetTargetProjectId sets TargetProjectId field to given value.

### HasTargetProjectId

`func (o *FlexBackupRestoreJob20241113) HasTargetProjectId() bool`

HasTargetProjectId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


