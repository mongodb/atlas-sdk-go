// Code based on the AtlasAPI V2 OpenAPI file

package admin

// LogIntegrationRequest Request schema for creating or updating a log integration.
type LogIntegrationRequest struct {
	// Array of log types exported by this integration. The specific log types available and maximum number of items depend on the integration type. See the integration-specific schema for details.
	LogTypes *[]string `json:"logTypes,omitempty"`
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
}

// NewLogIntegrationRequest instantiates a new LogIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogIntegrationRequest(type_ string) *LogIntegrationRequest {
	this := LogIntegrationRequest{}
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

// GetLogTypes returns the LogTypes field value if set, zero value otherwise
func (o *LogIntegrationRequest) GetLogTypes() []string {
	if o == nil || IsNil(o.LogTypes) {
		var ret []string
		return ret
	}
	return *o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationRequest) GetLogTypesOk() (*[]string, bool) {
	if o == nil || IsNil(o.LogTypes) {
		return nil, false
	}

	return o.LogTypes, true
}

// HasLogTypes returns a boolean if a field has been set.
func (o *LogIntegrationRequest) HasLogTypes() bool {
	if o != nil && !IsNil(o.LogTypes) {
		return true
	}

	return false
}

// SetLogTypes gets a reference to the given []string and assigns it to the LogTypes field.
func (o *LogIntegrationRequest) SetLogTypes(v []string) {
	o.LogTypes = &v
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
