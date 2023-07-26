## v20230202003

- `OnlineArchiveApi.CreateOnlineArchive` updated to take struct `BackupOnlineArchiveCreate` as input parameter
- `PerformanceAdvisorIndex` `Index` property corrected from `[]map[string]string` to `[]map[string]int`

## v20230202002

- Removed all internal golang structures starting with `admin.Nullable`.
- Data structures now use golang base types instead of `NullableString` or `NullableInt`. For example, the data structure
`ThirdPartyIntegration.ChannelName` `NullableString` now uses the` ChannelName *string` base type. 

This change affects the following models: 
```
admin/model_cluster_description_process_args.go
admin/model_live_import_validation.go
admin/model_live_migration_response.go
admin/model_system_status.go
admin/model_thrid_party_integration.go
```
