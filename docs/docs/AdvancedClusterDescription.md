# AdvancedClusterDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptDataRisksAndForceReplicaSetReconfig** | Pointer to **time.Time** | If reconfiguration is necessary to regain a primary due to a regional outage, submit this field alongside your topology reconfiguration to request a new regional outage resistant topology. Forced reconfigurations during an outage of the majority of electable nodes carry a risk of data loss if replicated writes (even majority committed writes) have not been replicated to the new primary node. MongoDB Atlas docs contain more information. To proceed with an operation which carries that risk, set **acceptDataRisksAndForceReplicaSetReconfig** to the current date. | [optional] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses [Cloud Backups](https://docs.atlas.mongodb.com/backup/cloud-backup/overview/) for dedicated clusters and [Shared Cluster Backups](https://docs.atlas.mongodb.com/backup/shared-tier/overview/) for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use backups. | [optional] [default to false]
**BiConnector** | Pointer to [**BiConnector**](BiConnector.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Configuration of nodes that comprise the cluster. | [optional] 
**ConnectionStrings** | Pointer to [**ClusterConnectionStrings**](ClusterConnectionStrings.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this cluster. This parameter expresses its value in ISO 8601 format in UTC. | [optional] [readonly] 
**DiskSizeGB** | Pointer to **float64** | Storage capacity that the host&#39;s root volume possesses expressed in gigabytes. Increase this number to add capacity. MongoDB Cloud requires this parameter if you set **replicationSpecs**. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value.  The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier. | [optional] 
**DiskWarmingMode** | Pointer to **string** | Disk warming mode selection. | [optional] [default to "FULLY_WARMED"]
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. To enable customer key management for encryption at rest, the cluster **replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize** setting must be &#x60;M10&#x60; or higher and &#x60;\&quot;backupEnabled\&quot; : false&#x60; or omitted entirely. | [optional] 
**GlobalClusterSelfManagedSharding** | Pointer to **bool** | Set this field to configure the Sharding Management Mode when creating a new Global Cluster.  When set to false, the management mode is set to Atlas-Managed Sharding. This mode fully manages the sharding of your Global Cluster and is built to provide a seamless deployment experience.  When set to true, the management mode is set to Self-Managed Sharding. This mode leaves the management of shards in your hands and is built to provide an advanced and flexible deployment experience.  This setting cannot be changed once the cluster is deployed. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Global Cluster. If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Global Cluster, you may specify this parameter. The request deletes any existing zones in a Global Cluster that you exclude from the request. | [optional] [readonly] 
**Labels** | Pointer to [**[]ComponentLabel**](ComponentLabel.md) | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn&#39;t display your labels.  Cluster labels are deprecated and will be removed in a future release. We strongly recommend that you use [resource tags](https://dochub.mongodb.org/core/add-cluster-tag-atlas) instead. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBMajorVersion** | Pointer to **string** | Major MongoDB version of the cluster. MongoDB Cloud deploys the cluster with the latest stable release of the specified version. | [optional] [default to "7.0"]
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the cluster runs. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the advanced cluster. | [optional] 
**Paused** | Pointer to **bool** | Flag that indicates whether the cluster is paused. | [optional] 
**PitEnabled** | Pointer to **bool** | Flag that indicates whether the cluster uses continuous cloud backups. | [optional] 
**ReplicationSpecs** | Pointer to [**[]ReplicationSpec**](ReplicationSpec.md) | List of settings that configure your cluster regions. For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global sharded clusters and replica sets, this array has one object representing where your clusters nodes deploy. | [optional] 
**RootCertType** | Pointer to **string** | Root Certificate Authority that MongoDB Cloud cluster uses. MongoDB Cloud supports Internet Security Research Group. | [optional] [default to "ISRGROOTX1"]
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of this cluster. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. If value is &#x60;CONTINUOUS&#x60;, you must not specify **mongoDBMajorVersion**. | [optional] [default to "LTS"]

## Methods

### NewAdvancedClusterDescription

`func NewAdvancedClusterDescription() *AdvancedClusterDescription`

NewAdvancedClusterDescription instantiates a new AdvancedClusterDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedClusterDescriptionWithDefaults

`func NewAdvancedClusterDescriptionWithDefaults() *AdvancedClusterDescription`

NewAdvancedClusterDescriptionWithDefaults instantiates a new AdvancedClusterDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *AdvancedClusterDescription) GetAcceptDataRisksAndForceReplicaSetReconfig() time.Time`

GetAcceptDataRisksAndForceReplicaSetReconfig returns the AcceptDataRisksAndForceReplicaSetReconfig field if non-nil, zero value otherwise.

### GetAcceptDataRisksAndForceReplicaSetReconfigOk

`func (o *AdvancedClusterDescription) GetAcceptDataRisksAndForceReplicaSetReconfigOk() (*time.Time, bool)`

GetAcceptDataRisksAndForceReplicaSetReconfigOk returns a tuple with the AcceptDataRisksAndForceReplicaSetReconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *AdvancedClusterDescription) SetAcceptDataRisksAndForceReplicaSetReconfig(v time.Time)`

SetAcceptDataRisksAndForceReplicaSetReconfig sets AcceptDataRisksAndForceReplicaSetReconfig field to given value.

### HasAcceptDataRisksAndForceReplicaSetReconfig

`func (o *AdvancedClusterDescription) HasAcceptDataRisksAndForceReplicaSetReconfig() bool`

HasAcceptDataRisksAndForceReplicaSetReconfig returns a boolean if a field has been set.
### GetBackupEnabled

`func (o *AdvancedClusterDescription) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *AdvancedClusterDescription) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *AdvancedClusterDescription) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *AdvancedClusterDescription) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.
### GetBiConnector

`func (o *AdvancedClusterDescription) GetBiConnector() BiConnector`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *AdvancedClusterDescription) GetBiConnectorOk() (*BiConnector, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *AdvancedClusterDescription) SetBiConnector(v BiConnector)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *AdvancedClusterDescription) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.
### GetClusterType

`func (o *AdvancedClusterDescription) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *AdvancedClusterDescription) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *AdvancedClusterDescription) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *AdvancedClusterDescription) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.
### GetConnectionStrings

`func (o *AdvancedClusterDescription) GetConnectionStrings() ClusterConnectionStrings`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *AdvancedClusterDescription) GetConnectionStringsOk() (*ClusterConnectionStrings, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *AdvancedClusterDescription) SetConnectionStrings(v ClusterConnectionStrings)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *AdvancedClusterDescription) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.
### GetCreateDate

`func (o *AdvancedClusterDescription) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *AdvancedClusterDescription) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *AdvancedClusterDescription) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *AdvancedClusterDescription) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.
### GetDiskSizeGB

`func (o *AdvancedClusterDescription) GetDiskSizeGB() float64`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *AdvancedClusterDescription) GetDiskSizeGBOk() (*float64, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *AdvancedClusterDescription) SetDiskSizeGB(v float64)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *AdvancedClusterDescription) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.
### GetDiskWarmingMode

`func (o *AdvancedClusterDescription) GetDiskWarmingMode() string`

GetDiskWarmingMode returns the DiskWarmingMode field if non-nil, zero value otherwise.

### GetDiskWarmingModeOk

`func (o *AdvancedClusterDescription) GetDiskWarmingModeOk() (*string, bool)`

GetDiskWarmingModeOk returns a tuple with the DiskWarmingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWarmingMode

`func (o *AdvancedClusterDescription) SetDiskWarmingMode(v string)`

SetDiskWarmingMode sets DiskWarmingMode field to given value.

### HasDiskWarmingMode

`func (o *AdvancedClusterDescription) HasDiskWarmingMode() bool`

HasDiskWarmingMode returns a boolean if a field has been set.
### GetEncryptionAtRestProvider

`func (o *AdvancedClusterDescription) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *AdvancedClusterDescription) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *AdvancedClusterDescription) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *AdvancedClusterDescription) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.
### GetGlobalClusterSelfManagedSharding

`func (o *AdvancedClusterDescription) GetGlobalClusterSelfManagedSharding() bool`

GetGlobalClusterSelfManagedSharding returns the GlobalClusterSelfManagedSharding field if non-nil, zero value otherwise.

### GetGlobalClusterSelfManagedShardingOk

`func (o *AdvancedClusterDescription) GetGlobalClusterSelfManagedShardingOk() (*bool, bool)`

GetGlobalClusterSelfManagedShardingOk returns a tuple with the GlobalClusterSelfManagedSharding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalClusterSelfManagedSharding

`func (o *AdvancedClusterDescription) SetGlobalClusterSelfManagedSharding(v bool)`

SetGlobalClusterSelfManagedSharding sets GlobalClusterSelfManagedSharding field to given value.

### HasGlobalClusterSelfManagedSharding

`func (o *AdvancedClusterDescription) HasGlobalClusterSelfManagedSharding() bool`

HasGlobalClusterSelfManagedSharding returns a boolean if a field has been set.
### GetGroupId

`func (o *AdvancedClusterDescription) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AdvancedClusterDescription) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AdvancedClusterDescription) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AdvancedClusterDescription) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *AdvancedClusterDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedClusterDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedClusterDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedClusterDescription) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLabels

`func (o *AdvancedClusterDescription) GetLabels() []ComponentLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AdvancedClusterDescription) GetLabelsOk() (*[]ComponentLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AdvancedClusterDescription) SetLabels(v []ComponentLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AdvancedClusterDescription) HasLabels() bool`

HasLabels returns a boolean if a field has been set.
### GetLinks

`func (o *AdvancedClusterDescription) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AdvancedClusterDescription) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AdvancedClusterDescription) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AdvancedClusterDescription) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMongoDBMajorVersion

`func (o *AdvancedClusterDescription) GetMongoDBMajorVersion() string`

GetMongoDBMajorVersion returns the MongoDBMajorVersion field if non-nil, zero value otherwise.

### GetMongoDBMajorVersionOk

`func (o *AdvancedClusterDescription) GetMongoDBMajorVersionOk() (*string, bool)`

GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBMajorVersion

`func (o *AdvancedClusterDescription) SetMongoDBMajorVersion(v string)`

SetMongoDBMajorVersion sets MongoDBMajorVersion field to given value.

### HasMongoDBMajorVersion

`func (o *AdvancedClusterDescription) HasMongoDBMajorVersion() bool`

HasMongoDBMajorVersion returns a boolean if a field has been set.
### GetMongoDBVersion

`func (o *AdvancedClusterDescription) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *AdvancedClusterDescription) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *AdvancedClusterDescription) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *AdvancedClusterDescription) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.
### GetName

`func (o *AdvancedClusterDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedClusterDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedClusterDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvancedClusterDescription) HasName() bool`

HasName returns a boolean if a field has been set.
### GetPaused

`func (o *AdvancedClusterDescription) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *AdvancedClusterDescription) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *AdvancedClusterDescription) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *AdvancedClusterDescription) HasPaused() bool`

HasPaused returns a boolean if a field has been set.
### GetPitEnabled

`func (o *AdvancedClusterDescription) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *AdvancedClusterDescription) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *AdvancedClusterDescription) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *AdvancedClusterDescription) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.
### GetReplicationSpecs

`func (o *AdvancedClusterDescription) GetReplicationSpecs() []ReplicationSpec`

GetReplicationSpecs returns the ReplicationSpecs field if non-nil, zero value otherwise.

### GetReplicationSpecsOk

`func (o *AdvancedClusterDescription) GetReplicationSpecsOk() (*[]ReplicationSpec, bool)`

GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecs

`func (o *AdvancedClusterDescription) SetReplicationSpecs(v []ReplicationSpec)`

SetReplicationSpecs sets ReplicationSpecs field to given value.

### HasReplicationSpecs

`func (o *AdvancedClusterDescription) HasReplicationSpecs() bool`

HasReplicationSpecs returns a boolean if a field has been set.
### GetRootCertType

`func (o *AdvancedClusterDescription) GetRootCertType() string`

GetRootCertType returns the RootCertType field if non-nil, zero value otherwise.

### GetRootCertTypeOk

`func (o *AdvancedClusterDescription) GetRootCertTypeOk() (*string, bool)`

GetRootCertTypeOk returns a tuple with the RootCertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCertType

`func (o *AdvancedClusterDescription) SetRootCertType(v string)`

SetRootCertType sets RootCertType field to given value.

### HasRootCertType

`func (o *AdvancedClusterDescription) HasRootCertType() bool`

HasRootCertType returns a boolean if a field has been set.
### GetStateName

`func (o *AdvancedClusterDescription) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *AdvancedClusterDescription) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *AdvancedClusterDescription) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *AdvancedClusterDescription) HasStateName() bool`

HasStateName returns a boolean if a field has been set.
### GetTags

`func (o *AdvancedClusterDescription) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdvancedClusterDescription) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdvancedClusterDescription) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdvancedClusterDescription) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *AdvancedClusterDescription) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *AdvancedClusterDescription) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *AdvancedClusterDescription) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *AdvancedClusterDescription) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.
### GetVersionReleaseSystem

`func (o *AdvancedClusterDescription) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *AdvancedClusterDescription) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *AdvancedClusterDescription) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *AdvancedClusterDescription) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


