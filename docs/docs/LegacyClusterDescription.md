# LegacyClusterDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**AutoScaling**](AutoScaling.md) |  | [optional] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] 
**BiConnector** | Pointer to [**BiConnector**](BiConnector.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Configuration of nodes that comprise the cluster. | [optional] 
**ConnectionStrings** | Pointer to [**ClusterDescriptionConnectionStrings**](ClusterDescriptionConnectionStrings.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this serverless instance. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**DiskSizeGB** | Pointer to **float64** | Storage capacity that the host&#39;s root volume possesses expressed in gigabytes. Increase this number to add capacity. MongoDB Cloud requires this parameter if you set **replicationSpecs**. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value.  The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier. | [optional] 
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of Encryption at Rest for the cluster. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster. | [optional] [readonly] 
**Labels** | Pointer to [**[]NDSLabel**](NDSLabel.md) | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn&#39;t display your labels. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBMajorVersion** | Pointer to **string** | Major MongoDB version of the cluster. MongoDB Cloud deploys the cluster with the latest stable release of the specified version. | [optional] [default to "6.0"]
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
**ReplicationFactor** | Pointer to **int** | Number of members that belong to the replica set. Each member retains a copy of your databases, providing high availability and data redundancy. Use **replicationSpecs** instead. | [optional] [default to 3]
**ReplicationSpec** | Pointer to [**map[string]RegionSpec**](RegionSpec.md) | Physical location where MongoDB Cloud provisions cluster nodes. | [optional] 
**ReplicationSpecs** | Pointer to [**[]LegacyReplicationSpec**](LegacyReplicationSpec.md) | List of settings that configure your cluster regions.  - For Global Clusters, each object in the array represents one zone where MongoDB Cloud deploys your clusters nodes. - For non-Global sharded clusters and replica sets, the single object represents where MongoDB Cloud deploys your clusters nodes. | [optional] 
**RootCertType** | Pointer to **string** | Root Certificate Authority that MongoDB Atlas clusters uses. MongoDB Cloud supports Internet Security Research Group. | [optional] [default to "ISRGROOTX1"]
**SrvAddress** | Pointer to **string** | Connection string that you can use to connect to the cluster. The &#x60;+srv&#x60; modifier forces the connection to use Transport Layer Security (TLS). The &#x60;mongoURI&#x60; parameter lists additional options. | [optional] [readonly] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of the cluster. | [optional] [readonly] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster. | [optional] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. If value is &#x60;CONTINUOUS&#x60;, you must not specify **mongoDBMajorVersion**. | [optional] [default to "LTS"]

## Methods

### NewLegacyClusterDescription

`func NewLegacyClusterDescription() *LegacyClusterDescription`

NewLegacyClusterDescription instantiates a new LegacyClusterDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyClusterDescriptionWithDefaults

`func NewLegacyClusterDescriptionWithDefaults() *LegacyClusterDescription`

NewLegacyClusterDescriptionWithDefaults instantiates a new LegacyClusterDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *LegacyClusterDescription) GetAutoScaling() AutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *LegacyClusterDescription) GetAutoScalingOk() (*AutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *LegacyClusterDescription) SetAutoScaling(v AutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *LegacyClusterDescription) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *LegacyClusterDescription) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *LegacyClusterDescription) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *LegacyClusterDescription) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *LegacyClusterDescription) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetBiConnector

`func (o *LegacyClusterDescription) GetBiConnector() BiConnector`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *LegacyClusterDescription) GetBiConnectorOk() (*BiConnector, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *LegacyClusterDescription) SetBiConnector(v BiConnector)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *LegacyClusterDescription) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.

### GetClusterType

`func (o *LegacyClusterDescription) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *LegacyClusterDescription) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *LegacyClusterDescription) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *LegacyClusterDescription) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetConnectionStrings

`func (o *LegacyClusterDescription) GetConnectionStrings() ClusterDescriptionConnectionStrings`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *LegacyClusterDescription) GetConnectionStringsOk() (*ClusterDescriptionConnectionStrings, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *LegacyClusterDescription) SetConnectionStrings(v ClusterDescriptionConnectionStrings)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *LegacyClusterDescription) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetCreateDate

`func (o *LegacyClusterDescription) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *LegacyClusterDescription) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *LegacyClusterDescription) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *LegacyClusterDescription) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *LegacyClusterDescription) GetDiskSizeGB() float64`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *LegacyClusterDescription) GetDiskSizeGBOk() (*float64, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *LegacyClusterDescription) SetDiskSizeGB(v float64)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *LegacyClusterDescription) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetEncryptionAtRestProvider

`func (o *LegacyClusterDescription) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *LegacyClusterDescription) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *LegacyClusterDescription) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *LegacyClusterDescription) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.

### GetGroupId

`func (o *LegacyClusterDescription) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LegacyClusterDescription) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LegacyClusterDescription) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LegacyClusterDescription) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *LegacyClusterDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyClusterDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyClusterDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyClusterDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *LegacyClusterDescription) GetLabels() []NDSLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LegacyClusterDescription) GetLabelsOk() (*[]NDSLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LegacyClusterDescription) SetLabels(v []NDSLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LegacyClusterDescription) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *LegacyClusterDescription) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LegacyClusterDescription) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LegacyClusterDescription) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LegacyClusterDescription) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMongoDBMajorVersion

`func (o *LegacyClusterDescription) GetMongoDBMajorVersion() string`

GetMongoDBMajorVersion returns the MongoDBMajorVersion field if non-nil, zero value otherwise.

### GetMongoDBMajorVersionOk

`func (o *LegacyClusterDescription) GetMongoDBMajorVersionOk() (*string, bool)`

GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBMajorVersion

`func (o *LegacyClusterDescription) SetMongoDBMajorVersion(v string)`

SetMongoDBMajorVersion sets MongoDBMajorVersion field to given value.

### HasMongoDBMajorVersion

`func (o *LegacyClusterDescription) HasMongoDBMajorVersion() bool`

HasMongoDBMajorVersion returns a boolean if a field has been set.

### GetMongoDBVersion

`func (o *LegacyClusterDescription) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *LegacyClusterDescription) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *LegacyClusterDescription) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *LegacyClusterDescription) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetMongoURI

`func (o *LegacyClusterDescription) GetMongoURI() string`

GetMongoURI returns the MongoURI field if non-nil, zero value otherwise.

### GetMongoURIOk

`func (o *LegacyClusterDescription) GetMongoURIOk() (*string, bool)`

GetMongoURIOk returns a tuple with the MongoURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURI

`func (o *LegacyClusterDescription) SetMongoURI(v string)`

SetMongoURI sets MongoURI field to given value.

### HasMongoURI

`func (o *LegacyClusterDescription) HasMongoURI() bool`

HasMongoURI returns a boolean if a field has been set.

### GetMongoURIUpdated

`func (o *LegacyClusterDescription) GetMongoURIUpdated() time.Time`

GetMongoURIUpdated returns the MongoURIUpdated field if non-nil, zero value otherwise.

### GetMongoURIUpdatedOk

`func (o *LegacyClusterDescription) GetMongoURIUpdatedOk() (*time.Time, bool)`

GetMongoURIUpdatedOk returns a tuple with the MongoURIUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIUpdated

`func (o *LegacyClusterDescription) SetMongoURIUpdated(v time.Time)`

SetMongoURIUpdated sets MongoURIUpdated field to given value.

### HasMongoURIUpdated

`func (o *LegacyClusterDescription) HasMongoURIUpdated() bool`

HasMongoURIUpdated returns a boolean if a field has been set.

### GetMongoURIWithOptions

`func (o *LegacyClusterDescription) GetMongoURIWithOptions() string`

GetMongoURIWithOptions returns the MongoURIWithOptions field if non-nil, zero value otherwise.

### GetMongoURIWithOptionsOk

`func (o *LegacyClusterDescription) GetMongoURIWithOptionsOk() (*string, bool)`

GetMongoURIWithOptionsOk returns a tuple with the MongoURIWithOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIWithOptions

`func (o *LegacyClusterDescription) SetMongoURIWithOptions(v string)`

SetMongoURIWithOptions sets MongoURIWithOptions field to given value.

### HasMongoURIWithOptions

`func (o *LegacyClusterDescription) HasMongoURIWithOptions() bool`

HasMongoURIWithOptions returns a boolean if a field has been set.

### GetName

`func (o *LegacyClusterDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyClusterDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyClusterDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegacyClusterDescription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumShards

`func (o *LegacyClusterDescription) GetNumShards() int`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *LegacyClusterDescription) GetNumShardsOk() (*int, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *LegacyClusterDescription) SetNumShards(v int)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *LegacyClusterDescription) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.

### GetPaused

`func (o *LegacyClusterDescription) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *LegacyClusterDescription) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *LegacyClusterDescription) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *LegacyClusterDescription) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPitEnabled

`func (o *LegacyClusterDescription) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *LegacyClusterDescription) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *LegacyClusterDescription) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *LegacyClusterDescription) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.

### GetProviderBackupEnabled

`func (o *LegacyClusterDescription) GetProviderBackupEnabled() bool`

GetProviderBackupEnabled returns the ProviderBackupEnabled field if non-nil, zero value otherwise.

### GetProviderBackupEnabledOk

`func (o *LegacyClusterDescription) GetProviderBackupEnabledOk() (*bool, bool)`

GetProviderBackupEnabledOk returns a tuple with the ProviderBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBackupEnabled

`func (o *LegacyClusterDescription) SetProviderBackupEnabled(v bool)`

SetProviderBackupEnabled sets ProviderBackupEnabled field to given value.

### HasProviderBackupEnabled

`func (o *LegacyClusterDescription) HasProviderBackupEnabled() bool`

HasProviderBackupEnabled returns a boolean if a field has been set.

### GetProviderSettings

`func (o *LegacyClusterDescription) GetProviderSettings() ClusterProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *LegacyClusterDescription) GetProviderSettingsOk() (*ClusterProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *LegacyClusterDescription) SetProviderSettings(v ClusterProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *LegacyClusterDescription) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *LegacyClusterDescription) GetReplicationFactor() int`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *LegacyClusterDescription) GetReplicationFactorOk() (*int, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *LegacyClusterDescription) SetReplicationFactor(v int)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *LegacyClusterDescription) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetReplicationSpec

`func (o *LegacyClusterDescription) GetReplicationSpec() map[string]RegionSpec`

GetReplicationSpec returns the ReplicationSpec field if non-nil, zero value otherwise.

### GetReplicationSpecOk

`func (o *LegacyClusterDescription) GetReplicationSpecOk() (*map[string]RegionSpec, bool)`

GetReplicationSpecOk returns a tuple with the ReplicationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpec

`func (o *LegacyClusterDescription) SetReplicationSpec(v map[string]RegionSpec)`

SetReplicationSpec sets ReplicationSpec field to given value.

### HasReplicationSpec

`func (o *LegacyClusterDescription) HasReplicationSpec() bool`

HasReplicationSpec returns a boolean if a field has been set.

### GetReplicationSpecs

`func (o *LegacyClusterDescription) GetReplicationSpecs() []LegacyReplicationSpec`

GetReplicationSpecs returns the ReplicationSpecs field if non-nil, zero value otherwise.

### GetReplicationSpecsOk

`func (o *LegacyClusterDescription) GetReplicationSpecsOk() (*[]LegacyReplicationSpec, bool)`

GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecs

`func (o *LegacyClusterDescription) SetReplicationSpecs(v []LegacyReplicationSpec)`

SetReplicationSpecs sets ReplicationSpecs field to given value.

### HasReplicationSpecs

`func (o *LegacyClusterDescription) HasReplicationSpecs() bool`

HasReplicationSpecs returns a boolean if a field has been set.

### GetRootCertType

`func (o *LegacyClusterDescription) GetRootCertType() string`

GetRootCertType returns the RootCertType field if non-nil, zero value otherwise.

### GetRootCertTypeOk

`func (o *LegacyClusterDescription) GetRootCertTypeOk() (*string, bool)`

GetRootCertTypeOk returns a tuple with the RootCertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCertType

`func (o *LegacyClusterDescription) SetRootCertType(v string)`

SetRootCertType sets RootCertType field to given value.

### HasRootCertType

`func (o *LegacyClusterDescription) HasRootCertType() bool`

HasRootCertType returns a boolean if a field has been set.

### GetSrvAddress

`func (o *LegacyClusterDescription) GetSrvAddress() string`

GetSrvAddress returns the SrvAddress field if non-nil, zero value otherwise.

### GetSrvAddressOk

`func (o *LegacyClusterDescription) GetSrvAddressOk() (*string, bool)`

GetSrvAddressOk returns a tuple with the SrvAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvAddress

`func (o *LegacyClusterDescription) SetSrvAddress(v string)`

SetSrvAddress sets SrvAddress field to given value.

### HasSrvAddress

`func (o *LegacyClusterDescription) HasSrvAddress() bool`

HasSrvAddress returns a boolean if a field has been set.

### GetStateName

`func (o *LegacyClusterDescription) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *LegacyClusterDescription) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *LegacyClusterDescription) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *LegacyClusterDescription) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### GetTags

`func (o *LegacyClusterDescription) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LegacyClusterDescription) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LegacyClusterDescription) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LegacyClusterDescription) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationProtectionEnabled

`func (o *LegacyClusterDescription) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *LegacyClusterDescription) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *LegacyClusterDescription) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *LegacyClusterDescription) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.

### GetVersionReleaseSystem

`func (o *LegacyClusterDescription) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *LegacyClusterDescription) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *LegacyClusterDescription) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *LegacyClusterDescription) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


