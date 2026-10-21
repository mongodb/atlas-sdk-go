# Go SDK examples

## Running Examples

For documentation about obtaining apiKey and apiSecret go to
https://docs.atlas.mongodb.com/configure-api-access.

```bash
export MONGODB_ATLAS_PUBLIC_KEY=somekey 
export MONGODB_ATLAS_PRIVATE_KEY=some-secret-key-for-gosdkapi
go run ./cluster/aws_cluster/aws.go
```

## Running Examples with Mocked Backend

SDK provides mocks using Testify and Mockery. 
One of the SDK examples covers usage of the mockery within tests.

```bash
go test ./mock/cluster_test.go 
```

## Examples Reference

### Authorization Code + PKCE Example

This example demonstrates browser-based user authentication and uses the
returned `code.Token` with the Atlas Admin API client.
Register `/my-app/callback` as a redirect path for the OAuth client before
running it.

```bash
export MONGODB_ATLAS_CLIENT_ID=<your-client-id>
export MONGODB_ATLAS_CALLBACK_PATH='/my-app/callback'
go run ./auth_code/auth_code.go
```

### Update Cluster Example

This example demonstrates how to safely change an Atlas cluster tier using a minimal PATCH payload.

```bash
export MONGODB_ATLAS_CLIENT_ID=<your-client-id>
export MONGODB_ATLAS_CLIENT_SECRET=<your-client-secret>
export MONGODB_ATLAS_PROJECT_ID=<your-project-id>
export MONGODB_ATLAS_CLUSTER_NAME=<your-cluster-name>
go run ./cluster/update_cluster/update_cluster.go
```

### Pause Cluster Example

This example demonstrates how to pause and unpause an Atlas cluster using a minimal PATCH payload.

```bash
export MONGODB_ATLAS_CLIENT_ID=<your-client-id>
export MONGODB_ATLAS_CLIENT_SECRET=<your-client-secret>
export MONGODB_ATLAS_PROJECT_ID=<your-project-id>
export MONGODB_ATLAS_CLUSTER_NAME=<your-cluster-name>
go run ./cluster/pause_cluster/pause_cluster.go
```

### Retry Example

This example provides automatic retries for all HTTP 500, 429 HTTP status errors.

```bash
go run ./retry/retry.go
```
