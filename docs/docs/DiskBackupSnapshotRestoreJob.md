# DiskBackupSnapshotRestoreJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to **bool** | Flag that indicates whether someone canceled this restore job. | [optional] [readonly] 
**Components** | Pointer to [**[]DiskBackupRestoreMember**](DiskBackupRestoreMember.md) | Information on the restore job for each replica set in the sharded cluster. | [optional] [readonly] 
**DeliveryType** | **string** | Human-readable label that categorizes the restore job to create. | 
**DeliveryUrl** | Pointer to **[]string** | One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;download\&quot;&#x60;. | [optional] [readonly] 
**DesiredTimestamp** | Pointer to [**ApiBSONTimestamp**](ApiBSONTimestamp.md) |  | [optional] 
**Expired** | Pointer to **bool** | Flag that indicates whether the restore job expired. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when the restore job expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Failed** | Pointer to **bool** | Flag that indicates whether the restore job failed. | [optional] [readonly] 
**FinishedAt** | Pointer to **time.Time** | Date and time when the restore job completed. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the restore job. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OplogInc** | Pointer to **int** | Oplog operation number from which you want to restore this snapshot. This number represents the second part of an Oplog timestamp. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;pointInTime\&quot;&#x60; and &#x60;oplogTs&#x60; exceeds &#x60;0&#x60;. | [optional] 
**OplogTs** | Pointer to **int** | Date and time from which you want to restore this snapshot. This parameter expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. This number represents the first part of an Oplog timestamp. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;pointInTime\&quot;&#x60; and &#x60;oplogTs&#x60; exceeds &#x60;0&#x60;. | [optional] 
**PointInTimeUTCSeconds** | Pointer to **int** | Date and time from which MongoDB Cloud restored this snapshot. This parameter expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;pointInTime\&quot;&#x60; and &#x60;pointInTimeUTCSeconds&#x60; exceeds &#x60;0&#x60;. | [optional] 
**PrivateDownloadDeliveryUrls** | Pointer to [**[]ApiPrivateDownloadDeliveryUrl**](ApiPrivateDownloadDeliveryUrl.md) | One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download and the corresponding private endpoint(s). MongoDB Cloud returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;download\&quot;&#x60; and the download can be performed privately. | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the snapshot. | [optional] 
**TargetClusterName** | Pointer to **string** | Human-readable label that identifies the target cluster to which the restore job restores the snapshot. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot;:&#x60; &#x60;\&quot;automated\&quot;&#x60;. Required for &#x60;automated&#x60; and &#x60;pointInTime&#x60; restore types. | [optional] 
**TargetGroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the target project for the specified &#x60;targetClusterName&#x60;. Required for &#x60;automated&#x60; and &#x60;pointInTime&#x60; restore types. | [optional] 
**Timestamp** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot associated with &#x60;snapshotId&#x60;. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewDiskBackupSnapshotRestoreJob

`func NewDiskBackupSnapshotRestoreJob(deliveryType string, ) *DiskBackupSnapshotRestoreJob`

NewDiskBackupSnapshotRestoreJob instantiates a new DiskBackupSnapshotRestoreJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotRestoreJobWithDefaults

`func NewDiskBackupSnapshotRestoreJobWithDefaults() *DiskBackupSnapshotRestoreJob`

NewDiskBackupSnapshotRestoreJobWithDefaults instantiates a new DiskBackupSnapshotRestoreJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *DiskBackupSnapshotRestoreJob) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *DiskBackupSnapshotRestoreJob) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *DiskBackupSnapshotRestoreJob) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *DiskBackupSnapshotRestoreJob) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.
### GetComponents

`func (o *DiskBackupSnapshotRestoreJob) GetComponents() []DiskBackupRestoreMember`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DiskBackupSnapshotRestoreJob) GetComponentsOk() (*[]DiskBackupRestoreMember, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DiskBackupSnapshotRestoreJob) SetComponents(v []DiskBackupRestoreMember)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DiskBackupSnapshotRestoreJob) HasComponents() bool`

HasComponents returns a boolean if a field has been set.
### GetDeliveryType

`func (o *DiskBackupSnapshotRestoreJob) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *DiskBackupSnapshotRestoreJob) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *DiskBackupSnapshotRestoreJob) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### GetDeliveryUrl

`func (o *DiskBackupSnapshotRestoreJob) GetDeliveryUrl() []string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *DiskBackupSnapshotRestoreJob) GetDeliveryUrlOk() (*[]string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *DiskBackupSnapshotRestoreJob) SetDeliveryUrl(v []string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *DiskBackupSnapshotRestoreJob) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.
### GetDesiredTimestamp

`func (o *DiskBackupSnapshotRestoreJob) GetDesiredTimestamp() ApiBSONTimestamp`

GetDesiredTimestamp returns the DesiredTimestamp field if non-nil, zero value otherwise.

### GetDesiredTimestampOk

`func (o *DiskBackupSnapshotRestoreJob) GetDesiredTimestampOk() (*ApiBSONTimestamp, bool)`

GetDesiredTimestampOk returns a tuple with the DesiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTimestamp

`func (o *DiskBackupSnapshotRestoreJob) SetDesiredTimestamp(v ApiBSONTimestamp)`

SetDesiredTimestamp sets DesiredTimestamp field to given value.

### HasDesiredTimestamp

`func (o *DiskBackupSnapshotRestoreJob) HasDesiredTimestamp() bool`

HasDesiredTimestamp returns a boolean if a field has been set.
### GetExpired

`func (o *DiskBackupSnapshotRestoreJob) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *DiskBackupSnapshotRestoreJob) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *DiskBackupSnapshotRestoreJob) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *DiskBackupSnapshotRestoreJob) HasExpired() bool`

HasExpired returns a boolean if a field has been set.
### GetExpiresAt

`func (o *DiskBackupSnapshotRestoreJob) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DiskBackupSnapshotRestoreJob) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DiskBackupSnapshotRestoreJob) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DiskBackupSnapshotRestoreJob) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.
### GetFailed

`func (o *DiskBackupSnapshotRestoreJob) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *DiskBackupSnapshotRestoreJob) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *DiskBackupSnapshotRestoreJob) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *DiskBackupSnapshotRestoreJob) HasFailed() bool`

HasFailed returns a boolean if a field has been set.
### GetFinishedAt

`func (o *DiskBackupSnapshotRestoreJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DiskBackupSnapshotRestoreJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DiskBackupSnapshotRestoreJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DiskBackupSnapshotRestoreJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.
### GetId

`func (o *DiskBackupSnapshotRestoreJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupSnapshotRestoreJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupSnapshotRestoreJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskBackupSnapshotRestoreJob) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLinks

`func (o *DiskBackupSnapshotRestoreJob) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotRestoreJob) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotRestoreJob) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotRestoreJob) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOplogInc

`func (o *DiskBackupSnapshotRestoreJob) GetOplogInc() int`

GetOplogInc returns the OplogInc field if non-nil, zero value otherwise.

### GetOplogIncOk

`func (o *DiskBackupSnapshotRestoreJob) GetOplogIncOk() (*int, bool)`

GetOplogIncOk returns a tuple with the OplogInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogInc

`func (o *DiskBackupSnapshotRestoreJob) SetOplogInc(v int)`

SetOplogInc sets OplogInc field to given value.

### HasOplogInc

`func (o *DiskBackupSnapshotRestoreJob) HasOplogInc() bool`

HasOplogInc returns a boolean if a field has been set.
### GetOplogTs

`func (o *DiskBackupSnapshotRestoreJob) GetOplogTs() int`

GetOplogTs returns the OplogTs field if non-nil, zero value otherwise.

### GetOplogTsOk

`func (o *DiskBackupSnapshotRestoreJob) GetOplogTsOk() (*int, bool)`

GetOplogTsOk returns a tuple with the OplogTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogTs

`func (o *DiskBackupSnapshotRestoreJob) SetOplogTs(v int)`

SetOplogTs sets OplogTs field to given value.

### HasOplogTs

`func (o *DiskBackupSnapshotRestoreJob) HasOplogTs() bool`

HasOplogTs returns a boolean if a field has been set.
### GetPointInTimeUTCSeconds

`func (o *DiskBackupSnapshotRestoreJob) GetPointInTimeUTCSeconds() int`

GetPointInTimeUTCSeconds returns the PointInTimeUTCSeconds field if non-nil, zero value otherwise.

### GetPointInTimeUTCSecondsOk

`func (o *DiskBackupSnapshotRestoreJob) GetPointInTimeUTCSecondsOk() (*int, bool)`

GetPointInTimeUTCSecondsOk returns a tuple with the PointInTimeUTCSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUTCSeconds

`func (o *DiskBackupSnapshotRestoreJob) SetPointInTimeUTCSeconds(v int)`

SetPointInTimeUTCSeconds sets PointInTimeUTCSeconds field to given value.

### HasPointInTimeUTCSeconds

`func (o *DiskBackupSnapshotRestoreJob) HasPointInTimeUTCSeconds() bool`

HasPointInTimeUTCSeconds returns a boolean if a field has been set.
### GetPrivateDownloadDeliveryUrls

`func (o *DiskBackupSnapshotRestoreJob) GetPrivateDownloadDeliveryUrls() []ApiPrivateDownloadDeliveryUrl`

GetPrivateDownloadDeliveryUrls returns the PrivateDownloadDeliveryUrls field if non-nil, zero value otherwise.

### GetPrivateDownloadDeliveryUrlsOk

`func (o *DiskBackupSnapshotRestoreJob) GetPrivateDownloadDeliveryUrlsOk() (*[]ApiPrivateDownloadDeliveryUrl, bool)`

GetPrivateDownloadDeliveryUrlsOk returns a tuple with the PrivateDownloadDeliveryUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDownloadDeliveryUrls

`func (o *DiskBackupSnapshotRestoreJob) SetPrivateDownloadDeliveryUrls(v []ApiPrivateDownloadDeliveryUrl)`

SetPrivateDownloadDeliveryUrls sets PrivateDownloadDeliveryUrls field to given value.

### HasPrivateDownloadDeliveryUrls

`func (o *DiskBackupSnapshotRestoreJob) HasPrivateDownloadDeliveryUrls() bool`

HasPrivateDownloadDeliveryUrls returns a boolean if a field has been set.
### GetSnapshotId

`func (o *DiskBackupSnapshotRestoreJob) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *DiskBackupSnapshotRestoreJob) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *DiskBackupSnapshotRestoreJob) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *DiskBackupSnapshotRestoreJob) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.
### GetTargetClusterName

`func (o *DiskBackupSnapshotRestoreJob) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *DiskBackupSnapshotRestoreJob) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *DiskBackupSnapshotRestoreJob) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### HasTargetClusterName

`func (o *DiskBackupSnapshotRestoreJob) HasTargetClusterName() bool`

HasTargetClusterName returns a boolean if a field has been set.
### GetTargetGroupId

`func (o *DiskBackupSnapshotRestoreJob) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *DiskBackupSnapshotRestoreJob) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *DiskBackupSnapshotRestoreJob) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *DiskBackupSnapshotRestoreJob) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.
### GetTimestamp

`func (o *DiskBackupSnapshotRestoreJob) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiskBackupSnapshotRestoreJob) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiskBackupSnapshotRestoreJob) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiskBackupSnapshotRestoreJob) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


