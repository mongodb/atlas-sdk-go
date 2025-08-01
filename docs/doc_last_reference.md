# SDK Reference 

## Reference Documentation for SDK Endpoints

All URIs are relative to *https://cloud.mongodb.com*

Class        | Method        | HTTP request  | Description   | 
------------ | ------------- | ------------- | ------------- |
*AWSClustersDNSApi* | [GetAwsCustomDns](./docs/AWSClustersDNSApi.md#getawscustomdns) | **Get** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Return One Custom DNS Configuration for Atlas Clusters on AWS |
*AWSClustersDNSApi* | [ToggleAwsCustomDns](./docs/AWSClustersDNSApi.md#toggleawscustomdns) | **Patch** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Update State of One Custom DNS Configuration for Atlas Clusters on AWS |
*AccessTrackingApi* | [ListAccessLogsByClusterName](./docs/AccessTrackingApi.md#listaccesslogsbyclustername) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/clusters/{clusterName} | Return Database Access History for One Cluster by Cluster Name |
*AccessTrackingApi* | [ListAccessLogsByHostname](./docs/AccessTrackingApi.md#listaccesslogsbyhostname) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/processes/{hostname} | Return Database Access History for One Cluster by Hostname |
*AlertConfigurationsApi* | [CreateAlertConfiguration](./docs/AlertConfigurationsApi.md#createalertconfiguration) | **Post** /api/atlas/v2/groups/{groupId}/alertConfigs | Create One Alert Configuration in One Project |
*AlertConfigurationsApi* | [DeleteAlertConfiguration](./docs/AlertConfigurationsApi.md#deletealertconfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Remove One Alert Configuration from One Project |
*AlertConfigurationsApi* | [GetAlertConfiguration](./docs/AlertConfigurationsApi.md#getalertconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Return One Alert Configuration from One Project |
*AlertConfigurationsApi* | [ListAlertConfigurationMatchersFieldNames](./docs/AlertConfigurationsApi.md#listalertconfigurationmatchersfieldnames) | **Get** /api/atlas/v2/alertConfigs/matchers/fieldNames | Return All Alert Configuration Matchers Field Names |
*AlertConfigurationsApi* | [ListAlertConfigurations](./docs/AlertConfigurationsApi.md#listalertconfigurations) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs | Return All Alert Configurations in One Project |
*AlertConfigurationsApi* | [ListAlertConfigurationsByAlertId](./docs/AlertConfigurationsApi.md#listalertconfigurationsbyalertid) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId}/alertConfigs | Return All Alert Configurations Set for One Alert |
*AlertConfigurationsApi* | [ToggleAlertConfiguration](./docs/AlertConfigurationsApi.md#togglealertconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Toggle State of One Alert Configuration in One Project |
*AlertConfigurationsApi* | [UpdateAlertConfiguration](./docs/AlertConfigurationsApi.md#updatealertconfiguration) | **Put** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Update One Alert Configuration in One Project |
*AlertsApi* | [AcknowledgeAlert](./docs/AlertsApi.md#acknowledgealert) | **Patch** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Acknowledge One Alert from One Project |
*AlertsApi* | [GetAlert](./docs/AlertsApi.md#getalert) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Return One Alert from One Project |
*AlertsApi* | [ListAlerts](./docs/AlertsApi.md#listalerts) | **Get** /api/atlas/v2/groups/{groupId}/alerts | Return All Alerts from One Project |
*AlertsApi* | [ListAlertsByAlertConfigurationId](./docs/AlertsApi.md#listalertsbyalertconfigurationid) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}/alerts | Return All Open Alerts for One Alert Configuration |
*AtlasSearchApi* | [CreateAtlasSearchDeployment](./docs/AtlasSearchApi.md#createatlassearchdeployment) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Create Search Nodes |
*AtlasSearchApi* | [CreateAtlasSearchIndex](./docs/AtlasSearchApi.md#createatlassearchindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Create One Atlas Search Index |
*AtlasSearchApi* | [CreateAtlasSearchIndexDeprecated](./docs/AtlasSearchApi.md#createatlassearchindexdeprecated) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index |
*AtlasSearchApi* | [DeleteAtlasSearchDeployment](./docs/AtlasSearchApi.md#deleteatlassearchdeployment) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Delete Search Nodes |
*AtlasSearchApi* | [DeleteAtlasSearchIndex](./docs/AtlasSearchApi.md#deleteatlassearchindex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Remove One Atlas Search Index by ID |
*AtlasSearchApi* | [DeleteAtlasSearchIndexByName](./docs/AtlasSearchApi.md#deleteatlassearchindexbyname) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Remove One Atlas Search Index by Name |
*AtlasSearchApi* | [DeleteAtlasSearchIndexDeprecated](./docs/AtlasSearchApi.md#deleteatlassearchindexdeprecated) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index |
*AtlasSearchApi* | [GetAtlasSearchDeployment](./docs/AtlasSearchApi.md#getatlassearchdeployment) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Return All Search Nodes |
*AtlasSearchApi* | [GetAtlasSearchIndex](./docs/AtlasSearchApi.md#getatlassearchindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Return One Atlas Search Index by ID |
*AtlasSearchApi* | [GetAtlasSearchIndexByName](./docs/AtlasSearchApi.md#getatlassearchindexbyname) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Return One Atlas Search Index by Name |
*AtlasSearchApi* | [GetAtlasSearchIndexDeprecated](./docs/AtlasSearchApi.md#getatlassearchindexdeprecated) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index |
*AtlasSearchApi* | [ListAtlasSearchIndexes](./docs/AtlasSearchApi.md#listatlassearchindexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection |
*AtlasSearchApi* | [ListAtlasSearchIndexesCluster](./docs/AtlasSearchApi.md#listatlassearchindexescluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Return All Atlas Search Indexes for One Cluster |
*AtlasSearchApi* | [ListAtlasSearchIndexesDeprecated](./docs/AtlasSearchApi.md#listatlassearchindexesdeprecated) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection |
*AtlasSearchApi* | [UpdateAtlasSearchDeployment](./docs/AtlasSearchApi.md#updateatlassearchdeployment) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Update Search Nodes |
*AtlasSearchApi* | [UpdateAtlasSearchIndex](./docs/AtlasSearchApi.md#updateatlassearchindex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Update One Atlas Search Index by ID |
*AtlasSearchApi* | [UpdateAtlasSearchIndexByName](./docs/AtlasSearchApi.md#updateatlassearchindexbyname) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Update One Atlas Search Index by Name |
*AtlasSearchApi* | [UpdateAtlasSearchIndexDeprecated](./docs/AtlasSearchApi.md#updateatlassearchindexdeprecated) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index |
*AuditingApi* | [GetAuditingConfiguration](./docs/AuditingApi.md#getauditingconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/auditLog | Return Auditing Configuration for One Project |
*AuditingApi* | [UpdateAuditingConfiguration](./docs/AuditingApi.md#updateauditingconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/auditLog | Update Auditing Configuration for One Project |
*CloudBackupsApi* | [CancelBackupRestoreJob](./docs/CloudBackupsApi.md#cancelbackuprestorejob) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job for One Cluster |
*CloudBackupsApi* | [CreateBackupExportJob](./docs/CloudBackupsApi.md#createbackupexportjob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Snapshot Export Job |
*CloudBackupsApi* | [CreateBackupRestoreJob](./docs/CloudBackupsApi.md#createbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Cluster |
*CloudBackupsApi* | [CreateExportBucket](./docs/CloudBackupsApi.md#createexportbucket) | **Post** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Create One Snapshot Export Bucket |
*CloudBackupsApi* | [CreateServerlessBackupRestoreJob](./docs/CloudBackupsApi.md#createserverlessbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Serverless Instance |
*CloudBackupsApi* | [DeleteAllBackupSchedules](./docs/CloudBackupsApi.md#deleteallbackupschedules) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules |
*CloudBackupsApi* | [DeleteExportBucket](./docs/CloudBackupsApi.md#deleteexportbucket) | **Delete** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Delete One Snapshot Export Bucket |
*CloudBackupsApi* | [DeleteReplicaSetBackup](./docs/CloudBackupsApi.md#deletereplicasetbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup |
*CloudBackupsApi* | [DeleteShardedClusterBackup](./docs/CloudBackupsApi.md#deleteshardedclusterbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup |
*CloudBackupsApi* | [DisableDataProtectionSettings](./docs/CloudBackupsApi.md#disabledataprotectionsettings) | **Delete** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Disable Backup Compliance Policy Settings |
*CloudBackupsApi* | [GetBackupExportJob](./docs/CloudBackupsApi.md#getbackupexportjob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Snapshot Export Job |
*CloudBackupsApi* | [GetBackupRestoreJob](./docs/CloudBackupsApi.md#getbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Cluster |
*CloudBackupsApi* | [GetBackupSchedule](./docs/CloudBackupsApi.md#getbackupschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule |
*CloudBackupsApi* | [GetDataProtectionSettings](./docs/CloudBackupsApi.md#getdataprotectionsettings) | **Get** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Return Backup Compliance Policy Settings |
*CloudBackupsApi* | [GetExportBucket](./docs/CloudBackupsApi.md#getexportbucket) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One Snapshot Export Bucket |
*CloudBackupsApi* | [GetReplicaSetBackup](./docs/CloudBackupsApi.md#getreplicasetbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup |
*CloudBackupsApi* | [GetServerlessBackup](./docs/CloudBackupsApi.md#getserverlessbackup) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance |
*CloudBackupsApi* | [GetServerlessBackupRestoreJob](./docs/CloudBackupsApi.md#getserverlessbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance |
*CloudBackupsApi* | [GetShardedClusterBackup](./docs/CloudBackupsApi.md#getshardedclusterbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup |
*CloudBackupsApi* | [ListBackupExportJobs](./docs/CloudBackupsApi.md#listbackupexportjobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Snapshot Export Jobs |
*CloudBackupsApi* | [ListBackupRestoreJobs](./docs/CloudBackupsApi.md#listbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster |
*CloudBackupsApi* | [ListExportBuckets](./docs/CloudBackupsApi.md#listexportbuckets) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Return All Snapshot Export Buckets |
*CloudBackupsApi* | [ListReplicaSetBackups](./docs/CloudBackupsApi.md#listreplicasetbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups |
*CloudBackupsApi* | [ListServerlessBackupRestoreJobs](./docs/CloudBackupsApi.md#listserverlessbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance |
*CloudBackupsApi* | [ListServerlessBackups](./docs/CloudBackupsApi.md#listserverlessbackups) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance |
*CloudBackupsApi* | [ListShardedClusterBackups](./docs/CloudBackupsApi.md#listshardedclusterbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups |
*CloudBackupsApi* | [TakeSnapshot](./docs/CloudBackupsApi.md#takesnapshot) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot |
*CloudBackupsApi* | [UpdateBackupSchedule](./docs/CloudBackupsApi.md#updatebackupschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster |
*CloudBackupsApi* | [UpdateDataProtectionSettings](./docs/CloudBackupsApi.md#updatedataprotectionsettings) | **Put** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Update Backup Compliance Policy Settings |
*CloudBackupsApi* | [UpdateSnapshotRetention](./docs/CloudBackupsApi.md#updatesnapshotretention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Update Expiration Date for One Cloud Backup |
*CloudMigrationServiceApi* | [CreateLinkToken](./docs/CloudMigrationServiceApi.md#createlinktoken) | **Post** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Create One Link-Token |
*CloudMigrationServiceApi* | [CreatePushMigration](./docs/CloudMigrationServiceApi.md#createpushmigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations | Migrate One Local Managed Cluster to MongoDB Atlas |
*CloudMigrationServiceApi* | [CutoverMigration](./docs/CloudMigrationServiceApi.md#cutovermigration) | **Put** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover | Cut Over One Migrated Cluster |
*CloudMigrationServiceApi* | [DeleteLinkToken](./docs/CloudMigrationServiceApi.md#deletelinktoken) | **Delete** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Remove One Link-Token |
*CloudMigrationServiceApi* | [GetPushMigration](./docs/CloudMigrationServiceApi.md#getpushmigration) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId} | Return One Migration Job |
*CloudMigrationServiceApi* | [GetValidationStatus](./docs/CloudMigrationServiceApi.md#getvalidationstatus) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/validate/{validationId} | Return One Migration Validation Job |
*CloudMigrationServiceApi* | [ListSourceProjects](./docs/CloudMigrationServiceApi.md#listsourceprojects) | **Get** /api/atlas/v2/orgs/{orgId}/liveMigrations/availableProjects | Return All Projects Available for Migration |
*CloudMigrationServiceApi* | [ValidateMigration](./docs/CloudMigrationServiceApi.md#validatemigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations/validate | Validate One Migration Request |
*CloudProviderAccessApi* | [AuthorizeCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#authorizecloudprovideraccessrole) | **Patch** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role |
*CloudProviderAccessApi* | [CreateCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#createcloudprovideraccessrole) | **Post** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role |
*CloudProviderAccessApi* | [DeauthorizeCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#deauthorizecloudprovideraccessrole) | **Delete** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role |
*CloudProviderAccessApi* | [GetCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#getcloudprovideraccessrole) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Return One Cloud Provider Access Role |
*CloudProviderAccessApi* | [ListCloudProviderAccessRoles](./docs/CloudProviderAccessApi.md#listcloudprovideraccessroles) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles |
*ClusterOutageSimulationApi* | [EndOutageSimulation](./docs/ClusterOutageSimulationApi.md#endoutagesimulation) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | End One Outage Simulation |
*ClusterOutageSimulationApi* | [GetOutageSimulation](./docs/ClusterOutageSimulationApi.md#getoutagesimulation) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Return One Outage Simulation |
*ClusterOutageSimulationApi* | [StartOutageSimulation](./docs/ClusterOutageSimulationApi.md#startoutagesimulation) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Start One Outage Simulation |
*ClustersApi* | [AutoScalingConfiguration](./docs/ClustersApi.md#autoscalingconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/autoScalingConfiguration | Return Auto Scaling Configuration for One Sharded Cluster |
*ClustersApi* | [CreateCluster](./docs/ClustersApi.md#createcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters | Create One Cluster in One Project |
*ClustersApi* | [DeleteCluster](./docs/ClustersApi.md#deletecluster) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Remove One Cluster from One Project |
*ClustersApi* | [GetCluster](./docs/ClustersApi.md#getcluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Return One Cluster from One Project |
*ClustersApi* | [GetClusterAdvancedConfiguration](./docs/ClustersApi.md#getclusteradvancedconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Return Advanced Configuration Options for One Cluster |
*ClustersApi* | [GetClusterStatus](./docs/ClustersApi.md#getclusterstatus) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status | Return Status of All Cluster Operations |
*ClustersApi* | [GetSampleDatasetLoadStatus](./docs/ClustersApi.md#getsampledatasetloadstatus) | **Get** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId} | Return Status of Sample Dataset Load for One Cluster |
*ClustersApi* | [GrantMongoDbEmployeeAccess](./docs/ClustersApi.md#grantmongodbemployeeaccess) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:grantMongoDBEmployeeAccess | Grant MongoDB Employee Cluster Access for One Cluster |
*ClustersApi* | [ListCloudProviderRegions](./docs/ClustersApi.md#listcloudproviderregions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/provider/regions | Return All Cloud Provider Regions |
*ClustersApi* | [ListClusters](./docs/ClustersApi.md#listclusters) | **Get** /api/atlas/v2/groups/{groupId}/clusters | Return All Clusters in One Project |
*ClustersApi* | [ListClustersForAllProjects](./docs/ClustersApi.md#listclustersforallprojects) | **Get** /api/atlas/v2/clusters | Return All Authorized Clusters in All Projects |
*ClustersApi* | [LoadSampleDataset](./docs/ClustersApi.md#loadsampledataset) | **Post** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name} | Load Sample Dataset into One Cluster |
*ClustersApi* | [PinFeatureCompatibilityVersion](./docs/ClustersApi.md#pinfeaturecompatibilityversion) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:pinFeatureCompatibilityVersion | Pin Feature Compatibility Version for One Cluster in One Project |
*ClustersApi* | [RevokeMongoDbEmployeeAccess](./docs/ClustersApi.md#revokemongodbemployeeaccess) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:revokeMongoDBEmployeeAccess | Revoke MongoDB Employee Cluster Access for One Cluster |
*ClustersApi* | [TestFailover](./docs/ClustersApi.md#testfailover) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restartPrimaries | Test Failover for One Cluster |
*ClustersApi* | [UnpinFeatureCompatibilityVersion](./docs/ClustersApi.md#unpinfeaturecompatibilityversion) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:unpinFeatureCompatibilityVersion | Unpin Feature Compatibility Version for One Cluster in One Project |
*ClustersApi* | [UpdateCluster](./docs/ClustersApi.md#updatecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Update One Cluster in One Project |
*ClustersApi* | [UpdateClusterAdvancedConfiguration](./docs/ClustersApi.md#updateclusteradvancedconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Update Advanced Configuration Options for One Cluster |
*ClustersApi* | [UpgradeSharedCluster](./docs/ClustersApi.md#upgradesharedcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade | Upgrade One Shared-Tier Cluster |
*ClustersApi* | [UpgradeSharedClusterToServerless](./docs/ClustersApi.md#upgradesharedclustertoserverless) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgradeToServerless | Upgrade One Shared-Tier Cluster to One Serverless Instance |
*CollectionLevelMetricsApi* | [GetCollStatsLatencyNamespaceClusterMeasurements](./docs/CollectionLevelMetricsApi.md#getcollstatslatencynamespaceclustermeasurements) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements | Return Cluster-Level Query Latency |
*CollectionLevelMetricsApi* | [GetCollStatsLatencyNamespaceHostMeasurements](./docs/CollectionLevelMetricsApi.md#getcollstatslatencynamespacehostmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/{databaseName}/{collectionName}/collStats/measurements | Return Host-Level Query Latency |
*CollectionLevelMetricsApi* | [GetCollStatsLatencyNamespaceMetrics](./docs/CollectionLevelMetricsApi.md#getcollstatslatencynamespacemetrics) | **Get** /api/atlas/v2/groups/{groupId}/collStats/metrics | Return all metric names |
*CollectionLevelMetricsApi* | [GetCollStatsLatencyNamespacesForCluster](./docs/CollectionLevelMetricsApi.md#getcollstatslatencynamespacesforcluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/collStats/namespaces | Return Ranked Namespaces from a Cluster |
*CollectionLevelMetricsApi* | [GetCollStatsLatencyNamespacesForHost](./docs/CollectionLevelMetricsApi.md#getcollstatslatencynamespacesforhost) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/collStats/namespaces | Return Ranked Namespaces from a Host |
*CollectionLevelMetricsApi* | [GetPinnedNamespaces](./docs/CollectionLevelMetricsApi.md#getpinnednamespaces) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Return Pinned Namespaces |
*CollectionLevelMetricsApi* | [PinNamespacesPatch](./docs/CollectionLevelMetricsApi.md#pinnamespacespatch) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Add Pinned Namespaces |
*CollectionLevelMetricsApi* | [PinNamespacesPut](./docs/CollectionLevelMetricsApi.md#pinnamespacesput) | **Put** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Pin Namespaces |
*CollectionLevelMetricsApi* | [UnpinNamespaces](./docs/CollectionLevelMetricsApi.md#unpinnamespaces) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/unpin | Unpin namespaces |
*CustomDatabaseRolesApi* | [CreateCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#createcustomdatabaserole) | **Post** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Create One Custom Role |
*CustomDatabaseRolesApi* | [DeleteCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#deletecustomdatabaserole) | **Delete** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project |
*CustomDatabaseRolesApi* | [GetCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#getcustomdatabaserole) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project |
*CustomDatabaseRolesApi* | [ListCustomDatabaseRoles](./docs/CustomDatabaseRolesApi.md#listcustomdatabaseroles) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project |
*CustomDatabaseRolesApi* | [UpdateCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#updatecustomdatabaserole) | **Patch** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project |
*DataFederationApi* | [CreateDataFederationPrivateEndpoint](./docs/DataFederationApi.md#createdatafederationprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Create One Federated Database Instance and Online Archive Private Endpoint for One Project |
*DataFederationApi* | [CreateFederatedDatabase](./docs/DataFederationApi.md#createfederateddatabase) | **Post** /api/atlas/v2/groups/{groupId}/dataFederation | Create One Federated Database Instance in One Project |
*DataFederationApi* | [CreateOneDataFederationQueryLimit](./docs/DataFederationApi.md#createonedatafederationquerylimit) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Configure One Query Limit for One Federated Database Instance |
*DataFederationApi* | [DeleteDataFederationPrivateEndpoint](./docs/DataFederationApi.md#deletedatafederationprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Remove One Federated Database Instance and Online Archive Private Endpoint from One Project |
*DataFederationApi* | [DeleteFederatedDatabase](./docs/DataFederationApi.md#deletefederateddatabase) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Remove One Federated Database Instance from One Project |
*DataFederationApi* | [DeleteOneDataFederationInstanceQueryLimit](./docs/DataFederationApi.md#deleteonedatafederationinstancequerylimit) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Delete One Query Limit for One Federated Database Instance |
*DataFederationApi* | [DownloadFederatedDatabaseQueryLogs](./docs/DataFederationApi.md#downloadfederateddatabasequerylogs) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz | Download Query Logs for One Federated Database Instance |
*DataFederationApi* | [GetDataFederationPrivateEndpoint](./docs/DataFederationApi.md#getdatafederationprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Return One Federated Database Instance and Online Archive Private Endpoint in One Project |
*DataFederationApi* | [GetFederatedDatabase](./docs/DataFederationApi.md#getfederateddatabase) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Return One Federated Database Instance in One Project |
*DataFederationApi* | [ListDataFederationPrivateEndpoints](./docs/DataFederationApi.md#listdatafederationprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Return All Federated Database Instance and Online Archive Private Endpoints in One Project |
*DataFederationApi* | [ListFederatedDatabases](./docs/DataFederationApi.md#listfederateddatabases) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation | Return All Federated Database Instances in One Project |
*DataFederationApi* | [ReturnFederatedDatabaseQueryLimit](./docs/DataFederationApi.md#returnfederateddatabasequerylimit) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Return One Federated Database Instance Query Limit for One Project |
*DataFederationApi* | [ReturnFederatedDatabaseQueryLimits](./docs/DataFederationApi.md#returnfederateddatabasequerylimits) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits | Return All Query Limits for One Federated Database Instance |
*DataFederationApi* | [UpdateFederatedDatabase](./docs/DataFederationApi.md#updatefederateddatabase) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Update One Federated Database Instance in One Project |
*DataLakePipelinesApi* | [CreatePipeline](./docs/DataLakePipelinesApi.md#createpipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines | Create One Data Lake Pipeline |
*DataLakePipelinesApi* | [DeletePipeline](./docs/DataLakePipelinesApi.md#deletepipeline) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Remove One Data Lake Pipeline |
*DataLakePipelinesApi* | [DeletePipelineRunDataset](./docs/DataLakePipelinesApi.md#deletepipelinerundataset) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Delete One Pipeline Run Dataset |
*DataLakePipelinesApi* | [GetPipeline](./docs/DataLakePipelinesApi.md#getpipeline) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Return One Data Lake Pipeline |
*DataLakePipelinesApi* | [GetPipelineRun](./docs/DataLakePipelinesApi.md#getpipelinerun) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Return One Data Lake Pipeline Run |
*DataLakePipelinesApi* | [ListPipelineRuns](./docs/DataLakePipelinesApi.md#listpipelineruns) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs | Return All Data Lake Pipeline Runs in One Project |
*DataLakePipelinesApi* | [ListPipelineSchedules](./docs/DataLakePipelinesApi.md#listpipelineschedules) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSchedules | Return All Ingestion Schedules for One Data Lake Pipeline |
*DataLakePipelinesApi* | [ListPipelineSnapshots](./docs/DataLakePipelinesApi.md#listpipelinesnapshots) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots | Return All Backup Snapshots for One Data Lake Pipeline |
*DataLakePipelinesApi* | [ListPipelines](./docs/DataLakePipelinesApi.md#listpipelines) | **Get** /api/atlas/v2/groups/{groupId}/pipelines | Return All Data Lake Pipelines in One Project |
*DataLakePipelinesApi* | [PausePipeline](./docs/DataLakePipelinesApi.md#pausepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/pause | Pause One Data Lake Pipeline |
*DataLakePipelinesApi* | [ResumePipeline](./docs/DataLakePipelinesApi.md#resumepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/resume | Resume One Data Lake Pipeline |
*DataLakePipelinesApi* | [TriggerSnapshotIngestion](./docs/DataLakePipelinesApi.md#triggersnapshotingestion) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/trigger | Trigger On-Demand Snapshot Ingestion |
*DataLakePipelinesApi* | [UpdatePipeline](./docs/DataLakePipelinesApi.md#updatepipeline) | **Patch** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Update One Data Lake Pipeline |
*DatabaseUsersApi* | [CreateDatabaseUser](./docs/DatabaseUsersApi.md#createdatabaseuser) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers | Create One Database User in One Project |
*DatabaseUsersApi* | [DeleteDatabaseUser](./docs/DatabaseUsersApi.md#deletedatabaseuser) | **Delete** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Remove One Database User from One Project |
*DatabaseUsersApi* | [GetDatabaseUser](./docs/DatabaseUsersApi.md#getdatabaseuser) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Return One Database User from One Project |
*DatabaseUsersApi* | [ListDatabaseUsers](./docs/DatabaseUsersApi.md#listdatabaseusers) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers | Return All Database Users in One Project |
*DatabaseUsersApi* | [UpdateDatabaseUser](./docs/DatabaseUsersApi.md#updatedatabaseuser) | **Patch** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Update One Database User in One Project |
*EncryptionAtRestUsingCustomerKeyManagementApi* | [CreateEncryptionAtRestPrivateEndpoint](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#createencryptionatrestprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints | Create One Private Endpoint in a Specified Region for Encryption at Rest Using Customer Key Management |
*EncryptionAtRestUsingCustomerKeyManagementApi* | [GetEncryptionAtRest](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#getencryptionatrest) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project |
*EncryptionAtRestUsingCustomerKeyManagementApi* | [GetEncryptionAtRestPrivateEndpoint](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#getencryptionatrestprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints/{endpointId} | Return One Private Endpoint for Encryption at Rest Using Customer Key Management |
*EncryptionAtRestUsingCustomerKeyManagementApi* | [GetEncryptionAtRestPrivateEndpointsForCloudProvider](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#getencryptionatrestprivateendpointsforcloudprovider) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints | Return Private Endpoints of a Cloud Provider for Encryption at Rest Using Customer Key Management for One Project |
*EncryptionAtRestUsingCustomerKeyManagementApi* | [RequestEncryptionAtRestPrivateEndpointDeletion](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#requestencryptionatrestprivateendpointdeletion) | **Delete** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints/{endpointId} | Delete One Private Endpoint for Encryption at Rest Using Customer Key Management |
*EncryptionAtRestUsingCustomerKeyManagementApi* | [UpdateEncryptionAtRest](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#updateencryptionatrest) | **Patch** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Update Encryption at Rest Configuration for One Project |
*EventsApi* | [GetOrganizationEvent](./docs/EventsApi.md#getorganizationevent) | **Get** /api/atlas/v2/orgs/{orgId}/events/{eventId} | Return One Event from One Organization |
*EventsApi* | [GetProjectEvent](./docs/EventsApi.md#getprojectevent) | **Get** /api/atlas/v2/groups/{groupId}/events/{eventId} | Return One Event from One Project |
*EventsApi* | [ListEventTypes](./docs/EventsApi.md#listeventtypes) | **Get** /api/atlas/v2/eventTypes | Return All Event Types |
*EventsApi* | [ListOrganizationEvents](./docs/EventsApi.md#listorganizationevents) | **Get** /api/atlas/v2/orgs/{orgId}/events | Return Events from One Organization |
*EventsApi* | [ListProjectEvents](./docs/EventsApi.md#listprojectevents) | **Get** /api/atlas/v2/groups/{groupId}/events | Return Events from One Project |
*FederatedAuthenticationApi* | [CreateIdentityProvider](./docs/FederatedAuthenticationApi.md#createidentityprovider) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Create One Identity Provider |
*FederatedAuthenticationApi* | [CreateRoleMapping](./docs/FederatedAuthenticationApi.md#createrolemapping) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Create One Role Mapping in One Organization Configuration |
*FederatedAuthenticationApi* | [DeleteFederationApp](./docs/FederatedAuthenticationApi.md#deletefederationapp) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId} | Delete One Federation Settings Instance |
*FederatedAuthenticationApi* | [DeleteIdentityProvider](./docs/FederatedAuthenticationApi.md#deleteidentityprovider) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Delete One Identity Provider |
*FederatedAuthenticationApi* | [DeleteRoleMapping](./docs/FederatedAuthenticationApi.md#deleterolemapping) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization |
*FederatedAuthenticationApi* | [GetConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#getconnectedorgconfig) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Organization Configuration from One Federation |
*FederatedAuthenticationApi* | [GetFederationSettings](./docs/FederatedAuthenticationApi.md#getfederationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization |
*FederatedAuthenticationApi* | [GetIdentityProvider](./docs/FederatedAuthenticationApi.md#getidentityprovider) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return One Identity Provider by ID |
*FederatedAuthenticationApi* | [GetIdentityProviderMetadata](./docs/FederatedAuthenticationApi.md#getidentityprovidermetadata) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return Metadata of One Identity Provider |
*FederatedAuthenticationApi* | [GetRoleMapping](./docs/FederatedAuthenticationApi.md#getrolemapping) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization |
*FederatedAuthenticationApi* | [ListConnectedOrgConfigs](./docs/FederatedAuthenticationApi.md#listconnectedorgconfigs) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Organization Configurations from One Federation |
*FederatedAuthenticationApi* | [ListIdentityProviders](./docs/FederatedAuthenticationApi.md#listidentityproviders) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Return All Identity Providers in One Federation |
*FederatedAuthenticationApi* | [ListRoleMappings](./docs/FederatedAuthenticationApi.md#listrolemappings) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization |
*FederatedAuthenticationApi* | [RemoveConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#removeconnectedorgconfig) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Organization Configuration from One Federation |
*FederatedAuthenticationApi* | [RevokeJwksFromIdentityProvider](./docs/FederatedAuthenticationApi.md#revokejwksfromidentityprovider) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/jwks | Revoke JWKS from One OIDC Identity Provider |
*FederatedAuthenticationApi* | [UpdateConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#updateconnectedorgconfig) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update One Organization Configuration in One Federation |
*FederatedAuthenticationApi* | [UpdateIdentityProvider](./docs/FederatedAuthenticationApi.md#updateidentityprovider) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update One Identity Provider |
*FederatedAuthenticationApi* | [UpdateRoleMapping](./docs/FederatedAuthenticationApi.md#updaterolemapping) | **Put** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization |
*FlexClustersApi* | [CreateFlexCluster](./docs/FlexClustersApi.md#createflexcluster) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters | Create One Flex Cluster in One Project |
*FlexClustersApi* | [DeleteFlexCluster](./docs/FlexClustersApi.md#deleteflexcluster) | **Delete** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Remove One Flex Cluster from One Project |
*FlexClustersApi* | [GetFlexCluster](./docs/FlexClustersApi.md#getflexcluster) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Return One Flex Cluster from One Project |
*FlexClustersApi* | [ListFlexClusters](./docs/FlexClustersApi.md#listflexclusters) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters | Return All Flex Clusters from One Project |
*FlexClustersApi* | [UpdateFlexCluster](./docs/FlexClustersApi.md#updateflexcluster) | **Patch** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Update One Flex Cluster in One Project |
*FlexClustersApi* | [UpgradeFlexCluster](./docs/FlexClustersApi.md#upgradeflexcluster) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters:tenantUpgrade | Upgrade One Flex Cluster |
*FlexRestoreJobsApi* | [CreateFlexBackupRestoreJob](./docs/FlexRestoreJobsApi.md#createflexbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs | Restore One Snapshot of One Flex Cluster |
*FlexRestoreJobsApi* | [GetFlexBackupRestoreJob](./docs/FlexRestoreJobsApi.md#getflexbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Flex Cluster |
*FlexRestoreJobsApi* | [ListFlexBackupRestoreJobs](./docs/FlexRestoreJobsApi.md#listflexbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs | Return All Restore Jobs for One Flex Cluster |
*FlexSnapshotsApi* | [DownloadFlexBackup](./docs/FlexSnapshotsApi.md#downloadflexbackup) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/download | Download One Flex Cluster Snapshot |
*FlexSnapshotsApi* | [GetFlexBackup](./docs/FlexSnapshotsApi.md#getflexbackup) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/snapshots/{snapshotId} | Return One Snapshot of One Flex Cluster |
*FlexSnapshotsApi* | [ListFlexBackups](./docs/FlexSnapshotsApi.md#listflexbackups) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/snapshots | Return All Snapshots of One Flex Cluster |
*GlobalClustersApi* | [CreateCustomZoneMapping](./docs/GlobalClustersApi.md#createcustomzonemapping) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Add One Custom Zone Mapping to One Global Cluster |
*GlobalClustersApi* | [CreateManagedNamespace](./docs/GlobalClustersApi.md#createmanagednamespace) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Create One Managed Namespace in One Global Cluster |
*GlobalClustersApi* | [DeleteAllCustomZoneMappings](./docs/GlobalClustersApi.md#deleteallcustomzonemappings) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Cluster |
*GlobalClustersApi* | [DeleteManagedNamespace](./docs/GlobalClustersApi.md#deletemanagednamespace) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Cluster |
*GlobalClustersApi* | [GetManagedNamespace](./docs/GlobalClustersApi.md#getmanagednamespace) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites | Return One Managed Namespace in One Global Cluster |
*InvoicesApi* | [CreateCostExplorerQueryProcess](./docs/InvoicesApi.md#createcostexplorerqueryprocess) | **Post** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage | Create One Cost Explorer Query Process |
*InvoicesApi* | [DownloadInvoiceCsv](./docs/InvoicesApi.md#downloadinvoicecsv) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Invoice as CSV for One Organization |
*InvoicesApi* | [GetCostExplorerQueryProcess](./docs/InvoicesApi.md#getcostexplorerqueryprocess) | **Get** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage/{token} | Return results from a given Cost Explorer query, or notify that the results are not ready yet. |
*InvoicesApi* | [GetInvoice](./docs/InvoicesApi.md#getinvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId} | Return One Invoice for One Organization |
*InvoicesApi* | [ListInvoices](./docs/InvoicesApi.md#listinvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices | Return All Invoices for One Organization |
*InvoicesApi* | [ListPendingInvoices](./docs/InvoicesApi.md#listpendinginvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization |
*InvoicesApi* | [QueryLineItemsFromSingleInvoice](./docs/InvoicesApi.md#querylineitemsfromsingleinvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/lineItems:search | Return All Line Items for One Invoice by Invoice ID |
*LDAPConfigurationApi* | [DeleteLdapConfiguration](./docs/LDAPConfigurationApi.md#deleteldapconfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove LDAP User to DN Mapping |
*LDAPConfigurationApi* | [GetLdapConfiguration](./docs/LDAPConfigurationApi.md#getldapconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return LDAP or X.509 Configuration |
*LDAPConfigurationApi* | [GetLdapConfigurationStatus](./docs/LDAPConfigurationApi.md#getldapconfigurationstatus) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return Status of LDAP Configuration Verification in One Project |
*LDAPConfigurationApi* | [SaveLdapConfiguration](./docs/LDAPConfigurationApi.md#saveldapconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Update LDAP or X.509 Configuration |
*LDAPConfigurationApi* | [VerifyLdapConfiguration](./docs/LDAPConfigurationApi.md#verifyldapconfiguration) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify LDAP Configuration in One Project |
*LegacyBackupApi* | [CreateLegacyBackupRestoreJob](./docs/LegacyBackupApi.md#createlegacybackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Create One Legacy Backup Restore Job |
*LegacyBackupApi* | [DeleteLegacySnapshot](./docs/LegacyBackupApi.md#deletelegacysnapshot) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Remove One Legacy Backup Snapshot |
*LegacyBackupApi* | [GetLegacyBackupCheckpoint](./docs/LegacyBackupApi.md#getlegacybackupcheckpoint) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId} | Return One Legacy Backup Checkpoint |
*LegacyBackupApi* | [GetLegacyBackupRestoreJob](./docs/LegacyBackupApi.md#getlegacybackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId} | Return One Legacy Backup Restore Job |
*LegacyBackupApi* | [GetLegacySnapshot](./docs/LegacyBackupApi.md#getlegacysnapshot) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Return One Legacy Backup Snapshot |
*LegacyBackupApi* | [GetLegacySnapshotSchedule](./docs/LegacyBackupApi.md#getlegacysnapshotschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Return One Snapshot Schedule |
*LegacyBackupApi* | [ListLegacyBackupCheckpoints](./docs/LegacyBackupApi.md#listlegacybackupcheckpoints) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints | Return All Legacy Backup Checkpoints |
*LegacyBackupApi* | [ListLegacyBackupRestoreJobs](./docs/LegacyBackupApi.md#listlegacybackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Return All Legacy Backup Restore Jobs |
*LegacyBackupApi* | [ListLegacySnapshots](./docs/LegacyBackupApi.md#listlegacysnapshots) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots | Return All Legacy Backup Snapshots |
*LegacyBackupApi* | [UpdateLegacySnapshotRetention](./docs/LegacyBackupApi.md#updatelegacysnapshotretention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Update Expiration Date for One Legacy Backup Snapshot |
*LegacyBackupApi* | [UpdateLegacySnapshotSchedule](./docs/LegacyBackupApi.md#updatelegacysnapshotschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Update Snapshot Schedule for One Cluster |
*MaintenanceWindowsApi* | [DeferMaintenanceWindow](./docs/MaintenanceWindowsApi.md#defermaintenancewindow) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/defer | Defer One Maintenance Window for One Project |
*MaintenanceWindowsApi* | [GetMaintenanceWindow](./docs/MaintenanceWindowsApi.md#getmaintenancewindow) | **Get** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Return One Maintenance Window for One Project |
*MaintenanceWindowsApi* | [ResetMaintenanceWindow](./docs/MaintenanceWindowsApi.md#resetmaintenancewindow) | **Delete** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Reset One Maintenance Window for One Project |
*MaintenanceWindowsApi* | [ToggleMaintenanceAutoDefer](./docs/MaintenanceWindowsApi.md#togglemaintenanceautodefer) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/autoDefer | Toggle Automatic Deferral of Maintenance for One Project |
*MaintenanceWindowsApi* | [UpdateMaintenanceWindow](./docs/MaintenanceWindowsApi.md#updatemaintenancewindow) | **Patch** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Update One Maintenance Window for One Project |
*MongoDBCloudUsersApi* | [AddOrganizationRole](./docs/MongoDBCloudUsersApi.md#addorganizationrole) | **Post** /api/atlas/v2/orgs/{orgId}/users/{userId}:addRole | Add One Organization Role to One MongoDB Cloud User |
*MongoDBCloudUsersApi* | [AddProjectRole](./docs/MongoDBCloudUsersApi.md#addprojectrole) | **Post** /api/atlas/v2/groups/{groupId}/users/{userId}:addRole | Add One Project Role to One MongoDB Cloud User |
*MongoDBCloudUsersApi* | [AddProjectUser](./docs/MongoDBCloudUsersApi.md#addprojectuser) | **Post** /api/atlas/v2/groups/{groupId}/users | Add One MongoDB Cloud User to One Project |
*MongoDBCloudUsersApi* | [AddUserToTeam](./docs/MongoDBCloudUsersApi.md#addusertoteam) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}:addUser | Add One MongoDB Cloud User to One Team |
*MongoDBCloudUsersApi* | [CreateOrganizationUser](./docs/MongoDBCloudUsersApi.md#createorganizationuser) | **Post** /api/atlas/v2/orgs/{orgId}/users | Add One MongoDB Cloud User to One Organization |
*MongoDBCloudUsersApi* | [CreateUser](./docs/MongoDBCloudUsersApi.md#createuser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User |
*MongoDBCloudUsersApi* | [GetOrganizationUser](./docs/MongoDBCloudUsersApi.md#getorganizationuser) | **Get** /api/atlas/v2/orgs/{orgId}/users/{userId} | Return One MongoDB Cloud User in One Organization |
*MongoDBCloudUsersApi* | [GetProjectUser](./docs/MongoDBCloudUsersApi.md#getprojectuser) | **Get** /api/atlas/v2/groups/{groupId}/users/{userId} | Return One MongoDB Cloud User in One Project |
*MongoDBCloudUsersApi* | [GetUser](./docs/MongoDBCloudUsersApi.md#getuser) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User by ID |
*MongoDBCloudUsersApi* | [GetUserByUsername](./docs/MongoDBCloudUsersApi.md#getuserbyusername) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User by Username |
*MongoDBCloudUsersApi* | [ListOrganizationUsers](./docs/MongoDBCloudUsersApi.md#listorganizationusers) | **Get** /api/atlas/v2/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization |
*MongoDBCloudUsersApi* | [ListProjectUsers](./docs/MongoDBCloudUsersApi.md#listprojectusers) | **Get** /api/atlas/v2/groups/{groupId}/users | Return All MongoDB Cloud Users in One Project |
*MongoDBCloudUsersApi* | [ListTeamUsers](./docs/MongoDBCloudUsersApi.md#listteamusers) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team |
*MongoDBCloudUsersApi* | [RemoveOrganizationRole](./docs/MongoDBCloudUsersApi.md#removeorganizationrole) | **Post** /api/atlas/v2/orgs/{orgId}/users/{userId}:removeRole | Remove One Organization Role from One MongoDB Cloud User |
*MongoDBCloudUsersApi* | [RemoveOrganizationUser](./docs/MongoDBCloudUsersApi.md#removeorganizationuser) | **Delete** /api/atlas/v2/orgs/{orgId}/users/{userId} | Remove One MongoDB Cloud User from One Organization |
*MongoDBCloudUsersApi* | [RemoveProjectRole](./docs/MongoDBCloudUsersApi.md#removeprojectrole) | **Post** /api/atlas/v2/groups/{groupId}/users/{userId}:removeRole | Remove One Project Role from One MongoDB Cloud User |
*MongoDBCloudUsersApi* | [RemoveProjectUser](./docs/MongoDBCloudUsersApi.md#removeprojectuser) | **Delete** /api/atlas/v2/groups/{groupId}/users/{userId} | Remove One MongoDB Cloud User from One Project |
*MongoDBCloudUsersApi* | [RemoveUserFromTeam](./docs/MongoDBCloudUsersApi.md#removeuserfromteam) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}:removeUser | Remove One MongoDB Cloud User from One Team |
*MongoDBCloudUsersApi* | [UpdateOrganizationUser](./docs/MongoDBCloudUsersApi.md#updateorganizationuser) | **Patch** /api/atlas/v2/orgs/{orgId}/users/{userId} | Update One MongoDB Cloud User in One Organization |
*MonitoringAndLogsApi* | [GetAtlasProcess](./docs/MonitoringAndLogsApi.md#getatlasprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId} | Return One MongoDB Process by ID |
*MonitoringAndLogsApi* | [GetDatabase](./docs/MonitoringAndLogsApi.md#getdatabase) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName} | Return One Database for One MongoDB Process |
*MonitoringAndLogsApi* | [GetDatabaseMeasurements](./docs/MonitoringAndLogsApi.md#getdatabasemeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements | Return Measurements for One Database in One MongoDB Process |
*MonitoringAndLogsApi* | [GetDiskMeasurements](./docs/MonitoringAndLogsApi.md#getdiskmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements | Return Measurements of One Disk for One MongoDB Process |
*MonitoringAndLogsApi* | [GetHostLogs](./docs/MonitoringAndLogsApi.md#gethostlogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Cluster Host in One Project |
*MonitoringAndLogsApi* | [GetHostMeasurements](./docs/MonitoringAndLogsApi.md#gethostmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/measurements | Return Measurements for One MongoDB Process |
*MonitoringAndLogsApi* | [GetIndexMetrics](./docs/MonitoringAndLogsApi.md#getindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements | Return Atlas Search Metrics for One Index in One Namespace |
*MonitoringAndLogsApi* | [GetMeasurements](./docs/MonitoringAndLogsApi.md#getmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/measurements | Return Atlas Search Hardware and Status Metrics |
*MonitoringAndLogsApi* | [ListAtlasProcesses](./docs/MonitoringAndLogsApi.md#listatlasprocesses) | **Get** /api/atlas/v2/groups/{groupId}/processes | Return All MongoDB Processes in One Project |
*MonitoringAndLogsApi* | [ListDatabases](./docs/MonitoringAndLogsApi.md#listdatabases) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases | Return Available Databases for One MongoDB Process |
*MonitoringAndLogsApi* | [ListDiskMeasurements](./docs/MonitoringAndLogsApi.md#listdiskmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName} | Return Measurements for One Disk |
*MonitoringAndLogsApi* | [ListDiskPartitions](./docs/MonitoringAndLogsApi.md#listdiskpartitions) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks | Return Available Disks for One MongoDB Process |
*MonitoringAndLogsApi* | [ListIndexMetrics](./docs/MonitoringAndLogsApi.md#listindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements | Return All Atlas Search Index Metrics for One Namespace |
*MonitoringAndLogsApi* | [ListMetricTypes](./docs/MonitoringAndLogsApi.md#listmetrictypes) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics | Return All Atlas Search Metric Types for One Process |
*NetworkPeeringApi* | [CreatePeeringConnection](./docs/NetworkPeeringApi.md#createpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/peers | Create One Network Peering Connection |
*NetworkPeeringApi* | [CreatePeeringContainer](./docs/NetworkPeeringApi.md#createpeeringcontainer) | **Post** /api/atlas/v2/groups/{groupId}/containers | Create One Network Peering Container |
*NetworkPeeringApi* | [DeletePeeringConnection](./docs/NetworkPeeringApi.md#deletepeeringconnection) | **Delete** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Remove One Network Peering Connection |
*NetworkPeeringApi* | [DeletePeeringContainer](./docs/NetworkPeeringApi.md#deletepeeringcontainer) | **Delete** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container |
*NetworkPeeringApi* | [DisablePeering](./docs/NetworkPeeringApi.md#disablepeering) | **Patch** /api/atlas/v2/groups/{groupId}/privateIpMode | Disable Connect via Peering-Only Mode for One Project |
*NetworkPeeringApi* | [GetPeeringConnection](./docs/NetworkPeeringApi.md#getpeeringconnection) | **Get** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project |
*NetworkPeeringApi* | [GetPeeringContainer](./docs/NetworkPeeringApi.md#getpeeringcontainer) | **Get** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Return One Network Peering Container |
*NetworkPeeringApi* | [ListPeeringConnections](./docs/NetworkPeeringApi.md#listpeeringconnections) | **Get** /api/atlas/v2/groups/{groupId}/peers | Return All Network Peering Connections in One Project |
*NetworkPeeringApi* | [ListPeeringContainerByCloudProvider](./docs/NetworkPeeringApi.md#listpeeringcontainerbycloudprovider) | **Get** /api/atlas/v2/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider |
*NetworkPeeringApi* | [ListPeeringContainers](./docs/NetworkPeeringApi.md#listpeeringcontainers) | **Get** /api/atlas/v2/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project |
*NetworkPeeringApi* | [UpdatePeeringConnection](./docs/NetworkPeeringApi.md#updatepeeringconnection) | **Patch** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Update One Network Peering Connection |
*NetworkPeeringApi* | [UpdatePeeringContainer](./docs/NetworkPeeringApi.md#updatepeeringcontainer) | **Patch** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Update One Network Peering Container |
*NetworkPeeringApi* | [VerifyConnectViaPeeringOnlyModeForOneProject](./docs/NetworkPeeringApi.md#verifyconnectviapeeringonlymodeforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/privateIpMode | Verify Connect via Peering-Only Mode for One Project |
*OnlineArchiveApi* | [CreateOnlineArchive](./docs/OnlineArchiveApi.md#createonlinearchive) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive |
*OnlineArchiveApi* | [DeleteOnlineArchive](./docs/OnlineArchiveApi.md#deleteonlinearchive) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive |
*OnlineArchiveApi* | [DownloadOnlineArchiveQueryLogs](./docs/OnlineArchiveApi.md#downloadonlinearchivequerylogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs |
*OnlineArchiveApi* | [GetOnlineArchive](./docs/OnlineArchiveApi.md#getonlinearchive) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive |
*OnlineArchiveApi* | [ListOnlineArchives](./docs/OnlineArchiveApi.md#listonlinearchives) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster |
*OnlineArchiveApi* | [UpdateOnlineArchive](./docs/OnlineArchiveApi.md#updateonlinearchive) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive |
*OrganizationsApi* | [CreateOrganization](./docs/OrganizationsApi.md#createorganization) | **Post** /api/atlas/v2/orgs | Create One Organization |
*OrganizationsApi* | [CreateOrganizationInvitation](./docs/OrganizationsApi.md#createorganizationinvitation) | **Post** /api/atlas/v2/orgs/{orgId}/invites | Invite One MongoDB Cloud User to One Atlas Organization |
*OrganizationsApi* | [DeleteOrganization](./docs/OrganizationsApi.md#deleteorganization) | **Delete** /api/atlas/v2/orgs/{orgId} | Remove One Organization |
*OrganizationsApi* | [DeleteOrganizationInvitation](./docs/OrganizationsApi.md#deleteorganizationinvitation) | **Delete** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Remove One Organization Invitation |
*OrganizationsApi* | [GetOrganization](./docs/OrganizationsApi.md#getorganization) | **Get** /api/atlas/v2/orgs/{orgId} | Return One Organization |
*OrganizationsApi* | [GetOrganizationInvitation](./docs/OrganizationsApi.md#getorganizationinvitation) | **Get** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Return One Organization Invitation |
*OrganizationsApi* | [GetOrganizationSettings](./docs/OrganizationsApi.md#getorganizationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/settings | Return Settings for One Organization |
*OrganizationsApi* | [ListOrganizationInvitations](./docs/OrganizationsApi.md#listorganizationinvitations) | **Get** /api/atlas/v2/orgs/{orgId}/invites | Return All Organization Invitations |
*OrganizationsApi* | [ListOrganizationProjects](./docs/OrganizationsApi.md#listorganizationprojects) | **Get** /api/atlas/v2/orgs/{orgId}/groups | Return All Projects in One Organization |
*OrganizationsApi* | [ListOrganizations](./docs/OrganizationsApi.md#listorganizations) | **Get** /api/atlas/v2/orgs | Return All Organizations |
*OrganizationsApi* | [UpdateOrganization](./docs/OrganizationsApi.md#updateorganization) | **Patch** /api/atlas/v2/orgs/{orgId} | Update One Organization |
*OrganizationsApi* | [UpdateOrganizationInvitation](./docs/OrganizationsApi.md#updateorganizationinvitation) | **Patch** /api/atlas/v2/orgs/{orgId}/invites | Update One Organization Invitation |
*OrganizationsApi* | [UpdateOrganizationInvitationById](./docs/OrganizationsApi.md#updateorganizationinvitationbyid) | **Patch** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Update One Organization Invitation by Invitation ID |
*OrganizationsApi* | [UpdateOrganizationRoles](./docs/OrganizationsApi.md#updateorganizationroles) | **Put** /api/atlas/v2/orgs/{orgId}/users/{userId}/roles | Update Organization Roles for One MongoDB Cloud User |
*OrganizationsApi* | [UpdateOrganizationSettings](./docs/OrganizationsApi.md#updateorganizationsettings) | **Patch** /api/atlas/v2/orgs/{orgId}/settings | Update Settings for One Organization |
*PerformanceAdvisorApi* | [DisableSlowOperationThresholding](./docs/PerformanceAdvisorApi.md#disableslowoperationthresholding) | **Delete** /api/atlas/v2/groups/{groupId}/managedSlowMs/disable | Disable Managed Slow Operation Threshold |
*PerformanceAdvisorApi* | [EnableSlowOperationThresholding](./docs/PerformanceAdvisorApi.md#enableslowoperationthresholding) | **Post** /api/atlas/v2/groups/{groupId}/managedSlowMs/enable | Enable Managed Slow Operation Threshold |
*PerformanceAdvisorApi* | [GetManagedSlowMs](./docs/PerformanceAdvisorApi.md#getmanagedslowms) | **Get** /api/atlas/v2/groups/{groupId}/managedSlowMs | Return Managed Slow Operation Threshold Status |
*PerformanceAdvisorApi* | [GetServerlessAutoIndexing](./docs/PerformanceAdvisorApi.md#getserverlessautoindexing) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing | Return Serverless Auto-Indexing Status |
*PerformanceAdvisorApi* | [ListClusterSuggestedIndexes](./docs/PerformanceAdvisorApi.md#listclustersuggestedindexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/suggestedIndexes | Return All Suggested Indexes |
*PerformanceAdvisorApi* | [ListDropIndexes](./docs/PerformanceAdvisorApi.md#listdropindexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/dropIndexSuggestions | Return All Suggested Indexes to Drop |
*PerformanceAdvisorApi* | [ListSchemaAdvice](./docs/PerformanceAdvisorApi.md#listschemaadvice) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/schemaAdvice | Return Schema Advice |
*PerformanceAdvisorApi* | [ListSlowQueries](./docs/PerformanceAdvisorApi.md#listslowqueries) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries |
*PerformanceAdvisorApi* | [ListSlowQueryNamespaces](./docs/PerformanceAdvisorApi.md#listslowquerynamespaces) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host |
*PerformanceAdvisorApi* | [ListSuggestedIndexes](./docs/PerformanceAdvisorApi.md#listsuggestedindexes) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return All Suggested Indexes |
*PerformanceAdvisorApi* | [SetServerlessAutoIndexing](./docs/PerformanceAdvisorApi.md#setserverlessautoindexing) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing | Set Serverless Auto-Indexing Status |
*PrivateEndpointServicesApi* | [CreatePrivateEndpoint](./docs/PrivateEndpointServicesApi.md#createprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint | Create One Private Endpoint for One Provider |
*PrivateEndpointServicesApi* | [CreatePrivateEndpointService](./docs/PrivateEndpointServicesApi.md#createprivateendpointservice) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService | Create One Private Endpoint Service for One Provider |
*PrivateEndpointServicesApi* | [DeletePrivateEndpoint](./docs/PrivateEndpointServicesApi.md#deleteprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Remove One Private Endpoint for One Provider |
*PrivateEndpointServicesApi* | [DeletePrivateEndpointService](./docs/PrivateEndpointServicesApi.md#deleteprivateendpointservice) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Remove One Private Endpoint Service for One Provider |
*PrivateEndpointServicesApi* | [GetPrivateEndpoint](./docs/PrivateEndpointServicesApi.md#getprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Return One Private Endpoint for One Provider |
*PrivateEndpointServicesApi* | [GetPrivateEndpointService](./docs/PrivateEndpointServicesApi.md#getprivateendpointservice) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Return One Private Endpoint Service for One Provider |
*PrivateEndpointServicesApi* | [GetRegionalizedPrivateEndpointSetting](./docs/PrivateEndpointServicesApi.md#getregionalizedprivateendpointsetting) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Return Regionalized Private Endpoint Status |
*PrivateEndpointServicesApi* | [ListPrivateEndpointServices](./docs/PrivateEndpointServicesApi.md#listprivateendpointservices) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService | Return All Private Endpoint Services for One Provider |
*PrivateEndpointServicesApi* | [ToggleRegionalizedPrivateEndpointSetting](./docs/PrivateEndpointServicesApi.md#toggleregionalizedprivateendpointsetting) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Toggle Regionalized Private Endpoint Status |
*ProgrammaticAPIKeysApi* | [AddProjectApiKey](./docs/ProgrammaticAPIKeysApi.md#addprojectapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Assign One Organization API Key to One Project |
*ProgrammaticAPIKeysApi* | [CreateApiKey](./docs/ProgrammaticAPIKeysApi.md#createapikey) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys | Create One Organization API Key |
*ProgrammaticAPIKeysApi* | [CreateApiKeyAccessList](./docs/ProgrammaticAPIKeysApi.md#createapikeyaccesslist) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Create One Access List Entry for One Organization API Key |
*ProgrammaticAPIKeysApi* | [CreateProjectApiKey](./docs/ProgrammaticAPIKeysApi.md#createprojectapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys | Create and Assign One Organization API Key to One Project |
*ProgrammaticAPIKeysApi* | [DeleteApiKey](./docs/ProgrammaticAPIKeysApi.md#deleteapikey) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Remove One Organization API Key |
*ProgrammaticAPIKeysApi* | [DeleteApiKeyAccessListEntry](./docs/ProgrammaticAPIKeysApi.md#deleteapikeyaccesslistentry) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Remove One Access List Entry for One Organization API Key |
*ProgrammaticAPIKeysApi* | [GetApiKey](./docs/ProgrammaticAPIKeysApi.md#getapikey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Return One Organization API Key |
*ProgrammaticAPIKeysApi* | [GetApiKeyAccessList](./docs/ProgrammaticAPIKeysApi.md#getapikeyaccesslist) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Return One Access List Entry for One Organization API Key |
*ProgrammaticAPIKeysApi* | [ListApiKeyAccessListsEntries](./docs/ProgrammaticAPIKeysApi.md#listapikeyaccesslistsentries) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Return All Access List Entries for One Organization API Key |
*ProgrammaticAPIKeysApi* | [ListApiKeys](./docs/ProgrammaticAPIKeysApi.md#listapikeys) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys | Return All Organization API Keys |
*ProgrammaticAPIKeysApi* | [ListProjectApiKeys](./docs/ProgrammaticAPIKeysApi.md#listprojectapikeys) | **Get** /api/atlas/v2/groups/{groupId}/apiKeys | Return All Organization API Keys Assigned to One Project |
*ProgrammaticAPIKeysApi* | [RemoveProjectApiKey](./docs/ProgrammaticAPIKeysApi.md#removeprojectapikey) | **Delete** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Unassign One Organization API Key from One Project |
*ProgrammaticAPIKeysApi* | [UpdateApiKey](./docs/ProgrammaticAPIKeysApi.md#updateapikey) | **Patch** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Update One Organization API Key |
*ProgrammaticAPIKeysApi* | [UpdateApiKeyRoles](./docs/ProgrammaticAPIKeysApi.md#updateapikeyroles) | **Patch** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Update Organization API Key Roles for One Project |
*ProjectIPAccessListApi* | [CreateProjectIpAccessList](./docs/ProjectIPAccessListApi.md#createprojectipaccesslist) | **Post** /api/atlas/v2/groups/{groupId}/accessList | Add Entries to Project IP Access List |
*ProjectIPAccessListApi* | [DeleteProjectIpAccessList](./docs/ProjectIPAccessListApi.md#deleteprojectipaccesslist) | **Delete** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Remove One Entry from One Project IP Access List |
*ProjectIPAccessListApi* | [GetProjectIpAccessListStatus](./docs/ProjectIPAccessListApi.md#getprojectipaccessliststatus) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue}/status | Return Status of One Project IP Access List Entry |
*ProjectIPAccessListApi* | [GetProjectIpList](./docs/ProjectIPAccessListApi.md#getprojectiplist) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Return One Project IP Access List Entry |
*ProjectIPAccessListApi* | [ListProjectIpAccessLists](./docs/ProjectIPAccessListApi.md#listprojectipaccesslists) | **Get** /api/atlas/v2/groups/{groupId}/accessList | Return All Project IP Access List Entries |
*ProjectsApi* | [AddUserToProject](./docs/ProjectsApi.md#addusertoproject) | **Post** /api/atlas/v2/groups/{groupId}/access | Add One MongoDB Cloud User to One Project |
*ProjectsApi* | [CreateProject](./docs/ProjectsApi.md#createproject) | **Post** /api/atlas/v2/groups | Create One Project |
*ProjectsApi* | [CreateProjectInvitation](./docs/ProjectsApi.md#createprojectinvitation) | **Post** /api/atlas/v2/groups/{groupId}/invites | Invite One MongoDB Cloud User to One Project |
*ProjectsApi* | [DeleteProject](./docs/ProjectsApi.md#deleteproject) | **Delete** /api/atlas/v2/groups/{groupId} | Remove One Project |
*ProjectsApi* | [DeleteProjectInvitation](./docs/ProjectsApi.md#deleteprojectinvitation) | **Delete** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Remove One Project Invitation |
*ProjectsApi* | [DeleteProjectLimit](./docs/ProjectsApi.md#deleteprojectlimit) | **Delete** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Remove One Project Limit |
*ProjectsApi* | [GetProject](./docs/ProjectsApi.md#getproject) | **Get** /api/atlas/v2/groups/{groupId} | Return One Project |
*ProjectsApi* | [GetProjectByName](./docs/ProjectsApi.md#getprojectbyname) | **Get** /api/atlas/v2/groups/byName/{groupName} | Return One Project by Name |
*ProjectsApi* | [GetProjectInvitation](./docs/ProjectsApi.md#getprojectinvitation) | **Get** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Return One Project Invitation |
*ProjectsApi* | [GetProjectLimit](./docs/ProjectsApi.md#getprojectlimit) | **Get** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Return One Limit for One Project |
*ProjectsApi* | [GetProjectLtsVersions](./docs/ProjectsApi.md#getprojectltsversions) | **Get** /api/atlas/v2/groups/{groupId}/mongoDBVersions | Return All Available MongoDB LTS Versions for Clusters in One Project |
*ProjectsApi* | [GetProjectSettings](./docs/ProjectsApi.md#getprojectsettings) | **Get** /api/atlas/v2/groups/{groupId}/settings | Return Project Settings |
*ProjectsApi* | [ListProjectInvitations](./docs/ProjectsApi.md#listprojectinvitations) | **Get** /api/atlas/v2/groups/{groupId}/invites | Return All Project Invitations |
*ProjectsApi* | [ListProjectLimits](./docs/ProjectsApi.md#listprojectlimits) | **Get** /api/atlas/v2/groups/{groupId}/limits | Return All Limits for One Project |
*ProjectsApi* | [ListProjects](./docs/ProjectsApi.md#listprojects) | **Get** /api/atlas/v2/groups | Return All Projects |
*ProjectsApi* | [MigrateProjectToAnotherOrg](./docs/ProjectsApi.md#migrateprojecttoanotherorg) | **Post** /api/atlas/v2/groups/{groupId}:migrate | Migrate One Project to Another Organization |
*ProjectsApi* | [ReturnAllIpAddresses](./docs/ProjectsApi.md#returnallipaddresses) | **Get** /api/atlas/v2/groups/{groupId}/ipAddresses | Return All IP Addresses for One Project |
*ProjectsApi* | [SetProjectLimit](./docs/ProjectsApi.md#setprojectlimit) | **Patch** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Set One Project Limit |
*ProjectsApi* | [UpdateProject](./docs/ProjectsApi.md#updateproject) | **Patch** /api/atlas/v2/groups/{groupId} | Update One Project |
*ProjectsApi* | [UpdateProjectInvitation](./docs/ProjectsApi.md#updateprojectinvitation) | **Patch** /api/atlas/v2/groups/{groupId}/invites | Update One Project Invitation |
*ProjectsApi* | [UpdateProjectInvitationById](./docs/ProjectsApi.md#updateprojectinvitationbyid) | **Patch** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Update One Project Invitation by Invitation ID |
*ProjectsApi* | [UpdateProjectRoles](./docs/ProjectsApi.md#updateprojectroles) | **Put** /api/atlas/v2/groups/{groupId}/users/{userId}/roles | Update Project Roles for One MongoDB Cloud User |
*ProjectsApi* | [UpdateProjectSettings](./docs/ProjectsApi.md#updateprojectsettings) | **Patch** /api/atlas/v2/groups/{groupId}/settings | Update Project Settings |
*PushBasedLogExportApi* | [CreatePushBasedLogConfiguration](./docs/PushBasedLogExportApi.md#createpushbasedlogconfiguration) | **Post** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Enable Push-Based Log Export for One Project |
*PushBasedLogExportApi* | [DeletePushBasedLogConfiguration](./docs/PushBasedLogExportApi.md#deletepushbasedlogconfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Disable Push-Based Log Export for One Project |
*PushBasedLogExportApi* | [GetPushBasedLogConfiguration](./docs/PushBasedLogExportApi.md#getpushbasedlogconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Return Push-Based Log Export Configuration for One Project |
*PushBasedLogExportApi* | [UpdatePushBasedLogConfiguration](./docs/PushBasedLogExportApi.md#updatepushbasedlogconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Update Push-Based Log Export for One Project |
*ResourcePoliciesApi* | [CreateOrgResourcePolicy](./docs/ResourcePoliciesApi.md#createorgresourcepolicy) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Create One Atlas Resource Policy |
*ResourcePoliciesApi* | [DeleteOrgResourcePolicy](./docs/ResourcePoliciesApi.md#deleteorgresourcepolicy) | **Delete** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Delete One Atlas Resource Policy |
*ResourcePoliciesApi* | [GetOrgResourcePolicy](./docs/ResourcePoliciesApi.md#getorgresourcepolicy) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Return One Atlas Resource Policy |
*ResourcePoliciesApi* | [GetResourcesNonCompliant](./docs/ResourcePoliciesApi.md#getresourcesnoncompliant) | **Get** /api/atlas/v2/orgs/{orgId}/nonCompliantResources | Return All Non-Compliant Resources |
*ResourcePoliciesApi* | [ListOrgResourcePolicies](./docs/ResourcePoliciesApi.md#listorgresourcepolicies) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Return All Atlas Resource Policies |
*ResourcePoliciesApi* | [UpdateOrgResourcePolicy](./docs/ResourcePoliciesApi.md#updateorgresourcepolicy) | **Patch** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Update One Atlas Resource Policy |
*ResourcePoliciesApi* | [ValidateAtlasResourcePolicy](./docs/ResourcePoliciesApi.md#validateatlasresourcepolicy) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies:validate | Validate One Atlas Resource Policy |
*RollingIndexApi* | [CreateRollingIndex](./docs/RollingIndexApi.md#createrollingindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/index | Create One Rolling Index |
*RootApi* | [GetSystemStatus](./docs/RootApi.md#getsystemstatus) | **Get** /api/atlas/v2 | Return the status of this MongoDB application |
*RootApi* | [ReturnAllControlPlaneIpAddresses](./docs/RootApi.md#returnallcontrolplaneipaddresses) | **Get** /api/atlas/v2/unauth/controlPlaneIPAddresses | Return All Control Plane IP Addresses |
*ServerlessInstancesApi* | [CreateServerlessInstance](./docs/ServerlessInstancesApi.md#createserverlessinstance) | **Post** /api/atlas/v2/groups/{groupId}/serverless | Create One Serverless Instance in One Project |
*ServerlessInstancesApi* | [DeleteServerlessInstance](./docs/ServerlessInstancesApi.md#deleteserverlessinstance) | **Delete** /api/atlas/v2/groups/{groupId}/serverless/{name} | Remove One Serverless Instance from One Project |
*ServerlessInstancesApi* | [GetServerlessInstance](./docs/ServerlessInstancesApi.md#getserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{name} | Return One Serverless Instance from One Project |
*ServerlessInstancesApi* | [ListServerlessInstances](./docs/ServerlessInstancesApi.md#listserverlessinstances) | **Get** /api/atlas/v2/groups/{groupId}/serverless | Return All Serverless Instances in One Project |
*ServerlessInstancesApi* | [UpdateServerlessInstance](./docs/ServerlessInstancesApi.md#updateserverlessinstance) | **Patch** /api/atlas/v2/groups/{groupId}/serverless/{name} | Update One Serverless Instance in One Project |
*ServerlessPrivateEndpointsApi* | [CreateServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#createserverlessprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Create One Private Endpoint for One Serverless Instance |
*ServerlessPrivateEndpointsApi* | [DeleteServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#deleteserverlessprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Remove One Private Endpoint for One Serverless Instance |
*ServerlessPrivateEndpointsApi* | [GetServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#getserverlessprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Return One Private Endpoint for One Serverless Instance |
*ServerlessPrivateEndpointsApi* | [ListServerlessPrivateEndpoints](./docs/ServerlessPrivateEndpointsApi.md#listserverlessprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Return All Private Endpoints for One Serverless Instance |
*ServerlessPrivateEndpointsApi* | [UpdateServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#updateserverlessprivateendpoint) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Update One Private Endpoint for One Serverless Instance |
*ServiceAccountsApi* | [AddProjectServiceAccount](./docs/ServiceAccountsApi.md#addprojectserviceaccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}:invite | Assign One Service Account to One Project |
*ServiceAccountsApi* | [CreateProjectServiceAccount](./docs/ServiceAccountsApi.md#createprojectserviceaccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts | Create One Project Service Account |
*ServiceAccountsApi* | [CreateProjectServiceAccountAccessList](./docs/ServiceAccountsApi.md#createprojectserviceaccountaccesslist) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Project Service Account |
*ServiceAccountsApi* | [CreateProjectServiceAccountSecret](./docs/ServiceAccountsApi.md#createprojectserviceaccountsecret) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets | Create One Project Service Account Secret |
*ServiceAccountsApi* | [CreateServiceAccount](./docs/ServiceAccountsApi.md#createserviceaccount) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Create One Organization Service Account |
*ServiceAccountsApi* | [CreateServiceAccountAccessList](./docs/ServiceAccountsApi.md#createserviceaccountaccesslist) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Organization Service Account |
*ServiceAccountsApi* | [CreateServiceAccountSecret](./docs/ServiceAccountsApi.md#createserviceaccountsecret) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets | Create One Organization Service Account Secret |
*ServiceAccountsApi* | [DeleteProjectServiceAccount](./docs/ServiceAccountsApi.md#deleteprojectserviceaccount) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Remove One Project Service Account |
*ServiceAccountsApi* | [DeleteProjectServiceAccountAccessListEntry](./docs/ServiceAccountsApi.md#deleteprojectserviceaccountaccesslistentry) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Project Service Account |
*ServiceAccountsApi* | [DeleteProjectServiceAccountSecret](./docs/ServiceAccountsApi.md#deleteprojectserviceaccountsecret) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Project Service Account Secret |
*ServiceAccountsApi* | [DeleteServiceAccount](./docs/ServiceAccountsApi.md#deleteserviceaccount) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Delete One Organization Service Account |
*ServiceAccountsApi* | [DeleteServiceAccountAccessListEntry](./docs/ServiceAccountsApi.md#deleteserviceaccountaccesslistentry) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Organization Service Account |
*ServiceAccountsApi* | [DeleteServiceAccountSecret](./docs/ServiceAccountsApi.md#deleteserviceaccountsecret) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Organization Service Account Secret |
*ServiceAccountsApi* | [GetProjectServiceAccount](./docs/ServiceAccountsApi.md#getprojectserviceaccount) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Return One Project Service Account |
*ServiceAccountsApi* | [GetServiceAccount](./docs/ServiceAccountsApi.md#getserviceaccount) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Return One Organization Service Account |
*ServiceAccountsApi* | [ListProjectServiceAccountAccessList](./docs/ServiceAccountsApi.md#listprojectserviceaccountaccesslist) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Project Service Account |
*ServiceAccountsApi* | [ListProjectServiceAccounts](./docs/ServiceAccountsApi.md#listprojectserviceaccounts) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts | Return All Project Service Accounts |
*ServiceAccountsApi* | [ListServiceAccountAccessList](./docs/ServiceAccountsApi.md#listserviceaccountaccesslist) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Organization Service Account |
*ServiceAccountsApi* | [ListServiceAccountProjects](./docs/ServiceAccountsApi.md#listserviceaccountprojects) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/groups | Return All Service Account Project Assignments |
*ServiceAccountsApi* | [ListServiceAccounts](./docs/ServiceAccountsApi.md#listserviceaccounts) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Return All Organization Service Accounts |
*ServiceAccountsApi* | [UpdateProjectServiceAccount](./docs/ServiceAccountsApi.md#updateprojectserviceaccount) | **Patch** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Update One Project Service Account |
*ServiceAccountsApi* | [UpdateServiceAccount](./docs/ServiceAccountsApi.md#updateserviceaccount) | **Patch** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Update One Organization Service Account |
*SharedTierRestoreJobsApi* | [CreateSharedClusterBackupRestoreJob](./docs/SharedTierRestoreJobsApi.md#createsharedclusterbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore | Create One Restore Job for One M2 or M5 Cluster |
*SharedTierRestoreJobsApi* | [GetSharedClusterBackupRestoreJob](./docs/SharedTierRestoreJobsApi.md#getsharedclusterbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId} | Return One Restore Job for One M2 or M5 Cluster |
*SharedTierRestoreJobsApi* | [ListSharedClusterBackupRestoreJobs](./docs/SharedTierRestoreJobsApi.md#listsharedclusterbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores | Return All Restore Jobs for One M2 or M5 Cluster |
*SharedTierSnapshotsApi* | [DownloadSharedClusterBackup](./docs/SharedTierSnapshotsApi.md#downloadsharedclusterbackup) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/download | Download One M2 or M5 Cluster Snapshot |
*SharedTierSnapshotsApi* | [GetSharedClusterBackup](./docs/SharedTierSnapshotsApi.md#getsharedclusterbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId} | Return One Snapshot of One M2 or M5 Cluster |
*SharedTierSnapshotsApi* | [ListSharedClusterBackups](./docs/SharedTierSnapshotsApi.md#listsharedclusterbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots | Return All Snapshots for One M2 or M5 Cluster |
*StreamsApi* | [AcceptVpcPeeringConnection](./docs/StreamsApi.md#acceptvpcpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:accept | Accept One Incoming VPC Peering Connection |
*StreamsApi* | [CreatePrivateLinkConnection](./docs/StreamsApi.md#createprivatelinkconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Create One Private Link Connection |
*StreamsApi* | [CreateStreamConnection](./docs/StreamsApi.md#createstreamconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Create One Stream Connection |
*StreamsApi* | [CreateStreamInstance](./docs/StreamsApi.md#createstreaminstance) | **Post** /api/atlas/v2/groups/{groupId}/streams | Create One Stream Instance |
*StreamsApi* | [CreateStreamInstanceWithSampleConnections](./docs/StreamsApi.md#createstreaminstancewithsampleconnections) | **Post** /api/atlas/v2/groups/{groupId}/streams:withSampleConnections | Create One Stream Instance with Sample Connections |
*StreamsApi* | [CreateStreamProcessor](./docs/StreamsApi.md#createstreamprocessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor | Create One Stream Processor |
*StreamsApi* | [DeletePrivateLinkConnection](./docs/StreamsApi.md#deleteprivatelinkconnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Delete One Private Link Connection |
*StreamsApi* | [DeleteStreamConnection](./docs/StreamsApi.md#deletestreamconnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Delete One Stream Connection |
*StreamsApi* | [DeleteStreamInstance](./docs/StreamsApi.md#deletestreaminstance) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Delete One Stream Instance |
*StreamsApi* | [DeleteStreamProcessor](./docs/StreamsApi.md#deletestreamprocessor) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Delete One Stream Processor |
*StreamsApi* | [DeleteVpcPeeringConnection](./docs/StreamsApi.md#deletevpcpeeringconnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id} | Delete One VPC Peering Connection |
*StreamsApi* | [DownloadStreamTenantAuditLogs](./docs/StreamsApi.md#downloadstreamtenantauditlogs) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/auditLogs | Download Audit Logs for One Atlas Stream Processing Instance |
*StreamsApi* | [GetAccountDetails](./docs/StreamsApi.md#getaccountdetails) | **Get** /api/atlas/v2/groups/{groupId}/streams/accountDetails | Return Account ID and VPC ID for One Project and Region |
*StreamsApi* | [GetActiveVpcPeeringConnections](./docs/StreamsApi.md#getactivevpcpeeringconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/activeVpcPeeringConnections | Return All Active Incoming VPC Peering Connections |
*StreamsApi* | [GetPrivateLinkConnection](./docs/StreamsApi.md#getprivatelinkconnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Return One Private Link Connection |
*StreamsApi* | [GetStreamConnection](./docs/StreamsApi.md#getstreamconnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Return One Stream Connection |
*StreamsApi* | [GetStreamInstance](./docs/StreamsApi.md#getstreaminstance) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Return One Stream Instance |
*StreamsApi* | [GetStreamProcessor](./docs/StreamsApi.md#getstreamprocessor) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Return One Stream Processor |
*StreamsApi* | [GetVpcPeeringConnections](./docs/StreamsApi.md#getvpcpeeringconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections | Return All VPC Peering Connections |
*StreamsApi* | [ListPrivateLinkConnections](./docs/StreamsApi.md#listprivatelinkconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Return All Private Link Connections |
*StreamsApi* | [ListStreamConnections](./docs/StreamsApi.md#liststreamconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Return All Connections of the Stream Instances |
*StreamsApi* | [ListStreamInstances](./docs/StreamsApi.md#liststreaminstances) | **Get** /api/atlas/v2/groups/{groupId}/streams | Return All Stream Instances in One Project |
*StreamsApi* | [ListStreamProcessors](./docs/StreamsApi.md#liststreamprocessors) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processors | Return All Stream Processors in One Stream Instance |
*StreamsApi* | [ModifyStreamProcessor](./docs/StreamsApi.md#modifystreamprocessor) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Update One Stream Processor |
*StreamsApi* | [RejectVpcPeeringConnection](./docs/StreamsApi.md#rejectvpcpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:reject | Reject One Incoming VPC Peering Connection |
*StreamsApi* | [StartStreamProcessor](./docs/StreamsApi.md#startstreamprocessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:start | Start One Stream Processor |
*StreamsApi* | [StartStreamProcessorWith](./docs/StreamsApi.md#startstreamprocessorwith) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:startWith | Start One Stream Processor With Options |
*StreamsApi* | [StopStreamProcessor](./docs/StreamsApi.md#stopstreamprocessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:stop | Stop One Stream Processor |
*StreamsApi* | [UpdateStreamConnection](./docs/StreamsApi.md#updatestreamconnection) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Update One Stream Connection |
*StreamsApi* | [UpdateStreamInstance](./docs/StreamsApi.md#updatestreaminstance) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Update One Stream Instance |
*TeamsApi* | [AddAllTeamsToProject](./docs/TeamsApi.md#addallteamstoproject) | **Post** /api/atlas/v2/groups/{groupId}/teams | Add One Team to One Project |
*TeamsApi* | [AddTeamUser](./docs/TeamsApi.md#addteamuser) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Assign MongoDB Cloud Users in One Organization to One Team |
*TeamsApi* | [CreateTeam](./docs/TeamsApi.md#createteam) | **Post** /api/atlas/v2/orgs/{orgId}/teams | Create One Team in One Organization |
*TeamsApi* | [DeleteTeam](./docs/TeamsApi.md#deleteteam) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Remove One Team from One Organization |
*TeamsApi* | [GetProjectTeam](./docs/TeamsApi.md#getprojectteam) | **Get** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Return One Team in One Project |
*TeamsApi* | [GetTeamById](./docs/TeamsApi.md#getteambyid) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Return One Team by ID |
*TeamsApi* | [GetTeamByName](./docs/TeamsApi.md#getteambyname) | **Get** /api/atlas/v2/orgs/{orgId}/teams/byName/{teamName} | Return One Team by Name |
*TeamsApi* | [ListOrganizationTeams](./docs/TeamsApi.md#listorganizationteams) | **Get** /api/atlas/v2/orgs/{orgId}/teams | Return All Teams in One Organization |
*TeamsApi* | [ListProjectTeams](./docs/TeamsApi.md#listprojectteams) | **Get** /api/atlas/v2/groups/{groupId}/teams | Return All Teams in One Project |
*TeamsApi* | [RemoveProjectTeam](./docs/TeamsApi.md#removeprojectteam) | **Delete** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Remove One Team from One Project |
*TeamsApi* | [RemoveTeamUser](./docs/TeamsApi.md#removeteamuser) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users/{userId} | Remove One MongoDB Cloud User from One Team |
*TeamsApi* | [RenameTeam](./docs/TeamsApi.md#renameteam) | **Patch** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Rename One Team |
*TeamsApi* | [UpdateTeamRoles](./docs/TeamsApi.md#updateteamroles) | **Patch** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Update Team Roles in One Project |
*ThirdPartyIntegrationsApi* | [CreateThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#createthirdpartyintegration) | **Post** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Configure One Third-Party Service Integration |
*ThirdPartyIntegrationsApi* | [DeleteThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#deletethirdpartyintegration) | **Delete** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Remove One Third-Party Service Integration |
*ThirdPartyIntegrationsApi* | [GetThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#getthirdpartyintegration) | **Get** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Return One Third-Party Service Integration |
*ThirdPartyIntegrationsApi* | [ListThirdPartyIntegrations](./docs/ThirdPartyIntegrationsApi.md#listthirdpartyintegrations) | **Get** /api/atlas/v2/groups/{groupId}/integrations | Return All Active Third-Party Service Integrations |
*ThirdPartyIntegrationsApi* | [UpdateThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#updatethirdpartyintegration) | **Put** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Update One Third-Party Service Integration |
*X509AuthenticationApi* | [CreateDatabaseUserCertificate](./docs/X509AuthenticationApi.md#createdatabaseusercertificate) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Create One X.509 Certificate for One Database User |
*X509AuthenticationApi* | [DisableCustomerManagedX509](./docs/X509AuthenticationApi.md#disablecustomermanagedx509) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/customerX509 | Disable Customer-Managed X.509 |
*X509AuthenticationApi* | [ListDatabaseUserCertificates](./docs/X509AuthenticationApi.md#listdatabaseusercertificates) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Return All X.509 Certificates Assigned to One Database User |


## Documentation For Models

 - [AWSCustomDNSEnabled](./docs/AWSCustomDNSEnabled.md)
 - [AWSKMSConfiguration](./docs/AWSKMSConfiguration.md)
 - [AccessListItem](./docs/AccessListItem.md)
 - [AccountDetails](./docs/AccountDetails.md)
 - [AcknowledgeAlert](./docs/AcknowledgeAlert.md)
 - [AddOrRemoveGroupRole](./docs/AddOrRemoveGroupRole.md)
 - [AddOrRemoveOrgRole](./docs/AddOrRemoveOrgRole.md)
 - [AddOrRemoveUserFromTeam](./docs/AddOrRemoveUserFromTeam.md)
 - [AddUserToTeam](./docs/AddUserToTeam.md)
 - [AdvancedAutoScalingSettings](./docs/AdvancedAutoScalingSettings.md)
 - [AdvancedComputeAutoScaling](./docs/AdvancedComputeAutoScaling.md)
 - [AdvancedDiskBackupSnapshotSchedulePolicy](./docs/AdvancedDiskBackupSnapshotSchedulePolicy.md)
 - [AlertViewForNdsGroup](./docs/AlertViewForNdsGroup.md)
 - [AlertsNotificationRootForGroup](./docs/AlertsNotificationRootForGroup.md)
 - [AlertsToggle](./docs/AlertsToggle.md)
 - [ApiAtlasCheckpoint](./docs/ApiAtlasCheckpoint.md)
 - [ApiAtlasClusterAdvancedConfiguration](./docs/ApiAtlasClusterAdvancedConfiguration.md)
 - [ApiAtlasFTSAnalyzers](./docs/ApiAtlasFTSAnalyzers.md)
 - [ApiAtlasFTSAnalyzersTokenizer](./docs/ApiAtlasFTSAnalyzersTokenizer.md)
 - [ApiAtlasFTSMappings](./docs/ApiAtlasFTSMappings.md)
 - [ApiAtlasInvalidPolicy](./docs/ApiAtlasInvalidPolicy.md)
 - [ApiAtlasInvalidPolicyErrorDetail](./docs/ApiAtlasInvalidPolicyErrorDetail.md)
 - [ApiAtlasInvalidResourcePolicyCreateError](./docs/ApiAtlasInvalidResourcePolicyCreateError.md)
 - [ApiAtlasNonCompliantResource](./docs/ApiAtlasNonCompliantResource.md)
 - [ApiAtlasPolicy](./docs/ApiAtlasPolicy.md)
 - [ApiAtlasPolicyCreate](./docs/ApiAtlasPolicyCreate.md)
 - [ApiAtlasPolicyMetadata](./docs/ApiAtlasPolicyMetadata.md)
 - [ApiAtlasResourcePolicy](./docs/ApiAtlasResourcePolicy.md)
 - [ApiAtlasResourcePolicyCreate](./docs/ApiAtlasResourcePolicyCreate.md)
 - [ApiAtlasResourcePolicyEdit](./docs/ApiAtlasResourcePolicyEdit.md)
 - [ApiAtlasResourcePolicyMetadata](./docs/ApiAtlasResourcePolicyMetadata.md)
 - [ApiAtlasSnapshotSchedule](./docs/ApiAtlasSnapshotSchedule.md)
 - [ApiAtlasUserMetadata](./docs/ApiAtlasUserMetadata.md)
 - [ApiBSONTimestamp](./docs/ApiBSONTimestamp.md)
 - [ApiCheckpointPart](./docs/ApiCheckpointPart.md)
 - [ApiError](./docs/ApiError.md)
 - [ApiHostViewAtlas](./docs/ApiHostViewAtlas.md)
 - [ApiKey](./docs/ApiKey.md)
 - [ApiKeyUserDetails](./docs/ApiKeyUserDetails.md)
 - [ApiMeasurementsGeneralViewAtlas](./docs/ApiMeasurementsGeneralViewAtlas.md)
 - [ApiPrivateDownloadDeliveryUrl](./docs/ApiPrivateDownloadDeliveryUrl.md)
 - [ApiPublicUsageDetailsQueryRequest](./docs/ApiPublicUsageDetailsQueryRequest.md)
 - [ApiSearchDeploymentRequest](./docs/ApiSearchDeploymentRequest.md)
 - [ApiSearchDeploymentResponse](./docs/ApiSearchDeploymentResponse.md)
 - [ApiSearchDeploymentSpec](./docs/ApiSearchDeploymentSpec.md)
 - [AtlasClusterOutageSimulationOutageFilter](./docs/AtlasClusterOutageSimulationOutageFilter.md)
 - [AtlasOrganization](./docs/AtlasOrganization.md)
 - [AtlasSearchAnalyzer](./docs/AtlasSearchAnalyzer.md)
 - [AtlasTenantClusterUpgradeRequest20240805](./docs/AtlasTenantClusterUpgradeRequest20240805.md)
 - [AuditLog](./docs/AuditLog.md)
 - [AuthFederationRoleMapping](./docs/AuthFederationRoleMapping.md)
 - [AutoExportPolicy](./docs/AutoExportPolicy.md)
 - [AvailableCloudProviderRegion](./docs/AvailableCloudProviderRegion.md)
 - [AvailableClustersDeployment](./docs/AvailableClustersDeployment.md)
 - [AzureKeyVault](./docs/AzureKeyVault.md)
 - [BackupComplianceOnDemandPolicyItem](./docs/BackupComplianceOnDemandPolicyItem.md)
 - [BackupComplianceScheduledPolicyItem](./docs/BackupComplianceScheduledPolicyItem.md)
 - [BackupLabel](./docs/BackupLabel.md)
 - [BackupOnlineArchive](./docs/BackupOnlineArchive.md)
 - [BackupOnlineArchiveCreate](./docs/BackupOnlineArchiveCreate.md)
 - [BackupRestoreJob](./docs/BackupRestoreJob.md)
 - [BackupRestoreJobDelivery](./docs/BackupRestoreJobDelivery.md)
 - [BackupSnapshot](./docs/BackupSnapshot.md)
 - [BackupSnapshotPart](./docs/BackupSnapshotPart.md)
 - [BackupSnapshotRetention](./docs/BackupSnapshotRetention.md)
 - [BackupTenantSnapshot](./docs/BackupTenantSnapshot.md)
 - [BadRequestDetail](./docs/BadRequestDetail.md)
 - [BaseNetworkPeeringConnectionSettings](./docs/BaseNetworkPeeringConnectionSettings.md)
 - [BaseSearchIndexCreateRequestDefinition](./docs/BaseSearchIndexCreateRequestDefinition.md)
 - [BaseSearchIndexResponseLatestDefinition](./docs/BaseSearchIndexResponseLatestDefinition.md)
 - [BiConnector](./docs/BiConnector.md)
 - [BillingInvoice](./docs/BillingInvoice.md)
 - [BillingInvoiceMetadata](./docs/BillingInvoiceMetadata.md)
 - [BillingPayment](./docs/BillingPayment.md)
 - [BillingRefund](./docs/BillingRefund.md)
 - [CloudAccessRoleAssignment](./docs/CloudAccessRoleAssignment.md)
 - [CloudAppUser](./docs/CloudAppUser.md)
 - [CloudCluster](./docs/CloudCluster.md)
 - [CloudDatabaseUser](./docs/CloudDatabaseUser.md)
 - [CloudProviderAccessAWSIAMRole](./docs/CloudProviderAccessAWSIAMRole.md)
 - [CloudProviderAccessAWSIAMRoleAllOf](./docs/CloudProviderAccessAWSIAMRoleAllOf.md)
 - [CloudProviderAccessAzureServicePrincipal](./docs/CloudProviderAccessAzureServicePrincipal.md)
 - [CloudProviderAccessAzureServicePrincipalAllOf](./docs/CloudProviderAccessAzureServicePrincipalAllOf.md)
 - [CloudProviderAccessFeatureUsage](./docs/CloudProviderAccessFeatureUsage.md)
 - [CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId](./docs/CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId.md)
 - [CloudProviderAccessRole](./docs/CloudProviderAccessRole.md)
 - [CloudProviderAccessRoleRequest](./docs/CloudProviderAccessRoleRequest.md)
 - [CloudProviderAccessRoleRequestUpdate](./docs/CloudProviderAccessRoleRequestUpdate.md)
 - [CloudProviderAccessRoles](./docs/CloudProviderAccessRoles.md)
 - [CloudProviderContainer](./docs/CloudProviderContainer.md)
 - [CloudProviderEndpointServiceRequest](./docs/CloudProviderEndpointServiceRequest.md)
 - [CloudProviderRegions](./docs/CloudProviderRegions.md)
 - [CloudRegionConfig20240805](./docs/CloudRegionConfig20240805.md)
 - [CloudSearchMetrics](./docs/CloudSearchMetrics.md)
 - [ClusterAutoScalingSettings](./docs/ClusterAutoScalingSettings.md)
 - [ClusterCloudProviderInstanceSize](./docs/ClusterCloudProviderInstanceSize.md)
 - [ClusterComputeAutoScaling](./docs/ClusterComputeAutoScaling.md)
 - [ClusterConnectionStrings](./docs/ClusterConnectionStrings.md)
 - [ClusterDescription20240805](./docs/ClusterDescription20240805.md)
 - [ClusterDescriptionAutoScalingModeConfiguration](./docs/ClusterDescriptionAutoScalingModeConfiguration.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpoint](./docs/ClusterDescriptionConnectionStringsPrivateEndpoint.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpointEndpoint](./docs/ClusterDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ClusterDescriptionProcessArgs20240805](./docs/ClusterDescriptionProcessArgs20240805.md)
 - [ClusterFreeAutoScaling](./docs/ClusterFreeAutoScaling.md)
 - [ClusterIPAddresses](./docs/ClusterIPAddresses.md)
 - [ClusterOutageSimulation](./docs/ClusterOutageSimulation.md)
 - [ClusterProviderSettings](./docs/ClusterProviderSettings.md)
 - [ClusterSearchIndex](./docs/ClusterSearchIndex.md)
 - [ClusterServerlessBackupOptions](./docs/ClusterServerlessBackupOptions.md)
 - [ClusterStatus](./docs/ClusterStatus.md)
 - [CollStatsRankedNamespaces](./docs/CollStatsRankedNamespaces.md)
 - [Collation](./docs/Collation.md)
 - [ComponentLabel](./docs/ComponentLabel.md)
 - [ConnectedOrgConfig](./docs/ConnectedOrgConfig.md)
 - [ConnectedOrgConfigRoleAssignment](./docs/ConnectedOrgConfigRoleAssignment.md)
 - [ControlPlaneIPAddresses](./docs/ControlPlaneIPAddresses.md)
 - [CostExplorerFilterRequestBody](./docs/CostExplorerFilterRequestBody.md)
 - [CostExplorerFilterResponse](./docs/CostExplorerFilterResponse.md)
 - [CreateAtlasOrganizationApiKey](./docs/CreateAtlasOrganizationApiKey.md)
 - [CreateAtlasProjectApiKey](./docs/CreateAtlasProjectApiKey.md)
 - [CreateDataProcessRegion](./docs/CreateDataProcessRegion.md)
 - [CreateEndpointRequest](./docs/CreateEndpointRequest.md)
 - [CreateGCPForwardingRuleRequest](./docs/CreateGCPForwardingRuleRequest.md)
 - [CreateOrganizationRequest](./docs/CreateOrganizationRequest.md)
 - [CreateOrganizationResponse](./docs/CreateOrganizationResponse.md)
 - [CreatePushBasedLogExportProjectRequest](./docs/CreatePushBasedLogExportProjectRequest.md)
 - [Criteria](./docs/Criteria.md)
 - [CustomZoneMappings](./docs/CustomZoneMappings.md)
 - [DBRoleToExecute](./docs/DBRoleToExecute.md)
 - [DBUserTLSX509Settings](./docs/DBUserTLSX509Settings.md)
 - [DataExpirationRule](./docs/DataExpirationRule.md)
 - [DataFederationAzureCloudProviderConfig](./docs/DataFederationAzureCloudProviderConfig.md)
 - [DataFederationGCPCloudProviderConfig](./docs/DataFederationGCPCloudProviderConfig.md)
 - [DataFederationLimit](./docs/DataFederationLimit.md)
 - [DataFederationTenantQueryLimit](./docs/DataFederationTenantQueryLimit.md)
 - [DataLakeAWSCloudProviderConfig](./docs/DataLakeAWSCloudProviderConfig.md)
 - [DataLakeApiBase](./docs/DataLakeApiBase.md)
 - [DataLakeAtlasStoreReadConcern](./docs/DataLakeAtlasStoreReadConcern.md)
 - [DataLakeAtlasStoreReadPreference](./docs/DataLakeAtlasStoreReadPreference.md)
 - [DataLakeAtlasStoreReadPreferenceTag](./docs/DataLakeAtlasStoreReadPreferenceTag.md)
 - [DataLakeCloudProviderConfig](./docs/DataLakeCloudProviderConfig.md)
 - [DataLakeDataProcessRegion](./docs/DataLakeDataProcessRegion.md)
 - [DataLakeDatabaseCollection](./docs/DataLakeDatabaseCollection.md)
 - [DataLakeDatabaseDataSourceSettings](./docs/DataLakeDatabaseDataSourceSettings.md)
 - [DataLakeDatabaseInstance](./docs/DataLakeDatabaseInstance.md)
 - [DataLakeIngestionPipeline](./docs/DataLakeIngestionPipeline.md)
 - [DataLakePipelinesPartitionField](./docs/DataLakePipelinesPartitionField.md)
 - [DataLakeStorage](./docs/DataLakeStorage.md)
 - [DataLakeStoreSettings](./docs/DataLakeStoreSettings.md)
 - [DataLakeTenant](./docs/DataLakeTenant.md)
 - [DataProcessRegion](./docs/DataProcessRegion.md)
 - [DataProtectionSettings20231001](./docs/DataProtectionSettings20231001.md)
 - [DatabaseInheritedRole](./docs/DatabaseInheritedRole.md)
 - [DatabasePermittedNamespaceResource](./docs/DatabasePermittedNamespaceResource.md)
 - [DatabasePrivilegeAction](./docs/DatabasePrivilegeAction.md)
 - [DatabaseRollingIndexRequest](./docs/DatabaseRollingIndexRequest.md)
 - [DatabaseUserRole](./docs/DatabaseUserRole.md)
 - [DatasetRetentionPolicy](./docs/DatasetRetentionPolicy.md)
 - [DedicatedHardwareSpec20240805](./docs/DedicatedHardwareSpec20240805.md)
 - [DeleteCopiedBackups20240805](./docs/DeleteCopiedBackups20240805.md)
 - [Destination](./docs/Destination.md)
 - [DiskBackupApiPolicyItem](./docs/DiskBackupApiPolicyItem.md)
 - [DiskBackupCopySetting20240805](./docs/DiskBackupCopySetting20240805.md)
 - [DiskBackupExportJob](./docs/DiskBackupExportJob.md)
 - [DiskBackupExportJobRequest](./docs/DiskBackupExportJobRequest.md)
 - [DiskBackupExportMember](./docs/DiskBackupExportMember.md)
 - [DiskBackupOnDemandSnapshotRequest](./docs/DiskBackupOnDemandSnapshotRequest.md)
 - [DiskBackupReplicaSet](./docs/DiskBackupReplicaSet.md)
 - [DiskBackupRestoreMember](./docs/DiskBackupRestoreMember.md)
 - [DiskBackupShardedClusterSnapshot](./docs/DiskBackupShardedClusterSnapshot.md)
 - [DiskBackupShardedClusterSnapshotMember](./docs/DiskBackupShardedClusterSnapshotMember.md)
 - [DiskBackupSnapshot](./docs/DiskBackupSnapshot.md)
 - [DiskBackupSnapshotExportBucketRequest](./docs/DiskBackupSnapshotExportBucketRequest.md)
 - [DiskBackupSnapshotExportBucketResponse](./docs/DiskBackupSnapshotExportBucketResponse.md)
 - [DiskBackupSnapshotRestoreJob](./docs/DiskBackupSnapshotRestoreJob.md)
 - [DiskBackupSnapshotSchedule20240805](./docs/DiskBackupSnapshotSchedule20240805.md)
 - [DiskGBAutoScaling](./docs/DiskGBAutoScaling.md)
 - [DropIndexSuggestionsIndex](./docs/DropIndexSuggestionsIndex.md)
 - [DropIndexSuggestionsResponse](./docs/DropIndexSuggestionsResponse.md)
 - [EARPrivateEndpoint](./docs/EARPrivateEndpoint.md)
 - [EmployeeAccessGrant](./docs/EmployeeAccessGrant.md)
 - [EncryptionAtRest](./docs/EncryptionAtRest.md)
 - [EndpointService](./docs/EndpointService.md)
 - [EventTypeDetails](./docs/EventTypeDetails.md)
 - [EventViewForNdsGroup](./docs/EventViewForNdsGroup.md)
 - [EventViewForOrg](./docs/EventViewForOrg.md)
 - [ExportStatus](./docs/ExportStatus.md)
 - [ExtraRetentionSetting](./docs/ExtraRetentionSetting.md)
 - [FTSMetric](./docs/FTSMetric.md)
 - [FederatedUser](./docs/FederatedUser.md)
 - [FederationIdentityProvider](./docs/FederationIdentityProvider.md)
 - [FederationIdentityProviderUpdate](./docs/FederationIdentityProviderUpdate.md)
 - [FederationOidcIdentityProvider](./docs/FederationOidcIdentityProvider.md)
 - [FederationOidcIdentityProviderUpdate](./docs/FederationOidcIdentityProviderUpdate.md)
 - [FieldTransformation](./docs/FieldTransformation.md)
 - [FieldViolation](./docs/FieldViolation.md)
 - [FlexBackupRestoreJob20241113](./docs/FlexBackupRestoreJob20241113.md)
 - [FlexBackupRestoreJobCreate20241113](./docs/FlexBackupRestoreJobCreate20241113.md)
 - [FlexBackupSettings20241113](./docs/FlexBackupSettings20241113.md)
 - [FlexBackupSnapshot20241113](./docs/FlexBackupSnapshot20241113.md)
 - [FlexBackupSnapshotDownloadCreate20241113](./docs/FlexBackupSnapshotDownloadCreate20241113.md)
 - [FlexClusterDescription20241113](./docs/FlexClusterDescription20241113.md)
 - [FlexClusterDescriptionCreate20241113](./docs/FlexClusterDescriptionCreate20241113.md)
 - [FlexClusterDescriptionUpdate20241113](./docs/FlexClusterDescriptionUpdate20241113.md)
 - [FlexClusterMetricThreshold](./docs/FlexClusterMetricThreshold.md)
 - [FlexConnectionStrings20241113](./docs/FlexConnectionStrings20241113.md)
 - [FlexProviderSettings20241113](./docs/FlexProviderSettings20241113.md)
 - [FlexProviderSettingsCreate20241113](./docs/FlexProviderSettingsCreate20241113.md)
 - [FreeComputeAutoScalingRules](./docs/FreeComputeAutoScalingRules.md)
 - [GCPConsumerForwardingRule](./docs/GCPConsumerForwardingRule.md)
 - [GeoSharding20240805](./docs/GeoSharding20240805.md)
 - [GoogleCloudKMS](./docs/GoogleCloudKMS.md)
 - [Group](./docs/Group.md)
 - [GroupAlertsConfig](./docs/GroupAlertsConfig.md)
 - [GroupIPAddresses](./docs/GroupIPAddresses.md)
 - [GroupInvitation](./docs/GroupInvitation.md)
 - [GroupInvitationRequest](./docs/GroupInvitationRequest.md)
 - [GroupInvitationUpdateRequest](./docs/GroupInvitationUpdateRequest.md)
 - [GroupMaintenanceWindow](./docs/GroupMaintenanceWindow.md)
 - [GroupMigrationRequest](./docs/GroupMigrationRequest.md)
 - [GroupPaginatedEvent](./docs/GroupPaginatedEvent.md)
 - [GroupRole](./docs/GroupRole.md)
 - [GroupRoleAssignment](./docs/GroupRoleAssignment.md)
 - [GroupService](./docs/GroupService.md)
 - [GroupServiceAccount](./docs/GroupServiceAccount.md)
 - [GroupServiceAccountRequest](./docs/GroupServiceAccountRequest.md)
 - [GroupServiceAccountRoleAssignment](./docs/GroupServiceAccountRoleAssignment.md)
 - [GroupServiceAccountUpdateRequest](./docs/GroupServiceAccountUpdateRequest.md)
 - [GroupSettings](./docs/GroupSettings.md)
 - [GroupUpdate](./docs/GroupUpdate.md)
 - [GroupUserRequest](./docs/GroupUserRequest.md)
 - [GroupUserResponse](./docs/GroupUserResponse.md)
 - [HardwareSpec20240805](./docs/HardwareSpec20240805.md)
 - [InboundControlPlaneCloudProviderIPAddresses](./docs/InboundControlPlaneCloudProviderIPAddresses.md)
 - [IndexOptions](./docs/IndexOptions.md)
 - [IngestionPipelineRun](./docs/IngestionPipelineRun.md)
 - [IngestionSink](./docs/IngestionSink.md)
 - [IngestionSource](./docs/IngestionSource.md)
 - [InvoiceLineItem](./docs/InvoiceLineItem.md)
 - [LDAPSecuritySettings](./docs/LDAPSecuritySettings.md)
 - [LDAPVerifyConnectivityJobRequest](./docs/LDAPVerifyConnectivityJobRequest.md)
 - [LDAPVerifyConnectivityJobRequestParams](./docs/LDAPVerifyConnectivityJobRequestParams.md)
 - [LDAPVerifyConnectivityJobRequestValidation](./docs/LDAPVerifyConnectivityJobRequestValidation.md)
 - [LegacyAtlasCluster](./docs/LegacyAtlasCluster.md)
 - [LegacyAtlasTenantClusterUpgradeRequest](./docs/LegacyAtlasTenantClusterUpgradeRequest.md)
 - [LegacyReplicationSpec](./docs/LegacyReplicationSpec.md)
 - [Link](./docs/Link.md)
 - [LinkAtlas](./docs/LinkAtlas.md)
 - [LiveImportAvailableProject](./docs/LiveImportAvailableProject.md)
 - [LiveImportValidation](./docs/LiveImportValidation.md)
 - [LiveMigrationRequest20240530](./docs/LiveMigrationRequest20240530.md)
 - [LiveMigrationResponse](./docs/LiveMigrationResponse.md)
 - [ManagedNamespaces](./docs/ManagedNamespaces.md)
 - [MdbAvailableVersion](./docs/MdbAvailableVersion.md)
 - [MeasurementDiskPartition](./docs/MeasurementDiskPartition.md)
 - [MeasurementsCollStatsLatencyCluster](./docs/MeasurementsCollStatsLatencyCluster.md)
 - [MeasurementsCollStatsLatencyHost](./docs/MeasurementsCollStatsLatencyHost.md)
 - [MeasurementsIndexes](./docs/MeasurementsIndexes.md)
 - [MeasurementsNonIndex](./docs/MeasurementsNonIndex.md)
 - [MesurementsDatabase](./docs/MesurementsDatabase.md)
 - [MetricDataPoint](./docs/MetricDataPoint.md)
 - [MetricDataPointAtlas](./docs/MetricDataPointAtlas.md)
 - [MetricsMeasurement](./docs/MetricsMeasurement.md)
 - [MetricsMeasurementAtlas](./docs/MetricsMeasurementAtlas.md)
 - [MongoDBAccessLogs](./docs/MongoDBAccessLogs.md)
 - [MongoDBAccessLogsList](./docs/MongoDBAccessLogsList.md)
 - [NamespaceObj](./docs/NamespaceObj.md)
 - [Namespaces](./docs/Namespaces.md)
 - [NamespacesRequest](./docs/NamespacesRequest.md)
 - [NetworkPermissionEntry](./docs/NetworkPermissionEntry.md)
 - [NetworkPermissionEntryStatus](./docs/NetworkPermissionEntryStatus.md)
 - [NumberMetricValue](./docs/NumberMetricValue.md)
 - [OnlineArchiveSchedule](./docs/OnlineArchiveSchedule.md)
 - [OrgFederationSettings](./docs/OrgFederationSettings.md)
 - [OrgGroup](./docs/OrgGroup.md)
 - [OrgPaginatedEvent](./docs/OrgPaginatedEvent.md)
 - [OrgServiceAccount](./docs/OrgServiceAccount.md)
 - [OrgServiceAccountRequest](./docs/OrgServiceAccountRequest.md)
 - [OrgServiceAccountUpdateRequest](./docs/OrgServiceAccountUpdateRequest.md)
 - [OrgUserRequest](./docs/OrgUserRequest.md)
 - [OrgUserResponse](./docs/OrgUserResponse.md)
 - [OrgUserRolesRequest](./docs/OrgUserRolesRequest.md)
 - [OrgUserRolesResponse](./docs/OrgUserRolesResponse.md)
 - [OrgUserUpdateRequest](./docs/OrgUserUpdateRequest.md)
 - [OrganizationInvitation](./docs/OrganizationInvitation.md)
 - [OrganizationInvitationGroupRoleAssignmentsRequest](./docs/OrganizationInvitationGroupRoleAssignmentsRequest.md)
 - [OrganizationInvitationRequest](./docs/OrganizationInvitationRequest.md)
 - [OrganizationInvitationUpdateRequest](./docs/OrganizationInvitationUpdateRequest.md)
 - [OrganizationSettings](./docs/OrganizationSettings.md)
 - [OutboundControlPlaneCloudProviderIPAddresses](./docs/OutboundControlPlaneCloudProviderIPAddresses.md)
 - [PaginatedAlert](./docs/PaginatedAlert.md)
 - [PaginatedAlertConfig](./docs/PaginatedAlertConfig.md)
 - [PaginatedApiApiUser](./docs/PaginatedApiApiUser.md)
 - [PaginatedApiAppUser](./docs/PaginatedApiAppUser.md)
 - [PaginatedApiAtlasCheckpoint](./docs/PaginatedApiAtlasCheckpoint.md)
 - [PaginatedApiAtlasDatabaseUser](./docs/PaginatedApiAtlasDatabaseUser.md)
 - [PaginatedApiAtlasDiskBackupExportJob](./docs/PaginatedApiAtlasDiskBackupExportJob.md)
 - [PaginatedApiAtlasEARPrivateEndpoint](./docs/PaginatedApiAtlasEARPrivateEndpoint.md)
 - [PaginatedApiAtlasFlexBackupRestoreJob20241113](./docs/PaginatedApiAtlasFlexBackupRestoreJob20241113.md)
 - [PaginatedApiAtlasFlexBackupSnapshot20241113](./docs/PaginatedApiAtlasFlexBackupSnapshot20241113.md)
 - [PaginatedApiAtlasProviderRegions](./docs/PaginatedApiAtlasProviderRegions.md)
 - [PaginatedApiAtlasServerlessBackupRestoreJob](./docs/PaginatedApiAtlasServerlessBackupRestoreJob.md)
 - [PaginatedApiAtlasServerlessBackupSnapshot](./docs/PaginatedApiAtlasServerlessBackupSnapshot.md)
 - [PaginatedApiInvoice](./docs/PaginatedApiInvoice.md)
 - [PaginatedApiInvoiceMetadata](./docs/PaginatedApiInvoiceMetadata.md)
 - [PaginatedApiStreamsConnection](./docs/PaginatedApiStreamsConnection.md)
 - [PaginatedApiStreamsPrivateLink](./docs/PaginatedApiStreamsPrivateLink.md)
 - [PaginatedApiStreamsStreamProcessorWithStats](./docs/PaginatedApiStreamsStreamProcessorWithStats.md)
 - [PaginatedApiStreamsTenant](./docs/PaginatedApiStreamsTenant.md)
 - [PaginatedApiUserAccessListResponse](./docs/PaginatedApiUserAccessListResponse.md)
 - [PaginatedAtlasGroup](./docs/PaginatedAtlasGroup.md)
 - [PaginatedAvailableVersion](./docs/PaginatedAvailableVersion.md)
 - [PaginatedBackupSnapshot](./docs/PaginatedBackupSnapshot.md)
 - [PaginatedBackupSnapshotExportBuckets](./docs/PaginatedBackupSnapshotExportBuckets.md)
 - [PaginatedCloudBackupReplicaSet](./docs/PaginatedCloudBackupReplicaSet.md)
 - [PaginatedCloudBackupRestoreJob](./docs/PaginatedCloudBackupRestoreJob.md)
 - [PaginatedCloudBackupShardedClusterSnapshot](./docs/PaginatedCloudBackupShardedClusterSnapshot.md)
 - [PaginatedCloudProviderContainer](./docs/PaginatedCloudProviderContainer.md)
 - [PaginatedClusterDescription20240805](./docs/PaginatedClusterDescription20240805.md)
 - [PaginatedConnectedOrgConfigs](./docs/PaginatedConnectedOrgConfigs.md)
 - [PaginatedContainerPeer](./docs/PaginatedContainerPeer.md)
 - [PaginatedDatabase](./docs/PaginatedDatabase.md)
 - [PaginatedDiskPartition](./docs/PaginatedDiskPartition.md)
 - [PaginatedEventTypeDetailsResponse](./docs/PaginatedEventTypeDetailsResponse.md)
 - [PaginatedFederationIdentityProvider](./docs/PaginatedFederationIdentityProvider.md)
 - [PaginatedFlexClusters20241113](./docs/PaginatedFlexClusters20241113.md)
 - [PaginatedGroupServiceAccounts](./docs/PaginatedGroupServiceAccounts.md)
 - [PaginatedGroupUser](./docs/PaginatedGroupUser.md)
 - [PaginatedHostViewAtlas](./docs/PaginatedHostViewAtlas.md)
 - [PaginatedIntegration](./docs/PaginatedIntegration.md)
 - [PaginatedNetworkAccess](./docs/PaginatedNetworkAccess.md)
 - [PaginatedOnlineArchive](./docs/PaginatedOnlineArchive.md)
 - [PaginatedOrgGroup](./docs/PaginatedOrgGroup.md)
 - [PaginatedOrgServiceAccounts](./docs/PaginatedOrgServiceAccounts.md)
 - [PaginatedOrgUser](./docs/PaginatedOrgUser.md)
 - [PaginatedOrganization](./docs/PaginatedOrganization.md)
 - [PaginatedPipelineRun](./docs/PaginatedPipelineRun.md)
 - [PaginatedPrivateNetworkEndpointIdEntry](./docs/PaginatedPrivateNetworkEndpointIdEntry.md)
 - [PaginatedPublicApiUsageDetailsLineItem](./docs/PaginatedPublicApiUsageDetailsLineItem.md)
 - [PaginatedRestoreJob](./docs/PaginatedRestoreJob.md)
 - [PaginatedRoleMapping](./docs/PaginatedRoleMapping.md)
 - [PaginatedServerlessInstanceDescription](./docs/PaginatedServerlessInstanceDescription.md)
 - [PaginatedServiceAccountGroup](./docs/PaginatedServiceAccountGroup.md)
 - [PaginatedServiceAccountIPAccessEntry](./docs/PaginatedServiceAccountIPAccessEntry.md)
 - [PaginatedSnapshot](./docs/PaginatedSnapshot.md)
 - [PaginatedTeam](./docs/PaginatedTeam.md)
 - [PaginatedTeamRole](./docs/PaginatedTeamRole.md)
 - [PaginatedTenantRestore](./docs/PaginatedTenantRestore.md)
 - [PaginatedTenantSnapshot](./docs/PaginatedTenantSnapshot.md)
 - [PaginatedUserCert](./docs/PaginatedUserCert.md)
 - [PartitionField](./docs/PartitionField.md)
 - [PemFileInfo](./docs/PemFileInfo.md)
 - [PemFileInfoUpdate](./docs/PemFileInfoUpdate.md)
 - [PerformanceAdvisorIndex](./docs/PerformanceAdvisorIndex.md)
 - [PerformanceAdvisorOpStats](./docs/PerformanceAdvisorOpStats.md)
 - [PerformanceAdvisorOperation](./docs/PerformanceAdvisorOperation.md)
 - [PerformanceAdvisorResponse](./docs/PerformanceAdvisorResponse.md)
 - [PerformanceAdvisorShape](./docs/PerformanceAdvisorShape.md)
 - [PerformanceAdvisorSlowQuery](./docs/PerformanceAdvisorSlowQuery.md)
 - [PerformanceAdvisorSlowQueryList](./docs/PerformanceAdvisorSlowQueryList.md)
 - [PerformanceAdvisorSlowQueryMetrics](./docs/PerformanceAdvisorSlowQueryMetrics.md)
 - [PinFCV](./docs/PinFCV.md)
 - [PinnedNamespaces](./docs/PinnedNamespaces.md)
 - [PipelineRunStats](./docs/PipelineRunStats.md)
 - [PrivateEndpointHostname](./docs/PrivateEndpointHostname.md)
 - [PrivateIPMode](./docs/PrivateIPMode.md)
 - [PrivateLinkEndpoint](./docs/PrivateLinkEndpoint.md)
 - [PrivateNetworkEndpointIdEntry](./docs/PrivateNetworkEndpointIdEntry.md)
 - [ProjectSettingItem](./docs/ProjectSettingItem.md)
 - [ProtectedHours](./docs/ProtectedHours.md)
 - [PublicApiUsageDetailsLineItem](./docs/PublicApiUsageDetailsLineItem.md)
 - [PushBasedLogExportProject](./docs/PushBasedLogExportProject.md)
 - [Raw](./docs/Raw.md)
 - [RegionSpec](./docs/RegionSpec.md)
 - [ReplicationSpec20240805](./docs/ReplicationSpec20240805.md)
 - [ResourceTag](./docs/ResourceTag.md)
 - [RestoreJobFileHash](./docs/RestoreJobFileHash.md)
 - [SampleDatasetStatus](./docs/SampleDatasetStatus.md)
 - [SchemaAdvisorItemRecommendation](./docs/SchemaAdvisorItemRecommendation.md)
 - [SchemaAdvisorNamespaceTriggers](./docs/SchemaAdvisorNamespaceTriggers.md)
 - [SchemaAdvisorResponse](./docs/SchemaAdvisorResponse.md)
 - [SchemaAdvisorTriggerDetails](./docs/SchemaAdvisorTriggerDetails.md)
 - [SearchIndexCreateRequest](./docs/SearchIndexCreateRequest.md)
 - [SearchIndexDefinitionVersion](./docs/SearchIndexDefinitionVersion.md)
 - [SearchIndexResponse](./docs/SearchIndexResponse.md)
 - [SearchIndexUpdateRequest](./docs/SearchIndexUpdateRequest.md)
 - [SearchIndexUpdateRequestDefinition](./docs/SearchIndexUpdateRequestDefinition.md)
 - [SearchMappings](./docs/SearchMappings.md)
 - [SearchSynonymMappingDefinition](./docs/SearchSynonymMappingDefinition.md)
 - [ServerlessBackupRestoreJob](./docs/ServerlessBackupRestoreJob.md)
 - [ServerlessBackupSnapshot](./docs/ServerlessBackupSnapshot.md)
 - [ServerlessConnectionStringsPrivateEndpointItem](./docs/ServerlessConnectionStringsPrivateEndpointItem.md)
 - [ServerlessConnectionStringsPrivateEndpointList](./docs/ServerlessConnectionStringsPrivateEndpointList.md)
 - [ServerlessInstanceDescription](./docs/ServerlessInstanceDescription.md)
 - [ServerlessInstanceDescriptionConnectionStrings](./docs/ServerlessInstanceDescriptionConnectionStrings.md)
 - [ServerlessInstanceDescriptionCreate](./docs/ServerlessInstanceDescriptionCreate.md)
 - [ServerlessInstanceDescriptionUpdate](./docs/ServerlessInstanceDescriptionUpdate.md)
 - [ServerlessProviderSettings](./docs/ServerlessProviderSettings.md)
 - [ServerlessTenantCreateRequest](./docs/ServerlessTenantCreateRequest.md)
 - [ServerlessTenantEndpoint](./docs/ServerlessTenantEndpoint.md)
 - [ServerlessTenantEndpointUpdate](./docs/ServerlessTenantEndpointUpdate.md)
 - [ServiceAccountGroup](./docs/ServiceAccountGroup.md)
 - [ServiceAccountIPAccessListEntry](./docs/ServiceAccountIPAccessListEntry.md)
 - [ServiceAccountSecret](./docs/ServiceAccountSecret.md)
 - [ServiceAccountSecretRequest](./docs/ServiceAccountSecretRequest.md)
 - [ShardEntry](./docs/ShardEntry.md)
 - [ShardKeys](./docs/ShardKeys.md)
 - [ShardingRequest](./docs/ShardingRequest.md)
 - [Source](./docs/Source.md)
 - [StateReason](./docs/StateReason.md)
 - [StreamConfig](./docs/StreamConfig.md)
 - [StreamProcessorMetricThreshold](./docs/StreamProcessorMetricThreshold.md)
 - [StreamsAWSConnectionConfig](./docs/StreamsAWSConnectionConfig.md)
 - [StreamsConnection](./docs/StreamsConnection.md)
 - [StreamsDLQ](./docs/StreamsDLQ.md)
 - [StreamsDataProcessRegion](./docs/StreamsDataProcessRegion.md)
 - [StreamsKafkaAuthentication](./docs/StreamsKafkaAuthentication.md)
 - [StreamsKafkaNetworking](./docs/StreamsKafkaNetworking.md)
 - [StreamsKafkaNetworkingAccess](./docs/StreamsKafkaNetworkingAccess.md)
 - [StreamsKafkaSecurity](./docs/StreamsKafkaSecurity.md)
 - [StreamsMatcher](./docs/StreamsMatcher.md)
 - [StreamsModifyStreamProcessor](./docs/StreamsModifyStreamProcessor.md)
 - [StreamsModifyStreamProcessorOptions](./docs/StreamsModifyStreamProcessorOptions.md)
 - [StreamsOptions](./docs/StreamsOptions.md)
 - [StreamsPrivateLinkConnection](./docs/StreamsPrivateLinkConnection.md)
 - [StreamsProcessor](./docs/StreamsProcessor.md)
 - [StreamsProcessorWithStats](./docs/StreamsProcessorWithStats.md)
 - [StreamsSampleConnections](./docs/StreamsSampleConnections.md)
 - [StreamsStartStreamProcessorWith](./docs/StreamsStartStreamProcessorWith.md)
 - [StreamsTenant](./docs/StreamsTenant.md)
 - [SynonymMappingStatusDetail](./docs/SynonymMappingStatusDetail.md)
 - [SynonymSource](./docs/SynonymSource.md)
 - [SystemStatus](./docs/SystemStatus.md)
 - [TargetOrg](./docs/TargetOrg.md)
 - [TargetOrgRequest](./docs/TargetOrgRequest.md)
 - [Team](./docs/Team.md)
 - [TeamResponse](./docs/TeamResponse.md)
 - [TeamRole](./docs/TeamRole.md)
 - [TeamUpdate](./docs/TeamUpdate.md)
 - [TenantRestore](./docs/TenantRestore.md)
 - [ThirdPartyIntegration](./docs/ThirdPartyIntegration.md)
 - [TriggerIngestionPipelineRequest](./docs/TriggerIngestionPipelineRequest.md)
 - [UpdateAtlasOrganizationApiKey](./docs/UpdateAtlasOrganizationApiKey.md)
 - [UpdateAtlasProjectApiKey](./docs/UpdateAtlasProjectApiKey.md)
 - [UpdateCustomDBRole](./docs/UpdateCustomDBRole.md)
 - [UpdateGroupRolesForUser](./docs/UpdateGroupRolesForUser.md)
 - [UpdateOrgRolesForUser](./docs/UpdateOrgRolesForUser.md)
 - [UsageDetailsFilterRequest](./docs/UsageDetailsFilterRequest.md)
 - [UserAccessListRequest](./docs/UserAccessListRequest.md)
 - [UserAccessListResponse](./docs/UserAccessListResponse.md)
 - [UserAccessRoleAssignment](./docs/UserAccessRoleAssignment.md)
 - [UserCert](./docs/UserCert.md)
 - [UserCustomDBRole](./docs/UserCustomDBRole.md)
 - [UserScope](./docs/UserScope.md)
 - [UserSecurity](./docs/UserSecurity.md)
 - [UserToDNMapping](./docs/UserToDNMapping.md)
 - [VPCPeeringActionChallenge](./docs/VPCPeeringActionChallenge.md)
 - [VectorSearchHostStatusDetail](./docs/VectorSearchHostStatusDetail.md)
 - [VectorSearchIndexDefinition](./docs/VectorSearchIndexDefinition.md)
 - [VectorSearchIndexStatusDetail](./docs/VectorSearchIndexStatusDetail.md)
 - [X509Certificate](./docs/X509Certificate.md)
 - [X509CertificateUpdate](./docs/X509CertificateUpdate.md)
 - [ZoneMapping](./docs/ZoneMapping.md)




