// Code based on the AtlasAPI V2 OpenAPI file

package admin

// LogIntegrationRequest Request schema for creating or updating a log integration.
type LogIntegrationRequest struct {
	// Array of log types exported by this integration. The specific log types available and maximum number of items depend on the integration type. See the integration-specific schema for details.
	LogTypes []string `json:"logTypes"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the log integration type.
	Type string `json:"type"`
	// Human-readable label that identifies the S3 bucket name for storing log files.
	BucketName *string `json:"bucketName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the AWS IAM role that MongoDB Cloud uses to access your S3 bucket.
	IamRoleId *string `json:"iamRoleId,omitempty"`
	// AWS KMS key ID or ARN for server-side encryption (optional). If not provided, uses bucket default encryption settings.
	KmsKey *string `json:"kmsKey,omitempty"`
	// S3 directory path prefix where the log files will be stored. MongoDB Cloud will add further sub-directories based on the log type.
	PrefixPath *string `json:"prefixPath,omitempty"`
	// Datadog API key for authentication.
	ApiKey *string `json:"apiKey,omitempty"`
	// Datadog site/region for log ingestion. Valid values: US1, US3, US5, EU, AP1, AP2, US1_FED.
	Region *string `json:"region,omitempty"`
	// GCS bucket name for storing log files.
	GcsBucketName *string `json:"gcsBucketName,omitempty"`
	// GCS object path prefix for organizing log files.
	GcsPrefixPath *string `json:"gcsPrefixPath,omitempty"`
	// Unique 24-character hexadecimal string that identifies the GCP service account role.
	GcsRoleId *string `json:"gcsRoleId,omitempty"`
	// OpenTelemetry collector endpoint URL. Must be HTTPS and not exceed 2048 characters.
	OtelEndpoint *string `json:"otelEndpoint,omitempty"`
	// HTTP headers for authentication and configuration. Maximum 10 headers, total size limit 2KB.
	OtelSuppliedHeaders *[]Header `json:"otelSuppliedHeaders,omitempty"`
	// Splunk HTTP Event Collector (HEC) token for authentication.
	SplunkHecToken *string `json:"splunkHecToken,omitempty"`
	// Splunk HTTP Event Collector (HEC) endpoint URL.
	SplunkHecUrl *string `json:"splunkHecUrl,omitempty"`
	// Azure Blob Storage container name for log files.
	AzureContainer *string `json:"azureContainer,omitempty"`
	// Blob path prefix for organizing log files within the container.
	AzurePrefixPath *string `json:"azurePrefixPath,omitempty"`
	// Unique 24-character hexadecimal string that identifies the Azure Service Principal.
	AzureServicePrincipalId *string `json:"azureServicePrincipalId,omitempty"`
	// Azure Storage Account name where logs will be stored.
	AzureStorageAccountName *string `json:"azureStorageAccountName,omitempty"`
}

// NewLogIntegrationRequest instantiates a new LogIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogIntegrationRequest(logTypes []string, type_ string) *LogIntegrationRequest {
	this := LogIntegrationRequest{}
	this.LogTypes = logTypes
	this.Type = type_
	return &this
}

// NewLogIntegrationRequestWithDefaults instantiates a new LogIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogIntegrationRequestWithDefaults() *LogIntegrationRequest {
	this := LogIntegrationRequest{}
	return &this
}

// GetLogTypes returns the LogTypes field value
func (o *LogIntegrationRequest) GetLogTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetLogTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTypes, true
}

// SetLogTypes sets field value
func (o *LogIntegrationRequest) SetLogTypes(v []string) {
	o.LogTypes = v
}

// GetType returns the Type field value
func (o *LogIntegrationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogIntegrationRequest) SetType(v string) {
	o.Type = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}

	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LogIntegrationRequest) SetBucketName(v string) {
	o.BucketName = &v
}

// GetIamRoleId returns the IamRoleId field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetIamRoleId() string {
	if o == nil || IsNil(o.IamRoleId) {
		var ret string
		return ret
	}
	return *o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetIamRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IamRoleId) {
		return nil, false
	}

	return o.IamRoleId, true
}

// HasIamRoleId returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasIamRoleId() bool {
	if o != nil && !IsNil(o.IamRoleId) {
		return true
	}

	return false
}

// SetIamRoleId gets a reference to the given string and assigns it to the IamRoleId field.
func (o *LogIntegrationRequest) SetIamRoleId(v string) {
	o.IamRoleId = &v
}

// GetKmsKey returns the KmsKey field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetKmsKey() string {
	if o == nil || IsNil(o.KmsKey) {
		var ret string
		return ret
	}
	return *o.KmsKey
}

// GetKmsKeyOk returns a tuple with the KmsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetKmsKeyOk() (*string, bool) {
	if o == nil || IsNil(o.KmsKey) {
		return nil, false
	}

	return o.KmsKey, true
}

// HasKmsKey returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasKmsKey() bool {
	if o != nil && !IsNil(o.KmsKey) {
		return true
	}

	return false
}

// SetKmsKey gets a reference to the given string and assigns it to the KmsKey field.
func (o *LogIntegrationRequest) SetKmsKey(v string) {
	o.KmsKey = &v
}

// GetPrefixPath returns the PrefixPath field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetPrefixPath() string {
	if o == nil || IsNil(o.PrefixPath) {
		var ret string
		return ret
	}
	return *o.PrefixPath
}

// GetPrefixPathOk returns a tuple with the PrefixPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetPrefixPathOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixPath) {
		return nil, false
	}

	return o.PrefixPath, true
}

// HasPrefixPath returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasPrefixPath() bool {
	if o != nil && !IsNil(o.PrefixPath) {
		return true
	}

	return false
}

// SetPrefixPath gets a reference to the given string and assigns it to the PrefixPath field.
func (o *LogIntegrationRequest) SetPrefixPath(v string) {
	o.PrefixPath = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}

	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *LogIntegrationRequest) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *LogIntegrationRequest) SetRegion(v string) {
	o.Region = &v
}

// GetGcsBucketName returns the GcsBucketName field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetGcsBucketName() string {
	if o == nil || IsNil(o.GcsBucketName) {
		var ret string
		return ret
	}
	return *o.GcsBucketName
}

// GetGcsBucketNameOk returns a tuple with the GcsBucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetGcsBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.GcsBucketName) {
		return nil, false
	}

	return o.GcsBucketName, true
}

// HasGcsBucketName returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasGcsBucketName() bool {
	if o != nil && !IsNil(o.GcsBucketName) {
		return true
	}

	return false
}

// SetGcsBucketName gets a reference to the given string and assigns it to the GcsBucketName field.
func (o *LogIntegrationRequest) SetGcsBucketName(v string) {
	o.GcsBucketName = &v
}

// GetGcsPrefixPath returns the GcsPrefixPath field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetGcsPrefixPath() string {
	if o == nil || IsNil(o.GcsPrefixPath) {
		var ret string
		return ret
	}
	return *o.GcsPrefixPath
}

// GetGcsPrefixPathOk returns a tuple with the GcsPrefixPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetGcsPrefixPathOk() (*string, bool) {
	if o == nil || IsNil(o.GcsPrefixPath) {
		return nil, false
	}

	return o.GcsPrefixPath, true
}

// HasGcsPrefixPath returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasGcsPrefixPath() bool {
	if o != nil && !IsNil(o.GcsPrefixPath) {
		return true
	}

	return false
}

// SetGcsPrefixPath gets a reference to the given string and assigns it to the GcsPrefixPath field.
func (o *LogIntegrationRequest) SetGcsPrefixPath(v string) {
	o.GcsPrefixPath = &v
}

// GetGcsRoleId returns the GcsRoleId field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetGcsRoleId() string {
	if o == nil || IsNil(o.GcsRoleId) {
		var ret string
		return ret
	}
	return *o.GcsRoleId
}

// GetGcsRoleIdOk returns a tuple with the GcsRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetGcsRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.GcsRoleId) {
		return nil, false
	}

	return o.GcsRoleId, true
}

// HasGcsRoleId returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasGcsRoleId() bool {
	if o != nil && !IsNil(o.GcsRoleId) {
		return true
	}

	return false
}

// SetGcsRoleId gets a reference to the given string and assigns it to the GcsRoleId field.
func (o *LogIntegrationRequest) SetGcsRoleId(v string) {
	o.GcsRoleId = &v
}

// GetOtelEndpoint returns the OtelEndpoint field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetOtelEndpoint() string {
	if o == nil || IsNil(o.OtelEndpoint) {
		var ret string
		return ret
	}
	return *o.OtelEndpoint
}

// GetOtelEndpointOk returns a tuple with the OtelEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetOtelEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.OtelEndpoint) {
		return nil, false
	}

	return o.OtelEndpoint, true
}

// HasOtelEndpoint returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasOtelEndpoint() bool {
	if o != nil && !IsNil(o.OtelEndpoint) {
		return true
	}

	return false
}

// SetOtelEndpoint gets a reference to the given string and assigns it to the OtelEndpoint field.
func (o *LogIntegrationRequest) SetOtelEndpoint(v string) {
	o.OtelEndpoint = &v
}

// GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetOtelSuppliedHeaders() []Header {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		var ret []Header
		return ret
	}
	return *o.OtelSuppliedHeaders
}

// GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetOtelSuppliedHeadersOk() (*[]Header, bool) {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		return nil, false
	}

	return o.OtelSuppliedHeaders, true
}

// HasOtelSuppliedHeaders returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasOtelSuppliedHeaders() bool {
	if o != nil && !IsNil(o.OtelSuppliedHeaders) {
		return true
	}

	return false
}

// SetOtelSuppliedHeaders gets a reference to the given []Header and assigns it to the OtelSuppliedHeaders field.
func (o *LogIntegrationRequest) SetOtelSuppliedHeaders(v []Header) {
	o.OtelSuppliedHeaders = &v
}

// GetSplunkHecToken returns the SplunkHecToken field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetSplunkHecToken() string {
	if o == nil || IsNil(o.SplunkHecToken) {
		var ret string
		return ret
	}
	return *o.SplunkHecToken
}

// GetSplunkHecTokenOk returns a tuple with the SplunkHecToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetSplunkHecTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SplunkHecToken) {
		return nil, false
	}

	return o.SplunkHecToken, true
}

// HasSplunkHecToken returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasSplunkHecToken() bool {
	if o != nil && !IsNil(o.SplunkHecToken) {
		return true
	}

	return false
}

// SetSplunkHecToken gets a reference to the given string and assigns it to the SplunkHecToken field.
func (o *LogIntegrationRequest) SetSplunkHecToken(v string) {
	o.SplunkHecToken = &v
}

// GetSplunkHecUrl returns the SplunkHecUrl field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetSplunkHecUrl() string {
	if o == nil || IsNil(o.SplunkHecUrl) {
		var ret string
		return ret
	}
	return *o.SplunkHecUrl
}

// GetSplunkHecUrlOk returns a tuple with the SplunkHecUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetSplunkHecUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SplunkHecUrl) {
		return nil, false
	}

	return o.SplunkHecUrl, true
}

// HasSplunkHecUrl returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasSplunkHecUrl() bool {
	if o != nil && !IsNil(o.SplunkHecUrl) {
		return true
	}

	return false
}

// SetSplunkHecUrl gets a reference to the given string and assigns it to the SplunkHecUrl field.
func (o *LogIntegrationRequest) SetSplunkHecUrl(v string) {
	o.SplunkHecUrl = &v
}

// GetAzureContainer returns the AzureContainer field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetAzureContainer() string {
	if o == nil || IsNil(o.AzureContainer) {
		var ret string
		return ret
	}
	return *o.AzureContainer
}

// GetAzureContainerOk returns a tuple with the AzureContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetAzureContainerOk() (*string, bool) {
	if o == nil || IsNil(o.AzureContainer) {
		return nil, false
	}

	return o.AzureContainer, true
}

// HasAzureContainer returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasAzureContainer() bool {
	if o != nil && !IsNil(o.AzureContainer) {
		return true
	}

	return false
}

// SetAzureContainer gets a reference to the given string and assigns it to the AzureContainer field.
func (o *LogIntegrationRequest) SetAzureContainer(v string) {
	o.AzureContainer = &v
}

// GetAzurePrefixPath returns the AzurePrefixPath field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetAzurePrefixPath() string {
	if o == nil || IsNil(o.AzurePrefixPath) {
		var ret string
		return ret
	}
	return *o.AzurePrefixPath
}

// GetAzurePrefixPathOk returns a tuple with the AzurePrefixPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetAzurePrefixPathOk() (*string, bool) {
	if o == nil || IsNil(o.AzurePrefixPath) {
		return nil, false
	}

	return o.AzurePrefixPath, true
}

// HasAzurePrefixPath returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasAzurePrefixPath() bool {
	if o != nil && !IsNil(o.AzurePrefixPath) {
		return true
	}

	return false
}

// SetAzurePrefixPath gets a reference to the given string and assigns it to the AzurePrefixPath field.
func (o *LogIntegrationRequest) SetAzurePrefixPath(v string) {
	o.AzurePrefixPath = &v
}

// GetAzureServicePrincipalId returns the AzureServicePrincipalId field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetAzureServicePrincipalId() string {
	if o == nil || IsNil(o.AzureServicePrincipalId) {
		var ret string
		return ret
	}
	return *o.AzureServicePrincipalId
}

// GetAzureServicePrincipalIdOk returns a tuple with the AzureServicePrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetAzureServicePrincipalIdOk() (*string, bool) {
	if o == nil || IsNil(o.AzureServicePrincipalId) {
		return nil, false
	}

	return o.AzureServicePrincipalId, true
}

// HasAzureServicePrincipalId returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasAzureServicePrincipalId() bool {
	if o != nil && !IsNil(o.AzureServicePrincipalId) {
		return true
	}

	return false
}

// SetAzureServicePrincipalId gets a reference to the given string and assigns it to the AzureServicePrincipalId field.
func (o *LogIntegrationRequest) SetAzureServicePrincipalId(v string) {
	o.AzureServicePrincipalId = &v
}

// GetAzureStorageAccountName returns the AzureStorageAccountName field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetAzureStorageAccountName() string {
	if o == nil || IsNil(o.AzureStorageAccountName) {
		var ret string
		return ret
	}
	return *o.AzureStorageAccountName
}

// GetAzureStorageAccountNameOk returns a tuple with the AzureStorageAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetAzureStorageAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AzureStorageAccountName) {
		return nil, false
	}

	return o.AzureStorageAccountName, true
}

// HasAzureStorageAccountName returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasAzureStorageAccountName() bool {
	if o != nil && !IsNil(o.AzureStorageAccountName) {
		return true
	}

	return false
}

// SetAzureStorageAccountName gets a reference to the given string and assigns it to the AzureStorageAccountName field.
func (o *LogIntegrationRequest) SetAzureStorageAccountName(v string) {
	o.AzureStorageAccountName = &v
}
