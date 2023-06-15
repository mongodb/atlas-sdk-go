# SDK Reference 

## Reference Documentation for SDK Endpoints

All URIs are relative to *https://cloud.mongodb.com*

Class        | Method        | Description   | Stability level   
------------ | ------------- | ------------- | ------------- | -------------
*AWSClustersDNSApi* | [GetAWSCustomDNS](./docs/AWSClustersDNSApi.md#getawscustomdns) |  Return One Custom DNS Configuration for Atlas Clusters on AWS | Experimental
*AWSClustersDNSApi* | [ToggleAWSCustomDNS](./docs/AWSClustersDNSApi.md#toggleawscustomdns) |  Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS | Experimental
*AccessTrackingApi* | [ListAccessLogsByClusterName](./docs/AccessTrackingApi.md#listaccesslogsbyclustername) |  Return Database Access History for One Cluster using Its Cluster Name | Stable
*AccessTrackingApi* | [ListAccessLogsByHostname](./docs/AccessTrackingApi.md#listaccesslogsbyhostname) |  Return Database Access History for One Cluster using Its Hostname | Stable
*AlertConfigurationsApi* | [CreateAlertConfiguration](./docs/AlertConfigurationsApi.md#createalertconfiguration) |  Create One Alert Configuration in One Project | Stable
*AlertConfigurationsApi* | [DeleteAlertConfiguration](./docs/AlertConfigurationsApi.md#deletealertconfiguration) |  Remove One Alert Configuration from One Project | Stable
*AlertConfigurationsApi* | [GetAlertConfiguration](./docs/AlertConfigurationsApi.md#getalertconfiguration) |  Return One Alert Configuration from One Project | Experimental
*AlertConfigurationsApi* | [ListAlertConfigurationMatchersFieldNames](./docs/AlertConfigurationsApi.md#listalertconfigurationmatchersfieldnames) |  Get All Alert Configuration Matchers Field Names | Stable
*AlertConfigurationsApi* | [ListAlertConfigurations](./docs/AlertConfigurationsApi.md#listalertconfigurations) |  Return All Alert Configurations for One Project | Stable
*AlertConfigurationsApi* | [ListAlertConfigurationsByAlertId](./docs/AlertConfigurationsApi.md#listalertconfigurationsbyalertid) |  Return All Alert Configurations Set for One Alert | Experimental
*AlertConfigurationsApi* | [ToggleAlertConfiguration](./docs/AlertConfigurationsApi.md#togglealertconfiguration) |  Toggle One State of One Alert Configuration in One Project | Stable
*AlertConfigurationsApi* | [UpdateAlertConfiguration](./docs/AlertConfigurationsApi.md#updatealertconfiguration) |  Update One Alert Configuration for One Project | Stable
*AlertsApi* | [AcknowledgeAlert](./docs/AlertsApi.md#acknowledgealert) |  Acknowledge One Alert from One Project | Stable
*AlertsApi* | [GetAlert](./docs/AlertsApi.md#getalert) |  Return One Alert from One Project | Stable
*AlertsApi* | [ListAlerts](./docs/AlertsApi.md#listalerts) |  Return All Alerts from One Project | Stable
*AlertsApi* | [ListAlertsByAlertConfigurationId](./docs/AlertsApi.md#listalertsbyalertconfigurationid) |  Return All Open Alerts for Alert Configuration | Experimental
*AtlasSearchApi* | [CreateAtlasSearchIndex](./docs/AtlasSearchApi.md#createatlassearchindex) |  Create One Atlas Search Index | Stable
*AtlasSearchApi* | [DeleteAtlasSearchIndex](./docs/AtlasSearchApi.md#deleteatlassearchindex) |  Remove One Atlas Search Index | Stable
*AtlasSearchApi* | [GetAtlasSearchIndex](./docs/AtlasSearchApi.md#getatlassearchindex) |  Return One Atlas Search Index | Stable
*AtlasSearchApi* | [ListAtlasSearchIndexes](./docs/AtlasSearchApi.md#listatlassearchindexes) |  Return All Atlas Search Indexes for One Collection | Stable
*AtlasSearchApi* | [UpdateAtlasSearchIndex](./docs/AtlasSearchApi.md#updateatlassearchindex) |  Update One Atlas Search Index | Stable
*AuditingApi* | [GetAuditingConfiguration](./docs/AuditingApi.md#getauditingconfiguration) |  Return the Auditing Configuration for One Project | Stable
*AuditingApi* | [UpdateAuditingConfiguration](./docs/AuditingApi.md#updateauditingconfiguration) |  Update Auditing Configuration for One Project | Experimental
*CloudBackupsApi* | [CancelBackupRestoreJob](./docs/CloudBackupsApi.md#cancelbackuprestorejob) |  Cancel One Restore Job of One Cluster | Experimental
*CloudBackupsApi* | [CreateBackupExportJob](./docs/CloudBackupsApi.md#createbackupexportjob) |  Create One Cloud Backup Snapshot Export Job | Stable
*CloudBackupsApi* | [CreateBackupRestoreJob](./docs/CloudBackupsApi.md#createbackuprestorejob) |  Restore One Snapshot of One Cluster | Stable
*CloudBackupsApi* | [CreateExportBucket](./docs/CloudBackupsApi.md#createexportbucket) |  Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports | Stable
*CloudBackupsApi* | [CreateServerlessBackupRestoreJob](./docs/CloudBackupsApi.md#createserverlessbackuprestorejob) |  Restore One Snapshot of One Serverless Instance | Stable
*CloudBackupsApi* | [DeleteAllBackupSchedules](./docs/CloudBackupsApi.md#deleteallbackupschedules) |  Remove All Cloud Backup Schedules | Stable
*CloudBackupsApi* | [DeleteExportBucket](./docs/CloudBackupsApi.md#deleteexportbucket) |  Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports | Stable
*CloudBackupsApi* | [DeleteReplicaSetBackup](./docs/CloudBackupsApi.md#deletereplicasetbackup) |  Remove One Replica Set Cloud Backup | Stable
*CloudBackupsApi* | [DeleteShardedClusterBackup](./docs/CloudBackupsApi.md#deleteshardedclusterbackup) |  Remove One Sharded Cluster Cloud Backup | Experimental
*CloudBackupsApi* | [GetBackupExportJob](./docs/CloudBackupsApi.md#getbackupexportjob) |  Return One Cloud Backup Snapshot Export Job | Stable
*CloudBackupsApi* | [GetBackupRestoreJob](./docs/CloudBackupsApi.md#getbackuprestorejob) |  Return One Restore Job of One Cluster | Stable
*CloudBackupsApi* | [GetBackupSchedule](./docs/CloudBackupsApi.md#getbackupschedule) |  Return One Cloud Backup Schedule | Stable
*CloudBackupsApi* | [GetDataProtectionSettings](./docs/CloudBackupsApi.md#getdataprotectionsettings) |  Return the Backup Compliance Policy settings | Experimental
*CloudBackupsApi* | [GetExportBucket](./docs/CloudBackupsApi.md#getexportbucket) |  Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports | Stable
*CloudBackupsApi* | [GetReplicaSetBackup](./docs/CloudBackupsApi.md#getreplicasetbackup) |  Return One Replica Set Cloud Backup | Stable
*CloudBackupsApi* | [GetServerlessBackup](./docs/CloudBackupsApi.md#getserverlessbackup) |  Return One Snapshot of One Serverless Instance | Stable
*CloudBackupsApi* | [GetServerlessBackupRestoreJob](./docs/CloudBackupsApi.md#getserverlessbackuprestorejob) |  Return One Restore Job for One Serverless Instance | Stable
*CloudBackupsApi* | [GetShardedClusterBackup](./docs/CloudBackupsApi.md#getshardedclusterbackup) |  Return One Sharded Cluster Cloud Backup | Experimental
*CloudBackupsApi* | [ListBackupExportJobs](./docs/CloudBackupsApi.md#listbackupexportjobs) |  Return All Cloud Backup Snapshot Export Jobs | Stable
*CloudBackupsApi* | [ListBackupRestoreJobs](./docs/CloudBackupsApi.md#listbackuprestorejobs) |  Return All Restore Jobs for One Cluster | Stable
*CloudBackupsApi* | [ListExportBuckets](./docs/CloudBackupsApi.md#listexportbuckets) |  Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports | Stable
*CloudBackupsApi* | [ListReplicaSetBackups](./docs/CloudBackupsApi.md#listreplicasetbackups) |  Return All Replica Set Cloud Backups | Stable
*CloudBackupsApi* | [ListServerlessBackupRestoreJobs](./docs/CloudBackupsApi.md#listserverlessbackuprestorejobs) |  Return All Restore Jobs for One Serverless Instance | Stable
*CloudBackupsApi* | [ListServerlessBackups](./docs/CloudBackupsApi.md#listserverlessbackups) |  Return All Snapshots of One Serverless Instance | Stable
*CloudBackupsApi* | [ListShardedClusterBackups](./docs/CloudBackupsApi.md#listshardedclusterbackups) |  Return All Sharded Cluster Cloud Backups | Experimental
*CloudBackupsApi* | [TakeSnapshot](./docs/CloudBackupsApi.md#takesnapshot) |  Take One On-Demand Snapshot | Stable
*CloudBackupsApi* | [UpdateBackupSchedule](./docs/CloudBackupsApi.md#updatebackupschedule) |  Update Cloud Backup Schedule for One Cluster | Stable
*CloudBackupsApi* | [UpdateDataProtectionSettings](./docs/CloudBackupsApi.md#updatedataprotectionsettings) |  Update or enable the Backup Compliance Policy settings | Experimental
*CloudBackupsApi* | [UpdateSnapshotRetention](./docs/CloudBackupsApi.md#updatesnapshotretention) |  Change Expiration Date for One Cloud Backup | Experimental
*CloudMigrationServiceApi* | [CreateLinkToken](./docs/CloudMigrationServiceApi.md#createlinktoken) |  Create One Link-Token | Stable
*CloudMigrationServiceApi* | [CreatePushMigration](./docs/CloudMigrationServiceApi.md#createpushmigration) |  Migrate One Local Managed Cluster to MongoDB Atlas | Stable
*CloudMigrationServiceApi* | [CutoverMigration](./docs/CloudMigrationServiceApi.md#cutovermigration) |  Cut Over the Migrated Cluster | Stable
*CloudMigrationServiceApi* | [DeleteLinkToken](./docs/CloudMigrationServiceApi.md#deletelinktoken) |  Remove One Link-Token | Stable
*CloudMigrationServiceApi* | [GetPushMigration](./docs/CloudMigrationServiceApi.md#getpushmigration) |  Return One Migration Job | Stable
*CloudMigrationServiceApi* | [GetValidationStatus](./docs/CloudMigrationServiceApi.md#getvalidationstatus) |  Return One Migration Validation Job | Stable
*CloudMigrationServiceApi* | [ListSourceProjects](./docs/CloudMigrationServiceApi.md#listsourceprojects) |  Return All Projects Available for Migration | Experimental
*CloudMigrationServiceApi* | [ValidateMigration](./docs/CloudMigrationServiceApi.md#validatemigration) |  Validate One Migration Request | Stable
*CloudProviderAccessApi* | [AuthorizeCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#authorizecloudprovideraccessrole) |  Authorize One Cloud Provider Access Role | Stable
*CloudProviderAccessApi* | [CreateCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#createcloudprovideraccessrole) |  Create One Cloud Provider Access Role | Stable
*CloudProviderAccessApi* | [DeauthorizeCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#deauthorizecloudprovideraccessrole) |  Deauthorize One Cloud Provider Access Role | Stable
*CloudProviderAccessApi* | [GetCloudProviderAccessRole](./docs/CloudProviderAccessApi.md#getcloudprovideraccessrole) |  Return specified Cloud Provider Access Role | Experimental
*CloudProviderAccessApi* | [ListCloudProviderAccessRoles](./docs/CloudProviderAccessApi.md#listcloudprovideraccessroles) |  Return All Cloud Provider Access Roles | Stable
*ClusterOutageSimulationApi* | [EndOutageSimulation](./docs/ClusterOutageSimulationApi.md#endoutagesimulation) |  End an Outage Simulation | Experimental
*ClusterOutageSimulationApi* | [GetOutageSimulation](./docs/ClusterOutageSimulationApi.md#getoutagesimulation) |  Return One Outage Simulation | Experimental
*ClusterOutageSimulationApi* | [StartOutageSimulation](./docs/ClusterOutageSimulationApi.md#startoutagesimulation) |  Start an Outage Simulation | Experimental
*ClustersApi* | [GetClusterAdvancedConfiguration](./docs/ClustersApi.md#getclusteradvancedconfiguration) |  Return One Advanced Configuration Options for One Cluster | Stable
*ClustersApi* | [GetClusterStatus](./docs/ClustersApi.md#getclusterstatus) |  Return Status of All Cluster Operations | Experimental
*ClustersApi* | [GetSampleDatasetLoadStatus](./docs/ClustersApi.md#getsampledatasetloadstatus) |  Check Status of Cluster Sample Dataset Request | Stable
*ClustersApi* | [ListCloudProviderRegions](./docs/ClustersApi.md#listcloudproviderregions) |  Return All Cloud Provider Regions | Stable
*ClustersApi* | [ListClustersForAllProjects](./docs/ClustersApi.md#listclustersforallprojects) |  Return All Authorized Clusters in All Projects | Experimental
*ClustersApi* | [LoadSampleDataset](./docs/ClustersApi.md#loadsampledataset) |  Load Sample Dataset Request into Cluster | Stable
*ClustersApi* | [UpdateClusterAdvancedConfiguration](./docs/ClustersApi.md#updateclusteradvancedconfiguration) |  Update Advanced Configuration Options for One Cluster | Stable
*ClustersApi* | [UpgradeSharedCluster](./docs/ClustersApi.md#upgradesharedcluster) |  Upgrade One Shared-tier Cluster | Experimental
*ClustersApi* | [UpgradeSharedClusterToServerless](./docs/ClustersApi.md#upgradesharedclustertoserverless) |  Upgrades One Shared-Tier Cluster to the Serverless Instance | Experimental
*CustomDatabaseRolesApi* | [CreateCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#createcustomdatabaserole) |  Create One Custom Role | Stable
*CustomDatabaseRolesApi* | [DeleteCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#deletecustomdatabaserole) |  Remove One Custom Role from One Project | Stable
*CustomDatabaseRolesApi* | [GetCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#getcustomdatabaserole) |  Return One Custom Role in One Project | Stable
*CustomDatabaseRolesApi* | [ListCustomDatabaseRoles](./docs/CustomDatabaseRolesApi.md#listcustomdatabaseroles) |  Return All Custom Roles in One Project | Stable
*CustomDatabaseRolesApi* | [UpdateCustomDatabaseRole](./docs/CustomDatabaseRolesApi.md#updatecustomdatabaserole) |  Update One Custom Role in One Project | Stable
*DataFederationApi* | [CreateDataFederationPrivateEndpoint](./docs/DataFederationApi.md#createdatafederationprivateendpoint) |  Create One Federated Database Instance and Online Archive Private Endpoint for One Project | Stable
*DataFederationApi* | [CreateFederatedDatabase](./docs/DataFederationApi.md#createfederateddatabase) |  Create One Federated Database Instance in One Project | Experimental
*DataFederationApi* | [CreateOneDataFederationQueryLimit](./docs/DataFederationApi.md#createonedatafederationquerylimit) |  Configure One Query Limit for One Federated Database Instance | Experimental
*DataFederationApi* | [DeleteDataFederationPrivateEndpoint](./docs/DataFederationApi.md#deletedatafederationprivateendpoint) |  Remove One Federated Database Instance and Online Archive Private Endpoint from One Project | Stable
*DataFederationApi* | [DeleteFederatedDatabase](./docs/DataFederationApi.md#deletefederateddatabase) |  Remove One Federated Database Instance from One Project | Experimental
*DataFederationApi* | [DeleteOneDataFederationInstanceQueryLimit](./docs/DataFederationApi.md#deleteonedatafederationinstancequerylimit) |  Delete One Query Limit For One Federated Database Instance | Experimental
*DataFederationApi* | [DownloadFederatedDatabaseQueryLogs](./docs/DataFederationApi.md#downloadfederateddatabasequerylogs) |  Download Query Logs for One Federated Database Instance | Experimental
*DataFederationApi* | [GetDataFederationPrivateEndpoint](./docs/DataFederationApi.md#getdatafederationprivateendpoint) |  Return One Federated Database Instance and Online Archive Private Endpoint in One Project | Stable
*DataFederationApi* | [GetFederatedDatabase](./docs/DataFederationApi.md#getfederateddatabase) |  Return One Federated Database Instance in One Project | Experimental
*DataFederationApi* | [ListDataFederationPrivateEndpoints](./docs/DataFederationApi.md#listdatafederationprivateendpoints) |  Return All Federated Database Instance and Online Archive Private Endpoints in One Project | Stable
*DataFederationApi* | [ListFederatedDatabases](./docs/DataFederationApi.md#listfederateddatabases) |  Return All Federated Database Instances in One Project | Experimental
*DataFederationApi* | [ReturnFederatedDatabaseQueryLimit](./docs/DataFederationApi.md#returnfederateddatabasequerylimit) |  Return One Federated Database Instance Query Limit for One Project | Experimental
*DataFederationApi* | [ReturnFederatedDatabaseQueryLimits](./docs/DataFederationApi.md#returnfederateddatabasequerylimits) |  Return All Query Limits for One Federated Database Instance | Experimental
*DataFederationApi* | [UpdateFederatedDatabase](./docs/DataFederationApi.md#updatefederateddatabase) |  Update One Federated Database Instance in One Project | Experimental
*DataLakePipelinesApi* | [CreatePipeline](./docs/DataLakePipelinesApi.md#createpipeline) |  Create One Data Lake Pipeline | Stable
*DataLakePipelinesApi* | [DeletePipeline](./docs/DataLakePipelinesApi.md#deletepipeline) |  Remove One Data Lake Pipeline | Stable
*DataLakePipelinesApi* | [DeletePipelineRunDataset](./docs/DataLakePipelinesApi.md#deletepipelinerundataset) |  Delete Pipeline Run Dataset | Stable
*DataLakePipelinesApi* | [GetPipeline](./docs/DataLakePipelinesApi.md#getpipeline) |  Return One Data Lake Pipeline | Stable
*DataLakePipelinesApi* | [GetPipelineRun](./docs/DataLakePipelinesApi.md#getpipelinerun) |  Return One Data Lake Pipeline Run | Stable
*DataLakePipelinesApi* | [ListPipelineRuns](./docs/DataLakePipelinesApi.md#listpipelineruns) |  Return All Data Lake Pipeline Runs from One Project | Stable
*DataLakePipelinesApi* | [ListPipelineSchedules](./docs/DataLakePipelinesApi.md#listpipelineschedules) |  Return Available Ingestion Schedules for One Data Lake Pipeline | Stable
*DataLakePipelinesApi* | [ListPipelineSnapshots](./docs/DataLakePipelinesApi.md#listpipelinesnapshots) |  Return Available Backup Snapshots for One Data Lake Pipeline | Stable
*DataLakePipelinesApi* | [ListPipelines](./docs/DataLakePipelinesApi.md#listpipelines) |  Return All Data Lake Pipelines from One Project | Stable
*DataLakePipelinesApi* | [PausePipeline](./docs/DataLakePipelinesApi.md#pausepipeline) |  Pause One Data Lake Pipeline | Stable
*DataLakePipelinesApi* | [ResumePipeline](./docs/DataLakePipelinesApi.md#resumepipeline) |  Resume One Data Lake Pipeline | Stable
*DataLakePipelinesApi* | [TriggerSnapshotIngestion](./docs/DataLakePipelinesApi.md#triggersnapshotingestion) |  Trigger on demand snapshot ingestion | Stable
*DataLakePipelinesApi* | [UpdatePipeline](./docs/DataLakePipelinesApi.md#updatepipeline) |  Update One Data Lake Pipeline | Stable
*DatabaseUsersApi* | [CreateDatabaseUser](./docs/DatabaseUsersApi.md#createdatabaseuser) |  Create One Database User in One Project | Stable
*DatabaseUsersApi* | [DeleteDatabaseUser](./docs/DatabaseUsersApi.md#deletedatabaseuser) |  Remove One Database User from One Project | Stable
*DatabaseUsersApi* | [GetDatabaseUser](./docs/DatabaseUsersApi.md#getdatabaseuser) |  Return One Database User from One Project | Stable
*DatabaseUsersApi* | [ListDatabaseUsers](./docs/DatabaseUsersApi.md#listdatabaseusers) |  Return All Database Users from One Project | Stable
*DatabaseUsersApi* | [UpdateDatabaseUser](./docs/DatabaseUsersApi.md#updatedatabaseuser) |  Update One Database User in One Project | Stable
*EncryptionAtRestUsingCustomerKeyManagementApi* | [GetEncryptionAtRest](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#getencryptionatrest) |  Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project | Stable
*EncryptionAtRestUsingCustomerKeyManagementApi* | [UpdateEncryptionAtRest](./docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#updateencryptionatrest) |  Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project | Experimental
*EventsApi* | [GetOrganizationEvent](./docs/EventsApi.md#getorganizationevent) |  Return One Event from One Organization | Experimental
*EventsApi* | [GetProjectEvent](./docs/EventsApi.md#getprojectevent) |  Return One Event from One Project | Experimental
*EventsApi* | [ListOrganizationEvents](./docs/EventsApi.md#listorganizationevents) |  Return All Events from One Organization | Stable
*EventsApi* | [ListProjectEvents](./docs/EventsApi.md#listprojectevents) |  Return All Events from One Project | Stable
*FederatedAuthenticationApi* | [CreateRoleMapping](./docs/FederatedAuthenticationApi.md#createrolemapping) |  Add One Role Mapping to One Organization | Experimental
*FederatedAuthenticationApi* | [DeleteFederationApp](./docs/FederatedAuthenticationApi.md#deletefederationapp) |  Delete the federation settings instance. | Experimental
*FederatedAuthenticationApi* | [DeleteRoleMapping](./docs/FederatedAuthenticationApi.md#deleterolemapping) |  Remove One Role Mapping from One Organization | Experimental
*FederatedAuthenticationApi* | [GetConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#getconnectedorgconfig) |  Return One Org Config Connected to One Federation | Experimental
*FederatedAuthenticationApi* | [GetFederationSettings](./docs/FederatedAuthenticationApi.md#getfederationsettings) |  Return Federation Settings for One Organization | Experimental
*FederatedAuthenticationApi* | [GetIdentityProvider](./docs/FederatedAuthenticationApi.md#getidentityprovider) |  Return one identity provider from the specified federation. | Experimental
*FederatedAuthenticationApi* | [GetIdentityProviderMetadata](./docs/FederatedAuthenticationApi.md#getidentityprovidermetadata) |  Return the metadata of one identity provider in the specified federation. | Experimental
*FederatedAuthenticationApi* | [GetRoleMapping](./docs/FederatedAuthenticationApi.md#getrolemapping) |  Return One Role Mapping from One Organization | Experimental
*FederatedAuthenticationApi* | [ListConnectedOrgConfigs](./docs/FederatedAuthenticationApi.md#listconnectedorgconfigs) |  Return All Connected Org Configs from the Federation | Experimental
*FederatedAuthenticationApi* | [ListIdentityProviders](./docs/FederatedAuthenticationApi.md#listidentityproviders) |  Return all identity providers from the specified federation. | Experimental
*FederatedAuthenticationApi* | [ListRoleMappings](./docs/FederatedAuthenticationApi.md#listrolemappings) |  Return All Role Mappings from One Organization | Experimental
*FederatedAuthenticationApi* | [RemoveConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#removeconnectedorgconfig) |  Remove One Org Config Connected to One Federation | Experimental
*FederatedAuthenticationApi* | [UpdateConnectedOrgConfig](./docs/FederatedAuthenticationApi.md#updateconnectedorgconfig) |  Update One Org Config Connected to One Federation | Experimental
*FederatedAuthenticationApi* | [UpdateIdentityProvider](./docs/FederatedAuthenticationApi.md#updateidentityprovider) |  Update the identity provider. | Experimental
*FederatedAuthenticationApi* | [UpdateRoleMapping](./docs/FederatedAuthenticationApi.md#updaterolemapping) |  Update One Role Mapping in One Organization | Experimental
*GlobalClustersApi* | [CreateCustomZoneMapping](./docs/GlobalClustersApi.md#createcustomzonemapping) |  Add One Entry to One Custom Zone Mapping | Experimental
*GlobalClustersApi* | [CreateManagedNamespace](./docs/GlobalClustersApi.md#createmanagednamespace) |  Create One Managed Namespace in One Global Multi-Cloud Cluster | Experimental
*GlobalClustersApi* | [DeleteAllCustomZoneMappings](./docs/GlobalClustersApi.md#deleteallcustomzonemappings) |  Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster | Experimental
*GlobalClustersApi* | [DeleteManagedNamespace](./docs/GlobalClustersApi.md#deletemanagednamespace) |  Remove One Managed Namespace from One Global Multi-Cloud Cluster | Experimental
*GlobalClustersApi* | [GetManagedNamespace](./docs/GlobalClustersApi.md#getmanagednamespace) |  Return One Managed Namespace in One Global Multi-Cloud Cluster | Stable
*InvoicesApi* | [DownloadInvoiceCSV](./docs/InvoicesApi.md#downloadinvoicecsv) |  Return One Organization Invoice as CSV | Experimental
*InvoicesApi* | [GetInvoice](./docs/InvoicesApi.md#getinvoice) |  Return One Organization Invoice | Experimental
*InvoicesApi* | [ListInvoices](./docs/InvoicesApi.md#listinvoices) |  Return All Invoices for One Organization | Experimental
*InvoicesApi* | [ListPendingInvoices](./docs/InvoicesApi.md#listpendinginvoices) |  Return All Pending Invoices for One Organization | Experimental
*LDAPConfigurationApi* | [DeleteLDAPConfiguration](./docs/LDAPConfigurationApi.md#deleteldapconfiguration) |  Remove the Current LDAP User to DN Mapping | Experimental
*LDAPConfigurationApi* | [GetLDAPConfiguration](./docs/LDAPConfigurationApi.md#getldapconfiguration) |  Return the Current LDAP or X.509 Configuration | Experimental
*LDAPConfigurationApi* | [GetLDAPConfigurationStatus](./docs/LDAPConfigurationApi.md#getldapconfigurationstatus) |  Return the Status of One Verify LDAP Configuration Request | Experimental
*LDAPConfigurationApi* | [SaveLDAPConfiguration](./docs/LDAPConfigurationApi.md#saveldapconfiguration) |  Edit the LDAP or X.509 Configuration | Experimental
*LDAPConfigurationApi* | [VerifyLDAPConfiguration](./docs/LDAPConfigurationApi.md#verifyldapconfiguration) |  Verify the LDAP Configuration in One Project | Experimental
*LegacyBackupApi* | [DeleteLegacySnapshot](./docs/LegacyBackupApi.md#deletelegacysnapshot) |  Remove One Legacy Backup Snapshot | Experimental
*LegacyBackupApi* | [GetLegacyBackupCheckpoint](./docs/LegacyBackupApi.md#getlegacybackupcheckpoint) |  Return One Legacy Backup Checkpoint | Experimental
*LegacyBackupApi* | [GetLegacyBackupRestoreJob](./docs/LegacyBackupApi.md#getlegacybackuprestorejob) |  Return One Legacy Backup Restore Job | Experimental
*LegacyBackupApi* | [GetLegacySnapshot](./docs/LegacyBackupApi.md#getlegacysnapshot) |  Return One Legacy Backup Snapshot | Experimental
*LegacyBackupApi* | [GetLegacySnapshotSchedule](./docs/LegacyBackupApi.md#getlegacysnapshotschedule) |  Return One Snapshot Schedule | Experimental
*LegacyBackupApi* | [ListLegacyBackupCheckpoints](./docs/LegacyBackupApi.md#listlegacybackupcheckpoints) |  Return All Legacy Backup Checkpoints | Experimental
*LegacyBackupApi* | [ListLegacyBackupRestoreJobs](./docs/LegacyBackupApi.md#listlegacybackuprestorejobs) |  Return All Legacy Backup Restore Jobs | Experimental
*LegacyBackupApi* | [ListLegacySnapshots](./docs/LegacyBackupApi.md#listlegacysnapshots) |  Return All Legacy Backup Snapshots | Experimental
*LegacyBackupApi* | [UpdateLegacySnapshotRetention](./docs/LegacyBackupApi.md#updatelegacysnapshotretention) |  Change One Legacy Backup Snapshot Expiration | Experimental
*LegacyBackupApi* | [UpdateLegacySnapshotSchedule](./docs/LegacyBackupApi.md#updatelegacysnapshotschedule) |  Update Snapshot Schedule for One Cluster | Experimental
*LegacyBackupRestoreJobsApi* | [CreateLegacyBackupRestoreJob](./docs/LegacyBackupRestoreJobsApi.md#createlegacybackuprestorejob) |  Create One Legacy Backup Restore Job | Experimental
*MaintenanceWindowsApi* | [DeferMaintenanceWindow](./docs/MaintenanceWindowsApi.md#defermaintenancewindow) |  Defer One Maintenance Window for One Project | Stable
*MaintenanceWindowsApi* | [GetMaintenanceWindow](./docs/MaintenanceWindowsApi.md#getmaintenancewindow) |  Return One Maintenance Window for One Project | Stable
*MaintenanceWindowsApi* | [ResetMaintenanceWindow](./docs/MaintenanceWindowsApi.md#resetmaintenancewindow) |  Reset One Maintenance Window for One Project | Stable
*MaintenanceWindowsApi* | [ToggleMaintenanceAutoDefer](./docs/MaintenanceWindowsApi.md#togglemaintenanceautodefer) |  Toggle Automatic Deferral of Maintenance for One Project | Experimental
*MaintenanceWindowsApi* | [UpdateMaintenanceWindow](./docs/MaintenanceWindowsApi.md#updatemaintenancewindow) |  Update Maintenance Window for One Project | Stable
*MongoDBCloudUsersApi* | [CreateUser](./docs/MongoDBCloudUsersApi.md#createuser) |  Create One MongoDB Cloud User | Stable
*MongoDBCloudUsersApi* | [GetUser](./docs/MongoDBCloudUsersApi.md#getuser) |  Return One MongoDB Cloud User using Its ID | Stable
*MongoDBCloudUsersApi* | [GetUserByUsername](./docs/MongoDBCloudUsersApi.md#getuserbyusername) |  Return One MongoDB Cloud User using Their Username | Stable
*MonitoringAndLogsApi* | [GetAtlasProcess](./docs/MonitoringAndLogsApi.md#getatlasprocess) |  Return One MongoDB Process by ID | Stable
*MonitoringAndLogsApi* | [GetDatabase](./docs/MonitoringAndLogsApi.md#getdatabase) |  Return One Database for a MongoDB Process | Experimental
*MonitoringAndLogsApi* | [GetDatabaseMeasurements](./docs/MonitoringAndLogsApi.md#getdatabasemeasurements) |  Return Measurements of One Database for One MongoDB Process | Stable
*MonitoringAndLogsApi* | [GetDiskMeasurements](./docs/MonitoringAndLogsApi.md#getdiskmeasurements) |  Return Measurements of One Disk for One MongoDB Process | Stable
*MonitoringAndLogsApi* | [GetHostLogs](./docs/MonitoringAndLogsApi.md#gethostlogs) |  Download Logs for One Cluster Host in One Project | Stable
*MonitoringAndLogsApi* | [GetHostMeasurements](./docs/MonitoringAndLogsApi.md#gethostmeasurements) |  Return Measurements for One MongoDB Process | Stable
*MonitoringAndLogsApi* | [GetIndexMetrics](./docs/MonitoringAndLogsApi.md#getindexmetrics) |  Return Atlas Search Metrics for One Index in One Specified Namespace | Experimental
*MonitoringAndLogsApi* | [GetMeasurements](./docs/MonitoringAndLogsApi.md#getmeasurements) |  Return Atlas Search Hardware and Status Metrics | Experimental
*MonitoringAndLogsApi* | [ListAtlasProcesses](./docs/MonitoringAndLogsApi.md#listatlasprocesses) |  Return All MongoDB Processes in One Project | Stable
*MonitoringAndLogsApi* | [ListDatabases](./docs/MonitoringAndLogsApi.md#listdatabases) |  Return Available Databases for One MongoDB Process | Stable
*MonitoringAndLogsApi* | [ListDiskMeasurements](./docs/MonitoringAndLogsApi.md#listdiskmeasurements) |  Return Measurements of One Disk | Experimental
*MonitoringAndLogsApi* | [ListDiskPartitions](./docs/MonitoringAndLogsApi.md#listdiskpartitions) |  Return Available Disks for One MongoDB Process | Stable
*MonitoringAndLogsApi* | [ListIndexMetrics](./docs/MonitoringAndLogsApi.md#listindexmetrics) |  Return All Atlas Search Index Metrics for One Namespace | Experimental
*MonitoringAndLogsApi* | [ListMetricTypes](./docs/MonitoringAndLogsApi.md#listmetrictypes) |  Return All Atlas Search Metric Types for One Process | Experimental
*MultiCloudClustersApi* | [CreateCluster](./docs/MultiCloudClustersApi.md#createcluster) |  Create One Multi-Cloud Cluster from One Project | Stable
*MultiCloudClustersApi* | [DeleteCluster](./docs/MultiCloudClustersApi.md#deletecluster) |  Remove One Multi-Cloud Cluster from One Project | Stable
*MultiCloudClustersApi* | [GetCluster](./docs/MultiCloudClustersApi.md#getcluster) |  Return One Multi-Cloud Cluster from One Project | Stable
*MultiCloudClustersApi* | [ListClusters](./docs/MultiCloudClustersApi.md#listclusters) |  Return All Clusters in One Project | Stable
*MultiCloudClustersApi* | [TestFailover](./docs/MultiCloudClustersApi.md#testfailover) |  Test Failover for One Multi-Cloud Cluster | Stable
*MultiCloudClustersApi* | [UpdateCluster](./docs/MultiCloudClustersApi.md#updatecluster) |  Modify One Multi-Cloud Cluster from One Project | Stable
*NetworkPeeringApi* | [CreatePeeringConnection](./docs/NetworkPeeringApi.md#createpeeringconnection) |  Create One New Network Peering Connection | Stable
*NetworkPeeringApi* | [CreatePeeringContainer](./docs/NetworkPeeringApi.md#createpeeringcontainer) |  Create One New Network Peering Container | Stable
*NetworkPeeringApi* | [DeletePeeringConnection](./docs/NetworkPeeringApi.md#deletepeeringconnection) |  Remove One Existing Network Peering Connection | Stable
*NetworkPeeringApi* | [DeletePeeringContainer](./docs/NetworkPeeringApi.md#deletepeeringcontainer) |  Remove One Network Peering Container | Stable
*NetworkPeeringApi* | [DisablePeering](./docs/NetworkPeeringApi.md#disablepeering) |  Disable Connect via Peering Only Mode for One Project | Experimental
*NetworkPeeringApi* | [GetPeeringConnection](./docs/NetworkPeeringApi.md#getpeeringconnection) |  Return One Network Peering Connection in One Project | Stable
*NetworkPeeringApi* | [GetPeeringContainer](./docs/NetworkPeeringApi.md#getpeeringcontainer) |  Return One Network Peering Container | Stable
*NetworkPeeringApi* | [ListPeeringConnections](./docs/NetworkPeeringApi.md#listpeeringconnections) |  Return All Network Peering Connections in One Project | Stable
*NetworkPeeringApi* | [ListPeeringContainerByCloudProvider](./docs/NetworkPeeringApi.md#listpeeringcontainerbycloudprovider) |  Return All Network Peering Containers in One Project for One Cloud Provider | Stable
*NetworkPeeringApi* | [ListPeeringContainers](./docs/NetworkPeeringApi.md#listpeeringcontainers) |  Return All Network Peering Containers in One Project | Stable
*NetworkPeeringApi* | [UpdatePeeringConnection](./docs/NetworkPeeringApi.md#updatepeeringconnection) |  Update One New Network Peering Connection | Experimental
*NetworkPeeringApi* | [UpdatePeeringContainer](./docs/NetworkPeeringApi.md#updatepeeringcontainer) |  Update One Network Peering Container | Experimental
*NetworkPeeringApi* | [VerifyConnectViaPeeringOnlyModeForOneProject](./docs/NetworkPeeringApi.md#verifyconnectviapeeringonlymodeforoneproject) |  Verify Connect via Peering Only Mode for One Project | Experimental
*OnlineArchiveApi* | [CreateOnlineArchive](./docs/OnlineArchiveApi.md#createonlinearchive) |  Create One Online Archive | Stable
*OnlineArchiveApi* | [DeleteOnlineArchive](./docs/OnlineArchiveApi.md#deleteonlinearchive) |  Remove One Online Archive | Stable
*OnlineArchiveApi* | [DownloadOnlineArchiveQueryLogs](./docs/OnlineArchiveApi.md#downloadonlinearchivequerylogs) |  Download Online Archive Query Logs | Experimental
*OnlineArchiveApi* | [GetOnlineArchive](./docs/OnlineArchiveApi.md#getonlinearchive) |  Return One Online Archive | Stable
*OnlineArchiveApi* | [ListOnlineArchives](./docs/OnlineArchiveApi.md#listonlinearchives) |  Return All Online Archives for One Cluster | Stable
*OnlineArchiveApi* | [UpdateOnlineArchive](./docs/OnlineArchiveApi.md#updateonlinearchive) |  Update One Online Archive | Stable
*OrganizationsApi* | [CreateOrganization](./docs/OrganizationsApi.md#createorganization) |  Create One Organization | Stable
*OrganizationsApi* | [CreateOrganizationInvitation](./docs/OrganizationsApi.md#createorganizationinvitation) |  Invite One MongoDB Cloud User to Join One Atlas Organization | Stable
*OrganizationsApi* | [DeleteOrganization](./docs/OrganizationsApi.md#deleteorganization) |  Remove One Organization | Stable
*OrganizationsApi* | [DeleteOrganizationInvitation](./docs/OrganizationsApi.md#deleteorganizationinvitation) |  Cancel One Organization Invitation | Stable
*OrganizationsApi* | [GetOrganization](./docs/OrganizationsApi.md#getorganization) |  Return One Organization | Stable
*OrganizationsApi* | [GetOrganizationInvitation](./docs/OrganizationsApi.md#getorganizationinvitation) |  Return One Organization Invitation | Stable
*OrganizationsApi* | [GetOrganizationSettings](./docs/OrganizationsApi.md#getorganizationsettings) |  Return Settings for One Organization | Experimental
*OrganizationsApi* | [ListOrganizationInvitations](./docs/OrganizationsApi.md#listorganizationinvitations) |  Return All Organization Invitations | Stable
*OrganizationsApi* | [ListOrganizationProjects](./docs/OrganizationsApi.md#listorganizationprojects) |  Return One or More Projects in One Organization | Stable
*OrganizationsApi* | [ListOrganizationUsers](./docs/OrganizationsApi.md#listorganizationusers) |  Return All MongoDB Cloud Users in One Organization | Stable
*OrganizationsApi* | [ListOrganizations](./docs/OrganizationsApi.md#listorganizations) |  Return All Organizations | Stable
*OrganizationsApi* | [RemoveOrganizationUser](./docs/OrganizationsApi.md#removeorganizationuser) |  Remove One MongoDB Cloud User from One Organization | Experimental
*OrganizationsApi* | [RenameOrganization](./docs/OrganizationsApi.md#renameorganization) |  Rename One Organization | Experimental
*OrganizationsApi* | [UpdateOrganizationInvitation](./docs/OrganizationsApi.md#updateorganizationinvitation) |  Update One Organization Invitation | Stable
*OrganizationsApi* | [UpdateOrganizationInvitationById](./docs/OrganizationsApi.md#updateorganizationinvitationbyid) |  Update One Organization Invitation by Invitation ID | Stable
*OrganizationsApi* | [UpdateOrganizationSettings](./docs/OrganizationsApi.md#updateorganizationsettings) |  Update Settings for One Organization | Experimental
*PerformanceAdvisorApi* | [DisableSlowOperationThresholding](./docs/PerformanceAdvisorApi.md#disableslowoperationthresholding) |  Disable Managed Slow Operation Threshold | Stable
*PerformanceAdvisorApi* | [EnableSlowOperationThresholding](./docs/PerformanceAdvisorApi.md#enableslowoperationthresholding) |  Enable Managed Slow Operation Threshold | Stable
*PerformanceAdvisorApi* | [ListSlowQueries](./docs/PerformanceAdvisorApi.md#listslowqueries) |  Return Slow Queries | Stable
*PerformanceAdvisorApi* | [ListSlowQueryNamespaces](./docs/PerformanceAdvisorApi.md#listslowquerynamespaces) |  Return All Namespaces for One Host | Stable
*PerformanceAdvisorApi* | [ListSuggestedIndexes](./docs/PerformanceAdvisorApi.md#listsuggestedindexes) |  Return Suggested Indexes | Stable
*PrivateEndpointServicesApi* | [CreatePrivateEndpoint](./docs/PrivateEndpointServicesApi.md#createprivateendpoint) |  Create One Private Endpoint for One Provider | Experimental
*PrivateEndpointServicesApi* | [CreatePrivateEndpointService](./docs/PrivateEndpointServicesApi.md#createprivateendpointservice) |  Create One Private Endpoint Service for One Provider | Stable
*PrivateEndpointServicesApi* | [DeletePrivateEndpoint](./docs/PrivateEndpointServicesApi.md#deleteprivateendpoint) |  Remove One Private Endpoint for One Provider | Stable
*PrivateEndpointServicesApi* | [DeletePrivateEndpointService](./docs/PrivateEndpointServicesApi.md#deleteprivateendpointservice) |  Remove One Private Endpoint Service for One Provider | Stable
*PrivateEndpointServicesApi* | [GetPrivateEndpoint](./docs/PrivateEndpointServicesApi.md#getprivateendpoint) |  Return One Private Endpoint for One Provider | Stable
*PrivateEndpointServicesApi* | [GetPrivateEndpointService](./docs/PrivateEndpointServicesApi.md#getprivateendpointservice) |  Return One Private Endpoint Service for One Provider | Stable
*PrivateEndpointServicesApi* | [GetRegionalizedPrivateEndpointSetting](./docs/PrivateEndpointServicesApi.md#getregionalizedprivateendpointsetting) |  Return Regionalized Private Endpoint Status | Stable
*PrivateEndpointServicesApi* | [ListPrivateEndpointServices](./docs/PrivateEndpointServicesApi.md#listprivateendpointservices) |  Return All Private Endpoint Services for One Provider | Stable
*PrivateEndpointServicesApi* | [ToggleRegionalizedPrivateEndpointSetting](./docs/PrivateEndpointServicesApi.md#toggleregionalizedprivateendpointsetting) |  Toggle Regionalized Private Endpoint Status | Stable
*ProgrammaticAPIKeysApi* | [AddProjectApiKey](./docs/ProgrammaticAPIKeysApi.md#addprojectapikey) |  Assign One Organization API Key to One Project | Experimental
*ProgrammaticAPIKeysApi* | [CreateApiKey](./docs/ProgrammaticAPIKeysApi.md#createapikey) |  Create One Organization API Key | Stable
*ProgrammaticAPIKeysApi* | [CreateApiKeyAccessList](./docs/ProgrammaticAPIKeysApi.md#createapikeyaccesslist) |  Create Access List Entries for One Organization API Key | Stable
*ProgrammaticAPIKeysApi* | [CreateProjectApiKey](./docs/ProgrammaticAPIKeysApi.md#createprojectapikey) |  Create and Assign One Organization API Key to One Project | Stable
*ProgrammaticAPIKeysApi* | [DeleteApiKey](./docs/ProgrammaticAPIKeysApi.md#deleteapikey) |  Remove One Organization API Key | Stable
*ProgrammaticAPIKeysApi* | [DeleteApiKeyAccessListEntry](./docs/ProgrammaticAPIKeysApi.md#deleteapikeyaccesslistentry) |  Remove One Access List Entry for One Organization API Key | Stable
*ProgrammaticAPIKeysApi* | [GetApiKey](./docs/ProgrammaticAPIKeysApi.md#getapikey) |  Return One Organization API Key | Stable
*ProgrammaticAPIKeysApi* | [GetApiKeyAccessList](./docs/ProgrammaticAPIKeysApi.md#getapikeyaccesslist) |  Return One Access List Entry for One Organization API Key | Experimental
*ProgrammaticAPIKeysApi* | [ListApiKeyAccessListsEntries](./docs/ProgrammaticAPIKeysApi.md#listapikeyaccesslistsentries) |  Return All Access List Entries for One Organization API Key | Stable
*ProgrammaticAPIKeysApi* | [ListApiKeys](./docs/ProgrammaticAPIKeysApi.md#listapikeys) |  Return All Organization API Keys | Stable
*ProgrammaticAPIKeysApi* | [ListProjectApiKeys](./docs/ProgrammaticAPIKeysApi.md#listprojectapikeys) |  Return All Organization API Keys Assigned to One Project | Stable
*ProgrammaticAPIKeysApi* | [RemoveProjectApiKey](./docs/ProgrammaticAPIKeysApi.md#removeprojectapikey) |  Unassign One Organization API Key from One Project | Stable
*ProgrammaticAPIKeysApi* | [UpdateApiKey](./docs/ProgrammaticAPIKeysApi.md#updateapikey) |  Update One Organization API Key | Stable
*ProgrammaticAPIKeysApi* | [UpdateApiKeyRoles](./docs/ProgrammaticAPIKeysApi.md#updateapikeyroles) |  Update Roles of One Organization API Key to One Project | Stable
*ProjectIPAccessListApi* | [CreateProjectIpAccessList](./docs/ProjectIPAccessListApi.md#createprojectipaccesslist) |  Add Entries to Project IP Access List | Stable
*ProjectIPAccessListApi* | [DeleteProjectIpAccessList](./docs/ProjectIPAccessListApi.md#deleteprojectipaccesslist) |  Remove One Entry from One Project IP Access List | Stable
*ProjectIPAccessListApi* | [GetProjectIpAccessListStatus](./docs/ProjectIPAccessListApi.md#getprojectipaccessliststatus) |  Return Status of One Project IP Access List Entry | Experimental
*ProjectIPAccessListApi* | [GetProjectIpList](./docs/ProjectIPAccessListApi.md#getprojectiplist) |  Return One Project IP Access List Entry | Stable
*ProjectIPAccessListApi* | [ListProjectIpAccessLists](./docs/ProjectIPAccessListApi.md#listprojectipaccesslists) |  Return Project IP Access List | Stable
*ProjectsApi* | [CreateProject](./docs/ProjectsApi.md#createproject) |  Create One Project | Stable
*ProjectsApi* | [CreateProjectInvitation](./docs/ProjectsApi.md#createprojectinvitation) |  Invite One MongoDB Cloud User to Join One Project | Stable
*ProjectsApi* | [DeleteProject](./docs/ProjectsApi.md#deleteproject) |  Remove One Project | Stable
*ProjectsApi* | [DeleteProjectInvitation](./docs/ProjectsApi.md#deleteprojectinvitation) |  Cancel One Project Invitation | Stable
*ProjectsApi* | [DeleteProjectLimit](./docs/ProjectsApi.md#deleteprojectlimit) |  Remove One Project Limit | Experimental
*ProjectsApi* | [GetProject](./docs/ProjectsApi.md#getproject) |  Return One Project | Stable
*ProjectsApi* | [GetProjectByName](./docs/ProjectsApi.md#getprojectbyname) |  Return One Project using Its Name | Stable
*ProjectsApi* | [GetProjectInvitation](./docs/ProjectsApi.md#getprojectinvitation) |  Return One Project Invitation | Stable
*ProjectsApi* | [GetProjectLimit](./docs/ProjectsApi.md#getprojectlimit) |  Return One Limit for One Project | Experimental
*ProjectsApi* | [GetProjectSettings](./docs/ProjectsApi.md#getprojectsettings) |  Return One Project Settings | Stable
*ProjectsApi* | [ListProjectInvitations](./docs/ProjectsApi.md#listprojectinvitations) |  Return All Project Invitations | Stable
*ProjectsApi* | [ListProjectLimits](./docs/ProjectsApi.md#listprojectlimits) |  Return All Limits for One Project | Experimental
*ProjectsApi* | [ListProjectUsers](./docs/ProjectsApi.md#listprojectusers) |  Return All Users in One Project | Stable
*ProjectsApi* | [ListProjects](./docs/ProjectsApi.md#listprojects) |  Return All Projects | Stable
*ProjectsApi* | [RemoveProjectUser](./docs/ProjectsApi.md#removeprojectuser) |  Remove One User from One Project | Stable
*ProjectsApi* | [SetProjectLimit](./docs/ProjectsApi.md#setprojectlimit) |  Set One Project Limit | Experimental
*ProjectsApi* | [UpdateProject](./docs/ProjectsApi.md#updateproject) |  Update One Project Name | Experimental
*ProjectsApi* | [UpdateProjectInvitation](./docs/ProjectsApi.md#updateprojectinvitation) |  Update One Project Invitation | Stable
*ProjectsApi* | [UpdateProjectInvitationById](./docs/ProjectsApi.md#updateprojectinvitationbyid) |  Update One Project Invitation by Invitation ID | Stable
*ProjectsApi* | [UpdateProjectSettings](./docs/ProjectsApi.md#updateprojectsettings) |  Update One Project Settings | Stable
*RollingIndexApi* | [CreateRollingIndex](./docs/RollingIndexApi.md#createrollingindex) |  Create One Rolling Index | Stable
*RootApi* | [GetSystemStatus](./docs/RootApi.md#getsystemstatus) |  Return the status of this MongoDB application | Experimental
*ServerlessInstancesApi* | [CreateServerlessInstance](./docs/ServerlessInstancesApi.md#createserverlessinstance) |  Create One Serverless Instance in One Project | Stable
*ServerlessInstancesApi* | [DeleteServerlessInstance](./docs/ServerlessInstancesApi.md#deleteserverlessinstance) |  Remove One Serverless Instance from One Project | Stable
*ServerlessInstancesApi* | [GetServerlessInstance](./docs/ServerlessInstancesApi.md#getserverlessinstance) |  Return One Serverless Instance from One Project | Stable
*ServerlessInstancesApi* | [ListServerlessInstances](./docs/ServerlessInstancesApi.md#listserverlessinstances) |  Return All Serverless Instances from One Project | Stable
*ServerlessInstancesApi* | [UpdateServerlessInstance](./docs/ServerlessInstancesApi.md#updateserverlessinstance) |  Update One Serverless Instance in One Project | Stable
*ServerlessPrivateEndpointsApi* | [CreateServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#createserverlessprivateendpoint) |  Create One Private Endpoint for One Serverless Instance | Experimental
*ServerlessPrivateEndpointsApi* | [DeleteServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#deleteserverlessprivateendpoint) |  Remove One Private Endpoint for One Serverless Instance | Experimental
*ServerlessPrivateEndpointsApi* | [GetServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#getserverlessprivateendpoint) |  Return One Private Endpoint for One Serverless Instance | Experimental
*ServerlessPrivateEndpointsApi* | [ListServerlessPrivateEndpoints](./docs/ServerlessPrivateEndpointsApi.md#listserverlessprivateendpoints) |  Return All Private Endpoints for One Serverless Instance | Stable
*ServerlessPrivateEndpointsApi* | [UpdateServerlessPrivateEndpoint](./docs/ServerlessPrivateEndpointsApi.md#updateserverlessprivateendpoint) |  Update One Private Endpoint for One Serverless Instance | Experimental
*SharedTierRestoreJobsApi* | [CreateSharedClusterBackupRestoreJob](./docs/SharedTierRestoreJobsApi.md#createsharedclusterbackuprestorejob) |  Create One Restore Job from One M2 or M5 Cluster | Experimental
*SharedTierRestoreJobsApi* | [GetSharedClusterBackupRestoreJob](./docs/SharedTierRestoreJobsApi.md#getsharedclusterbackuprestorejob) |  Return One Restore Job for One M2 or M5 Cluster | Experimental
*SharedTierRestoreJobsApi* | [ListSharedClusterBackupRestoreJobs](./docs/SharedTierRestoreJobsApi.md#listsharedclusterbackuprestorejobs) |  Return All Restore Jobs for One M2 or M5 Cluster | Experimental
*SharedTierSnapshotsApi* | [DownloadSharedClusterBackup](./docs/SharedTierSnapshotsApi.md#downloadsharedclusterbackup) |  Download One M2 or M5 Cluster Snapshot | Experimental
*SharedTierSnapshotsApi* | [GetSharedClusterBackup](./docs/SharedTierSnapshotsApi.md#getsharedclusterbackup) |  Return One Snapshot for One M2 or M5 Cluster | Experimental
*SharedTierSnapshotsApi* | [ListSharedClusterBackups](./docs/SharedTierSnapshotsApi.md#listsharedclusterbackups) |  Return All Snapshots for One M2 or M5 Cluster | Experimental
*TeamsApi* | [AddAllTeamsToProject](./docs/TeamsApi.md#addallteamstoproject) |  Add One or More Teams to One Project | Stable
*TeamsApi* | [AddTeamUser](./docs/TeamsApi.md#addteamuser) |  Assign MongoDB Cloud Users from One Organization to One Team | Stable
*TeamsApi* | [CreateTeam](./docs/TeamsApi.md#createteam) |  Create One Team in One Organization | Stable
*TeamsApi* | [DeleteTeam](./docs/TeamsApi.md#deleteteam) |  Remove One Team from One Organization | Stable
*TeamsApi* | [GetTeamById](./docs/TeamsApi.md#getteambyid) |  Return One Team using its ID | Stable
*TeamsApi* | [GetTeamByName](./docs/TeamsApi.md#getteambyname) |  Return One Team using its Name | Stable
*TeamsApi* | [ListOrganizationTeams](./docs/TeamsApi.md#listorganizationteams) |  Return All Teams in One Organization | Stable
*TeamsApi* | [ListProjectTeams](./docs/TeamsApi.md#listprojectteams) |  Return All Teams in One Project | Stable
*TeamsApi* | [ListTeamUsers](./docs/TeamsApi.md#listteamusers) |  Return All MongoDB Cloud Users Assigned to One Team | Stable
*TeamsApi* | [RemoveProjectTeam](./docs/TeamsApi.md#removeprojectteam) |  Remove One Team from One Project | Stable
*TeamsApi* | [RemoveTeamUser](./docs/TeamsApi.md#removeteamuser) |  Remove One MongoDB Cloud User from One Team | Stable
*TeamsApi* | [RenameTeam](./docs/TeamsApi.md#renameteam) |  Rename One Team | Experimental
*TeamsApi* | [UpdateTeamRoles](./docs/TeamsApi.md#updateteamroles) |  Update Team Roles in One Project | Stable
*ThirdPartyIntegrationsApi* | [CreateThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#createthirdpartyintegration) |  Configure One Third-Party Service Integration | Stable
*ThirdPartyIntegrationsApi* | [DeleteThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#deletethirdpartyintegration) |  Remove One Third-Party Service Integration | Stable
*ThirdPartyIntegrationsApi* | [GetThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#getthirdpartyintegration) |  Return One Third-Party Service Integration | Stable
*ThirdPartyIntegrationsApi* | [ListThirdPartyIntegrations](./docs/ThirdPartyIntegrationsApi.md#listthirdpartyintegrations) |  Return All Active Third-Party Service Integrations | Stable
*ThirdPartyIntegrationsApi* | [UpdateThirdPartyIntegration](./docs/ThirdPartyIntegrationsApi.md#updatethirdpartyintegration) |  Update One Third-Party Service Integration | Experimental
*X509AuthenticationApi* | [CreateDatabaseUserCertificate](./docs/X509AuthenticationApi.md#createdatabaseusercertificate) |  Create One X.509 Certificate for One MongoDB User | Stable
*X509AuthenticationApi* | [DisableCustomerManagedX509](./docs/X509AuthenticationApi.md#disablecustomermanagedx509) |  Disable Customer-Managed X.509 | Stable
*X509AuthenticationApi* | [ListDatabaseUserCertificates](./docs/X509AuthenticationApi.md#listdatabaseusercertificates) |  Return All X.509 Certificates Assigned to One MongoDB User | Stable


## Documentation For Models

 - [AWSCloudProviderContainer](./docs/AWSCloudProviderContainer.md)
 - [AWSCloudProviderSettings](./docs/AWSCloudProviderSettings.md)
 - [AWSComputeAutoScaling](./docs/AWSComputeAutoScaling.md)
 - [AWSCustomDNSEnabled](./docs/AWSCustomDNSEnabled.md)
 - [AWSInterfaceEndpoint](./docs/AWSInterfaceEndpoint.md)
 - [AWSKMSConfiguration](./docs/AWSKMSConfiguration.md)
 - [AWSPeerVpc](./docs/AWSPeerVpc.md)
 - [AccessListItem](./docs/AccessListItem.md)
 - [AddUserToTeam](./docs/AddUserToTeam.md)
 - [AdvancedAutoScalingSettings](./docs/AdvancedAutoScalingSettings.md)
 - [AdvancedClusterDescription](./docs/AdvancedClusterDescription.md)
 - [AdvancedComputeAutoScaling](./docs/AdvancedComputeAutoScaling.md)
 - [AlertViewForNdsGroup](./docs/AlertViewForNdsGroup.md)
 - [AlertsNotificationRootForGroup](./docs/AlertsNotificationRootForGroup.md)
 - [AlertsThresholdInteger](./docs/AlertsThresholdInteger.md)
 - [AlertsToggle](./docs/AlertsToggle.md)
 - [AuditLog](./docs/AuditLog.md)
 - [AuthFederationRoleMapping](./docs/AuthFederationRoleMapping.md)
 - [AutoExportPolicy](./docs/AutoExportPolicy.md)
 - [AvailableClustersDeployment](./docs/AvailableClustersDeployment.md)
 - [AzureCloudProviderContainer](./docs/AzureCloudProviderContainer.md)
 - [AzureCloudProviderSettings](./docs/AzureCloudProviderSettings.md)
 - [AzureComputeAutoScalingRules](./docs/AzureComputeAutoScalingRules.md)
 - [AzureKeyVault](./docs/AzureKeyVault.md)
 - [AzurePeerNetwork](./docs/AzurePeerNetwork.md)
 - [AzurePrivateEndpoint](./docs/AzurePrivateEndpoint.md)
 - [BSONTimestamp](./docs/BSONTimestamp.md)
 - [BackupLabel](./docs/BackupLabel.md)
 - [BackupOnlineArchive](./docs/BackupOnlineArchive.md)
 - [BackupRestoreJob](./docs/BackupRestoreJob.md)
 - [BackupRestoreJobDelivery](./docs/BackupRestoreJobDelivery.md)
 - [BackupSnapshot](./docs/BackupSnapshot.md)
 - [BackupSnapshotPart](./docs/BackupSnapshotPart.md)
 - [BackupSnapshotRetention](./docs/BackupSnapshotRetention.md)
 - [BackupTenantSnapshot](./docs/BackupTenantSnapshot.md)
 - [BiConnector](./docs/BiConnector.md)
 - [BillingInvoice](./docs/BillingInvoice.md)
 - [BillingPayment](./docs/BillingPayment.md)
 - [BillingRefund](./docs/BillingRefund.md)
 - [Checkpoint](./docs/Checkpoint.md)
 - [CheckpointPart](./docs/CheckpointPart.md)
 - [CloudCluster](./docs/CloudCluster.md)
 - [CloudDatabaseUser](./docs/CloudDatabaseUser.md)
 - [CloudGCPProviderSettings](./docs/CloudGCPProviderSettings.md)
 - [CloudProviderAWSAutoScaling](./docs/CloudProviderAWSAutoScaling.md)
 - [CloudProviderAccessAWSIAMRole](./docs/CloudProviderAccessAWSIAMRole.md)
 - [CloudProviderAccessAzureServicePrincipal](./docs/CloudProviderAccessAzureServicePrincipal.md)
 - [CloudProviderAccessDataLakeFeatureUsage](./docs/CloudProviderAccessDataLakeFeatureUsage.md)
 - [CloudProviderAccessEncryptionAtRestFeatureUsage](./docs/CloudProviderAccessEncryptionAtRestFeatureUsage.md)
 - [CloudProviderAccessExportSnapshotFeatureUsage](./docs/CloudProviderAccessExportSnapshotFeatureUsage.md)
 - [CloudProviderAccessFeatureUsage](./docs/CloudProviderAccessFeatureUsage.md)
 - [CloudProviderAccessFeatureUsageDataLakeFeatureId](./docs/CloudProviderAccessFeatureUsageDataLakeFeatureId.md)
 - [CloudProviderAccessFeatureUsageExportSnapshotFeatureId](./docs/CloudProviderAccessFeatureUsageExportSnapshotFeatureId.md)
 - [CloudProviderAccessRole](./docs/CloudProviderAccessRole.md)
 - [CloudProviderAccessRoles](./docs/CloudProviderAccessRoles.md)
 - [CloudProviderAzureAutoScaling](./docs/CloudProviderAzureAutoScaling.md)
 - [CloudProviderContainer](./docs/CloudProviderContainer.md)
 - [CloudProviderEndpointServiceRequest](./docs/CloudProviderEndpointServiceRequest.md)
 - [CloudProviderGCPAutoScaling](./docs/CloudProviderGCPAutoScaling.md)
 - [CloudProviderRegions](./docs/CloudProviderRegions.md)
 - [CloudRegionConfig](./docs/CloudRegionConfig.md)
 - [CloudRoleAssignment](./docs/CloudRoleAssignment.md)
 - [CloudSearchMetrics](./docs/CloudSearchMetrics.md)
 - [CloudUser](./docs/CloudUser.md)
 - [ClusterAutoScalingSettings](./docs/ClusterAutoScalingSettings.md)
 - [ClusterComputeAutoScaling](./docs/ClusterComputeAutoScaling.md)
 - [ClusterConnectionStrings](./docs/ClusterConnectionStrings.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpoint](./docs/ClusterDescriptionConnectionStringsPrivateEndpoint.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpointEndpoint](./docs/ClusterDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ClusterDescriptionProcessArgs](./docs/ClusterDescriptionProcessArgs.md)
 - [ClusterFreeAutoScaling](./docs/ClusterFreeAutoScaling.md)
 - [ClusterFreeProviderSettings](./docs/ClusterFreeProviderSettings.md)
 - [ClusterOutageSimulation](./docs/ClusterOutageSimulation.md)
 - [ClusterOutageSimulationOutageFilter](./docs/ClusterOutageSimulationOutageFilter.md)
 - [ClusterProviderSettings](./docs/ClusterProviderSettings.md)
 - [ClusterSearchIndex](./docs/ClusterSearchIndex.md)
 - [ClusterServerlessBackupOptions](./docs/ClusterServerlessBackupOptions.md)
 - [ClusterStatus](./docs/ClusterStatus.md)
 - [Collation](./docs/Collation.md)
 - [ComponentLabel](./docs/ComponentLabel.md)
 - [ConnectedOrgConfig](./docs/ConnectedOrgConfig.md)
 - [CreateAWSEndpointRequest](./docs/CreateAWSEndpointRequest.md)
 - [CreateAzureEndpointRequest](./docs/CreateAzureEndpointRequest.md)
 - [CreateEndpointRequest](./docs/CreateEndpointRequest.md)
 - [CreateGCPEndpointGroupRequest](./docs/CreateGCPEndpointGroupRequest.md)
 - [CreateGCPForwardingRuleRequest](./docs/CreateGCPForwardingRuleRequest.md)
 - [CreateOrganizationKey](./docs/CreateOrganizationKey.md)
 - [CreateOrganizationRequest](./docs/CreateOrganizationRequest.md)
 - [CreateOrganizationResponse](./docs/CreateOrganizationResponse.md)
 - [CreateProjectKey](./docs/CreateProjectKey.md)
 - [Criteria](./docs/Criteria.md)
 - [CustomCriteria](./docs/CustomCriteria.md)
 - [DBUserTLSX509Settings](./docs/DBUserTLSX509Settings.md)
 - [DLSIngestionSink](./docs/DLSIngestionSink.md)
 - [DailySchedule](./docs/DailySchedule.md)
 - [DataFederationLimit](./docs/DataFederationLimit.md)
 - [DataFederationQueryLimit](./docs/DataFederationQueryLimit.md)
 - [DataFederationTenantQueryLimit](./docs/DataFederationTenantQueryLimit.md)
 - [DataLakeAWSCloudProviderConfig](./docs/DataLakeAWSCloudProviderConfig.md)
 - [DataLakeApiBase](./docs/DataLakeApiBase.md)
 - [DataLakeAtlasStoreInstance](./docs/DataLakeAtlasStoreInstance.md)
 - [DataLakeAtlasStoreReadPreference](./docs/DataLakeAtlasStoreReadPreference.md)
 - [DataLakeAtlasStoreReadPreferenceTag](./docs/DataLakeAtlasStoreReadPreferenceTag.md)
 - [DataLakeCloudProviderConfig](./docs/DataLakeCloudProviderConfig.md)
 - [DataLakeDataProcessRegion](./docs/DataLakeDataProcessRegion.md)
 - [DataLakeDatabaseCollection](./docs/DataLakeDatabaseCollection.md)
 - [DataLakeDatabaseDataSourceSettings](./docs/DataLakeDatabaseDataSourceSettings.md)
 - [DataLakeDatabaseInstance](./docs/DataLakeDatabaseInstance.md)
 - [DataLakeHTTPStore](./docs/DataLakeHTTPStore.md)
 - [DataLakeIngestionPipeline](./docs/DataLakeIngestionPipeline.md)
 - [DataLakeOnlineArchiveStore](./docs/DataLakeOnlineArchiveStore.md)
 - [DataLakeS3StoreSettings](./docs/DataLakeS3StoreSettings.md)
 - [DataLakeStorage](./docs/DataLakeStorage.md)
 - [DataLakeStoreSettings](./docs/DataLakeStoreSettings.md)
 - [DataLakeTenant](./docs/DataLakeTenant.md)
 - [DataMetricThreshold](./docs/DataMetricThreshold.md)
 - [DataProtectionSettings](./docs/DataProtectionSettings.md)
 - [DatabaseInheritedRole](./docs/DatabaseInheritedRole.md)
 - [DatabasePermittedNamespaceResource](./docs/DatabasePermittedNamespaceResource.md)
 - [DatabasePrivilegeAction](./docs/DatabasePrivilegeAction.md)
 - [DatabaseUserRole](./docs/DatabaseUserRole.md)
 - [Datadog](./docs/Datadog.md)
 - [DateCriteria](./docs/DateCriteria.md)
 - [DedicatedHardwareSpec](./docs/DedicatedHardwareSpec.md)
 - [DefaultLimit](./docs/DefaultLimit.md)
 - [DefaultSchedule](./docs/DefaultSchedule.md)
 - [DeleteCopiedBackups](./docs/DeleteCopiedBackups.md)
 - [Destination](./docs/Destination.md)
 - [DiskBackupBaseRestoreMember](./docs/DiskBackupBaseRestoreMember.md)
 - [DiskBackupCopySetting](./docs/DiskBackupCopySetting.md)
 - [DiskBackupExportJob](./docs/DiskBackupExportJob.md)
 - [DiskBackupExportJobRequest](./docs/DiskBackupExportJobRequest.md)
 - [DiskBackupOnDemandSnapshotRequest](./docs/DiskBackupOnDemandSnapshotRequest.md)
 - [DiskBackupReplicaSet](./docs/DiskBackupReplicaSet.md)
 - [DiskBackupShardedClusterSnapshot](./docs/DiskBackupShardedClusterSnapshot.md)
 - [DiskBackupShardedClusterSnapshotMember](./docs/DiskBackupShardedClusterSnapshotMember.md)
 - [DiskBackupSnapshot](./docs/DiskBackupSnapshot.md)
 - [DiskBackupSnapshotAWSExportBucket](./docs/DiskBackupSnapshotAWSExportBucket.md)
 - [DiskBackupSnapshotRestoreJob](./docs/DiskBackupSnapshotRestoreJob.md)
 - [DiskBackupSnapshotSchedule](./docs/DiskBackupSnapshotSchedule.md)
 - [DiskGBAutoScaling](./docs/DiskGBAutoScaling.md)
 - [EncryptionAtRest](./docs/EncryptionAtRest.md)
 - [EndpointService](./docs/EndpointService.md)
 - [Error](./docs/Error.md)
 - [EventViewForNdsGroup](./docs/EventViewForNdsGroup.md)
 - [EventViewForOrg](./docs/EventViewForOrg.md)
 - [ExportStatus](./docs/ExportStatus.md)
 - [FTSAnalyzers](./docs/FTSAnalyzers.md)
 - [FTSAnalyzersCharFiltersInner](./docs/FTSAnalyzersCharFiltersInner.md)
 - [FTSAnalyzersCharFiltersInnerMappings](./docs/FTSAnalyzersCharFiltersInnerMappings.md)
 - [FTSAnalyzersTokenFiltersInner](./docs/FTSAnalyzersTokenFiltersInner.md)
 - [FTSAnalyzersTokenizer](./docs/FTSAnalyzersTokenizer.md)
 - [FTSMappings](./docs/FTSMappings.md)
 - [FTSMetric](./docs/FTSMetric.md)
 - [FederatedUser](./docs/FederatedUser.md)
 - [FederationIdentityProvider](./docs/FederationIdentityProvider.md)
 - [FieldTransformation](./docs/FieldTransformation.md)
 - [GCPCloudProviderContainer](./docs/GCPCloudProviderContainer.md)
 - [GCPComputeAutoScaling](./docs/GCPComputeAutoScaling.md)
 - [GCPConsumerForwardingRule](./docs/GCPConsumerForwardingRule.md)
 - [GCPPeerVpc](./docs/GCPPeerVpc.md)
 - [GeoSharding](./docs/GeoSharding.md)
 - [GoogleCloudKMS](./docs/GoogleCloudKMS.md)
 - [Group](./docs/Group.md)
 - [GroupAlertsConfig](./docs/GroupAlertsConfig.md)
 - [GroupInvitation](./docs/GroupInvitation.md)
 - [GroupInvitationRequest](./docs/GroupInvitationRequest.md)
 - [GroupInvitationUpdateRequest](./docs/GroupInvitationUpdateRequest.md)
 - [GroupMaintenanceWindow](./docs/GroupMaintenanceWindow.md)
 - [GroupName](./docs/GroupName.md)
 - [GroupPaginatedEvent](./docs/GroupPaginatedEvent.md)
 - [GroupSettings](./docs/GroupSettings.md)
 - [HardwareSpec](./docs/HardwareSpec.md)
 - [HostMetricValue](./docs/HostMetricValue.md)
 - [HostViewAtlas](./docs/HostViewAtlas.md)
 - [IndexOptions](./docs/IndexOptions.md)
 - [IngestionPipelineRun](./docs/IngestionPipelineRun.md)
 - [IngestionSink](./docs/IngestionSink.md)
 - [IngestionSource](./docs/IngestionSource.md)
 - [InvoiceLineItem](./docs/InvoiceLineItem.md)
 - [Key](./docs/Key.md)
 - [KeyUser](./docs/KeyUser.md)
 - [LDAPSecuritySettings](./docs/LDAPSecuritySettings.md)
 - [LDAPVerifyConnectivityJobRequest](./docs/LDAPVerifyConnectivityJobRequest.md)
 - [LDAPVerifyConnectivityJobRequestParams](./docs/LDAPVerifyConnectivityJobRequestParams.md)
 - [LDAPVerifyConnectivityJobRequestValidation](./docs/LDAPVerifyConnectivityJobRequestValidation.md)
 - [LegacyAtlasCluster](./docs/LegacyAtlasCluster.md)
 - [LegacyReplicationSpec](./docs/LegacyReplicationSpec.md)
 - [Link](./docs/Link.md)
 - [LinkAtlas](./docs/LinkAtlas.md)
 - [LiveImportAvailableProject](./docs/LiveImportAvailableProject.md)
 - [LiveImportValidation](./docs/LiveImportValidation.md)
 - [LiveMigrationRequest](./docs/LiveMigrationRequest.md)
 - [LiveMigrationResponse](./docs/LiveMigrationResponse.md)
 - [ManagedNamespace](./docs/ManagedNamespace.md)
 - [ManagedNamespaces](./docs/ManagedNamespaces.md)
 - [MeasurementDiskPartition](./docs/MeasurementDiskPartition.md)
 - [MeasurementsGeneralViewAtlas](./docs/MeasurementsGeneralViewAtlas.md)
 - [MeasurementsIndexes](./docs/MeasurementsIndexes.md)
 - [MeasurementsNonIndex](./docs/MeasurementsNonIndex.md)
 - [MesurementsDatabase](./docs/MesurementsDatabase.md)
 - [MetricDataPoint](./docs/MetricDataPoint.md)
 - [MetricDataPointAtlas](./docs/MetricDataPointAtlas.md)
 - [MetricsMeasurement](./docs/MetricsMeasurement.md)
 - [MetricsMeasurementAtlas](./docs/MetricsMeasurementAtlas.md)
 - [MicrosoftTeams](./docs/MicrosoftTeams.md)
 - [MongoDBAccessLogs](./docs/MongoDBAccessLogs.md)
 - [MongoDBAccessLogsList](./docs/MongoDBAccessLogsList.md)
 - [MonthlySchedule](./docs/MonthlySchedule.md)
 - [NamespaceObj](./docs/NamespaceObj.md)
 - [Namespaces](./docs/Namespaces.md)
 - [NetworkContainerPeer](./docs/NetworkContainerPeer.md)
 - [NetworkPermissionEntry](./docs/NetworkPermissionEntry.md)
 - [NetworkPermissionEntryStatus](./docs/NetworkPermissionEntryStatus.md)
 - [NewRelic](./docs/NewRelic.md)
 - [OnDemandCpsSnapshotSource](./docs/OnDemandCpsSnapshotSource.md)
 - [OnlineArchiveSchedule](./docs/OnlineArchiveSchedule.md)
 - [OpsGenie](./docs/OpsGenie.md)
 - [OrgFederationSettings](./docs/OrgFederationSettings.md)
 - [OrgGroup](./docs/OrgGroup.md)
 - [OrgPaginatedEvent](./docs/OrgPaginatedEvent.md)
 - [Organization](./docs/Organization.md)
 - [OrganizationInvitation](./docs/OrganizationInvitation.md)
 - [OrganizationInvitationRequest](./docs/OrganizationInvitationRequest.md)
 - [OrganizationInvitationUpdateRequest](./docs/OrganizationInvitationUpdateRequest.md)
 - [OrganizationSettings](./docs/OrganizationSettings.md)
 - [PagerDuty](./docs/PagerDuty.md)
 - [PaginatedAdvancedClusterDescription](./docs/PaginatedAdvancedClusterDescription.md)
 - [PaginatedAlert](./docs/PaginatedAlert.md)
 - [PaginatedAlertConfig](./docs/PaginatedAlertConfig.md)
 - [PaginatedApiApiUser](./docs/PaginatedApiApiUser.md)
 - [PaginatedApiAppUser](./docs/PaginatedApiAppUser.md)
 - [PaginatedApiAtlasCheckpoint](./docs/PaginatedApiAtlasCheckpoint.md)
 - [PaginatedApiAtlasDatabaseUser](./docs/PaginatedApiAtlasDatabaseUser.md)
 - [PaginatedApiAtlasDiskBackupExportJob](./docs/PaginatedApiAtlasDiskBackupExportJob.md)
 - [PaginatedApiAtlasProviderRegions](./docs/PaginatedApiAtlasProviderRegions.md)
 - [PaginatedApiAtlasServerlessBackupRestoreJob](./docs/PaginatedApiAtlasServerlessBackupRestoreJob.md)
 - [PaginatedApiAtlasServerlessBackupSnapshot](./docs/PaginatedApiAtlasServerlessBackupSnapshot.md)
 - [PaginatedApiInvoice](./docs/PaginatedApiInvoice.md)
 - [PaginatedApiUserAccessList](./docs/PaginatedApiUserAccessList.md)
 - [PaginatedAppUser](./docs/PaginatedAppUser.md)
 - [PaginatedAtlasGroup](./docs/PaginatedAtlasGroup.md)
 - [PaginatedBackupSnapshot](./docs/PaginatedBackupSnapshot.md)
 - [PaginatedBackupSnapshotExportBucket](./docs/PaginatedBackupSnapshotExportBucket.md)
 - [PaginatedCloudBackupReplicaSet](./docs/PaginatedCloudBackupReplicaSet.md)
 - [PaginatedCloudBackupRestoreJob](./docs/PaginatedCloudBackupRestoreJob.md)
 - [PaginatedCloudBackupShardedClusterSnapshot](./docs/PaginatedCloudBackupShardedClusterSnapshot.md)
 - [PaginatedCloudProviderContainer](./docs/PaginatedCloudProviderContainer.md)
 - [PaginatedContainerPeer](./docs/PaginatedContainerPeer.md)
 - [PaginatedDatabase](./docs/PaginatedDatabase.md)
 - [PaginatedDiskPartition](./docs/PaginatedDiskPartition.md)
 - [PaginatedHostViewAtlas](./docs/PaginatedHostViewAtlas.md)
 - [PaginatedIntegration](./docs/PaginatedIntegration.md)
 - [PaginatedNetworkAccess](./docs/PaginatedNetworkAccess.md)
 - [PaginatedOnlineArchive](./docs/PaginatedOnlineArchive.md)
 - [PaginatedOrgGroup](./docs/PaginatedOrgGroup.md)
 - [PaginatedOrganization](./docs/PaginatedOrganization.md)
 - [PaginatedPipelineRun](./docs/PaginatedPipelineRun.md)
 - [PaginatedPrivateNetworkEndpointIdEntry](./docs/PaginatedPrivateNetworkEndpointIdEntry.md)
 - [PaginatedRestoreJob](./docs/PaginatedRestoreJob.md)
 - [PaginatedServerlessInstanceDescription](./docs/PaginatedServerlessInstanceDescription.md)
 - [PaginatedSnapshot](./docs/PaginatedSnapshot.md)
 - [PaginatedTeam](./docs/PaginatedTeam.md)
 - [PaginatedTeamRole](./docs/PaginatedTeamRole.md)
 - [PaginatedTenantRestore](./docs/PaginatedTenantRestore.md)
 - [PaginatedTenantSnapshot](./docs/PaginatedTenantSnapshot.md)
 - [PaginatedUserCert](./docs/PaginatedUserCert.md)
 - [PartitionField](./docs/PartitionField.md)
 - [PemFileInfo](./docs/PemFileInfo.md)
 - [PerformanceAdvisorIndex](./docs/PerformanceAdvisorIndex.md)
 - [PerformanceAdvisorOpStats](./docs/PerformanceAdvisorOpStats.md)
 - [PerformanceAdvisorOperation](./docs/PerformanceAdvisorOperation.md)
 - [PerformanceAdvisorResponse](./docs/PerformanceAdvisorResponse.md)
 - [PerformanceAdvisorShape](./docs/PerformanceAdvisorShape.md)
 - [PerformanceAdvisorSlowQuery](./docs/PerformanceAdvisorSlowQuery.md)
 - [PerformanceAdvisorSlowQueryList](./docs/PerformanceAdvisorSlowQueryList.md)
 - [PeriodicCpsSnapshotSource](./docs/PeriodicCpsSnapshotSource.md)
 - [PipelineRunStats](./docs/PipelineRunStats.md)
 - [PipelinesPartitionField](./docs/PipelinesPartitionField.md)
 - [Policy](./docs/Policy.md)
 - [PolicyItem](./docs/PolicyItem.md)
 - [PrivateGCPEndpointGroup](./docs/PrivateGCPEndpointGroup.md)
 - [PrivateIPMode](./docs/PrivateIPMode.md)
 - [PrivateLinkEndpoint](./docs/PrivateLinkEndpoint.md)
 - [PrivateNetworkEndpointIdEntry](./docs/PrivateNetworkEndpointIdEntry.md)
 - [ProjectSettingItem](./docs/ProjectSettingItem.md)
 - [Prometheus](./docs/Prometheus.md)
 - [RPUMetricThreshold](./docs/RPUMetricThreshold.md)
 - [Raw](./docs/Raw.md)
 - [RawMetricThreshold](./docs/RawMetricThreshold.md)
 - [RegionSpec](./docs/RegionSpec.md)
 - [ReplicationSpec](./docs/ReplicationSpec.md)
 - [ResourceTag](./docs/ResourceTag.md)
 - [RestoreJobFileHash](./docs/RestoreJobFileHash.md)
 - [RoleAssignment](./docs/RoleAssignment.md)
 - [RollingIndexRequest](./docs/RollingIndexRequest.md)
 - [SamlIdentityProviderUpdate](./docs/SamlIdentityProviderUpdate.md)
 - [SampleDatasetStatus](./docs/SampleDatasetStatus.md)
 - [SearchSynonymMappingDefinition](./docs/SearchSynonymMappingDefinition.md)
 - [ServerlessAWSTenantEndpointUpdate](./docs/ServerlessAWSTenantEndpointUpdate.md)
 - [ServerlessAzureTenantEndpointUpdate](./docs/ServerlessAzureTenantEndpointUpdate.md)
 - [ServerlessBackupRestoreJob](./docs/ServerlessBackupRestoreJob.md)
 - [ServerlessBackupSnapshot](./docs/ServerlessBackupSnapshot.md)
 - [ServerlessConnectionStringsPrivateEndpointItem](./docs/ServerlessConnectionStringsPrivateEndpointItem.md)
 - [ServerlessConnectionStringsPrivateEndpointList](./docs/ServerlessConnectionStringsPrivateEndpointList.md)
 - [ServerlessInstanceDescription](./docs/ServerlessInstanceDescription.md)
 - [ServerlessInstanceDescriptionConnectionStrings](./docs/ServerlessInstanceDescriptionConnectionStrings.md)
 - [ServerlessInstanceDescriptionCreate](./docs/ServerlessInstanceDescriptionCreate.md)
 - [ServerlessInstanceDescriptionUpdate](./docs/ServerlessInstanceDescriptionUpdate.md)
 - [ServerlessMetricThreshold](./docs/ServerlessMetricThreshold.md)
 - [ServerlessProviderSettings](./docs/ServerlessProviderSettings.md)
 - [ServerlessTenantCreateRequest](./docs/ServerlessTenantCreateRequest.md)
 - [ServerlessTenantEndpoint](./docs/ServerlessTenantEndpoint.md)
 - [ServerlessTenantEndpointUpdate](./docs/ServerlessTenantEndpointUpdate.md)
 - [Slack](./docs/Slack.md)
 - [SnapshotSchedule](./docs/SnapshotSchedule.md)
 - [Source](./docs/Source.md)
 - [SynonymSource](./docs/SynonymSource.md)
 - [SystemStatus](./docs/SystemStatus.md)
 - [TargetOrg](./docs/TargetOrg.md)
 - [TargetOrgRequest](./docs/TargetOrgRequest.md)
 - [Team](./docs/Team.md)
 - [TeamResponse](./docs/TeamResponse.md)
 - [TeamRole](./docs/TeamRole.md)
 - [TenantRestore](./docs/TenantRestore.md)
 - [ThridPartyIntegration](./docs/ThridPartyIntegration.md)
 - [TimeMetricThreshold](./docs/TimeMetricThreshold.md)
 - [TokenFilterasciiFolding](./docs/TokenFilterasciiFolding.md)
 - [TokenFilterdaitchMokotoffSoundex](./docs/TokenFilterdaitchMokotoffSoundex.md)
 - [TokenFilteredgeGram](./docs/TokenFilteredgeGram.md)
 - [TokenFiltericuFolding](./docs/TokenFiltericuFolding.md)
 - [TokenFiltericuNormalizer](./docs/TokenFiltericuNormalizer.md)
 - [TokenFilterlength](./docs/TokenFilterlength.md)
 - [TokenFilterlowercase](./docs/TokenFilterlowercase.md)
 - [TokenFilternGram](./docs/TokenFilternGram.md)
 - [TokenFilterregex](./docs/TokenFilterregex.md)
 - [TokenFilterreverse](./docs/TokenFilterreverse.md)
 - [TokenFiltershingle](./docs/TokenFiltershingle.md)
 - [TokenFiltersnowballStemming](./docs/TokenFiltersnowballStemming.md)
 - [TokenFilterstopword](./docs/TokenFilterstopword.md)
 - [TokenFiltertrim](./docs/TokenFiltertrim.md)
 - [TriggerIngestionPipelineRequest](./docs/TriggerIngestionPipelineRequest.md)
 - [UpdateCustomDBRole](./docs/UpdateCustomDBRole.md)
 - [UserAccessList](./docs/UserAccessList.md)
 - [UserCert](./docs/UserCert.md)
 - [UserCustomDBRole](./docs/UserCustomDBRole.md)
 - [UserRoleAssignment](./docs/UserRoleAssignment.md)
 - [UserScope](./docs/UserScope.md)
 - [UserSecurity](./docs/UserSecurity.md)
 - [UserToDNMapping](./docs/UserToDNMapping.md)
 - [VictorOps](./docs/VictorOps.md)
 - [Webhook](./docs/Webhook.md)
 - [WeeklySchedule](./docs/WeeklySchedule.md)
 - [X509Certificate](./docs/X509Certificate.md)




