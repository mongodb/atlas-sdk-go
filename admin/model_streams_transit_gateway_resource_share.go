// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsTransitGatewayResourceShare struct for StreamsTransitGatewayResourceShare
type StreamsTransitGatewayResourceShare struct {
	// Provider for the transit gateway resources.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// AWS Region name.
	RegionName *string `json:"regionName,omitempty"`
	// AWS transit gateway ID.
	TgwId *string `json:"tgwId,omitempty"`
	// AWS Transit Gateway resource share ARN.
	TgwResourceShareArn *string `json:"tgwResourceShareArn,omitempty"`
	// AWS Transit Gateway resource share invitation ARN.
	// Read only field.
	TgwResourceShareInvitationArn *string `json:"tgwResourceShareInvitationArn,omitempty"`
}

// NewStreamsTransitGatewayResourceShare instantiates a new StreamsTransitGatewayResourceShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsTransitGatewayResourceShare() *StreamsTransitGatewayResourceShare {
	this := StreamsTransitGatewayResourceShare{}
	return &this
}

// NewStreamsTransitGatewayResourceShareWithDefaults instantiates a new StreamsTransitGatewayResourceShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsTransitGatewayResourceShareWithDefaults() *StreamsTransitGatewayResourceShare {
	this := StreamsTransitGatewayResourceShare{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *StreamsTransitGatewayResourceShare) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayResourceShare) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *StreamsTransitGatewayResourceShare) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *StreamsTransitGatewayResourceShare) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsTransitGatewayResourceShare) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayResourceShare) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsTransitGatewayResourceShare) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsTransitGatewayResourceShare) SetLinks(v []Link) {
	o.Links = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *StreamsTransitGatewayResourceShare) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayResourceShare) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *StreamsTransitGatewayResourceShare) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *StreamsTransitGatewayResourceShare) SetRegionName(v string) {
	o.RegionName = &v
}

// GetTgwId returns the TgwId field value if set, zero value otherwise
func (o *StreamsTransitGatewayResourceShare) GetTgwId() string {
	if o == nil || IsNil(o.TgwId) {
		var ret string
		return ret
	}
	return *o.TgwId
}

// GetTgwIdOk returns a tuple with the TgwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayResourceShare) GetTgwIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwId) {
		return nil, false
	}

	return o.TgwId, true
}

// HasTgwId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayResourceShare) HasTgwId() bool {
	if o != nil && !IsNil(o.TgwId) {
		return true
	}

	return false
}

// SetTgwId gets a reference to the given string and assigns it to the TgwId field.
func (o *StreamsTransitGatewayResourceShare) SetTgwId(v string) {
	o.TgwId = &v
}

// GetTgwResourceShareArn returns the TgwResourceShareArn field value if set, zero value otherwise
func (o *StreamsTransitGatewayResourceShare) GetTgwResourceShareArn() string {
	if o == nil || IsNil(o.TgwResourceShareArn) {
		var ret string
		return ret
	}
	return *o.TgwResourceShareArn
}

// GetTgwResourceShareArnOk returns a tuple with the TgwResourceShareArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayResourceShare) GetTgwResourceShareArnOk() (*string, bool) {
	if o == nil || IsNil(o.TgwResourceShareArn) {
		return nil, false
	}

	return o.TgwResourceShareArn, true
}

// HasTgwResourceShareArn returns a boolean if a field has been set.
func (o *StreamsTransitGatewayResourceShare) HasTgwResourceShareArn() bool {
	if o != nil && !IsNil(o.TgwResourceShareArn) {
		return true
	}

	return false
}

// SetTgwResourceShareArn gets a reference to the given string and assigns it to the TgwResourceShareArn field.
func (o *StreamsTransitGatewayResourceShare) SetTgwResourceShareArn(v string) {
	o.TgwResourceShareArn = &v
}

// GetTgwResourceShareInvitationArn returns the TgwResourceShareInvitationArn field value if set, zero value otherwise
func (o *StreamsTransitGatewayResourceShare) GetTgwResourceShareInvitationArn() string {
	if o == nil || IsNil(o.TgwResourceShareInvitationArn) {
		var ret string
		return ret
	}
	return *o.TgwResourceShareInvitationArn
}

// GetTgwResourceShareInvitationArnOk returns a tuple with the TgwResourceShareInvitationArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayResourceShare) GetTgwResourceShareInvitationArnOk() (*string, bool) {
	if o == nil || IsNil(o.TgwResourceShareInvitationArn) {
		return nil, false
	}

	return o.TgwResourceShareInvitationArn, true
}

// HasTgwResourceShareInvitationArn returns a boolean if a field has been set.
func (o *StreamsTransitGatewayResourceShare) HasTgwResourceShareInvitationArn() bool {
	if o != nil && !IsNil(o.TgwResourceShareInvitationArn) {
		return true
	}

	return false
}

// SetTgwResourceShareInvitationArn gets a reference to the given string and assigns it to the TgwResourceShareInvitationArn field.
func (o *StreamsTransitGatewayResourceShare) SetTgwResourceShareInvitationArn(v string) {
	o.TgwResourceShareInvitationArn = &v
}
