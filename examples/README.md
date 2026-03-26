# Go SDK examples

## Running Examples

For documentation about obtaining apiKey and apiSecret go to
https://docs.atlas.mongodb.com/configure-api-access.

```bash
export MONGODB_ATLAS_PUBLIC_KEY=somekey 
export MONGODB_ATLAS_PRIVATE_KEY=some-secret-key-for-gosdkapi
go run ./aws_cluster/aws.go
```

## Running Examples with Mocked Backend

SDK provides mocks using Testify and Mockery. 
One of the SDK examples covers usage of the mockery within tests.

```bash
go test ./mock/cluster_test.go 
```

## Examples Reference

### Update Cluster Example

Example demonstrates how to safely update an Atlas cluster using minimal PATCH payloads,
covering tier resize and pause/unpause operations.

```bash
export MONGODB_ATLAS_PUBLIC_KEY=<your-public-key>
export MONGODB_ATLAS_PRIVATE_KEY=<your-private-key>
export MONGODB_ATLAS_PROJECT_ID=<your-project-id>
export MONGODB_ATLAS_CLUSTER_NAME=<your-cluster-name>
go run ./update_cluster/update_cluster.go
```

### Retry Example

Example provides automatic retries for all HTTP 500, 429 HTTP status errors.

```bash
go run ./retry/retry.go
```
