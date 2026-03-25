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

Demonstrates the recommended approach for updating an existing Atlas cluster
using the PATCH API, covering:

- **Changing the cluster tier (instance size)**:
  1. GET the cluster
  2. Copy only `ReplicationSpecs` into a minimal `ClusterDescription20240805`
  3. Update the instance size
  4. PATCH with that minimal payload

- **Pausing and unpausing a cluster**: 
  1. Set the `Paused` root-level field in a minimal `ClusterDescription20240805`
  2. PATCH with that minimal payload.

Key concepts documented in the example:

- The PATCH endpoint operates on **root-level fields only**. Include only the
  root fields you want to change; omit everything else.
- You must **not** include read-only root fields (`ConnectionStrings`, `CreateDate`, `Id`,
  `MongoDBVersion`, `StateName`, `GroupId`)  in the PATCH payload or the API will return HTTP 400.
- The following fields within `ReplicationSpecs`, are read-only: `ZoneId`, `StateName`, and `CreateDate`. The API silently ignores them, but it is best practice not to
  include them.

To run the example against a real Atlas cluster, set the following credentials:

```bash
export MONGODB_ATLAS_PUBLIC_KEY=somekey
export MONGODB_ATLAS_PRIVATE_KEY=some-secret-key-for-gosdkapi
export MONGODB_ATLAS_PROJECT_ID=<your-project-id>
export MONGODB_ATLAS_CLUSTER_NAME=<your-cluster-name>
go run ./update_cluster/update_cluster.go
```

To run the mock-based tests for this example without a real Atlas cluster:

```bash
go test ./mock/ -run TestChangeClusterTier -run TestSetClusterPaused -run TestSetClusterUnpaused -v
```

### Retry Example

Example provides automatic retries for all HTTP 500, 429 HTTP status errors.

```bash
go run ./retry/retry.go
```
