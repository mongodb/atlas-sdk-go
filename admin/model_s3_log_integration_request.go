// Code based on the AtlasAPI V2 OpenAPI file

package admin

// S3LogIntegrationRequest Request schema for creating or updating an S3 log export integration.
type S3LogIntegrationRequest struct {
	// Human-readable label that identifies the S3 bucket name for storing log files.
	BucketName string `json:"bucketName"`
	// Unique 24-hexadecimal digit string that identifies the AWS IAM role that MongoDB Cloud uses to access your S3 bucket.
	IamRoleId string `json:"iamRoleId"`
	// AWS KMS key ID or ARN for server-side encryption (optional). If not provided, uses bucket default encryption settings.
	KmsKey *string `json:"kmsKey,omitempty"`
	// Array of log types to export to S3. Valid values: MONGOD, MONGOS, MONGOD_AUDIT, MONGOS_AUDIT.
	LogTypes []string `json:"logTypes"`
	// S3 directory path prefix where the log files will be stored. MongoDB Cloud will add further sub-directories based on the log type.
	PrefixPath string `json:"prefixPath"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the log integration type.
	Type string `json:"type"`
}

// NewS3LogIntegrationRequest instantiates a new S3LogIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3LogIntegrationRequest(bucketName string, iamRoleId string, logTypes []string, prefixPath string, type_ string) *S3LogIntegrationRequest {
	this := S3LogIntegrationRequest{}
	this.Type = type_
	this.BucketName = bucketName
	this.IamRoleId = iamRoleId
	this.LogTypes = logTypes
	this.PrefixPath = prefixPath
	return &this
}

// NewS3LogIntegrationRequestWithDefaults instantiates a new S3LogIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3LogIntegrationRequestWithDefaults() *S3LogIntegrationRequest {
	this := S3LogIntegrationRequest{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *S3LogIntegrationRequest) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequest) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *S3LogIntegrationRequest) SetBucketName(v string) {
	o.BucketName = v
}

// GetIamRoleId returns the IamRoleId field value
func (o *S3LogIntegrationRequest) GetIamRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequest) GetIamRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IamRoleId, true
}

// SetIamRoleId sets field value
func (o *S3LogIntegrationRequest) SetIamRoleId(v string) {
	o.IamRoleId = v
}

// GetKmsKey returns the KmsKey field value if set, zero value otherwise
func (o *S3LogIntegrationRequest) GetKmsKey() string {
	if o == nil || IsNil(o.KmsKey) {
		var ret string
		return ret
	}
	return *o.KmsKey
}

// GetKmsKeyOk returns a tuple with the KmsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequest) GetKmsKeyOk() (*string, bool) {
	if o == nil || IsNil(o.KmsKey) {
		return nil, false
	}

	return o.KmsKey, true
}

// HasKmsKey returns a boolean if a field has been set.
func (o *S3LogIntegrationRequest) HasKmsKey() bool {
	if o != nil && !IsNil(o.KmsKey) {
		return true
	}

	return false
}

// SetKmsKey gets a reference to the given string and assigns it to the KmsKey field.
func (o *S3LogIntegrationRequest) SetKmsKey(v string) {
	o.KmsKey = &v
}

// GetLogTypes returns the LogTypes field value
func (o *S3LogIntegrationRequest) GetLogTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequest) GetLogTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTypes, true
}

// SetLogTypes sets field value
func (o *S3LogIntegrationRequest) SetLogTypes(v []string) {
	o.LogTypes = v
}

// GetPrefixPath returns the PrefixPath field value
func (o *S3LogIntegrationRequest) GetPrefixPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrefixPath
}

// GetPrefixPathOk returns a tuple with the PrefixPath field value
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequest) GetPrefixPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixPath, true
}

// SetPrefixPath sets field value
func (o *S3LogIntegrationRequest) SetPrefixPath(v string) {
	o.PrefixPath = v
}

// GetType returns the Type field value
func (o *S3LogIntegrationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *S3LogIntegrationRequest) SetType(v string) {
	o.Type = v
}
