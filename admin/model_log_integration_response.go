// Code based on the AtlasAPI V2 OpenAPI file

package admin

// LogIntegrationResponse Response schema for log integration operations.
type LogIntegrationResponse struct {
	// Unique 24-character hexadecimal digit string that identifies the log integration configuration.
	Id string `json:"id"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the log integration type.
	Type string `json:"type"`
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
}

// NewLogIntegrationResponse instantiates a new LogIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogIntegrationResponse(id string, type_ string) *LogIntegrationResponse {
	this := LogIntegrationResponse{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewLogIntegrationResponseWithDefaults instantiates a new LogIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogIntegrationResponseWithDefaults() *LogIntegrationResponse {
	this := LogIntegrationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LogIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *LogIntegrationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogIntegrationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogIntegrationResponse) SetType(v string) {
	o.Type = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise
func (o *LogIntegrationResponse) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationResponse) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}

	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LogIntegrationResponse) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LogIntegrationResponse) SetBucketName(v string) {
	o.BucketName = &v
}

// GetIamRoleId returns the IamRoleId field value if set, zero value otherwise
func (o *LogIntegrationResponse) GetIamRoleId() string {
	if o == nil || IsNil(o.IamRoleId) {
		var ret string
		return ret
	}
	return *o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationResponse) GetIamRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IamRoleId) {
		return nil, false
	}

	return o.IamRoleId, true
}

// HasIamRoleId returns a boolean if a field has been set.
func (o *LogIntegrationResponse) HasIamRoleId() bool {
	if o != nil && !IsNil(o.IamRoleId) {
		return true
	}

	return false
}

// SetIamRoleId gets a reference to the given string and assigns it to the IamRoleId field.
func (o *LogIntegrationResponse) SetIamRoleId(v string) {
	o.IamRoleId = &v
}

// GetKmsKey returns the KmsKey field value if set, zero value otherwise
func (o *LogIntegrationResponse) GetKmsKey() string {
	if o == nil || IsNil(o.KmsKey) {
		var ret string
		return ret
	}
	return *o.KmsKey
}

// GetKmsKeyOk returns a tuple with the KmsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationResponse) GetKmsKeyOk() (*string, bool) {
	if o == nil || IsNil(o.KmsKey) {
		return nil, false
	}

	return o.KmsKey, true
}

// HasKmsKey returns a boolean if a field has been set.
func (o *LogIntegrationResponse) HasKmsKey() bool {
	if o != nil && !IsNil(o.KmsKey) {
		return true
	}

	return false
}

// SetKmsKey gets a reference to the given string and assigns it to the KmsKey field.
func (o *LogIntegrationResponse) SetKmsKey(v string) {
	o.KmsKey = &v
}

// GetLogTypes returns the LogTypes field value if set, zero value otherwise
func (o *LogIntegrationResponse) GetLogTypes() []string {
	if o == nil || IsNil(o.LogTypes) {
		var ret []string
		return ret
	}
	return *o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationResponse) GetLogTypesOk() (*[]string, bool) {
	if o == nil || IsNil(o.LogTypes) {
		return nil, false
	}

	return o.LogTypes, true
}

// HasLogTypes returns a boolean if a field has been set.
func (o *LogIntegrationResponse) HasLogTypes() bool {
	if o != nil && !IsNil(o.LogTypes) {
		return true
	}

	return false
}

// SetLogTypes gets a reference to the given []string and assigns it to the LogTypes field.
func (o *LogIntegrationResponse) SetLogTypes(v []string) {
	o.LogTypes = &v
}

// GetPrefixPath returns the PrefixPath field value if set, zero value otherwise
func (o *LogIntegrationResponse) GetPrefixPath() string {
	if o == nil || IsNil(o.PrefixPath) {
		var ret string
		return ret
	}
	return *o.PrefixPath
}

// GetPrefixPathOk returns a tuple with the PrefixPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIntegrationResponse) GetPrefixPathOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixPath) {
		return nil, false
	}

	return o.PrefixPath, true
}

// HasPrefixPath returns a boolean if a field has been set.
func (o *LogIntegrationResponse) HasPrefixPath() bool {
	if o != nil && !IsNil(o.PrefixPath) {
		return true
	}

	return false
}

// SetPrefixPath gets a reference to the given string and assigns it to the PrefixPath field.
func (o *LogIntegrationResponse) SetPrefixPath(v string) {
	o.PrefixPath = &v
}
