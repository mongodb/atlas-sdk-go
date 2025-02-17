// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsAWSConnectionBaseConfig The base configuration of AWS connection used in stream processors.
type StreamsAWSConnectionBaseConfig struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account.
	RoleArn *string `json:"roleArn,omitempty"`
}

// NewStreamsAWSConnectionBaseConfig instantiates a new StreamsAWSConnectionBaseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsAWSConnectionBaseConfig() *StreamsAWSConnectionBaseConfig {
	this := StreamsAWSConnectionBaseConfig{}
	return &this
}

// NewStreamsAWSConnectionBaseConfigWithDefaults instantiates a new StreamsAWSConnectionBaseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsAWSConnectionBaseConfigWithDefaults() *StreamsAWSConnectionBaseConfig {
	this := StreamsAWSConnectionBaseConfig{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsAWSConnectionBaseConfig) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsAWSConnectionBaseConfig) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsAWSConnectionBaseConfig) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsAWSConnectionBaseConfig) SetLinks(v []Link) {
	o.Links = &v
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise
func (o *StreamsAWSConnectionBaseConfig) GetRoleArn() string {
	if o == nil || IsNil(o.RoleArn) {
		var ret string
		return ret
	}
	return *o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsAWSConnectionBaseConfig) GetRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.RoleArn) {
		return nil, false
	}

	return o.RoleArn, true
}

// HasRoleArn returns a boolean if a field has been set.
func (o *StreamsAWSConnectionBaseConfig) HasRoleArn() bool {
	if o != nil && !IsNil(o.RoleArn) {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given string and assigns it to the RoleArn field.
func (o *StreamsAWSConnectionBaseConfig) SetRoleArn(v string) {
	o.RoleArn = &v
}
