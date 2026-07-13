# SDK Reference 

## Reference Documentation for SDK Endpoints

All URIs are relative to *https://cloud.mongodb.com*

Class        | Method        | HTTP request  | Description   | 
------------ | ------------- | ------------- | ------------- |
*AWSClustersDNSAPI* | [GetAwsCustomDns](./docs/AWSClustersDNSApi.md#getawscustomdns) | **Get** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Return One Custom DNS Configuration for Atlas Clusters on AWS |
*AWSClustersDNSAPI* | [ToggleAwsCustomDns](./docs/AWSClustersDNSApi.md#toggleawscustomdns) | **Patch** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Update State of One Custom DNS Configuration for Atlas Clusters on AWS |
*AccessTrackingAPI* | [GetAccessHistoryCluster](./docs/AccessTrackingApi.md#getaccesshistorycluster) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/clusters/{clusterName} | Return Database Access History for One Cluster by Cluster Name |
*AccessTrackingAPI* | [GetAccessHistoryProcess](./docs/AccessTrackingApi.md#getaccesshistoryprocess) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/processes/{hostname} | Return Database Access History for One Cluster by Hostname |
*ActivityFeedAPI* | [GetGroupActivityFeed](./docs/ActivityFeedApi.md#getgroupactivityfeed) | **Get** /api/atlas/v2/groups/{groupId}/activityFeed | Return Pre-Filtered Activity Feed Link for One Project |
*ActivityFeedAPI* | [GetOrgActivityFeed](./docs/ActivityFeedApi.md#getorgactivityfeed) | **Get** /api/atlas/v2/orgs/{orgId}/activityFeed | Return Pre-Filtered Activity Feed Link for One Organization |
*AlertConfigurationsAPI* | [CreateAlertConfig](./docs/AlertConfigurationsApi.md#createalertconfig) | **Post** /api/atlas/v2/groups/{groupId}/alertConfigs | Create One Alert Configuration in One Project |
*AlertConfigurationsAPI* | [DeleteAlertConfig](./docs/AlertConfigurationsApi.md#deletealertconfig) | **Delete** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Remove One Alert Configuration from One Project |
*AlertConfigurationsAPI* | [GetAlertConfig](./docs/AlertConfigurationsApi.md#getalertconfig) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Return One Alert Configuration from One Project |
*AlertConfigurationsAPI* | [GetAlertConfigs](./docs/AlertConfigurationsApi.md#getalertconfigs) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId}/alertConfigs | Return All Alert Configurations Set for One Alert |
*AlertConfigurationsAPI* | [ListAlertConfigs](./docs/AlertConfigurationsApi.md#listalertconfigs) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs | Return All Alert Configurations in One Project |
*AlertConfigurationsAPI* | [ListMatcherFieldNames](./docs/AlertConfigurationsApi.md#listmatcherfieldnames) | **Get** /api/atlas/v2/alertConfigs/matchers/fieldNames | Return All Alert Configuration Matchers Field Names |
*AlertConfigurationsAPI* | [ToggleAlertConfig](./docs/AlertConfigurationsApi.md#togglealertconfig) | **Patch** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Toggle State of One Alert Configuration in One Project |
*AlertConfigurationsAPI* | [UpdateAlertConfig](./docs/AlertConfigurationsApi.md#updatealertconfig) | **Put** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Update One Alert Configuration in One Project |
*AlertsAPI* | [AcknowledgeAlert](./docs/AlertsApi.md#acknowledgealert) | **Patch** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Acknowledge One Alert from One Project |
*AlertsAPI* | [GetAlert](./docs/AlertsApi.md#getalert) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Return One Alert from One Project |
*AlertsAPI* | [GetAlertConfigAlerts](./docs/AlertsApi.md#getalertconfigalerts) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}/alerts | Return All Open Alerts for One Alert Configuration |
*AlertsAPI* | [ListAlerts](./docs/AlertsApi.md#listalerts) | **Get** /api/atlas/v2/groups/{groupId}/alerts | Return All Alerts from One Project |
*AtlasSearchAPI* | [CreateClusterFtsIndex](./docs/AtlasSearchApi.md#createclusterftsindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index |
*AtlasSearchAPI* | [CreateClusterSearchDeployment](./docs/AtlasSearchApi.md#createclustersearchdeployment) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Create Search Nodes |
*AtlasSearchAPI* | [CreateClusterSearchIndex](./docs/AtlasSearchApi.md#createclustersearchindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Create One Atlas Search Index |
*AtlasSearchAPI* | [DeleteClusterFtsIndex](./docs/AtlasSearchApi.md#deleteclusterftsindex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index |
*AtlasSearchAPI* | [DeleteClusterSearchDeployment](./docs/AtlasSearchApi.md#deleteclustersearchdeployment) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Delete Search Nodes |
*AtlasSearchAPI* | [DeleteClusterSearchIndex](./docs/AtlasSearchApi.md#deleteclustersearchindex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Remove One Atlas Search Index by ID |
*AtlasSearchAPI* | [DeleteIndexByName](./docs/AtlasSearchApi.md#deleteindexbyname) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Remove One Atlas Search Index by Name |
*AtlasSearchAPI* | [GetClusterFtsIndex](./docs/AtlasSearchApi.md#getclusterftsindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index |
*AtlasSearchAPI* | [GetClusterSearchDeployment](./docs/AtlasSearchApi.md#getclustersearchdeployment) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Return All Search Nodes |
*AtlasSearchAPI* | [GetClusterSearchIndex](./docs/AtlasSearchApi.md#getclustersearchindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Return One Atlas Search Index by ID |
*AtlasSearchAPI* | [GetIndexByName](./docs/AtlasSearchApi.md#getindexbyname) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Return One Atlas Search Index by Name |
*AtlasSearchAPI* | [ListClusterFtsIndex](./docs/AtlasSearchApi.md#listclusterftsindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection |
*AtlasSearchAPI* | [ListClusterSearchIndexes](./docs/AtlasSearchApi.md#listclustersearchindexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Return All Atlas Search Indexes for One Cluster |
*AtlasSearchAPI* | [ListSearchIndex](./docs/AtlasSearchApi.md#listsearchindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection |
*AtlasSearchAPI* | [UpdateClusterFtsIndex](./docs/AtlasSearchApi.md#updateclusterftsindex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index |
*AtlasSearchAPI* | [UpdateClusterSearchDeployment](./docs/AtlasSearchApi.md#updateclustersearchdeployment) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Update Search Nodes |
*AtlasSearchAPI* | [UpdateClusterSearchIndex](./docs/AtlasSearchApi.md#updateclustersearchindex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Update One Atlas Search Index by ID |
*AtlasSearchAPI* | [UpdateIndexByName](./docs/AtlasSearchApi.md#updateindexbyname) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Update One Atlas Search Index by Name |
*AuditingAPI* | [GetGroupAuditLog](./docs/AuditingApi.md#getgroupauditlog) | **Get** /api/atlas/v2/groups/{groupId}/auditLog | Return Auditing Configuration for One Project |
*AuditingAPI* | [UpdateAuditLog](./docs/AuditingApi.md#updateauditlog) | **Patch** /api/atlas/v2/groups/{groupId}/auditLog | Update Auditing Configuration for One Project |
*CloudBackupsAPI* | [CancelBackupRestoreJob](./docs/CloudBackupsApi.md#cancelbackuprestorejob) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job for One Cluster |
*CloudBackupsAPI* | [CreateBackupExport](./docs/CloudBackupsApi.md#createbackupexport) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Snapshot Export Job |
*CloudBackupsAPI* | [CreateBackupPrivateEndpoint](./docs/CloudBackupsApi.md#createbackupprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/backup/{cloudProvider}/privateEndpoints | Create One Object Storage Private Endpoint for Cloud Backups for One Cloud Provider in One Project |
*CloudBackupsAPI* | [CreateBackupRestoreJob](./docs/CloudBackupsApi.md#createbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Create One Restore Job of One Cluster |
*CloudBackupsAPI* | [CreateExportBucket](./docs/CloudBackupsApi.md#createexportbucket) | **Post** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Create One Snapshot Export Bucket |
*CloudBackupsAPI* | [CreateServerlessRestoreJob](./docs/CloudBackupsApi.md#createserverlessrestorejob) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Create One Restore Job for One Serverless Instance |
*CloudBackupsAPI* | [DeleteBackupPrivateEndpoint](./docs/CloudBackupsApi.md#deletebackupprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/backup/{cloudProvider}/privateEndpoints/{endpointId} | Delete One Object Storage Private Endpoint for Cloud Backups for One Cloud Provider from One Project |
*CloudBackupsAPI* | [DeleteBackupShardedCluster](./docs/CloudBackupsApi.md#deletebackupshardedcluster) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup |
*CloudBackupsAPI* | [DeleteClusterBackupSchedule](./docs/CloudBackupsApi.md#deleteclusterbackupschedule) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules |
*CloudBackupsAPI* | [DeleteClusterBackupSnapshot](./docs/CloudBackupsApi.md#deleteclusterbackupsnapshot) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup |
*CloudBackupsAPI* | [DeleteExportBucket](./docs/CloudBackupsApi.md#deleteexportbucket) | **Delete** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Delete One Snapshot Export Bucket |
*CloudBackupsAPI* | [DisableCompliancePolicy](./docs/CloudBackupsApi.md#disablecompliancepolicy) | **Delete** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Disable Backup Compliance Policy Settings |
*CloudBackupsAPI* | [GetBackupExport](./docs/CloudBackupsApi.md#getbackupexport) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Snapshot Export Job |
*CloudBackupsAPI* | [GetBackupPrivateEndpoint](./docs/CloudBackupsApi.md#getbackupprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/backup/{cloudProvider}/privateEndpoints/{endpointId} | Return One Object Storage Private Endpoint for Cloud Backups for One Cloud Provider in One Project |
*CloudBackupsAPI* | [GetBackupRestoreJob](./docs/CloudBackupsApi.md#getbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Cluster |
*CloudBackupsAPI* | [GetBackupSchedule](./docs/CloudBackupsApi.md#getbackupschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule |
*CloudBackupsAPI* | [GetBackupShardedCluster](./docs/CloudBackupsApi.md#getbackupshardedcluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup |
*CloudBackupsAPI* | [GetClusterBackupSnapshot](./docs/CloudBackupsApi.md#getclusterbackupsnapshot) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup |
*CloudBackupsAPI* | [GetCompliancePolicy](./docs/CloudBackupsApi.md#getcompliancepolicy) | **Get** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Return Backup Compliance Policy Settings |
*CloudBackupsAPI* | [GetExportBucket](./docs/CloudBackupsApi.md#getexportbucket) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One Snapshot Export Bucket |
*CloudBackupsAPI* | [GetServerlessBackupSnapshot](./docs/CloudBackupsApi.md#getserverlessbackupsnapshot) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance |
*CloudBackupsAPI* | [GetServerlessRestoreJob](./docs/CloudBackupsApi.md#getserverlessrestorejob) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance |
*CloudBackupsAPI* | [ListBackupExports](./docs/CloudBackupsApi.md#listbackupexports) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Snapshot Export Jobs |
*CloudBackupsAPI* | [ListBackupPrivateEndpoints](./docs/CloudBackupsApi.md#listbackupprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/backup/{cloudProvider}/privateEndpoints | Return Object Storage Private Endpoints for Cloud Backups for One Cloud Provider in One Project |
*CloudBackupsAPI* | [ListBackupRestoreJobs](./docs/CloudBackupsApi.md#listbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster |
*CloudBackupsAPI* | [ListBackupShardedClusters](./docs/CloudBackupsApi.md#listbackupshardedclusters) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups |
*CloudBackupsAPI* | [ListBackupSnapshots](./docs/CloudBackupsApi.md#listbackupsnapshots) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups |
*CloudBackupsAPI* | [ListExportBuckets](./docs/CloudBackupsApi.md#listexportbuckets) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Return All Snapshot Export Buckets |
*CloudBackupsAPI* | [ListServerlessBackupSnapshots](./docs/CloudBackupsApi.md#listserverlessbackupsnapshots) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance |
*CloudBackupsAPI* | [ListServerlessRestoreJobs](./docs/CloudBackupsApi.md#listserverlessrestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance |
*CloudBackupsAPI* | [TakeSnapshots](./docs/CloudBackupsApi.md#takesnapshots) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot |
*CloudBackupsAPI* | [UpdateBackupExportBucket](./docs/CloudBackupsApi.md#updatebackupexportbucket) | **Patch** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Update One Export Bucket Private Networking Settings |
*CloudBackupsAPI* | [UpdateBackupSchedule](./docs/CloudBackupsApi.md#updatebackupschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster |
*CloudBackupsAPI* | [UpdateBackupSnapshot](./docs/CloudBackupsApi.md#updatebackupsnapshot) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Update Expiration Date for One Cloud Backup |
*CloudBackupsAPI* | [UpdateCompliancePolicy](./docs/CloudBackupsApi.md#updatecompliancepolicy) | **Put** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Update Backup Compliance Policy Settings |
*CloudMigrationServiceAPI* | [CreateGroupLiveMigration](./docs/CloudMigrationServiceApi.md#creategrouplivemigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations | Create One Migration for One Local Managed Cluster to MongoDB Atlas |
*CloudMigrationServiceAPI* | [CreateLinkToken](./docs/CloudMigrationServiceApi.md#createlinktoken) | **Post** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Create One Link-Token |
*CloudMigrationServiceAPI* | [CutoverMigration](./docs/CloudMigrationServiceApi.md#cutovermigration) | **Put** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover | Cut Over One Migrated Cluster |
*CloudMigrationServiceAPI* | [DeleteLinkTokens](./docs/CloudMigrationServiceApi.md#deletelinktokens) | **Delete** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Remove One Link-Token |
*CloudMigrationServiceAPI* | [GetGroupLiveMigration](./docs/CloudMigrationServiceApi.md#getgrouplivemigration) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId} | Return One Migration Job |
*CloudMigrationServiceAPI* | [GetMigrationValidateStatus](./docs/CloudMigrationServiceApi.md#getmigrationvalidatestatus) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/validate/{validationId} | Return One Migration Validation Job |
*CloudMigrationServiceAPI* | [ListAvailableProjects](./docs/CloudMigrationServiceApi.md#listavailableprojects) | **Get** /api/atlas/v2/orgs/{orgId}/liveMigrations/availableProjects | Return All Projects Available for Migration |
*CloudMigrationServiceAPI* | [ValidateLiveMigrations](./docs/CloudMigrationServiceApi.md#validatelivemigrations) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations/validate | Validate One Migration Request |
*CloudProviderAccessAPI* | [AuthorizeProviderAccessRole](./docs/CloudProviderAccessApi.md#authorizeprovideraccessrole) | **Patch** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role |
*CloudProviderAccessAPI* | [CreateCloudProviderAccess](./docs/CloudProviderAccessApi.md#createcloudprovideraccess) | **Post** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role |
*CloudProviderAccessAPI* | [DeauthorizeProviderAccessRole](./docs/CloudProviderAccessApi.md#deauthorizeprovideraccessrole) | **Delete** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role |
*CloudProviderAccessAPI* | [GetCloudProviderAccess](./docs/CloudProviderAccessApi.md#getcloudprovideraccess) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Return One Cloud Provider Access Role |
*CloudProviderAccessAPI* | [ListCloudProviderAccess](./docs/CloudProviderAccessApi.md#listcloudprovideraccess) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles |
*ClusterOutageSimulationAPI* | [EndOutageSimulation](./docs/ClusterOutageSimulationApi.md#endoutagesimulation) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | End One Outage Simulation |
*ClusterOutageSimulationAPI* | [GetOutageSimulation](./docs/ClusterOutageSimulationApi.md#getoutagesimulation) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Return One Outage Simulation |
*ClusterOutageSimulationAPI* | [StartOutageSimulation](./docs/ClusterOutageSimulationApi.md#startoutagesimulation) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Start One Outage Simulation |
*ClustersAPI* | [AutoScalingConfiguration](./docs/ClustersApi.md#autoscalingconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/autoScalingConfiguration | Return Auto Scaling Configuration for One Sharded Cluster |
*ClustersAPI* | [CreateCluster](./docs/ClustersApi.md#createcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters | Create One Cluster in One Project |
*ClustersAPI* | [DeleteCluster](./docs/ClustersApi.md#deletecluster) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Remove One Cluster from One Project |
*ClustersAPI* | [GetCluster](./docs/ClustersApi.md#getcluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Return One Cluster from One Project |
*ClustersAPI* | [GetClusterStatus](./docs/ClustersApi.md#getclusterstatus) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status | Return Status of All Cluster Operations |
*ClustersAPI* | [GetProcessArgs](./docs/ClustersApi.md#getprocessargs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Return Advanced Configuration Options for One Cluster |
*ClustersAPI* | [GetSampleDatasetLoad](./docs/ClustersApi.md#getsampledatasetload) | **Get** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId} | Return Status of Sample Dataset Load for One Cluster |
*ClustersAPI* | [GrantMongoEmployeeAccess](./docs/ClustersApi.md#grantmongoemployeeaccess) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:grantMongoDBEmployeeAccess | Grant MongoDB Employee Cluster Access for One Cluster |
*ClustersAPI* | [ListClusterDetails](./docs/ClustersApi.md#listclusterdetails) | **Get** /api/atlas/v2/clusters | Return All Authorized Clusters in All Projects |
*ClustersAPI* | [ListClusterProviderRegions](./docs/ClustersApi.md#listclusterproviderregions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/provider/regions | Return All Cloud Provider Regions |
*ClustersAPI* | [ListClusters](./docs/ClustersApi.md#listclusters) | **Get** /api/atlas/v2/groups/{groupId}/clusters | Return All Clusters in One Project |
*ClustersAPI* | [PinFeatureCompatibilityVersion](./docs/ClustersApi.md#pinfeaturecompatibilityversion) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:pinFeatureCompatibilityVersion | Pin Feature Compatibility Version for One Cluster in One Project |
*ClustersAPI* | [RequestSampleDatasetLoad](./docs/ClustersApi.md#requestsampledatasetload) | **Post** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name} | Load Sample Dataset into One Cluster |
*ClustersAPI* | [RestartPrimaries](./docs/ClustersApi.md#restartprimaries) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restartPrimaries | Test Failover for One Cluster |
*ClustersAPI* | [RevokeMongoEmployeeAccess](./docs/ClustersApi.md#revokemongoemployeeaccess) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:revokeMongoDBEmployeeAccess | Revoke MongoDB Employee Cluster Access for One Cluster |
*ClustersAPI* | [UnpinFeatureCompatibilityVersion](./docs/ClustersApi.md#unpinfeaturecompatibilityversion) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}:unpinFeatureCompatibilityVersion | Unpin Feature Compatibility Version for One Cluster in One Project |
*ClustersAPI* | [UpdateCluster](./docs/ClustersApi.md#updatecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Update One Cluster in One Project |
*ClustersAPI* | [UpdateProcessArgs](./docs/ClustersApi.md#updateprocessargs) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Update Advanced Configuration Options for One Cluster |
*ClustersAPI* | [UpgradeClusterToServerless](./docs/ClustersApi.md#upgradeclustertoserverless) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgradeToServerless | Upgrade One Shared-Tier Cluster to One Serverless Instance |
*ClustersAPI* | [UpgradeTenantUpgrade](./docs/ClustersApi.md#upgradetenantupgrade) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade | Upgrade One Shared-Tier Cluster |
*ClustersAPI* | [ValidateGroupClusterConfigurations](./docs/ClustersApi.md#validategroupclusterconfigurations) | **Post** /api/atlas/v2/groups/{groupId}/clusterConfigurations:validate | Validate One Cluster Configuration |
*CollectionLevelMetricsAPI* | [GetClusterNamespaces](./docs/CollectionLevelMetricsApi.md#getclusternamespaces) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/collStats/namespaces | Return Ranked Namespaces from One Cluster |
*CollectionLevelMetricsAPI* | [GetProcessNamespaces](./docs/CollectionLevelMetricsApi.md#getprocessnamespaces) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/collStats/namespaces | Return Ranked Namespaces from One Host |
*CollectionLevelMetricsAPI* | [ListCollStatMeasurements](./docs/CollectionLevelMetricsApi.md#listcollstatmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements | Return Cluster-Level Query Latency |
*CollectionLevelMetricsAPI* | [ListCollStatMetrics](./docs/CollectionLevelMetricsApi.md#listcollstatmetrics) | **Get** /api/atlas/v2/groups/{groupId}/collStats/metrics | Return All Metric Names |
*CollectionLevelMetricsAPI* | [ListPinnedNamespaces](./docs/CollectionLevelMetricsApi.md#listpinnednamespaces) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Return Pinned Namespaces |
*CollectionLevelMetricsAPI* | [ListProcessMeasurements](./docs/CollectionLevelMetricsApi.md#listprocessmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/{databaseName}/{collectionName}/collStats/measurements | Return Host-Level Query Latency |
*CollectionLevelMetricsAPI* | [PinNamespaces](./docs/CollectionLevelMetricsApi.md#pinnamespaces) | **Put** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Pin Namespaces |
*CollectionLevelMetricsAPI* | [UnpinNamespaces](./docs/CollectionLevelMetricsApi.md#unpinnamespaces) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/unpin | Unpin Namespaces |
*CollectionLevelMetricsAPI* | [UpdatePinnedNamespaces](./docs/CollectionLevelMetricsApi.md#updatepinnednamespaces) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Add Pinned Namespaces |
*CustomDatabaseRolesAPI* | [CreateCustomDbRole](./docs/CustomDatabaseRolesApi.md#createcustomdbrole) | **Post** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Create One Custom Role |
*CustomDatabaseRolesAPI* | [DeleteCustomDbRole](./docs/CustomDatabaseRolesApi.md#deletecustomdbrole) | **Delete** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project |
*CustomDatabaseRolesAPI* | [GetCustomDbRole](./docs/CustomDatabaseRolesApi.md#getcustomdbrole) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project |
*CustomDatabaseRolesAPI* | [ListCustomDbRoles](./docs/CustomDatabaseRolesApi.md#listcustomdbroles) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project |
*CustomDatabaseRolesAPI* | [UpdateCustomDbRole](./docs/CustomDatabaseRolesApi.md#updatecustomdbrole) | **Patch** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project |
*DataFederationAPI* | [CreateDataFederation](./docs/DataFederationApi.md#createdatafederation) | **Post** /api/atlas/v2/groups/{groupId}/dataFederation | Create One Federated Database Instance in One Project |
*DataFederationAPI* | [CreatePrivateEndpointId](./docs/DataFederationApi.md#createprivateendpointid) | **Post** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Create One Federated Database Instance and Online Archive Private Endpoint for One Project |
*DataFederationAPI* | [DeleteDataFederation](./docs/DataFederationApi.md#deletedatafederation) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Remove One Federated Database Instance from One Project |
*DataFederationAPI* | [DeleteDataFederationLimit](./docs/DataFederationApi.md#deletedatafederationlimit) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Delete One Query Limit for One Federated Database Instance |
*DataFederationAPI* | [DeletePrivateEndpointId](./docs/DataFederationApi.md#deleteprivateendpointid) | **Delete** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Remove One Federated Database Instance and Online Archive Private Endpoint from One Project |
*DataFederationAPI* | [DownloadFederationQueryLogs](./docs/DataFederationApi.md#downloadfederationquerylogs) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz | Download Query Logs for One Federated Database Instance |
*DataFederationAPI* | [GetDataFederation](./docs/DataFederationApi.md#getdatafederation) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Return One Federated Database Instance in One Project |
*DataFederationAPI* | [GetDataFederationLimit](./docs/DataFederationApi.md#getdatafederationlimit) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Return One Federated Database Instance Query Limit for One Project |
*DataFederationAPI* | [GetPrivateEndpointId](./docs/DataFederationApi.md#getprivateendpointid) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Return One Federated Database Instance and Online Archive Private Endpoint in One Project |
*DataFederationAPI* | [ListDataFederation](./docs/DataFederationApi.md#listdatafederation) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation | Return All Federated Database Instances in One Project |
*DataFederationAPI* | [ListDataFederationLimits](./docs/DataFederationApi.md#listdatafederationlimits) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits | Return All Query Limits for One Federated Database Instance |
*DataFederationAPI* | [ListPrivateEndpointIds](./docs/DataFederationApi.md#listprivateendpointids) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Return All Federated Database Instance and Online Archive Private Endpoints in One Project |
*DataFederationAPI* | [SetDataFederationLimit](./docs/DataFederationApi.md#setdatafederationlimit) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Configure One Query Limit for One Federated Database Instance |
*DataFederationAPI* | [UpdateDataFederation](./docs/DataFederationApi.md#updatedatafederation) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Update One Federated Database Instance in One Project |
*DataLakePipelinesAPI* | [CreatePipeline](./docs/DataLakePipelinesApi.md#createpipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines | Create One Data Lake Pipeline |
*DataLakePipelinesAPI* | [DeletePipeline](./docs/DataLakePipelinesApi.md#deletepipeline) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Remove One Data Lake Pipeline |
*DataLakePipelinesAPI* | [DeletePipelineRun](./docs/DataLakePipelinesApi.md#deletepipelinerun) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Delete One Pipeline Run Dataset |
*DataLakePipelinesAPI* | [GetAvailablePipelineSchedules](./docs/DataLakePipelinesApi.md#getavailablepipelineschedules) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSchedules | Return All Ingestion Schedules for One Data Lake Pipeline |
*DataLakePipelinesAPI* | [GetAvailablePipelineSnapshots](./docs/DataLakePipelinesApi.md#getavailablepipelinesnapshots) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots | Return All Backup Snapshots for One Data Lake Pipeline |
*DataLakePipelinesAPI* | [GetPipeline](./docs/DataLakePipelinesApi.md#getpipeline) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Return One Data Lake Pipeline |
*DataLakePipelinesAPI* | [GetPipelineRun](./docs/DataLakePipelinesApi.md#getpipelinerun) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Return One Data Lake Pipeline Run |
*DataLakePipelinesAPI* | [ListPipelineRuns](./docs/DataLakePipelinesApi.md#listpipelineruns) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs | Return All Data Lake Pipeline Runs in One Project |
*DataLakePipelinesAPI* | [ListPipelines](./docs/DataLakePipelinesApi.md#listpipelines) | **Get** /api/atlas/v2/groups/{groupId}/pipelines | Return All Data Lake Pipelines in One Project |
*DataLakePipelinesAPI* | [PausePipeline](./docs/DataLakePipelinesApi.md#pausepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/pause | Pause One Data Lake Pipeline |
*DataLakePipelinesAPI* | [ResumePipeline](./docs/DataLakePipelinesApi.md#resumepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/resume | Resume One Data Lake Pipeline |
*DataLakePipelinesAPI* | [TriggerPipeline](./docs/DataLakePipelinesApi.md#triggerpipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/trigger | Trigger On-Demand Snapshot Ingestion |
*DataLakePipelinesAPI* | [UpdatePipeline](./docs/DataLakePipelinesApi.md#updatepipeline) | **Patch** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Update One Data Lake Pipeline |
*DatabaseUsersAPI* | [CreateDatabaseUser](./docs/DatabaseUsersApi.md#createdatabaseuser) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers | Create One Database User in One Project |
*DatabaseUsersAPI* | [DeleteDatabaseUser](./docs/DatabaseUsersApi.md#deletedatabaseuser) | **Delete** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Remove One Database User from One Project |
*DatabaseUsersAPI* | [GetDatabaseUser](./docs/DatabaseUsersApi.md#getdatabaseuser) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Return One Database User from One Project |
*DatabaseUsersAPI* | [ListDatabaseUsers](./docs/DatabaseUsersApi.md#listdatabaseusers) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers | Return All Database Users in One Project |
*DatabaseUsersAPI* | [UpdateDatabaseUser](./docs/DatabaseUsersApi.md#updatedatabaseuser) | **Patch** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Update One Database User in One Project |
*EncryptionAtRestUsingCustomerKeyManagementAPI* | [CreateRestPrivateEndpoint](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#createrestprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints | Create One Private Endpoint for Encryption at Rest Using Customer Key Management for One Cloud Provider in One Project |
*EncryptionAtRestUsingCustomerKeyManagementAPI* | [GetEncryptionAtRest](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#getencryptionatrest) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest Using Customer-Managed Keys for One Project |
*EncryptionAtRestUsingCustomerKeyManagementAPI* | [GetRestPrivateEndpoint](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#getrestprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints/{endpointId} | Return One Private Endpoint for Encryption at Rest Using Customer Key Management for One Cloud Provider in One Project |
*EncryptionAtRestUsingCustomerKeyManagementAPI* | [ListRestPrivateEndpoints](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#listrestprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints | Return Private Endpoints for Encryption at Rest Using Customer Key Management for One Cloud Provider in One Project |
*EncryptionAtRestUsingCustomerKeyManagementAPI* | [RequestPrivateEndpointDeletion](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#requestprivateendpointdeletion) | **Delete** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints/{endpointId} | Delete One Private Endpoint for Encryption at Rest Using Customer Key Management for One Cloud Provider from One Project |
*EncryptionAtRestUsingCustomerKeyManagementAPI* | [UpdateEncryptionAtRest](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#updateencryptionatrest) | **Patch** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Update Encryption at Rest Configuration in One Project |
*EventsAPI* | [GetGroupEvent](./docs/EventsApi.md#getgroupevent) | **Get** /api/atlas/v2/groups/{groupId}/events/{eventId} | Return One Event from One Project |
*EventsAPI* | [GetOrgEvent](./docs/EventsApi.md#getorgevent) | **Get** /api/atlas/v2/orgs/{orgId}/events/{eventId} | Return One Event from One Organization |
*EventsAPI* | [ListEventTypes](./docs/EventsApi.md#listeventtypes) | **Get** /api/atlas/v2/eventTypes | Return All Event Types |
*EventsAPI* | [ListGroupEvents](./docs/EventsApi.md#listgroupevents) | **Get** /api/atlas/v2/groups/{groupId}/events | Return Events from One Project |
*EventsAPI* | [ListOrgEvents](./docs/EventsApi.md#listorgevents) | **Get** /api/atlas/v2/orgs/{orgId}/events | Return Events from One Organization |
*FederatedAuthenticationAPI* | [CreateIdentityProvider](./docs/FederatedAuthenticationApi.md#createidentityprovider) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Create One Identity Provider |
*FederatedAuthenticationAPI* | [CreateRoleMapping](./docs/FederatedAuthenticationApi.md#createrolemapping) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Create One Role Mapping in One Organization Configuration |
*FederatedAuthenticationAPI* | [DeleteFederationSetting](./docs/FederatedAuthenticationApi.md#deletefederationsetting) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId} | Delete One Federation Settings Instance |
*FederatedAuthenticationAPI* | [DeleteIdentityProvider](./docs/FederatedAuthenticationApi.md#deleteidentityprovider) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Delete One Identity Provider |
*FederatedAuthenticationAPI* | [DeleteRoleMapping](./docs/FederatedAuthenticationApi.md#deleterolemapping) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization |
*FederatedAuthenticationAPI* | [GetConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#getconnectedorgconfig) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Organization Configuration from One Federation |
*FederatedAuthenticationAPI* | [GetFederationSettings](./docs/FederatedAuthenticationApi.md#getfederationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization |
*FederatedAuthenticationAPI* | [GetIdentityProvider](./docs/FederatedAuthenticationApi.md#getidentityprovider) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return One Identity Provider by ID |
*FederatedAuthenticationAPI* | [GetIdentityProviderMetadata](./docs/FederatedAuthenticationApi.md#getidentityprovidermetadata) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return Metadata of One Identity Provider |
*FederatedAuthenticationAPI* | [GetRoleMapping](./docs/FederatedAuthenticationApi.md#getrolemapping) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization |
*FederatedAuthenticationAPI* | [ListConnectedOrgConfigs](./docs/FederatedAuthenticationApi.md#listconnectedorgconfigs) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Organization Configurations from One Federation |
*FederatedAuthenticationAPI* | [ListIdentityProviders](./docs/FederatedAuthenticationApi.md#listidentityproviders) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Return All Identity Providers in One Federation |
*FederatedAuthenticationAPI* | [ListRoleMappings](./docs/FederatedAuthenticationApi.md#listrolemappings) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization |
*FederatedAuthenticationAPI* | [RemoveConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#removeconnectedorgconfig) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Organization Configuration from One Federation |
*FederatedAuthenticationAPI* | [RevokeIdentityProviderJwks](./docs/FederatedAuthenticationApi.md#revokeidentityproviderjwks) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/jwks | Revoke JWKS from One OIDC Identity Provider |
*FederatedAuthenticationAPI* | [UpdateConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#updateconnectedorgconfig) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update One Organization Configuration in One Federation |
*FederatedAuthenticationAPI* | [UpdateIdentityProvider](./docs/FederatedAuthenticationApi.md#updateidentityprovider) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update One Identity Provider |
*FederatedAuthenticationAPI* | [UpdateRoleMapping](./docs/FederatedAuthenticationApi.md#updaterolemapping) | **Put** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization |
*FlexClustersAPI* | [CreateFlexCluster](./docs/FlexClustersApi.md#createflexcluster) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters | Create One Flex Cluster in One Project |
*FlexClustersAPI* | [DeleteFlexCluster](./docs/FlexClustersApi.md#deleteflexcluster) | **Delete** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Remove One Flex Cluster from One Project |
*FlexClustersAPI* | [GetFlexCluster](./docs/FlexClustersApi.md#getflexcluster) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Return One Flex Cluster from One Project |
*FlexClustersAPI* | [ListFlexClusters](./docs/FlexClustersApi.md#listflexclusters) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters | Return All Flex Clusters from One Project |
*FlexClustersAPI* | [TenantUpgrade](./docs/FlexClustersApi.md#tenantupgrade) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters:tenantUpgrade | Upgrade One Flex Cluster |
*FlexClustersAPI* | [UpdateFlexCluster](./docs/FlexClustersApi.md#updateflexcluster) | **Patch** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Update One Flex Cluster in One Project |
*FlexRestoreJobsAPI* | [CreateFlexRestoreJob](./docs/FlexRestoreJobsApi.md#createflexrestorejob) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs | Create One Restore Job for One Flex Cluster |
*FlexRestoreJobsAPI* | [GetFlexRestoreJob](./docs/FlexRestoreJobsApi.md#getflexrestorejob) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Flex Cluster |
*FlexRestoreJobsAPI* | [ListFlexRestoreJobs](./docs/FlexRestoreJobsApi.md#listflexrestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs | Return All Restore Jobs for One Flex Cluster |
*FlexSnapshotsAPI* | [DownloadFlexBackup](./docs/FlexSnapshotsApi.md#downloadflexbackup) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/download | Download One Flex Cluster Snapshot |
*FlexSnapshotsAPI* | [GetFlexBackupSnapshot](./docs/FlexSnapshotsApi.md#getflexbackupsnapshot) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/snapshots/{snapshotId} | Return One Snapshot for One Flex Cluster |
*FlexSnapshotsAPI* | [ListFlexBackupSnapshots](./docs/FlexSnapshotsApi.md#listflexbackupsnapshots) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/snapshots | Return All Snapshots for One Flex Cluster |
*GlobalClustersAPI* | [CreateCustomZoneMapping](./docs/GlobalClustersApi.md#createcustomzonemapping) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Add One Custom Zone Mapping to One Global Cluster |
*GlobalClustersAPI* | [CreateManagedNamespace](./docs/GlobalClustersApi.md#createmanagednamespace) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Create One Managed Namespace in One Global Cluster |
*GlobalClustersAPI* | [DeleteCustomZoneMapping](./docs/GlobalClustersApi.md#deletecustomzonemapping) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Cluster |
*GlobalClustersAPI* | [DeleteManagedNamespaces](./docs/GlobalClustersApi.md#deletemanagednamespaces) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Cluster |
*GlobalClustersAPI* | [GetClusterGlobalWrites](./docs/GlobalClustersApi.md#getclusterglobalwrites) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites | Return One Managed Namespace in One Global Cluster |
*InvoicesAPI* | [CreateCostExplorerProcess](./docs/InvoicesApi.md#createcostexplorerprocess) | **Post** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage | Create One Cost Explorer Query Process |
*InvoicesAPI* | [GetCostExplorerUsage](./docs/InvoicesApi.md#getcostexplorerusage) | **Get** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage/{token} | Return Usage Details for One Cost Explorer Query |
*InvoicesAPI* | [GetInvoice](./docs/InvoicesApi.md#getinvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId} | Return One Invoice for One Organization |
*InvoicesAPI* | [GetInvoiceCsv](./docs/InvoicesApi.md#getinvoicecsv) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Invoice as CSV for One Organization |
*InvoicesAPI* | [GetSku](./docs/InvoicesApi.md#getsku) | **Get** /api/atlas/v2/skus/{skuId} | Return One Stock Keeping Unit |
*InvoicesAPI* | [ListInvoicePending](./docs/InvoicesApi.md#listinvoicepending) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization |
*InvoicesAPI* | [ListInvoices](./docs/InvoicesApi.md#listinvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices | Return All Invoices for One Organization |
*InvoicesAPI* | [ListSkus](./docs/InvoicesApi.md#listskus) | **Get** /api/atlas/v2/skus | Return All Stock Keeping Units |
*InvoicesAPI* | [SearchInvoiceLineItems](./docs/InvoicesApi.md#searchinvoicelineitems) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/lineItems:search | Return All Line Items for One Invoice by Invoice ID |
*LDAPConfigurationAPI* | [DeleteLdapUserMapping](./docs/LDAPConfigurationApi.md#deleteldapusermapping) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove LDAP User to DN Mapping |
*LDAPConfigurationAPI* | [GetUserSecurity](./docs/LDAPConfigurationApi.md#getusersecurity) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return LDAP or X.509 Configuration |
*LDAPConfigurationAPI* | [GetUserSecurityVerify](./docs/LDAPConfigurationApi.md#getusersecurityverify) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return Status of LDAP Configuration Verification in One Project |
*LDAPConfigurationAPI* | [UpdateUserSecurity](./docs/LDAPConfigurationApi.md#updateusersecurity) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Update LDAP or X.509 Configuration |
*LDAPConfigurationAPI* | [VerifyUserSecurityLdap](./docs/LDAPConfigurationApi.md#verifyusersecurityldap) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify LDAP Configuration in One Project |
*LegacyBackupAPI* | [CreateClusterRestoreJob](./docs/LegacyBackupApi.md#createclusterrestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Create One Legacy Backup Restore Job |
*LegacyBackupAPI* | [DeleteClusterSnapshot](./docs/LegacyBackupApi.md#deleteclustersnapshot) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Remove One Legacy Backup Snapshot |
*LegacyBackupAPI* | [GetClusterBackupCheckpoint](./docs/LegacyBackupApi.md#getclusterbackupcheckpoint) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId} | Return One Legacy Backup Checkpoint |
*LegacyBackupAPI* | [GetClusterRestoreJob](./docs/LegacyBackupApi.md#getclusterrestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId} | Return One Legacy Backup Restore Job |
*LegacyBackupAPI* | [GetClusterSnapshot](./docs/LegacyBackupApi.md#getclustersnapshot) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Return One Legacy Backup Snapshot |
*LegacyBackupAPI* | [GetClusterSnapshotSchedule](./docs/LegacyBackupApi.md#getclustersnapshotschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Return One Snapshot Schedule |
*LegacyBackupAPI* | [ListClusterBackupCheckpoints](./docs/LegacyBackupApi.md#listclusterbackupcheckpoints) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints | Return All Legacy Backup Checkpoints |
*LegacyBackupAPI* | [ListClusterRestoreJobs](./docs/LegacyBackupApi.md#listclusterrestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Return All Legacy Backup Restore Jobs |
*LegacyBackupAPI* | [ListClusterSnapshots](./docs/LegacyBackupApi.md#listclustersnapshots) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots | Return All Legacy Backup Snapshots |
*LegacyBackupAPI* | [UpdateClusterSnapshot](./docs/LegacyBackupApi.md#updateclustersnapshot) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Update Expiration Date for One Legacy Backup Snapshot |
*LegacyBackupAPI* | [UpdateClusterSnapshotSchedule](./docs/LegacyBackupApi.md#updateclustersnapshotschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Update Snapshot Schedule for One Cluster |
*MaintenanceWindowsAPI* | [DeferMaintenanceWindow](./docs/MaintenanceWindowsApi.md#defermaintenancewindow) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/defer | Defer One Maintenance Window for One Project |
*MaintenanceWindowsAPI* | [GetMaintenanceWindow](./docs/MaintenanceWindowsApi.md#getmaintenancewindow) | **Get** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Return One Maintenance Window for One Project |
*MaintenanceWindowsAPI* | [ResetMaintenanceWindow](./docs/MaintenanceWindowsApi.md#resetmaintenancewindow) | **Delete** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Reset One Maintenance Window for One Project |
*MaintenanceWindowsAPI* | [ToggleMaintenanceAutoDefer](./docs/MaintenanceWindowsApi.md#togglemaintenanceautodefer) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/autoDefer | Toggle Automatic Deferral of Maintenance for One Project |
*MaintenanceWindowsAPI* | [UpdateMaintenanceWindow](./docs/MaintenanceWindowsApi.md#updatemaintenancewindow) | **Patch** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Update One Maintenance Window for One Project |
*MongoDBCloudUsersAPI* | [AddGroupUserRole](./docs/MongoDBCloudUsersApi.md#addgroupuserrole) | **Post** /api/atlas/v2/groups/{groupId}/users/{userId}:addRole | Add One Project Role to One MongoDB Cloud User |
*MongoDBCloudUsersAPI* | [AddGroupUsers](./docs/MongoDBCloudUsersApi.md#addgroupusers) | **Post** /api/atlas/v2/groups/{groupId}/users | Add One MongoDB Cloud User to One Project |
*MongoDBCloudUsersAPI* | [AddOrgRole](./docs/MongoDBCloudUsersApi.md#addorgrole) | **Post** /api/atlas/v2/orgs/{orgId}/users/{userId}:addRole | Add One Organization Role to One MongoDB Cloud User |
*MongoDBCloudUsersAPI* | [AddOrgTeamUser](./docs/MongoDBCloudUsersApi.md#addorgteamuser) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}:addUser | Add One MongoDB Cloud User to One Team |
*MongoDBCloudUsersAPI* | [CreateOrgUser](./docs/MongoDBCloudUsersApi.md#createorguser) | **Post** /api/atlas/v2/orgs/{orgId}/users | Add One MongoDB Cloud User to One Organization |
*MongoDBCloudUsersAPI* | [CreateUser](./docs/MongoDBCloudUsersApi.md#createuser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User |
*MongoDBCloudUsersAPI* | [GetGroupUser](./docs/MongoDBCloudUsersApi.md#getgroupuser) | **Get** /api/atlas/v2/groups/{groupId}/users/{userId} | Return One MongoDB Cloud User in One Project |
*MongoDBCloudUsersAPI* | [GetOrgUser](./docs/MongoDBCloudUsersApi.md#getorguser) | **Get** /api/atlas/v2/orgs/{orgId}/users/{userId} | Return One MongoDB Cloud User in One Organization |
*MongoDBCloudUsersAPI* | [GetUser](./docs/MongoDBCloudUsersApi.md#getuser) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User by ID |
*MongoDBCloudUsersAPI* | [GetUserByName](./docs/MongoDBCloudUsersApi.md#getuserbyname) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User by Username |
*MongoDBCloudUsersAPI* | [ListGroupUsers](./docs/MongoDBCloudUsersApi.md#listgroupusers) | **Get** /api/atlas/v2/groups/{groupId}/users | Return All MongoDB Cloud Users in One Project |
*MongoDBCloudUsersAPI* | [ListOrgUsers](./docs/MongoDBCloudUsersApi.md#listorgusers) | **Get** /api/atlas/v2/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization |
*MongoDBCloudUsersAPI* | [ListTeamUsers](./docs/MongoDBCloudUsersApi.md#listteamusers) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team |
*MongoDBCloudUsersAPI* | [RemoveGroupUser](./docs/MongoDBCloudUsersApi.md#removegroupuser) | **Delete** /api/atlas/v2/groups/{groupId}/users/{userId} | Remove One MongoDB Cloud User from One Project |
*MongoDBCloudUsersAPI* | [RemoveGroupUserRole](./docs/MongoDBCloudUsersApi.md#removegroupuserrole) | **Post** /api/atlas/v2/groups/{groupId}/users/{userId}:removeRole | Remove One Project Role from One MongoDB Cloud User |
*MongoDBCloudUsersAPI* | [RemoveOrgRole](./docs/MongoDBCloudUsersApi.md#removeorgrole) | **Post** /api/atlas/v2/orgs/{orgId}/users/{userId}:removeRole | Remove One Organization Role from One MongoDB Cloud User |
*MongoDBCloudUsersAPI* | [RemoveOrgTeamUser](./docs/MongoDBCloudUsersApi.md#removeorgteamuser) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}:removeUser | Remove One MongoDB Cloud User from One Team |
*MongoDBCloudUsersAPI* | [RemoveOrgUser](./docs/MongoDBCloudUsersApi.md#removeorguser) | **Delete** /api/atlas/v2/orgs/{orgId}/users/{userId} | Remove One MongoDB Cloud User from One Organization |
*MongoDBCloudUsersAPI* | [UpdateOrgUser](./docs/MongoDBCloudUsersApi.md#updateorguser) | **Patch** /api/atlas/v2/orgs/{orgId}/users/{userId} | Update One MongoDB Cloud User in One Organization |
*MonitoringAndLogsAPI* | [DownloadClusterLog](./docs/MonitoringAndLogsApi.md#downloadclusterlog) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Cluster Host in One Project |
*MonitoringAndLogsAPI* | [GetDatabase](./docs/MonitoringAndLogsApi.md#getdatabase) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName} | Return One Database for One MongoDB Process |
*MonitoringAndLogsAPI* | [GetDatabaseMeasurements](./docs/MonitoringAndLogsApi.md#getdatabasemeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements | Return Measurements for One Database in One MongoDB Process |
*MonitoringAndLogsAPI* | [GetGroupProcess](./docs/MonitoringAndLogsApi.md#getgroupprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId} | Return One MongoDB Process by ID |
*MonitoringAndLogsAPI* | [GetIndexMeasurements](./docs/MonitoringAndLogsApi.md#getindexmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements | Return Atlas Search Metrics for One Index in One Namespace |
*MonitoringAndLogsAPI* | [GetProcessDisk](./docs/MonitoringAndLogsApi.md#getprocessdisk) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName} | Return Measurements for One Disk |
*MonitoringAndLogsAPI* | [GetProcessDiskMeasurements](./docs/MonitoringAndLogsApi.md#getprocessdiskmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements | Return Measurements of One Disk for One MongoDB Process |
*MonitoringAndLogsAPI* | [GetProcessMeasurements](./docs/MonitoringAndLogsApi.md#getprocessmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/measurements | Return Measurements for One MongoDB Process |
*MonitoringAndLogsAPI* | [ListDatabases](./docs/MonitoringAndLogsApi.md#listdatabases) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases | Return Available Databases for One MongoDB Process |
*MonitoringAndLogsAPI* | [ListGroupProcesses](./docs/MonitoringAndLogsApi.md#listgroupprocesses) | **Get** /api/atlas/v2/groups/{groupId}/processes | Return All MongoDB Processes in One Project |
*MonitoringAndLogsAPI* | [ListHostFtsMetrics](./docs/MonitoringAndLogsApi.md#listhostftsmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics | Return All Atlas Search Metric Types for One Process |
*MonitoringAndLogsAPI* | [ListIndexMeasurements](./docs/MonitoringAndLogsApi.md#listindexmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements | Return All Atlas Search Index Metrics for One Namespace |
*MonitoringAndLogsAPI* | [ListMeasurements](./docs/MonitoringAndLogsApi.md#listmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/measurements | Return Atlas Search Hardware and Status Metrics |
*MonitoringAndLogsAPI* | [ListProcessDisks](./docs/MonitoringAndLogsApi.md#listprocessdisks) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks | Return Available Disks for One MongoDB Process |
*NetworkPeeringAPI* | [CreateGroupContainer](./docs/NetworkPeeringApi.md#creategroupcontainer) | **Post** /api/atlas/v2/groups/{groupId}/containers | Create One Network Peering Container |
*NetworkPeeringAPI* | [CreateGroupPeer](./docs/NetworkPeeringApi.md#creategrouppeer) | **Post** /api/atlas/v2/groups/{groupId}/peers | Create One Network Peering Connection |
*NetworkPeeringAPI* | [DeleteGroupContainer](./docs/NetworkPeeringApi.md#deletegroupcontainer) | **Delete** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container |
*NetworkPeeringAPI* | [DeleteGroupPeer](./docs/NetworkPeeringApi.md#deletegrouppeer) | **Delete** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Remove One Network Peering Connection |
*NetworkPeeringAPI* | [DisablePeering](./docs/NetworkPeeringApi.md#disablepeering) | **Patch** /api/atlas/v2/groups/{groupId}/privateIpMode | Disable Connect via Peering-Only Mode for One Project |
*NetworkPeeringAPI* | [GetGroupContainer](./docs/NetworkPeeringApi.md#getgroupcontainer) | **Get** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Return One Network Peering Container |
*NetworkPeeringAPI* | [GetGroupPeer](./docs/NetworkPeeringApi.md#getgrouppeer) | **Get** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project |
*NetworkPeeringAPI* | [ListGroupContainerAll](./docs/NetworkPeeringApi.md#listgroupcontainerall) | **Get** /api/atlas/v2/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project |
*NetworkPeeringAPI* | [ListGroupContainers](./docs/NetworkPeeringApi.md#listgroupcontainers) | **Get** /api/atlas/v2/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider |
*NetworkPeeringAPI* | [ListGroupPeers](./docs/NetworkPeeringApi.md#listgrouppeers) | **Get** /api/atlas/v2/groups/{groupId}/peers | Return All Network Peering Connections in One Project |
*NetworkPeeringAPI* | [UpdateGroupContainer](./docs/NetworkPeeringApi.md#updategroupcontainer) | **Patch** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Update One Network Peering Container |
*NetworkPeeringAPI* | [UpdateGroupPeer](./docs/NetworkPeeringApi.md#updategrouppeer) | **Patch** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Update One Network Peering Connection |
*NetworkPeeringAPI* | [VerifyPrivateIpMode](./docs/NetworkPeeringApi.md#verifyprivateipmode) | **Get** /api/atlas/v2/groups/{groupId}/privateIpMode | Verify Connect via Peering-Only Mode for One Project |
*OnlineArchiveAPI* | [CreateOnlineArchive](./docs/OnlineArchiveApi.md#createonlinearchive) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive |
*OnlineArchiveAPI* | [DeleteOnlineArchive](./docs/OnlineArchiveApi.md#deleteonlinearchive) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive |
*OnlineArchiveAPI* | [DownloadQueryLogs](./docs/OnlineArchiveApi.md#downloadquerylogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs |
*OnlineArchiveAPI* | [GetOnlineArchive](./docs/OnlineArchiveApi.md#getonlinearchive) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive |
*OnlineArchiveAPI* | [ListOnlineArchives](./docs/OnlineArchiveApi.md#listonlinearchives) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster |
*OnlineArchiveAPI* | [UpdateOnlineArchive](./docs/OnlineArchiveApi.md#updateonlinearchive) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive |
*OrganizationsAPI* | [CreateOrg](./docs/OrganizationsApi.md#createorg) | **Post** /api/atlas/v2/orgs | Create One Organization |
*OrganizationsAPI* | [CreateOrgInvite](./docs/OrganizationsApi.md#createorginvite) | **Post** /api/atlas/v2/orgs/{orgId}/invites | Create Invitation for One MongoDB Cloud User in One Organization |
*OrganizationsAPI* | [DeleteOrg](./docs/OrganizationsApi.md#deleteorg) | **Delete** /api/atlas/v2/orgs/{orgId} | Remove One Organization |
*OrganizationsAPI* | [DeleteOrgInvite](./docs/OrganizationsApi.md#deleteorginvite) | **Delete** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Remove One Invitation from One Organization |
*OrganizationsAPI* | [GetOrg](./docs/OrganizationsApi.md#getorg) | **Get** /api/atlas/v2/orgs/{orgId} | Return One Organization |
*OrganizationsAPI* | [GetOrgGroups](./docs/OrganizationsApi.md#getorggroups) | **Get** /api/atlas/v2/orgs/{orgId}/groups | Return All Projects in One Organization |
*OrganizationsAPI* | [GetOrgInvite](./docs/OrganizationsApi.md#getorginvite) | **Get** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Return One Invitation in One Organization by Invitation ID |
*OrganizationsAPI* | [GetOrgSettings](./docs/OrganizationsApi.md#getorgsettings) | **Get** /api/atlas/v2/orgs/{orgId}/settings | Return Settings for One Organization |
*OrganizationsAPI* | [ListOrgInvites](./docs/OrganizationsApi.md#listorginvites) | **Get** /api/atlas/v2/orgs/{orgId}/invites | Return All Invitations in One Organization |
*OrganizationsAPI* | [ListOrgs](./docs/OrganizationsApi.md#listorgs) | **Get** /api/atlas/v2/orgs | Return All Organizations |
*OrganizationsAPI* | [UpdateOrg](./docs/OrganizationsApi.md#updateorg) | **Patch** /api/atlas/v2/orgs/{orgId} | Update One Organization |
*OrganizationsAPI* | [UpdateOrgInviteById](./docs/OrganizationsApi.md#updateorginvitebyid) | **Patch** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Update One Invitation in One Organization by Invitation ID |
*OrganizationsAPI* | [UpdateOrgInvites](./docs/OrganizationsApi.md#updateorginvites) | **Patch** /api/atlas/v2/orgs/{orgId}/invites | Update One Invitation in One Organization |
*OrganizationsAPI* | [UpdateOrgSettings](./docs/OrganizationsApi.md#updateorgsettings) | **Patch** /api/atlas/v2/orgs/{orgId}/settings | Update Settings for One Organization |
*OrganizationsAPI* | [UpdateOrgUserRoles](./docs/OrganizationsApi.md#updateorguserroles) | **Put** /api/atlas/v2/orgs/{orgId}/users/{userId}/roles | Update Organization Roles for One MongoDB Cloud User |
*PerformanceAdvisorAPI* | [DisableManagedSlowMs](./docs/PerformanceAdvisorApi.md#disablemanagedslowms) | **Delete** /api/atlas/v2/groups/{groupId}/managedSlowMs/disable | Disable Managed Slow Operation Threshold |
*PerformanceAdvisorAPI* | [EnableManagedSlowMs](./docs/PerformanceAdvisorApi.md#enablemanagedslowms) | **Post** /api/atlas/v2/groups/{groupId}/managedSlowMs/enable | Enable Managed Slow Operation Threshold |
*PerformanceAdvisorAPI* | [GetManagedSlowMs](./docs/PerformanceAdvisorApi.md#getmanagedslowms) | **Get** /api/atlas/v2/groups/{groupId}/managedSlowMs | Return Managed Slow Operation Threshold Status |
*PerformanceAdvisorAPI* | [GetServerlessAutoIndexing](./docs/PerformanceAdvisorApi.md#getserverlessautoindexing) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing | Return Serverless Auto-Indexing Status |
*PerformanceAdvisorAPI* | [ListClusterSuggestedIndexes](./docs/PerformanceAdvisorApi.md#listclustersuggestedindexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/suggestedIndexes | Return All Suggested Indexes |
*PerformanceAdvisorAPI* | [ListDropIndexSuggestions](./docs/PerformanceAdvisorApi.md#listdropindexsuggestions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/dropIndexSuggestions | Return All Suggested Indexes to Drop |
*PerformanceAdvisorAPI* | [ListPerformanceAdvisorNamespaces](./docs/PerformanceAdvisorApi.md#listperformanceadvisornamespaces) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host |
*PerformanceAdvisorAPI* | [ListSchemaAdvice](./docs/PerformanceAdvisorApi.md#listschemaadvice) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/schemaAdvice | Return Schema Advice |
*PerformanceAdvisorAPI* | [ListSlowQueryLogs](./docs/PerformanceAdvisorApi.md#listslowquerylogs) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries |
*PerformanceAdvisorAPI* | [ListSuggestedIndexes](./docs/PerformanceAdvisorApi.md#listsuggestedindexes) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return All Suggested Indexes |
*PerformanceAdvisorAPI* | [SetServerlessAutoIndexing](./docs/PerformanceAdvisorApi.md#setserverlessautoindexing) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing | Set Serverless Auto-Indexing Status |
*PrivateEndpointServicesAPI* | [CreatePrivateEndpoint](./docs/PrivateEndpointServicesApi.md#createprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint | Create One Private Endpoint for One Provider |
*PrivateEndpointServicesAPI* | [CreatePrivateEndpointService](./docs/PrivateEndpointServicesApi.md#createprivateendpointservice) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService | Create One Private Endpoint Service for One Provider |
*PrivateEndpointServicesAPI* | [DeletePrivateEndpoint](./docs/PrivateEndpointServicesApi.md#deleteprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Remove One Private Endpoint for One Provider |
*PrivateEndpointServicesAPI* | [DeletePrivateEndpointService](./docs/PrivateEndpointServicesApi.md#deleteprivateendpointservice) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Remove One Private Endpoint Service for One Provider |
*PrivateEndpointServicesAPI* | [GetPrivateEndpoint](./docs/PrivateEndpointServicesApi.md#getprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Return One Private Endpoint for One Provider |
*PrivateEndpointServicesAPI* | [GetPrivateEndpointService](./docs/PrivateEndpointServicesApi.md#getprivateendpointservice) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Return One Private Endpoint Service for One Provider |
*PrivateEndpointServicesAPI* | [GetRegionalEndpointMode](./docs/PrivateEndpointServicesApi.md#getregionalendpointmode) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Return Regionalized Private Endpoint Status |
*PrivateEndpointServicesAPI* | [ListPrivateEndpointService](./docs/PrivateEndpointServicesApi.md#listprivateendpointservice) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService | Return All Private Endpoint Services for One Provider |
*PrivateEndpointServicesAPI* | [ToggleRegionalEndpointMode](./docs/PrivateEndpointServicesApi.md#toggleregionalendpointmode) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Toggle Regionalized Private Endpoint Status |
*PrivateEndpointServicesAPI* | [UpdatePrivateEndpointService](./docs/PrivateEndpointServicesApi.md#updateprivateendpointservice) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService/{endpointServiceId} | Update One Private Endpoint Service for One Provider |
*ProgrammaticAPIKeysAPI* | [AddGroupApiKey](./docs/ProgrammaticAPIKeysApi.md#addgroupapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Assign One Organization API Key to One Project |
*ProgrammaticAPIKeysAPI* | [CreateGroupApiKey](./docs/ProgrammaticAPIKeysApi.md#creategroupapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys | Create and Assign One Organization API Key to One Project |
*ProgrammaticAPIKeysAPI* | [CreateOrgAccessEntry](./docs/ProgrammaticAPIKeysApi.md#createorgaccessentry) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Create One Access List Entry for One Organization API Key |
*ProgrammaticAPIKeysAPI* | [CreateOrgApiKey](./docs/ProgrammaticAPIKeysApi.md#createorgapikey) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys | Create One Organization API Key |
*ProgrammaticAPIKeysAPI* | [DeleteAccessEntry](./docs/ProgrammaticAPIKeysApi.md#deleteaccessentry) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Remove One Access List Entry for One Organization API Key |
*ProgrammaticAPIKeysAPI* | [DeleteOrgApiKey](./docs/ProgrammaticAPIKeysApi.md#deleteorgapikey) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Remove One Organization API Key |
*ProgrammaticAPIKeysAPI* | [GetOrgAccessEntry](./docs/ProgrammaticAPIKeysApi.md#getorgaccessentry) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Return One Access List Entry for One Organization API Key |
*ProgrammaticAPIKeysAPI* | [GetOrgApiKey](./docs/ProgrammaticAPIKeysApi.md#getorgapikey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Return One Organization API Key |
*ProgrammaticAPIKeysAPI* | [ListGroupApiKeys](./docs/ProgrammaticAPIKeysApi.md#listgroupapikeys) | **Get** /api/atlas/v2/groups/{groupId}/apiKeys | Return All Organization API Keys Assigned to One Project |
*ProgrammaticAPIKeysAPI* | [ListOrgAccessEntries](./docs/ProgrammaticAPIKeysApi.md#listorgaccessentries) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Return All Access List Entries for One Organization API Key |
*ProgrammaticAPIKeysAPI* | [ListOrgApiKeys](./docs/ProgrammaticAPIKeysApi.md#listorgapikeys) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys | Return All Organization API Keys |
*ProgrammaticAPIKeysAPI* | [RemoveGroupApiKey](./docs/ProgrammaticAPIKeysApi.md#removegroupapikey) | **Delete** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Unassign One Organization API Key from One Project |
*ProgrammaticAPIKeysAPI* | [UpdateApiKeyRoles](./docs/ProgrammaticAPIKeysApi.md#updateapikeyroles) | **Patch** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Update Organization API Key Roles for One Project |
*ProgrammaticAPIKeysAPI* | [UpdateOrgApiKey](./docs/ProgrammaticAPIKeysApi.md#updateorgapikey) | **Patch** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Update One Organization API Key |
*ProjectIPAccessListAPI* | [CreateAccessListEntry](./docs/ProjectIPAccessListApi.md#createaccesslistentry) | **Post** /api/atlas/v2/groups/{groupId}/accessList | Add Entries to Project IP Access List |
*ProjectIPAccessListAPI* | [DeleteAccessListEntry](./docs/ProjectIPAccessListApi.md#deleteaccesslistentry) | **Delete** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Remove One Entry from One Project IP Access List |
*ProjectIPAccessListAPI* | [GetAccessListEntry](./docs/ProjectIPAccessListApi.md#getaccesslistentry) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Return One Project IP Access List Entry |
*ProjectIPAccessListAPI* | [GetAccessListStatus](./docs/ProjectIPAccessListApi.md#getaccessliststatus) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue}/status | Return Status of One Project IP Access List Entry |
*ProjectIPAccessListAPI* | [ListAccessListEntries](./docs/ProjectIPAccessListApi.md#listaccesslistentries) | **Get** /api/atlas/v2/groups/{groupId}/accessList | Return All Project IP Access List Entries |
*ProjectsAPI* | [AddGroupUser](./docs/ProjectsApi.md#addgroupuser) | **Post** /api/atlas/v2/groups/{groupId}/access | Add One MongoDB Cloud User to One Project |
*ProjectsAPI* | [CreateGroup](./docs/ProjectsApi.md#creategroup) | **Post** /api/atlas/v2/groups | Create One Project |
*ProjectsAPI* | [CreateGroupInvite](./docs/ProjectsApi.md#creategroupinvite) | **Post** /api/atlas/v2/groups/{groupId}/invites | Create Invitation for One MongoDB Cloud User in One Project |
*ProjectsAPI* | [DeleteGroup](./docs/ProjectsApi.md#deletegroup) | **Delete** /api/atlas/v2/groups/{groupId} | Remove One Project |
*ProjectsAPI* | [DeleteGroupInvite](./docs/ProjectsApi.md#deletegroupinvite) | **Delete** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Remove One Invitation from One Project |
*ProjectsAPI* | [DeleteGroupLimit](./docs/ProjectsApi.md#deletegrouplimit) | **Delete** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Remove One Project Limit |
*ProjectsAPI* | [GetGroup](./docs/ProjectsApi.md#getgroup) | **Get** /api/atlas/v2/groups/{groupId} | Return One Project |
*ProjectsAPI* | [GetGroupByName](./docs/ProjectsApi.md#getgroupbyname) | **Get** /api/atlas/v2/groups/byName/{groupName} | Return One Project by Name |
*ProjectsAPI* | [GetGroupInvite](./docs/ProjectsApi.md#getgroupinvite) | **Get** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Return One Invitation in One Project by Invitation ID |
*ProjectsAPI* | [GetGroupIpAddresses](./docs/ProjectsApi.md#getgroupipaddresses) | **Get** /api/atlas/v2/groups/{groupId}/ipAddresses | Return All IP Addresses for One Project |
*ProjectsAPI* | [GetGroupLimit](./docs/ProjectsApi.md#getgrouplimit) | **Get** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Return One Limit for One Project |
*ProjectsAPI* | [GetGroupSettings](./docs/ProjectsApi.md#getgroupsettings) | **Get** /api/atlas/v2/groups/{groupId}/settings | Return Project Settings |
*ProjectsAPI* | [GetMongoDbVersions](./docs/ProjectsApi.md#getmongodbversions) | **Get** /api/atlas/v2/groups/{groupId}/mongoDBVersions | Return All Available MongoDB LTS Versions for Clusters in One Project |
*ProjectsAPI* | [ListGroupInvites](./docs/ProjectsApi.md#listgroupinvites) | **Get** /api/atlas/v2/groups/{groupId}/invites | Return All Invitations in One Project |
*ProjectsAPI* | [ListGroupLimits](./docs/ProjectsApi.md#listgrouplimits) | **Get** /api/atlas/v2/groups/{groupId}/limits | Return All Limits for One Project |
*ProjectsAPI* | [ListGroups](./docs/ProjectsApi.md#listgroups) | **Get** /api/atlas/v2/groups | Return All Projects |
*ProjectsAPI* | [MigrateGroup](./docs/ProjectsApi.md#migrategroup) | **Post** /api/atlas/v2/groups/{groupId}:migrate | Migrate One Project to Another Organization |
*ProjectsAPI* | [SetGroupLimit](./docs/ProjectsApi.md#setgrouplimit) | **Patch** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Set One Project Limit |
*ProjectsAPI* | [UpdateGroup](./docs/ProjectsApi.md#updategroup) | **Patch** /api/atlas/v2/groups/{groupId} | Update One Project |
*ProjectsAPI* | [UpdateGroupInvites](./docs/ProjectsApi.md#updategroupinvites) | **Patch** /api/atlas/v2/groups/{groupId}/invites | Update One Invitation in One Project |
*ProjectsAPI* | [UpdateGroupSettings](./docs/ProjectsApi.md#updategroupsettings) | **Patch** /api/atlas/v2/groups/{groupId}/settings | Update Project Settings |
*ProjectsAPI* | [UpdateGroupUserRoles](./docs/ProjectsApi.md#updategroupuserroles) | **Put** /api/atlas/v2/groups/{groupId}/users/{userId}/roles | Update Project Roles for One MongoDB Cloud User |
*ProjectsAPI* | [UpdateInviteById](./docs/ProjectsApi.md#updateinvitebyid) | **Patch** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Update One Invitation in One Project by Invitation ID |
*PushBasedLogExportAPI* | [CreateGroupLogIntegration](./docs/PushBasedLogExportApi.md#creategrouplogintegration) | **Post** /api/atlas/v2/groups/{groupId}/logIntegrations | Create One Log Integration |
*PushBasedLogExportAPI* | [CreateLogExport](./docs/PushBasedLogExportApi.md#createlogexport) | **Post** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Create One Push-Based Log Export Configuration in One Project |
*PushBasedLogExportAPI* | [DeleteGroupLogIntegration](./docs/PushBasedLogExportApi.md#deletegrouplogintegration) | **Delete** /api/atlas/v2/groups/{groupId}/logIntegrations/{id} | Remove One Log Integration |
*PushBasedLogExportAPI* | [DeleteLogExport](./docs/PushBasedLogExportApi.md#deletelogexport) | **Delete** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Disable Push-Based Log Export for One Project |
*PushBasedLogExportAPI* | [GetGroupLogIntegration](./docs/PushBasedLogExportApi.md#getgrouplogintegration) | **Get** /api/atlas/v2/groups/{groupId}/logIntegrations/{id} | Return One Log Integration |
*PushBasedLogExportAPI* | [GetLogExport](./docs/PushBasedLogExportApi.md#getlogexport) | **Get** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Return One Push-Based Log Export Configuration in One Project |
*PushBasedLogExportAPI* | [ListGroupLogIntegrations](./docs/PushBasedLogExportApi.md#listgrouplogintegrations) | **Get** /api/atlas/v2/groups/{groupId}/logIntegrations | Return All Active Log Integrations |
*PushBasedLogExportAPI* | [UpdateGroupLogIntegration](./docs/PushBasedLogExportApi.md#updategrouplogintegration) | **Put** /api/atlas/v2/groups/{groupId}/logIntegrations/{id} | Update One Log Integration |
*PushBasedLogExportAPI* | [UpdateLogExport](./docs/PushBasedLogExportApi.md#updatelogexport) | **Patch** /api/atlas/v2/groups/{groupId}/pushBasedLogExport | Update One Push-Based Log Export Configuration in One Project |
*QueryShapeInsightsAPI* | [GetClusterQueryShape](./docs/QueryShapeInsightsApi.md#getclusterqueryshape) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapes/{queryShapeHash} | Return One Query Shape |
*QueryShapeInsightsAPI* | [GetQueryShapeDetails](./docs/QueryShapeInsightsApi.md#getqueryshapedetails) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapeInsights/{queryShapeHash}/details | Return Query Shape Details |
*QueryShapeInsightsAPI* | [ListClusterQueryShapes](./docs/QueryShapeInsightsApi.md#listclusterqueryshapes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapes | Return All Query Shapes |
*QueryShapeInsightsAPI* | [ListQueryShapeSummaries](./docs/QueryShapeInsightsApi.md#listqueryshapesummaries) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapeInsights/summaries | Return Query Statistic Summaries |
*QueryShapeInsightsAPI* | [UpdateClusterQueryShape](./docs/QueryShapeInsightsApi.md#updateclusterqueryshape) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapes/{queryShapeHash} | Update Query Shape Rejection Status |
*RateLimitingAPI* | [GetRateLimit](./docs/RateLimitingApi.md#getratelimit) | **Get** /api/atlas/v2/rateLimits/{endpointSetId} | Return One Rate Limit |
*RateLimitingAPI* | [ListRateLimits](./docs/RateLimitingApi.md#listratelimits) | **Get** /api/atlas/v2/rateLimits | Return All Rate Limits |
*ResourcePoliciesAPI* | [CreateOrgResourcePolicy](./docs/ResourcePoliciesApi.md#createorgresourcepolicy) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Create One Atlas Resource Policy |
*ResourcePoliciesAPI* | [DeleteOrgResourcePolicy](./docs/ResourcePoliciesApi.md#deleteorgresourcepolicy) | **Delete** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Delete One Atlas Resource Policy |
*ResourcePoliciesAPI* | [GetNonCompliantResources](./docs/ResourcePoliciesApi.md#getnoncompliantresources) | **Get** /api/atlas/v2/orgs/{orgId}/nonCompliantResources | Return All Non-Compliant Resources |
*ResourcePoliciesAPI* | [GetOrgResourcePolicy](./docs/ResourcePoliciesApi.md#getorgresourcepolicy) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Return One Atlas Resource Policy |
*ResourcePoliciesAPI* | [ListOrgResourcePolicies](./docs/ResourcePoliciesApi.md#listorgresourcepolicies) | **Get** /api/atlas/v2/orgs/{orgId}/resourcePolicies | Return All Atlas Resource Policies |
*ResourcePoliciesAPI* | [UpdateOrgResourcePolicy](./docs/ResourcePoliciesApi.md#updateorgresourcepolicy) | **Patch** /api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId} | Update One Atlas Resource Policy |
*ResourcePoliciesAPI* | [ValidateResourcePolicies](./docs/ResourcePoliciesApi.md#validateresourcepolicies) | **Post** /api/atlas/v2/orgs/{orgId}/resourcePolicies:validate | Validate One Atlas Resource Policy |
*RollingIndexAPI* | [CreateRollingIndex](./docs/RollingIndexApi.md#createrollingindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/index | Create One Rolling Index |
*RootAPI* | [GetSystemStatus](./docs/RootApi.md#getsystemstatus) | **Get** /api/atlas/v2 | Return the Status of This MongoDB Application |
*RootAPI* | [ListControlPlaneAddresses](./docs/RootApi.md#listcontrolplaneaddresses) | **Get** /api/atlas/v2/unauth/controlPlaneIPAddresses | Return All Control Plane IP Addresses |
*ServerlessInstancesAPI* | [GetServerlessInstance](./docs/ServerlessInstancesApi.md#getserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{name} | Return One Serverless Instance from One Project |
*ServerlessInstancesAPI* | [ListServerlessInstances](./docs/ServerlessInstancesApi.md#listserverlessinstances) | **Get** /api/atlas/v2/groups/{groupId}/serverless | Return All Serverless Instances in One Project |
*ServerlessPrivateEndpointsAPI* | [CreateServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#createserverlessprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Create One Private Endpoint for One Serverless Instance |
*ServerlessPrivateEndpointsAPI* | [DeleteServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#deleteserverlessprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Remove One Private Endpoint for One Serverless Instance |
*ServerlessPrivateEndpointsAPI* | [GetServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#getserverlessprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Return One Private Endpoint for One Serverless Instance |
*ServerlessPrivateEndpointsAPI* | [ListServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#listserverlessprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Return All Private Endpoints for One Serverless Instance |
*ServerlessPrivateEndpointsAPI* | [UpdateServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#updateserverlessprivateendpoint) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Update One Private Endpoint for One Serverless Instance |
*ServiceAccountsAPI* | [CreateAccessList](./docs/ServiceAccountsApi.md#createaccesslist) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Project Service Account |
*ServiceAccountsAPI* | [CreateGroupSecret](./docs/ServiceAccountsApi.md#creategroupsecret) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets | Create One Project Service Account Secret |
*ServiceAccountsAPI* | [CreateGroupServiceAccount](./docs/ServiceAccountsApi.md#creategroupserviceaccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts | Create One Project Service Account |
*ServiceAccountsAPI* | [CreateOrgAccessList](./docs/ServiceAccountsApi.md#createorgaccesslist) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Add Access List Entries for One Organization Service Account |
*ServiceAccountsAPI* | [CreateOrgSecret](./docs/ServiceAccountsApi.md#createorgsecret) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets | Create One Organization Service Account Secret |
*ServiceAccountsAPI* | [CreateOrgServiceAccount](./docs/ServiceAccountsApi.md#createorgserviceaccount) | **Post** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Create One Organization Service Account |
*ServiceAccountsAPI* | [DeleteGroupAccessEntry](./docs/ServiceAccountsApi.md#deletegroupaccessentry) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Project Service Account |
*ServiceAccountsAPI* | [DeleteGroupSecret](./docs/ServiceAccountsApi.md#deletegroupsecret) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Project Service Account Secret |
*ServiceAccountsAPI* | [DeleteGroupServiceAccount](./docs/ServiceAccountsApi.md#deletegroupserviceaccount) | **Delete** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Remove One Project Service Account |
*ServiceAccountsAPI* | [DeleteOrgAccessEntry](./docs/ServiceAccountsApi.md#deleteorgaccessentry) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList/{ipAddress} | Remove One Access List Entry from One Organization Service Account |
*ServiceAccountsAPI* | [DeleteOrgSecret](./docs/ServiceAccountsApi.md#deleteorgsecret) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets/{secretId} | Delete One Organization Service Account Secret |
*ServiceAccountsAPI* | [DeleteOrgServiceAccount](./docs/ServiceAccountsApi.md#deleteorgserviceaccount) | **Delete** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Delete One Organization Service Account |
*ServiceAccountsAPI* | [GetGroupServiceAccount](./docs/ServiceAccountsApi.md#getgroupserviceaccount) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Return One Project Service Account |
*ServiceAccountsAPI* | [GetOrgServiceAccount](./docs/ServiceAccountsApi.md#getorgserviceaccount) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Return One Organization Service Account |
*ServiceAccountsAPI* | [GetServiceAccountGroups](./docs/ServiceAccountsApi.md#getserviceaccountgroups) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/groups | Return All Service Account Project Assignments |
*ServiceAccountsAPI* | [InviteGroupServiceAccount](./docs/ServiceAccountsApi.md#invitegroupserviceaccount) | **Post** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}:invite | Assign One Service Account to One Project |
*ServiceAccountsAPI* | [ListAccessList](./docs/ServiceAccountsApi.md#listaccesslist) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Project Service Account |
*ServiceAccountsAPI* | [ListGroupServiceAccounts](./docs/ServiceAccountsApi.md#listgroupserviceaccounts) | **Get** /api/atlas/v2/groups/{groupId}/serviceAccounts | Return All Project Service Accounts |
*ServiceAccountsAPI* | [ListOrgAccessList](./docs/ServiceAccountsApi.md#listorgaccesslist) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList | Return All Access List Entries for One Organization Service Account |
*ServiceAccountsAPI* | [ListOrgServiceAccounts](./docs/ServiceAccountsApi.md#listorgserviceaccounts) | **Get** /api/atlas/v2/orgs/{orgId}/serviceAccounts | Return All Organization Service Accounts |
*ServiceAccountsAPI* | [UpdateGroupServiceAccount](./docs/ServiceAccountsApi.md#updategroupserviceaccount) | **Patch** /api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId} | Update One Project Service Account |
*ServiceAccountsAPI* | [UpdateOrgServiceAccount](./docs/ServiceAccountsApi.md#updateorgserviceaccount) | **Patch** /api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId} | Update One Organization Service Account |
*StreamsAPI* | [AcceptVpcPeeringConnection](./docs/StreamsApi.md#acceptvpcpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:accept | Accept One Incoming VPC Peering Connection |
*StreamsAPI* | [CreateFailoverConnection](./docs/StreamsApi.md#createfailoverconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections | Create One Failover Stream Connection |
*StreamsAPI* | [CreatePrivateLinkConnection](./docs/StreamsApi.md#createprivatelinkconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Create One Private Link Connection |
*StreamsAPI* | [CreateStreamConnection](./docs/StreamsApi.md#createstreamconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Create One Stream Connection |
*StreamsAPI* | [CreateStreamProcessor](./docs/StreamsApi.md#createstreamprocessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor | Create One Stream Processor |
*StreamsAPI* | [CreateStreamWorkspace](./docs/StreamsApi.md#createstreamworkspace) | **Post** /api/atlas/v2/groups/{groupId}/streams | Create One Stream Workspace |
*StreamsAPI* | [DeletePrivateLinkConnection](./docs/StreamsApi.md#deleteprivatelinkconnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Delete One Private Link Connection |
*StreamsAPI* | [DeleteStreamConnection](./docs/StreamsApi.md#deletestreamconnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Delete One Stream Connection |
*StreamsAPI* | [DeleteStreamFailoverConnection](./docs/StreamsApi.md#deletestreamfailoverconnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections/{failoverConnectionId} | Delete One Stream Failover Connection |
*StreamsAPI* | [DeleteStreamProcessor](./docs/StreamsApi.md#deletestreamprocessor) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Delete One Stream Processor |
*StreamsAPI* | [DeleteStreamWorkspace](./docs/StreamsApi.md#deletestreamworkspace) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Delete One Stream Workspace |
*StreamsAPI* | [DeleteVpcPeeringConnection](./docs/StreamsApi.md#deletevpcpeeringconnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id} | Delete One VPC Peering Connection |
*StreamsAPI* | [DownloadAuditLogs](./docs/StreamsApi.md#downloadauditlogs) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/auditLogs | Download Audit Logs for One Atlas Stream Processing Workspace |
*StreamsAPI* | [DownloadOperationalLogs](./docs/StreamsApi.md#downloadoperationallogs) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}:downloadOperationalLogs | Download Operational Logs for One Atlas Stream Processing Workspace |
*StreamsAPI* | [GetAccountDetails](./docs/StreamsApi.md#getaccountdetails) | **Get** /api/atlas/v2/groups/{groupId}/streams/accountDetails | Return Account ID and VPC ID for One Project and Region |
*StreamsAPI* | [GetPrivateLinkConnection](./docs/StreamsApi.md#getprivatelinkconnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Return One Private Link Connection |
*StreamsAPI* | [GetStreamConnection](./docs/StreamsApi.md#getstreamconnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Return One Stream Connection |
*StreamsAPI* | [GetStreamFailoverConnection](./docs/StreamsApi.md#getstreamfailoverconnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections/{failoverConnectionId} | Return One Stream Failover Connection |
*StreamsAPI* | [GetStreamProcessor](./docs/StreamsApi.md#getstreamprocessor) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Return One Stream Processor |
*StreamsAPI* | [GetStreamProcessors](./docs/StreamsApi.md#getstreamprocessors) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processors | Return All Stream Processors in One Stream Workspace |
*StreamsAPI* | [GetStreamWorkspace](./docs/StreamsApi.md#getstreamworkspace) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Return One Stream Workspace |
*StreamsAPI* | [ListActivePeeringConnections](./docs/StreamsApi.md#listactivepeeringconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/activeVpcPeeringConnections | Return All Active Incoming VPC Peering Connections |
*StreamsAPI* | [ListFailoverConnections](./docs/StreamsApi.md#listfailoverconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections | Return All Stream Failover Connections |
*StreamsAPI* | [ListPrivateLinkConnections](./docs/StreamsApi.md#listprivatelinkconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Return All Private Link Connections |
*StreamsAPI* | [ListStreamConnections](./docs/StreamsApi.md#liststreamconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Return All Connections of the Stream Workspaces |
*StreamsAPI* | [ListStreamWorkspaces](./docs/StreamsApi.md#liststreamworkspaces) | **Get** /api/atlas/v2/groups/{groupId}/streams | Return All Stream Workspaces in One Project |
*StreamsAPI* | [ListVpcPeeringConnections](./docs/StreamsApi.md#listvpcpeeringconnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections | Return All VPC Peering Connections |
*StreamsAPI* | [RejectVpcPeeringConnection](./docs/StreamsApi.md#rejectvpcpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:reject | Reject One Incoming VPC Peering Connection |
*StreamsAPI* | [StartStreamProcessor](./docs/StreamsApi.md#startstreamprocessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:start | Start One Stream Processor |
*StreamsAPI* | [StartStreamProcessorWith](./docs/StreamsApi.md#startstreamprocessorwith) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:startWith | Start One Stream Processor With Options |
*StreamsAPI* | [StopStreamProcessor](./docs/StreamsApi.md#stopstreamprocessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:stop | Stop One Stream Processor |
*StreamsAPI* | [UpdateStreamConnection](./docs/StreamsApi.md#updatestreamconnection) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Update One Stream Connection |
*StreamsAPI* | [UpdateStreamFailoverConnection](./docs/StreamsApi.md#updatestreamfailoverconnection) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections/{failoverConnectionId} | Update One Stream Failover Connection |
*StreamsAPI* | [UpdateStreamProcessor](./docs/StreamsApi.md#updatestreamprocessor) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Update One Stream Processor |
*StreamsAPI* | [UpdateStreamWorkspace](./docs/StreamsApi.md#updatestreamworkspace) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Update One Stream Workspace |
*StreamsAPI* | [WithStreamSampleConnections](./docs/StreamsApi.md#withstreamsampleconnections) | **Post** /api/atlas/v2/groups/{groupId}/streams:withSampleConnections | Create One Stream Workspace with Sample Connections |
*TeamsAPI* | [AddGroupTeams](./docs/TeamsApi.md#addgroupteams) | **Post** /api/atlas/v2/groups/{groupId}/teams | Add Multiple Teams to One Project |
*TeamsAPI* | [AddTeamUsers](./docs/TeamsApi.md#addteamusers) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Assign MongoDB Cloud Users in One Organization to One Team |
*TeamsAPI* | [CreateOrgTeam](./docs/TeamsApi.md#createorgteam) | **Post** /api/atlas/v2/orgs/{orgId}/teams | Create One Team in One Organization |
*TeamsAPI* | [DeleteOrgTeam](./docs/TeamsApi.md#deleteorgteam) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Remove One Team from One Organization |
*TeamsAPI* | [GetGroupTeam](./docs/TeamsApi.md#getgroupteam) | **Get** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Return One Team in One Project |
*TeamsAPI* | [GetOrgTeam](./docs/TeamsApi.md#getorgteam) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Return One Team by ID |
*TeamsAPI* | [GetTeamByName](./docs/TeamsApi.md#getteambyname) | **Get** /api/atlas/v2/orgs/{orgId}/teams/byName/{teamName} | Return One Team by Name |
*TeamsAPI* | [ListGroupTeams](./docs/TeamsApi.md#listgroupteams) | **Get** /api/atlas/v2/groups/{groupId}/teams | Return All Teams in One Project |
*TeamsAPI* | [ListOrgTeams](./docs/TeamsApi.md#listorgteams) | **Get** /api/atlas/v2/orgs/{orgId}/teams | Return All Teams in One Organization |
*TeamsAPI* | [RemoveGroupTeam](./docs/TeamsApi.md#removegroupteam) | **Delete** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Remove One Team from One Project |
*TeamsAPI* | [RemoveUserFromTeam](./docs/TeamsApi.md#removeuserfromteam) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users/{userId} | Remove One MongoDB Cloud User from One Team |
*TeamsAPI* | [RenameOrgTeam](./docs/TeamsApi.md#renameorgteam) | **Patch** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Rename One Team |
*TeamsAPI* | [UpdateGroupTeam](./docs/TeamsApi.md#updategroupteam) | **Patch** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Update Team Roles in One Project |
*ThirdPartyIntegrationsAPI* | [CreateGroupIntegration](./docs/ThirdPartyIntegrationsApi.md#creategroupintegration) | **Post** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Create One Third-Party Service Integration |
*ThirdPartyIntegrationsAPI* | [DeleteGroupIntegration](./docs/ThirdPartyIntegrationsApi.md#deletegroupintegration) | **Delete** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Remove One Third-Party Service Integration |
*ThirdPartyIntegrationsAPI* | [GetGroupIntegration](./docs/ThirdPartyIntegrationsApi.md#getgroupintegration) | **Get** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Return One Third-Party Service Integration |
*ThirdPartyIntegrationsAPI* | [ListGroupIntegrations](./docs/ThirdPartyIntegrationsApi.md#listgroupintegrations) | **Get** /api/atlas/v2/groups/{groupId}/integrations | Return All Active Third-Party Service Integrations |
*ThirdPartyIntegrationsAPI* | [UpdateGroupIntegration](./docs/ThirdPartyIntegrationsApi.md#updategroupintegration) | **Put** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Update One Third-Party Service Integration |
*X509AuthenticationAPI* | [CreateDatabaseUserCert](./docs/X509AuthenticationApi.md#createdatabaseusercert) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Create One X.509 Certificate for One Database User |
*X509AuthenticationAPI* | [DisableSecurityCustomerX509](./docs/X509AuthenticationApi.md#disablesecuritycustomerx509) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/customerX509 | Disable Customer-Managed X.509 |
*X509AuthenticationAPI* | [ListDatabaseUserCerts](./docs/X509AuthenticationApi.md#listdatabaseusercerts) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Return All X.509 Certificates Assigned to One Database User |


## Documentation For Models

 - [AWSCustomDNSEnabled](./docs/AWSCustomDNSEnabled.md)
 - [AWSKMSConfiguration](./docs/AWSKMSConfiguration.md)
 - [AccessListItem](./docs/AccessListItem.md)
 - [AccountDetails](./docs/AccountDetails.md)
 - [AcknowledgeAlert](./docs/AcknowledgeAlert.md)
 - [ActivityFeedLinkResponse](./docs/ActivityFeedLinkResponse.md)
 - [AddOrRemoveGroupRole](./docs/AddOrRemoveGroupRole.md)
 - [AddOrRemoveOrgRole](./docs/AddOrRemoveOrgRole.md)
 - [AddOrRemoveUserFromTeam](./docs/AddOrRemoveUserFromTeam.md)
 - [AddUserToTeam](./docs/AddUserToTeam.md)
 - [AdditionalData](./docs/AdditionalData.md)
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
 - [ApiAtlasModifyEndpointServiceRequest](./docs/ApiAtlasModifyEndpointServiceRequest.md)
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
 - [ApiSearchDeploymentEffectiveSpec](./docs/ApiSearchDeploymentEffectiveSpec.md)
 - [ApiSearchDeploymentRequest](./docs/ApiSearchDeploymentRequest.md)
 - [ApiSearchDeploymentRequestSpec](./docs/ApiSearchDeploymentRequestSpec.md)
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
 - [AzureConnection](./docs/AzureConnection.md)
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
 - [CloudProviderAccessAzureServicePrincipal](./docs/CloudProviderAccessAzureServicePrincipal.md)
 - [CloudProviderAccessFeatureUsage](./docs/CloudProviderAccessFeatureUsage.md)
 - [CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId](./docs/CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId.md)
 - [CloudProviderAccessGCPServiceAccount](./docs/CloudProviderAccessGCPServiceAccount.md)
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
 - [ClusterConfigurationValidation](./docs/ClusterConfigurationValidation.md)
 - [ClusterConfigurationValidationError](./docs/ClusterConfigurationValidationError.md)
 - [ClusterConfigurationValidationResult](./docs/ClusterConfigurationValidationResult.md)
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
 - [CollStatsLatencyNamespaceMetric](./docs/CollStatsLatencyNamespaceMetric.md)
 - [CollStatsLatencyNamespaceMetrics](./docs/CollStatsLatencyNamespaceMetrics.md)
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
 - [CustomSessionTimeouts](./docs/CustomSessionTimeouts.md)
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
 - [DiskBackupCopyPolicyItem](./docs/DiskBackupCopyPolicyItem.md)
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
 - [DiskBackupSnapshotAWSExportBucketResponse](./docs/DiskBackupSnapshotAWSExportBucketResponse.md)
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
 - [Header](./docs/Header.md)
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
 - [LogIntegrationRequest](./docs/LogIntegrationRequest.md)
 - [LogIntegrationResponse](./docs/LogIntegrationResponse.md)
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
 - [ObjectStoragePrivateEndpointRequest](./docs/ObjectStoragePrivateEndpointRequest.md)
 - [ObjectStoragePrivateEndpointResponse](./docs/ObjectStoragePrivateEndpointResponse.md)
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
 - [PaginatedApiAtlasObjectStoragePrivateEndpointResponse](./docs/PaginatedApiAtlasObjectStoragePrivateEndpointResponse.md)
 - [PaginatedApiAtlasProviderRegions](./docs/PaginatedApiAtlasProviderRegions.md)
 - [PaginatedApiAtlasServerlessBackupRestoreJob](./docs/PaginatedApiAtlasServerlessBackupRestoreJob.md)
 - [PaginatedApiAtlasServerlessBackupSnapshot](./docs/PaginatedApiAtlasServerlessBackupSnapshot.md)
 - [PaginatedApiInvoice](./docs/PaginatedApiInvoice.md)
 - [PaginatedApiInvoiceMetadata](./docs/PaginatedApiInvoiceMetadata.md)
 - [PaginatedApiSKU](./docs/PaginatedApiSKU.md)
 - [PaginatedApiStreamsConnection](./docs/PaginatedApiStreamsConnection.md)
 - [PaginatedApiStreamsFailoverConnection](./docs/PaginatedApiStreamsFailoverConnection.md)
 - [PaginatedApiStreamsPrivateLink](./docs/PaginatedApiStreamsPrivateLink.md)
 - [PaginatedApiStreamsStreamProcessorWithStats](./docs/PaginatedApiStreamsStreamProcessorWithStats.md)
 - [PaginatedApiStreamsTenant](./docs/PaginatedApiStreamsTenant.md)
 - [PaginatedApiStreamsVPCPeeringConnection](./docs/PaginatedApiStreamsVPCPeeringConnection.md)
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
 - [PaginatedLogIntegrationResponse](./docs/PaginatedLogIntegrationResponse.md)
 - [PaginatedNetworkAccess](./docs/PaginatedNetworkAccess.md)
 - [PaginatedOnlineArchive](./docs/PaginatedOnlineArchive.md)
 - [PaginatedOrgGroup](./docs/PaginatedOrgGroup.md)
 - [PaginatedOrgServiceAccounts](./docs/PaginatedOrgServiceAccounts.md)
 - [PaginatedOrgUser](./docs/PaginatedOrgUser.md)
 - [PaginatedOrganization](./docs/PaginatedOrganization.md)
 - [PaginatedPipelineRun](./docs/PaginatedPipelineRun.md)
 - [PaginatedPrivateNetworkEndpointIdEntry](./docs/PaginatedPrivateNetworkEndpointIdEntry.md)
 - [PaginatedPublicApiUsageDetailsLineItem](./docs/PaginatedPublicApiUsageDetailsLineItem.md)
 - [PaginatedQueryShapes](./docs/PaginatedQueryShapes.md)
 - [PaginatedRateLimitEndpointSets](./docs/PaginatedRateLimitEndpointSets.md)
 - [PaginatedRestoreJob](./docs/PaginatedRestoreJob.md)
 - [PaginatedRoleMapping](./docs/PaginatedRoleMapping.md)
 - [PaginatedServerlessInstanceDescription](./docs/PaginatedServerlessInstanceDescription.md)
 - [PaginatedServiceAccountGroup](./docs/PaginatedServiceAccountGroup.md)
 - [PaginatedServiceAccountIPAccessEntry](./docs/PaginatedServiceAccountIPAccessEntry.md)
 - [PaginatedSnapshot](./docs/PaginatedSnapshot.md)
 - [PaginatedTeam](./docs/PaginatedTeam.md)
 - [PaginatedTeamRole](./docs/PaginatedTeamRole.md)
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
 - [QueryShapeResponse](./docs/QueryShapeResponse.md)
 - [QueryShapeSeenMetadata](./docs/QueryShapeSeenMetadata.md)
 - [QueryShapeUpdateRequest](./docs/QueryShapeUpdateRequest.md)
 - [QueryStatsDetailsResponse](./docs/QueryStatsDetailsResponse.md)
 - [QueryStatsSummary](./docs/QueryStatsSummary.md)
 - [QueryStatsSummaryListResponse](./docs/QueryStatsSummaryListResponse.md)
 - [RateLimitEndpointSetCapacity](./docs/RateLimitEndpointSetCapacity.md)
 - [RateLimitEndpointSetEndpoint](./docs/RateLimitEndpointSetEndpoint.md)
 - [RateLimitEndpointSetRefillDurationSeconds](./docs/RateLimitEndpointSetRefillDurationSeconds.md)
 - [RateLimitEndpointSetRefillRate](./docs/RateLimitEndpointSetRefillRate.md)
 - [RateLimitEndpointSetResponse](./docs/RateLimitEndpointSetResponse.md)
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
 - [SchemaRegistryAuthentication](./docs/SchemaRegistryAuthentication.md)
 - [SearchIndexCreateRequest](./docs/SearchIndexCreateRequest.md)
 - [SearchIndexDefinitionVersion](./docs/SearchIndexDefinitionVersion.md)
 - [SearchIndexResponse](./docs/SearchIndexResponse.md)
 - [SearchIndexUpdateRequest](./docs/SearchIndexUpdateRequest.md)
 - [SearchIndexUpdateRequestDefinition](./docs/SearchIndexUpdateRequestDefinition.md)
 - [SearchMappings](./docs/SearchMappings.md)
 - [SearchSynonymMappingDefinition](./docs/SearchSynonymMappingDefinition.md)
 - [SearchTypeSets](./docs/SearchTypeSets.md)
 - [ServerlessBackupRestoreJob](./docs/ServerlessBackupRestoreJob.md)
 - [ServerlessBackupSnapshot](./docs/ServerlessBackupSnapshot.md)
 - [ServerlessConnectionStringsPrivateEndpointItem](./docs/ServerlessConnectionStringsPrivateEndpointItem.md)
 - [ServerlessConnectionStringsPrivateEndpointList](./docs/ServerlessConnectionStringsPrivateEndpointList.md)
 - [ServerlessInstanceDescription](./docs/ServerlessInstanceDescription.md)
 - [ServerlessInstanceDescriptionConnectionStrings](./docs/ServerlessInstanceDescriptionConnectionStrings.md)
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
 - [SkuResponse](./docs/SkuResponse.md)
 - [Source](./docs/Source.md)
 - [StateReason](./docs/StateReason.md)
 - [StreamConfig](./docs/StreamConfig.md)
 - [StreamProcessorMetricThreshold](./docs/StreamProcessorMetricThreshold.md)
 - [StreamsAWSConnectionConfig](./docs/StreamsAWSConnectionConfig.md)
 - [StreamsConnection](./docs/StreamsConnection.md)
 - [StreamsDLQ](./docs/StreamsDLQ.md)
 - [StreamsDataProcessRegion](./docs/StreamsDataProcessRegion.md)
 - [StreamsGCPConnectionConfig](./docs/StreamsGCPConnectionConfig.md)
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
 - [StreamsProcessorStatus](./docs/StreamsProcessorStatus.md)
 - [StreamsProcessorWithStats](./docs/StreamsProcessorWithStats.md)
 - [StreamsPublicPrivateLinkNetworking](./docs/StreamsPublicPrivateLinkNetworking.md)
 - [StreamsPublicPrivateLinkNetworkingAccess](./docs/StreamsPublicPrivateLinkNetworkingAccess.md)
 - [StreamsSampleConnections](./docs/StreamsSampleConnections.md)
 - [StreamsStartProcessorFailover](./docs/StreamsStartProcessorFailover.md)
 - [StreamsStartStreamProcessorWith](./docs/StreamsStartStreamProcessorWith.md)
 - [StreamsTenant](./docs/StreamsTenant.md)
 - [StreamsTenantUpdateRequest](./docs/StreamsTenantUpdateRequest.md)
 - [SynonymMappingStatusDetail](./docs/SynonymMappingStatusDetail.md)
 - [SynonymSource](./docs/SynonymSource.md)
 - [SystemStatus](./docs/SystemStatus.md)
 - [TargetOrg](./docs/TargetOrg.md)
 - [TargetOrgRequest](./docs/TargetOrgRequest.md)
 - [Team](./docs/Team.md)
 - [TeamResponse](./docs/TeamResponse.md)
 - [TeamRole](./docs/TeamRole.md)
 - [TeamUpdate](./docs/TeamUpdate.md)
 - [ThirdPartyIntegration](./docs/ThirdPartyIntegration.md)
 - [TriggerIngestionPipelineRequest](./docs/TriggerIngestionPipelineRequest.md)
 - [UpdateAtlasOrganizationApiKey](./docs/UpdateAtlasOrganizationApiKey.md)
 - [UpdateAtlasProjectApiKey](./docs/UpdateAtlasProjectApiKey.md)
 - [UpdateCustomDBRole](./docs/UpdateCustomDBRole.md)
 - [UpdateGroupRolesForUser](./docs/UpdateGroupRolesForUser.md)
 - [UpdateOrgRolesForUser](./docs/UpdateOrgRolesForUser.md)
 - [UpdateRequirePrivateNetworkingRequest](./docs/UpdateRequirePrivateNetworkingRequest.md)
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
 - [VPCPeeringConnection](./docs/VPCPeeringConnection.md)
 - [VectorSearchHostStatusDetail](./docs/VectorSearchHostStatusDetail.md)
 - [VectorSearchIndexDefinition](./docs/VectorSearchIndexDefinition.md)
 - [VectorSearchIndexStatusDetail](./docs/VectorSearchIndexStatusDetail.md)
 - [X509Certificate](./docs/X509Certificate.md)
 - [X509CertificateUpdate](./docs/X509CertificateUpdate.md)
 - [ZoneMapping](./docs/ZoneMapping.md)




