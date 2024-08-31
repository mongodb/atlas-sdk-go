# ClusterDescription20240805

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptDataRisksAndForceReplicaSetReconfig** | Pointer to **time.Time** | If reconfiguration is necessary to regain a primary due to a regional outage, submit this field alongside your topology reconfiguration to request a new regional outage resistant topology. Forced reconfigurations during an outage of the majority of electable nodes carry a risk of data loss if replicated writes (even majority committed writes) have not been replicated to the new primary node. MongoDB Atlas docs contain more information. To proceed with an operation which carries that risk, set **acceptDataRisksAndForceReplicaSetReconfig** to the current date. | [optional] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses [Cloud Backups](https://docs.atlas.mongodb.com/backup/cloud-backup/overview/) for dedicated clusters and [Shared Cluster Backups](https://docs.atlas.mongodb.com/backup/shared-tier/overview/) for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use backups. | [optional] [default to false]
**BiConnector** | Pointer to [**BiConnector**](BiConnector.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Configuration of nodes that comprise the cluster. | [optional] 
**ConnectionStrings** | Pointer to [**ClusterConnectionStrings**](ClusterConnectionStrings.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this cluster. This parameter expresses its value in ISO 8601 format in UTC. | [optional] [readonly] 
**DiskWarmingMode** | Pointer to **string** | Disk warming mode selection. | [optional] [default to "FULLY_WARMED"]
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. To enable customer key management for encryption at rest, the cluster **replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize** setting must be &#x60;M10&#x60; or higher and &#x60;\&quot;backupEnabled\&quot; : false&#x60; or omitted entirely. | [optional] 
**FeatureCompatibilityVersion** | Pointer to **string** | Feature compatibility version of the cluster. | [optional] [readonly] 
**FeatureCompatibilityVersionExpirationDate** | Pointer to **time.Time** | Feature compatibility version expiration date. | [optional] [readonly] 
**GlobalClusterSelfManagedSharding** | Pointer to **bool** | Set this field to configure the Sharding Management Mode when creating a new Global Cluster.  When set to false, the management mode is set to Atlas-Managed Sharding. This mode fully manages the sharding of your Global Cluster and is built to provide a seamless deployment experience.  When set to true, the management mode is set to Self-Managed Sharding. This mode leaves the management of shards in your hands and is built to provide an advanced and flexible deployment experience.  This setting cannot be changed once the cluster is deployed. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster. | [optional] [readonly] 
**Labels** | Pointer to [**[]ComponentLabel**](ComponentLabel.md) | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn&#39;t display your labels.  Cluster labels are deprecated and will be removed in a future release. We strongly recommend that you use [resource tags](https://dochub.mongodb.org/core/add-cluster-tag-atlas) instead. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBMajorVersion** | Pointer to **string** | MongoDB major version of the cluster.  On creation: Choose from the available versions of MongoDB, or leave unspecified for the current recommended default in the MongoDB Cloud platform. The recommended version is a recent Long Term Support version. The default is not guaranteed to be the most recently released version throughout the entire release cycle. For versions available in a specific project, see the linked documentation or use the API endpoint for [project LTS versions endpoint](#tag/Projects/operation/getProjectLTSVersions).   On update: Increase version only by 1 major version at a time. If the cluster is pinned to a MongoDB feature compatibility version exactly one major version below the current MongoDB version, the MongoDB version can be downgraded to the previous major version. | [optional] 
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the cluster runs. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the cluster. | [optional] 
**Paused** | Pointer to **bool** | Flag that indicates whether the cluster is paused. | [optional] 
**PitEnabled** | Pointer to **bool** | Flag that indicates whether the cluster uses continuous cloud backups. | [optional] 
**ReplicaSetScalingStrategy** | Pointer to **string** | Set this field to configure the replica set scaling mode for your cluster.  By default, Atlas scales under WORKLOAD_TYPE. This mode allows Atlas to scale your analytics nodes in parallel to your operational nodes.  When configured as SEQUENTIAL, Atlas scales all nodes sequentially. This mode is intended for steady-state workloads and applications performing latency-sensitive secondary reads.  When configured as NODE_TYPE, Atlas scales your electable nodes in parallel with your read-only and analytics nodes. This mode is intended for large, dynamic workloads requiring frequent and timely cluster tier scaling. This is the fastest scaling strategy, but it might impact latency of workloads when performing extensive secondary reads. | [optional] [default to "WORKLOAD_TYPE"]
**ReplicationSpecs** | Pointer to [**[]ReplicationSpec20240805**](ReplicationSpec20240805.md) | List of settings that configure your cluster regions. This array has one object per shard representing node configurations in each shard. For replica sets there is only one object representing node configurations. | [optional] 
**RootCertType** | Pointer to **string** | Root Certificate Authority that MongoDB Cloud cluster uses. MongoDB Cloud supports Internet Security Research Group. | [optional] [default to "ISRGROOTX1"]
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of this cluster. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. If value is &#x60;CONTINUOUS&#x60;, you must not specify **mongoDBMajorVersion**. | [optional] [default to "LTS"]

## Methods

### NewClusterDescription20240805

`func NewClusterDescription20240805() *ClusterDescription20240805`

NewClusterDescription20240805 instantiates a new ClusterDescription20240805 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDescription20240805WithDefaults

`func NewClusterDescription20240805WithDefaults() *ClusterDescription20240805`

NewClusterDescription20240805WithDefaults instantiates a new ClusterDescription20240805 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *ClusterDescription20240805) GetAcceptDataRisksAndForceReplicaSetReconfig() time.Time`

GetAcceptDataRisksAndForceReplicaSetReconfig returns the AcceptDataRisksAndForceReplicaSetReconfig field if non-nil, zero value otherwise.

### GetAcceptDataRisksAndForceReplicaSetReconfigOk

`func (o *ClusterDescription20240805) GetAcceptDataRisksAndForceReplicaSetReconfigOk() (*time.Time, bool)`

GetAcceptDataRisksAndForceReplicaSetReconfigOk returns a tuple with the AcceptDataRisksAndForceReplicaSetReconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *ClusterDescription20240805) SetAcceptDataRisksAndForceReplicaSetReconfig(v time.Time)`

SetAcceptDataRisksAndForceReplicaSetReconfig sets AcceptDataRisksAndForceReplicaSetReconfig field to given value.

### HasAcceptDataRisksAndForceReplicaSetReconfig

`func (o *ClusterDescription20240805) HasAcceptDataRisksAndForceReplicaSetReconfig() bool`

HasAcceptDataRisksAndForceReplicaSetReconfig returns a boolean if a field has been set.
### GetBackupEnabled

`func (o *ClusterDescription20240805) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *ClusterDescription20240805) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *ClusterDescription20240805) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *ClusterDescription20240805) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.
### GetBiConnector

`func (o *ClusterDescription20240805) GetBiConnector() BiConnector`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *ClusterDescription20240805) GetBiConnectorOk() (*BiConnector, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *ClusterDescription20240805) SetBiConnector(v BiConnector)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *ClusterDescription20240805) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.
### GetClusterType

`func (o *ClusterDescription20240805) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterDescription20240805) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterDescription20240805) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *ClusterDescription20240805) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.
### GetConnectionStrings

`func (o *ClusterDescription20240805) GetConnectionStrings() ClusterConnectionStrings`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *ClusterDescription20240805) GetConnectionStringsOk() (*ClusterConnectionStrings, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *ClusterDescription20240805) SetConnectionStrings(v ClusterConnectionStrings)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *ClusterDescription20240805) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.
### GetCreateDate

`func (o *ClusterDescription20240805) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ClusterDescription20240805) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ClusterDescription20240805) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ClusterDescription20240805) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.
### GetDiskWarmingMode

`func (o *ClusterDescription20240805) GetDiskWarmingMode() string`

GetDiskWarmingMode returns the DiskWarmingMode field if non-nil, zero value otherwise.

### GetDiskWarmingModeOk

`func (o *ClusterDescription20240805) GetDiskWarmingModeOk() (*string, bool)`

GetDiskWarmingModeOk returns a tuple with the DiskWarmingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWarmingMode

`func (o *ClusterDescription20240805) SetDiskWarmingMode(v string)`

SetDiskWarmingMode sets DiskWarmingMode field to given value.

### HasDiskWarmingMode

`func (o *ClusterDescription20240805) HasDiskWarmingMode() bool`

HasDiskWarmingMode returns a boolean if a field has been set.
### GetEncryptionAtRestProvider

`func (o *ClusterDescription20240805) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *ClusterDescription20240805) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *ClusterDescription20240805) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *ClusterDescription20240805) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.
### GetFeatureCompatibilityVersion

`func (o *ClusterDescription20240805) GetFeatureCompatibilityVersion() string`

GetFeatureCompatibilityVersion returns the FeatureCompatibilityVersion field if non-nil, zero value otherwise.

### GetFeatureCompatibilityVersionOk

`func (o *ClusterDescription20240805) GetFeatureCompatibilityVersionOk() (*string, bool)`

GetFeatureCompatibilityVersionOk returns a tuple with the FeatureCompatibilityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCompatibilityVersion

`func (o *ClusterDescription20240805) SetFeatureCompatibilityVersion(v string)`

SetFeatureCompatibilityVersion sets FeatureCompatibilityVersion field to given value.

### HasFeatureCompatibilityVersion

`func (o *ClusterDescription20240805) HasFeatureCompatibilityVersion() bool`

HasFeatureCompatibilityVersion returns a boolean if a field has been set.
### GetFeatureCompatibilityVersionExpirationDate

`func (o *ClusterDescription20240805) GetFeatureCompatibilityVersionExpirationDate() time.Time`

GetFeatureCompatibilityVersionExpirationDate returns the FeatureCompatibilityVersionExpirationDate field if non-nil, zero value otherwise.

### GetFeatureCompatibilityVersionExpirationDateOk

`func (o *ClusterDescription20240805) GetFeatureCompatibilityVersionExpirationDateOk() (*time.Time, bool)`

GetFeatureCompatibilityVersionExpirationDateOk returns a tuple with the FeatureCompatibilityVersionExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCompatibilityVersionExpirationDate

`func (o *ClusterDescription20240805) SetFeatureCompatibilityVersionExpirationDate(v time.Time)`

SetFeatureCompatibilityVersionExpirationDate sets FeatureCompatibilityVersionExpirationDate field to given value.

### HasFeatureCompatibilityVersionExpirationDate

`func (o *ClusterDescription20240805) HasFeatureCompatibilityVersionExpirationDate() bool`

HasFeatureCompatibilityVersionExpirationDate returns a boolean if a field has been set.
### GetGlobalClusterSelfManagedSharding

`func (o *ClusterDescription20240805) GetGlobalClusterSelfManagedSharding() bool`

GetGlobalClusterSelfManagedSharding returns the GlobalClusterSelfManagedSharding field if non-nil, zero value otherwise.

### GetGlobalClusterSelfManagedShardingOk

`func (o *ClusterDescription20240805) GetGlobalClusterSelfManagedShardingOk() (*bool, bool)`

GetGlobalClusterSelfManagedShardingOk returns a tuple with the GlobalClusterSelfManagedSharding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalClusterSelfManagedSharding

`func (o *ClusterDescription20240805) SetGlobalClusterSelfManagedSharding(v bool)`

SetGlobalClusterSelfManagedSharding sets GlobalClusterSelfManagedSharding field to given value.

### HasGlobalClusterSelfManagedSharding

`func (o *ClusterDescription20240805) HasGlobalClusterSelfManagedSharding() bool`

HasGlobalClusterSelfManagedSharding returns a boolean if a field has been set.
### GetGroupId

`func (o *ClusterDescription20240805) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ClusterDescription20240805) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ClusterDescription20240805) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ClusterDescription20240805) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *ClusterDescription20240805) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterDescription20240805) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterDescription20240805) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterDescription20240805) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLabels

`func (o *ClusterDescription20240805) GetLabels() []ComponentLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterDescription20240805) GetLabelsOk() (*[]ComponentLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterDescription20240805) SetLabels(v []ComponentLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterDescription20240805) HasLabels() bool`

HasLabels returns a boolean if a field has been set.
### GetLinks

`func (o *ClusterDescription20240805) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClusterDescription20240805) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClusterDescription20240805) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClusterDescription20240805) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMongoDBMajorVersion

`func (o *ClusterDescription20240805) GetMongoDBMajorVersion() string`

GetMongoDBMajorVersion returns the MongoDBMajorVersion field if non-nil, zero value otherwise.

### GetMongoDBMajorVersionOk

`func (o *ClusterDescription20240805) GetMongoDBMajorVersionOk() (*string, bool)`

GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBMajorVersion

`func (o *ClusterDescription20240805) SetMongoDBMajorVersion(v string)`

SetMongoDBMajorVersion sets MongoDBMajorVersion field to given value.

### HasMongoDBMajorVersion

`func (o *ClusterDescription20240805) HasMongoDBMajorVersion() bool`

HasMongoDBMajorVersion returns a boolean if a field has been set.
### GetMongoDBVersion

`func (o *ClusterDescription20240805) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ClusterDescription20240805) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ClusterDescription20240805) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ClusterDescription20240805) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.
### GetName

`func (o *ClusterDescription20240805) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterDescription20240805) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterDescription20240805) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterDescription20240805) HasName() bool`

HasName returns a boolean if a field has been set.
### GetPaused

`func (o *ClusterDescription20240805) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ClusterDescription20240805) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ClusterDescription20240805) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ClusterDescription20240805) HasPaused() bool`

HasPaused returns a boolean if a field has been set.
### GetPitEnabled

`func (o *ClusterDescription20240805) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *ClusterDescription20240805) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *ClusterDescription20240805) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *ClusterDescription20240805) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.
### GetReplicaSetScalingStrategy

`func (o *ClusterDescription20240805) GetReplicaSetScalingStrategy() string`

GetReplicaSetScalingStrategy returns the ReplicaSetScalingStrategy field if non-nil, zero value otherwise.

### GetReplicaSetScalingStrategyOk

`func (o *ClusterDescription20240805) GetReplicaSetScalingStrategyOk() (*string, bool)`

GetReplicaSetScalingStrategyOk returns a tuple with the ReplicaSetScalingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetScalingStrategy

`func (o *ClusterDescription20240805) SetReplicaSetScalingStrategy(v string)`

SetReplicaSetScalingStrategy sets ReplicaSetScalingStrategy field to given value.

### HasReplicaSetScalingStrategy

`func (o *ClusterDescription20240805) HasReplicaSetScalingStrategy() bool`

HasReplicaSetScalingStrategy returns a boolean if a field has been set.
### GetReplicationSpecs

`func (o *ClusterDescription20240805) GetReplicationSpecs() []ReplicationSpec20240805`

GetReplicationSpecs returns the ReplicationSpecs field if non-nil, zero value otherwise.

### GetReplicationSpecsOk

`func (o *ClusterDescription20240805) GetReplicationSpecsOk() (*[]ReplicationSpec20240805, bool)`

GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecs

`func (o *ClusterDescription20240805) SetReplicationSpecs(v []ReplicationSpec20240805)`

SetReplicationSpecs sets ReplicationSpecs field to given value.

### HasReplicationSpecs

`func (o *ClusterDescription20240805) HasReplicationSpecs() bool`

HasReplicationSpecs returns a boolean if a field has been set.
### GetRootCertType

`func (o *ClusterDescription20240805) GetRootCertType() string`

GetRootCertType returns the RootCertType field if non-nil, zero value otherwise.

### GetRootCertTypeOk

`func (o *ClusterDescription20240805) GetRootCertTypeOk() (*string, bool)`

GetRootCertTypeOk returns a tuple with the RootCertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCertType

`func (o *ClusterDescription20240805) SetRootCertType(v string)`

SetRootCertType sets RootCertType field to given value.

### HasRootCertType

`func (o *ClusterDescription20240805) HasRootCertType() bool`

HasRootCertType returns a boolean if a field has been set.
### GetStateName

`func (o *ClusterDescription20240805) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ClusterDescription20240805) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ClusterDescription20240805) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ClusterDescription20240805) HasStateName() bool`

HasStateName returns a boolean if a field has been set.
### GetTags

`func (o *ClusterDescription20240805) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterDescription20240805) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterDescription20240805) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterDescription20240805) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *ClusterDescription20240805) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ClusterDescription20240805) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ClusterDescription20240805) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ClusterDescription20240805) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.
### GetVersionReleaseSystem

`func (o *ClusterDescription20240805) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *ClusterDescription20240805) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *ClusterDescription20240805) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *ClusterDescription20240805) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


