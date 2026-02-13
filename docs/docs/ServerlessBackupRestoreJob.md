# ServerlessBackupRestoreJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to **bool** | Flag that indicates whether someone canceled this restore job. | [optional] [readonly] 
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
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the snapshot. | [optional] 
**TargetClusterName** | **string** | Human-readable label that identifies the target cluster to which the restore job restores the snapshot. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot;:&#x60; &#x60;\&quot;automated\&quot;&#x60;. | 
**TargetGroupId** | **string** | Unique 24-hexadecimal digit string that identifies the target project for the specified &#x60;targetClusterName&#x60;. | 
**Timestamp** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot associated with &#x60;snapshotId&#x60;. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewServerlessBackupRestoreJob

`func NewServerlessBackupRestoreJob(deliveryType string, targetClusterName string, targetGroupId string, ) *ServerlessBackupRestoreJob`

NewServerlessBackupRestoreJob instantiates a new ServerlessBackupRestoreJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessBackupRestoreJobWithDefaults

`func NewServerlessBackupRestoreJobWithDefaults() *ServerlessBackupRestoreJob`

NewServerlessBackupRestoreJobWithDefaults instantiates a new ServerlessBackupRestoreJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *ServerlessBackupRestoreJob) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *ServerlessBackupRestoreJob) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *ServerlessBackupRestoreJob) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *ServerlessBackupRestoreJob) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.
### GetDeliveryType

`func (o *ServerlessBackupRestoreJob) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ServerlessBackupRestoreJob) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ServerlessBackupRestoreJob) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### GetDeliveryUrl

`func (o *ServerlessBackupRestoreJob) GetDeliveryUrl() []string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *ServerlessBackupRestoreJob) GetDeliveryUrlOk() (*[]string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *ServerlessBackupRestoreJob) SetDeliveryUrl(v []string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *ServerlessBackupRestoreJob) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.
### GetDesiredTimestamp

`func (o *ServerlessBackupRestoreJob) GetDesiredTimestamp() ApiBSONTimestamp`

GetDesiredTimestamp returns the DesiredTimestamp field if non-nil, zero value otherwise.

### GetDesiredTimestampOk

`func (o *ServerlessBackupRestoreJob) GetDesiredTimestampOk() (*ApiBSONTimestamp, bool)`

GetDesiredTimestampOk returns a tuple with the DesiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTimestamp

`func (o *ServerlessBackupRestoreJob) SetDesiredTimestamp(v ApiBSONTimestamp)`

SetDesiredTimestamp sets DesiredTimestamp field to given value.

### HasDesiredTimestamp

`func (o *ServerlessBackupRestoreJob) HasDesiredTimestamp() bool`

HasDesiredTimestamp returns a boolean if a field has been set.
### GetExpired

`func (o *ServerlessBackupRestoreJob) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ServerlessBackupRestoreJob) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ServerlessBackupRestoreJob) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ServerlessBackupRestoreJob) HasExpired() bool`

HasExpired returns a boolean if a field has been set.
### GetExpiresAt

`func (o *ServerlessBackupRestoreJob) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ServerlessBackupRestoreJob) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ServerlessBackupRestoreJob) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ServerlessBackupRestoreJob) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.
### GetFailed

`func (o *ServerlessBackupRestoreJob) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ServerlessBackupRestoreJob) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ServerlessBackupRestoreJob) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ServerlessBackupRestoreJob) HasFailed() bool`

HasFailed returns a boolean if a field has been set.
### GetFinishedAt

`func (o *ServerlessBackupRestoreJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ServerlessBackupRestoreJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ServerlessBackupRestoreJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ServerlessBackupRestoreJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.
### GetId

`func (o *ServerlessBackupRestoreJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerlessBackupRestoreJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerlessBackupRestoreJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerlessBackupRestoreJob) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLinks

`func (o *ServerlessBackupRestoreJob) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerlessBackupRestoreJob) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerlessBackupRestoreJob) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerlessBackupRestoreJob) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOplogInc

`func (o *ServerlessBackupRestoreJob) GetOplogInc() int`

GetOplogInc returns the OplogInc field if non-nil, zero value otherwise.

### GetOplogIncOk

`func (o *ServerlessBackupRestoreJob) GetOplogIncOk() (*int, bool)`

GetOplogIncOk returns a tuple with the OplogInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogInc

`func (o *ServerlessBackupRestoreJob) SetOplogInc(v int)`

SetOplogInc sets OplogInc field to given value.

### HasOplogInc

`func (o *ServerlessBackupRestoreJob) HasOplogInc() bool`

HasOplogInc returns a boolean if a field has been set.
### GetOplogTs

`func (o *ServerlessBackupRestoreJob) GetOplogTs() int`

GetOplogTs returns the OplogTs field if non-nil, zero value otherwise.

### GetOplogTsOk

`func (o *ServerlessBackupRestoreJob) GetOplogTsOk() (*int, bool)`

GetOplogTsOk returns a tuple with the OplogTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogTs

`func (o *ServerlessBackupRestoreJob) SetOplogTs(v int)`

SetOplogTs sets OplogTs field to given value.

### HasOplogTs

`func (o *ServerlessBackupRestoreJob) HasOplogTs() bool`

HasOplogTs returns a boolean if a field has been set.
### GetPointInTimeUTCSeconds

`func (o *ServerlessBackupRestoreJob) GetPointInTimeUTCSeconds() int`

GetPointInTimeUTCSeconds returns the PointInTimeUTCSeconds field if non-nil, zero value otherwise.

### GetPointInTimeUTCSecondsOk

`func (o *ServerlessBackupRestoreJob) GetPointInTimeUTCSecondsOk() (*int, bool)`

GetPointInTimeUTCSecondsOk returns a tuple with the PointInTimeUTCSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUTCSeconds

`func (o *ServerlessBackupRestoreJob) SetPointInTimeUTCSeconds(v int)`

SetPointInTimeUTCSeconds sets PointInTimeUTCSeconds field to given value.

### HasPointInTimeUTCSeconds

`func (o *ServerlessBackupRestoreJob) HasPointInTimeUTCSeconds() bool`

HasPointInTimeUTCSeconds returns a boolean if a field has been set.
### GetSnapshotId

`func (o *ServerlessBackupRestoreJob) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ServerlessBackupRestoreJob) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ServerlessBackupRestoreJob) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ServerlessBackupRestoreJob) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.
### GetTargetClusterName

`func (o *ServerlessBackupRestoreJob) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *ServerlessBackupRestoreJob) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *ServerlessBackupRestoreJob) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### GetTargetGroupId

`func (o *ServerlessBackupRestoreJob) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ServerlessBackupRestoreJob) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ServerlessBackupRestoreJob) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### GetTimestamp

`func (o *ServerlessBackupRestoreJob) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServerlessBackupRestoreJob) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServerlessBackupRestoreJob) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServerlessBackupRestoreJob) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


