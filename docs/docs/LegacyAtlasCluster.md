# LegacyAtlasCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptDataRisksAndForceReplicaSetReconfig** | Pointer to **time.Time** | If reconfiguration is necessary to regain a primary due to a regional outage, submit this field alongside your topology reconfiguration to request a new regional outage resistant topology. Forced reconfigurations during an outage of the majority of electable nodes carry a risk of data loss if replicated writes (even majority committed writes) have not been replicated to the new primary node. MongoDB Atlas docs contain more information. To proceed with an operation which carries that risk, set **acceptDataRisksAndForceReplicaSetReconfig** to the current date. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] 
**AdvancedConfiguration** | Pointer to [**ApiAtlasClusterAdvancedConfiguration**](ApiAtlasClusterAdvancedConfiguration.md) |  | [optional] 
**AutoScaling** | Pointer to [**ClusterAutoScalingSettings**](ClusterAutoScalingSettings.md) |  | [optional] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] 
**BiConnector** | Pointer to [**BiConnector**](BiConnector.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Configuration of nodes that comprise the cluster. | [optional] 
**ConfigServerManagementMode** | Pointer to **string** | Config Server Management Mode for creating or updating a sharded cluster.  When configured as ATLAS_MANAGED, atlas may automatically switch the cluster&#39;s config server type for optimal performance and savings.  When configured as FIXED_TO_DEDICATED, the cluster will always use a dedicated config server. | [optional] [default to "ATLAS_MANAGED"]
**ConfigServerType** | Pointer to **string** | Describes a sharded cluster&#39;s config server type. | [optional] [readonly] 
**ConnectionStrings** | Pointer to [**ClusterConnectionStrings**](ClusterConnectionStrings.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this serverless instance. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**DiskSizeGB** | Pointer to **float64** | Storage capacity of instance data volumes expressed in gigabytes. Increase this number to add capacity.   This value is not configurable on M0/M2/M5 clusters.   MongoDB Cloud requires this parameter if you set **replicationSpecs**.   If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value.    Storage charge calculations depend on whether you choose the default value or a custom value.   The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier. | [optional] 
**DiskWarmingMode** | Pointer to **string** | Disk warming mode selection. | [optional] [default to "FULLY_WARMED"]
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. To enable customer key management for encryption at rest, the cluster **replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize** setting must be &#x60;M10&#x60; or higher and &#x60;\&quot;backupEnabled\&quot; : false&#x60; or omitted entirely. | [optional] 
**FeatureCompatibilityVersion** | Pointer to **string** | Feature compatibility version of the cluster. | [optional] [readonly] 
**FeatureCompatibilityVersionExpirationDate** | Pointer to **time.Time** | Feature compatibility version expiration date. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GlobalClusterSelfManagedSharding** | Pointer to **bool** | Set this field to configure the Sharding Management Mode when creating a new Global Cluster.  When set to false, the management mode is set to Atlas-Managed Sharding. This mode fully manages the sharding of your Global Cluster and is built to provide a seamless deployment experience.  When set to true, the management mode is set to Self-Managed Sharding. This mode leaves the management of shards in your hands and is built to provide an advanced and flexible deployment experience.  This setting cannot be changed once the cluster is deployed. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster. | [optional] [readonly] 
**Labels** | Pointer to [**[]ComponentLabel**](ComponentLabel.md) | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn&#39;t display your labels.  Cluster labels are deprecated and will be removed in a future release. We strongly recommend that you use Resource Tags instead. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBEmployeeAccessGrant** | Pointer to [**EmployeeAccessGrant**](EmployeeAccessGrant.md) |  | [optional] 
**MongoDBMajorVersion** | Pointer to **string** | MongoDB major version of the cluster.  On creation: Choose from the available versions of MongoDB, or leave unspecified for the current recommended default in the MongoDB Cloud platform. The recommended version is a recent Long Term Support version. The default is not guaranteed to be the most recently released version throughout the entire release cycle. For versions available in a specific project, see the linked documentation or use the API endpoint for [project LTS versions endpoint](#tag/Projects/operation/getProjectLTSVersions).   On update: Increase version only by 1 major version at a time. If the cluster is pinned to a MongoDB feature compatibility version exactly one major version below the current MongoDB version, the MongoDB version can be downgraded to the previous major version. | [optional] 
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the cluster runs. | [optional] 
**MongoURI** | Pointer to **string** | Base connection string that you can use to connect to the cluster. MongoDB Cloud displays the string only after the cluster starts, not while it builds the cluster. | [optional] [readonly] 
**MongoURIUpdated** | Pointer to **time.Time** | Date and time when someone last updated the connection string. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**MongoURIWithOptions** | Pointer to **string** | Connection string that you can use to connect to the cluster including the &#x60;replicaSet&#x60;, &#x60;ssl&#x60;, and &#x60;authSource&#x60; query parameters with values appropriate for the cluster. You may need to add MongoDB database users. The response returns this parameter once the cluster can receive requests, not while it builds the cluster. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the cluster. | [optional] 
**NumShards** | Pointer to **int** | Number of shards up to 50 to deploy for a sharded cluster. The resource returns &#x60;1&#x60; to indicate a replica set and values of &#x60;2&#x60; and higher to indicate a sharded cluster. The returned value equals the number of shards in the cluster. | [optional] [default to 1]
**Paused** | Pointer to **bool** | Flag that indicates whether the cluster is paused. | [optional] 
**PitEnabled** | Pointer to **bool** | Flag that indicates whether the cluster uses continuous cloud backups. | [optional] 
**ProviderBackupEnabled** | Pointer to **bool** | Flag that indicates whether the M10 or higher cluster can perform Cloud Backups. If set to &#x60;true&#x60;, the cluster can perform backups. If this and **backupEnabled** are set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] 
**ProviderSettings** | Pointer to [**ClusterProviderSettings**](ClusterProviderSettings.md) |  | [optional] 
**ReplicaSetScalingStrategy** | Pointer to **string** | Set this field to configure the replica set scaling mode for your cluster.  By default, Atlas scales under WORKLOAD_TYPE. This mode allows Atlas to scale your analytics nodes in parallel to your operational nodes.  When configured as SEQUENTIAL, Atlas scales all nodes sequentially. This mode is intended for steady-state workloads and applications performing latency-sensitive secondary reads.  When configured as NODE_TYPE, Atlas scales your electable nodes in parallel with your read-only and analytics nodes. This mode is intended for large, dynamic workloads requiring frequent and timely cluster tier scaling. This is the fastest scaling strategy, but it might impact latency of workloads when performing extensive secondary reads. | [optional] [default to "WORKLOAD_TYPE"]
**ReplicationFactor** | Pointer to **int** | Number of members that belong to the replica set. Each member retains a copy of your databases, providing high availability and data redundancy. Use **replicationSpecs** instead. | [optional] [default to 3]
**ReplicationSpec** | Pointer to [**map[string]RegionSpec**](RegionSpec.md) | Physical location where MongoDB Cloud provisions cluster nodes. | [optional] 
**ReplicationSpecs** | Pointer to [**[]LegacyReplicationSpec**](LegacyReplicationSpec.md) | List of settings that configure your cluster regions.  - For Global Clusters, each object in the array represents one zone where MongoDB Cloud deploys your clusters nodes. - For non-Global sharded clusters and replica sets, the single object represents where MongoDB Cloud deploys your clusters nodes. | [optional] 
**RootCertType** | Pointer to **string** | Root Certificate Authority that MongoDB Atlas cluster uses. MongoDB Cloud supports Internet Security Research Group. | [optional] [default to "ISRGROOTX1"]
**SrvAddress** | Pointer to **string** | Connection string that you can use to connect to the cluster. The &#x60;+srv&#x60; modifier forces the connection to use Transport Layer Security (TLS). The &#x60;mongoURI&#x60; parameter lists additional options. | [optional] [readonly] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of the cluster. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. If value is &#x60;CONTINUOUS&#x60;, you must not specify **mongoDBMajorVersion**. | [optional] [default to "LTS"]

## Methods

### NewLegacyAtlasCluster

`func NewLegacyAtlasCluster() *LegacyAtlasCluster`

NewLegacyAtlasCluster instantiates a new LegacyAtlasCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyAtlasClusterWithDefaults

`func NewLegacyAtlasClusterWithDefaults() *LegacyAtlasCluster`

NewLegacyAtlasClusterWithDefaults instantiates a new LegacyAtlasCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *LegacyAtlasCluster) GetAcceptDataRisksAndForceReplicaSetReconfig() time.Time`

GetAcceptDataRisksAndForceReplicaSetReconfig returns the AcceptDataRisksAndForceReplicaSetReconfig field if non-nil, zero value otherwise.

### GetAcceptDataRisksAndForceReplicaSetReconfigOk

`func (o *LegacyAtlasCluster) GetAcceptDataRisksAndForceReplicaSetReconfigOk() (*time.Time, bool)`

GetAcceptDataRisksAndForceReplicaSetReconfigOk returns a tuple with the AcceptDataRisksAndForceReplicaSetReconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptDataRisksAndForceReplicaSetReconfig

`func (o *LegacyAtlasCluster) SetAcceptDataRisksAndForceReplicaSetReconfig(v time.Time)`

SetAcceptDataRisksAndForceReplicaSetReconfig sets AcceptDataRisksAndForceReplicaSetReconfig field to given value.

### HasAcceptDataRisksAndForceReplicaSetReconfig

`func (o *LegacyAtlasCluster) HasAcceptDataRisksAndForceReplicaSetReconfig() bool`

HasAcceptDataRisksAndForceReplicaSetReconfig returns a boolean if a field has been set.
### GetAdvancedConfiguration

`func (o *LegacyAtlasCluster) GetAdvancedConfiguration() ApiAtlasClusterAdvancedConfiguration`

GetAdvancedConfiguration returns the AdvancedConfiguration field if non-nil, zero value otherwise.

### GetAdvancedConfigurationOk

`func (o *LegacyAtlasCluster) GetAdvancedConfigurationOk() (*ApiAtlasClusterAdvancedConfiguration, bool)`

GetAdvancedConfigurationOk returns a tuple with the AdvancedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfiguration

`func (o *LegacyAtlasCluster) SetAdvancedConfiguration(v ApiAtlasClusterAdvancedConfiguration)`

SetAdvancedConfiguration sets AdvancedConfiguration field to given value.

### HasAdvancedConfiguration

`func (o *LegacyAtlasCluster) HasAdvancedConfiguration() bool`

HasAdvancedConfiguration returns a boolean if a field has been set.
### GetAutoScaling

`func (o *LegacyAtlasCluster) GetAutoScaling() ClusterAutoScalingSettings`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *LegacyAtlasCluster) GetAutoScalingOk() (*ClusterAutoScalingSettings, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *LegacyAtlasCluster) SetAutoScaling(v ClusterAutoScalingSettings)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *LegacyAtlasCluster) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.
### GetBackupEnabled

`func (o *LegacyAtlasCluster) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *LegacyAtlasCluster) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *LegacyAtlasCluster) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *LegacyAtlasCluster) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.
### GetBiConnector

`func (o *LegacyAtlasCluster) GetBiConnector() BiConnector`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *LegacyAtlasCluster) GetBiConnectorOk() (*BiConnector, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *LegacyAtlasCluster) SetBiConnector(v BiConnector)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *LegacyAtlasCluster) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.
### GetClusterType

`func (o *LegacyAtlasCluster) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *LegacyAtlasCluster) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *LegacyAtlasCluster) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *LegacyAtlasCluster) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.
### GetConfigServerManagementMode

`func (o *LegacyAtlasCluster) GetConfigServerManagementMode() string`

GetConfigServerManagementMode returns the ConfigServerManagementMode field if non-nil, zero value otherwise.

### GetConfigServerManagementModeOk

`func (o *LegacyAtlasCluster) GetConfigServerManagementModeOk() (*string, bool)`

GetConfigServerManagementModeOk returns a tuple with the ConfigServerManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigServerManagementMode

`func (o *LegacyAtlasCluster) SetConfigServerManagementMode(v string)`

SetConfigServerManagementMode sets ConfigServerManagementMode field to given value.

### HasConfigServerManagementMode

`func (o *LegacyAtlasCluster) HasConfigServerManagementMode() bool`

HasConfigServerManagementMode returns a boolean if a field has been set.
### GetConfigServerType

`func (o *LegacyAtlasCluster) GetConfigServerType() string`

GetConfigServerType returns the ConfigServerType field if non-nil, zero value otherwise.

### GetConfigServerTypeOk

`func (o *LegacyAtlasCluster) GetConfigServerTypeOk() (*string, bool)`

GetConfigServerTypeOk returns a tuple with the ConfigServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigServerType

`func (o *LegacyAtlasCluster) SetConfigServerType(v string)`

SetConfigServerType sets ConfigServerType field to given value.

### HasConfigServerType

`func (o *LegacyAtlasCluster) HasConfigServerType() bool`

HasConfigServerType returns a boolean if a field has been set.
### GetConnectionStrings

`func (o *LegacyAtlasCluster) GetConnectionStrings() ClusterConnectionStrings`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *LegacyAtlasCluster) GetConnectionStringsOk() (*ClusterConnectionStrings, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *LegacyAtlasCluster) SetConnectionStrings(v ClusterConnectionStrings)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *LegacyAtlasCluster) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.
### GetCreateDate

`func (o *LegacyAtlasCluster) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *LegacyAtlasCluster) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *LegacyAtlasCluster) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *LegacyAtlasCluster) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.
### GetDiskSizeGB

`func (o *LegacyAtlasCluster) GetDiskSizeGB() float64`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *LegacyAtlasCluster) GetDiskSizeGBOk() (*float64, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *LegacyAtlasCluster) SetDiskSizeGB(v float64)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *LegacyAtlasCluster) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.
### GetDiskWarmingMode

`func (o *LegacyAtlasCluster) GetDiskWarmingMode() string`

GetDiskWarmingMode returns the DiskWarmingMode field if non-nil, zero value otherwise.

### GetDiskWarmingModeOk

`func (o *LegacyAtlasCluster) GetDiskWarmingModeOk() (*string, bool)`

GetDiskWarmingModeOk returns a tuple with the DiskWarmingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWarmingMode

`func (o *LegacyAtlasCluster) SetDiskWarmingMode(v string)`

SetDiskWarmingMode sets DiskWarmingMode field to given value.

### HasDiskWarmingMode

`func (o *LegacyAtlasCluster) HasDiskWarmingMode() bool`

HasDiskWarmingMode returns a boolean if a field has been set.
### GetEncryptionAtRestProvider

`func (o *LegacyAtlasCluster) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *LegacyAtlasCluster) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *LegacyAtlasCluster) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *LegacyAtlasCluster) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.
### GetFeatureCompatibilityVersion

`func (o *LegacyAtlasCluster) GetFeatureCompatibilityVersion() string`

GetFeatureCompatibilityVersion returns the FeatureCompatibilityVersion field if non-nil, zero value otherwise.

### GetFeatureCompatibilityVersionOk

`func (o *LegacyAtlasCluster) GetFeatureCompatibilityVersionOk() (*string, bool)`

GetFeatureCompatibilityVersionOk returns a tuple with the FeatureCompatibilityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCompatibilityVersion

`func (o *LegacyAtlasCluster) SetFeatureCompatibilityVersion(v string)`

SetFeatureCompatibilityVersion sets FeatureCompatibilityVersion field to given value.

### HasFeatureCompatibilityVersion

`func (o *LegacyAtlasCluster) HasFeatureCompatibilityVersion() bool`

HasFeatureCompatibilityVersion returns a boolean if a field has been set.
### GetFeatureCompatibilityVersionExpirationDate

`func (o *LegacyAtlasCluster) GetFeatureCompatibilityVersionExpirationDate() time.Time`

GetFeatureCompatibilityVersionExpirationDate returns the FeatureCompatibilityVersionExpirationDate field if non-nil, zero value otherwise.

### GetFeatureCompatibilityVersionExpirationDateOk

`func (o *LegacyAtlasCluster) GetFeatureCompatibilityVersionExpirationDateOk() (*time.Time, bool)`

GetFeatureCompatibilityVersionExpirationDateOk returns a tuple with the FeatureCompatibilityVersionExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCompatibilityVersionExpirationDate

`func (o *LegacyAtlasCluster) SetFeatureCompatibilityVersionExpirationDate(v time.Time)`

SetFeatureCompatibilityVersionExpirationDate sets FeatureCompatibilityVersionExpirationDate field to given value.

### HasFeatureCompatibilityVersionExpirationDate

`func (o *LegacyAtlasCluster) HasFeatureCompatibilityVersionExpirationDate() bool`

HasFeatureCompatibilityVersionExpirationDate returns a boolean if a field has been set.
### GetGlobalClusterSelfManagedSharding

`func (o *LegacyAtlasCluster) GetGlobalClusterSelfManagedSharding() bool`

GetGlobalClusterSelfManagedSharding returns the GlobalClusterSelfManagedSharding field if non-nil, zero value otherwise.

### GetGlobalClusterSelfManagedShardingOk

`func (o *LegacyAtlasCluster) GetGlobalClusterSelfManagedShardingOk() (*bool, bool)`

GetGlobalClusterSelfManagedShardingOk returns a tuple with the GlobalClusterSelfManagedSharding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalClusterSelfManagedSharding

`func (o *LegacyAtlasCluster) SetGlobalClusterSelfManagedSharding(v bool)`

SetGlobalClusterSelfManagedSharding sets GlobalClusterSelfManagedSharding field to given value.

### HasGlobalClusterSelfManagedSharding

`func (o *LegacyAtlasCluster) HasGlobalClusterSelfManagedSharding() bool`

HasGlobalClusterSelfManagedSharding returns a boolean if a field has been set.
### GetGroupId

`func (o *LegacyAtlasCluster) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LegacyAtlasCluster) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LegacyAtlasCluster) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LegacyAtlasCluster) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *LegacyAtlasCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyAtlasCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyAtlasCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyAtlasCluster) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLabels

`func (o *LegacyAtlasCluster) GetLabels() []ComponentLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LegacyAtlasCluster) GetLabelsOk() (*[]ComponentLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LegacyAtlasCluster) SetLabels(v []ComponentLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LegacyAtlasCluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.
### GetLinks

`func (o *LegacyAtlasCluster) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LegacyAtlasCluster) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LegacyAtlasCluster) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LegacyAtlasCluster) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMongoDBEmployeeAccessGrant

`func (o *LegacyAtlasCluster) GetMongoDBEmployeeAccessGrant() EmployeeAccessGrant`

GetMongoDBEmployeeAccessGrant returns the MongoDBEmployeeAccessGrant field if non-nil, zero value otherwise.

### GetMongoDBEmployeeAccessGrantOk

`func (o *LegacyAtlasCluster) GetMongoDBEmployeeAccessGrantOk() (*EmployeeAccessGrant, bool)`

GetMongoDBEmployeeAccessGrantOk returns a tuple with the MongoDBEmployeeAccessGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBEmployeeAccessGrant

`func (o *LegacyAtlasCluster) SetMongoDBEmployeeAccessGrant(v EmployeeAccessGrant)`

SetMongoDBEmployeeAccessGrant sets MongoDBEmployeeAccessGrant field to given value.

### HasMongoDBEmployeeAccessGrant

`func (o *LegacyAtlasCluster) HasMongoDBEmployeeAccessGrant() bool`

HasMongoDBEmployeeAccessGrant returns a boolean if a field has been set.
### GetMongoDBMajorVersion

`func (o *LegacyAtlasCluster) GetMongoDBMajorVersion() string`

GetMongoDBMajorVersion returns the MongoDBMajorVersion field if non-nil, zero value otherwise.

### GetMongoDBMajorVersionOk

`func (o *LegacyAtlasCluster) GetMongoDBMajorVersionOk() (*string, bool)`

GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBMajorVersion

`func (o *LegacyAtlasCluster) SetMongoDBMajorVersion(v string)`

SetMongoDBMajorVersion sets MongoDBMajorVersion field to given value.

### HasMongoDBMajorVersion

`func (o *LegacyAtlasCluster) HasMongoDBMajorVersion() bool`

HasMongoDBMajorVersion returns a boolean if a field has been set.
### GetMongoDBVersion

`func (o *LegacyAtlasCluster) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *LegacyAtlasCluster) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *LegacyAtlasCluster) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *LegacyAtlasCluster) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.
### GetMongoURI

`func (o *LegacyAtlasCluster) GetMongoURI() string`

GetMongoURI returns the MongoURI field if non-nil, zero value otherwise.

### GetMongoURIOk

`func (o *LegacyAtlasCluster) GetMongoURIOk() (*string, bool)`

GetMongoURIOk returns a tuple with the MongoURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURI

`func (o *LegacyAtlasCluster) SetMongoURI(v string)`

SetMongoURI sets MongoURI field to given value.

### HasMongoURI

`func (o *LegacyAtlasCluster) HasMongoURI() bool`

HasMongoURI returns a boolean if a field has been set.
### GetMongoURIUpdated

`func (o *LegacyAtlasCluster) GetMongoURIUpdated() time.Time`

GetMongoURIUpdated returns the MongoURIUpdated field if non-nil, zero value otherwise.

### GetMongoURIUpdatedOk

`func (o *LegacyAtlasCluster) GetMongoURIUpdatedOk() (*time.Time, bool)`

GetMongoURIUpdatedOk returns a tuple with the MongoURIUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIUpdated

`func (o *LegacyAtlasCluster) SetMongoURIUpdated(v time.Time)`

SetMongoURIUpdated sets MongoURIUpdated field to given value.

### HasMongoURIUpdated

`func (o *LegacyAtlasCluster) HasMongoURIUpdated() bool`

HasMongoURIUpdated returns a boolean if a field has been set.
### GetMongoURIWithOptions

`func (o *LegacyAtlasCluster) GetMongoURIWithOptions() string`

GetMongoURIWithOptions returns the MongoURIWithOptions field if non-nil, zero value otherwise.

### GetMongoURIWithOptionsOk

`func (o *LegacyAtlasCluster) GetMongoURIWithOptionsOk() (*string, bool)`

GetMongoURIWithOptionsOk returns a tuple with the MongoURIWithOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIWithOptions

`func (o *LegacyAtlasCluster) SetMongoURIWithOptions(v string)`

SetMongoURIWithOptions sets MongoURIWithOptions field to given value.

### HasMongoURIWithOptions

`func (o *LegacyAtlasCluster) HasMongoURIWithOptions() bool`

HasMongoURIWithOptions returns a boolean if a field has been set.
### GetName

`func (o *LegacyAtlasCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyAtlasCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyAtlasCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegacyAtlasCluster) HasName() bool`

HasName returns a boolean if a field has been set.
### GetNumShards

`func (o *LegacyAtlasCluster) GetNumShards() int`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *LegacyAtlasCluster) GetNumShardsOk() (*int, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *LegacyAtlasCluster) SetNumShards(v int)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *LegacyAtlasCluster) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.
### GetPaused

`func (o *LegacyAtlasCluster) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *LegacyAtlasCluster) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *LegacyAtlasCluster) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *LegacyAtlasCluster) HasPaused() bool`

HasPaused returns a boolean if a field has been set.
### GetPitEnabled

`func (o *LegacyAtlasCluster) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *LegacyAtlasCluster) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *LegacyAtlasCluster) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *LegacyAtlasCluster) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.
### GetProviderBackupEnabled

`func (o *LegacyAtlasCluster) GetProviderBackupEnabled() bool`

GetProviderBackupEnabled returns the ProviderBackupEnabled field if non-nil, zero value otherwise.

### GetProviderBackupEnabledOk

`func (o *LegacyAtlasCluster) GetProviderBackupEnabledOk() (*bool, bool)`

GetProviderBackupEnabledOk returns a tuple with the ProviderBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBackupEnabled

`func (o *LegacyAtlasCluster) SetProviderBackupEnabled(v bool)`

SetProviderBackupEnabled sets ProviderBackupEnabled field to given value.

### HasProviderBackupEnabled

`func (o *LegacyAtlasCluster) HasProviderBackupEnabled() bool`

HasProviderBackupEnabled returns a boolean if a field has been set.
### GetProviderSettings

`func (o *LegacyAtlasCluster) GetProviderSettings() ClusterProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *LegacyAtlasCluster) GetProviderSettingsOk() (*ClusterProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *LegacyAtlasCluster) SetProviderSettings(v ClusterProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *LegacyAtlasCluster) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.
### GetReplicaSetScalingStrategy

`func (o *LegacyAtlasCluster) GetReplicaSetScalingStrategy() string`

GetReplicaSetScalingStrategy returns the ReplicaSetScalingStrategy field if non-nil, zero value otherwise.

### GetReplicaSetScalingStrategyOk

`func (o *LegacyAtlasCluster) GetReplicaSetScalingStrategyOk() (*string, bool)`

GetReplicaSetScalingStrategyOk returns a tuple with the ReplicaSetScalingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetScalingStrategy

`func (o *LegacyAtlasCluster) SetReplicaSetScalingStrategy(v string)`

SetReplicaSetScalingStrategy sets ReplicaSetScalingStrategy field to given value.

### HasReplicaSetScalingStrategy

`func (o *LegacyAtlasCluster) HasReplicaSetScalingStrategy() bool`

HasReplicaSetScalingStrategy returns a boolean if a field has been set.
### GetReplicationFactor

`func (o *LegacyAtlasCluster) GetReplicationFactor() int`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *LegacyAtlasCluster) GetReplicationFactorOk() (*int, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *LegacyAtlasCluster) SetReplicationFactor(v int)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *LegacyAtlasCluster) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.
### GetReplicationSpec

`func (o *LegacyAtlasCluster) GetReplicationSpec() map[string]RegionSpec`

GetReplicationSpec returns the ReplicationSpec field if non-nil, zero value otherwise.

### GetReplicationSpecOk

`func (o *LegacyAtlasCluster) GetReplicationSpecOk() (*map[string]RegionSpec, bool)`

GetReplicationSpecOk returns a tuple with the ReplicationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpec

`func (o *LegacyAtlasCluster) SetReplicationSpec(v map[string]RegionSpec)`

SetReplicationSpec sets ReplicationSpec field to given value.

### HasReplicationSpec

`func (o *LegacyAtlasCluster) HasReplicationSpec() bool`

HasReplicationSpec returns a boolean if a field has been set.
### GetReplicationSpecs

`func (o *LegacyAtlasCluster) GetReplicationSpecs() []LegacyReplicationSpec`

GetReplicationSpecs returns the ReplicationSpecs field if non-nil, zero value otherwise.

### GetReplicationSpecsOk

`func (o *LegacyAtlasCluster) GetReplicationSpecsOk() (*[]LegacyReplicationSpec, bool)`

GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecs

`func (o *LegacyAtlasCluster) SetReplicationSpecs(v []LegacyReplicationSpec)`

SetReplicationSpecs sets ReplicationSpecs field to given value.

### HasReplicationSpecs

`func (o *LegacyAtlasCluster) HasReplicationSpecs() bool`

HasReplicationSpecs returns a boolean if a field has been set.
### GetRootCertType

`func (o *LegacyAtlasCluster) GetRootCertType() string`

GetRootCertType returns the RootCertType field if non-nil, zero value otherwise.

### GetRootCertTypeOk

`func (o *LegacyAtlasCluster) GetRootCertTypeOk() (*string, bool)`

GetRootCertTypeOk returns a tuple with the RootCertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCertType

`func (o *LegacyAtlasCluster) SetRootCertType(v string)`

SetRootCertType sets RootCertType field to given value.

### HasRootCertType

`func (o *LegacyAtlasCluster) HasRootCertType() bool`

HasRootCertType returns a boolean if a field has been set.
### GetSrvAddress

`func (o *LegacyAtlasCluster) GetSrvAddress() string`

GetSrvAddress returns the SrvAddress field if non-nil, zero value otherwise.

### GetSrvAddressOk

`func (o *LegacyAtlasCluster) GetSrvAddressOk() (*string, bool)`

GetSrvAddressOk returns a tuple with the SrvAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvAddress

`func (o *LegacyAtlasCluster) SetSrvAddress(v string)`

SetSrvAddress sets SrvAddress field to given value.

### HasSrvAddress

`func (o *LegacyAtlasCluster) HasSrvAddress() bool`

HasSrvAddress returns a boolean if a field has been set.
### GetStateName

`func (o *LegacyAtlasCluster) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *LegacyAtlasCluster) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *LegacyAtlasCluster) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *LegacyAtlasCluster) HasStateName() bool`

HasStateName returns a boolean if a field has been set.
### GetTags

`func (o *LegacyAtlasCluster) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LegacyAtlasCluster) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LegacyAtlasCluster) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LegacyAtlasCluster) HasTags() bool`

HasTags returns a boolean if a field has been set.
### GetTerminationProtectionEnabled

`func (o *LegacyAtlasCluster) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *LegacyAtlasCluster) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *LegacyAtlasCluster) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *LegacyAtlasCluster) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.
### GetVersionReleaseSystem

`func (o *LegacyAtlasCluster) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *LegacyAtlasCluster) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *LegacyAtlasCluster) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *LegacyAtlasCluster) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


