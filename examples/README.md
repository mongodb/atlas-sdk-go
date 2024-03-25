## Go SDK examples

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
