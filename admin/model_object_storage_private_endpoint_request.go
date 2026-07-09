// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ObjectStoragePrivateEndpointRequest struct for ObjectStoragePrivateEndpointRequest
type ObjectStoragePrivateEndpointRequest struct {
	// Human-readable label that identifies the cloud provider.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Cloud provider region of the S3 bucket that the Object Storage private endpoint accesses. For same-region endpoints, this is also the region where the VPC interface endpoint is deployed.
	RegionName *string `json:"regionName,omitempty"`
	// Cloud provider region in which the VPC interface endpoint is deployed. Omit to deploy the interface endpoint in the same region as the S3 bucket (same-region endpoint). Set to a region different from `regionName` to create a cross-region endpoint.
	VpcRegionName *string `json:"vpcRegionName,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ObjectStoragePrivateEndpointRequest) MarshalJSON() ([]byte, error) {
	type noMethod ObjectStoragePrivateEndpointRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewObjectStoragePrivateEndpointRequest instantiates a new ObjectStoragePrivateEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePrivateEndpointRequest() *ObjectStoragePrivateEndpointRequest {
	this := ObjectStoragePrivateEndpointRequest{}
	return &this
}

// NewObjectStoragePrivateEndpointRequestWithDefaults instantiates a new ObjectStoragePrivateEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePrivateEndpointRequestWithDefaults() *ObjectStoragePrivateEndpointRequest {
	this := ObjectStoragePrivateEndpointRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ObjectStoragePrivateEndpointRequest) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetCloudProviderNil() {
	o.CloudProvider = nil
	o.NullFields = append(o.NullFields, "CloudProvider")
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ObjectStoragePrivateEndpointRequest) SetRegionName(v string) {
	o.RegionName = &v
}

// SetRegionNameNil sets RegionName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetRegionNameNil() {
	o.RegionName = nil
	o.NullFields = append(o.NullFields, "RegionName")
}

// GetVpcRegionName returns the VpcRegionName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetVpcRegionName() string {
	if o == nil || IsNil(o.VpcRegionName) {
		var ret string
		return ret
	}
	return *o.VpcRegionName
}

// GetVpcRegionNameOk returns a tuple with the VpcRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetVpcRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.VpcRegionName) {
		return nil, false
	}

	return o.VpcRegionName, true
}

// HasVpcRegionName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasVpcRegionName() bool {
	if o != nil && !IsNil(o.VpcRegionName) {
		return true
	}

	return false
}

// SetVpcRegionName gets a reference to the given string and assigns it to the VpcRegionName field.
func (o *ObjectStoragePrivateEndpointRequest) SetVpcRegionName(v string) {
	o.VpcRegionName = &v
}

// SetVpcRegionNameNil sets VpcRegionName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetVpcRegionNameNil() {
	o.VpcRegionName = nil
	o.NullFields = append(o.NullFields, "VpcRegionName")
}
