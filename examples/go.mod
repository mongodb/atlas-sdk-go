module go.mongodb.org/atlas-sdk/v20240805005/examples

go 1.22.0

toolchain go1.23.1

replace go.mongodb.org/atlas-sdk/v20240805005 => ../
replace go.mongodb.org/atlas-sdk/v20240805005/auth/credentials => ../auth/credentials

require (
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/mongodb-forks/digest v1.1.0
	go.mongodb.org/atlas-sdk/v20240805005 v20240805005.1.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.9.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
