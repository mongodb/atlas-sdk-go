
# # v20230201002.0.0

## Notes

 The Atlas Golang SDK official v20230201002.0.0 release. 

## Breaking Changes

1. All golang structures starting from `admin.Nullable` have been removed.
Structures weren't part of the API and should not be used by end users.
2. Data structures now use golang base types instead of `NullableString` or `NullableInt`.
For example:

`ThridPartyIntegration.ChannelName NullableString` => `ChannelName *string`

Following list of the models have been affected:
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
