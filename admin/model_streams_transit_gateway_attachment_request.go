// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsTransitGatewayAttachmentRequest struct for StreamsTransitGatewayAttachmentRequest
type StreamsTransitGatewayAttachmentRequest struct {
	// Provider for the transit gateway resources.
	// Write only field.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// AWS region name.
	// Write only field.
	RegionName *string `json:"regionName,omitempty"`
	// AWS transit gateway ID.
	TgwId *string `json:"tgwId,omitempty"`
	// AWS VPC ID.
	// Write only field.
	VpcId *string `json:"vpcId,omitempty"`
}

// NewStreamsTransitGatewayAttachmentRequest instantiates a new StreamsTransitGatewayAttachmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsTransitGatewayAttachmentRequest() *StreamsTransitGatewayAttachmentRequest {
	this := StreamsTransitGatewayAttachmentRequest{}
	return &this
}

// NewStreamsTransitGatewayAttachmentRequestWithDefaults instantiates a new StreamsTransitGatewayAttachmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsTransitGatewayAttachmentRequestWithDefaults() *StreamsTransitGatewayAttachmentRequest {
	this := StreamsTransitGatewayAttachmentRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *StreamsTransitGatewayAttachmentRequest) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayAttachmentRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *StreamsTransitGatewayAttachmentRequest) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *StreamsTransitGatewayAttachmentRequest) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *StreamsTransitGatewayAttachmentRequest) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayAttachmentRequest) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *StreamsTransitGatewayAttachmentRequest) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *StreamsTransitGatewayAttachmentRequest) SetRegionName(v string) {
	o.RegionName = &v
}

// GetTgwId returns the TgwId field value if set, zero value otherwise
func (o *StreamsTransitGatewayAttachmentRequest) GetTgwId() string {
	if o == nil || IsNil(o.TgwId) {
		var ret string
		return ret
	}
	return *o.TgwId
}

// GetTgwIdOk returns a tuple with the TgwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayAttachmentRequest) GetTgwIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwId) {
		return nil, false
	}

	return o.TgwId, true
}

// HasTgwId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayAttachmentRequest) HasTgwId() bool {
	if o != nil && !IsNil(o.TgwId) {
		return true
	}

	return false
}

// SetTgwId gets a reference to the given string and assigns it to the TgwId field.
func (o *StreamsTransitGatewayAttachmentRequest) SetTgwId(v string) {
	o.TgwId = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise
func (o *StreamsTransitGatewayAttachmentRequest) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayAttachmentRequest) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}

	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayAttachmentRequest) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *StreamsTransitGatewayAttachmentRequest) SetVpcId(v string) {
	o.VpcId = &v
}
