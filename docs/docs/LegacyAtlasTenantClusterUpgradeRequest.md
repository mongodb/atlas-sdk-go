# LegacyAtlasTenantClusterUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptDataRisksAndForceReplicaSetReconfig** | Pointer to **time.Time** | If reconfiguration is necessary to regain a primary due to a regional outage, submit this field alongside your topology reconfiguration to request a new regional outage resistant topology. Forced reconfigurations during an outage of the majority of electable nodes carry a risk of data loss if replicated writes (even majority committed writes) have not been replicated to the new primary node. MongoDB Atlas docs contain more information. To proceed with an operation which carries that risk, set **acceptDataRisksAndForceReplicaSetReconfig** to the current date. | [optional] 
**AutoScaling** | Pointer to [**ClusterAutoScalingSettings**](ClusterAutoScalingSettings.md) |  | [optional] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] 
**BiConnector** | Pointer to [**BiConnector**](BiConnector.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Configuration of nodes that comprise the cluster. | [optional] 
**ConnectionStrings** | Pointer to [**ClusterConnectionStrings**](ClusterConnectionStrings.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this serverless instance. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**DiskSizeGB** | Pointer to **float64** | Storage capacity of instance data volumes expressed in gigabytes. Increase this number to add capacity.   This value is not configurable on M0/M2/M5 clusters.   MongoDB Cloud requires this parameter if you set **replicationSpecs**.   If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value.    Storage charge calculations depend on whether you choose the default value or a custom value.   The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier. | [optional] 
**DiskWarmingMode** | Pointer to **string** | Disk warming mode selection. | [optional] [default to "FULLY_WARMED"]
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of Encryption at Rest for the cluster. | [optional] 
**FeatureCompatibilityVersion** | Pointer to **string** | Feature compatibility version of the cluster. | [optional] [readonly] 
**FeatureCompatibilityVersionExpirationDate** | Pointer to **time.Time** | Feature compatibility version expiration date. | [optional] [readonly] 
**GlobalClusterSelfManagedSharding** | Pointer to **bool** | Set this field to configure the Sharding Management Mode when creating a new Global Cluster.  When set to false, the management mode is set to Atlas-Managed Sharding. This mode fully manages the sharding of your Global Cluster and is built to provide a seamless deployment experience.  When set to true, the management mode is set to Self-Managed Sharding. This mode leaves the management of shards in your hands and is built to provide an advanced and flexible deployment experience.  This setting cannot be changed once the cluster is deployed. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster. | [optional] [readonly] 
**Labels** | Pointer to [**[]ComponentLabel**](ComponentLabel.md) | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn&#39;t display your labels.  Cluster labels are deprecated and will be removed in a future release. We strongly recommend that you use [resource tags](https://dochub.mongodb.org/core/add-cluster-tag-atlas) instead. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBEmployeeAccessGrant** | Pointer to [**EmployeeAccessGrant**](EmployeeAccessGrant.md) |  | [optional] 
**MongoDBMajorVersion** | Pointer to **string** | MongoDB major version of the cluster.  On creation: Choose from the available versions of MongoDB, or leave unspecified for the current recommended default in the MongoDB Cloud platform. The recommended version is a recent Long Term Support version. The default is not guaranteed to be the most recently released version throughout the entire release cycle. For versions available in a specific project, see the linked documentation or use the API endpoint for [project LTS versions endpoint](#tag/Projects/operation/getProjectLTSVersions).   On update: Increase version only by 1 major version at a time. If the cluster is pinned to a MongoDB feature compatibility version exactly one major version below the current MongoDB version, the MongoDB version can be downgraded to the previous major version. | [optional] 
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the cluster runs. | [optional] 
**MongoURI** | Pointer to **string** | Base connection string that you can use to connect to the cluster. MongoDB Cloud displays the string only after the cluster starts, not while it builds the cluster. | [optional] [readonly] 
**MongoURIUpdated** | Pointer to **time.Time** | Date and time when someone last updated the connection string. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**MongoURIWithOptions** | Pointer to **string** | Connection string that you can use to connect to the cluster including the &#x60;replicaSet&#x60;, &#x60;ssl&#x60;, and &#x60;authSource&#x60; query parameters with values appropriate for the cluster. You may need to add MongoDB database users. The response returns this parameter once the cluster can receive requests, not while it builds the cluster. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies the cluster. | 
**NumShards** | Pointer to **int** | Number of shards up to 50 to deploy for a sharded cluster. The resource returns &#x60;1&#x60; to indicate a replica set and values of &#x60;2&#x60; and higher to indicate a sharded cluster. The returned value equals the number of shards in the cluster. | [optional] [default to 1]
**Paused** | Pointer to **bool** | Flag that indicates whether the cluster is paused. | [optional] 
**PitEnabled** | Pointer to **bool** | Flag that indicates whether the cluster uses continuous cloud backups. | [optional] 
**ProviderBackupEnabled** | Pointer to **bool** | Flag that indicates whether the M10 or higher cluster can perform Cloud Backups. If set to &#x60;true&#x60;, the cluster can perform backups. If this and **backupEnabled** are set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] 
**ProviderSettings** | Pointer to [**ClusterProviderSettings**](ClusterProviderSettings.md) |  | [optional] 
**ReplicaSetScalingStrategy** | Pointer to **string** | Set this field to configure the replica set scaling mode for your cluster.  By default, Atlas scales under WORKLOAD_TYPE. This mode allows Atlas to scale your analytics nodes in parallel to your operational nodes.  When configured as SEQUENTIAL, Atlas scales all nodes sequentially. This mode is intended for steady-state workloads and applications performing latency-sensitive secondary reads.  When configured as NODE_TYPE, Atlas scales your electable nodes in parallel with your read-only and analytics nodes. This mode is intended for large, dynamic workloads requiring frequent and timely cluster tier scaling. This is the fastest scaling strategy, but it might impact latency of workloads when performing extensive secondary reads. | [optional] [default to "WORKLOAD_TYPE"]
**ReplicationFactor** | Pointer to **int** | Number of members that belong to the replica set. Each member retains a copy of your databases, providing high availability and data redundancy. Use **replicationSpecs** instead. | [optional] [default to 3]
**ReplicationSpec** | Pointer to [**map[string]RegionSpec**](RegionSpec.md) | Physical location where MongoDB Cloud provisions cluster nodes. | [optional] 
**ReplicationSpecs** | Pointer to [**[]LegacyReplicationSpec**](LegacyReplicationSpec.md) | List of settings that configure your cluster regions.  - For Global Clusters, each object in the array represents one zone where MongoDB Cloud deploys your clusters nodes. - For non-Global sharded clusters and replica sets, the single object represents where MongoDB Cloud deploys your clusters nodes. | [optional] 
**RootCertType** | Pointer to **string** | Root Certificate Authority that MongoDB Atlas clusters uses. MongoDB Cloud supports Internet Security Research Group. | [optional] [default to "ISRGROOTX1"]
**SrvAddress** | Pointer to **string** | Connection string that you can use to connect to the cluster. The &#x60;+srv&#x60; modifier forces the connection to use Transport Layer Security (TLS). The &#x60;mongoURI&#x60; parameter lists additional options. | [optional] [readonly] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of the cluster. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. If value is &#x60;CONTINUOUS&#x60;, you must not specify **mongoDBMajorVersion**. | [optional] [default to "LTS"]

## Methods

### NewLegacyAtlasTenantClusterUpgradeRequest

`func NewLegacyAtlasTenantClusterUpgradeRequest(name string, ) *LegacyAtlasTenantClusterUpgradeRequest`

NewLegacyAtlasTenantClusterUpgradeRequest instantiates a new LegacyAtlasTenantClusterUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyAtlasTenantClusterUpgradeRequestWithDefaults

`func NewLegacyAtlasTenantClusterUpgradeRequestWithDefaults() *LegacyAtlasTenantClusterUpgradeRequest`

NewLegacyAtlasTenantClusterUpgradeRequestWithDefaults instantiates a new LegacyAtlasTenantClusterUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetAcceptDataRisksAndForceReplicaSetReconfig() time.Time`

GetAcceptDataRisksAndForceReplicaSetReconfig returns the AcceptDataRisksAndForceReplicaSetReconfig field if non-nil, zero value otherwise.

### GetAcceptDataRisksAndForceReplicaSetReconfigOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetAcceptDataRisksAndForceReplicaSetReconfigOk() (*time.Time, bool)`

GetAcceptDataRisksAndForceReplicaSetReconfigOk returns a tuple with the AcceptDataRisksAndForceReplicaSetReconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetAcceptDataRisksAndForceReplicaSetReconfig(v time.Time)`

SetAcceptDataRisksAndForceReplicaSetReconfig sets AcceptDataRisksAndForceReplicaSetReconfig field to given value.

### HasAcceptDataRisksAndForceReplicaSetReconfig

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasAcceptDataRisksAndForceReplicaSetReconfig() bool`

HasAcceptDataRisksAndForceReplicaSetReconfig returns a boolean if a field has been set.
### GetAutoScaling

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetAutoScaling() ClusterAutoScalingSettings`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetAutoScalingOk() (*ClusterAutoScalingSettings, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetAutoScaling(v ClusterAutoScalingSettings)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.
### GetBackupEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.
### GetBiConnector

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetBiConnector() BiConnector`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetBiConnectorOk() (*BiConnector, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetBiConnector(v BiConnector)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.
### GetClusterType

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.
### GetConnectionStrings

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetConnectionStrings() ClusterConnectionStrings`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetConnectionStringsOk() (*ClusterConnectionStrings, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetConnectionStrings(v ClusterConnectionStrings)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.
### GetCreateDate

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.
### GetDiskSizeGB

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetDiskSizeGB() float64`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetDiskSizeGBOk() (*float64, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetDiskSizeGB(v float64)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.
### GetDiskWarmingMode

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetDiskWarmingMode() string`

GetDiskWarmingMode returns the DiskWarmingMode field if non-nil, zero value otherwise.

### GetDiskWarmingModeOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetDiskWarmingModeOk() (*string, bool)`

GetDiskWarmingModeOk returns a tuple with the DiskWarmingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWarmingMode

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetDiskWarmingMode(v string)`

SetDiskWarmingMode sets DiskWarmingMode field to given value.

### HasDiskWarmingMode

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasDiskWarmingMode() bool`

HasDiskWarmingMode returns a boolean if a field has been set.
### GetEncryptionAtRestProvider

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.
### GetFeatureCompatibilityVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetFeatureCompatibilityVersion() string`

GetFeatureCompatibilityVersion returns the FeatureCompatibilityVersion field if non-nil, zero value otherwise.

### GetFeatureCompatibilityVersionOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetFeatureCompatibilityVersionOk() (*string, bool)`

GetFeatureCompatibilityVersionOk returns a tuple with the FeatureCompatibilityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCompatibilityVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetFeatureCompatibilityVersion(v string)`

SetFeatureCompatibilityVersion sets FeatureCompatibilityVersion field to given value.

### HasFeatureCompatibilityVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasFeatureCompatibilityVersion() bool`

HasFeatureCompatibilityVersion returns a boolean if a field has been set.
### GetFeatureCompatibilityVersionExpirationDate

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetFeatureCompatibilityVersionExpirationDate() time.Time`

GetFeatureCompatibilityVersionExpirationDate returns the FeatureCompatibilityVersionExpirationDate field if non-nil, zero value otherwise.

### GetFeatureCompatibilityVersionExpirationDateOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetFeatureCompatibilityVersionExpirationDateOk() (*time.Time, bool)`

GetFeatureCompatibilityVersionExpirationDateOk returns a tuple with the FeatureCompatibilityVersionExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCompatibilityVersionExpirationDate

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetFeatureCompatibilityVersionExpirationDate(v time.Time)`

SetFeatureCompatibilityVersionExpirationDate sets FeatureCompatibilityVersionExpirationDate field to given value.

### HasFeatureCompatibilityVersionExpirationDate

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasFeatureCompatibilityVersionExpirationDate() bool`

HasFeatureCompatibilityVersionExpirationDate returns a boolean if a field has been set.
### GetGlobalClusterSelfManagedSharding

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetGlobalClusterSelfManagedSharding() bool`

GetGlobalClusterSelfManagedSharding returns the GlobalClusterSelfManagedSharding field if non-nil, zero value otherwise.

### GetGlobalClusterSelfManagedShardingOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetGlobalClusterSelfManagedShardingOk() (*bool, bool)`

GetGlobalClusterSelfManagedShardingOk returns a tuple with the GlobalClusterSelfManagedSharding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalClusterSelfManagedSharding

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetGlobalClusterSelfManagedSharding(v bool)`

SetGlobalClusterSelfManagedSharding sets GlobalClusterSelfManagedSharding field to given value.

### HasGlobalClusterSelfManagedSharding

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasGlobalClusterSelfManagedSharding() bool`

HasGlobalClusterSelfManagedSharding returns a boolean if a field has been set.
### GetGroupId

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLabels

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetLabels() []ComponentLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetLabelsOk() (*[]ComponentLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetLabels(v []ComponentLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.
### GetLinks

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMongoDBEmployeeAccessGrant

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoDBEmployeeAccessGrant() EmployeeAccessGrant`

GetMongoDBEmployeeAccessGrant returns the MongoDBEmployeeAccessGrant field if non-nil, zero value otherwise.

### GetMongoDBEmployeeAccessGrantOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoDBEmployeeAccessGrantOk() (*EmployeeAccessGrant, bool)`

GetMongoDBEmployeeAccessGrantOk returns a tuple with the MongoDBEmployeeAccessGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBEmployeeAccessGrant

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetMongoDBEmployeeAccessGrant(v EmployeeAccessGrant)`

SetMongoDBEmployeeAccessGrant sets MongoDBEmployeeAccessGrant field to given value.

### HasMongoDBEmployeeAccessGrant

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasMongoDBEmployeeAccessGrant() bool`

HasMongoDBEmployeeAccessGrant returns a boolean if a field has been set.
### GetMongoDBMajorVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoDBMajorVersion() string`

GetMongoDBMajorVersion returns the MongoDBMajorVersion field if non-nil, zero value otherwise.

### GetMongoDBMajorVersionOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoDBMajorVersionOk() (*string, bool)`

GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBMajorVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetMongoDBMajorVersion(v string)`

SetMongoDBMajorVersion sets MongoDBMajorVersion field to given value.

### HasMongoDBMajorVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasMongoDBMajorVersion() bool`

HasMongoDBMajorVersion returns a boolean if a field has been set.
### GetMongoDBVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.
### GetMongoURI

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoURI() string`

GetMongoURI returns the MongoURI field if non-nil, zero value otherwise.

### GetMongoURIOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoURIOk() (*string, bool)`

GetMongoURIOk returns a tuple with the MongoURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURI

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetMongoURI(v string)`

SetMongoURI sets MongoURI field to given value.

### HasMongoURI

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasMongoURI() bool`

HasMongoURI returns a boolean if a field has been set.
### GetMongoURIUpdated

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoURIUpdated() time.Time`

GetMongoURIUpdated returns the MongoURIUpdated field if non-nil, zero value otherwise.

### GetMongoURIUpdatedOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoURIUpdatedOk() (*time.Time, bool)`

GetMongoURIUpdatedOk returns a tuple with the MongoURIUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIUpdated

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetMongoURIUpdated(v time.Time)`

SetMongoURIUpdated sets MongoURIUpdated field to given value.

### HasMongoURIUpdated

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasMongoURIUpdated() bool`

HasMongoURIUpdated returns a boolean if a field has been set.
### GetMongoURIWithOptions

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoURIWithOptions() string`

GetMongoURIWithOptions returns the MongoURIWithOptions field if non-nil, zero value otherwise.

### GetMongoURIWithOptionsOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetMongoURIWithOptionsOk() (*string, bool)`

GetMongoURIWithOptionsOk returns a tuple with the MongoURIWithOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIWithOptions

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetMongoURIWithOptions(v string)`

SetMongoURIWithOptions sets MongoURIWithOptions field to given value.

### HasMongoURIWithOptions

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasMongoURIWithOptions() bool`

HasMongoURIWithOptions returns a boolean if a field has been set.
### GetName

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetName(v string)`

SetName sets Name field to given value.

### GetNumShards

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetNumShards() int`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetNumShardsOk() (*int, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetNumShards(v int)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.
### GetPaused

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasPaused() bool`

HasPaused returns a boolean if a field has been set.
### GetPitEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.
### GetProviderBackupEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetProviderBackupEnabled() bool`

GetProviderBackupEnabled returns the ProviderBackupEnabled field if non-nil, zero value otherwise.

### GetProviderBackupEnabledOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetProviderBackupEnabledOk() (*bool, bool)`

GetProviderBackupEnabledOk returns a tuple with the ProviderBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBackupEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetProviderBackupEnabled(v bool)`

SetProviderBackupEnabled sets ProviderBackupEnabled field to given value.

### HasProviderBackupEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasProviderBackupEnabled() bool`

HasProviderBackupEnabled returns a boolean if a field has been set.
### GetProviderSettings

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetProviderSettings() ClusterProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetProviderSettingsOk() (*ClusterProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetProviderSettings(v ClusterProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.
### GetReplicaSetScalingStrategy

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicaSetScalingStrategy() string`

GetReplicaSetScalingStrategy returns the ReplicaSetScalingStrategy field if non-nil, zero value otherwise.

### GetReplicaSetScalingStrategyOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicaSetScalingStrategyOk() (*string, bool)`

GetReplicaSetScalingStrategyOk returns a tuple with the ReplicaSetScalingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetScalingStrategy

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetReplicaSetScalingStrategy(v string)`

SetReplicaSetScalingStrategy sets ReplicaSetScalingStrategy field to given value.

### HasReplicaSetScalingStrategy

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasReplicaSetScalingStrategy() bool`

HasReplicaSetScalingStrategy returns a boolean if a field has been set.
### GetReplicationFactor

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicationFactor() int`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicationFactorOk() (*int, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetReplicationFactor(v int)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.
### GetReplicationSpec

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicationSpec() map[string]RegionSpec`

GetReplicationSpec returns the ReplicationSpec field if non-nil, zero value otherwise.

### GetReplicationSpecOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicationSpecOk() (*map[string]RegionSpec, bool)`

GetReplicationSpecOk returns a tuple with the ReplicationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpec

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetReplicationSpec(v map[string]RegionSpec)`

SetReplicationSpec sets ReplicationSpec field to given value.

### HasReplicationSpec

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasReplicationSpec() bool`

HasReplicationSpec returns a boolean if a field has been set.
### GetReplicationSpecs

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicationSpecs() []LegacyReplicationSpec`

GetReplicationSpecs returns the ReplicationSpecs field if non-nil, zero value otherwise.

### GetReplicationSpecsOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetReplicationSpecsOk() (*[]LegacyReplicationSpec, bool)`

GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecs

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetReplicationSpecs(v []LegacyReplicationSpec)`

SetReplicationSpecs sets ReplicationSpecs field to given value.

### HasReplicationSpecs

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasReplicationSpecs() bool`

HasReplicationSpecs returns a boolean if a field has been set.
### GetRootCertType

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetRootCertType() string`

GetRootCertType returns the RootCertType field if non-nil, zero value otherwise.

### GetRootCertTypeOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetRootCertTypeOk() (*string, bool)`

GetRootCertTypeOk returns a tuple with the RootCertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCertType

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetRootCertType(v string)`

SetRootCertType sets RootCertType field to given value.

### HasRootCertType

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasRootCertType() bool`

HasRootCertType returns a boolean if a field has been set.
### GetSrvAddress

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetSrvAddress() string`

GetSrvAddress returns the SrvAddress field if non-nil, zero value otherwise.

### GetSrvAddressOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetSrvAddressOk() (*string, bool)`

GetSrvAddressOk returns a tuple with the SrvAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvAddress

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetSrvAddress(v string)`

SetSrvAddress sets SrvAddress field to given value.

### HasSrvAddress

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasSrvAddress() bool`

HasSrvAddress returns a boolean if a field has been set.
### GetStateName

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasStateName() bool`

HasStateName returns a boolean if a field has been set.
### GetTags

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.
### GetVersionReleaseSystem

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *LegacyAtlasTenantClusterUpgradeRequest) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *LegacyAtlasTenantClusterUpgradeRequest) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *LegacyAtlasTenantClusterUpgradeRequest) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


