module github.com/mongodb/atlas-sdk-go/examples

go 1.24.0

replace github.com/mongodb/atlas-sdk-go => ../

require (
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/mongodb-forks/digest v1.1.0
)

require github.com/mongodb/atlas-sdk-go v1.0.0 // indirect

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.11.1
	golang.org/x/oauth2 v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
