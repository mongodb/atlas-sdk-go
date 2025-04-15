# DiskBackupRestoreMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadUrl** | Pointer to **string** | One Uniform Resource Locator that points to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;download\&quot;&#x60;. | [optional] [readonly] 
**PrivateDownloadDeliveryUrls** | Pointer to [**[]ApiPrivateDownloadDeliveryUrl**](ApiPrivateDownloadDeliveryUrl.md) | One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download and the corresponding private endpoint(s). MongoDB Cloud returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;download\&quot;&#x60; and the download can be performed privately. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label that identifies the replica set on the sharded cluster. | [optional] [readonly] 

## Methods

### NewDiskBackupRestoreMember

`func NewDiskBackupRestoreMember() *DiskBackupRestoreMember`

NewDiskBackupRestoreMember instantiates a new DiskBackupRestoreMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupRestoreMemberWithDefaults

`func NewDiskBackupRestoreMemberWithDefaults() *DiskBackupRestoreMember`

NewDiskBackupRestoreMemberWithDefaults instantiates a new DiskBackupRestoreMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadUrl

`func (o *DiskBackupRestoreMember) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *DiskBackupRestoreMember) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *DiskBackupRestoreMember) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *DiskBackupRestoreMember) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.
### GetPrivateDownloadDeliveryUrls

`func (o *DiskBackupRestoreMember) GetPrivateDownloadDeliveryUrls() []ApiPrivateDownloadDeliveryUrl`

GetPrivateDownloadDeliveryUrls returns the PrivateDownloadDeliveryUrls field if non-nil, zero value otherwise.

### GetPrivateDownloadDeliveryUrlsOk

`func (o *DiskBackupRestoreMember) GetPrivateDownloadDeliveryUrlsOk() (*[]ApiPrivateDownloadDeliveryUrl, bool)`

GetPrivateDownloadDeliveryUrlsOk returns a tuple with the PrivateDownloadDeliveryUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDownloadDeliveryUrls

`func (o *DiskBackupRestoreMember) SetPrivateDownloadDeliveryUrls(v []ApiPrivateDownloadDeliveryUrl)`

SetPrivateDownloadDeliveryUrls sets PrivateDownloadDeliveryUrls field to given value.

### HasPrivateDownloadDeliveryUrls

`func (o *DiskBackupRestoreMember) HasPrivateDownloadDeliveryUrls() bool`

HasPrivateDownloadDeliveryUrls returns a boolean if a field has been set.
### GetReplicaSetName

`func (o *DiskBackupRestoreMember) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *DiskBackupRestoreMember) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *DiskBackupRestoreMember) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *DiskBackupRestoreMember) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


