# CloudCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertCount** | Pointer to **int** | Whole number that indicates the quantity of alerts open on the cluster. | [optional] [readonly] 
**AuthEnabled** | Pointer to **bool** | Flag that indicates whether authentication is required to access the nodes in this cluster. | [optional] [readonly] 
**Availability** | Pointer to **string** | Term that expresses how many nodes of the cluster can be accessed when MongoDB Cloud receives this request. This parameter returns &#x60;available&#x60; when all nodes are accessible, &#x60;warning&#x60; only when some nodes in the cluster can be accessed, &#x60;unavailable&#x60; when the cluster can&#39;t be accessed, or &#x60;dead&#x60; when the cluster has been deactivated. | [optional] [readonly] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] [readonly] 
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the cluster. Each &#x60;&#x60;clusterId&#x60;&#x60; is used only once across all MongoDB Cloud deployments. | [optional] [readonly] 
**DataSizeBytes** | Pointer to **int64** | Total size of the data stored on each node in the cluster. The resource expresses this value in bytes. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the cluster. | [optional] [readonly] 
**NodeCount** | Pointer to **int** | Whole number that indicates the quantity of nodes that comprise the cluster. | [optional] [readonly] 
**SslEnabled** | Pointer to **bool** | Flag that indicates whether TLS authentication is required to access the nodes in this cluster. | [optional] [readonly] 
**Type** | Pointer to **string** | Human-readable label that indicates the cluster type. | [optional] [readonly] 
**Versions** | Pointer to **[]string** | List that contains the versions of MongoDB that each node in the cluster runs. | [optional] [readonly] 

## Methods

### NewCloudCluster

`func NewCloudCluster() *CloudCluster`

NewCloudCluster instantiates a new CloudCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClusterWithDefaults

`func NewCloudClusterWithDefaults() *CloudCluster`

NewCloudClusterWithDefaults instantiates a new CloudCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertCount

`func (o *CloudCluster) GetAlertCount() int`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *CloudCluster) GetAlertCountOk() (*int, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *CloudCluster) SetAlertCount(v int)`

SetAlertCount sets AlertCount field to given value.

### HasAlertCount

`func (o *CloudCluster) HasAlertCount() bool`

HasAlertCount returns a boolean if a field has been set.
### GetAuthEnabled

`func (o *CloudCluster) GetAuthEnabled() bool`

GetAuthEnabled returns the AuthEnabled field if non-nil, zero value otherwise.

### GetAuthEnabledOk

`func (o *CloudCluster) GetAuthEnabledOk() (*bool, bool)`

GetAuthEnabledOk returns a tuple with the AuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEnabled

`func (o *CloudCluster) SetAuthEnabled(v bool)`

SetAuthEnabled sets AuthEnabled field to given value.

### HasAuthEnabled

`func (o *CloudCluster) HasAuthEnabled() bool`

HasAuthEnabled returns a boolean if a field has been set.
### GetAvailability

`func (o *CloudCluster) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CloudCluster) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CloudCluster) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *CloudCluster) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.
### GetBackupEnabled

`func (o *CloudCluster) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *CloudCluster) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *CloudCluster) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *CloudCluster) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.
### GetClusterId

`func (o *CloudCluster) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloudCluster) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloudCluster) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloudCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.
### GetDataSizeBytes

`func (o *CloudCluster) GetDataSizeBytes() int64`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *CloudCluster) GetDataSizeBytesOk() (*int64, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *CloudCluster) SetDataSizeBytes(v int64)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### HasDataSizeBytes

`func (o *CloudCluster) HasDataSizeBytes() bool`

HasDataSizeBytes returns a boolean if a field has been set.
### GetName

`func (o *CloudCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCluster) HasName() bool`

HasName returns a boolean if a field has been set.
### GetNodeCount

`func (o *CloudCluster) GetNodeCount() int`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CloudCluster) GetNodeCountOk() (*int, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CloudCluster) SetNodeCount(v int)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *CloudCluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.
### GetSslEnabled

`func (o *CloudCluster) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *CloudCluster) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *CloudCluster) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *CloudCluster) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.
### GetType

`func (o *CloudCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudCluster) HasType() bool`

HasType returns a boolean if a field has been set.
### GetVersions

`func (o *CloudCluster) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CloudCluster) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CloudCluster) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CloudCluster) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


