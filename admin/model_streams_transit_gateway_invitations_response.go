// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsTransitGatewayInvitationsResponse struct for StreamsTransitGatewayInvitationsResponse
type StreamsTransitGatewayInvitationsResponse struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// AWS transit gateway ID.
	TgwId *string `json:"tgwId,omitempty"`
	// AWS Transit Gateway resource share ARN.
	TgwResourceShareArn *string `json:"tgwResourceShareArn,omitempty"`
	// AWS Transit Gateway resource share invitation ARN.
	TgwResourceShareInvitationArn *string `json:"tgwResourceShareInvitationArn,omitempty"`
}

// NewStreamsTransitGatewayInvitationsResponse instantiates a new StreamsTransitGatewayInvitationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsTransitGatewayInvitationsResponse() *StreamsTransitGatewayInvitationsResponse {
	this := StreamsTransitGatewayInvitationsResponse{}
	return &this
}

// NewStreamsTransitGatewayInvitationsResponseWithDefaults instantiates a new StreamsTransitGatewayInvitationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsTransitGatewayInvitationsResponseWithDefaults() *StreamsTransitGatewayInvitationsResponse {
	this := StreamsTransitGatewayInvitationsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsTransitGatewayInvitationsResponse) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayInvitationsResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsTransitGatewayInvitationsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsTransitGatewayInvitationsResponse) SetLinks(v []Link) {
	o.Links = &v
}

// GetTgwId returns the TgwId field value if set, zero value otherwise
func (o *StreamsTransitGatewayInvitationsResponse) GetTgwId() string {
	if o == nil || IsNil(o.TgwId) {
		var ret string
		return ret
	}
	return *o.TgwId
}

// GetTgwIdOk returns a tuple with the TgwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayInvitationsResponse) GetTgwIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwId) {
		return nil, false
	}

	return o.TgwId, true
}

// HasTgwId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayInvitationsResponse) HasTgwId() bool {
	if o != nil && !IsNil(o.TgwId) {
		return true
	}

	return false
}

// SetTgwId gets a reference to the given string and assigns it to the TgwId field.
func (o *StreamsTransitGatewayInvitationsResponse) SetTgwId(v string) {
	o.TgwId = &v
}

// GetTgwResourceShareArn returns the TgwResourceShareArn field value if set, zero value otherwise
func (o *StreamsTransitGatewayInvitationsResponse) GetTgwResourceShareArn() string {
	if o == nil || IsNil(o.TgwResourceShareArn) {
		var ret string
		return ret
	}
	return *o.TgwResourceShareArn
}

// GetTgwResourceShareArnOk returns a tuple with the TgwResourceShareArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayInvitationsResponse) GetTgwResourceShareArnOk() (*string, bool) {
	if o == nil || IsNil(o.TgwResourceShareArn) {
		return nil, false
	}

	return o.TgwResourceShareArn, true
}

// HasTgwResourceShareArn returns a boolean if a field has been set.
func (o *StreamsTransitGatewayInvitationsResponse) HasTgwResourceShareArn() bool {
	if o != nil && !IsNil(o.TgwResourceShareArn) {
		return true
	}

	return false
}

// SetTgwResourceShareArn gets a reference to the given string and assigns it to the TgwResourceShareArn field.
func (o *StreamsTransitGatewayInvitationsResponse) SetTgwResourceShareArn(v string) {
	o.TgwResourceShareArn = &v
}

// GetTgwResourceShareInvitationArn returns the TgwResourceShareInvitationArn field value if set, zero value otherwise
func (o *StreamsTransitGatewayInvitationsResponse) GetTgwResourceShareInvitationArn() string {
	if o == nil || IsNil(o.TgwResourceShareInvitationArn) {
		var ret string
		return ret
	}
	return *o.TgwResourceShareInvitationArn
}

// GetTgwResourceShareInvitationArnOk returns a tuple with the TgwResourceShareInvitationArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayInvitationsResponse) GetTgwResourceShareInvitationArnOk() (*string, bool) {
	if o == nil || IsNil(o.TgwResourceShareInvitationArn) {
		return nil, false
	}

	return o.TgwResourceShareInvitationArn, true
}

// HasTgwResourceShareInvitationArn returns a boolean if a field has been set.
func (o *StreamsTransitGatewayInvitationsResponse) HasTgwResourceShareInvitationArn() bool {
	if o != nil && !IsNil(o.TgwResourceShareInvitationArn) {
		return true
	}

	return false
}

// SetTgwResourceShareInvitationArn gets a reference to the given string and assigns it to the TgwResourceShareInvitationArn field.
func (o *StreamsTransitGatewayInvitationsResponse) SetTgwResourceShareInvitationArn(v string) {
	o.TgwResourceShareInvitationArn = &v
}
