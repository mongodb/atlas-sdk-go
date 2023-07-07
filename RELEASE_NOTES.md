
# # v20230201002.0.0

## Notes

 The Atlas Golang SDK official v20230201002.0.0 release. 

## Breaking Changes

- Removed all internal golang structures starting with `admin.Nullable`.
- Data structures now use golang base types instead of `NullableString` or `NullableInt`. For example, the data structure
`ThirdPartyIntegration.ChannelName` NullableString now uses the ChannelName *string base type. This change affects the following models: 
```
admin/model_cluster_description_process_args.go
admin/model_live_import_validation.go
admin/model_live_migration_response.go
admin/model_system_status.go
admin/model_thrid_party_integration.go
```

## SDK documentation
 
Please refer to the official documentation

https://www.mongodb.com/docs/atlas/sdk/

## API documentation

SDK is based on Atlas Admin v2 API.
Currently supported version: 2023-02-01

For API documentation please refer to: 
https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/2023-02-01/


# v20230201001.0.0

## Notes

 The Atlas Golang SDK official v20230201001.0.0 release. 

## SDK documentation
 
Please refer to the official documentation

https://www.mongodb.com/docs/atlas/sdk/

## API documentation

SDK is based on Atlas Admin v2 API.
Currently supported version: 2023-02-01

For API documentation please refer to: 
https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/2023-02-01/
