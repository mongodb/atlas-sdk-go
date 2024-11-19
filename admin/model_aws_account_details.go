// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AWSAccountDetails struct for AWSAccountDetails
type AWSAccountDetails struct {
	// The AWS Account ID.
	AwsAccountId *string `json:"awsAccountId,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// The VPC ID.
	VpcId *string `json:"vpcId,omitempty"`
}

// NewAWSAccountDetails instantiates a new AWSAccountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSAccountDetails() *AWSAccountDetails {
	this := AWSAccountDetails{}
	return &this
}

// NewAWSAccountDetailsWithDefaults instantiates a new AWSAccountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSAccountDetailsWithDefaults() *AWSAccountDetails {
	this := AWSAccountDetails{}
	return &this
}

// GetAwsAccountId returns the AwsAccountId field value if set, zero value otherwise
func (o *AWSAccountDetails) GetAwsAccountId() string {
	if o == nil || IsNil(o.AwsAccountId) {
		var ret string
		return ret
	}
	return *o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountDetails) GetAwsAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountId) {
		return nil, false
	}

	return o.AwsAccountId, true
}

// HasAwsAccountId returns a boolean if a field has been set.
func (o *AWSAccountDetails) HasAwsAccountId() bool {
	if o != nil && !IsNil(o.AwsAccountId) {
		return true
	}

	return false
}

// SetAwsAccountId gets a reference to the given string and assigns it to the AwsAccountId field.
func (o *AWSAccountDetails) SetAwsAccountId(v string) {
	o.AwsAccountId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *AWSAccountDetails) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountDetails) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AWSAccountDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *AWSAccountDetails) SetLinks(v []Link) {
	o.Links = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise
func (o *AWSAccountDetails) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountDetails) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}

	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *AWSAccountDetails) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *AWSAccountDetails) SetVpcId(v string) {
	o.VpcId = &v
}
