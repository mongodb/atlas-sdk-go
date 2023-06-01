# SDK Reference 

## Documentation for SDK Endpoints

All URIs are relative to *https://cloud.mongodb.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AWSClustersDNSApi* | [GetAWSCustomDNS](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSClustersDNSApi.md#getawscustomdns) | **Get** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Return One Custom DNS Configuration for Atlas Clusters on AWS
*AWSClustersDNSApi* | [ToggleAWSCustomDNS](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSClustersDNSApi.md#toggleawscustomdns) | **Patch** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS
*AccessTrackingApi* | [ListAccessLogsByClusterName](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AccessTrackingApi.md#listaccesslogsbyclustername) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/clusters/{clusterName} | Return Database Access History for One Cluster using Its Cluster Name
*AccessTrackingApi* | [ListAccessLogsByHostname](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AccessTrackingApi.md#listaccesslogsbyhostname) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/processes/{hostname} | Return Database Access History for One Cluster using Its Hostname
*AlertConfigurationsApi* | [CreateAlertConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#createalertconfiguration) | **Post** /api/atlas/v2/groups/{groupId}/alertConfigs | Create One Alert Configuration in One Project
*AlertConfigurationsApi* | [DeleteAlertConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#deletealertconfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Remove One Alert Configuration from One Project
*AlertConfigurationsApi* | [GetAlertConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#getalertconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Return One Alert Configuration from One Project
*AlertConfigurationsApi* | [ListAlertConfigurationMatchersFieldNames](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#listalertconfigurationmatchersfieldnames) | **Get** /api/atlas/v2/alertConfigs/matchers/fieldNames | Get All Alert Configuration Matchers Field Names
*AlertConfigurationsApi* | [ListAlertConfigurations](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#listalertconfigurations) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs | Return All Alert Configurations for One Project
*AlertConfigurationsApi* | [ListAlertConfigurationsByAlertId](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#listalertconfigurationsbyalertid) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId}/alertConfigs | Return All Alert Configurations Set for One Alert
*AlertConfigurationsApi* | [ToggleAlertConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#togglealertconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Toggle One State of One Alert Configuration in One Project
*AlertConfigurationsApi* | [UpdateAlertConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigurationsApi.md#updatealertconfiguration) | **Put** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Update One Alert Configuration for One Project
*AlertsApi* | [AcknowledgeAlert](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertsApi.md#acknowledgealert) | **Patch** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Acknowledge One Alert from One Project
*AlertsApi* | [GetAlert](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertsApi.md#getalert) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Return One Alert from One Project
*AlertsApi* | [ListAlerts](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertsApi.md#listalerts) | **Get** /api/atlas/v2/groups/{groupId}/alerts | Return All Alerts from One Project
*AlertsApi* | [ListAlertsByAlertConfigurationId](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertsApi.md#listalertsbyalertconfigurationid) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}/alerts | Return All Open Alerts for Alert Configuration
*AtlasSearchApi* | [CreateAtlasSearchIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AtlasSearchApi.md#createatlassearchindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index
*AtlasSearchApi* | [DeleteAtlasSearchIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AtlasSearchApi.md#deleteatlassearchindex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index
*AtlasSearchApi* | [GetAtlasSearchIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AtlasSearchApi.md#getatlassearchindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index
*AtlasSearchApi* | [ListAtlasSearchIndexes](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AtlasSearchApi.md#listatlassearchindexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
*AtlasSearchApi* | [UpdateAtlasSearchIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AtlasSearchApi.md#updateatlassearchindex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index
*AuditingApi* | [GetAuditingConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AuditingApi.md#getauditingconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/auditLog | Return the Auditing Configuration for One Project
*AuditingApi* | [UpdateAuditingConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AuditingApi.md#updateauditingconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/auditLog | Update Auditing Configuration for One Project
*CloudBackupsApi* | [CancelBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#cancelbackuprestorejob) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job of One Cluster
*CloudBackupsApi* | [CreateBackupExportJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#createbackupexportjob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Cloud Backup Snapshot Export Job
*CloudBackupsApi* | [CreateBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#createbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Cluster
*CloudBackupsApi* | [CreateExportBucket](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#createexportbucket) | **Post** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [CreateServerlessBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#createserverlessbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Serverless Instance
*CloudBackupsApi* | [DeleteAllBackupSchedules](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#deleteallbackupschedules) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules
*CloudBackupsApi* | [DeleteExportBucket](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#deleteexportbucket) | **Delete** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [DeleteReplicaSetBackup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#deletereplicasetbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup
*CloudBackupsApi* | [DeleteShardedClusterBackup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#deleteshardedclusterbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup
*CloudBackupsApi* | [GetBackupExportJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getbackupexportjob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Cloud Backup Snapshot Export Job
*CloudBackupsApi* | [GetBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job of One Cluster
*CloudBackupsApi* | [GetBackupSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getbackupschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule
*CloudBackupsApi* | [GetDataProtectionSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getdataprotectionsettings) | **Get** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Return the Backup Compliance Policy settings
*CloudBackupsApi* | [GetExportBucket](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getexportbucket) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [GetReplicaSetBackup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getreplicasetbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup
*CloudBackupsApi* | [GetServerlessBackup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getserverlessbackup) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance
*CloudBackupsApi* | [GetServerlessBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getserverlessbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance
*CloudBackupsApi* | [GetShardedClusterBackup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#getshardedclusterbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup
*CloudBackupsApi* | [ListBackupExportJobs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#listbackupexportjobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Cloud Backup Snapshot Export Jobs
*CloudBackupsApi* | [ListBackupRestoreJobs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#listbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster
*CloudBackupsApi* | [ListExportBuckets](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#listexportbuckets) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [ListReplicaSetBackups](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#listreplicasetbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups
*CloudBackupsApi* | [ListServerlessBackupRestoreJobs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#listserverlessbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance
*CloudBackupsApi* | [ListServerlessBackups](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#listserverlessbackups) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance
*CloudBackupsApi* | [ListShardedClusterBackups](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#listshardedclusterbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups
*CloudBackupsApi* | [TakeSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#takesnapshot) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot
*CloudBackupsApi* | [UpdateBackupSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#updatebackupschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster
*CloudBackupsApi* | [UpdateDataProtectionSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#updatedataprotectionsettings) | **Put** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Update or enable the Backup Compliance Policy settings
*CloudBackupsApi* | [UpdateSnapshotRetention](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudBackupsApi.md#updatesnapshotretention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Change Expiration Date for One Cloud Backup
*CloudMigrationServiceApi* | [CreateLinkToken](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#createlinktoken) | **Post** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Create One Link-Token
*CloudMigrationServiceApi* | [CreatePushMigration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#createpushmigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations | Migrate One Local Managed Cluster to MongoDB Atlas
*CloudMigrationServiceApi* | [CutoverMigration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#cutovermigration) | **Put** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover | Cut Over the Migrated Cluster
*CloudMigrationServiceApi* | [DeleteLinkToken](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#deletelinktoken) | **Delete** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Remove One Link-Token
*CloudMigrationServiceApi* | [GetPushMigration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#getpushmigration) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId} | Return One Migration Job
*CloudMigrationServiceApi* | [GetValidationStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#getvalidationstatus) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/validate/{validationId} | Return One Migration Validation Job
*CloudMigrationServiceApi* | [ListSourceProjects](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#listsourceprojects) | **Get** /api/atlas/v2/orgs/{orgId}/liveMigrations/availableProjects | Return All Projects Available for Migration
*CloudMigrationServiceApi* | [ValidateMigration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudMigrationServiceApi.md#validatemigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations/validate | Validate One Migration Request
*CloudProviderAccessApi* | [AuthorizeCloudProviderAccessRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessApi.md#authorizecloudprovideraccessrole) | **Patch** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role
*CloudProviderAccessApi* | [CreateCloudProviderAccessRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessApi.md#createcloudprovideraccessrole) | **Post** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role
*CloudProviderAccessApi* | [DeauthorizeCloudProviderAccessRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessApi.md#deauthorizecloudprovideraccessrole) | **Delete** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role
*CloudProviderAccessApi* | [GetCloudProviderAccessRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessApi.md#getcloudprovideraccessrole) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Return specified Cloud Provider Access Role
*CloudProviderAccessApi* | [ListCloudProviderAccessRoles](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessApi.md#listcloudprovideraccessroles) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles
*ClusterOutageSimulationApi* | [EndOutageSimulation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterOutageSimulationApi.md#endoutagesimulation) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | End an Outage Simulation
*ClusterOutageSimulationApi* | [GetOutageSimulation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterOutageSimulationApi.md#getoutagesimulation) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Return One Outage Simulation
*ClusterOutageSimulationApi* | [StartOutageSimulation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterOutageSimulationApi.md#startoutagesimulation) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Start an Outage Simulation
*ClustersApi* | [GetClusterAdvancedConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#getclusteradvancedconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Return One Advanced Configuration Options for One Cluster
*ClustersApi* | [GetClusterStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#getclusterstatus) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status | Return Status of All Cluster Operations
*ClustersApi* | [GetSampleDatasetLoadStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#getsampledatasetloadstatus) | **Get** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId} | Check Status of Cluster Sample Dataset Request
*ClustersApi* | [ListCloudProviderRegions](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#listcloudproviderregions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/provider/regions | Return All Cloud Provider Regions
*ClustersApi* | [ListClustersForAllProjects](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#listclustersforallprojects) | **Get** /api/atlas/v2/clusters | Return All Authorized Clusters in All Projects
*ClustersApi* | [LoadSampleDataset](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#loadsampledataset) | **Post** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name} | Load Sample Dataset Request into Cluster
*ClustersApi* | [UpdateClusterAdvancedConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#updateclusteradvancedconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Update Advanced Configuration Options for One Cluster
*ClustersApi* | [UpgradeSharedCluster](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#upgradesharedcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade | Upgrade One Shared-tier Cluster
*ClustersApi* | [UpgradeSharedClusterToServerless](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClustersApi.md#upgradesharedclustertoserverless) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgradeToServerless | Upgrades One Shared-Tier Cluster to the Serverless Instance
*CustomDatabaseRolesApi* | [CreateCustomDatabaseRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomDatabaseRolesApi.md#createcustomdatabaserole) | **Post** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Create One Custom Role
*CustomDatabaseRolesApi* | [DeleteCustomDatabaseRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomDatabaseRolesApi.md#deletecustomdatabaserole) | **Delete** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project
*CustomDatabaseRolesApi* | [GetCustomDatabaseRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomDatabaseRolesApi.md#getcustomdatabaserole) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project
*CustomDatabaseRolesApi* | [ListCustomDatabaseRoles](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomDatabaseRolesApi.md#listcustomdatabaseroles) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project
*CustomDatabaseRolesApi* | [UpdateCustomDatabaseRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomDatabaseRolesApi.md#updatecustomdatabaserole) | **Patch** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project
*DataFederationApi* | [CreateDataFederationPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#createdatafederationprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Create One Federated Database Instance and Online Archive Private Endpoint for One Project
*DataFederationApi* | [CreateFederatedDatabase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#createfederateddatabase) | **Post** /api/atlas/v2/groups/{groupId}/dataFederation | Create One Federated Database Instance in One Project
*DataFederationApi* | [CreateOneDataFederationQueryLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#createonedatafederationquerylimit) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Configure One Query Limit for One Federated Database Instance
*DataFederationApi* | [DeleteDataFederationPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#deletedatafederationprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Remove One Federated Database Instance and Online Archive Private Endpoint from One Project
*DataFederationApi* | [DeleteFederatedDatabase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#deletefederateddatabase) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Remove One Federated Database Instance from One Project
*DataFederationApi* | [DeleteOneDataFederationInstanceQueryLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#deleteonedatafederationinstancequerylimit) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Delete One Query Limit For One Federated Database Instance
*DataFederationApi* | [DownloadFederatedDatabaseQueryLogs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#downloadfederateddatabasequerylogs) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz | Download Query Logs for One Federated Database Instance
*DataFederationApi* | [GetDataFederationPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#getdatafederationprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Return One Federated Database Instance and Online Archive Private Endpoint in One Project
*DataFederationApi* | [GetFederatedDatabase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#getfederateddatabase) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Return One Federated Database Instance in One Project
*DataFederationApi* | [ListDataFederationPrivateEndpoints](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#listdatafederationprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Return All Federated Database Instance and Online Archive Private Endpoints in One Project
*DataFederationApi* | [ListFederatedDatabases](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#listfederateddatabases) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation | Return All Federated Database Instances in One Project
*DataFederationApi* | [ReturnFederatedDatabaseQueryLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#returnfederateddatabasequerylimit) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Return One Federated Database Instance Query Limit for One Project
*DataFederationApi* | [ReturnFederatedDatabaseQueryLimits](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#returnfederateddatabasequerylimits) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits | Return All Query Limits for One Federated Database Instance
*DataFederationApi* | [UpdateFederatedDatabase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationApi.md#updatefederateddatabase) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Update One Federated Database Instance in One Project
*DataLakePipelinesApi* | [CreatePipeline](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#createpipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines | Create One Data Lake Pipeline
*DataLakePipelinesApi* | [DeletePipeline](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#deletepipeline) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Remove One Data Lake Pipeline
*DataLakePipelinesApi* | [DeletePipelineRunDataset](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#deletepipelinerundataset) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Delete Pipeline Run Dataset
*DataLakePipelinesApi* | [GetPipeline](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#getpipeline) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Return One Data Lake Pipeline
*DataLakePipelinesApi* | [GetPipelineRun](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#getpipelinerun) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Return One Data Lake Pipeline Run
*DataLakePipelinesApi* | [ListPipelineRuns](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#listpipelineruns) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs | Return All Data Lake Pipeline Runs from One Project
*DataLakePipelinesApi* | [ListPipelineSchedules](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#listpipelineschedules) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSchedules | Return Available Ingestion Schedules for One Data Lake Pipeline
*DataLakePipelinesApi* | [ListPipelineSnapshots](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#listpipelinesnapshots) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots | Return Available Backup Snapshots for One Data Lake Pipeline
*DataLakePipelinesApi* | [ListPipelines](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#listpipelines) | **Get** /api/atlas/v2/groups/{groupId}/pipelines | Return All Data Lake Pipelines from One Project
*DataLakePipelinesApi* | [PausePipeline](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#pausepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/pause | Pause One Data Lake Pipeline
*DataLakePipelinesApi* | [ResumePipeline](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#resumepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/resume | Resume One Data Lake Pipeline
*DataLakePipelinesApi* | [TriggerSnapshotIngestion](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#triggersnapshotingestion) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/trigger | Trigger on demand snapshot ingestion
*DataLakePipelinesApi* | [UpdatePipeline](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakePipelinesApi.md#updatepipeline) | **Patch** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Update One Data Lake Pipeline
*DatabaseUsersApi* | [CreateDatabaseUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DatabaseUsersApi.md#createdatabaseuser) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers | Create One Database User in One Project
*DatabaseUsersApi* | [DeleteDatabaseUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DatabaseUsersApi.md#deletedatabaseuser) | **Delete** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Remove One Database User from One Project
*DatabaseUsersApi* | [GetDatabaseUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DatabaseUsersApi.md#getdatabaseuser) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Return One Database User from One Project
*DatabaseUsersApi* | [ListDatabaseUsers](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DatabaseUsersApi.md#listdatabaseusers) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers | Return All Database Users from One Project
*DatabaseUsersApi* | [UpdateDatabaseUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DatabaseUsersApi.md#updatedatabaseuser) | **Patch** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Update One Database User in One Project
*EncryptionAtRestUsingCustomerKeyManagementApi* | [GetEncryptionAtRest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#getencryptionatrest) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project
*EncryptionAtRestUsingCustomerKeyManagementApi* | [UpdateEncryptionAtRest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#updateencryptionatrest) | **Patch** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project
*EventsApi* | [GetOrganizationEvent](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EventsApi.md#getorganizationevent) | **Get** /api/atlas/v2/orgs/{orgId}/events/{eventId} | Return One Event from One Organization
*EventsApi* | [GetProjectEvent](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EventsApi.md#getprojectevent) | **Get** /api/atlas/v2/groups/{groupId}/events/{eventId} | Return One Event from One Project
*EventsApi* | [ListOrganizationEvents](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EventsApi.md#listorganizationevents) | **Get** /api/atlas/v2/orgs/{orgId}/events | Return All Events from One Organization
*EventsApi* | [ListProjectEvents](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EventsApi.md#listprojectevents) | **Get** /api/atlas/v2/groups/{groupId}/events | Return All Events from One Project
*FederatedAuthenticationApi* | [CreateRoleMapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#createrolemapping) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Add One Role Mapping to One Organization
*FederatedAuthenticationApi* | [DeleteFederationApp](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#deletefederationapp) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId} | Delete the federation settings instance.
*FederatedAuthenticationApi* | [DeleteRoleMapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#deleterolemapping) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization
*FederatedAuthenticationApi* | [GetConnectedOrgConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#getconnectedorgconfig) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [GetFederationSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#getfederationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization
*FederatedAuthenticationApi* | [GetIdentityProvider](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#getidentityprovider) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return one identity provider from the specified federation.
*FederatedAuthenticationApi* | [GetIdentityProviderMetadata](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#getidentityprovidermetadata) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return the metadata of one identity provider in the specified federation.
*FederatedAuthenticationApi* | [GetRoleMapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#getrolemapping) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization
*FederatedAuthenticationApi* | [ListConnectedOrgConfigs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#listconnectedorgconfigs) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Connected Org Configs from the Federation
*FederatedAuthenticationApi* | [ListIdentityProviders](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#listidentityproviders) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Return all identity providers from the specified federation.
*FederatedAuthenticationApi* | [ListRoleMappings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#listrolemappings) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization
*FederatedAuthenticationApi* | [RemoveConnectedOrgConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#removeconnectedorgconfig) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [UpdateConnectedOrgConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#updateconnectedorgconfig) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [UpdateIdentityProvider](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#updateidentityprovider) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update the identity provider.
*FederatedAuthenticationApi* | [UpdateRoleMapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedAuthenticationApi.md#updaterolemapping) | **Put** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization
*GlobalClustersApi* | [CreateCustomZoneMapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GlobalClustersApi.md#createcustomzonemapping) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Add One Entry to One Custom Zone Mapping
*GlobalClustersApi* | [CreateManagedNamespace](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GlobalClustersApi.md#createmanagednamespace) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Create One Managed Namespace in One Global Multi-Cloud Cluster
*GlobalClustersApi* | [DeleteAllCustomZoneMappings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GlobalClustersApi.md#deleteallcustomzonemappings) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster
*GlobalClustersApi* | [DeleteManagedNamespace](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GlobalClustersApi.md#deletemanagednamespace) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Multi-Cloud Cluster
*GlobalClustersApi* | [GetManagedNamespace](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GlobalClustersApi.md#getmanagednamespace) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites | Return One Managed Namespace in One Global Multi-Cloud Cluster
*InvoicesApi* | [DownloadInvoiceCSV](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/InvoicesApi.md#downloadinvoicecsv) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Organization Invoice as CSV
*InvoicesApi* | [GetInvoice](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/InvoicesApi.md#getinvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId} | Return One Organization Invoice
*InvoicesApi* | [ListInvoices](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/InvoicesApi.md#listinvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices | Return All Invoices for One Organization
*InvoicesApi* | [ListPendingInvoices](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/InvoicesApi.md#listpendinginvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization
*LDAPConfigurationApi* | [DeleteLDAPConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LDAPConfigurationApi.md#deleteldapconfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove the Current LDAP User to DN Mapping
*LDAPConfigurationApi* | [GetLDAPConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LDAPConfigurationApi.md#getldapconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return the Current LDAP or X.509 Configuration
*LDAPConfigurationApi* | [GetLDAPConfigurationStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LDAPConfigurationApi.md#getldapconfigurationstatus) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return the Status of One Verify LDAP Configuration Request
*LDAPConfigurationApi* | [SaveLDAPConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LDAPConfigurationApi.md#saveldapconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Edit the LDAP or X.509 Configuration
*LDAPConfigurationApi* | [VerifyLDAPConfiguration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LDAPConfigurationApi.md#verifyldapconfiguration) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify the LDAP Configuration in One Project
*LegacyBackupApi* | [DeleteLegacySnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#deletelegacysnapshot) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Remove One Legacy Backup Snapshot
*LegacyBackupApi* | [GetLegacyBackupCheckpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#getlegacybackupcheckpoint) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId} | Return One Legacy Backup Checkpoint
*LegacyBackupApi* | [GetLegacyBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#getlegacybackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId} | Return One Legacy Backup Restore Job
*LegacyBackupApi* | [GetLegacySnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#getlegacysnapshot) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Return One Legacy Backup Snapshot
*LegacyBackupApi* | [GetLegacySnapshotSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#getlegacysnapshotschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Return One Snapshot Schedule
*LegacyBackupApi* | [ListLegacyBackupCheckpoints](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#listlegacybackupcheckpoints) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints | Return All Legacy Backup Checkpoints
*LegacyBackupApi* | [ListLegacyBackupRestoreJobs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#listlegacybackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Return All Legacy Backup Restore Jobs
*LegacyBackupApi* | [ListLegacySnapshots](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#listlegacysnapshots) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots | Return All Legacy Backup Snapshots
*LegacyBackupApi* | [UpdateLegacySnapshotRetention](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#updatelegacysnapshotretention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Change One Legacy Backup Snapshot Expiration
*LegacyBackupApi* | [UpdateLegacySnapshotSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupApi.md#updatelegacysnapshotschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Update Snapshot Schedule for One Cluster
*LegacyBackupRestoreJobsApi* | [CreateLegacyBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyBackupRestoreJobsApi.md#createlegacybackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Create One Legacy Backup Restore Job
*MaintenanceWindowsApi* | [DeferMaintenanceWindow](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MaintenanceWindowsApi.md#defermaintenancewindow) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/defer | Defer One Maintenance Window for One Project
*MaintenanceWindowsApi* | [GetMaintenanceWindow](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MaintenanceWindowsApi.md#getmaintenancewindow) | **Get** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Return One Maintenance Window for One Project
*MaintenanceWindowsApi* | [ResetMaintenanceWindow](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MaintenanceWindowsApi.md#resetmaintenancewindow) | **Delete** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Reset One Maintenance Window for One Project
*MaintenanceWindowsApi* | [ToggleMaintenanceAutoDefer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MaintenanceWindowsApi.md#togglemaintenanceautodefer) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/autoDefer | Toggle Automatic Deferral of Maintenance for One Project
*MaintenanceWindowsApi* | [UpdateMaintenanceWindow](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MaintenanceWindowsApi.md#updatemaintenancewindow) | **Patch** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Update Maintenance Window for One Project
*MongoDBCloudUsersApi* | [CreateUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MongoDBCloudUsersApi.md#createuser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User
*MongoDBCloudUsersApi* | [GetUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MongoDBCloudUsersApi.md#getuser) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User using Its ID
*MongoDBCloudUsersApi* | [GetUserByUsername](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MongoDBCloudUsersApi.md#getuserbyusername) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User using Their Username
*MonitoringAndLogsApi* | [GetAtlasProcess](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#getatlasprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId} | Return One MongoDB Process by ID
*MonitoringAndLogsApi* | [GetDatabase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#getdatabase) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName} | Return One Database for a MongoDB Process
*MonitoringAndLogsApi* | [GetDatabaseMeasurements](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#getdatabasemeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements | Return Measurements of One Database for One MongoDB Process
*MonitoringAndLogsApi* | [GetDiskMeasurements](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#getdiskmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements | Return Measurements of One Disk for One MongoDB Process
*MonitoringAndLogsApi* | [GetHostLogs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#gethostlogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Multi-Cloud Cluster Host in One Project
*MonitoringAndLogsApi* | [GetHostMeasurements](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#gethostmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/measurements | Return Measurements for One MongoDB Process
*MonitoringAndLogsApi* | [GetIndexMetrics](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#getindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements | Return Atlas Search Metrics for One Index in One Specified Namespace
*MonitoringAndLogsApi* | [GetMeasurements](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#getmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/measurements | Return Atlas Search Hardware and Status Metrics
*MonitoringAndLogsApi* | [ListAtlasProcesses](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#listatlasprocesses) | **Get** /api/atlas/v2/groups/{groupId}/processes | Return All MongoDB Processes in One Project
*MonitoringAndLogsApi* | [ListDatabases](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#listdatabases) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases | Return Available Databases for One MongoDB Process
*MonitoringAndLogsApi* | [ListDiskMeasurements](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#listdiskmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName} | Return Measurements of One Disk
*MonitoringAndLogsApi* | [ListDiskPartitions](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#listdiskpartitions) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks | Return Available Disks for One MongoDB Process
*MonitoringAndLogsApi* | [ListIndexMetrics](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#listindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements | Return All Atlas Search Index Metrics for One Namespace
*MonitoringAndLogsApi* | [ListMetricTypes](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonitoringAndLogsApi.md#listmetrictypes) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics | Return All Atlas Search Metric Types for One Process
*MultiCloudClustersApi* | [CreateCluster](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MultiCloudClustersApi.md#createcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters | Create One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [DeleteCluster](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MultiCloudClustersApi.md#deletecluster) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Remove One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [GetCluster](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MultiCloudClustersApi.md#getcluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Return One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [ListClusters](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MultiCloudClustersApi.md#listclusters) | **Get** /api/atlas/v2/groups/{groupId}/clusters | Return All Multi-Cloud Clusters from One Project
*MultiCloudClustersApi* | [TestFailover](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MultiCloudClustersApi.md#testfailover) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restartPrimaries | Test Failover for One Multi-Cloud Cluster
*MultiCloudClustersApi* | [UpdateCluster](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MultiCloudClustersApi.md#updatecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Modify One Multi-Cloud Cluster from One Project
*NetworkPeeringApi* | [CreatePeeringConnection](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#createpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/peers | Create One New Network Peering Connection
*NetworkPeeringApi* | [CreatePeeringContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#createpeeringcontainer) | **Post** /api/atlas/v2/groups/{groupId}/containers | Create One New Network Peering Container
*NetworkPeeringApi* | [DeletePeeringConnection](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#deletepeeringconnection) | **Delete** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Remove One Existing Network Peering Connection
*NetworkPeeringApi* | [DeletePeeringContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#deletepeeringcontainer) | **Delete** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container
*NetworkPeeringApi* | [DisablePeering](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#disablepeering) | **Patch** /api/atlas/v2/groups/{groupId}/privateIpMode | Disable Connect via Peering Only Mode for One Project
*NetworkPeeringApi* | [GetPeeringConnection](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#getpeeringconnection) | **Get** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project
*NetworkPeeringApi* | [GetPeeringContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#getpeeringcontainer) | **Get** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Return One Network Peering Container
*NetworkPeeringApi* | [ListPeeringConnections](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#listpeeringconnections) | **Get** /api/atlas/v2/groups/{groupId}/peers | Return All Network Peering Connections in One Project
*NetworkPeeringApi* | [ListPeeringContainerByCloudProvider](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#listpeeringcontainerbycloudprovider) | **Get** /api/atlas/v2/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider
*NetworkPeeringApi* | [ListPeeringContainers](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#listpeeringcontainers) | **Get** /api/atlas/v2/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project
*NetworkPeeringApi* | [UpdatePeeringConnection](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#updatepeeringconnection) | **Patch** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Update One New Network Peering Connection
*NetworkPeeringApi* | [UpdatePeeringContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#updatepeeringcontainer) | **Patch** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Update One Network Peering Container
*NetworkPeeringApi* | [VerifyConnectViaPeeringOnlyModeForOneProject](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPeeringApi.md#verifyconnectviapeeringonlymodeforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/privateIpMode | Verify Connect via Peering Only Mode for One Project
*OnlineArchiveApi* | [CreateOnlineArchive](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchiveApi.md#createonlinearchive) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive
*OnlineArchiveApi* | [DeleteOnlineArchive](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchiveApi.md#deleteonlinearchive) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive
*OnlineArchiveApi* | [DownloadOnlineArchiveQueryLogs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchiveApi.md#downloadonlinearchivequerylogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs
*OnlineArchiveApi* | [GetOnlineArchive](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchiveApi.md#getonlinearchive) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive
*OnlineArchiveApi* | [ListOnlineArchives](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchiveApi.md#listonlinearchives) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster
*OnlineArchiveApi* | [UpdateOnlineArchive](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchiveApi.md#updateonlinearchive) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive
*OrganizationsApi* | [CreateOrganization](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#createorganization) | **Post** /api/atlas/v2/orgs | Create One Organization
*OrganizationsApi* | [CreateOrganizationInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#createorganizationinvitation) | **Post** /api/atlas/v2/orgs/{orgId}/invites | Invite One MongoDB Cloud User to Join One Atlas Organization
*OrganizationsApi* | [DeleteOrganization](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#deleteorganization) | **Delete** /api/atlas/v2/orgs/{orgId} | Remove One Organization
*OrganizationsApi* | [DeleteOrganizationInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#deleteorganizationinvitation) | **Delete** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Cancel One Organization Invitation
*OrganizationsApi* | [GetOrganization](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#getorganization) | **Get** /api/atlas/v2/orgs/{orgId} | Return One Organization
*OrganizationsApi* | [GetOrganizationInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#getorganizationinvitation) | **Get** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Return One Organization Invitation
*OrganizationsApi* | [GetOrganizationSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#getorganizationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/settings | Return Settings for One Organization
*OrganizationsApi* | [ListOrganizationInvitations](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#listorganizationinvitations) | **Get** /api/atlas/v2/orgs/{orgId}/invites | Return All Organization Invitations
*OrganizationsApi* | [ListOrganizationProjects](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#listorganizationprojects) | **Get** /api/atlas/v2/orgs/{orgId}/groups | Return One or More Projects in One Organization
*OrganizationsApi* | [ListOrganizationUsers](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#listorganizationusers) | **Get** /api/atlas/v2/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization
*OrganizationsApi* | [ListOrganizations](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#listorganizations) | **Get** /api/atlas/v2/orgs | Return All Organizations
*OrganizationsApi* | [RenameOrganization](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#renameorganization) | **Patch** /api/atlas/v2/orgs/{orgId} | Rename One Organization
*OrganizationsApi* | [UpdateOrganizationInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#updateorganizationinvitation) | **Patch** /api/atlas/v2/orgs/{orgId}/invites | Update One Organization Invitation
*OrganizationsApi* | [UpdateOrganizationInvitationById](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#updateorganizationinvitationbyid) | **Patch** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Update One Organization Invitation by Invitation ID
*OrganizationsApi* | [UpdateOrganizationSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationsApi.md#updateorganizationsettings) | **Patch** /api/atlas/v2/orgs/{orgId}/settings | Update Settings for One Organization
*PerformanceAdvisorApi* | [DisableSlowOperationThresholding](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorApi.md#disableslowoperationthresholding) | **Delete** /api/atlas/v2/groups/{groupId}/managedSlowMs/disable | Disable Managed Slow Operation Threshold
*PerformanceAdvisorApi* | [EnableSlowOperationThresholding](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorApi.md#enableslowoperationthresholding) | **Post** /api/atlas/v2/groups/{groupId}/managedSlowMs/enable | Enable Managed Slow Operation Threshold
*PerformanceAdvisorApi* | [ListSlowQueries](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorApi.md#listslowqueries) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries
*PerformanceAdvisorApi* | [ListSlowQueryNamespaces](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorApi.md#listslowquerynamespaces) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host
*PerformanceAdvisorApi* | [ListSuggestedIndexes](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorApi.md#listsuggestedindexes) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return Suggested Indexes
*PrivateEndpointServicesApi* | [CreatePrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#createprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint | Create One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [CreatePrivateEndpointService](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#createprivateendpointservice) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService | Create One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [DeletePrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#deleteprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Remove One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [DeletePrivateEndpointService](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#deleteprivateendpointservice) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Remove One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [GetPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#getprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Return One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [GetPrivateEndpointService](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#getprivateendpointservice) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Return One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [GetRegionalizedPrivateEndpointSetting](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#getregionalizedprivateendpointsetting) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Return Regionalized Private Endpoint Status
*PrivateEndpointServicesApi* | [ListPrivateEndpointServices](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#listprivateendpointservices) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService | Return All Private Endpoint Services for One Provider
*PrivateEndpointServicesApi* | [ToggleRegionalizedPrivateEndpointSetting](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateEndpointServicesApi.md#toggleregionalizedprivateendpointsetting) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Toggle Regionalized Private Endpoint Status
*ProgrammaticAPIKeysApi* | [AddProjectApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#addprojectapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Assign One Organization API Key to One Project
*ProgrammaticAPIKeysApi* | [CreateApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#createapikey) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys | Create One Organization API Key
*ProgrammaticAPIKeysApi* | [CreateApiKeyAccessList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#createapikeyaccesslist) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Create Access List Entries for One Organization API Key
*ProgrammaticAPIKeysApi* | [CreateProjectApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#createprojectapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys | Create and Assign One Organization API Key to One Project
*ProgrammaticAPIKeysApi* | [DeleteApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#deleteapikey) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Remove One Organization API Key
*ProgrammaticAPIKeysApi* | [DeleteApiKeyAccessListEntry](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#deleteapikeyaccesslistentry) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Remove One Access List Entry for One Organization API Key
*ProgrammaticAPIKeysApi* | [GetApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#getapikey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Return One Organization API Key
*ProgrammaticAPIKeysApi* | [GetApiKeyAccessList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#getapikeyaccesslist) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Return One Access List Entry for One Organization API Key
*ProgrammaticAPIKeysApi* | [ListApiKeyAccessListsEntries](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#listapikeyaccesslistsentries) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Return All Access List Entries for One Organization API Key
*ProgrammaticAPIKeysApi* | [ListApiKeys](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#listapikeys) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys | Return All Organization API Keys
*ProgrammaticAPIKeysApi* | [ListProjectApiKeys](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#listprojectapikeys) | **Get** /api/atlas/v2/groups/{groupId}/apiKeys | Return All Organization API Keys Assigned to One Project
*ProgrammaticAPIKeysApi* | [RemoveProjectApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#removeprojectapikey) | **Delete** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Unassign One Organization API Key from One Project
*ProgrammaticAPIKeysApi* | [UpdateApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#updateapikey) | **Patch** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Update One Organization API Key
*ProgrammaticAPIKeysApi* | [UpdateApiKeyRoles](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProgrammaticAPIKeysApi.md#updateapikeyroles) | **Patch** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Update Roles of One Organization API Key to One Project
*ProjectIPAccessListApi* | [CreateProjectIpAccessList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectIPAccessListApi.md#createprojectipaccesslist) | **Post** /api/atlas/v2/groups/{groupId}/accessList | Add Entries to Project IP Access List
*ProjectIPAccessListApi* | [DeleteProjectIpAccessList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectIPAccessListApi.md#deleteprojectipaccesslist) | **Delete** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Remove One Entry from One Project IP Access List
*ProjectIPAccessListApi* | [GetProjectIpAccessListStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectIPAccessListApi.md#getprojectipaccessliststatus) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue}/status | Return Status of One Project IP Access List Entry
*ProjectIPAccessListApi* | [GetProjectIpList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectIPAccessListApi.md#getprojectiplist) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Return One Project IP Access List Entry
*ProjectIPAccessListApi* | [ListProjectIpAccessLists](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectIPAccessListApi.md#listprojectipaccesslists) | **Get** /api/atlas/v2/groups/{groupId}/accessList | Return Project IP Access List
*ProjectsApi* | [CreateProject](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#createproject) | **Post** /api/atlas/v2/groups | Create One Project
*ProjectsApi* | [CreateProjectInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#createprojectinvitation) | **Post** /api/atlas/v2/groups/{groupId}/invites | Invite One MongoDB Cloud User to Join One Project
*ProjectsApi* | [DeleteProject](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#deleteproject) | **Delete** /api/atlas/v2/groups/{groupId} | Remove One Project
*ProjectsApi* | [DeleteProjectInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#deleteprojectinvitation) | **Delete** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Cancel One Project Invitation
*ProjectsApi* | [DeleteProjectLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#deleteprojectlimit) | **Delete** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Remove One Project Limit
*ProjectsApi* | [GetProject](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#getproject) | **Get** /api/atlas/v2/groups/{groupId} | Return One Project
*ProjectsApi* | [GetProjectByName](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#getprojectbyname) | **Get** /api/atlas/v2/groups/byName/{groupName} | Return One Project using Its Name
*ProjectsApi* | [GetProjectInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#getprojectinvitation) | **Get** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Return One Project Invitation
*ProjectsApi* | [GetProjectLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#getprojectlimit) | **Get** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Return One Limit for One Project
*ProjectsApi* | [GetProjectSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#getprojectsettings) | **Get** /api/atlas/v2/groups/{groupId}/settings | Return One Project Settings
*ProjectsApi* | [ListProjectInvitations](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#listprojectinvitations) | **Get** /api/atlas/v2/groups/{groupId}/invites | Return All Project Invitations
*ProjectsApi* | [ListProjectLimits](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#listprojectlimits) | **Get** /api/atlas/v2/groups/{groupId}/limits | Return All Limits for One Project
*ProjectsApi* | [ListProjectUsers](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#listprojectusers) | **Get** /api/atlas/v2/groups/{groupId}/users | Return All Users in One Project
*ProjectsApi* | [ListProjects](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#listprojects) | **Get** /api/atlas/v2/groups | Return All Projects
*ProjectsApi* | [RemoveProjectUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#removeprojectuser) | **Delete** /api/atlas/v2/groups/{groupId}/users/{userId} | Remove One User from One Project
*ProjectsApi* | [SetProjectLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#setprojectlimit) | **Patch** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Set One Project Limit
*ProjectsApi* | [UpdateProject](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#updateproject) | **Patch** /api/atlas/v2/groups/{groupId} | Update One Project Name
*ProjectsApi* | [UpdateProjectInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#updateprojectinvitation) | **Patch** /api/atlas/v2/groups/{groupId}/invites | Update One Project Invitation
*ProjectsApi* | [UpdateProjectInvitationById](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#updateprojectinvitationbyid) | **Patch** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Update One Project Invitation by Invitation ID
*ProjectsApi* | [UpdateProjectSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectsApi.md#updateprojectsettings) | **Patch** /api/atlas/v2/groups/{groupId}/settings | Update One Project Settings
*RollingIndexApi* | [CreateRollingIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RollingIndexApi.md#createrollingindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/index | Create One Rolling Index
*RootApi* | [GetSystemStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RootApi.md#getsystemstatus) | **Get** /api/atlas/v2 | Return the status of this MongoDB application
*ServerlessInstancesApi* | [CreateServerlessInstance](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstancesApi.md#createserverlessinstance) | **Post** /api/atlas/v2/groups/{groupId}/serverless | Create One Serverless Instance in One Project
*ServerlessInstancesApi* | [DeleteServerlessInstance](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstancesApi.md#deleteserverlessinstance) | **Delete** /api/atlas/v2/groups/{groupId}/serverless/{name} | Remove One Serverless Instance from One Project
*ServerlessInstancesApi* | [GetServerlessInstance](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstancesApi.md#getserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{name} | Return One Serverless Instance from One Project
*ServerlessInstancesApi* | [ListServerlessInstances](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstancesApi.md#listserverlessinstances) | **Get** /api/atlas/v2/groups/{groupId}/serverless | Return All Serverless Instances from One Project
*ServerlessInstancesApi* | [UpdateServerlessInstance](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstancesApi.md#updateserverlessinstance) | **Patch** /api/atlas/v2/groups/{groupId}/serverless/{name} | Update One Serverless Instance in One Project
*ServerlessPrivateEndpointsApi* | [CreateServerlessPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessPrivateEndpointsApi.md#createserverlessprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Create One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [DeleteServerlessPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessPrivateEndpointsApi.md#deleteserverlessprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Remove One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [GetServerlessPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessPrivateEndpointsApi.md#getserverlessprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Return One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [ListServerlessPrivateEndpoints](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessPrivateEndpointsApi.md#listserverlessprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Return All Private Endpoints for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [UpdateServerlessPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessPrivateEndpointsApi.md#updateserverlessprivateendpoint) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Update One Private Endpoint for One Serverless Instance
*SharedTierRestoreJobsApi* | [CreateSharedClusterBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SharedTierRestoreJobsApi.md#createsharedclusterbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore | Create One Restore Job from One M2 or M5 Cluster
*SharedTierRestoreJobsApi* | [GetSharedClusterBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SharedTierRestoreJobsApi.md#getsharedclusterbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId} | Return One Restore Job for One M2 or M5 Cluster
*SharedTierRestoreJobsApi* | [ListSharedClusterBackupRestoreJobs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SharedTierRestoreJobsApi.md#listsharedclusterbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores | Return All Restore Jobs for One M2 or M5 Cluster
*SharedTierSnapshotsApi* | [DownloadSharedClusterBackup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SharedTierSnapshotsApi.md#downloadsharedclusterbackup) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/download | Download One M2 or M5 Cluster Snapshot
*SharedTierSnapshotsApi* | [GetSharedClusterBackup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SharedTierSnapshotsApi.md#getsharedclusterbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId} | Return One Snapshot for One M2 or M5 Cluster
*SharedTierSnapshotsApi* | [ListSharedClusterBackups](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SharedTierSnapshotsApi.md#listsharedclusterbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots | Return All Snapshots for One M2 or M5 Cluster
*TeamsApi* | [AddAllTeamsToProject](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#addallteamstoproject) | **Post** /api/atlas/v2/groups/{groupId}/teams | Add One or More Teams to One Project
*TeamsApi* | [AddTeamUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#addteamuser) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Assign MongoDB Cloud Users from One Organization to One Team
*TeamsApi* | [CreateTeam](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#createteam) | **Post** /api/atlas/v2/orgs/{orgId}/teams | Create One Team in One Organization
*TeamsApi* | [DeleteTeam](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#deleteteam) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Remove One Team from One Organization
*TeamsApi* | [GetTeamById](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#getteambyid) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Return One Team using its ID
*TeamsApi* | [GetTeamByName](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#getteambyname) | **Get** /api/atlas/v2/orgs/{orgId}/teams/byName/{teamName} | Return One Team using its Name
*TeamsApi* | [ListOrganizationTeams](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#listorganizationteams) | **Get** /api/atlas/v2/orgs/{orgId}/teams | Return All Teams in One Organization
*TeamsApi* | [ListProjectTeams](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#listprojectteams) | **Get** /api/atlas/v2/groups/{groupId}/teams | Return All Teams in One Project
*TeamsApi* | [ListTeamUsers](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#listteamusers) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team
*TeamsApi* | [RemoveProjectTeam](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#removeprojectteam) | **Delete** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Remove One Team from One Project
*TeamsApi* | [RemoveTeamUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#removeteamuser) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users/{userId} | Remove One MongoDB Cloud User from One Team
*TeamsApi* | [RenameTeam](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#renameteam) | **Patch** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Rename One Team
*TeamsApi* | [UpdateTeamRoles](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamsApi.md#updateteamroles) | **Patch** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Update Team Roles in One Project
*TestApi* | [VersionedExample](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TestApi.md#versionedexample) | **Get** /api/atlas/v2/example/info | Example resource info for versioning of the Atlas API
*ThirdPartyIntegrationsApi* | [CreateThirdPartyIntegration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ThirdPartyIntegrationsApi.md#createthirdpartyintegration) | **Post** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Configure One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [DeleteThirdPartyIntegration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ThirdPartyIntegrationsApi.md#deletethirdpartyintegration) | **Delete** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Remove One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [GetThirdPartyIntegration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ThirdPartyIntegrationsApi.md#getthirdpartyintegration) | **Get** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Return One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [ListThirdPartyIntegrations](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ThirdPartyIntegrationsApi.md#listthirdpartyintegrations) | **Get** /api/atlas/v2/groups/{groupId}/integrations | Return All Active Third-Party Service Integrations
*ThirdPartyIntegrationsApi* | [UpdateThirdPartyIntegration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ThirdPartyIntegrationsApi.md#updatethirdpartyintegration) | **Put** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Update One Third-Party Service Integration
*X509AuthenticationApi* | [CreateDatabaseUserCertificate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/X509AuthenticationApi.md#createdatabaseusercertificate) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Create One X.509 Certificate for One MongoDB User
*X509AuthenticationApi* | [DisableCustomerManagedX509](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/X509AuthenticationApi.md#disablecustomermanagedx509) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/customerX509 | Disable Customer-Managed X.509
*X509AuthenticationApi* | [ListDatabaseUserCertificates](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/X509AuthenticationApi.md#listdatabaseusercertificates) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Return All X.509 Certificates Assigned to One MongoDB User


## Documentation For Models

 - [AWSAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSAutoScaling.md)
 - [AWSCloudProviderContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSCloudProviderContainer.md)
 - [AWSComputeAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSComputeAutoScaling.md)
 - [AWSCustomDNSEnabled](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSCustomDNSEnabled.md)
 - [AWSInterfaceEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSInterfaceEndpoint.md)
 - [AWSKMS](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSKMS.md)
 - [AWSPeerVpc](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSPeerVpc.md)
 - [AWSPrivateLinkConnection](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSPrivateLinkConnection.md)
 - [AWSProviderSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AWSProviderSettings.md)
 - [AccessListItem](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AccessListItem.md)
 - [AddUserToTeam](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AddUserToTeam.md)
 - [AlertConfigViewForNdsGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertConfigViewForNdsGroup.md)
 - [AlertViewForNdsGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AlertViewForNdsGroup.md)
 - [ApiUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ApiUser.md)
 - [AppUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AppUser.md)
 - [AuditLog](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AuditLog.md)
 - [AutoExportPolicy](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AutoExportPolicy.md)
 - [AutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AutoScaling.md)
 - [AutoScalingV15](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AutoScalingV15.md)
 - [AvailableDeployment](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AvailableDeployment.md)
 - [AvailableProject](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AvailableProject.md)
 - [AvailableRegion](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AvailableRegion.md)
 - [AzureAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzureAutoScaling.md)
 - [AzureCloudProviderContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzureCloudProviderContainer.md)
 - [AzureComputeAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzureComputeAutoScaling.md)
 - [AzureKeyVault](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzureKeyVault.md)
 - [AzurePeerNetwork](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzurePeerNetwork.md)
 - [AzurePrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzurePrivateEndpoint.md)
 - [AzurePrivateLinkConnection](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzurePrivateLinkConnection.md)
 - [AzureProviderSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/AzureProviderSettings.md)
 - [BSONTimestamp](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/BSONTimestamp.md)
 - [BiConnector](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/BiConnector.md)
 - [CharFilterhtmlStrip](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CharFilterhtmlStrip.md)
 - [CharFiltericuNormalize](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CharFiltericuNormalize.md)
 - [CharFiltermapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CharFiltermapping.md)
 - [CharFiltermappingMappings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CharFiltermappingMappings.md)
 - [CharFilterpersian](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CharFilterpersian.md)
 - [Checkpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Checkpoint.md)
 - [CheckpointPart](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CheckpointPart.md)
 - [CloudProviderAccess](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccess.md)
 - [CloudProviderAccessAWSIAMRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessAWSIAMRole.md)
 - [CloudProviderAccessAzureServicePrincipal](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessAzureServicePrincipal.md)
 - [CloudProviderAccessDataLakeFeatureUsage](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessDataLakeFeatureUsage.md)
 - [CloudProviderAccessEncryptionAtRestFeatureUsage](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessEncryptionAtRestFeatureUsage.md)
 - [CloudProviderAccessExportSnapshotFeatureUsage](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessExportSnapshotFeatureUsage.md)
 - [CloudProviderAccessFeatureUsage](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessFeatureUsage.md)
 - [CloudProviderAccessFeatureUsageDataLakeFeatureId](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessFeatureUsageDataLakeFeatureId.md)
 - [CloudProviderAccessFeatureUsageExportSnapshotFeatureId](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessFeatureUsageExportSnapshotFeatureId.md)
 - [CloudProviderAccessRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderAccessRole.md)
 - [CloudProviderContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CloudProviderContainer.md)
 - [Cluster](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Cluster.md)
 - [ClusterDescriptionConnectionStrings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterDescriptionConnectionStrings.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterDescriptionConnectionStringsPrivateEndpoint.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpointEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ClusterDescriptionProcessArgs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterDescriptionProcessArgs.md)
 - [ClusterDescriptionV15](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterDescriptionV15.md)
 - [ClusterOutageSimulation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterOutageSimulation.md)
 - [ClusterOutageSimulationOutageFilter](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterOutageSimulationOutageFilter.md)
 - [ClusterProviderSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterProviderSettings.md)
 - [ClusterStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ClusterStatus.md)
 - [Collation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Collation.md)
 - [ComputeAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ComputeAutoScaling.md)
 - [ComputeAutoScalingV15](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ComputeAutoScalingV15.md)
 - [ConnectedOrgConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ConnectedOrgConfig.md)
 - [ContainerPeer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ContainerPeer.md)
 - [CreateAWSEndpointRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateAWSEndpointRequest.md)
 - [CreateApiKey](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateApiKey.md)
 - [CreateAzureEndpointRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateAzureEndpointRequest.md)
 - [CreateEndpointRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateEndpointRequest.md)
 - [CreateEndpointServiceRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateEndpointServiceRequest.md)
 - [CreateGCPEndpointGroupRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateGCPEndpointGroupRequest.md)
 - [CreateGCPForwardingRuleRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateGCPForwardingRuleRequest.md)
 - [CreateOrganizationRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateOrganizationRequest.md)
 - [CreateOrganizationResponse](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CreateOrganizationResponse.md)
 - [Criteria](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Criteria.md)
 - [CustomCriteria](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomCriteria.md)
 - [CustomDBRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomDBRole.md)
 - [CustomerX509](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/CustomerX509.md)
 - [DBAction](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DBAction.md)
 - [DBResource](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DBResource.md)
 - [DLSIngestionSink](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DLSIngestionSink.md)
 - [DailySchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DailySchedule.md)
 - [DataFederationQueryLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationQueryLimit.md)
 - [DataFederationTenantQueryLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataFederationTenantQueryLimit.md)
 - [DataLakeAWSCloudProviderConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeAWSCloudProviderConfig.md)
 - [DataLakeAtlasStore](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeAtlasStore.md)
 - [DataLakeAtlasStoreReadPreference](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeAtlasStoreReadPreference.md)
 - [DataLakeAtlasStoreReadPreferenceTag](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeAtlasStoreReadPreferenceTag.md)
 - [DataLakeCloudProviderConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeCloudProviderConfig.md)
 - [DataLakeDataProcessRegion](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeDataProcessRegion.md)
 - [DataLakeDatabase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeDatabase.md)
 - [DataLakeDatabaseCollection](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeDatabaseCollection.md)
 - [DataLakeDatabaseDataSource](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeDatabaseDataSource.md)
 - [DataLakeHTTPStore](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeHTTPStore.md)
 - [DataLakeOnlineArchiveStore](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeOnlineArchiveStore.md)
 - [DataLakeRegion](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeRegion.md)
 - [DataLakeS3Store](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeS3Store.md)
 - [DataLakeStorage](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeStorage.md)
 - [DataLakeStore](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeStore.md)
 - [DataLakeTenant](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeTenant.md)
 - [DataLakeView](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataLakeView.md)
 - [DataMetricThreshold](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataMetricThreshold.md)
 - [DataMetricUnits](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataMetricUnits.md)
 - [DataProtectionSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DataProtectionSettings.md)
 - [Database](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Database.md)
 - [DatabaseUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DatabaseUser.md)
 - [Datadog](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Datadog.md)
 - [DatadogNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DatadogNotification.md)
 - [DateCriteria](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DateCriteria.md)
 - [DedicatedHardwareSpec](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DedicatedHardwareSpec.md)
 - [DefaultLimit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DefaultLimit.md)
 - [DefaultSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DefaultSchedule.md)
 - [DeleteCopiedBackups](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DeleteCopiedBackups.md)
 - [Destination](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Destination.md)
 - [DiskBackupBaseRestoreMember](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupBaseRestoreMember.md)
 - [DiskBackupCopySetting](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupCopySetting.md)
 - [DiskBackupExportJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupExportJob.md)
 - [DiskBackupExportJobRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupExportJobRequest.md)
 - [DiskBackupOnDemandSnapshotRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupOnDemandSnapshotRequest.md)
 - [DiskBackupReplicaSet](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupReplicaSet.md)
 - [DiskBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupRestoreJob.md)
 - [DiskBackupShardedClusterSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupShardedClusterSnapshot.md)
 - [DiskBackupShardedClusterSnapshotMember](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupShardedClusterSnapshotMember.md)
 - [DiskBackupSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupSnapshot.md)
 - [DiskBackupSnapshotAWSExportBucket](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupSnapshotAWSExportBucket.md)
 - [DiskBackupSnapshotSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskBackupSnapshotSchedule.md)
 - [DiskGBAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskGBAutoScaling.md)
 - [DiskPartition](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/DiskPartition.md)
 - [EmailNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EmailNotification.md)
 - [EncryptionAtRest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EncryptionAtRest.md)
 - [Endpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Endpoint.md)
 - [EndpointService](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EndpointService.md)
 - [Error](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Error.md)
 - [EventViewForNdsGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EventViewForNdsGroup.md)
 - [EventViewForOrg](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/EventViewForOrg.md)
 - [ExampleResourceResponseView20230201](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ExampleResourceResponseView20230201.md)
 - [ExportStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ExportStatus.md)
 - [FTSAnalyzers](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSAnalyzers.md)
 - [FTSAnalyzersCharFiltersInner](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSAnalyzersCharFiltersInner.md)
 - [FTSAnalyzersTokenFiltersInner](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSAnalyzersTokenFiltersInner.md)
 - [FTSAnalyzersTokenizer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSAnalyzersTokenizer.md)
 - [FTSIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSIndex.md)
 - [FTSMappings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSMappings.md)
 - [FTSMetric](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSMetric.md)
 - [FTSMetrics](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSMetrics.md)
 - [FTSSynonymMappingDefinition](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FTSSynonymMappingDefinition.md)
 - [FederatedUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FederatedUser.md)
 - [FieldTransformation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FieldTransformation.md)
 - [FreeAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FreeAutoScaling.md)
 - [FreeProviderSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/FreeProviderSettings.md)
 - [GCPAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPAutoScaling.md)
 - [GCPCloudProviderContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPCloudProviderContainer.md)
 - [GCPComputeAutoScaling](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPComputeAutoScaling.md)
 - [GCPConsumerForwardingRule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPConsumerForwardingRule.md)
 - [GCPEndpointGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPEndpointGroup.md)
 - [GCPEndpointService](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPEndpointService.md)
 - [GCPPeerVpc](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPPeerVpc.md)
 - [GCPProviderSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GCPProviderSettings.md)
 - [GeoSharding](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GeoSharding.md)
 - [GoogleCloudKMS](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GoogleCloudKMS.md)
 - [Group](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Group.md)
 - [GroupInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupInvitation.md)
 - [GroupInvitationRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupInvitationRequest.md)
 - [GroupInvitationUpdateRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupInvitationUpdateRequest.md)
 - [GroupMaintenanceWindow](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupMaintenanceWindow.md)
 - [GroupName](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupName.md)
 - [GroupNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupNotification.md)
 - [GroupPaginatedEvent](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupPaginatedEvent.md)
 - [GroupSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/GroupSettings.md)
 - [HardwareSpec](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/HardwareSpec.md)
 - [HipChatNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/HipChatNotification.md)
 - [HostMetricValue](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/HostMetricValue.md)
 - [HostViewAtlas](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/HostViewAtlas.md)
 - [IdentityProvider](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/IdentityProvider.md)
 - [IndexOptions](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/IndexOptions.md)
 - [IndexRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/IndexRequest.md)
 - [IngestionPipeline](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/IngestionPipeline.md)
 - [IngestionPipelineRun](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/IngestionPipelineRun.md)
 - [IngestionSink](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/IngestionSink.md)
 - [IngestionSource](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/IngestionSource.md)
 - [InheritedRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/InheritedRole.md)
 - [InstanceSize](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/InstanceSize.md)
 - [Integration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Integration.md)
 - [Invoice](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Invoice.md)
 - [Key](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Key.md)
 - [Label](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Label.md)
 - [LegacyClusterDescription](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyClusterDescription.md)
 - [LegacyReplicationSpec](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LegacyReplicationSpec.md)
 - [Limit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Limit.md)
 - [LineItem](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LineItem.md)
 - [Link](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Link.md)
 - [LinkAtlas](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LinkAtlas.md)
 - [LiveMigrationRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LiveMigrationRequest.md)
 - [LiveMigrationResponse](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/LiveMigrationResponse.md)
 - [ManagedNamespace](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ManagedNamespace.md)
 - [ManagedNamespaces](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ManagedNamespaces.md)
 - [MatcherField](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MatcherField.md)
 - [Measurement](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Measurement.md)
 - [MeasurementViewAtlas](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MeasurementViewAtlas.md)
 - [MeasurementsGeneralViewAtlas](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MeasurementsGeneralViewAtlas.md)
 - [MeasurementsIndexes](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MeasurementsIndexes.md)
 - [MeasurementsNonIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MeasurementsNonIndex.md)
 - [MetricDataPoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MetricDataPoint.md)
 - [MetricDataPointViewAtlas](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MetricDataPointViewAtlas.md)
 - [MicrosoftTeams](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MicrosoftTeams.md)
 - [MicrosoftTeamsNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MicrosoftTeamsNotification.md)
 - [MongoDBAccessLogs](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MongoDBAccessLogs.md)
 - [MongoDBAccessLogsList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MongoDBAccessLogsList.md)
 - [MonthlySchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/MonthlySchedule.md)
 - [NDSLDAP](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NDSLDAP.md)
 - [NDSLDAPVerifyConnectivityJobRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NDSLDAPVerifyConnectivityJobRequest.md)
 - [NDSLDAPVerifyConnectivityJobRequestParams](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NDSLDAPVerifyConnectivityJobRequestParams.md)
 - [NDSLDAPVerifyConnectivityJobRequestValidation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NDSLDAPVerifyConnectivityJobRequestValidation.md)
 - [NDSLabel](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NDSLabel.md)
 - [NDSUserToDNMapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NDSUserToDNMapping.md)
 - [NamespaceObj](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NamespaceObj.md)
 - [Namespaces](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Namespaces.md)
 - [NetworkPermissionEntry](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPermissionEntry.md)
 - [NetworkPermissionEntryStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NetworkPermissionEntryStatus.md)
 - [NewRelic](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NewRelic.md)
 - [NotificationViewForNdsGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/NotificationViewForNdsGroup.md)
 - [OnDemandCpsSnapshotSource](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnDemandCpsSnapshotSource.md)
 - [OnlineArchive](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchive.md)
 - [OnlineArchiveSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OnlineArchiveSchedule.md)
 - [Operator](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Operator.md)
 - [OpsGenie](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OpsGenie.md)
 - [OpsGenieNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OpsGenieNotification.md)
 - [OrgFederationSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrgFederationSettings.md)
 - [OrgGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrgGroup.md)
 - [OrgNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrgNotification.md)
 - [OrgPaginatedEvent](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrgPaginatedEvent.md)
 - [Organization](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Organization.md)
 - [OrganizationInvitation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationInvitation.md)
 - [OrganizationInvitationRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationInvitationRequest.md)
 - [OrganizationInvitationUpdateRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationInvitationUpdateRequest.md)
 - [OrganizationSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/OrganizationSettings.md)
 - [PagerDuty](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PagerDuty.md)
 - [PagerDutyNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PagerDutyNotification.md)
 - [PaginatedAlert](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedAlert.md)
 - [PaginatedAlertConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedAlertConfig.md)
 - [PaginatedApiApiUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiApiUser.md)
 - [PaginatedApiAppUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiAppUser.md)
 - [PaginatedApiAtlasCheckpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiAtlasCheckpoint.md)
 - [PaginatedApiAtlasDatabaseUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiAtlasDatabaseUser.md)
 - [PaginatedApiAtlasDiskBackupExportJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiAtlasDiskBackupExportJob.md)
 - [PaginatedApiAtlasProviderRegions](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiAtlasProviderRegions.md)
 - [PaginatedApiAtlasServerlessBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiAtlasServerlessBackupRestoreJob.md)
 - [PaginatedApiAtlasServerlessBackupSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiAtlasServerlessBackupSnapshot.md)
 - [PaginatedApiInvoice](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiInvoice.md)
 - [PaginatedApiUserAccessList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedApiUserAccessList.md)
 - [PaginatedAppUser](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedAppUser.md)
 - [PaginatedAtlasGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedAtlasGroup.md)
 - [PaginatedBackupSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedBackupSnapshot.md)
 - [PaginatedBackupSnapshotExportBucket](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedBackupSnapshotExportBucket.md)
 - [PaginatedCloudBackupReplicaSet](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedCloudBackupReplicaSet.md)
 - [PaginatedCloudBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedCloudBackupRestoreJob.md)
 - [PaginatedCloudBackupShardedClusterSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedCloudBackupShardedClusterSnapshot.md)
 - [PaginatedCloudProviderContainer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedCloudProviderContainer.md)
 - [PaginatedClusterDescriptionV15](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedClusterDescriptionV15.md)
 - [PaginatedContainerPeer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedContainerPeer.md)
 - [PaginatedDatabase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedDatabase.md)
 - [PaginatedDiskPartition](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedDiskPartition.md)
 - [PaginatedHostViewAtlas](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedHostViewAtlas.md)
 - [PaginatedIntegration](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedIntegration.md)
 - [PaginatedNetworkAccess](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedNetworkAccess.md)
 - [PaginatedOnlineArchive](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedOnlineArchive.md)
 - [PaginatedOrgGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedOrgGroup.md)
 - [PaginatedOrganization](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedOrganization.md)
 - [PaginatedPipelineRun](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedPipelineRun.md)
 - [PaginatedPrivateNetworkEndpointIdEntry](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedPrivateNetworkEndpointIdEntry.md)
 - [PaginatedRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedRestoreJob.md)
 - [PaginatedServerlessInstanceDescription](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedServerlessInstanceDescription.md)
 - [PaginatedSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedSnapshot.md)
 - [PaginatedTeam](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedTeam.md)
 - [PaginatedTeamRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedTeamRole.md)
 - [PaginatedTenantRestore](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedTenantRestore.md)
 - [PaginatedTenantSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedTenantSnapshot.md)
 - [PaginatedUserCert](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PaginatedUserCert.md)
 - [PartitionField](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PartitionField.md)
 - [Payment](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Payment.md)
 - [PemFileInfo](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PemFileInfo.md)
 - [PerformanceAdvisorIndex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorIndex.md)
 - [PerformanceAdvisorOpStats](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorOpStats.md)
 - [PerformanceAdvisorOperation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorOperation.md)
 - [PerformanceAdvisorResponse](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorResponse.md)
 - [PerformanceAdvisorShape](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorShape.md)
 - [PerformanceAdvisorSlowQuery](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorSlowQuery.md)
 - [PerformanceAdvisorSlowQueryList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PerformanceAdvisorSlowQueryList.md)
 - [PeriodicCpsSnapshotSource](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PeriodicCpsSnapshotSource.md)
 - [PipelineRunStats](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PipelineRunStats.md)
 - [Policy](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Policy.md)
 - [PolicyItem](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PolicyItem.md)
 - [PrivateIPMode](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateIPMode.md)
 - [PrivateNetworkEndpointIdEntry](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/PrivateNetworkEndpointIdEntry.md)
 - [ProjectSettingItem](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProjectSettingItem.md)
 - [Prometheus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Prometheus.md)
 - [ProviderInstanceSize](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProviderInstanceSize.md)
 - [ProviderRegions](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ProviderRegions.md)
 - [RPUMetricThreshold](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RPUMetricThreshold.md)
 - [Raw](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Raw.md)
 - [RawMetricThreshold](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RawMetricThreshold.md)
 - [RawMetricUnits](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RawMetricUnits.md)
 - [Refund](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Refund.md)
 - [RegionConfig](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RegionConfig.md)
 - [RegionSpec](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RegionSpec.md)
 - [ReplicationSpec](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ReplicationSpec.md)
 - [ResourceEventType](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ResourceEventType.md)
 - [RestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RestoreJob.md)
 - [RestoreJobDelivery](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RestoreJobDelivery.md)
 - [RestoreJobFileHash](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RestoreJobFileHash.md)
 - [Role](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Role.md)
 - [RoleAssignment](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RoleAssignment.md)
 - [RoleMapping](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/RoleMapping.md)
 - [SMSNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SMSNotification.md)
 - [SamlIdentityProviderUpdate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SamlIdentityProviderUpdate.md)
 - [SampleDatasetStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SampleDatasetStatus.md)
 - [ServerlessAWSTenantEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessAWSTenantEndpoint.md)
 - [ServerlessAWSTenantEndpointUpdate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessAWSTenantEndpointUpdate.md)
 - [ServerlessAzureTenantEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessAzureTenantEndpoint.md)
 - [ServerlessAzureTenantEndpointUpdate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessAzureTenantEndpointUpdate.md)
 - [ServerlessBackupOptions](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessBackupOptions.md)
 - [ServerlessBackupRestoreJob](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessBackupRestoreJob.md)
 - [ServerlessBackupSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessBackupSnapshot.md)
 - [ServerlessEventTypeViewAlertable](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessEventTypeViewAlertable.md)
 - [ServerlessInstanceDescription](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstanceDescription.md)
 - [ServerlessInstanceDescriptionConnectionStrings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstanceDescriptionConnectionStrings.md)
 - [ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint.md)
 - [ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ServerlessInstanceDescriptionCreate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstanceDescriptionCreate.md)
 - [ServerlessInstanceDescriptionUpdate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessInstanceDescriptionUpdate.md)
 - [ServerlessMetricThreshold](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessMetricThreshold.md)
 - [ServerlessMetricUnits](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessMetricUnits.md)
 - [ServerlessProviderSettings](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessProviderSettings.md)
 - [ServerlessTenantEndpoint](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessTenantEndpoint.md)
 - [ServerlessTenantEndpointCreate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessTenantEndpointCreate.md)
 - [ServerlessTenantEndpointUpdate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ServerlessTenantEndpointUpdate.md)
 - [Slack](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Slack.md)
 - [SlackNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SlackNotification.md)
 - [Snapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Snapshot.md)
 - [SnapshotPart](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SnapshotPart.md)
 - [SnapshotRetention](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SnapshotRetention.md)
 - [SnapshotSchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SnapshotSchedule.md)
 - [Source](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Source.md)
 - [SynonymSource](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SynonymSource.md)
 - [SystemStatus](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/SystemStatus.md)
 - [Tag](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Tag.md)
 - [TargetOrg](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TargetOrg.md)
 - [TargetOrgRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TargetOrgRequest.md)
 - [Team](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Team.md)
 - [TeamNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamNotification.md)
 - [TeamResponse](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamResponse.md)
 - [TeamRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TeamRole.md)
 - [TenantRestore](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TenantRestore.md)
 - [TenantSnapshot](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TenantSnapshot.md)
 - [ThresholdViewInteger](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/ThresholdViewInteger.md)
 - [TimeMetricThreshold](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TimeMetricThreshold.md)
 - [TimeMetricUnits](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TimeMetricUnits.md)
 - [Toggle](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Toggle.md)
 - [TokenFilterasciiFolding](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilterasciiFolding.md)
 - [TokenFilterdaitchMokotoffSoundex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilterdaitchMokotoffSoundex.md)
 - [TokenFilteredgeGram](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilteredgeGram.md)
 - [TokenFiltericuFolding](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFiltericuFolding.md)
 - [TokenFiltericuNormalizer](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFiltericuNormalizer.md)
 - [TokenFilterlength](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilterlength.md)
 - [TokenFilterlowercase](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilterlowercase.md)
 - [TokenFilternGram](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilternGram.md)
 - [TokenFilterregex](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilterregex.md)
 - [TokenFilterreverse](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilterreverse.md)
 - [TokenFiltershingle](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFiltershingle.md)
 - [TokenFiltersnowballStemming](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFiltersnowballStemming.md)
 - [TokenFilterstopword](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFilterstopword.md)
 - [TokenFiltertrim](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenFiltertrim.md)
 - [TokenizeredgeGram](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenizeredgeGram.md)
 - [Tokenizerkeyword](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Tokenizerkeyword.md)
 - [TokenizernGram](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenizernGram.md)
 - [TokenizerregexCaptureGroup](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenizerregexCaptureGroup.md)
 - [TokenizerregexSplit](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenizerregexSplit.md)
 - [Tokenizerstandard](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Tokenizerstandard.md)
 - [TokenizeruaxUrlEmail](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TokenizeruaxUrlEmail.md)
 - [Tokenizerwhitespace](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Tokenizerwhitespace.md)
 - [TriggerIngestionRequest](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/TriggerIngestionRequest.md)
 - [UpdateCustomDBRole](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UpdateCustomDBRole.md)
 - [UserAccessList](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UserAccessList.md)
 - [UserCert](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UserCert.md)
 - [UserEventTypeViewForOrg](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UserEventTypeViewForOrg.md)
 - [UserNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UserNotification.md)
 - [UserRoleAssignment](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UserRoleAssignment.md)
 - [UserScope](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UserScope.md)
 - [UserSecurity](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/UserSecurity.md)
 - [Validation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Validation.md)
 - [VictorOps](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/VictorOps.md)
 - [VictorOpsNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/VictorOpsNotification.md)
 - [Webhook](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/Webhook.md)
 - [WebhookNotification](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/WebhookNotification.md)
 - [WeeklySchedule](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/WeeklySchedule.md)
 - [X509Certificate](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/docs/X509Certificate.md)




