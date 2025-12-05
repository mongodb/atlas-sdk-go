// Code based on the AtlasAPI V2 OpenAPI file

package admin

// S3LogIntegrationRequestAllOf struct for S3LogIntegrationRequestAllOf
type S3LogIntegrationRequestAllOf struct {
	// Human-readable label that identifies the S3 bucket name for storing log files.
	BucketName *string `json:"bucketName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the AWS IAM role that MongoDB Cloud uses to access your S3 bucket.
	IamRoleId *string `json:"iamRoleId,omitempty"`
	// AWS KMS key ID or ARN for server-side encryption (optional). If not provided, uses bucket default encryption settings.
	KmsKey *string `json:"kmsKey,omitempty"`
	// Array of log types to export to S3. Valid values: MONGOD, MONGOS, MONGOD_AUDIT, MONGOS_AUDIT.
	LogTypes *[]string `json:"logTypes,omitempty"`
	// S3 directory path prefix where the log files will be stored. MongoDB Cloud will add further sub-directories based on the log type.
	PrefixPath *string `json:"prefixPath,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the log integration type.
	Type *string `json:"type,omitempty"`
}

// NewS3LogIntegrationRequestAllOf instantiates a new S3LogIntegrationRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3LogIntegrationRequestAllOf() *S3LogIntegrationRequestAllOf {
	this := S3LogIntegrationRequestAllOf{}
	return &this
}

// NewS3LogIntegrationRequestAllOfWithDefaults instantiates a new S3LogIntegrationRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3LogIntegrationRequestAllOfWithDefaults() *S3LogIntegrationRequestAllOf {
	this := S3LogIntegrationRequestAllOf{}
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise
func (o *S3LogIntegrationRequestAllOf) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequestAllOf) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}

	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *S3LogIntegrationRequestAllOf) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *S3LogIntegrationRequestAllOf) SetBucketName(v string) {
	o.BucketName = &v
}

// GetIamRoleId returns the IamRoleId field value if set, zero value otherwise
func (o *S3LogIntegrationRequestAllOf) GetIamRoleId() string {
	if o == nil || IsNil(o.IamRoleId) {
		var ret string
		return ret
	}
	return *o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequestAllOf) GetIamRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IamRoleId) {
		return nil, false
	}

	return o.IamRoleId, true
}

// HasIamRoleId returns a boolean if a field has been set.
func (o *S3LogIntegrationRequestAllOf) HasIamRoleId() bool {
	if o != nil && !IsNil(o.IamRoleId) {
		return true
	}

	return false
}

// SetIamRoleId gets a reference to the given string and assigns it to the IamRoleId field.
func (o *S3LogIntegrationRequestAllOf) SetIamRoleId(v string) {
	o.IamRoleId = &v
}

// GetKmsKey returns the KmsKey field value if set, zero value otherwise
func (o *S3LogIntegrationRequestAllOf) GetKmsKey() string {
	if o == nil || IsNil(o.KmsKey) {
		var ret string
		return ret
	}
	return *o.KmsKey
}

// GetKmsKeyOk returns a tuple with the KmsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequestAllOf) GetKmsKeyOk() (*string, bool) {
	if o == nil || IsNil(o.KmsKey) {
		return nil, false
	}

	return o.KmsKey, true
}

// HasKmsKey returns a boolean if a field has been set.
func (o *S3LogIntegrationRequestAllOf) HasKmsKey() bool {
	if o != nil && !IsNil(o.KmsKey) {
		return true
	}

	return false
}

// SetKmsKey gets a reference to the given string and assigns it to the KmsKey field.
func (o *S3LogIntegrationRequestAllOf) SetKmsKey(v string) {
	o.KmsKey = &v
}

// GetLogTypes returns the LogTypes field value if set, zero value otherwise
func (o *S3LogIntegrationRequestAllOf) GetLogTypes() []string {
	if o == nil || IsNil(o.LogTypes) {
		var ret []string
		return ret
	}
	return *o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequestAllOf) GetLogTypesOk() (*[]string, bool) {
	if o == nil || IsNil(o.LogTypes) {
		return nil, false
	}

	return o.LogTypes, true
}

// HasLogTypes returns a boolean if a field has been set.
func (o *S3LogIntegrationRequestAllOf) HasLogTypes() bool {
	if o != nil && !IsNil(o.LogTypes) {
		return true
	}

	return false
}

// SetLogTypes gets a reference to the given []string and assigns it to the LogTypes field.
func (o *S3LogIntegrationRequestAllOf) SetLogTypes(v []string) {
	o.LogTypes = &v
}

// GetPrefixPath returns the PrefixPath field value if set, zero value otherwise
func (o *S3LogIntegrationRequestAllOf) GetPrefixPath() string {
	if o == nil || IsNil(o.PrefixPath) {
		var ret string
		return ret
	}
	return *o.PrefixPath
}

// GetPrefixPathOk returns a tuple with the PrefixPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequestAllOf) GetPrefixPathOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixPath) {
		return nil, false
	}

	return o.PrefixPath, true
}

// HasPrefixPath returns a boolean if a field has been set.
func (o *S3LogIntegrationRequestAllOf) HasPrefixPath() bool {
	if o != nil && !IsNil(o.PrefixPath) {
		return true
	}

	return false
}

// SetPrefixPath gets a reference to the given string and assigns it to the PrefixPath field.
func (o *S3LogIntegrationRequestAllOf) SetPrefixPath(v string) {
	o.PrefixPath = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *S3LogIntegrationRequestAllOf) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LogIntegrationRequestAllOf) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *S3LogIntegrationRequestAllOf) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *S3LogIntegrationRequestAllOf) SetType(v string) {
	o.Type = &v
}
